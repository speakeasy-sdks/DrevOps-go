# GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowBranchDataTrends

Trends aggregated across a workflow or branch for a project.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `P95DurationSecs`                                            | *float32*                                                    | :heavy_check_mark:                                           | The 95th percentile duration among a group of workflow runs. |
| `SuccessRate`                                                | *float32*                                                    | :heavy_check_mark:                                           | The trend value for the success rate.                        |
| `TotalCreditsUsed`                                           | *float32*                                                    | :heavy_check_mark:                                           | The trend value for total credits consumed.                  |
| `TotalRuns`                                                  | *float32*                                                    | :heavy_check_mark:                                           | The trend value for total number of runs.                    |