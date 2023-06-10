# Project

### Available Operations

* [CreateCheckoutKey](#createcheckoutkey) - Create a new checkout key
* [CreateEnvVar](#createenvvar) - Create an environment variable
* [DeleteCheckoutKey](#deletecheckoutkey) - Delete a checkout key
* [DeleteEnvVar](#deleteenvvar) - Delete an environment variable
* [GetCheckoutKey](#getcheckoutkey) - Get a checkout key
* [GetEnvVar](#getenvvar) - Get a masked environment variable
* [GetProjectBySlug](#getprojectbyslug) - Get a project
* [ListCheckoutKeys](#listcheckoutkeys) - Get all checkout keys
* [ListEnvVars](#listenvvars) - List all environment variables

## CreateCheckoutKey

Creates a new checkout key. This API request is only usable with a user API token.

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
    res, err := s.Project.CreateCheckoutKey(ctx, operations.CreateCheckoutKeyRequest{
        RequestBody: &operations.CreateCheckoutKeyCheckoutKeyInput{
            Type: operations.CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeDeployKey,
        },
        ProjectSlug: "voluptate",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CheckoutKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateCheckoutKeyRequest](../../models/operations/createcheckoutkeyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateCheckoutKeyResponse](../../models/operations/createcheckoutkeyresponse.md), error**


## CreateEnvVar

Creates a new environment variable.

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
    res, err := s.Project.CreateEnvVar(ctx, operations.CreateEnvVarRequest{
        RequestBody: &operations.CreateEnvVarEnvironmentVariablePair{
            Name: "foo",
            Value: "xxxx1234",
        },
        ProjectSlug: "dignissimos",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentVariablePair != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateEnvVarRequest](../../models/operations/createenvvarrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateEnvVarResponse](../../models/operations/createenvvarresponse.md), error**


## DeleteCheckoutKey

Deletes the checkout key.

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
    res, err := s.Project.DeleteCheckoutKey(ctx, operations.DeleteCheckoutKeyRequest{
        Fingerprint: "reiciendis",
        ProjectSlug: "amet",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteCheckoutKeyRequest](../../models/operations/deletecheckoutkeyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteCheckoutKeyResponse](../../models/operations/deletecheckoutkeyresponse.md), error**


## DeleteEnvVar

Deletes the environment variable named :name.

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
    res, err := s.Project.DeleteEnvVar(ctx, operations.DeleteEnvVarRequest{
        Name: "Mr. Bradley Bogan",
        ProjectSlug: "odio",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteEnvVarRequest](../../models/operations/deleteenvvarrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteEnvVarResponse](../../models/operations/deleteenvvarresponse.md), error**


## GetCheckoutKey

Returns an individual checkout key.

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
    res, err := s.Project.GetCheckoutKey(ctx, operations.GetCheckoutKeyRequest{
        Fingerprint: "quaerat",
        ProjectSlug: "accusamus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CheckoutKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetCheckoutKeyRequest](../../models/operations/getcheckoutkeyrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetCheckoutKeyResponse](../../models/operations/getcheckoutkeyresponse.md), error**


## GetEnvVar

Returns the masked value of environment variable :name.

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
    res, err := s.Project.GetEnvVar(ctx, operations.GetEnvVarRequest{
        Name: "Jan Hodkiewicz",
        ProjectSlug: "atque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentVariablePair != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetEnvVarRequest](../../models/operations/getenvvarrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetEnvVarResponse](../../models/operations/getenvvarresponse.md), error**


## GetProjectBySlug

Retrieves a project by project slug.

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
    res, err := s.Project.GetProjectBySlug(ctx, operations.GetProjectBySlugRequest{
        ProjectSlug: "sit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Project != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetProjectBySlugRequest](../../models/operations/getprojectbyslugrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetProjectBySlugResponse](../../models/operations/getprojectbyslugresponse.md), error**


## ListCheckoutKeys

Returns a sequence of checkout keys for `:project`.

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
    res, err := s.Project.ListCheckoutKeys(ctx, operations.ListCheckoutKeysRequest{
        ProjectSlug: "fugiat",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CheckoutKeyListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListCheckoutKeysRequest](../../models/operations/listcheckoutkeysrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListCheckoutKeysResponse](../../models/operations/listcheckoutkeysresponse.md), error**


## ListEnvVars

Returns four 'x' characters, in addition to the last four ASCII characters of the value, consistent with the display of environment variable values on the CircleCI website.

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
    res, err := s.Project.ListEnvVars(ctx, operations.ListEnvVarsRequest{
        ProjectSlug: "ab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentVariableListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListEnvVarsRequest](../../models/operations/listenvvarsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListEnvVarsResponse](../../models/operations/listenvvarsresponse.md), error**

