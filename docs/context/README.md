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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()    
    req := operations.AddEnvironmentVariableToContextRequest{
        RequestBody: &operations.AddEnvironmentVariableToContextRequestBody{
            Value: "some-secret-value",
        },
        ContextID: "05dfc2dd-f7cc-478c-a1ba-928fc816742c",
        EnvVarName: "cum",
    }

    res, err := s.Context.AddEnvironmentVariableToContext(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddEnvironmentVariableToContext200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()    
    req := operations.CreateContextRequestBody{
        Name: "Edna Mante II",
        Owner: operations.CreateContextRequestBodyOwner{},
    }

    res, err := s.Context.CreateContext(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Context != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()    
    req := operations.DeleteContextRequest{
        ContextID: "929396fe-a759-46eb-90fa-aa2352c59559",
    }

    res, err := s.Context.DeleteContext(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()    
    req := operations.DeleteEnvironmentVariableFromContextRequest{
        ContextID: "07aff1a3-a2fa-4946-b739-251aa52c3f5a",
        EnvVarName: "possimus",
    }

    res, err := s.Context.DeleteEnvironmentVariableFromContext(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()    
    req := operations.GetContextRequest{
        ContextID: "019da1ff-e78f-4097-b007-4f15471b5e6e",
    }

    res, err := s.Context.GetContext(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Context != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()    
    req := operations.ListContextsRequest{
        OwnerID: sdk.String("13b99d48-8e1e-491e-850a-d2abd4426980"),
        OwnerSlug: sdk.String("magni"),
        OwnerType: operations.ListContextsOwnerTypeEnumOrganization.ToPointer(),
        PageToken: sdk.String("ipsam"),
    }

    res, err := s.Context.ListContexts(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListContexts200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()    
    req := operations.ListEnvironmentVariablesFromContextRequest{
        ContextID: "02a94bb4-f63c-4969-a9a3-efa77dfb14cd",
        PageToken: sdk.String("ea"),
    }

    res, err := s.Context.ListEnvironmentVariablesFromContext(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListEnvironmentVariablesFromContext200ApplicationJSONObject != nil {
        // handle response
    }
}
```
