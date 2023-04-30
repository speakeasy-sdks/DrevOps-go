// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CreateWebhookRequestBodyEventsEnum string

const (
	CreateWebhookRequestBodyEventsEnumWorkflowCompleted CreateWebhookRequestBodyEventsEnum = "workflow-completed"
	CreateWebhookRequestBodyEventsEnumJobCompleted      CreateWebhookRequestBodyEventsEnum = "job-completed"
)

func (e CreateWebhookRequestBodyEventsEnum) ToPointer() *CreateWebhookRequestBodyEventsEnum {
	return &e
}

func (e *CreateWebhookRequestBodyEventsEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "workflow-completed":
		fallthrough
	case "job-completed":
		*e = CreateWebhookRequestBodyEventsEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateWebhookRequestBodyEventsEnum: %s", s)
	}
}

// CreateWebhookRequestBodyScopeTypeEnum - Type of the scope being used
type CreateWebhookRequestBodyScopeTypeEnum string

const (
	CreateWebhookRequestBodyScopeTypeEnumProject CreateWebhookRequestBodyScopeTypeEnum = "project"
)

func (e CreateWebhookRequestBodyScopeTypeEnum) ToPointer() *CreateWebhookRequestBodyScopeTypeEnum {
	return &e
}

func (e *CreateWebhookRequestBodyScopeTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "project":
		*e = CreateWebhookRequestBodyScopeTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateWebhookRequestBodyScopeTypeEnum: %s", s)
	}
}

// CreateWebhookRequestBodyScope - The scope in which the relevant events that will trigger webhooks
type CreateWebhookRequestBodyScope struct {
	// ID of the scope being used (at the moment, only project ID is supported)
	ID string `json:"id"`
	// Type of the scope being used
	Type CreateWebhookRequestBodyScopeTypeEnum `json:"type"`
}

// CreateWebhookRequestBody - The parameters for a create webhook request
type CreateWebhookRequestBody struct {
	// Events that will trigger the webhook
	Events []CreateWebhookRequestBodyEventsEnum `json:"events"`
	// Name of the webhook
	Name string `json:"name"`
	// The scope in which the relevant events that will trigger webhooks
	Scope CreateWebhookRequestBodyScope `json:"scope"`
	// Secret used to build an HMAC hash of the payload and passed as a header in the webhook request
	SigningSecret string `json:"signing-secret"`
	// URL to deliver the webhook to. Note: protocol must be included as well (only https is supported)
	URL string `json:"url"`
	// Whether to enforce TLS certificate verification when delivering the webhook
	VerifyTLS bool `json:"verify-tls"`
}

// CreateWebhookDefaultApplicationJSON - Error response.
type CreateWebhookDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateWebhookWebhookEventsEnum string

const (
	CreateWebhookWebhookEventsEnumWorkflowCompleted CreateWebhookWebhookEventsEnum = "workflow-completed"
	CreateWebhookWebhookEventsEnumJobCompleted      CreateWebhookWebhookEventsEnum = "job-completed"
)

func (e CreateWebhookWebhookEventsEnum) ToPointer() *CreateWebhookWebhookEventsEnum {
	return &e
}

func (e *CreateWebhookWebhookEventsEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "workflow-completed":
		fallthrough
	case "job-completed":
		*e = CreateWebhookWebhookEventsEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateWebhookWebhookEventsEnum: %s", s)
	}
}

// CreateWebhookWebhookScope - The scope in which the relevant events that will trigger webhooks
type CreateWebhookWebhookScope struct {
	// ID of the scope being used (at the moment, only project ID is supported)
	ID string `json:"id"`
	// Type of the scope being used
	Type string `json:"type"`
}

// CreateWebhookWebhook - A webhook
type CreateWebhookWebhook struct {
	// The date and time the webhook was created.
	CreatedAt time.Time `json:"created-at"`
	// Events that will trigger the webhook
	Events []CreateWebhookWebhookEventsEnum `json:"events"`
	// The unique ID of the webhook
	ID string `json:"id"`
	// Name of the webhook
	Name string `json:"name"`
	// The scope in which the relevant events that will trigger webhooks
	Scope CreateWebhookWebhookScope `json:"scope"`
	// Masked value of the secret used to build an HMAC hash of the payload and passed as a header in the webhook request
	SigningSecret string `json:"signing-secret"`
	// The date and time the webhook was last updated.
	UpdatedAt time.Time `json:"updated-at"`
	// URL to deliver the webhook to. Note: protocol must be included as well (only https is supported)
	URL string `json:"url"`
	// Whether to enforce TLS certificate verification when delivering the webhook
	VerifyTLS bool `json:"verify-tls"`
}

type CreateWebhookResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A webhook
	Webhook *CreateWebhookWebhook
	// Error response.
	CreateWebhookDefaultApplicationJSONObject *CreateWebhookDefaultApplicationJSON
}
