# Workflow

### Available Operations

* [ApprovePendingApprovalJobByID](#approvependingapprovaljobbyid) - Approve a job
* [CancelWorkflow](#cancelworkflow) - Cancel a workflow
* [GetWorkflowByID](#getworkflowbyid) - Get a workflow
* [ListWorkflowJobs](#listworkflowjobs) - Get a workflow's jobs
* [RerunWorkflow](#rerunworkflow) - Rerun a workflow

## ApprovePendingApprovalJobByID

Approves a pending approval job in a workflow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Workflow.ApprovePendingApprovalJobByID(ctx, operations.ApprovePendingApprovalJobByIDRequest{
        ApprovalRequestID: "a19f1d17-0513-439d-8808-6a1840394c26",
        ID: "071f93f5-f064-42da-87af-515cc413aa63",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

## CancelWorkflow

Cancels a running workflow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Workflow.CancelWorkflow(ctx, operations.CancelWorkflowRequest{
        ID: "aae8d678-64db-4b67-9fd5-e60b375ed4f6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

## GetWorkflowByID

Returns summary fields of a workflow by ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Workflow.GetWorkflowByID(ctx, operations.GetWorkflowByIDRequest{
        ID: "fbee41f3-3317-4fe3-9b60-eb1ea426555b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Workflow != nil {
        // handle response
    }
}
```

## ListWorkflowJobs

Returns a sequence of jobs for a workflow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Workflow.ListWorkflowJobs(ctx, operations.ListWorkflowJobsRequest{
        ID: "a3c28744-ed53-4b88-b3a8-d8f5c0b2f2fb",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowJobListResponse != nil {
        // handle response
    }
}
```

## RerunWorkflow

Reruns a workflow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Workflow.RerunWorkflow(ctx, operations.RerunWorkflowRequest{
        RequestBody: &operations.RerunWorkflowRerunWorkflowParameters{
            EnableSSH: sdk.Bool(false),
            FromFailed: sdk.Bool(false),
            Jobs: []string{
                "b194a276-b269-416f-a1f0-8f4294e3698f",
                "447f603e-8b44-45e8-8ca5-5efd20e457e1",
            },
            SparseTree: sdk.Bool(false),
        },
        ID: "858b6a89-fbe3-4a5a-a8e4-824d0ab40750",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RerunWorkflow202ApplicationJSONObject != nil {
        // handle response
    }
}
```
