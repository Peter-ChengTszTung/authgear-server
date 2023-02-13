package latte

import (
	"context"
	"errors"
	"fmt"

	"github.com/authgear/authgear-server/pkg/api"
	"github.com/authgear/authgear-server/pkg/api/apierrors"
	"github.com/authgear/authgear-server/pkg/api/model"
	"github.com/authgear/authgear-server/pkg/lib/authn/identity"
	"github.com/authgear/authgear-server/pkg/lib/workflow"
	"github.com/authgear/authgear-server/pkg/util/errorutil"
	"github.com/authgear/authgear-server/pkg/util/validation"
)

func init() {
	workflow.RegisterPrivateIntent(&IntentCreateLoginID{})
}

var IntentCreateLoginIDSchema = validation.NewSimpleSchema(`{}`)

type IntentCreateLoginID struct {
	UserID      string               `json:"user_id"`
	LoginIDType model.LoginIDKeyType `json:"login_id_type"`
	LoginIDKey  string               `json:"login_id_key"`
}

func (*IntentCreateLoginID) Kind() string {
	return "latte.IntentCreateLoginID"
}

func (*IntentCreateLoginID) JSONSchema() *validation.SimpleSchema {
	return IntentCreateLoginIDSchema
}

func (*IntentCreateLoginID) CanReactTo(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) ([]workflow.Input, error) {
	switch len(w.Nodes) {
	case 0:
		return []workflow.Input{
			&InputTakeLoginID{},
		}, nil
	case 1:
		return nil, nil
	case 2:
		return nil, nil
	case 3:
		return nil, nil
	default:
		return nil, workflow.ErrEOF
	}
}

func (i *IntentCreateLoginID) ReactTo(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow, input workflow.Input) (*workflow.Node, error) {
	var inputTakeLoginID inputTakeLoginID

	switch len(w.Nodes) {
	case 0:
		if workflow.AsInput(input, &inputTakeLoginID) {
			loginID := inputTakeLoginID.GetLoginID()
			spec := &identity.Spec{
				Type: model.IdentityTypeLoginID,
				LoginID: &identity.LoginIDSpec{
					Type:  i.LoginIDType,
					Key:   i.LoginIDKey,
					Value: loginID,
				},
			}

			info, err := deps.Identities.New(i.UserID, spec, identity.NewIdentityOptions{
				LoginIDEmailByPassBlocklistAllowlist: false,
			})
			if err != nil {
				return nil, err
			}

			duplicate, err := deps.Identities.CheckDuplicated(info)
			if err != nil && !errors.Is(err, identity.ErrIdentityAlreadyExists) {
				return nil, err
			}
			// Either err == nil, or err == ErrIdentityAlreadyExists and duplicate is non-nil.
			if err != nil {
				spec := info.ToSpec()
				otherSpec := duplicate.ToSpec()
				return nil, identityFillDetails(api.ErrDuplicatedIdentity, &spec, &otherSpec)
			}

			return workflow.NewNodeSimple(&NodeCreateIdentity{
				IdentityInfo: info,
			}), nil
		}
	case 1:
		iden := i.identityInfo(w)
		return workflow.NewNodeSimple(&NodePopulateStandardAttributes{
			IdentityInfo: iden,
		}), nil
	case 2:
		iden := i.identityInfo(w)
		return workflow.NewSubWorkflow(&IntentVerifyIdentity{
			IdentityInfo: iden,
			IsFromSignUp: true,
		}), nil
	case 3:
		iden := i.identityInfo(w)
		return workflow.NewSubWorkflow(&IntentCreateOOBOTPAuthenticatorForLoginID{
			IdentityInfo: iden,
		}), nil
	}

	return nil, workflow.ErrIncompatibleInput
}

func (i *IntentCreateLoginID) identityInfo(w *workflow.Workflow) *identity.Info {
	node, ok := workflow.FindSingleNode[*NodeCreateIdentity](w)
	if !ok {
		panic(fmt.Errorf("workflow: expected NodeCreateIdentity"))
	}

	return node.IdentityInfo
}

func (*IntentCreateLoginID) GetEffects(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) (effs []workflow.Effect, err error) {
	return nil, nil
}

func (*IntentCreateLoginID) OutputData(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) (interface{}, error) {
	return nil, nil
}

func identityFillDetails(err error, spec *identity.Spec, otherSpec *identity.Spec) error {
	details := errorutil.Details{}

	if spec != nil {
		details["IdentityTypeIncoming"] = apierrors.APIErrorDetail.Value(spec.Type)
		switch spec.Type {
		case model.IdentityTypeLoginID:
			details["LoginIDTypeIncoming"] = apierrors.APIErrorDetail.Value(spec.LoginID.Type)
		case model.IdentityTypeOAuth:
			details["OAuthProviderTypeIncoming"] = apierrors.APIErrorDetail.Value(spec.OAuth.ProviderID.Type)
		}
	}

	if otherSpec != nil {
		details["IdentityTypeExisting"] = apierrors.APIErrorDetail.Value(otherSpec.Type)
		switch otherSpec.Type {
		case model.IdentityTypeLoginID:
			details["LoginIDTypeExisting"] = apierrors.APIErrorDetail.Value(otherSpec.LoginID.Type)
		case model.IdentityTypeOAuth:
			details["OAuthProviderTypeExisting"] = apierrors.APIErrorDetail.Value(otherSpec.OAuth.ProviderID.Type)
		}
	}

	return errorutil.WithDetails(err, details)
}
