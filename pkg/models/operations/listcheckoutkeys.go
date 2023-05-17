// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ListCheckoutKeysRequest struct {
	// Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

// ListCheckoutKeysDefaultApplicationJSON - Error response.
type ListCheckoutKeysDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum - The type of checkout key. This may be either `deploy-key` or `github-user-key`.
type ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum string

const (
	ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnumDeployKey     ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum = "deploy-key"
	ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnumGithubUserKey ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum = "github-user-key"
)

func (e ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum) ToPointer() *ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum {
	return &e
}

func (e *ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deploy-key":
		fallthrough
	case "github-user-key":
		*e = ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum: %v", v)
	}
}

type ListCheckoutKeysCheckoutKeyListResponseCheckoutKey struct {
	// The date and time the checkout key was created.
	CreatedAt time.Time `json:"created-at"`
	// An SSH key fingerprint.
	Fingerprint string `json:"fingerprint"`
	// A boolean value that indicates if this key is preferred.
	Preferred bool `json:"preferred"`
	// A public SSH key.
	PublicKey string `json:"public-key"`
	// The type of checkout key. This may be either `deploy-key` or `github-user-key`.
	Type ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum `json:"type"`
}

// ListCheckoutKeysCheckoutKeyListResponse - A sequence of checkout keys.
type ListCheckoutKeysCheckoutKeyListResponse struct {
	Items []ListCheckoutKeysCheckoutKeyListResponseCheckoutKey `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

type ListCheckoutKeysResponse struct {
	// A sequence of checkout keys.
	CheckoutKeyListResponse *ListCheckoutKeysCheckoutKeyListResponse
	ContentType             string
	StatusCode              int
	RawResponse             *http.Response
	// Error response.
	ListCheckoutKeysDefaultApplicationJSONObject *ListCheckoutKeysDefaultApplicationJSON
}
