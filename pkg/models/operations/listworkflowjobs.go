// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ListWorkflowJobsRequest struct {
	// The unique ID of the workflow.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

// ListWorkflowJobsDefaultApplicationJSON - Error response.
type ListWorkflowJobsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// ListWorkflowJobsWorkflowJobListResponseJobStatus - The current status of the job.
type ListWorkflowJobsWorkflowJobListResponseJobStatus string

const (
	ListWorkflowJobsWorkflowJobListResponseJobStatusSuccess            ListWorkflowJobsWorkflowJobListResponseJobStatus = "success"
	ListWorkflowJobsWorkflowJobListResponseJobStatusRunning            ListWorkflowJobsWorkflowJobListResponseJobStatus = "running"
	ListWorkflowJobsWorkflowJobListResponseJobStatusNotRun             ListWorkflowJobsWorkflowJobListResponseJobStatus = "not_run"
	ListWorkflowJobsWorkflowJobListResponseJobStatusFailed             ListWorkflowJobsWorkflowJobListResponseJobStatus = "failed"
	ListWorkflowJobsWorkflowJobListResponseJobStatusRetried            ListWorkflowJobsWorkflowJobListResponseJobStatus = "retried"
	ListWorkflowJobsWorkflowJobListResponseJobStatusQueued             ListWorkflowJobsWorkflowJobListResponseJobStatus = "queued"
	ListWorkflowJobsWorkflowJobListResponseJobStatusNotRunning         ListWorkflowJobsWorkflowJobListResponseJobStatus = "not_running"
	ListWorkflowJobsWorkflowJobListResponseJobStatusInfrastructureFail ListWorkflowJobsWorkflowJobListResponseJobStatus = "infrastructure_fail"
	ListWorkflowJobsWorkflowJobListResponseJobStatusTimedout           ListWorkflowJobsWorkflowJobListResponseJobStatus = "timedout"
	ListWorkflowJobsWorkflowJobListResponseJobStatusOnHold             ListWorkflowJobsWorkflowJobListResponseJobStatus = "on_hold"
	ListWorkflowJobsWorkflowJobListResponseJobStatusTerminatedUnknown  ListWorkflowJobsWorkflowJobListResponseJobStatus = "terminated-unknown"
	ListWorkflowJobsWorkflowJobListResponseJobStatusBlocked            ListWorkflowJobsWorkflowJobListResponseJobStatus = "blocked"
	ListWorkflowJobsWorkflowJobListResponseJobStatusCanceled           ListWorkflowJobsWorkflowJobListResponseJobStatus = "canceled"
	ListWorkflowJobsWorkflowJobListResponseJobStatusUnauthorized       ListWorkflowJobsWorkflowJobListResponseJobStatus = "unauthorized"
)

func (e ListWorkflowJobsWorkflowJobListResponseJobStatus) ToPointer() *ListWorkflowJobsWorkflowJobListResponseJobStatus {
	return &e
}

func (e *ListWorkflowJobsWorkflowJobListResponseJobStatus) UnmarshalJSON(data []byte) error {
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
	case "retried":
		fallthrough
	case "queued":
		fallthrough
	case "not_running":
		fallthrough
	case "infrastructure_fail":
		fallthrough
	case "timedout":
		fallthrough
	case "on_hold":
		fallthrough
	case "terminated-unknown":
		fallthrough
	case "blocked":
		fallthrough
	case "canceled":
		fallthrough
	case "unauthorized":
		*e = ListWorkflowJobsWorkflowJobListResponseJobStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListWorkflowJobsWorkflowJobListResponseJobStatus: %v", v)
	}
}

// ListWorkflowJobsWorkflowJobListResponseJobType - The type of job.
type ListWorkflowJobsWorkflowJobListResponseJobType string

const (
	ListWorkflowJobsWorkflowJobListResponseJobTypeBuild    ListWorkflowJobsWorkflowJobListResponseJobType = "build"
	ListWorkflowJobsWorkflowJobListResponseJobTypeApproval ListWorkflowJobsWorkflowJobListResponseJobType = "approval"
)

func (e ListWorkflowJobsWorkflowJobListResponseJobType) ToPointer() *ListWorkflowJobsWorkflowJobListResponseJobType {
	return &e
}

func (e *ListWorkflowJobsWorkflowJobListResponseJobType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "build":
		fallthrough
	case "approval":
		*e = ListWorkflowJobsWorkflowJobListResponseJobType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListWorkflowJobsWorkflowJobListResponseJobType: %v", v)
	}
}

// ListWorkflowJobsWorkflowJobListResponseJob - Job
type ListWorkflowJobsWorkflowJobListResponseJob struct {
	// The unique ID of the job.
	ApprovalRequestID *string `json:"approval_request_id,omitempty"`
	// The unique ID of the user.
	ApprovedBy *string `json:"approved_by,omitempty"`
	// The unique ID of the user.
	CanceledBy *string `json:"canceled_by,omitempty"`
	// A sequence of the unique job IDs for the jobs that this job depends upon in the workflow.
	Dependencies []string `json:"dependencies"`
	// The unique ID of the job.
	ID string `json:"id"`
	// The number of the job.
	JobNumber *int64 `json:"job_number,omitempty"`
	// The name of the job.
	Name string `json:"name"`
	// The project-slug for the job.
	ProjectSlug string `json:"project_slug"`
	// The date and time the job started.
	StartedAt time.Time `json:"started_at"`
	// The current status of the job.
	Status ListWorkflowJobsWorkflowJobListResponseJobStatus `json:"status"`
	// The time when the job stopped.
	StoppedAt *time.Time `json:"stopped_at,omitempty"`
	// The type of job.
	Type ListWorkflowJobsWorkflowJobListResponseJobType `json:"type"`
}

// ListWorkflowJobsWorkflowJobListResponse - A paginated sequence of jobs.
type ListWorkflowJobsWorkflowJobListResponse struct {
	Items []ListWorkflowJobsWorkflowJobListResponseJob `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

type ListWorkflowJobsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A paginated sequence of jobs.
	WorkflowJobListResponse *ListWorkflowJobsWorkflowJobListResponse
	// Error response.
	ListWorkflowJobsDefaultApplicationJSONObject *ListWorkflowJobsDefaultApplicationJSON
}
