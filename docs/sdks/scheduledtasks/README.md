# ScheduledTasks

## Overview

### Available Operations

* [List](#list) - List scheduled tasks
* [Create](#create) - Create a scheduled task
* [Get](#get) - Get a scheduled task
* [Update](#update) - Update a scheduled task
* [Delete](#delete) - Delete a scheduled task
* [TriggerRun](#triggerrun) - Trigger force run

## List

Get a list of scheduled tasks with optional filters

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/tasks/scheduled" method="get" path="/tasks/scheduled" -->
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

    res, err := s.ScheduledTasks.List(ctx, nil, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListScheduledTasksResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | Limit                                                    |
| `offset`                                                 | **int64*                                                 | :heavy_minus_sign:                                       | Offset                                                   |
| `connectionID`                                           | **string*                                                | :heavy_minus_sign:                                       | Filter by connection ID                                  |
| `entityType`                                             | **string*                                                | :heavy_minus_sign:                                       | Filter by entity type                                    |
| `interval`                                               | **string*                                                | :heavy_minus_sign:                                       | Filter by interval                                       |
| `enabled`                                                | **bool*                                                  | :heavy_minus_sign:                                       | Filter by enabled status                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTasksScheduledResponse](../../models/operations/gettasksscheduledresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a new scheduled task for data export

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/tasks/scheduled" method="post" path="/tasks/scheduled" -->
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

    res, err := s.ScheduledTasks.Create(ctx, components.DtoCreateScheduledTaskRequest{
        ConnectionID: "<id>",
        EntityType: components.TypesScheduledTaskEntityTypeInvoice,
        Interval: components.TypesScheduledTaskIntervalCustom,
        JobConfig: components.TypesS3JobConfig{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoScheduledTaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.DtoCreateScheduledTaskRequest](../../models/components/dtocreatescheduledtaskrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PostTasksScheduledResponse](../../models/operations/posttasksscheduledresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a scheduled task by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/tasks/scheduled/{id}" method="get" path="/tasks/scheduled/{id}" -->
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

    res, err := s.ScheduledTasks.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoScheduledTaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Scheduled Task ID                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTasksScheduledIDResponse](../../models/operations/gettasksscheduledidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a scheduled task by ID - Only enabled field can be changed (pause/resume)

### Example Usage

<!-- UsageSnippet language="go" operationID="put_/tasks/scheduled/{id}" method="put" path="/tasks/scheduled/{id}" -->
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

    res, err := s.ScheduledTasks.Update(ctx, "<id>", components.DtoUpdateScheduledTaskRequest{
        Enabled: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoScheduledTaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `id`                                                                                                 | *string*                                                                                             | :heavy_check_mark:                                                                                   | Scheduled Task ID                                                                                    |
| `body`                                                                                               | [components.DtoUpdateScheduledTaskRequest](../../models/components/dtoupdatescheduledtaskrequest.md) | :heavy_check_mark:                                                                                   | Update request (enabled: true/false to pause/resume)                                                 |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PutTasksScheduledIDResponse](../../models/operations/puttasksscheduledidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Archive a scheduled task by ID (soft delete) - Sets status to archived and deletes from Temporal

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_/tasks/scheduled/{id}" method="delete" path="/tasks/scheduled/{id}" -->
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

    res, err := s.ScheduledTasks.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Scheduled Task ID                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteTasksScheduledIDResponse](../../models/operations/deletetasksscheduledidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## TriggerRun

Trigger a force run export immediately for a scheduled task with optional custom time range

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/tasks/scheduled/{id}/run" method="post" path="/tasks/scheduled/{id}/run" -->
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

    res, err := s.ScheduledTasks.TriggerRun(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTriggerForceRunResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |
| `id`                                                                                          | *string*                                                                                      | :heavy_check_mark:                                                                            | Scheduled Task ID                                                                             |
| `body`                                                                                        | [*components.DtoTriggerForceRunRequest](../../models/components/dtotriggerforcerunrequest.md) | :heavy_minus_sign:                                                                            | Optional start and end time for custom range                                                  |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |

### Response

**[*operations.PostTasksScheduledIDRunResponse](../../models/operations/posttasksscheduledidrunresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 404                      | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |