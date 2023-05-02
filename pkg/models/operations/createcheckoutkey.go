// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum - The type of checkout key to create. This may be either `deploy-key` or `user-key`.
type CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum string

const (
	CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnumUserKey   CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum = "user-key"
	CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnumDeployKey CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum = "deploy-key"
)

func (e CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum) ToPointer() *CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum {
	return &e
}

func (e *CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "user-key":
		fallthrough
	case "deploy-key":
		*e = CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum: %s", s)
	}
}

type CreateCheckoutKeyCheckoutKeyInput struct {
	// The type of checkout key to create. This may be either `deploy-key` or `user-key`.
	Type CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum `json:"type"`
}

type CreateCheckoutKeyRequest struct {
	RequestBody *CreateCheckoutKeyCheckoutKeyInput `request:"mediaType=application/json"`
	// Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

// CreateCheckoutKeyDefaultApplicationJSON - Error response.
type CreateCheckoutKeyDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum - The type of checkout key. This may be either `deploy-key` or `github-user-key`.
type CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum string

const (
	CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnumDeployKey     CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum = "deploy-key"
	CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnumGithubUserKey CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum = "github-user-key"
)

func (e CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum) ToPointer() *CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum {
	return &e
}

func (e *CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "deploy-key":
		fallthrough
	case "github-user-key":
		*e = CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum: %s", s)
	}
}

// CreateCheckoutKeyCheckoutKey - The checkout key.
type CreateCheckoutKeyCheckoutKey struct {
	// The date and time the checkout key was created.
	CreatedAt time.Time `json:"created-at"`
	// An SSH key fingerprint.
	Fingerprint string `json:"fingerprint"`
	// A boolean value that indicates if this key is preferred.
	Preferred bool `json:"preferred"`
	// A public SSH key.
	PublicKey string `json:"public-key"`
	// The type of checkout key. This may be either `deploy-key` or `github-user-key`.
	Type CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum `json:"type"`
}

type CreateCheckoutKeyResponse struct {
	// The checkout key.
	CheckoutKey *CreateCheckoutKeyCheckoutKey
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Error response.
	CreateCheckoutKeyDefaultApplicationJSONObject *CreateCheckoutKeyDefaultApplicationJSON
}
