package webapp

import (
	"net/http"

	"github.com/authgear/authgear-server/pkg/auth/config"
	"github.com/authgear/authgear-server/pkg/auth/dependency/auth"
	"github.com/authgear/authgear-server/pkg/auth/dependency/identity/loginid"
	interactionflows "github.com/authgear/authgear-server/pkg/auth/dependency/interaction/flows"
	"github.com/authgear/authgear-server/pkg/db"
	"github.com/authgear/authgear-server/pkg/httproute"
	"github.com/authgear/authgear-server/pkg/template"
	"github.com/authgear/authgear-server/pkg/validation"
)

const (
	TemplateItemTypeAuthUIEnterLoginIDHTML config.TemplateItemType = "auth_ui_enter_login_id.html"
)

var TemplateAuthUIEnterLoginIDHTML = template.Spec{
	Type:        TemplateItemTypeAuthUIEnterLoginIDHTML,
	IsHTML:      true,
	Translation: TemplateItemTypeAuthUITranslationJSON,
	Defines:     defines,
	Components:  components,
	Default: `<!DOCTYPE html>
<html>
{{ template "auth_ui_html_head.html" . }}
<body class="page">
<div class="content">

{{ template "auth_ui_header.html" . }}

<div class="simple-form vertical-form form-fields-container">

<div class="nav-bar">
	<button class="btn back-btn" type="button" title="{{ localize "back-button-title" }}"></button>
</div>

<div class="title primary-txt">
	{{ if $.OldLoginIDValue }}
	{{ localize "enter-login-id-page-title--change" $.LoginIDKey }}
	{{ else }}
	{{ localize "enter-login-id-page-title--add" $.LoginIDKey }}
	{{ end }}
</div>

{{ template "ERROR" . }}

<form class="vertical-form form-fields-container" method="post" novalidate>

{{ $.CSRFField }}

{{ if eq .LoginIDInputType "phone" }}
<div class="phone-input">
	<select class="input select primary-txt" name="x_calling_code">
		{{ range $.CountryCallingCodes }}
		<option
			value="{{ . }}"
			{{ if $.x_calling_code }}{{ if eq $.x_calling_code . }}
			selected
			{{ end }}{{ end }}
			>
			+{{ . }}
		</option>
		{{ end }}
	</select>
	<input class="input text-input primary-txt" type="text" inputmode="numeric" pattern="[0-9]*" name="x_national_number" placeholder="{{ localize "phone-number-placeholder" }}">
</div>
{{ else }}
<input class="input text-input primary-txt" type="{{ .LoginIDInputType }}" name="x_login_id" placeholder="{{ localize "login-id-placeholder" .LoginIDType }}">
{{ end }}

<button class="btn primary-btn align-self-flex-end" type="submit" name="x_action" value="add_or_update">{{ localize "next-button-label" }}</button>

</form>

{{ if .OldLoginIDValue }}
<form class="enter-login-id-remove-form" method="post" novalidate>
{{ $.CSRFField }}
<button class="anchor" type="submit" name="x_action" value="remove">{{ localize "disconnect-button-label" }}</button>
{{ end }}
</form>

</div>
{{ template "auth_ui_footer.html" . }}

</div>
</body>
</html>
`,
}

type EnterLoginIDViewModel struct {
	LoginIDKey       string
	LoginIDType      string
	OldLoginIDValue  string
	LoginIDInputType string
}

func NewEnterLoginIDViewModel(state *interactionflows.State) EnterLoginIDViewModel {
	loginIDKey, _ := state.Extra[interactionflows.ExtraLoginIDKey].(string)
	loginIDType, _ := state.Extra[interactionflows.ExtraLoginIDType].(string)
	loginIDInputType, _ := state.Extra[interactionflows.ExtraLoginIDInputType].(string)
	oldLoginIDValue, _ := state.Extra[interactionflows.ExtraOldLoginID].(string)

	return EnterLoginIDViewModel{
		LoginIDKey:       loginIDKey,
		LoginIDType:      loginIDType,
		LoginIDInputType: loginIDInputType,
		OldLoginIDValue:  oldLoginIDValue,
	}
}

