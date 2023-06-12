# CircleCi

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/DrevOps-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
    res, err := s.Context.AddEnvironmentVariableToContext(ctx, operations.AddEnvironmentVariableToContextRequest{
        RequestBody: &operations.AddEnvironmentVariableToContextRequestBody{
            Value: "some-secret-value",
        },
        ContextID: "89bd9d8d-69a6-474e-8f46-7cc8796ed151",
        EnvVarName: "deserunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddEnvironmentVariableToContext200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Context](docs/sdks/context/README.md)

* [AddEnvironmentVariableToContext](docs/sdks/context/README.md#addenvironmentvariabletocontext) - Add or update an environment variable
* [CreateContext](docs/sdks/context/README.md#createcontext) - Create a new context
* [DeleteContext](docs/sdks/context/README.md#deletecontext) - Delete a context
* [DeleteEnvironmentVariableFromContext](docs/sdks/context/README.md#deleteenvironmentvariablefromcontext) - Remove an environment variable
* [GetContext](docs/sdks/context/README.md#getcontext) - Get a context
* [ListContexts](docs/sdks/context/README.md#listcontexts) - List contexts
* [ListEnvironmentVariablesFromContext](docs/sdks/context/README.md#listenvironmentvariablesfromcontext) - List environment variables

### [Insights](docs/sdks/insights/README.md)

* [GetAllInsightsBranches](docs/sdks/insights/README.md#getallinsightsbranches) - Get all branches for a project
* [GetFlakyTests](docs/sdks/insights/README.md#getflakytests) - Get flaky tests for a project
* [GetJobTimeseries](docs/sdks/insights/README.md#getjobtimeseries) - Job timeseries data
* [GetOrgSummaryData](docs/sdks/insights/README.md#getorgsummarydata) - Get summary metrics with trends for the entire org, and for each project.
* [GetProjectWorkflowJobMetrics](docs/sdks/insights/README.md#getprojectworkflowjobmetrics) - Get summary metrics for a project workflow's jobs.
* [GetProjectWorkflowMetrics](docs/sdks/insights/README.md#getprojectworkflowmetrics) - Get summary metrics for a project's workflows
* [GetProjectWorkflowRuns](docs/sdks/insights/README.md#getprojectworkflowruns) - Get recent runs of a workflow
* [GetProjectWorkflowTestMetrics](docs/sdks/insights/README.md#getprojectworkflowtestmetrics) - Get test metrics for a project's workflows
* [GetProjectWorkflowsPageData](docs/sdks/insights/README.md#getprojectworkflowspagedata) - Get summary metrics and trends for a project across it's workflows and branches
* [GetWorkflowSummary](docs/sdks/insights/README.md#getworkflowsummary) - Get metrics and trends for workflows

### [Job](docs/sdks/job/README.md)

* [CancelJob](docs/sdks/job/README.md#canceljob) - Cancel job
* [GetJobArtifacts](docs/sdks/job/README.md#getjobartifacts) - Get a job's artifacts
* [GetJobDetails](docs/sdks/job/README.md#getjobdetails) - Get job details
* [GetTests](docs/sdks/job/README.md#gettests) - Get test metadata

### [Pipeline](docs/sdks/pipeline/README.md)

* [ContinuePipeline](docs/sdks/pipeline/README.md#continuepipeline) - Continue a pipeline
* [GetPipelineByID](docs/sdks/pipeline/README.md#getpipelinebyid) - Get a pipeline by ID
* [GetPipelineByNumber](docs/sdks/pipeline/README.md#getpipelinebynumber) - Get a pipeline by pipeline number
* [GetPipelineConfigByID](docs/sdks/pipeline/README.md#getpipelineconfigbyid) - Get a pipeline's configuration
* [ListMyPipelines](docs/sdks/pipeline/README.md#listmypipelines) - Get your pipelines
* [ListPipelines](docs/sdks/pipeline/README.md#listpipelines) - Get a list of pipelines
* [ListPipelinesForProject](docs/sdks/pipeline/README.md#listpipelinesforproject) - Get all pipelines
* [ListWorkflowsByPipelineID](docs/sdks/pipeline/README.md#listworkflowsbypipelineid) - Get a pipeline's workflows
* [TriggerPipeline](docs/sdks/pipeline/README.md#triggerpipeline) - Trigger a new pipeline

### [Project](docs/sdks/project/README.md)

* [CreateCheckoutKey](docs/sdks/project/README.md#createcheckoutkey) - Create a new checkout key
* [CreateEnvVar](docs/sdks/project/README.md#createenvvar) - Create an environment variable
* [DeleteCheckoutKey](docs/sdks/project/README.md#deletecheckoutkey) - Delete a checkout key
* [DeleteEnvVar](docs/sdks/project/README.md#deleteenvvar) - Delete an environment variable
* [GetCheckoutKey](docs/sdks/project/README.md#getcheckoutkey) - Get a checkout key
* [GetEnvVar](docs/sdks/project/README.md#getenvvar) - Get a masked environment variable
* [GetProjectBySlug](docs/sdks/project/README.md#getprojectbyslug) - Get a project
* [ListCheckoutKeys](docs/sdks/project/README.md#listcheckoutkeys) - Get all checkout keys
* [ListEnvVars](docs/sdks/project/README.md#listenvvars) - List all environment variables

### [Schedule](docs/sdks/schedule/README.md)

* [CreateSchedule](docs/sdks/schedule/README.md#createschedule) - Create a schedule
* [DeleteScheduleByID](docs/sdks/schedule/README.md#deleteschedulebyid) - Delete a schedule
* [GetScheduleByID](docs/sdks/schedule/README.md#getschedulebyid) - Get a schedule
* [ListSchedulesForProject](docs/sdks/schedule/README.md#listschedulesforproject) - Get all schedules
* [UpdateSchedule](docs/sdks/schedule/README.md#updateschedule) - Update a schedule

### [User](docs/sdks/user/README.md)

* [GetCollaborations](docs/sdks/user/README.md#getcollaborations) - Collaborations
* [GetCurrentUser](docs/sdks/user/README.md#getcurrentuser) - User Information
* [GetUser](docs/sdks/user/README.md#getuser) - User Information

### [Webhook](docs/sdks/webhook/README.md)

* [CreateWebhook](docs/sdks/webhook/README.md#createwebhook) - Create a webhook
* [DeleteWebhook](docs/sdks/webhook/README.md#deletewebhook) - Delete a webhook
* [GetWebhookByID](docs/sdks/webhook/README.md#getwebhookbyid) - Get a webhook
* [GetWebhooks](docs/sdks/webhook/README.md#getwebhooks) - List webhooks
* [UpdateWebhook](docs/sdks/webhook/README.md#updatewebhook) - Update a webhook

### [Workflow](docs/sdks/workflow/README.md)

* [ApprovePendingApprovalJobByID](docs/sdks/workflow/README.md#approvependingapprovaljobbyid) - Approve a job
* [CancelWorkflow](docs/sdks/workflow/README.md#cancelworkflow) - Cancel a workflow
* [GetWorkflowByID](docs/sdks/workflow/README.md#getworkflowbyid) - Get a workflow
* [ListWorkflowJobs](docs/sdks/workflow/README.md#listworkflowjobs) - Get a workflow's jobs
* [RerunWorkflow](docs/sdks/workflow/README.md#rerunworkflow) - Rerun a workflow
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
