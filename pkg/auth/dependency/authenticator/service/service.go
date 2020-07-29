package service

import (
	"github.com/authgear/authgear-server/pkg/auth/dependency/authenticator"
	"github.com/authgear/authgear-server/pkg/auth/dependency/authenticator/oob"
	"github.com/authgear/authgear-server/pkg/auth/dependency/authenticator/password"
	"github.com/authgear/authgear-server/pkg/auth/dependency/authenticator/totp"
	"github.com/authgear/authgear-server/pkg/auth/dependency/identity"
	"github.com/authgear/authgear-server/pkg/core/authn"
)

type PasswordAuthenticatorProvider interface {
	Get(userID, id string) (*password.Authenticator, error)
	List(userID string) ([]*password.Authenticator, error)
	New(userID string, password string) (*password.Authenticator, error)
	// WithPassword returns new authenticator pointer if password is changed
	// Otherwise original authenticator will be returned
	WithPassword(a *password.Authenticator, password string) (*password.Authenticator, error)
	Create(*password.Authenticator) error
	UpdatePassword(*password.Authenticator) error
	Delete(*password.Authenticator) error
	Authenticate(a *password.Authenticator, password string) error
}

type TOTPAuthenticatorProvider interface {
	Get(userID, id string) (*totp.Authenticator, error)
	List(userID string) ([]*totp.Authenticator, error)
	New(userID string) *totp.Authenticator
	Create(*totp.Authenticator) error
	Delete(*totp.Authenticator) error
	Authenticate(candidates []*totp.Authenticator, code string) *totp.Authenticator
}

type OOBOTPAuthenticatorProvider interface {
	Get(userID, id string) (*oob.Authenticator, error)
	List(userID string) ([]*oob.Authenticator, error)
	New(userID string, channel authn.AuthenticatorOOBChannel, phone string, email string) *oob.Authenticator
	Create(*oob.Authenticator) error
	Delete(*oob.Authenticator) error
	Authenticate(secret string, channel authn.AuthenticatorOOBChannel, code string) error
}

type Service struct {
	Password PasswordAuthenticatorProvider
	TOTP     TOTPAuthenticatorProvider
	OOBOTP   OOBOTPAuthenticatorProvider
}

func (s *Service) Get(userID string, typ authn.AuthenticatorType, id string) (*authenticator.Info, error) {
	switch typ {
	case authn.AuthenticatorTypePassword:
		p, err := s.Password.Get(userID, id)
		if err != nil {
			return nil, err
		}
		return passwordToAuthenticatorInfo(p), nil

	case authn.AuthenticatorTypeTOTP:
		t, err := s.TOTP.Get(userID, id)
		if err != nil {
			return nil, err
		}
		return totpToAuthenticatorInfo(t), nil

	case authn.AuthenticatorTypeOOB:
		o, err := s.OOBOTP.Get(userID, id)
		if err != nil {
			return nil, err
		}
		return oobotpToAuthenticatorInfo(o), nil
	}

	panic("authenticator: unknown authenticator type " + typ)
}

func (s *Service) List(userID string, typ authn.AuthenticatorType) ([]*authenticator.Info, error) {
	var ais []*authenticator.Info
	switch typ {
	case authn.AuthenticatorTypePassword:
		as, err := s.Password.List(userID)
		if err != nil {
			return nil, err
		}
		for _, a := range as {
			ais = append(ais, passwordToAuthenticatorInfo(a))
		}

	case authn.AuthenticatorTypeTOTP:
		as, err := s.TOTP.List(userID)
		if err != nil {
			return nil, err
		}
		for _, a := range as {
			ais = append(ais, totpToAuthenticatorInfo(a))
		}

	case authn.AuthenticatorTypeOOB:
		as, err := s.OOBOTP.List(userID)
		if err != nil {
			return nil, err
		}
		for _, a := range as {
			ais = append(ais, oobotpToAuthenticatorInfo(a))
		}

	default:
		panic("authenticator: unknown authenticator type " + typ)
	}
	return ais, nil
}

func (s *Service) ListByIdentity(ii *identity.Info) (ais []*authenticator.Info, err error) {
	// This function takes IdentityInfo instead of IdentitySpec because
	// The login ID value in IdentityInfo is normalized.
	switch ii.Type {
	case authn.IdentityTypeOAuth:
		// OAuth Identity does not have associated authenticators.
		return
	case authn.IdentityTypeLoginID:
		// Login ID Identity has password, TOTP and OOB OTP.
		// Note that we only return OOB OTP associated with the login ID.
		var pas []*password.Authenticator
		pas, err = s.Password.List(ii.UserID)
		if err != nil {
			return
		}
		for _, pa := range pas {
			ais = append(ais, passwordToAuthenticatorInfo(pa))
		}

		var tas []*totp.Authenticator
		tas, err = s.TOTP.List(ii.UserID)
		if err != nil {
			return
		}
		for _, ta := range tas {
			ais = append(ais, totpToAuthenticatorInfo(ta))
		}

		loginID := ii.Claims[identity.IdentityClaimLoginIDValue]
		var oas []*oob.Authenticator
		oas, err = s.OOBOTP.List(ii.UserID)
		if err != nil {
			return
		}
		for _, oa := range oas {
			if oa.Email == loginID || oa.Phone == loginID {
				ais = append(ais, oobotpToAuthenticatorInfo(oa))
			}
		}
	case authn.IdentityTypeAnonymous:
		// Anonymous Identity does not have associated authenticators.
		return
	default:
		panic("v: unknown identity type " + ii.Type)
	}

	return
}

