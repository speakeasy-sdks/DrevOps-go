# AddEnvironmentVariableToContext200ApplicationJSON1


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ContextID`                                             | *string*                                                | :heavy_check_mark:                                      | ID of the context (UUID)                                |                                                         |
| `CreatedAt`                                             | [time.Time](https://pkg.go.dev/time#Time)               | :heavy_check_mark:                                      | The date and time the environment variable was created. | 2015-09-21T17:29:21.042Z                                |
| `UpdatedAt`                                             | [time.Time](https://pkg.go.dev/time#Time)               | :heavy_check_mark:                                      | The date and time the environment variable was updated  | 2015-09-21T17:29:21.042Z                                |
| `Variable`                                              | *string*                                                | :heavy_check_mark:                                      | The name of the environment variable                    | POSTGRES_USER                                           |