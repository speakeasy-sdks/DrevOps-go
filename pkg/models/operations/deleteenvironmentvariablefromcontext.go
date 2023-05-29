// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteEnvironmentVariableFromContextRequest struct {
	// ID of the context (UUID)
	ContextID string `pathParam:"style=simple,explode=false,name=context-id"`
	// The name of the environment variable
	EnvVarName string `pathParam:"style=simple,explode=false,name=env-var-name"`
}

// DeleteEnvironmentVariableFromContextDefaultApplicationJSON - Error response.
type DeleteEnvironmentVariableFromContextDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// DeleteEnvironmentVariableFromContextMessageResponse - message response
type DeleteEnvironmentVariableFromContextMessageResponse struct {
	// A human-readable message
	Message string `json:"message"`
}

type DeleteEnvironmentVariableFromContextResponse struct {
	ContentType string
	// A confirmation message
	MessageResponse *DeleteEnvironmentVariableFromContextMessageResponse
	StatusCode      int
	RawResponse     *http.Response
	// Error response.
	DeleteEnvironmentVariableFromContextDefaultApplicationJSONObject *DeleteEnvironmentVariableFromContextDefaultApplicationJSON
}