func (s *Service) New(spec *authenticator.Spec, secret string) (*authenticator.Info, error) {
	switch spec.Type {
	case authn.AuthenticatorTypePassword:
		p, err := s.Password.New(spec.UserID, secret)
		if err != nil {
			return nil, err
		}
		return passwordToAuthenticatorInfo(p), nil

	case authn.AuthenticatorTypeTOTP:
		t := s.TOTP.New(spec.UserID)
		return totpToAuthenticatorInfo(t), nil

	case authn.AuthenticatorTypeOOB:
		channel := spec.Props[authenticator.AuthenticatorPropOOBOTPChannelType].(string)
		var phone, email string
		switch authn.AuthenticatorOOBChannel(channel) {
		case authn.AuthenticatorOOBChannelSMS:
			phone = spec.Props[authenticator.AuthenticatorPropOOBOTPPhone].(string)
		case authn.AuthenticatorOOBChannelEmail:
			email = spec.Props[authenticator.AuthenticatorPropOOBOTPEmail].(string)
		}
		o := s.OOBOTP.New(spec.UserID, authn.AuthenticatorOOBChannel(channel), phone, email)
		return oobotpToAuthenticatorInfo(o), nil
	}

	panic("authenticator: unknown authenticator type " + spec.Type)
}

func (s *Service) WithSecret(ai *authenticator.Info, secret string) (bool, *authenticator.Info, error) {
	changed := false
	switch ai.Type {
	case authn.AuthenticatorTypePassword:
		a := passwordFromAuthenticatorInfo(ai)
		newAuth, err := s.Password.WithPassword(a, secret)
		if err != nil {
			return false, nil, err
		}
		changed = (newAuth != a)
		return changed, passwordToAuthenticatorInfo(newAuth), nil
	}

	panic("authenticator: update authenticator is not supported for type " + ai.Type)
}

func (s *Service) Create(info *authenticator.Info) error {
	switch info.Type {
	case authn.AuthenticatorTypePassword:
		a := passwordFromAuthenticatorInfo(info)
		if err := s.Password.Create(a); err != nil {
			return err
		}

	case authn.AuthenticatorTypeTOTP:
		a := totpFromAuthenticatorInfo(info)
		if err := s.TOTP.Create(a); err != nil {
			return err
		}

	case authn.AuthenticatorTypeOOB:
		a := oobotpFromAuthenticatorInfo(info)
		if err := s.OOBOTP.Create(a); err != nil {
			return err
		}

	default:
		panic("authenticator: unknown authenticator type " + info.Type)
	}

	return nil
}

func (s *Service) Update(info *authenticator.Info) error {
	switch info.Type {
	case authn.AuthenticatorTypePassword:
		a := passwordFromAuthenticatorInfo(info)
		if err := s.Password.UpdatePassword(a); err != nil {
			return err
		}
	default:
		panic("authenticator: unknown authenticator type for update" + info.Type)
	}

	return nil
}

func (s *Service) Delete(info *authenticator.Info) error {
	switch info.Type {
	case authn.AuthenticatorTypePassword:
		a := passwordFromAuthenticatorInfo(info)
		if err := s.Password.Delete(a); err != nil {
			return err
		}

	case authn.AuthenticatorTypeTOTP:
		a := totpFromAuthenticatorInfo(info)
		if err := s.TOTP.Delete(a); err != nil {
			return err
		}

	case authn.AuthenticatorTypeOOB:
		a := oobotpFromAuthenticatorInfo(info)
		if err := s.OOBOTP.Delete(a); err != nil {
			return err
		}
	default:
		panic("authenticator: delete authenticator is not supported yet for type " + info.Type)
	}

	return nil
}

func (s *Service) Authenticate(spec *authenticator.Spec, state map[string]string, secret string) (*authenticator.Info, error) {
	switch spec.Type {
	case authn.AuthenticatorTypePassword:
		ps, err := s.Password.List(spec.UserID)
		if err != nil {
			return nil, err
		}
		if len(ps) != 1 {
			return nil, authenticator.ErrAuthenticatorNotFound
		}

		if s.Password.Authenticate(ps[0], secret) != nil {
			return nil, authenticator.ErrInvalidCredentials
		}
		return passwordToAuthenticatorInfo(ps[0]), nil

	case authn.AuthenticatorTypeTOTP:
		ts, err := s.TOTP.List(spec.UserID)
		if err != nil {
			return nil, err
		}

		t := s.TOTP.Authenticate(ts, secret)
		if t == nil {
			return nil, authenticator.ErrInvalidCredentials
		}
		return totpToAuthenticatorInfo(t), nil

	case authn.AuthenticatorTypeOOB:
		panic("authenticator: unsupported OOB authenticator")
	}

	panic("authenticator: unknown authenticator type " + spec.Type)
}

func (s *Service) VerifySecret(info *authenticator.Info, state map[string]string, secret string) error {
	switch info.Type {
	case authn.AuthenticatorTypePassword:
		a := passwordFromAuthenticatorInfo(info)
		if s.Password.Authenticate(a, secret) != nil {
			return authenticator.ErrInvalidCredentials
		}
		return nil

	case authn.AuthenticatorTypeTOTP:
		a := totpFromAuthenticatorInfo(info)
		if s.TOTP.Authenticate([]*totp.Authenticator{a}, secret) != nil {
			return authenticator.ErrInvalidCredentials
		}
		return nil

	case authn.AuthenticatorTypeOOB:
		a := oobotpFromAuthenticatorInfo(info)
		otpSecret := state[authenticator.AuthenticatorStateOOBOTPSecret]
		if s.OOBOTP.Authenticate(otpSecret, a.Channel, secret) != nil {
			return authenticator.ErrInvalidCredentials
		}
		return nil
	}

	panic("authenticator: unhandled authenticator type " + info.Type)
}
