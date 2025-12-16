# Groups

## Overview

### Available Operations

* [Create](#create) - Create a group
* [Search](#search) - Get groups
* [GetByID](#getbyid) - Get a group
* [Delete](#delete) - Delete a group

## Create

Create a new group for organizing entities (prices, plans, customers, etc.)

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/groups" method="post" path="/groups" -->
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

    res, err := s.Groups.Create(ctx, components.DtoCreateGroupRequest{
        EntityType: "<value>",
        LookupKey: "<value>",
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGroupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.DtoCreateGroupRequest](../../models/components/dtocreategrouprequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PostGroupsResponse](../../models/operations/postgroupsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Search

Get groups with optional filtering via query parameters

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/groups/search" method="post" path="/groups/search" -->
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

    res, err := s.Groups.Search(ctx, nil, nil, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListGroupsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `entityType`                                             | **string*                                                | :heavy_minus_sign:                                       | Filter by entity type (e.g., 'price')                    |
| `name`                                                   | **string*                                                | :heavy_minus_sign:                                       | Filter by group name (contains search)                   |
| `lookupKey`                                              | **string*                                                | :heavy_minus_sign:                                       | Filter by lookup key (exact match)                       |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | Number of items to return (default: 20)                  |
| `offset`                                                 | **int64*                                                 | :heavy_minus_sign:                                       | Number of items to skip (default: 0)                     |
| `sortBy`                                                 | **string*                                                | :heavy_minus_sign:                                       | Field to sort by (name, created_at, updated_at)          |
| `sortOrder`                                              | **string*                                                | :heavy_minus_sign:                                       | Sort order (asc, desc)                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostGroupsSearchResponse](../../models/operations/postgroupssearchresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetByID

Get a group by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/groups/{id}" method="get" path="/groups/{id}" -->
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

    res, err := s.Groups.GetByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGroupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Group ID                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetGroupsIDResponse](../../models/operations/getgroupsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a group and remove all entity associations

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/groups/{id}" method="delete" path="/groups/{id}" -->
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

    res, err := s.Groups.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Group ID                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteGroupsIDResponse](../../models/operations/deletegroupsidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |