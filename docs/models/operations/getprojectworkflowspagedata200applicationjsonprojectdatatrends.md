# GetProjectWorkflowsPageData200ApplicationJSONProjectDataTrends

Metric trends aggregated across all workflows and branches for a project.


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `SuccessRate`                                       | *float32*                                           | :heavy_check_mark:                                  | The trend value for the success rate.               |
| `Throughput`                                        | *float32*                                           | :heavy_check_mark:                                  | Trend value for the average number of runs per day. |
| `TotalCreditsUsed`                                  | *float32*                                           | :heavy_check_mark:                                  | The trend value for total credits consumed.         |
| `TotalDurationSecs`                                 | *float32*                                           | :heavy_check_mark:                                  | Trend value for total duration.                     |
| `TotalRuns`                                         | *float32*                                           | :heavy_check_mark:                                  | The trend value for total number of runs.           |