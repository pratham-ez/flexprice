# Plans

## Overview

### Available Operations

* [List](#list) - Get plans
* [Create](#create) - Create a new plan
* [Search](#search) - List plans by filter
* [GetByID](#getbyid) - Get a plan
* [Update](#update) - Update a plan
* [Delete](#delete) - Delete a plan
* [SyncSubscriptions](#syncsubscriptions) - Synchronize plan prices
* [GetEntitlements](#getentitlements) - Get plan entitlements

## List

Get plans with optional filtering

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/plans" method="get" path="/plans" -->
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

    res, err := s.Plans.List(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListPlansResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `endTime`                                                                                   | **string*                                                                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `expand`                                                                                    | **string*                                                                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `limit`                                                                                     | **int64*                                                                                    | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `lookupKey`                                                                                 | **string*                                                                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `offset`                                                                                    | **int64*                                                                                    | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `order`                                                                                     | [*operations.GetPlansQueryParamOrder](../../models/operations/getplansqueryparamorder.md)   | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `planIds`                                                                                   | []*string*                                                                                  | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `startTime`                                                                                 | **string*                                                                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `status`                                                                                    | [*operations.GetPlansQueryParamStatus](../../models/operations/getplansqueryparamstatus.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.GetPlansResponse](../../models/operations/getplansresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a new plan with the specified configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/plans" method="post" path="/plans" -->
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

    res, err := s.Plans.Create(ctx, components.DtoCreatePlanRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPlanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.DtoCreatePlanRequest](../../models/components/dtocreateplanrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.PostPlansResponse](../../models/operations/postplansresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Search

List plans by filter

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/plans/search" method="post" path="/plans/search" -->
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

    res, err := s.Plans.Search(ctx, components.TypesPlanFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListPlansResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.TypesPlanFilter](../../models/components/typesplanfilter.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.PostPlansSearchResponse](../../models/operations/postplanssearchresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByID

Get a plan by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/plans/{id}" method="get" path="/plans/{id}" -->
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

    res, err := s.Plans.GetByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPlanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Plan ID                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPlansIDResponse](../../models/operations/getplansidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a plan by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/plans/{id}" method="put" path="/plans/{id}" -->
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

    res, err := s.Plans.Update(ctx, "<id>", components.DtoUpdatePlanRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPlanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `id`                                                                               | *string*                                                                           | :heavy_check_mark:                                                                 | Plan ID                                                                            |
| `body`                                                                             | [components.DtoUpdatePlanRequest](../../models/components/dtoupdateplanrequest.md) | :heavy_check_mark:                                                                 | Plan update                                                                        |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.PutPlansIDResponse](../../models/operations/putplansidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a plan by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/plans/{id}" method="delete" path="/plans/{id}" -->
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

    res, err := s.Plans.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GinH != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Plan ID                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeletePlansIDResponse](../../models/operations/deleteplansidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## SyncSubscriptions

Synchronize current plan prices with all existing active subscriptions

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/plans/{id}/sync/subscriptions" method="post" path="/plans/{id}/sync/subscriptions" -->
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

    res, err := s.Plans.SyncSubscriptions(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ModelsTemporalWorkflowResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Plan ID                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostPlansIDSyncSubscriptionsResponse](../../models/operations/postplansidsyncsubscriptionsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404, 422                 | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetEntitlements

Get all entitlements for a plan

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/plans/{id}/entitlements" method="get" path="/plans/{id}/entitlements" -->
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

    res, err := s.Plans.GetEntitlements(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPlanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Plan ID                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPlansIDEntitlementsResponse](../../models/operations/getplansidentitlementsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |