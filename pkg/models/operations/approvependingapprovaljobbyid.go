// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ApprovePendingApprovalJobByIDRequest struct {
	// The ID of the job being approved.
	ApprovalRequestID string `pathParam:"style=simple,explode=false,name=approval_request_id"`
	// The unique ID of the workflow.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

// ApprovePendingApprovalJobByIDDefaultApplicationJSON - Error response.
type ApprovePendingApprovalJobByIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// ApprovePendingApprovalJobByIDMessageResponse - message response
type ApprovePendingApprovalJobByIDMessageResponse struct {
	// A human-readable message
	Message string `json:"message"`
}

type ApprovePendingApprovalJobByIDResponse struct {
	ContentType string
	// A confirmation message.
	MessageResponse *ApprovePendingApprovalJobByIDMessageResponse
	StatusCode      int
	RawResponse     *http.Response
	// Error response.
	ApprovePendingApprovalJobByIDDefaultApplicationJSONObject *ApprovePendingApprovalJobByIDDefaultApplicationJSON
}
