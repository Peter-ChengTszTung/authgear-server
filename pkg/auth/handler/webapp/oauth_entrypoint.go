package webapp

import (
	"net/http"

	"github.com/authgear/authgear-server/pkg/util/httproute"
)

func ConfigureOAuthEntrypointRoute(route httproute.Route) httproute.Route {
	return route.
		WithMethods("OPTIONS", "GET").
		WithPathPattern("/_internals/oauth_entrypoint")
}

type OAuthEntrypointHandler struct{}

func (h *OAuthEntrypointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/flows/select_account", http.StatusFound)
}
