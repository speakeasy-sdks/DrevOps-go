// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GetWebhookByIDRequest struct {
	// ID of the webhook (UUID)
	WebhookID string `pathParam:"style=simple,explode=false,name=webhook-id"`
}

// GetWebhookByIDDefaultApplicationJSON - Error response.
type GetWebhookByIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetWebhookByIDWebhookEventsEnum string

const (
	GetWebhookByIDWebhookEventsEnumWorkflowCompleted GetWebhookByIDWebhookEventsEnum = "workflow-completed"
	GetWebhookByIDWebhookEventsEnumJobCompleted      GetWebhookByIDWebhookEventsEnum = "job-completed"
)

func (e GetWebhookByIDWebhookEventsEnum) ToPointer() *GetWebhookByIDWebhookEventsEnum {
	return &e
}

func (e *GetWebhookByIDWebhookEventsEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workflow-completed":
		fallthrough
	case "job-completed":
		*e = GetWebhookByIDWebhookEventsEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetWebhookByIDWebhookEventsEnum: %v", v)
	}
}

// GetWebhookByIDWebhookScope - The scope in which the relevant events that will trigger webhooks
type GetWebhookByIDWebhookScope struct {
	// ID of the scope being used (at the moment, only project ID is supported)
	ID string `json:"id"`
	// Type of the scope being used
	Type string `json:"type"`
}

// GetWebhookByIDWebhook - A webhook
type GetWebhookByIDWebhook struct {
	// The date and time the webhook was created.
	CreatedAt time.Time `json:"created-at"`
	// Events that will trigger the webhook
	Events []GetWebhookByIDWebhookEventsEnum `json:"events"`
	// The unique ID of the webhook
	ID string `json:"id"`
	// Name of the webhook
	Name string `json:"name"`
	// The scope in which the relevant events that will trigger webhooks
	Scope GetWebhookByIDWebhookScope `json:"scope"`
	// Masked value of the secret used to build an HMAC hash of the payload and passed as a header in the webhook request
	SigningSecret string `json:"signing-secret"`
	// The date and time the webhook was last updated.
	UpdatedAt time.Time `json:"updated-at"`
	// URL to deliver the webhook to. Note: protocol must be included as well (only https is supported)
	URL string `json:"url"`
	// Whether to enforce TLS certificate verification when delivering the webhook
	VerifyTLS bool `json:"verify-tls"`
}

type GetWebhookByIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A webhook
	Webhook *GetWebhookByIDWebhook
	// Error response.
	GetWebhookByIDDefaultApplicationJSONObject *GetWebhookByIDDefaultApplicationJSON
}
