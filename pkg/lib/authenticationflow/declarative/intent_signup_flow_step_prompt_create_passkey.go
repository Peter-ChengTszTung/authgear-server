package declarative

import (
	"context"

	"github.com/iawaknahc/jsonschema/pkg/jsonpointer"

	"github.com/authgear/authgear-server/pkg/api/model"
	authflow "github.com/authgear/authgear-server/pkg/lib/authenticationflow"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator"
	"github.com/authgear/authgear-server/pkg/lib/config"
)

func init() {
	authflow.RegisterIntent(&IntentSignupFlowStepPromptCreatePasskey{})
}

type IntentSignupFlowStepPromptCreatePasskey struct {
	StepID      string        `json:"step_id,omitempty"`
	JSONPointer jsonpointer.T `json:"json_pointer,omitempty"`
	UserID      string        `json:"user_id,omitempty"`
}

var _ FlowStep = &IntentSignupFlowStepPromptCreatePasskey{}

func (i *IntentSignupFlowStepPromptCreatePasskey) GetID() string {
	return i.StepID
}

func (i *IntentSignupFlowStepPromptCreatePasskey) GetJSONPointer() jsonpointer.T {
	return i.JSONPointer
}

var _ authflow.Intent = &IntentSignupFlowStepPromptCreatePasskey{}

func (*IntentSignupFlowStepPromptCreatePasskey) Kind() string {
	return "IntentSignupFlowStepPromptCreatePasskey"
}

func (i *IntentSignupFlowStepPromptCreatePasskey) CanReactTo(ctx context.Context, deps *authflow.Dependencies, flows authflow.Flows) (authflow.InputSchema, error) {
	// Find out whether we need to prompt.
	if len(flows.Nearest.Nodes) == 0 {
		return nil, nil
	}

	return nil, authflow.ErrEOF
}

func (i *IntentSignupFlowStepPromptCreatePasskey) ReactTo(ctx context.Context, deps *authflow.Dependencies, flows authflow.Flows, _ authflow.Input) (*authflow.Node, error) {
	// See if any used identification can use passkey.
	passkeyCanBeUsed := false
	milestones := authflow.FindAllMilestones[MilestoneIdentificationMethod](flows.Root)
	for _, m := range milestones {
		i := m.MilestoneIdentificationMethod()
		for _, a := range i.PrimaryAuthentications() {
			if a == config.AuthenticationFlowAuthenticationPrimaryPasskey {
				passkeyCanBeUsed = true
			}
		}
	}

	// No identification used can use passkey.
	// Do not prompt.
	if !passkeyCanBeUsed {
		return authflow.NewNodeSimple(&NodeSentinel{}), nil
	}

	ais, err := deps.Authenticators.List(
		i.UserID,
		authenticator.KeepKind(authenticator.KindPrimary),
		authenticator.KeepType(model.AuthenticatorTypePasskey),
	)
	if err != nil {
		return nil, err
	}

	// The user has at least one passkey. Do not need to prompt them.
	if len(ais) > 0 {
		return authflow.NewNodeSimple(&NodeSentinel{}), nil
	}

	// Otherwise it is OK to prompt.
	n, err := NewNodePromptCreatePasskey(deps, &NodePromptCreatePasskey{
		JSONPointer: i.JSONPointer,
		UserID:      i.UserID,
	})
	if err != nil {
		return nil, err
	}

	return authflow.NewNodeSimple(n), nil
}