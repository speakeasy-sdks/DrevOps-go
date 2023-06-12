# GetOrgSummaryData200ApplicationJSONOrgDataMetrics

Metrics for a single org metrics.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `SuccessRate`                                                    | *float32*                                                        | :heavy_check_mark:                                               | N/A                                                              |
| `Throughput`                                                     | *float32*                                                        | :heavy_check_mark:                                               | The average number of runs per day.                              |
| `TotalCreditsUsed`                                               | *int64*                                                          | :heavy_check_mark:                                               | The total credits consumed over the current timeseries interval. |
| `TotalDurationSecs`                                              | *int64*                                                          | :heavy_check_mark:                                               | Total duration, in seconds.                                      |
| `TotalRuns`                                                      | *int64*                                                          | :heavy_check_mark:                                               | The total number of runs.                                        |