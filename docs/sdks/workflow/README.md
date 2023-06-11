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
        ApprovalRequestID: "21aa6f1e-674b-4db0-8f15-756082d68ea1",
        ID: "9f1d1705-1339-4d08-886a-1840394c2607",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ApprovePendingApprovalJobByIDRequest](../../models/operations/approvependingapprovaljobbyidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.ApprovePendingApprovalJobByIDResponse](../../models/operations/approvependingapprovaljobbyidresponse.md), error**


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
        ID: "1f93f5f0-642d-4ac7-af51-5cc413aa63aa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CancelWorkflowRequest](../../models/operations/cancelworkflowrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CancelWorkflowResponse](../../models/operations/cancelworkflowresponse.md), error**


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
        ID: "e8d67864-dbb6-475f-95e6-0b375ed4f6fb",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Workflow != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetWorkflowByIDRequest](../../models/operations/getworkflowbyidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetWorkflowByIDResponse](../../models/operations/getworkflowbyidresponse.md), error**


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
        ID: "ee41f333-17fe-435b-a0eb-1ea426555ba3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowJobListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListWorkflowJobsRequest](../../models/operations/listworkflowjobsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListWorkflowJobsResponse](../../models/operations/listworkflowjobsresponse.md), error**


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
                "28744ed5-3b88-4f3a-8d8f-5c0b2f2fb7b1",
                "94a276b2-6916-4fe1-b08f-4294e3698f44",
                "7f603e8b-445e-480c-a55e-fd20e457e185",
                "8b6a89fb-e3a5-4aa8-a482-4d0ab4075088",
            },
            SparseTree: sdk.Bool(false),
        },
        ID: "e5186206-5e90-44f3-b119-4b8abf603a79",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RerunWorkflow202ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RerunWorkflowRequest](../../models/operations/rerunworkflowrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.RerunWorkflowResponse](../../models/operations/rerunworkflowresponse.md), error**

