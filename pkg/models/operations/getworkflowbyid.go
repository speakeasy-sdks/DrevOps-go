// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GetWorkflowByIDRequest struct {
	// The unique ID of the workflow.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

// GetWorkflowByIDDefaultApplicationJSON - Error response.
type GetWorkflowByIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// GetWorkflowByIDWorkflowStatusEnum - The current status of the workflow.
type GetWorkflowByIDWorkflowStatusEnum string

const (
	GetWorkflowByIDWorkflowStatusEnumSuccess      GetWorkflowByIDWorkflowStatusEnum = "success"
	GetWorkflowByIDWorkflowStatusEnumRunning      GetWorkflowByIDWorkflowStatusEnum = "running"
	GetWorkflowByIDWorkflowStatusEnumNotRun       GetWorkflowByIDWorkflowStatusEnum = "not_run"
	GetWorkflowByIDWorkflowStatusEnumFailed       GetWorkflowByIDWorkflowStatusEnum = "failed"
	GetWorkflowByIDWorkflowStatusEnumError        GetWorkflowByIDWorkflowStatusEnum = "error"
	GetWorkflowByIDWorkflowStatusEnumFailing      GetWorkflowByIDWorkflowStatusEnum = "failing"
	GetWorkflowByIDWorkflowStatusEnumOnHold       GetWorkflowByIDWorkflowStatusEnum = "on_hold"
	GetWorkflowByIDWorkflowStatusEnumCanceled     GetWorkflowByIDWorkflowStatusEnum = "canceled"
	GetWorkflowByIDWorkflowStatusEnumUnauthorized GetWorkflowByIDWorkflowStatusEnum = "unauthorized"
)

func (e GetWorkflowByIDWorkflowStatusEnum) ToPointer() *GetWorkflowByIDWorkflowStatusEnum {
	return &e
}

func (e *GetWorkflowByIDWorkflowStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "running":
		fallthrough
	case "not_run":
		fallthrough
	case "failed":
		fallthrough
	case "error":
		fallthrough
	case "failing":
		fallthrough
	case "on_hold":
		fallthrough
	case "canceled":
		fallthrough
	case "unauthorized":
		*e = GetWorkflowByIDWorkflowStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetWorkflowByIDWorkflowStatusEnum: %v", v)
	}
}

// GetWorkflowByIDWorkflowTagEnum - Tag used for the workflow
type GetWorkflowByIDWorkflowTagEnum string

const (
	GetWorkflowByIDWorkflowTagEnumSetup GetWorkflowByIDWorkflowTagEnum = "setup"
)

func (e GetWorkflowByIDWorkflowTagEnum) ToPointer() *GetWorkflowByIDWorkflowTagEnum {
	return &e
}

func (e *GetWorkflowByIDWorkflowTagEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "setup":
		*e = GetWorkflowByIDWorkflowTagEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetWorkflowByIDWorkflowTagEnum: %v", v)
	}
}

// GetWorkflowByIDWorkflow - A workflow
type GetWorkflowByIDWorkflow struct {
	CanceledBy *string `json:"canceled_by,omitempty"`
	// The date and time the workflow was created.
	CreatedAt time.Time `json:"created_at"`
	ErroredBy *string   `json:"errored_by,omitempty"`
	// The unique ID of the workflow.
	ID string `json:"id"`
	// The name of the workflow.
	Name string `json:"name"`
	// The ID of the pipeline this workflow belongs to.
	PipelineID string `json:"pipeline_id"`
	// The number of the pipeline this workflow belongs to.
	PipelineNumber int64 `json:"pipeline_number"`
	// The project-slug for the pipeline this workflow belongs to.
	ProjectSlug string `json:"project_slug"`
	StartedBy   string `json:"started_by"`
	// The current status of the workflow.
	Status GetWorkflowByIDWorkflowStatusEnum `json:"status"`
	// The date and time the workflow stopped.
	StoppedAt time.Time `json:"stopped_at"`
	// Tag used for the workflow
	Tag *GetWorkflowByIDWorkflowTagEnum `json:"tag,omitempty"`
}

type GetWorkflowByIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A workflow object.
	Workflow *GetWorkflowByIDWorkflow
	// Error response.
	GetWorkflowByIDDefaultApplicationJSONObject *GetWorkflowByIDDefaultApplicationJSON
}
