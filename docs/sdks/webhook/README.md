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
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Webhook.CreateWebhook(ctx, operations.CreateWebhookRequestBody{
        Events: []CreateWebhookRequestBodyEvents{
            operations.CreateWebhookRequestBodyEventsJobCompleted,
            operations.CreateWebhookRequestBodyEventsJobCompleted,
        },
        Name: "Nathaniel Ryan",
        Scope: operations.CreateWebhookRequestBodyScope{
            ID: "6146c3e2-50fb-4008-842e-141aac366c8d",
            Type: operations.CreateWebhookRequestBodyScopeTypeProject,
        },
        SigningSecret: "nulla",
        URL: "voluptas",
        VerifyTLS: false,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateWebhookRequestBody](../../models/operations/createwebhookrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateWebhookResponse](../../models/operations/createwebhookresponse.md), error**


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
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Webhook.DeleteWebhook(ctx, operations.DeleteWebhookRequest{
        WebhookID: "b1442907-4747-478a-bbd4-66d28c10ab3c",
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
| `request`                                                                          | [operations.DeleteWebhookRequest](../../models/operations/deletewebhookrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DeleteWebhookResponse](../../models/operations/deletewebhookresponse.md), error**


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
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Webhook.GetWebhookByID(ctx, operations.GetWebhookByIDRequest{
        WebhookID: "dca42519-04e5-423c-be0b-c7178e4796f2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetWebhookByIDRequest](../../models/operations/getwebhookbyidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetWebhookByIDResponse](../../models/operations/getwebhookbyidresponse.md), error**


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
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Webhook.GetWebhooks(ctx, operations.GetWebhooksRequest{
        ScopeID: "a70c6882-82aa-4482-962f-222e9817ee17",
        ScopeType: operations.GetWebhooksScopeTypeProject,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetWebhooks200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetWebhooksRequest](../../models/operations/getwebhooksrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetWebhooksResponse](../../models/operations/getwebhooksresponse.md), error**


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
            APIKeyHeader: sdk.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Webhook.UpdateWebhook(ctx, operations.UpdateWebhookRequest{
        RequestBody: &operations.UpdateWebhookRequestBody{
            Events: []UpdateWebhookRequestBodyEvents{
                operations.UpdateWebhookRequestBodyEventsJobCompleted,
                operations.UpdateWebhookRequestBodyEventsJobCompleted,
                operations.UpdateWebhookRequestBodyEventsWorkflowCompleted,
                operations.UpdateWebhookRequestBodyEventsWorkflowCompleted,
            },
            Name: sdk.String("Cecil Pollich"),
            SigningSecret: sdk.String("occaecati"),
            URL: sdk.String("minima"),
            VerifyTLS: sdk.Bool(false),
        },
        WebhookID: "bc0ab3c2-0c4f-4378-9fd8-71f99dd2efd1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateWebhookRequest](../../models/operations/updatewebhookrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateWebhookResponse](../../models/operations/updatewebhookresponse.md), error**

