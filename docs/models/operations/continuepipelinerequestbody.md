# ContinuePipelineRequestBody


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Configuration`                                            | *string*                                                   | :heavy_check_mark:                                         | A configuration string for the pipeline.                   |
| `ContinuationKey`                                          | *string*                                                   | :heavy_check_mark:                                         | A pipeline continuation key.                               |
| `Parameters`                                               | map[string]*interface{}*                                   | :heavy_minus_sign:                                         | An object containing pipeline parameters and their values. |