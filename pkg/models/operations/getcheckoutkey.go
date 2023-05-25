// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GetCheckoutKeyRequest struct {
	// An SSH key fingerprint.
	Fingerprint string `pathParam:"style=simple,explode=false,name=fingerprint"`
	// Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

// GetCheckoutKeyDefaultApplicationJSON - Error response.
type GetCheckoutKeyDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// GetCheckoutKeyCheckoutKeyCheckoutKeyType - The type of checkout key. This may be either `deploy-key` or `github-user-key`.
type GetCheckoutKeyCheckoutKeyCheckoutKeyType string

const (
	GetCheckoutKeyCheckoutKeyCheckoutKeyTypeDeployKey     GetCheckoutKeyCheckoutKeyCheckoutKeyType = "deploy-key"
	GetCheckoutKeyCheckoutKeyCheckoutKeyTypeGithubUserKey GetCheckoutKeyCheckoutKeyCheckoutKeyType = "github-user-key"
)

func (e GetCheckoutKeyCheckoutKeyCheckoutKeyType) ToPointer() *GetCheckoutKeyCheckoutKeyCheckoutKeyType {
	return &e
}

func (e *GetCheckoutKeyCheckoutKeyCheckoutKeyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deploy-key":
		fallthrough
	case "github-user-key":
		*e = GetCheckoutKeyCheckoutKeyCheckoutKeyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetCheckoutKeyCheckoutKeyCheckoutKeyType: %v", v)
	}
}

// GetCheckoutKeyCheckoutKey - The checkout key.
type GetCheckoutKeyCheckoutKey struct {
	// The date and time the checkout key was created.
	CreatedAt time.Time `json:"created-at"`
	// An SSH key fingerprint.
	Fingerprint string `json:"fingerprint"`
	// A boolean value that indicates if this key is preferred.
	Preferred bool `json:"preferred"`
	// A public SSH key.
	PublicKey string `json:"public-key"`
	// The type of checkout key. This may be either `deploy-key` or `github-user-key`.
	Type GetCheckoutKeyCheckoutKeyCheckoutKeyType `json:"type"`
}

type GetCheckoutKeyResponse struct {
	// The checkout key.
	CheckoutKey *GetCheckoutKeyCheckoutKey
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Error response.
	GetCheckoutKeyDefaultApplicationJSONObject *GetCheckoutKeyDefaultApplicationJSON
}
