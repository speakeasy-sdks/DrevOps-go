# GetWorkflowSummary200ApplicationJSONMetricsDurationMetrics

Metrics relating to the duration of runs for a workflow.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Max`                                                            | *int64*                                                          | :heavy_check_mark:                                               | The max duration, in seconds, among a group of runs.             |
| `Mean`                                                           | *int64*                                                          | :heavy_check_mark:                                               | The mean duration, in seconds, among a group of runs.            |
| `Median`                                                         | *int64*                                                          | :heavy_check_mark:                                               | The median duration, in seconds, among a group of runs.          |
| `Min`                                                            | *int64*                                                          | :heavy_check_mark:                                               | The minimum duration, in seconds, among a group of runs.         |
| `P95`                                                            | *int64*                                                          | :heavy_check_mark:                                               | The 95th percentile duration, in seconds, among a group of runs. |
| `StandardDeviation`                                              | *float32*                                                        | :heavy_check_mark:                                               | The standard deviation, in seconds, among a group of runs.       |