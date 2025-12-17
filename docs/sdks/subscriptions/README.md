# Subscriptions

## Overview

### Available Operations

* [List](#list) - List subscriptions
* [Create](#create) - Create subscription
* [AddAddon](#addaddon) - Add addon to subscription
* [RemoveAddon](#removeaddon) - Remove addon from subscription
* [UpdateLineItem](#updatelineitem) - Update subscription line item
* [DeleteLineItem](#deletelineitem) - Delete subscription line item
* [Search](#search) - List subscriptions by filter
* [GetUsage](#getusage) - Get usage by subscription
* [Get](#get) - Get subscription
* [Activate](#activate) - Activate draft subscription
* [Cancel](#cancel) - Cancel subscription
* [ExecuteChange](#executechange) - Execute subscription plan change
* [PreviewPlanChange](#previewplanchange) - Preview subscription plan change
* [GetEntitlements](#getentitlements) - Get subscription entitlements
* [GetUpcomingGrants](#getupcominggrants) - Get upcoming credit grant applications
* [Pause](#pause) - Pause a subscription
* [ListPauses](#listpauses) - List all pauses for a subscription
* [Resume](#resume) - Resume a paused subscription

## List

Get subscriptions with optional filtering

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/subscriptions" method="get" path="/subscriptions" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.List(ctx, operations.GetSubscriptionsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListSubscriptionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetSubscriptionsRequest](../../models/operations/getsubscriptionsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetSubscriptionsResponse](../../models/operations/getsubscriptionsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a new subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/subscriptions" method="post" path="/subscriptions" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.Create(ctx, components.DtoCreateSubscriptionRequest{
        BillingCadence: components.TypesBillingCadenceRecurring,
        BillingPeriod: components.TypesBillingPeriodHalfYearly,
        Currency: "Nakfa",
        PlanID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.DtoCreateSubscriptionRequest](../../models/components/dtocreatesubscriptionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PostSubscriptionsResponse](../../models/operations/postsubscriptionsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## AddAddon

Add an addon to a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/subscriptions/addon" method="post" path="/subscriptions/addon" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.AddAddon(ctx, components.DtoAddAddonToSubscriptionRequest{
        AddonID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoAddonAssociationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [components.DtoAddAddonToSubscriptionRequest](../../models/components/dtoaddaddontosubscriptionrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.PostSubscriptionsAddonResponse](../../models/operations/postsubscriptionsaddonresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## RemoveAddon

Remove an addon from a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/subscriptions/addon" method="delete" path="/subscriptions/addon" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.RemoveAddon(ctx, components.DtoRemoveAddonRequest{
        AddonAssociationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GinH != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.DtoRemoveAddonRequest](../../models/components/dtoremoveaddonrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeleteSubscriptionsAddonResponse](../../models/operations/deletesubscriptionsaddonresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## UpdateLineItem

Update a subscription line item by terminating the existing one and creating a new one

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/subscriptions/lineitems/{id}" method="put" path="/subscriptions/lineitems/{id}" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.UpdateLineItem(ctx, "<id>", components.DtoUpdateSubscriptionLineItemRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionLineItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `id`                                                                                                               | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | Line Item ID                                                                                                       |
| `body`                                                                                                             | [components.DtoUpdateSubscriptionLineItemRequest](../../models/components/dtoupdatesubscriptionlineitemrequest.md) | :heavy_check_mark:                                                                                                 | Update Line Item Request                                                                                           |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.PutSubscriptionsLineitemsIDResponse](../../models/operations/putsubscriptionslineitemsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## DeleteLineItem

Delete a subscription line item by setting its end date

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/subscriptions/lineitems/{id}" method="delete" path="/subscriptions/lineitems/{id}" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.DeleteLineItem(ctx, "<id>", components.DtoDeleteSubscriptionLineItemRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionLineItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `id`                                                                                                               | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | Line Item ID                                                                                                       |
| `body`                                                                                                             | [components.DtoDeleteSubscriptionLineItemRequest](../../models/components/dtodeletesubscriptionlineitemrequest.md) | :heavy_check_mark:                                                                                                 | Delete Line Item Request                                                                                           |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.DeleteSubscriptionsLineitemsIDResponse](../../models/operations/deletesubscriptionslineitemsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Search

List subscriptions by filter

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/subscriptions/search" method="post" path="/subscriptions/search" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.Search(ctx, components.TypesSubscriptionFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListSubscriptionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.TypesSubscriptionFilter](../../models/components/typessubscriptionfilter.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PostSubscriptionsSearchResponse](../../models/operations/postsubscriptionssearchresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetUsage

Get usage for a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/subscriptions/usage" method="post" path="/subscriptions/usage" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.GetUsage(ctx, components.DtoGetUsageBySubscriptionRequest{
        EndTime: gosdk.Pointer("2024-03-20T00:00:00Z"),
        LifetimeUsage: gosdk.Pointer(false),
        StartTime: gosdk.Pointer("2024-03-13T00:00:00Z"),
        SubscriptionID: "123",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetUsageBySubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [components.DtoGetUsageBySubscriptionRequest](../../models/components/dtogetusagebysubscriptionrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.PostSubscriptionsUsageResponse](../../models/operations/postsubscriptionsusageresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a subscription by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/subscriptions/{id}" method="get" path="/subscriptions/{id}" -->
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

    res, err := s.Subscriptions.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Subscription ID                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSubscriptionsIDResponse](../../models/operations/getsubscriptionsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Activate

Activate a draft subscription with a new start date

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/subscriptions/{id}/activate" method="post" path="/subscriptions/{id}/activate" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.Activate(ctx, "<id>", components.DtoActivateDraftSubscriptionRequest{
        StartDate: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `id`                                                                                                             | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Subscription ID                                                                                                  |
| `body`                                                                                                           | [components.DtoActivateDraftSubscriptionRequest](../../models/components/dtoactivatedraftsubscriptionrequest.md) | :heavy_check_mark:                                                                                               | Activate Draft Subscription Request                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.PostSubscriptionsIDActivateResponse](../../models/operations/postsubscriptionsidactivateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Cancel

Cancel a subscription with enhanced proration support

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/subscriptions/{id}/cancel" method="post" path="/subscriptions/{id}/cancel" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.Cancel(ctx, "<id>", components.DtoCancelSubscriptionRequest{
        CancellationType: components.TypesCancellationTypeEndOfPeriod,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCancelSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | Subscription ID                                                                                    |
| `body`                                                                                             | [components.DtoCancelSubscriptionRequest](../../models/components/dtocancelsubscriptionrequest.md) | :heavy_check_mark:                                                                                 | Cancel Subscription Request                                                                        |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PostSubscriptionsIDCancelResponse](../../models/operations/postsubscriptionsidcancelresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ExecuteChange

Execute a subscription plan change, including proration and invoice generation

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/subscriptions/{id}/change/execute" method="post" path="/subscriptions/{id}/change/execute" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.ExecuteChange(ctx, "<id>", components.DtoSubscriptionChangeRequest{
        BillingCadence: components.TypesBillingCadenceOnetime,
        BillingCycle: components.TypesBillingCycleAnniversary,
        BillingPeriod: components.TypesBillingPeriodHalfYearly,
        ProrationBehavior: components.TypesProrationBehaviorNone,
        TargetPlanID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionChangeExecuteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | Subscription ID                                                                                    |
| `body`                                                                                             | [components.DtoSubscriptionChangeRequest](../../models/components/dtosubscriptionchangerequest.md) | :heavy_check_mark:                                                                                 | Subscription change request                                                                        |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PostSubscriptionsIDChangeExecuteResponse](../../models/operations/postsubscriptionsidchangeexecuteresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## PreviewPlanChange

Preview the impact of changing a subscription's plan, including proration calculations

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/subscriptions/{id}/change/preview" method="post" path="/subscriptions/{id}/change/preview" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.PreviewPlanChange(ctx, "<id>", components.DtoSubscriptionChangeRequest{
        BillingCadence: components.TypesBillingCadenceRecurring,
        BillingCycle: components.TypesBillingCycleAnniversary,
        BillingPeriod: components.TypesBillingPeriodWeekly,
        ProrationBehavior: components.TypesProrationBehaviorCreateProrations,
        TargetPlanID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionChangePreviewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | Subscription ID                                                                                    |
| `body`                                                                                             | [components.DtoSubscriptionChangeRequest](../../models/components/dtosubscriptionchangerequest.md) | :heavy_check_mark:                                                                                 | Subscription change preview request                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PostSubscriptionsIDChangePreviewResponse](../../models/operations/postsubscriptionsidchangepreviewresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetEntitlements

Get all entitlements for a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/subscriptions/{id}/entitlements" method="get" path="/subscriptions/{id}/entitlements" -->
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

    res, err := s.Subscriptions.GetEntitlements(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Subscription ID                                          |
| `featureIds`                                             | []*string*                                               | :heavy_minus_sign:                                       | Feature IDs to filter by                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSubscriptionsIDEntitlementsResponse](../../models/operations/getsubscriptionsidentitlementsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetUpcomingGrants

Get upcoming credit grant applications for a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/subscriptions/{id}/grants/upcoming" method="get" path="/subscriptions/{id}/grants/upcoming" -->
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

    res, err := s.Subscriptions.GetUpcomingGrants(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListCreditGrantApplicationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Subscription ID                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSubscriptionsIDGrantsUpcomingResponse](../../models/operations/getsubscriptionsidgrantsupcomingresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Pause

Pause a subscription with the specified parameters

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/subscriptions/{id}/pause" method="post" path="/subscriptions/{id}/pause" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.Pause(ctx, "<id>", components.DtoPauseSubscriptionRequest{
        PauseMode: components.TypesPauseModePeriodEnd,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionPauseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | Subscription ID                                                                                  |
| `body`                                                                                           | [components.DtoPauseSubscriptionRequest](../../models/components/dtopausesubscriptionrequest.md) | :heavy_check_mark:                                                                               | Pause subscription request                                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PostSubscriptionsIDPauseResponse](../../models/operations/postsubscriptionsidpauseresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ListPauses

List all pauses for a subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/subscriptions/{id}/pauses" method="get" path="/subscriptions/{id}/pauses" -->
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

    res, err := s.Subscriptions.ListPauses(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListSubscriptionPausesResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Subscription ID                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSubscriptionsIDPausesResponse](../../models/operations/getsubscriptionsidpausesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Resume

Resume a paused subscription with the specified parameters

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/subscriptions/{id}/resume" method="post" path="/subscriptions/{id}/resume" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.Resume(ctx, "<id>", components.DtoResumeSubscriptionRequest{
        ResumeMode: components.TypesResumeModeScheduled,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionPauseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | Subscription ID                                                                                    |
| `body`                                                                                             | [components.DtoResumeSubscriptionRequest](../../models/components/dtoresumesubscriptionrequest.md) | :heavy_check_mark:                                                                                 | Resume subscription request                                                                        |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PostSubscriptionsIDResumeResponse](../../models/operations/postsubscriptionsidresumeresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |