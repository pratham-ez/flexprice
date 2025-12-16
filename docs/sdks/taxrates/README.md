# TaxRates

## Overview

### Available Operations

* [GetAll](#getall) - Get tax rates
* [Create](#create) - Create a tax rate
* [Get](#get) - Get a tax rate
* [Update](#update) - Update a tax rate
* [Delete](#delete) - Delete a tax rate

## GetAll

Get tax rates

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/taxes/rates" method="get" path="/taxes/rates" -->
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

    res, err := s.TaxRates.GetAll(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxRateResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `endTime`                                                                                             | **string*                                                                                             | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `expand`                                                                                              | **string*                                                                                             | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `limit`                                                                                               | **int64*                                                                                              | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `offset`                                                                                              | **int64*                                                                                              | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `order`                                                                                               | [*operations.GetTaxesRatesQueryParamOrder](../../models/operations/gettaxesratesqueryparamorder.md)   | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `scope`                                                                                               | [*operations.QueryParamScope](../../models/operations/queryparamscope.md)                             | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `startTime`                                                                                           | **string*                                                                                             | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `status`                                                                                              | [*operations.GetTaxesRatesQueryParamStatus](../../models/operations/gettaxesratesqueryparamstatus.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `taxrateCodes`                                                                                        | []*string*                                                                                            | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `taxrateIds`                                                                                          | []*string*                                                                                            | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |

### Response

**[*operations.GetTaxesRatesResponse](../../models/operations/gettaxesratesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a tax rate

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/taxes/rates" method="post" path="/taxes/rates" -->
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

    res, err := s.TaxRates.Create(ctx, components.DtoCreateTaxRateRequest{
        Code: "<value>",
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxRateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.DtoCreateTaxRateRequest](../../models/components/dtocreatetaxraterequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PostTaxesRatesResponse](../../models/operations/posttaxesratesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a tax rate

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/taxes/rates/{id}" method="get" path="/taxes/rates/{id}" -->
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

    res, err := s.TaxRates.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxRateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Tax rate ID                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTaxesRatesIDResponse](../../models/operations/gettaxesratesidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a tax rate

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/taxes/rates/{id}" method="put" path="/taxes/rates/{id}" -->
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

    res, err := s.TaxRates.Update(ctx, "<id>", components.DtoUpdateTaxRateRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxRateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Tax rate ID                                                                              |
| `body`                                                                                   | [components.DtoUpdateTaxRateRequest](../../models/components/dtoupdatetaxraterequest.md) | :heavy_check_mark:                                                                       | Tax rate to update                                                                       |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PutTaxesRatesIDResponse](../../models/operations/puttaxesratesidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a tax rate

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/taxes/rates/{id}" method="delete" path="/taxes/rates/{id}" -->
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

    res, err := s.TaxRates.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Tax rate ID                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteTaxesRatesIDResponse](../../models/operations/deletetaxesratesidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |