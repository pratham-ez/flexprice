# Features

## Overview

### Available Operations

* [List](#list) - List features
* [Create](#create) - Create a new feature
* [Search](#search) - List features by filter
* [GetByID](#getbyid) - Get a feature by ID
* [Update](#update) - Update a feature
* [Delete](#delete) - Delete a feature

## List

List features with optional filtering

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/features" method="get" path="/features" -->
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

    res, err := s.Features.List(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListFeaturesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `endTime`                                                                                         | **string*                                                                                         | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `expand`                                                                                          | **string*                                                                                         | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `featureIds`                                                                                      | []*string*                                                                                        | :heavy_minus_sign:                                                                                | Feature specific filters                                                                          |
| `limit`                                                                                           | **int64*                                                                                          | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `lookupKey`                                                                                       | **string*                                                                                         | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `lookupKeys`                                                                                      | []*string*                                                                                        | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `meterIds`                                                                                        | []*string*                                                                                        | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `nameContains`                                                                                    | **string*                                                                                         | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `offset`                                                                                          | **int64*                                                                                          | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `order`                                                                                           | [*operations.GetFeaturesQueryParamOrder](../../models/operations/getfeaturesqueryparamorder.md)   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `startTime`                                                                                       | **string*                                                                                         | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `status`                                                                                          | [*operations.GetFeaturesQueryParamStatus](../../models/operations/getfeaturesqueryparamstatus.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.GetFeaturesResponse](../../models/operations/getfeaturesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a new feature

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/features" method="post" path="/features" -->
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

    res, err := s.Features.Create(ctx, components.DtoCreateFeatureRequest{
        Meter: &components.DtoCreateMeterRequest{
            Aggregation: components.MeterAggregation{},
            EventName: "api_request",
            Name: "API Usage Meter",
            ResetUsage: components.TypesResetUsageNever,
        },
        Name: "<value>",
        Type: components.TypesFeatureTypeMetered,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoFeatureResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.DtoCreateFeatureRequest](../../models/components/dtocreatefeaturerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PostFeaturesResponse](../../models/operations/postfeaturesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Search

List features by filter

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/features/search" method="post" path="/features/search" -->
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

    res, err := s.Features.Search(ctx, components.TypesFeatureFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListFeaturesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.TypesFeatureFilter](../../models/components/typesfeaturefilter.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.PostFeaturesSearchResponse](../../models/operations/postfeaturessearchresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByID

Get a feature by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/features/{id}" method="get" path="/features/{id}" -->
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

    res, err := s.Features.GetByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoFeatureResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Feature ID                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetFeaturesIDResponse](../../models/operations/getfeaturesidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a feature by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/features/{id}" method="put" path="/features/{id}" -->
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

    res, err := s.Features.Update(ctx, "<id>", components.DtoUpdateFeatureRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoFeatureResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Feature ID                                                                               |
| `body`                                                                                   | [components.DtoUpdateFeatureRequest](../../models/components/dtoupdatefeaturerequest.md) | :heavy_check_mark:                                                                       | Feature update data                                                                      |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PutFeaturesIDResponse](../../models/operations/putfeaturesidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a feature by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/features/{id}" method="delete" path="/features/{id}" -->
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

    res, err := s.Features.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Feature ID                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteFeaturesIDResponse](../../models/operations/deletefeaturesidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |