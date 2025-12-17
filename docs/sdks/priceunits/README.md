# PriceUnits

## Overview

### Available Operations

* [List](#list) - List price units
* [Create](#create) - Create a new price unit
* [GetByCode](#getbycode) - Get a price unit by code
* [Search](#search) - List price units by filter
* [GetByID](#getbyid) - Get a price unit by ID
* [Update](#update) - Update a price unit
* [Archive](#archive) - Archive a price unit

## List

Get a paginated list of price units with optional filtering

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/prices/units" method="get" path="/prices/units" -->
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

    res, err := s.PriceUnits.List(ctx, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListPriceUnitsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `status`                                                 | **string*                                                | :heavy_minus_sign:                                       | Filter by status                                         |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | Limit number of results                                  |
| `offset`                                                 | **int64*                                                 | :heavy_minus_sign:                                       | Offset for pagination                                    |
| `sort`                                                   | **string*                                                | :heavy_minus_sign:                                       | Sort field                                               |
| `order`                                                  | **string*                                                | :heavy_minus_sign:                                       | Sort order (asc/desc)                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPricesUnitsResponse](../../models/operations/getpricesunitsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a new price unit with the provided details

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/prices/units" method="post" path="/prices/units" -->
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

    res, err := s.PriceUnits.Create(ctx, components.DtoCreatePriceUnitRequest{
        BaseCurrency: "<value>",
        Code: "<value>",
        ConversionRate: "<value>",
        Name: "<value>",
        Symbol: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPriceUnitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.DtoCreatePriceUnitRequest](../../models/components/dtocreatepriceunitrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PostPricesUnitsResponse](../../models/operations/postpricesunitsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByCode

Get a price unit by code

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/prices/units/code/{code}" method="get" path="/prices/units/code/{code}" -->
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

    res, err := s.PriceUnits.GetByCode(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPriceUnitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `code`                                                   | *string*                                                 | :heavy_check_mark:                                       | Price unit code                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPricesUnitsCodeCodeResponse](../../models/operations/getpricesunitscodecoderesponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Search

List price units by filter

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/prices/units/search" method="post" path="/prices/units/search" -->
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

    res, err := s.PriceUnits.Search(ctx, components.PriceunitPriceUnitFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListPriceUnitsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.PriceunitPriceUnitFilter](../../models/components/priceunitpriceunitfilter.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PostPricesUnitsSearchResponse](../../models/operations/postpricesunitssearchresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByID

Get a price unit by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/prices/units/{id}" method="get" path="/prices/units/{id}" -->
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

    res, err := s.PriceUnits.GetByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPriceUnitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Price unit ID                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPricesUnitsIDResponse](../../models/operations/getpricesunitsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update an existing price unit with the provided details. Only name, symbol, precision, and conversion_rate can be updated. Status changes are not allowed.

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/prices/units/{id}" method="put" path="/prices/units/{id}" -->
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

    res, err := s.PriceUnits.Update(ctx, "<id>", components.DtoUpdatePriceUnitRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPriceUnitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `id`                                                                                         | *string*                                                                                     | :heavy_check_mark:                                                                           | Price unit ID                                                                                |
| `body`                                                                                       | [components.DtoUpdatePriceUnitRequest](../../models/components/dtoupdatepriceunitrequest.md) | :heavy_check_mark:                                                                           | Price unit details to update                                                                 |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PutPricesUnitsIDResponse](../../models/operations/putpricesunitsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Archive

Archive an existing price unit. The unit will be marked as archived and cannot be used in new prices.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/prices/units/{id}" method="delete" path="/prices/units/{id}" -->
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

    res, err := s.PriceUnits.Archive(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Price unit ID                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeletePricesUnitsIDResponse](../../models/operations/deletepricesunitsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |