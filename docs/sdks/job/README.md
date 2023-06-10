# Job

### Available Operations

* [CancelJob](#canceljob) - Cancel job
* [GetJobArtifacts](#getjobartifacts) - Get a job's artifacts
* [GetJobDetails](#getjobdetails) - Get job details
* [GetTests](#gettests) - Get test metadata

## CancelJob

Cancel job with a given job number.

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
    res, err := s.Job.CancelJob(ctx, operations.CancelJobRequest{
        JobNumber: "et",
        ProjectSlug: "excepturi",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.CancelJobRequest](../../models/operations/canceljobrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.CancelJobResponse](../../models/operations/canceljobresponse.md), error**


## GetJobArtifacts

Returns a job's artifacts.

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
    res, err := s.Job.GetJobArtifacts(ctx, operations.GetJobArtifactsRequest{
        JobNumber: "ullam",
        ProjectSlug: "provident",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ArtifactListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetJobArtifactsRequest](../../models/operations/getjobartifactsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetJobArtifactsResponse](../../models/operations/getjobartifactsresponse.md), error**


## GetJobDetails

Returns job details.

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
    res, err := s.Job.GetJobDetails(ctx, operations.GetJobDetailsRequest{
        JobNumber: "quos",
        ProjectSlug: "sint",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JobDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetJobDetailsRequest](../../models/operations/getjobdetailsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetJobDetailsResponse](../../models/operations/getjobdetailsresponse.md), error**


## GetTests

Get test metadata for a build. In the rare case where there is more than 250MB of test data on the job, no results will be returned.

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
    res, err := s.Job.GetTests(ctx, operations.GetTestsRequest{
        JobNumber: "accusantium",
        ProjectSlug: "mollitia",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TestsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetTestsRequest](../../models/operations/gettestsrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetTestsResponse](../../models/operations/gettestsresponse.md), error**

