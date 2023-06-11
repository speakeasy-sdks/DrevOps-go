# GetPipelineByIDPipelineTrigger

A summary of the trigger.


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `Actor`                                                                                               | [GetPipelineByIDPipelineTriggerActor](../../models/operations/getpipelinebyidpipelinetriggeractor.md) | :heavy_check_mark:                                                                                    | The user who triggered the Pipeline.                                                                  |
| `ReceivedAt`                                                                                          | [time.Time](https://pkg.go.dev/time#Time)                                                             | :heavy_check_mark:                                                                                    | The date and time the trigger was received.                                                           |
| `Type`                                                                                                | [GetPipelineByIDPipelineTriggerType](../../models/operations/getpipelinebyidpipelinetriggertype.md)   | :heavy_check_mark:                                                                                    | The type of trigger.                                                                                  |