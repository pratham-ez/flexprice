# Webhooks

## Overview

### Available Operations

* [PostWebhooksChargebeeTenantIDEnvironmentID](#postwebhookschargebeetenantidenvironmentid) - Handle Chargebee webhook events
* [PostWebhooksHubspotTenantIDEnvironmentID](#postwebhookshubspottenantidenvironmentid) - Handle HubSpot webhook events
* [PostWebhooksNomodTenantIDEnvironmentID](#postwebhooksnomodtenantidenvironmentid) - Handle Nomod webhook events
* [PostWebhooksQuickbooksTenantIDEnvironmentID](#postwebhooksquickbookstenantidenvironmentid) - Handle QuickBooks webhook events
* [PostWebhooksRazorpayTenantIDEnvironmentID](#postwebhooksrazorpaytenantidenvironmentid) - Handle Razorpay webhook events
* [PostWebhooksStripeTenantIDEnvironmentID](#postwebhooksstripetenantidenvironmentid) - Handle Stripe webhook events

## PostWebhooksChargebeeTenantIDEnvironmentID

Process incoming Chargebee webhook events for payment status updates

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhooks/chargebee/{tenant_id}/{environment_id}" method="post" path="/webhooks/chargebee/{tenant_id}/{environment_id}" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Webhooks.PostWebhooksChargebeeTenantIDEnvironmentID(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenantID`                                               | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `environmentID`                                          | *string*                                                 | :heavy_check_mark:                                       | Environment ID                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[map[string]any](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## PostWebhooksHubspotTenantIDEnvironmentID

Process incoming HubSpot webhook events for deal closed won and customer creation

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhooks/hubspot/{tenant_id}/{environment_id}" method="post" path="/webhooks/hubspot/{tenant_id}/{environment_id}" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Webhooks.PostWebhooksHubspotTenantIDEnvironmentID(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenantID`                                               | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `environmentID`                                          | *string*                                                 | :heavy_check_mark:                                       | Environment ID                                           |
| `xHubSpotSignatureV3`                                    | *string*                                                 | :heavy_check_mark:                                       | HubSpot webhook signature                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[map[string]any](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## PostWebhooksNomodTenantIDEnvironmentID

Process incoming Nomod webhook events for payment and invoice payments

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhooks/nomod/{tenant_id}/{environment_id}" method="post" path="/webhooks/nomod/{tenant_id}/{environment_id}" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Webhooks.PostWebhooksNomodTenantIDEnvironmentID(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenantID`                                               | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `environmentID`                                          | *string*                                                 | :heavy_check_mark:                                       | Environment ID                                           |
| `xAPIKey`                                                | **string*                                                | :heavy_minus_sign:                                       | Nomod webhook secret (if configured)                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[map[string]any](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## PostWebhooksQuickbooksTenantIDEnvironmentID

Process incoming QuickBooks webhook events for payment sync

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhooks/quickbooks/{tenant_id}/{environment_id}" method="post" path="/webhooks/quickbooks/{tenant_id}/{environment_id}" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Webhooks.PostWebhooksQuickbooksTenantIDEnvironmentID(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenantID`                                               | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `environmentID`                                          | *string*                                                 | :heavy_check_mark:                                       | Environment ID                                           |
| `intuitSignature`                                        | **string*                                                | :heavy_minus_sign:                                       | QuickBooks webhook signature                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[map[string]any](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## PostWebhooksRazorpayTenantIDEnvironmentID

Process incoming Razorpay webhook events for payment capture and failure

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhooks/razorpay/{tenant_id}/{environment_id}" method="post" path="/webhooks/razorpay/{tenant_id}/{environment_id}" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Webhooks.PostWebhooksRazorpayTenantIDEnvironmentID(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenantID`                                               | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `environmentID`                                          | *string*                                                 | :heavy_check_mark:                                       | Environment ID                                           |
| `xRazorpaySignature`                                     | *string*                                                 | :heavy_check_mark:                                       | Razorpay webhook signature                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[map[string]any](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## PostWebhooksStripeTenantIDEnvironmentID

Process incoming Stripe webhook events for payment status updates and customer creation

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhooks/stripe/{tenant_id}/{environment_id}" method="post" path="/webhooks/stripe/{tenant_id}/{environment_id}" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Webhooks.PostWebhooksStripeTenantIDEnvironmentID(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenantID`                                               | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `environmentID`                                          | *string*                                                 | :heavy_check_mark:                                       | Environment ID                                           |
| `stripeSignature`                                        | *string*                                                 | :heavy_check_mark:                                       | Stripe webhook signature                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[map[string]any](../../.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |