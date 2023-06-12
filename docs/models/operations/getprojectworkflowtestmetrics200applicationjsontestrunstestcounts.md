# GetProjectWorkflowTestMetrics200ApplicationJSONTestRunsTestCounts

Test counts for a given pipeline number


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Error`                                     | *int64*                                     | :heavy_check_mark:                          | The number of tests with the error status   |
| `Failure`                                   | *int64*                                     | :heavy_check_mark:                          | The number of tests with the failure status |
| `Skipped`                                   | *int64*                                     | :heavy_check_mark:                          | The number of tests with the skipped status |
| `Success`                                   | *int64*                                     | :heavy_check_mark:                          | The number of tests with the success status |
| `Total`                                     | *int64*                                     | :heavy_check_mark:                          | The total number of tests                   |