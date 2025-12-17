# Prices

## Overview

### Available Operations

* [List](#list) - Get prices
* [Create](#create) - Create a new price
* [BulkCreate](#bulkcreate) - Create multiple prices in bulk
* [GetByID](#getbyid) - Get a price by ID
* [Update](#update) - Update a price
* [Delete](#delete) - Delete a price

## List

Get prices with the specified filter

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/prices" method="get" path="/prices" -->
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

    res, err := s.Prices.List(ctx, operations.GetPricesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListPricesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetPricesRequest](../../models/operations/getpricesrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetPricesResponse](../../models/operations/getpricesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a new price with the specified configuration. Supports both regular and price unit configurations.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/prices" method="post" path="/prices" -->
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

    res, err := s.Prices.Create(ctx, components.DtoCreatePriceRequest{
        BillingCadence: components.TypesBillingCadenceRecurring,
        BillingModel: components.TypesBillingModelTiered,
        BillingPeriod: components.TypesBillingPeriodDaily,
        Currency: "Iceland Krona",
        EntityID: "<id>",
        EntityType: components.TypesPriceEntityTypeCostsheet,
        InvoiceCadence: components.TypesInvoiceCadenceAdvance,
        PriceUnitType: components.TypesPriceUnitTypeCustom,
        Type: components.TypesPriceTypeUsage,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPriceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.DtoCreatePriceRequest](../../models/components/dtocreatepricerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PostPricesResponse](../../models/operations/postpricesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## BulkCreate

Create multiple prices with the specified configurations. Supports both regular and price unit configurations.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/prices/bulk" method="post" path="/prices/bulk" -->
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

    res, err := s.Prices.BulkCreate(ctx, components.DtoCreateBulkPriceRequest{
        Items: []components.DtoCreatePriceRequest{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreateBulkPriceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.DtoCreateBulkPriceRequest](../../models/components/dtocreatebulkpricerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PostPricesBulkResponse](../../models/operations/postpricesbulkresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByID

Get a price by ID with expanded meter and price unit information

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/prices/{id}" method="get" path="/prices/{id}" -->
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

    res, err := s.Prices.GetByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPriceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Price ID                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPricesIDResponse](../../models/operations/getpricesidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a price with the specified configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/prices/{id}" method="put" path="/prices/{id}" -->
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

    res, err := s.Prices.Update(ctx, "<id>", components.DtoUpdatePriceRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPriceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | Price ID                                                                             |
| `body`                                                                               | [components.DtoUpdatePriceRequest](../../models/components/dtoupdatepricerequest.md) | :heavy_check_mark:                                                                   | Price configuration                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PutPricesIDResponse](../../models/operations/putpricesidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a price

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/prices/{id}" method="delete" path="/prices/{id}" -->
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

    res, err := s.Prices.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Price ID                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeletePricesIDResponse](../../models/operations/deletepricesidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |