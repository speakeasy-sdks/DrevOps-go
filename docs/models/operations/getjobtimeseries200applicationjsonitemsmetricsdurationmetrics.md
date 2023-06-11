# GetJobTimeseries200ApplicationJSONItemsMetricsDurationMetrics

Metrics relating to the duration of runs for a workflow.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Max`                                                            | *int64*                                                          | :heavy_check_mark:                                               | The max duration, in seconds, among a group of runs.             |
| `Median`                                                         | *int64*                                                          | :heavy_check_mark:                                               | The median duration, in seconds, among a group of runs.          |
| `Min`                                                            | *int64*                                                          | :heavy_check_mark:                                               | The minimum duration, in seconds, among a group of runs.         |
| `P95`                                                            | *int64*                                                          | :heavy_check_mark:                                               | The 95th percentile duration, in seconds, among a group of runs. |
| `Total`                                                          | *int64*                                                          | :heavy_check_mark:                                               | The total duration, in seconds, added across a group of runs.    |