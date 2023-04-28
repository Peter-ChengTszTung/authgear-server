package nonblocking

import (
	"github.com/authgear/authgear-server/pkg/api/event"
)

const (
	ProjectAppUpdated event.Type = "project.app.updated"
)

type ProjectAppUpdatedEventPayload struct {
	AppConfigDiff    string   `json:"app_config_diff"`
	UpdatedSecrets   []string `json:"updated_secrets"`
	UpdatedResources []string `json:"updated_resources"`
}

func (e *ProjectAppUpdatedEventPayload) NonBlockingEventType() event.Type {
	return ProjectAppUpdated
}

func (e *ProjectAppUpdatedEventPayload) UserID() string {
	return ""
}

func (e *ProjectAppUpdatedEventPayload) GetTriggeredBy() event.TriggeredByType {
	return event.TriggeredByPortal
}

func (e *ProjectAppUpdatedEventPayload) FillContext(ctx *event.Context) {
}

func (e *ProjectAppUpdatedEventPayload) ForHook() bool {
	return false
}

func (e *ProjectAppUpdatedEventPayload) ForAudit() bool {
	return true
}

func (e *ProjectAppUpdatedEventPayload) ReindexUserNeeded() bool {
	return false
}

func (e *ProjectAppUpdatedEventPayload) IsUserDeleted() bool {
	return false
}

var _ event.NonBlockingPayload = &ProjectAppUpdatedEventPayload{}