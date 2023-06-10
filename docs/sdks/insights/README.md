# Insights

### Available Operations

* [GetAllInsightsBranches](#getallinsightsbranches) - Get all branches for a project
* [GetFlakyTests](#getflakytests) - Get flaky tests for a project
* [GetJobTimeseries](#getjobtimeseries) - Job timeseries data
* [GetOrgSummaryData](#getorgsummarydata) - Get summary metrics with trends for the entire org, and for each project.
* [GetProjectWorkflowJobMetrics](#getprojectworkflowjobmetrics) - Get summary metrics for a project workflow's jobs.
* [GetProjectWorkflowMetrics](#getprojectworkflowmetrics) - Get summary metrics for a project's workflows
* [GetProjectWorkflowRuns](#getprojectworkflowruns) - Get recent runs of a workflow
* [GetProjectWorkflowTestMetrics](#getprojectworkflowtestmetrics) - Get test metrics for a project's workflows
* [GetProjectWorkflowsPageData](#getprojectworkflowspagedata) - Get summary metrics and trends for a project across it's workflows and branches
* [GetWorkflowSummary](#getworkflowsummary) - Get metrics and trends for workflows

## GetAllInsightsBranches

Get a list of all branches for a specified project. The list will only contain branches currently available within Insights. The maximum number of branches returned by this endpoint is 5,000.

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
    res, err := s.Insights.GetAllInsightsBranches(ctx, operations.GetAllInsightsBranchesRequest{
        ProjectSlug: "aliquid",
        WorkflowName: sdk.String("laborum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllInsightsBranches200ApplicationJSONAny != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAllInsightsBranchesRequest](../../models/operations/getallinsightsbranchesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetAllInsightsBranchesResponse](../../models/operations/getallinsightsbranchesresponse.md), error**


## GetFlakyTests

Get a list of flaky tests for a given project. Flaky tests are branch agnostic. 
             A flaky test is a test that passed and failed in the same commit.

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
    res, err := s.Insights.GetFlakyTests(ctx, operations.GetFlakyTestsRequest{
        ProjectSlug: "accusamus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetFlakyTests200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetFlakyTestsRequest](../../models/operations/getflakytestsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetFlakyTestsResponse](../../models/operations/getflakytestsresponse.md), error**


## GetJobTimeseries

Get timeseries data for all jobs within a workflow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
	"CircleCi/pkg/types"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Insights.GetJobTimeseries(ctx, operations.GetJobTimeseriesRequest{
        Branch: sdk.String("non"),
        EndDate: types.MustTimeFromString("2022-05-17T08:24:52.669Z"),
        Granularity: operations.GetJobTimeseriesGranularityHourly.ToPointer(),
        ProjectSlug: "delectus",
        StartDate: types.MustTimeFromString("2021-10-28T10:05:29.849Z"),
        WorkflowName: "nam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetJobTimeseries200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetJobTimeseriesRequest](../../models/operations/getjobtimeseriesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetJobTimeseriesResponse](../../models/operations/getjobtimeseriesresponse.md), error**


## GetOrgSummaryData

Gets aggregated summary metrics with trends for the entire org. 
              Also gets aggregated metrics and trends for each project belonging to the org.

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
    res, err := s.Insights.GetOrgSummaryData(ctx, operations.GetOrgSummaryDataRequest{
        OrgSlug: "id",
        ProjectNames: &operations.GetOrgSummaryDataProjectNames{},
        ReportingWindow: operations.GetOrgSummaryDataReportingWindowLast24Hours.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetOrgSummaryData200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetOrgSummaryDataRequest](../../models/operations/getorgsummarydatarequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetOrgSummaryDataResponse](../../models/operations/getorgsummarydataresponse.md), error**


## GetProjectWorkflowJobMetrics

Get summary metrics for a project workflow's jobs. Job runs going back at most 90 days are included in the aggregation window. Metrics are refreshed daily, and thus may not include executions from the last 24 hours. Please note that Insights is not a financial reporting tool and should not be used for precise credit reporting.  Credit reporting from Insights does not use the same source of truth as the billing information that is found in the Plan Overview page in the CircleCI UI, nor does the underlying data have the same data accuracy guarantees as the billing information in the CircleCI UI.  This may lead to discrepancies between credits reported from Insights and the billing information in the Plan Overview page of the CircleCI UI.  For precise credit reporting, always use the Plan Overview page in the CircleCI UI.

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
    res, err := s.Insights.GetProjectWorkflowJobMetrics(ctx, operations.GetProjectWorkflowJobMetricsRequest{
        AllBranches: sdk.Bool(false),
        Branch: sdk.String("deleniti"),
        PageToken: sdk.String("sapiente"),
        ProjectSlug: "amet",
        ReportingWindow: operations.GetProjectWorkflowJobMetricsReportingWindowLast30Days.ToPointer(),
        WorkflowName: "nisi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetProjectWorkflowJobMetrics200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetProjectWorkflowJobMetricsRequest](../../models/operations/getprojectworkflowjobmetricsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetProjectWorkflowJobMetricsResponse](../../models/operations/getprojectworkflowjobmetricsresponse.md), error**


## GetProjectWorkflowMetrics

Get summary metrics for a project's workflows.  Workflow runs going back at most 90 days are included in the aggregation window. Metrics are refreshed daily, and thus may not include executions from the last 24 hours.  Please note that Insights is not a financial reporting tool and should not be used for precise credit reporting.  Credit reporting from Insights does not use the same source of truth as the billing information that is found in the Plan Overview page in the CircleCI UI, nor does the underlying data have the same data accuracy guarantees as the billing information in the CircleCI UI.  This may lead to discrepancies between credits reported from Insights and the billing information in the Plan Overview page of the CircleCI UI.  For precise credit reporting, always use the Plan Overview page in the CircleCI UI.

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
    res, err := s.Insights.GetProjectWorkflowMetrics(ctx, operations.GetProjectWorkflowMetricsRequest{
        AllBranches: sdk.Bool(false),
        Branch: sdk.String("vel"),
        PageToken: sdk.String("natus"),
        ProjectSlug: "omnis",
        ReportingWindow: operations.GetProjectWorkflowMetricsReportingWindowLast24Hours.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetProjectWorkflowMetrics200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetProjectWorkflowMetricsRequest](../../models/operations/getprojectworkflowmetricsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetProjectWorkflowMetricsResponse](../../models/operations/getprojectworkflowmetricsresponse.md), error**


## GetProjectWorkflowRuns

Get recent runs of a workflow. Runs going back at most 90 days are returned. Please note that Insights is not a financial reporting tool and should not be used for precise credit reporting.  Credit reporting from Insights does not use the same source of truth as the billing information that is found in the Plan Overview page in the CircleCI UI, nor does the underlying data have the same data accuracy guarantees as the billing information in the CircleCI UI.  This may lead to discrepancies between credits reported from Insights and the billing information in the Plan Overview page of the CircleCI UI.  For precise credit reporting, always use the Plan Overview page in the CircleCI UI.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
	"CircleCi/pkg/types"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Insights.GetProjectWorkflowRuns(ctx, operations.GetProjectWorkflowRunsRequest{
        AllBranches: sdk.Bool(false),
        Branch: sdk.String("perferendis"),
        EndDate: types.MustTimeFromString("2022-09-12T22:12:15.897Z"),
        PageToken: sdk.String("distinctio"),
        ProjectSlug: "id",
        StartDate: types.MustTimeFromString("2022-09-17T02:55:11.783Z"),
        WorkflowName: "suscipit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetProjectWorkflowRuns200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetProjectWorkflowRunsRequest](../../models/operations/getprojectworkflowrunsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetProjectWorkflowRunsResponse](../../models/operations/getprojectworkflowrunsresponse.md), error**


## GetProjectWorkflowTestMetrics

Get test metrics for a project's workflows. Currently tests metrics are calculated based on 10 most recent workflow runs.

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
    res, err := s.Insights.GetProjectWorkflowTestMetrics(ctx, operations.GetProjectWorkflowTestMetricsRequest{
        AllBranches: sdk.Bool(false),
        Branch: sdk.String("natus"),
        ProjectSlug: "nobis",
        WorkflowName: "eum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetProjectWorkflowTestMetrics200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetProjectWorkflowTestMetricsRequest](../../models/operations/getprojectworkflowtestmetricsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetProjectWorkflowTestMetricsResponse](../../models/operations/getprojectworkflowtestmetricsresponse.md), error**


## GetProjectWorkflowsPageData

Get summary metrics and trends for a project at workflow and branch level. 
             Workflow runs going back at most 90 days are included in the aggregation window. 
             Trends are only supported upto last 30 days. 
             Please note that Insights is not a financial reporting tool and should not be used for precise credit reporting.  Credit reporting from Insights does not use the same source of truth as the billing information that is found in the Plan Overview page in the CircleCI UI, nor does the underlying data have the same data accuracy guarantees as the billing information in the CircleCI UI.  This may lead to discrepancies between credits reported from Insights and the billing information in the Plan Overview page of the CircleCI UI.  For precise credit reporting, always use the Plan Overview page in the CircleCI UI.

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
    res, err := s.Insights.GetProjectWorkflowsPageData(ctx, operations.GetProjectWorkflowsPageDataRequest{
        Branches: &operations.GetProjectWorkflowsPageDataBranches{},
        ProjectSlug: "vero",
        ReportingWindow: operations.GetProjectWorkflowsPageDataReportingWindowLast7Days.ToPointer(),
        WorkflowNames: &operations.GetProjectWorkflowsPageDataWorkflowNames{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetProjectWorkflowsPageData200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetProjectWorkflowsPageDataRequest](../../models/operations/getprojectworkflowspagedatarequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetProjectWorkflowsPageDataResponse](../../models/operations/getprojectworkflowspagedataresponse.md), error**


## GetWorkflowSummary

Get the metrics and trends for a particular workflow on a single branch or all branches

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
    res, err := s.Insights.GetWorkflowSummary(ctx, operations.GetWorkflowSummaryRequest{
        AllBranches: sdk.Bool(false),
        Branches: &operations.GetWorkflowSummaryBranches{},
        ProjectSlug: "architecto",
        WorkflowName: "magnam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetWorkflowSummary200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetWorkflowSummaryRequest](../../models/operations/getworkflowsummaryrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetWorkflowSummaryResponse](../../models/operations/getworkflowsummaryresponse.md), error**

