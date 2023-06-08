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
        JobNumber: "deleniti",
        ProjectSlug: "facilis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

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
        JobNumber: "in",
        ProjectSlug: "architecto",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ArtifactListResponse != nil {
        // handle response
    }
}
```

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
        JobNumber: "architecto",
        ProjectSlug: "repudiandae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JobDetails != nil {
        // handle response
    }
}
```

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
        JobNumber: "ullam",
        ProjectSlug: "expedita",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TestsResponse != nil {
        // handle response
    }
}
```
