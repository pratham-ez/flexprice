# Addons

## Overview

### Available Operations

* [List](#list) - List addons
* [Create](#create) - Create addon
* [GetByLookupKey](#getbylookupkey) - Get addon by lookup key
* [Search](#search) - List addons by filter
* [Get](#get) - Get addon
* [Update](#update) - Update addon
* [Delete](#delete) - Delete addon

## List

Get addons with optional filtering

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/addons" method="get" path="/addons" -->
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

    res, err := s.Addons.List(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListAddonsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                     | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ctx`                                                         | [context.Context](https://pkg.go.dev/context#Context)         | :heavy_check_mark:                                            | The context to use for the request.                           |
| `addonIds`                                                    | []*string*                                                    | :heavy_minus_sign:                                            | N/A                                                           |
| `addonType`                                                   | [*operations.AddonType](../../models/operations/addontype.md) | :heavy_minus_sign:                                            | N/A                                                           |
| `endTime`                                                     | **string*                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `expand`                                                      | **string*                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `limit`                                                       | **int64*                                                      | :heavy_minus_sign:                                            | N/A                                                           |
| `lookupKeys`                                                  | []*string*                                                    | :heavy_minus_sign:                                            | N/A                                                           |
| `offset`                                                      | **int64*                                                      | :heavy_minus_sign:                                            | N/A                                                           |
| `order`                                                       | [*operations.Order](../../models/operations/order.md)         | :heavy_minus_sign:                                            | N/A                                                           |
| `startTime`                                                   | **string*                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `status`                                                      | [*operations.Status](../../models/operations/status.md)       | :heavy_minus_sign:                                            | N/A                                                           |
| `opts`                                                        | [][operations.Option](../../models/operations/option.md)      | :heavy_minus_sign:                                            | The options for this request.                                 |

### Response

**[*operations.GetAddonsResponse](../../models/operations/getaddonsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a new addon

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/addons" method="post" path="/addons" -->
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

    res, err := s.Addons.Create(ctx, components.DtoCreateAddonRequest{
        LookupKey: "<value>",
        Name: "<value>",
        Type: components.TypesAddonTypeOnetime,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreateAddonResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.DtoCreateAddonRequest](../../models/components/dtocreateaddonrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PostAddonsResponse](../../models/operations/postaddonsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByLookupKey

Get an addon by lookup key

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/addons/lookup/{lookup_key}" method="get" path="/addons/lookup/{lookup_key}" -->
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

    res, err := s.Addons.GetByLookupKey(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoAddonResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `lookupKey`                                              | *string*                                                 | :heavy_check_mark:                                       | Addon Lookup Key                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAddonsLookupLookupKeyResponse](../../models/operations/getaddonslookuplookupkeyresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Search

List addons by filter

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/addons/search" method="post" path="/addons/search" -->
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

    res, err := s.Addons.Search(ctx, components.TypesAddonFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListAddonsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.TypesAddonFilter](../../models/components/typesaddonfilter.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.PostAddonsSearchResponse](../../models/operations/postaddonssearchresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get an addon by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/addons/{id}" method="get" path="/addons/{id}" -->
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

    res, err := s.Addons.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoAddonResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Addon ID                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAddonsIDResponse](../../models/operations/getaddonsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update an existing addon

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/addons/{id}" method="put" path="/addons/{id}" -->
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

    res, err := s.Addons.Update(ctx, "<id>", components.DtoUpdateAddonRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoAddonResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | Addon ID                                                                             |
| `body`                                                                               | [components.DtoUpdateAddonRequest](../../models/components/dtoupdateaddonrequest.md) | :heavy_check_mark:                                                                   | Update Addon Request                                                                 |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PutAddonsIDResponse](../../models/operations/putaddonsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete an addon

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/addons/{id}" method="delete" path="/addons/{id}" -->
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

    res, err := s.Addons.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Addon ID                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteAddonsIDResponse](../../models/operations/deleteaddonsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |