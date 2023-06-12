# GetJobDetailsJobDetailsMessages

Message from CircleCI execution platform.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Message`                                                       | *string*                                                        | :heavy_check_mark:                                              | Information describing message.                                 |
| `Reason`                                                        | **string*                                                       | :heavy_minus_sign:                                              | Value describing the reason for message to be added to the job. |
| `Type`                                                          | *string*                                                        | :heavy_check_mark:                                              | Message type.                                                   |