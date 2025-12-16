# EntityIntegrationMappings

## Overview

### Available Operations

* [List](#list) - List entity integration mappings
* [Create](#create) - Create entity integration mapping
* [GetByID](#getbyid) - Get entity integration mapping
* [Delete](#delete) - Delete entity integration mapping

## List

Retrieve a list of entity integration mappings with optional filtering

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/entity-integration-mappings" method="get" path="/entity-integration-mappings" -->
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

    res, err := s.EntityIntegrationMappings.List(ctx, nil, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListEntityIntegrationMappingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `entityID`                                               | **string*                                                | :heavy_minus_sign:                                       | Filter by FlexPrice entity ID                            |
| `entityType`                                             | **string*                                                | :heavy_minus_sign:                                       | Filter by entity type                                    |
| `providerType`                                           | **string*                                                | :heavy_minus_sign:                                       | Filter by provider type                                  |
| `providerEntityID`                                       | **string*                                                | :heavy_minus_sign:                                       | Filter by provider entity ID                             |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | Number of results to return (default: 20, max: 100)      |
| `offset`                                                 | **int64*                                                 | :heavy_minus_sign:                                       | Pagination offset (default: 0)                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetEntityIntegrationMappingsResponse](../../models/operations/getentityintegrationmappingsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 401                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a new entity integration mapping

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/entity-integration-mappings" method="post" path="/entity-integration-mappings" -->
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

    res, err := s.EntityIntegrationMappings.Create(ctx, components.DtoCreateEntityIntegrationMappingRequest{
        EntityID: "<id>",
        EntityType: components.TypesIntegrationEntityTypeAddon,
        ProviderEntityID: "<id>",
        ProviderType: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoEntityIntegrationMappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [components.DtoCreateEntityIntegrationMappingRequest](../../models/components/dtocreateentityintegrationmappingrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.PostEntityIntegrationMappingsResponse](../../models/operations/postentityintegrationmappingsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 401, 409                 | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByID

Retrieve a specific entity integration mapping by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/entity-integration-mappings/{id}" method="get" path="/entity-integration-mappings/{id}" -->
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

    res, err := s.EntityIntegrationMappings.GetByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoEntityIntegrationMappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Entity integration mapping ID                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetEntityIntegrationMappingsIDResponse](../../models/operations/getentityintegrationmappingsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 401, 404                 | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete an entity integration mapping

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/entity-integration-mappings/{id}" method="delete" path="/entity-integration-mappings/{id}" -->
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

    res, err := s.EntityIntegrationMappings.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Entity integration mapping ID                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteEntityIntegrationMappingsIDResponse](../../models/operations/deleteentityintegrationmappingsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 401, 404                 | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |