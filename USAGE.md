<!-- Start SDK Example Usage -->
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
    res, err := s.Context.AddEnvironmentVariableToContext(ctx, operations.AddEnvironmentVariableToContextRequest{
        RequestBody: &operations.AddEnvironmentVariableToContextRequestBody{
            Value: "some-secret-value",
        },
        ContextID: "89bd9d8d-69a6-474e-8f46-7cc8796ed151",
        EnvVarName: "deserunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddEnvironmentVariableToContext200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->