const RemoveLoginIDRequest = "RemoveLoginIDRequest"

var EnterLoginIDSchema = validation.NewMultipartSchema("").
	Add(RemoveLoginIDRequest, `
		{
			"type": "object",
			"properties": {
				"x_login_id_key": { "type": "string" },
				"x_old_login_id_value": { "type": "string" }
			},
			"required": ["x_login_id_key", "x_old_login_id_value"]
		}
	`).Instantiate()

func ConfigureEnterLoginIDRoute(route httproute.Route) httproute.Route {
	return route.
		WithMethods("OPTIONS", "POST", "GET").
		WithPathPattern("/enter_login_id")
}

type EnterLoginIDInteractions interface {
	RemoveLoginID(state *interactionflows.State, userID string, loginID loginid.LoginID) (*interactionflows.WebAppResult, error)
	UpdateLoginID(state *interactionflows.State, userID string, oldLoginID loginid.LoginID, newLoginID loginid.LoginID) (*interactionflows.WebAppResult, error)
	AddLoginID(state *interactionflows.State, userID string, loginID loginid.LoginID) (*interactionflows.WebAppResult, error)
}

type EnterLoginIDHandler struct {
	Database      *db.Handle
	State         StateService
	BaseViewModel *BaseViewModeler
	Renderer      Renderer
	Interactions  EnterLoginIDInteractions
	Responder     Responder
}

func (h *EnterLoginIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "GET" {
		state, err := h.State.RestoreReadOnlyState(r, false)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		baseViewModel := h.BaseViewModel.ViewModel(r, state.Error)
		enterLoginIDViewModel := NewEnterLoginIDViewModel(state)

		data := map[string]interface{}{}

		Embed(data, baseViewModel)
		Embed(data, enterLoginIDViewModel)

		h.Renderer.Render(w, r, TemplateItemTypeAuthUIEnterLoginIDHTML, data)
		return
	}

	if r.Method == "POST" && r.Form.Get("x_action") == "remove" {
		h.Database.WithTx(func() error {
			state, err := h.State.CloneState(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return err
			}

			var result *interactionflows.WebAppResult
			defer func() {
				h.State.UpdateState(state, result, err)
				h.Responder.Respond(w, r, state, result, err)
			}()

			enterLoginIDViewModel := NewEnterLoginIDViewModel(state)

			userID := auth.GetSession(r.Context()).AuthnAttrs().UserID

			result, err = h.Interactions.RemoveLoginID(state, userID, loginid.LoginID{
				Key:   enterLoginIDViewModel.LoginIDKey,
				Value: enterLoginIDViewModel.OldLoginIDValue,
			})
			if err != nil {
				return err
			}

			return nil
		})
	}

	if r.Method == "POST" && r.Form.Get("x_action") == "add_or_update" {
		h.Database.WithTx(func() error {
			state, err := h.State.CloneState(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return err
			}

			var result *interactionflows.WebAppResult
			defer func() {
				h.State.UpdateState(state, result, err)
				h.Responder.Respond(w, r, state, result, err)
			}()

			enterLoginIDViewModel := NewEnterLoginIDViewModel(state)
			oldLoginID := enterLoginIDViewModel.OldLoginIDValue
			newLoginID := r.Form.Get("x_login_id")

			userID := auth.GetSession(r.Context()).AuthnAttrs().UserID

			if oldLoginID != "" {
				result, err = h.Interactions.UpdateLoginID(
					state,
					userID,
					loginid.LoginID{
						Key:   enterLoginIDViewModel.LoginIDKey,
						Value: oldLoginID,
					},
					loginid.LoginID{
						Key:   enterLoginIDViewModel.LoginIDKey,
						Value: newLoginID,
					},
				)
			} else {
				result, err = h.Interactions.AddLoginID(state, userID, loginid.LoginID{
					Key:   enterLoginIDViewModel.LoginIDKey,
					Value: newLoginID,
				})
			}
			if err != nil {
				return err
			}

			return nil
		})
	}
}
