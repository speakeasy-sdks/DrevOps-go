# Pipeline

### Available Operations

* [ContinuePipeline](#continuepipeline) - Continue a pipeline
* [GetPipelineByID](#getpipelinebyid) - Get a pipeline by ID
* [GetPipelineByNumber](#getpipelinebynumber) - Get a pipeline by pipeline number
* [GetPipelineConfigByID](#getpipelineconfigbyid) - Get a pipeline's configuration
* [ListMyPipelines](#listmypipelines) - Get your pipelines
* [ListPipelines](#listpipelines) - Get a list of pipelines
* [ListPipelinesForProject](#listpipelinesforproject) - Get all pipelines
* [ListWorkflowsByPipelineID](#listworkflowsbypipelineid) - Get a pipeline's workflows
* [TriggerPipeline](#triggerpipeline) - Trigger a new pipeline

## ContinuePipeline

Continue a pipeline from the setup phase.

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
    res, err := s.Pipeline.ContinuePipeline(ctx, operations.ContinuePipelineRequestBody{
        Configuration: "nihil",
        ContinuationKey: "repellat",
        Parameters: map[string]interface{}{
            "sed": "saepe",
            "pariatur": "accusantium",
            "consequuntur": "praesentium",
            "natus": "magni",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

## GetPipelineByID

Returns a pipeline by the pipeline ID.

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
    res, err := s.Pipeline.GetPipelineByID(ctx, operations.GetPipelineByIDRequest{
        PipelineID: "1cddc692-601f-4b57-ab0d-5f0d30c5fbb2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pipeline != nil {
        // handle response
    }
}
```

## GetPipelineByNumber

Returns a pipeline by the pipeline number.

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
    res, err := s.Pipeline.GetPipelineByNumber(ctx, operations.GetPipelineByNumberRequest{
        PipelineNumber: "quis",
        ProjectSlug: "totam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pipeline != nil {
        // handle response
    }
}
```

## GetPipelineConfigByID

Returns a pipeline's configuration by ID.

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
    res, err := s.Pipeline.GetPipelineConfigByID(ctx, operations.GetPipelineConfigByIDRequest{
        PipelineID: "7053202c-73d5-4fe9-b90c-28909b3fe49a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PipelineConfig != nil {
        // handle response
    }
}
```

## ListMyPipelines

Returns a sequence of all pipelines for this project triggered by the user.

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
    res, err := s.Pipeline.ListMyPipelines(ctx, operations.ListMyPipelinesRequest{
        PageToken: sdk.String("deleniti"),
        ProjectSlug: "pariatur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PipelineListResponse != nil {
        // handle response
    }
}
```

## ListPipelines

Returns all pipelines for the most recently built projects (max 250) you follow in an organization.

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
    res, err := s.Pipeline.ListPipelines(ctx, operations.ListPipelinesRequest{
        Mine: sdk.Bool(false),
        OrgSlug: sdk.String("provident"),
        PageToken: sdk.String("nobis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PipelineListResponse != nil {
        // handle response
    }
}
```

## ListPipelinesForProject

Returns all pipelines for this project.

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
    res, err := s.Pipeline.ListPipelinesForProject(ctx, operations.ListPipelinesForProjectRequest{
        Branch: sdk.String("libero"),
        PageToken: sdk.String("delectus"),
        ProjectSlug: "quaerat",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PipelineListResponse != nil {
        // handle response
    }
}
```

## ListWorkflowsByPipelineID

Returns a paginated list of workflows by pipeline ID.

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
    res, err := s.Pipeline.ListWorkflowsByPipelineID(ctx, operations.ListWorkflowsByPipelineIDRequest{
        PageToken: sdk.String("quos"),
        PipelineID: "633323f9-b77f-43a4-9006-74ebf69280d1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowListResponse != nil {
        // handle response
    }
}
```

## TriggerPipeline

Triggers a new pipeline on the project.

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
    res, err := s.Pipeline.TriggerPipeline(ctx, operations.TriggerPipelineRequest{
        RequestBody: &operations.TriggerPipelineTriggerPipelineParameters{
            Branch: sdk.String("feature/design-new-api"),
            Parameters: map[string]interface{}{
                "dolorum": "iusto",
                "voluptate": "dolorum",
                "deleniti": "omnis",
            },
            Tag: sdk.String("v3.1.4159"),
        },
        ProjectSlug: "necessitatibus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PipelineCreation != nil {
        // handle response
    }
}
```
