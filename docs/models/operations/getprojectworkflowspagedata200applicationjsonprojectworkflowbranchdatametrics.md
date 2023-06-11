# GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowBranchDataMetrics

Metrics aggregated across a workflow or branchfor a project.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `P95DurationSecs`                                                | *float32*                                                        | :heavy_check_mark:                                               | The 95th percentile duration among a group of workflow runs.     |
| `SuccessRate`                                                    | *float32*                                                        | :heavy_check_mark:                                               | N/A                                                              |
| `TotalCreditsUsed`                                               | *int64*                                                          | :heavy_check_mark:                                               | The total credits consumed over the current timeseries interval. |
| `TotalRuns`                                                      | *int64*                                                          | :heavy_check_mark:                                               | The total number of runs.                                        |