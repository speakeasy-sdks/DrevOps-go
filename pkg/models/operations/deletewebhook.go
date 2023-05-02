// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteWebhookRequest struct {
	// ID of the webhook (UUID)
	WebhookID string `pathParam:"style=simple,explode=false,name=webhook-id"`
}

// DeleteWebhookDefaultApplicationJSON - Error response.
type DeleteWebhookDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// DeleteWebhookMessageResponse - message response
type DeleteWebhookMessageResponse struct {
	// A human-readable message
	Message string `json:"message"`
}

type DeleteWebhookResponse struct {
	ContentType string
	// A confirmation message
	MessageResponse *DeleteWebhookMessageResponse
	StatusCode      int
	RawResponse     *http.Response
	// Error response.
	DeleteWebhookDefaultApplicationJSONObject *DeleteWebhookDefaultApplicationJSON
}
