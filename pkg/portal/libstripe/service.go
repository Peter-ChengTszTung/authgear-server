package libstripe

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	relay "github.com/authgear/graphql-go-relay"
	goredis "github.com/go-redis/redis/v8"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
	"github.com/stripe/stripe-go/v72/webhook"

	"github.com/authgear/authgear-server/pkg/lib/infra/redis/globalredis"
	portalconfig "github.com/authgear/authgear-server/pkg/portal/config"
	"github.com/authgear/authgear-server/pkg/portal/model"
	"github.com/authgear/authgear-server/pkg/util/clock"
	"github.com/authgear/authgear-server/pkg/util/duration"
	"github.com/authgear/authgear-server/pkg/util/log"
	"github.com/authgear/authgear-server/pkg/util/redisutil"
	"github.com/authgear/authgear-server/pkg/util/timeutil"
)

const RedisCacheKeySubscriptionPlans = "cache:portal:subscription-plans"

type Logger struct{ *log.Logger }

func NewLogger(lf *log.Factory) Logger { return Logger{lf.New("stripe")} }

func NewClientAPI(stripeConfig *portalconfig.StripeConfig, logger Logger) *client.API {
	clientAPI := &client.API{}
	clientAPI.Init(stripeConfig.SecretKey, &stripe.Backends{
		API: stripe.GetBackendWithConfig(stripe.APIBackend, &stripe.BackendConfig{
			LeveledLogger: logger,
		}),
	})
	return clientAPI
}

type PlanService interface {
	ListPlans() ([]*model.Plan, error)
}

type Cache interface {
	Get(context.Context, redisutil.SimpleCmdable, redisutil.Item) ([]byte, error)
}

type EndpointsProvider interface {
	BillingEndpointURL(relayGlobalAppID string) *url.URL
	BillingRedirectEndpointURL(relayGlobalAppID string) *url.URL
}

type Service struct {
	ClientAPI         *client.API
	Logger            Logger
	Context           context.Context
	Plans             PlanService
	GlobalRedisHandle *globalredis.Handle
	Cache             Cache
	Clock             clock.Clock
	StripeConfig      *portalconfig.StripeConfig
	Endpoints         EndpointsProvider
}

func (s *Service) FetchSubscriptionPlans() (subscriptionPlans []*SubscriptionPlan, err error) {
	item := redisutil.Item{
		Key:        RedisCacheKeySubscriptionPlans,
		Expiration: duration.PerHour,
		Do:         s.fetchSubscriptionPlans,
	}

	err = s.GlobalRedisHandle.WithConn(func(conn *goredis.Conn) error {
		bytes, err := s.Cache.Get(s.Context, conn, item)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &subscriptionPlans)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}

	return
}

func (s *Service) CreateCheckoutSession(appID string, customerEmail string, subscriptionPlan *SubscriptionPlan) (*CheckoutSession, error) {
	relayGlobalAppID := relay.ToGlobalID("App", appID)
	billingPageURL := s.Endpoints.BillingEndpointURL(relayGlobalAppID).String()
	billingRedirectPageURL := s.Endpoints.BillingRedirectEndpointURL(relayGlobalAppID).String()
	successURL := fmt.Sprintf("%s?session_id={CHECKOUT_SESSION_ID}", billingRedirectPageURL)
	cancelURL := billingPageURL

	params := &stripe.CheckoutSessionParams{
		Params: stripe.Params{
			Context: s.Context,
			Metadata: map[string]string{
				MetadataKeyAppID:    appID,
				MetadataKeyPlanName: subscriptionPlan.Name,
			},
		},
		SuccessURL:         &successURL,
		CancelURL:          &cancelURL,
		Mode:               stripe.String(string(stripe.CheckoutSessionModeSetup)),
		PaymentMethodTypes: []*string{stripe.String(string(stripe.PaymentMethodTypeCard))},
		CustomerCreation:   stripe.String(string(stripe.CheckoutSessionCustomerCreationAlways)),
	}

	if customerEmail != "" {
		// If the customer email is empty
		// The customer will be asked to enter their email address during the checkout process
		params.CustomerEmail = &customerEmail
	}

	checkoutSession, err := s.ClientAPI.CheckoutSessions.New(params)
	if err != nil {
		return nil, err
	}

	return NewCheckoutSession(checkoutSession), nil
}

func (s *Service) GetSubscriptionPlan(planName string) (*SubscriptionPlan, error) {
	subscriptionPlans, err := s.FetchSubscriptionPlans()
	if err != nil {
		return nil, err
	}
	return s.getStripeSubscription(planName, subscriptionPlans)
}

func (s *Service) FetchCheckoutSession(checkoutSessionID string) (*CheckoutSession, error) {
	checkoutSession, err := s.ClientAPI.CheckoutSessions.Get(checkoutSessionID, &stripe.CheckoutSessionParams{
		Params: stripe.Params{
			Context: s.Context,
		},
	})
	if err != nil {
		return nil, err
	}

	return NewCheckoutSession(checkoutSession), nil
}

