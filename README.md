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


### [Context](docs/context/README.md)

* [AddEnvironmentVariableToContext](docs/context/README.md#addenvironmentvariabletocontext) - Add or update an environment variable
* [CreateContext](docs/context/README.md#createcontext) - Create a new context
* [DeleteContext](docs/context/README.md#deletecontext) - Delete a context
* [DeleteEnvironmentVariableFromContext](docs/context/README.md#deleteenvironmentvariablefromcontext) - Remove an environment variable
* [GetContext](docs/context/README.md#getcontext) - Get a context
* [ListContexts](docs/context/README.md#listcontexts) - List contexts
* [ListEnvironmentVariablesFromContext](docs/context/README.md#listenvironmentvariablesfromcontext) - List environment variables

### [Insights](docs/insights/README.md)

* [GetAllInsightsBranches](docs/insights/README.md#getallinsightsbranches) - Get all branches for a project
* [GetFlakyTests](docs/insights/README.md#getflakytests) - Get flaky tests for a project
* [GetJobTimeseries](docs/insights/README.md#getjobtimeseries) - Job timeseries data
* [GetOrgSummaryData](docs/insights/README.md#getorgsummarydata) - Get summary metrics with trends for the entire org, and for each project.
* [GetProjectWorkflowJobMetrics](docs/insights/README.md#getprojectworkflowjobmetrics) - Get summary metrics for a project workflow's jobs.
* [GetProjectWorkflowMetrics](docs/insights/README.md#getprojectworkflowmetrics) - Get summary metrics for a project's workflows
* [GetProjectWorkflowRuns](docs/insights/README.md#getprojectworkflowruns) - Get recent runs of a workflow
* [GetProjectWorkflowTestMetrics](docs/insights/README.md#getprojectworkflowtestmetrics) - Get test metrics for a project's workflows
* [GetProjectWorkflowsPageData](docs/insights/README.md#getprojectworkflowspagedata) - Get summary metrics and trends for a project across it's workflows and branches
* [GetWorkflowSummary](docs/insights/README.md#getworkflowsummary) - Get metrics and trends for workflows

### [Job](docs/job/README.md)

* [CancelJob](docs/job/README.md#canceljob) - Cancel job
* [GetJobArtifacts](docs/job/README.md#getjobartifacts) - Get a job's artifacts
* [GetJobDetails](docs/job/README.md#getjobdetails) - Get job details
* [GetTests](docs/job/README.md#gettests) - Get test metadata

### [Pipeline](docs/pipeline/README.md)

* [ContinuePipeline](docs/pipeline/README.md#continuepipeline) - Continue a pipeline
* [GetPipelineByID](docs/pipeline/README.md#getpipelinebyid) - Get a pipeline by ID
* [GetPipelineByNumber](docs/pipeline/README.md#getpipelinebynumber) - Get a pipeline by pipeline number
* [GetPipelineConfigByID](docs/pipeline/README.md#getpipelineconfigbyid) - Get a pipeline's configuration
* [ListMyPipelines](docs/pipeline/README.md#listmypipelines) - Get your pipelines
* [ListPipelines](docs/pipeline/README.md#listpipelines) - Get a list of pipelines
* [ListPipelinesForProject](docs/pipeline/README.md#listpipelinesforproject) - Get all pipelines
* [ListWorkflowsByPipelineID](docs/pipeline/README.md#listworkflowsbypipelineid) - Get a pipeline's workflows
* [TriggerPipeline](docs/pipeline/README.md#triggerpipeline) - Trigger a new pipeline

### [Project](docs/project/README.md)

* [CreateCheckoutKey](docs/project/README.md#createcheckoutkey) - Create a new checkout key
* [CreateEnvVar](docs/project/README.md#createenvvar) - Create an environment variable
* [DeleteCheckoutKey](docs/project/README.md#deletecheckoutkey) - Delete a checkout key
* [DeleteEnvVar](docs/project/README.md#deleteenvvar) - Delete an environment variable
* [GetCheckoutKey](docs/project/README.md#getcheckoutkey) - Get a checkout key
* [GetEnvVar](docs/project/README.md#getenvvar) - Get a masked environment variable
* [GetProjectBySlug](docs/project/README.md#getprojectbyslug) - Get a project
* [ListCheckoutKeys](docs/project/README.md#listcheckoutkeys) - Get all checkout keys
* [ListEnvVars](docs/project/README.md#listenvvars) - List all environment variables

### [Schedule](docs/schedule/README.md)

* [CreateSchedule](docs/schedule/README.md#createschedule) - Create a schedule
* [DeleteScheduleByID](docs/schedule/README.md#deleteschedulebyid) - Delete a schedule
* [GetScheduleByID](docs/schedule/README.md#getschedulebyid) - Get a schedule
* [ListSchedulesForProject](docs/schedule/README.md#listschedulesforproject) - Get all schedules
* [UpdateSchedule](docs/schedule/README.md#updateschedule) - Update a schedule

### [User](docs/user/README.md)

* [GetCollaborations](docs/user/README.md#getcollaborations) - Collaborations
* [GetCurrentUser](docs/user/README.md#getcurrentuser) - User Information
* [GetUser](docs/user/README.md#getuser) - User Information

### [Webhook](docs/webhook/README.md)

* [CreateWebhook](docs/webhook/README.md#createwebhook) - Create a webhook
* [DeleteWebhook](docs/webhook/README.md#deletewebhook) - Delete a webhook
* [GetWebhookByID](docs/webhook/README.md#getwebhookbyid) - Get a webhook
* [GetWebhooks](docs/webhook/README.md#getwebhooks) - List webhooks
* [UpdateWebhook](docs/webhook/README.md#updatewebhook) - Update a webhook

### [Workflow](docs/workflow/README.md)

* [ApprovePendingApprovalJobByID](docs/workflow/README.md#approvependingapprovaljobbyid) - Approve a job
* [CancelWorkflow](docs/workflow/README.md#cancelworkflow) - Cancel a workflow
* [GetWorkflowByID](docs/workflow/README.md#getworkflowbyid) - Get a workflow
* [ListWorkflowJobs](docs/workflow/README.md#listworkflowjobs) - Get a workflow's jobs
* [RerunWorkflow](docs/workflow/README.md#rerunworkflow) - Rerun a workflow
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
