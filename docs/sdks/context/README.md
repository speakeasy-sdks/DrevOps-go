# Context

### Available Operations

* [AddEnvironmentVariableToContext](#addenvironmentvariabletocontext) - Add or update an environment variable
* [CreateContext](#createcontext) - Create a new context
* [DeleteContext](#deletecontext) - Delete a context
* [DeleteEnvironmentVariableFromContext](#deleteenvironmentvariablefromcontext) - Remove an environment variable
* [GetContext](#getcontext) - Get a context
* [ListContexts](#listcontexts) - List contexts
* [ListEnvironmentVariablesFromContext](#listenvironmentvariablesfromcontext) - List environment variables

## AddEnvironmentVariableToContext

Create or update an environment variable within a context. Returns information about the environment variable, not including its value.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Context.AddEnvironmentVariableToContext(ctx, operations.AddEnvironmentVariableToContextRequest{
        RequestBody: &operations.AddEnvironmentVariableToContextRequestBody{
            Value: "some-secret-value",
        },
        ContextID: "05dfc2dd-f7cc-478c-a1ba-928fc816742c",
        EnvVarName: "cum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddEnvironmentVariableToContext200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.AddEnvironmentVariableToContextRequest](../../models/operations/addenvironmentvariabletocontextrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.AddEnvironmentVariableToContextResponse](../../models/operations/addenvironmentvariabletocontextresponse.md), error**


## CreateContext

Create a new context

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Context.CreateContext(ctx, operations.CreateContextRequestBody{
        Name: "Edna Mante II",
        Owner: operations.CreateContextRequestBodyOwner{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Context != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateContextRequestBody](../../models/operations/createcontextrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateContextResponse](../../models/operations/createcontextresponse.md), error**


## DeleteContext

Delete a context

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Context.DeleteContext(ctx, operations.DeleteContextRequest{
        ContextID: "929396fe-a759-46eb-90fa-aa2352c59559",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteContextRequest](../../models/operations/deletecontextrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DeleteContextResponse](../../models/operations/deletecontextresponse.md), error**


## DeleteEnvironmentVariableFromContext

Delete an environment variable from a context.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Context.DeleteEnvironmentVariableFromContext(ctx, operations.DeleteEnvironmentVariableFromContextRequest{
        ContextID: "07aff1a3-a2fa-4946-b739-251aa52c3f5a",
        EnvVarName: "possimus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.DeleteEnvironmentVariableFromContextRequest](../../models/operations/deleteenvironmentvariablefromcontextrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.DeleteEnvironmentVariableFromContextResponse](../../models/operations/deleteenvironmentvariablefromcontextresponse.md), error**


## GetContext

Returns basic information about a context.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Context.GetContext(ctx, operations.GetContextRequest{
        ContextID: "019da1ff-e78f-4097-b007-4f15471b5e6e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Context != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetContextRequest](../../models/operations/getcontextrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetContextResponse](../../models/operations/getcontextresponse.md), error**


## ListContexts

List all contexts for an owner.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Context.ListContexts(ctx, operations.ListContextsRequest{
        OwnerID: sdk.String("13b99d48-8e1e-491e-850a-d2abd4426980"),
        OwnerSlug: sdk.String("magni"),
        OwnerType: operations.ListContextsOwnerTypeOrganization.ToPointer(),
        PageToken: sdk.String("ipsam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListContexts200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListContextsRequest](../../models/operations/listcontextsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListContextsResponse](../../models/operations/listcontextsresponse.md), error**


## ListEnvironmentVariablesFromContext

List information about environment variables in a context, not including their values.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"CircleCi"
	"CircleCi/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Context.ListEnvironmentVariablesFromContext(ctx, operations.ListEnvironmentVariablesFromContextRequest{
        ContextID: "02a94bb4-f63c-4969-a9a3-efa77dfb14cd",
        PageToken: sdk.String("ea"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListEnvironmentVariablesFromContext200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ListEnvironmentVariablesFromContextRequest](../../models/operations/listenvironmentvariablesfromcontextrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.ListEnvironmentVariablesFromContextResponse](../../models/operations/listenvironmentvariablesfromcontextresponse.md), error**