func (s *Service) ConstructEvent(r *http.Request) (Event, error) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	sig := r.Header.Get("Stripe-Signature")
	stripeEvent, err := webhook.ConstructEvent(payload, sig, s.StripeConfig.WebhookSigningKey)
	if err != nil {
		return nil, err
	}

	event, err := s.constructEvent(&stripeEvent)
	if errors.Is(err, ErrUnknownEvent) {
		s.Logger.WithField("payload", string(payload)).Info("unhandled event")
	}
	return event, err
}

func (s *Service) CreateSubscriptionIfNotExists(checkoutSessionID string, subscriptionPlans []*SubscriptionPlan) error {
	// Fetch the checkout session
	expandSetupIntentPaymentMethod := "setup_intent.payment_method"
	expandCustomerSubscriptions := "customer.subscriptions"
	checkoutSession, err := s.ClientAPI.CheckoutSessions.Get(checkoutSessionID, &stripe.CheckoutSessionParams{
		Params: stripe.Params{
			Context: s.Context,
			Expand:  []*string{&expandSetupIntentPaymentMethod, &expandCustomerSubscriptions},
		},
	})
	if err != nil {
		return err
	}

	planName := checkoutSession.Metadata[MetadataKeyPlanName]
	appID := checkoutSession.Metadata[MetadataKeyAppID]

	// Find the subscription plan
	subscriptionPlan, err := s.getStripeSubscription(planName, subscriptionPlans)
	if err != nil {
		return err
	}

	// Update invoice settings default
	customerID := &checkoutSession.Customer.ID
	pm := checkoutSession.SetupIntent.PaymentMethod
	customerParams := &stripe.CustomerParams{
		Params: stripe.Params{
			Context: s.Context,
		},
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(pm.ID),
		},
	}

	_, err = s.ClientAPI.Customers.Update(*customerID, customerParams)
	if err != nil {
		return fmt.Errorf("failed to update customer default payment method: %w", err)
	}

	// Check if the custom has subscription to avoid duplicate subscription
	if checkoutSession.Customer.Subscriptions != nil && len(checkoutSession.Customer.Subscriptions.Data) > 0 {
		return ErrCustomerAlreadySubscribed
	}

	// Check if the app has subscription to avoid duplicate subscription
	hasSubscription := false
	iter := s.ClientAPI.Subscriptions.Search(&stripe.SubscriptionSearchParams{
		SearchParams: stripe.SearchParams{
			Context: s.Context,
			Query:   fmt.Sprintf("status:'active' AND metadata['app_id']: '%s'", appID),
		},
	})
	for iter.Next() {
		hasSubscription = true
	}
	if err := iter.Err(); err != nil {
		return fmt.Errorf("failed to search app's subscription: %w", err)
	}
	if hasSubscription {
		return ErrAppAlreadySubscribed
	}

	// Create subscription
	subscriptionItems := []*stripe.SubscriptionItemsParams{}
	for _, p := range subscriptionPlan.Prices {
		subscriptionItems = append(subscriptionItems, &stripe.SubscriptionItemsParams{
			Price: stripe.String(p.StripePriceID),
		})
	}

	billingCycleAnchor := s.Clock.NowUTC().AddDate(0, 1, 0)
	billingCycleAnchor = timeutil.FirstDayOfTheMonth(billingCycleAnchor)
	billingCycleAnchorUnix := billingCycleAnchor.Unix()
	_, err = s.ClientAPI.Subscriptions.New(&stripe.SubscriptionParams{
		Params: stripe.Params{
			Context: s.Context,
			Metadata: map[string]string{
				MetadataKeyAppID:    appID,
				MetadataKeyPlanName: planName,
			},
		},
		Customer:           customerID,
		Items:              subscriptionItems,
		BillingCycleAnchor: &billingCycleAnchorUnix,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) fetchSubscriptionPlans() ([]byte, error) {
	plans, err := s.Plans.ListPlans()
	if err != nil {
		return nil, err
	}

	products, err := s.fetchProducts()
	if err != nil {
		return nil, err
	}
	subscriptionPlans, err := s.convertToSubscriptionPlans(plans, products)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(subscriptionPlans)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (s *Service) fetchProducts() ([]*stripe.Product, error) {
	var products []*stripe.Product

	expandDefaultPrice := "data.default_price"
	listProductParams := &stripe.ProductListParams{
		ListParams: stripe.ListParams{
			Context: s.Context,
			Expand:  []*string{&expandDefaultPrice},
		},
		Active: stripe.Bool(true),
	}
	iter := s.ClientAPI.Products.List(listProductParams)
	for iter.Next() {
		product := iter.Current().(*stripe.Product)
		products = append(products, product)
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Service) convertToSubscriptionPlans(plans []*model.Plan, products []*stripe.Product) ([]*SubscriptionPlan, error) {
	knownPlanNames := make(map[string]struct{})
	for _, plan := range plans {
		knownPlanNames[plan.Name] = struct{}{}
	}

	m := make(map[string]*SubscriptionPlan)
	usagePrices := []*Price{}
	for _, product := range products {
		price, err := NewPrice(product)
		if err != nil {
			// skip the unknown product
			continue
		}
		switch price.Type {
		case model.PriceTypeFixed:
			// New SubscriptionPlan for the fixed price products
			planName := product.Metadata[MetadataKeyPlanName]
			// There could exist some unknown Products on Stripe.
			// We tolerate that.
			_, ok := knownPlanNames[planName]
			if !ok {
				continue
			}
			// If there are multiple fixed price products have the same plan name
			// Add the price to the same SubscriptionPlan
			if _, exists := m[planName]; !exists {
				m[planName] = NewSubscriptionPlan(planName)
			}
			m[planName].Prices = append(m[planName].Prices, price)
		case model.PriceTypeUsage:
			usagePrices = append(usagePrices, price)
		}
	}

	var out []*SubscriptionPlan
	for _, subscriptionPlan := range m {
		// Add usage prices to all subscription plans
		subscriptionPlan.Prices = append(subscriptionPlan.Prices, usagePrices...)
		out = append(out, subscriptionPlan)
	}

	return out, nil
}

func (s *Service) getStripeSubscription(planName string, subscriptionPlans []*SubscriptionPlan) (*SubscriptionPlan, error) {
	var subscriptionPlan *SubscriptionPlan
	for _, sp := range subscriptionPlans {
		if sp.Name == planName {
			subscriptionPlan = sp
			break
		}
	}
	if subscriptionPlan == nil {
		return nil, fmt.Errorf("subscription plan not found")
	}

	return subscriptionPlan, nil
}

func (s *Service) constructEvent(stripeEvent *stripe.Event) (Event, error) {
	switch stripeEvent.Type {
	case string(EventTypeCheckoutSessionCompleted):
		object := stripeEvent.Data.Object
		checkoutSessionID, ok := object["id"].(string)
		if !ok {
			return nil, ErrUnknownEvent
		}
		customerID, ok := object["customer"].(string)
		if !ok {
			return nil, ErrUnknownEvent
		}
		metadata, ok := object["metadata"].(map[string]interface{})
		if !ok {
			return nil, ErrUnknownEvent
		}
		appID, ok := metadata[MetadataKeyAppID].(string)
		if !ok {
			return nil, ErrUnknownEvent
		}
		planName, ok := metadata[MetadataKeyPlanName].(string)
		if !ok {
			return nil, ErrUnknownEvent
		}
		return &CheckoutSessionCompletedEvent{
			AppID:                   appID,
			PlanName:                planName,
			StripeCheckoutSessionID: checkoutSessionID,
			StripeCustomerID:        customerID,
		}, nil
	case string(EventTypeCustomerSubscriptionCreated),
		string(EventTypeCustomerSubscriptionUpdated):
		object := stripeEvent.Data.Object
		subscriptionID, ok := object["id"].(string)
		if !ok {
			return nil, ErrUnknownEvent
		}

		subscriptionStatus, ok := object["status"].(string)
		if !ok {
			return nil, ErrUnknownEvent
		}
		customerID, ok := object["customer"].(string)
		if !ok {
			return nil, ErrUnknownEvent
		}
		metadata, ok := object["metadata"].(map[string]interface{})
		if !ok {
			return nil, ErrUnknownEvent
		}
		appID, ok := metadata[MetadataKeyAppID].(string)
		if !ok {
			return nil, ErrUnknownEvent
		}
		planName, ok := metadata[MetadataKeyPlanName].(string)
		if !ok {
			return nil, ErrUnknownEvent
		}
		if stripeEvent.Type == string(EventTypeCustomerSubscriptionCreated) {
			return &CustomerSubscriptionCreatedEvent{
				&CustomerSubscriptionEvent{
					StripeSubscriptionID:     subscriptionID,
					StripeCustomerID:         customerID,
					AppID:                    appID,
					PlanName:                 planName,
					StripeSubscriptionStatus: stripe.SubscriptionStatus(subscriptionStatus),
				},
			}, nil
		}

		return &CustomerSubscriptionUpdatedEvent{
			&CustomerSubscriptionEvent{
				StripeSubscriptionID:     subscriptionID,
				StripeCustomerID:         customerID,
				AppID:                    appID,
				PlanName:                 planName,
				StripeSubscriptionStatus: stripe.SubscriptionStatus(subscriptionStatus),
			},
		}, nil
	default:
		return nil, ErrUnknownEvent
	}
}

func (s *Service) GenerateCustomerPortalSession(appID string, customerID string) (*stripe.BillingPortalSession, error) {
	u := s.Endpoints.BillingEndpointURL(relay.ToGlobalID("App", appID))

	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(customerID),
		ReturnURL: stripe.String(u.String()),
	}

	return s.ClientAPI.BillingPortalSessions.New(params)
}