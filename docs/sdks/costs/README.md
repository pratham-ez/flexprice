# Costs

## Overview

### Available Operations

* [Create](#create) - Create a new costsheet
* [GetActive](#getactive) - Get active costsheet for tenant
* [GetAnalytics](#getanalytics) - Get combined revenue and cost analytics
* [Search](#search) - List costsheets by filter
* [GetByID](#getbyid) - Get a costsheet by ID
* [Update](#update) - Update a costsheet
* [Delete](#delete) - Delete a costsheet

## Create

Create a new costsheet with the specified name

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/costs" method="post" path="/costs" -->
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

    res, err := s.Costs.Create(ctx, components.DtoCreateCostsheetRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreateCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.DtoCreateCostsheetRequest](../../models/components/dtocreatecostsheetrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PostCostsResponse](../../models/operations/postcostsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 409                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetActive

Get the active costsheet for the current tenant

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/costs/active" method="get" path="/costs/active" -->
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

    res, err := s.Costs.GetActive(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCostsActiveResponse](../../models/operations/getcostsactiveresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 404                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetAnalytics

Retrieve combined analytics with ROI, margin, and detailed breakdowns. If start_time and end_time are not provided, defaults to last 7 days.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/costs/analytics" method="post" path="/costs/analytics" -->
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

    res, err := s.Costs.GetAnalytics(ctx, components.DtoGetCostAnalyticsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetDetailedCostAnalyticsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.DtoGetCostAnalyticsRequest](../../models/components/dtogetcostanalyticsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PostCostsAnalyticsResponse](../../models/operations/postcostsanalyticsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Search

List costsheet records by filter with POST body

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/costs/search" method="post" path="/costs/search" -->
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

    res, err := s.Costs.Search(ctx, components.CostsheetFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.CostsheetFilter](../../models/components/costsheetfilter.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.PostCostsSearchResponse](../../models/operations/postcostssearchresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByID

Get a costsheet by ID with optional price expansion

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/costs/{id}" method="get" path="/costs/{id}" -->
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

    res, err := s.Costs.GetByID(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `id`                                                      | *string*                                                  | :heavy_check_mark:                                        | Costsheet ID                                              |
| `expand`                                                  | **string*                                                 | :heavy_minus_sign:                                        | Comma-separated list of fields to expand (e.g., 'prices') |
| `opts`                                                    | [][operations.Option](../../models/operations/option.md)  | :heavy_minus_sign:                                        | The options for this request.                             |

### Response

**[*operations.GetCostsIDResponse](../../models/operations/getcostsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a costsheet with the specified configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/costs/{id}" method="put" path="/costs/{id}" -->
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

    res, err := s.Costs.Update(ctx, "<id>", components.DtoUpdateCostsheetRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoUpdateCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `id`                                                                                         | *string*                                                                                     | :heavy_check_mark:                                                                           | Costsheet ID                                                                                 |
| `body`                                                                                       | [components.DtoUpdateCostsheetRequest](../../models/components/dtoupdatecostsheetrequest.md) | :heavy_check_mark:                                                                           | Costsheet configuration                                                                      |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PutCostsIDResponse](../../models/operations/putcostsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404, 409                 | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Soft delete a costsheet by setting its status to deleted

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/costs/{id}" method="delete" path="/costs/{id}" -->
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

    res, err := s.Costs.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoDeleteCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Costsheet ID                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteCostsIDResponse](../../models/operations/deletecostsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |