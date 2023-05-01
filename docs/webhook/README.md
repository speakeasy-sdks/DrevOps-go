# Webhook

### Available Operations

* [CreateWebhook](#createwebhook) - Create a webhook
* [DeleteWebhook](#deletewebhook) - Delete a webhook
* [GetWebhookByID](#getwebhookbyid) - Get a webhook
* [GetWebhooks](#getwebhooks) - List webhooks
* [UpdateWebhook](#updatewebhook) - Update a webhook

## CreateWebhook

Create a webhook

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
    req := operations.CreateWebhookRequestBody{
        Events: []CreateWebhookRequestBodyEventsEnum{
            operations.CreateWebhookRequestBodyEventsEnumWorkflowCompleted,
        },
        Name: "Vicky Lynch",
        Scope: operations.CreateWebhookRequestBodyScope{
            ID: "6b144290-7474-4778-a7bd-466d28c10ab3",
            Type: operations.CreateWebhookRequestBodyScopeTypeEnumProject,
        },
        SigningSecret: "quo",
        URL: "illum",
        VerifyTLS: false,
    }

    res, err := s.Webhook.CreateWebhook(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

## DeleteWebhook

Delete a webhook

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
    req := operations.DeleteWebhookRequest{
        WebhookID: "ca425190-4e52-43c7-a0bc-7178e4796f2a",
    }

    res, err := s.Webhook.DeleteWebhook(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

## GetWebhookByID

Get a webhook by id.

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
    req := operations.GetWebhookByIDRequest{
        WebhookID: "70c68828-2aa4-4825-a2f2-22e9817ee17c",
    }

    res, err := s.Webhook.GetWebhookByID(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

## GetWebhooks

Get a list of webhook that match the given scope-type and scope-id

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
    req := operations.GetWebhooksRequest{
        ScopeID: "be61e6b7-b95b-4c0a-b3c2-0c4f3789fd87",
        ScopeType: operations.GetWebhooksScopeTypeEnumProject,
    }

    res, err := s.Webhook.GetWebhooks(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetWebhooks200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## UpdateWebhook

Update a webhook

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
    req := operations.UpdateWebhookRequest{
        RequestBody: &operations.UpdateWebhookRequestBody{
            Events: []UpdateWebhookRequestBodyEventsEnum{
                operations.UpdateWebhookRequestBodyEventsEnumJobCompleted,
            },
            Name: sdk.String("Kirk Stracke"),
            SigningSecret: sdk.String("eveniet"),
            URL: sdk.String("asperiores"),
            VerifyTLS: sdk.Bool(false),
        },
        WebhookID: "d121aa6f-1e67-44bd-b04f-15756082d68e",
    }

    res, err := s.Webhook.UpdateWebhook(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```
