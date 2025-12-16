# Webhooks

## Overview

### Available Operations

* [ProcessChargebee](#processchargebee) - Handle Chargebee webhook events
* [HandleHubspot](#handlehubspot) - Handle HubSpot webhook events
* [ProcessNomod](#processnomod) - Handle Nomod webhook events
* [HandleQuickbooks](#handlequickbooks) - Handle QuickBooks webhook events
* [HandleRazorpay](#handlerazorpay) - Handle Razorpay webhook events
* [ProcessStripe](#processstripe) - Handle Stripe webhook events

## ProcessChargebee

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

    res, err := s.Webhooks.ProcessChargebee(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.PostWebhooksChargebeeTenantIDEnvironmentIDResponse](../../models/operations/postwebhookschargebeetenantidenvironmentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## HandleHubspot

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

    res, err := s.Webhooks.HandleHubspot(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.PostWebhooksHubspotTenantIDEnvironmentIDResponse](../../models/operations/postwebhookshubspottenantidenvironmentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## ProcessNomod

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

    res, err := s.Webhooks.ProcessNomod(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.PostWebhooksNomodTenantIDEnvironmentIDResponse](../../models/operations/postwebhooksnomodtenantidenvironmentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## HandleQuickbooks

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

    res, err := s.Webhooks.HandleQuickbooks(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.PostWebhooksQuickbooksTenantIDEnvironmentIDResponse](../../models/operations/postwebhooksquickbookstenantidenvironmentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## HandleRazorpay

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

    res, err := s.Webhooks.HandleRazorpay(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.PostWebhooksRazorpayTenantIDEnvironmentIDResponse](../../models/operations/postwebhooksrazorpaytenantidenvironmentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |

## ProcessStripe

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

    res, err := s.Webhooks.ProcessStripe(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.PostWebhooksStripeTenantIDEnvironmentIDResponse](../../models/operations/postwebhooksstripetenantidenvironmentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 4XX, 5XX           | \*/\*              |