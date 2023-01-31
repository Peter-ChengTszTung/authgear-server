package hook

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/authgear/authgear-server/pkg/api/apierrors"
	"github.com/authgear/authgear-server/pkg/api/event"
	"github.com/authgear/authgear-server/pkg/util/resource"
)

type DenoHookImpl struct {
	Context           context.Context
	SyncDenoClient    SyncDenoClient
	AsyncDenoClient   AsyncDenoClient
	DenoClientFactory DenoClientFactory
	ResourceManager   ResourceManager
}

var _ DenoHook = &DenoHookImpl{}

func (h *DenoHookImpl) SupportURL(u *url.URL) bool {
	return u.Scheme == "authgeardeno"
}

func (h *DenoHookImpl) RunSync(u *url.URL, input interface{}, timeout *time.Duration) (interface{}, error) {
	script, err := h.loadScript(u)
	if err != nil {
		return nil, err
	}

	var client DenoClient = h.SyncDenoClient
	if timeout != nil {
		client = h.DenoClientFactory(*timeout)
	}

	out, err := client.Run(h.Context, string(script), input)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (h *DenoHookImpl) DeliverBlockingEvent(u *url.URL, e *event.Event) (*event.HookResponse, error) {
	out, err := h.RunSync(u, e, nil)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(out)
	if err != nil {
		return nil, err
	}

	hookResp, err := event.ParseHookResponse(bytes.NewReader(b))
	if err != nil {
		apiError := apierrors.AsAPIError(err)
		err = WebHookInvalidResponse.NewWithInfo("invalid response body", apiError.Info)
		return nil, err
	}

	return hookResp, nil
}

func (h *DenoHookImpl) DeliverNonBlockingEvent(u *url.URL, e *event.Event) error {
	script, err := h.loadScript(u)
	if err != nil {
		return err
	}

	_, err = h.AsyncDenoClient.Run(h.Context, string(script), e)
	if err != nil {
		return err
	}

	return nil
}

func (h *DenoHookImpl) loadScript(u *url.URL) ([]byte, error) {
	out, err := h.ResourceManager.Read(DenoFile, resource.AppFile{
		Path: h.rel(u.Path),
	})
	if err != nil {
		return nil, err
	}

	return out.([]byte), nil
}

// rel is a simplified version of filepath.Rel.
func (h *DenoHookImpl) rel(p string) string {
	return strings.TrimPrefix(p, "/")
}
