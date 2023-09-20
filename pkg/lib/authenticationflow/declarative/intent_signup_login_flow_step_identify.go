package declarative

import (
	"context"
	"fmt"

	"github.com/iawaknahc/jsonschema/pkg/jsonpointer"

	authflow "github.com/authgear/authgear-server/pkg/lib/authenticationflow"
	"github.com/authgear/authgear-server/pkg/lib/config"
)

func init() {
	authflow.RegisterIntent(&IntentSignupLoginFlowStepIdentify{})
}

type IntentSignupLoginFlowStepIdentifyData struct {
	Candidates []IdentificationCandidate `json:"candidates"`
}

var _ authflow.Data = IntentSignupLoginFlowStepIdentifyData{}

func (IntentSignupLoginFlowStepIdentifyData) Data() {}

type IntentSignupLoginFlowStepIdentify struct {
	JSONPointer jsonpointer.T             `json:"json_pointer,omitempty"`
	StepName    string                    `json:"step_name,omitempty"`
	Candidates  []IdentificationCandidate `json:"candidates"`
}

var _ authflow.Intent = &IntentSignupLoginFlowStepIdentify{}
var _ authflow.DataOutputer = &IntentSignupLoginFlowStepIdentify{}

func NewIntentSignupLoginFlowStepIdentify(ctx context.Context, deps *authflow.Dependencies, i *IntentSignupLoginFlowStepIdentify) (*IntentSignupLoginFlowStepIdentify, error) {
	current, err := authflow.FlowObject(authflow.GetFlowRootObject(ctx), i.JSONPointer)
	if err != nil {
		return nil, err
	}
	step := i.step(current)

	candidates := []IdentificationCandidate{}
	for _, b := range step.OneOf {
		switch b.Identification {
		case config.AuthenticationFlowIdentificationEmail:
			fallthrough
		case config.AuthenticationFlowIdentificationPhone:
			fallthrough
		case config.AuthenticationFlowIdentificationUsername:
			c := NewIdentificationCandidateLoginID(b.Identification)
			candidates = append(candidates, c)
		case config.AuthenticationFlowIdentificationOAuth:
			oauthCandidates := NewIdentificationCandidatesOAuth(
				deps.Config.Identity.OAuth,
				deps.FeatureConfig.Identity.OAuth.Providers,
			)
			candidates = append(candidates, oauthCandidates...)
		case config.AuthenticationFlowIdentificationPasskey:
			// Passkey is for login only.
			requestOptions, err := deps.PasskeyRequestOptionsService.MakeModalRequestOptions()
			if err != nil {
				return nil, err
			}
			c := NewIdentificationCandidatePasskey(requestOptions)
			candidates = append(candidates, c)
		}
	}

	i.Candidates = candidates
	return i, nil
}

func (*IntentSignupLoginFlowStepIdentify) Kind() string {
	return "IntentSignupLoginFlowStepIdentify"
}

func (i *IntentSignupLoginFlowStepIdentify) CanReactTo(ctx context.Context, deps *authflow.Dependencies, flows authflow.Flows) (authflow.InputSchema, error) {
	// Let the input to select which identification method to use.
	if len(flows.Nearest.Nodes) == 0 {
		return &InputSchemaStepIdentify{
			JSONPointer: i.JSONPointer,
			Candidates:  i.Candidates,
		}, nil
	}

	return nil, authflow.ErrEOF
}

func (i *IntentSignupLoginFlowStepIdentify) ReactTo(ctx context.Context, deps *authflow.Dependencies, flows authflow.Flows, input authflow.Input) (*authflow.Node, error) {
	current, err := authflow.FlowObject(authflow.GetFlowRootObject(ctx), i.JSONPointer)
	if err != nil {
		return nil, err
	}
	step := i.step(current)

	if len(flows.Nearest.Nodes) == 0 {
		var inputTakeIdentificationMethod inputTakeIdentificationMethod
		if authflow.AsInput(input, &inputTakeIdentificationMethod) {
			identification := inputTakeIdentificationMethod.GetIdentificationMethod()
			_, err := i.checkIdentificationMethod(deps, step, identification)
			if err != nil {
				return nil, err
			}

			switch identification {
			case config.AuthenticationFlowIdentificationEmail:
				fallthrough
			case config.AuthenticationFlowIdentificationPhone:
				fallthrough
			case config.AuthenticationFlowIdentificationUsername:
				break
				//return authflow.NewNodeSimple(&NodeLookupIdentityLoginID{
				//	JSONPointer:    authflow.JSONPointerForOneOf(i.JSONPointer, idx),
				//	Identification: identification,
				//}), nil
			case config.AuthenticationFlowIdentificationOAuth:
				break
				//return authflow.NewSubFlow(&IntentLookupIdentityOAuth{
				//	JSONPointer:    authflow.JSONPointerForOneOf(i.JSONPointer, idx),
				//	Identification: identification,
				//}), nil
			case config.AuthenticationFlowIdentificationPasskey:
				break
				//return authflow.NewNodeSimple(&NodeLookupIdentityPasskey{
				//	JSONPointer:    authflow.JSONPointerForOneOf(i.JSONPointer, idx),
				//	Identification: identification,
				//}), nil
			}
		}
	}

	return nil, authflow.ErrIncompatibleInput
}

func (i *IntentSignupLoginFlowStepIdentify) OutputData(ctx context.Context, deps *authflow.Dependencies, flows authflow.Flows) (authflow.Data, error) {
	return IntentSignupLoginFlowStepIdentifyData{
		Candidates: i.Candidates,
	}, nil
}

func (i *IntentSignupLoginFlowStepIdentify) checkIdentificationMethod(deps *authflow.Dependencies, step *config.AuthenticationFlowSignupLoginFlowStep, im config.AuthenticationFlowIdentification) (idx int, err error) {
	idx = -1

	for index, branch := range step.OneOf {
		branch := branch
		if im == branch.Identification {
			idx = index
		}
	}

	if idx >= 0 {
		return
	}

	err = authflow.ErrIncompatibleInput
	return
}

func (i *IntentSignupLoginFlowStepIdentify) step(o config.AuthenticationFlowObject) *config.AuthenticationFlowSignupLoginFlowStep {
	step, ok := o.(*config.AuthenticationFlowSignupLoginFlowStep)
	if !ok {
		panic(fmt.Errorf("flow object is %T", o))
	}

	return step
}
