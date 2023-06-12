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
        Configuration: "reiciendis",
        ContinuationKey: "mollitia",
        Parameters: map[string]interface{}{
            "eum": "dolor",
            "necessitatibus": "odit",
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

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ContinuePipelineRequestBody](../../models/operations/continuepipelinerequestbody.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ContinuePipelineResponse](../../models/operations/continuepipelineresponse.md), error**


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
        PipelineID: "516fe4c8-b711-4e5b-bfd2-ed028921cddc",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetPipelineByIDRequest](../../models/operations/getpipelinebyidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetPipelineByIDResponse](../../models/operations/getpipelinebyidresponse.md), error**


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
        PipelineNumber: "ea",
        ProjectSlug: "excepturi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetPipelineByNumberRequest](../../models/operations/getpipelinebynumberrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetPipelineByNumberResponse](../../models/operations/getpipelinebynumberresponse.md), error**


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
        PipelineID: "2601fb57-6b0d-45f0-930c-5fbb25870532",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PipelineConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPipelineConfigByIDRequest](../../models/operations/getpipelineconfigbyidrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetPipelineConfigByIDResponse](../../models/operations/getpipelineconfigbyidresponse.md), error**


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
        PageToken: sdk.String("perferendis"),
        ProjectSlug: "dolores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PipelineListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListMyPipelinesRequest](../../models/operations/listmypipelinesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListMyPipelinesResponse](../../models/operations/listmypipelinesresponse.md), error**


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
        OrgSlug: sdk.String("minus"),
        PageToken: sdk.String("quam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PipelineListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListPipelinesRequest](../../models/operations/listpipelinesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListPipelinesResponse](../../models/operations/listpipelinesresponse.md), error**


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
        Branch: sdk.String("dolor"),
        PageToken: sdk.String("vero"),
        ProjectSlug: "nostrum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PipelineListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListPipelinesForProjectRequest](../../models/operations/listpipelinesforprojectrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.ListPipelinesForProjectResponse](../../models/operations/listpipelinesforprojectresponse.md), error**


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
        PageToken: sdk.String("hic"),
        PipelineID: "e9b90c28-909b-43fe-89a8-d9cbf4863332",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListWorkflowsByPipelineIDRequest](../../models/operations/listworkflowsbypipelineidrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.ListWorkflowsByPipelineIDResponse](../../models/operations/listworkflowsbypipelineidresponse.md), error**


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
                "hic": "excepturi",
            },
            Tag: sdk.String("v3.1.4159"),
        },
        ProjectSlug: "cum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PipelineCreation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.TriggerPipelineRequest](../../models/operations/triggerpipelinerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.TriggerPipelineResponse](../../models/operations/triggerpipelineresponse.md), error**

