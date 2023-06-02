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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Project.CreateCheckoutKey(ctx, operations.CreateCheckoutKeyRequest{
        RequestBody: &operations.CreateCheckoutKeyCheckoutKeyInput{
            Type: operations.CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeDeployKey,
        },
        ProjectSlug: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CheckoutKey != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Project.CreateEnvVar(ctx, operations.CreateEnvVarRequest{
        RequestBody: &operations.CreateEnvVarEnvironmentVariablePair{
            Name: "foo",
            Value: "xxxx1234",
        },
        ProjectSlug: "asperiores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentVariablePair != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Project.DeleteCheckoutKey(ctx, operations.DeleteCheckoutKeyRequest{
        Fingerprint: "nihil",
        ProjectSlug: "ipsum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Project.DeleteEnvVar(ctx, operations.DeleteEnvVarRequest{
        Name: "Alberta Ullrich",
        ProjectSlug: "perferendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Project.GetCheckoutKey(ctx, operations.GetCheckoutKeyRequest{
        Fingerprint: "amet",
        ProjectSlug: "optio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CheckoutKey != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Project.GetEnvVar(ctx, operations.GetEnvVarRequest{
        Name: "Tommy Turner",
        ProjectSlug: "provident",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentVariablePair != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Project.GetProjectBySlug(ctx, operations.GetProjectBySlugRequest{
        ProjectSlug: "minima",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Project != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Project.ListCheckoutKeys(ctx, operations.ListCheckoutKeysRequest{
        ProjectSlug: "repellendus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CheckoutKeyListResponse != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.Project.ListEnvVars(ctx, operations.ListEnvVarsRequest{
        ProjectSlug: "totam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentVariableListResponse != nil {
        // handle response
    }
}
```
