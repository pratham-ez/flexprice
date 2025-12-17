# Customers

## Overview

### Available Operations

* [List](#list) - Get customers
* [Create](#create) - Create a customer
* [GetByLookupKey](#getbylookupkey) - Get a customer by lookup key
* [Search](#search) - List customers by filter
* [GetUsageSummary](#getusagesummary) - Get customer usage summary
* [GetByID](#getbyid) - Get a customer
* [Update](#update) - Update a customer
* [Delete](#delete) - Delete a customer
* [GetEntitlements](#getentitlements) - Get customer entitlements
* [GetUpcomingGrants](#getupcominggrants) - Get upcoming credit grant applications

## List

Get customers

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/customers" method="get" path="/customers" -->
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

    res, err := s.Customers.List(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListCustomersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |
| `customerIds`                                                                                       | []*string*                                                                                          | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `email`                                                                                             | **string*                                                                                           | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `endTime`                                                                                           | **string*                                                                                           | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `expand`                                                                                            | **string*                                                                                           | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `externalID`                                                                                        | **string*                                                                                           | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `externalIds`                                                                                       | []*string*                                                                                          | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `limit`                                                                                             | **int64*                                                                                            | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `offset`                                                                                            | **int64*                                                                                            | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `order`                                                                                             | [*operations.GetCustomersQueryParamOrder](../../models/operations/getcustomersqueryparamorder.md)   | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `parentCustomerIds`                                                                                 | []*string*                                                                                          | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `startTime`                                                                                         | **string*                                                                                           | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `status`                                                                                            | [*operations.GetCustomersQueryParamStatus](../../models/operations/getcustomersqueryparamstatus.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |

### Response

**[*operations.GetCustomersResponse](../../models/operations/getcustomersresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a customer

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/customers" method="post" path="/customers" -->
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

    res, err := s.Customers.Create(ctx, components.DtoCreateCustomerRequest{
        ExternalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.DtoCreateCustomerRequest](../../models/components/dtocreatecustomerrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PostCustomersResponse](../../models/operations/postcustomersresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByLookupKey

Get a customer by lookup key (external_id)

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/customers/lookup/{lookup_key}" method="get" path="/customers/lookup/{lookup_key}" -->
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

    res, err := s.Customers.GetByLookupKey(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `lookupKey`                                              | *string*                                                 | :heavy_check_mark:                                       | Customer Lookup Key (external_id)                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCustomersLookupLookupKeyResponse](../../models/operations/getcustomerslookuplookupkeyresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Search

List customers by filter

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/customers/search" method="post" path="/customers/search" -->
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

    res, err := s.Customers.Search(ctx, components.TypesCustomerFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListCustomersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.TypesCustomerFilter](../../models/components/typescustomerfilter.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.PostCustomersSearchResponse](../../models/operations/postcustomerssearchresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetUsageSummary

Get customer usage summary by customer_id or customer_lookup_key (external_customer_id)

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/customers/usage" method="get" path="/customers/usage" -->
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

    res, err := s.Customers.GetUsageSummary(ctx, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCustomerUsageSummaryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `customerID`                                             | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `customerLookupKey`                                      | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `featureIds`                                             | []*string*                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `featureLookupKeys`                                      | []*string*                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `subscriptionIds`                                        | []*string*                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCustomersUsageResponse](../../models/operations/getcustomersusageresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByID

Get a customer

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/customers/{id}" method="get" path="/customers/{id}" -->
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

    res, err := s.Customers.GetByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Customer ID                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCustomersIDResponse](../../models/operations/getcustomersidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a customer

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/customers/{id}" method="put" path="/customers/{id}" -->
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

    res, err := s.Customers.Update(ctx, "<id>", components.DtoUpdateCustomerRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `id`                                                                                       | *string*                                                                                   | :heavy_check_mark:                                                                         | Customer ID                                                                                |
| `body`                                                                                     | [components.DtoUpdateCustomerRequest](../../models/components/dtoupdatecustomerrequest.md) | :heavy_check_mark:                                                                         | Customer                                                                                   |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PutCustomersIDResponse](../../models/operations/putcustomersidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a customer

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/customers/{id}" method="delete" path="/customers/{id}" -->
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

    res, err := s.Customers.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Customer ID                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteCustomersIDResponse](../../models/operations/deletecustomersidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetEntitlements

Get customer entitlements

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/customers/{id}/entitlements" method="get" path="/customers/{id}/entitlements" -->
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

    res, err := s.Customers.GetEntitlements(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCustomerEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Customer ID                                              |
| `featureIds`                                             | []*string*                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `subscriptionIds`                                        | []*string*                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCustomersIDEntitlementsResponse](../../models/operations/getcustomersidentitlementsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetUpcomingGrants

Get upcoming credit grant applications for a customer

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/customers/{id}/grants/upcoming" method="get" path="/customers/{id}/grants/upcoming" -->
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

    res, err := s.Customers.GetUpcomingGrants(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Customer ID                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCustomersIDGrantsUpcomingResponse](../../models/operations/getcustomersidgrantsupcomingresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |