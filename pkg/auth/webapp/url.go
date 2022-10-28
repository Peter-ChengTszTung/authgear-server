package webapp

import (
	"net/url"
	"strings"

	"github.com/authgear/authgear-server/pkg/util/httputil"
)

func MakeURL(u *url.URL, path string, inQuery url.Values) *url.URL {
	uu := httputil.HostRelative(u)
	uu.RawQuery = inQuery.Encode()
	if path != "" {
		uu.Path = path
	}
	return uu
}

func PreserveQuery(q url.Values) url.Values {
	outQuery := url.Values{}
	for key := range q {
		// Preserve any query parameter that does not start with q_
		if !strings.HasPrefix(key, "q_") {
			outQuery.Set(key, q.Get(key))
		}
	}
	return outQuery
}
