# ScheduledTasks

## Overview

### Available Operations

* [get_tasks_scheduled](#get_tasks_scheduled) - List scheduled tasks
* [post_tasks_scheduled](#post_tasks_scheduled) - Create a scheduled task
* [post_tasks_scheduled_schedule_update_billing_period](#post_tasks_scheduled_schedule_update_billing_period) - Schedule update billing period
* [get_tasks_scheduled_id_](#get_tasks_scheduled_id_) - Get a scheduled task
* [put_tasks_scheduled_id_](#put_tasks_scheduled_id_) - Update a scheduled task
* [delete_tasks_scheduled_id_](#delete_tasks_scheduled_id_) - Delete a scheduled task
* [post_tasks_scheduled_id_run](#post_tasks_scheduled_id_run) - Trigger force run

## get_tasks_scheduled

Get a list of scheduled tasks with optional filters

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/tasks/scheduled" method="get" path="/tasks/scheduled" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.scheduled_tasks.get_tasks_scheduled()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `limit`                                                             | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | Limit                                                               |
| `offset`                                                            | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | Offset                                                              |
| `connection_id`                                                     | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by connection ID                                             |
| `entity_type`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by entity type                                               |
| `interval`                                                          | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by interval                                                  |
| `enabled`                                                           | *Optional[bool]*                                                    | :heavy_minus_sign:                                                  | Filter by enabled status                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoListScheduledTasksResponse](../../models/components/dtolistscheduledtasksresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_tasks_scheduled

Create a new scheduled task for data export

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/tasks/scheduled" method="post" path="/tasks/scheduled" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.scheduled_tasks.post_tasks_scheduled(connection_id="<id>", entity_type="credit_topups", interval="15MIN", job_config={})

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `connection_id`                                                                                    | *str*                                                                                              | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `entity_type`                                                                                      | [components.TypesScheduledTaskEntityType](../../models/components/typesscheduledtaskentitytype.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `interval`                                                                                         | [components.TypesScheduledTaskInterval](../../models/components/typesscheduledtaskinterval.md)     | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `job_config`                                                                                       | [components.TypesS3JobConfig](../../models/components/typess3jobconfig.md)                         | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `enabled`                                                                                          | *Optional[bool]*                                                                                   | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `retries`                                                                                          | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                   | :heavy_minus_sign:                                                                                 | Configuration to override the default retry behavior of the client.                                |

### Response

**[components.DtoScheduledTaskResponse](../../models/components/dtoscheduledtaskresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_tasks_scheduled_schedule_update_billing_period

Schedule an update billing period workflow

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/tasks/scheduled/schedule-update-billing-period" method="post" path="/tasks/scheduled/schedule-update-billing-period" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.scheduled_tasks.post_tasks_scheduled_schedule_update_billing_period(request={})

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `request`                                                                                                                                          | [operations.PostTasksScheduledScheduleUpdateBillingPeriodRequest](../../models/operations/posttasksscheduledscheduleupdatebillingperiodrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `retries`                                                                                                                                          | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                   | :heavy_minus_sign:                                                                                                                                 | Configuration to override the default retry behavior of the client.                                                                                |

### Response

**[operations.PostTasksScheduledScheduleUpdateBillingPeriodResponse](../../models/operations/posttasksscheduledscheduleupdatebillingperiodresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_tasks_scheduled_id_

Get a scheduled task by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/tasks/scheduled/{id}" method="get" path="/tasks/scheduled/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.scheduled_tasks.get_tasks_scheduled_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Scheduled Task ID                                                   |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoScheduledTaskResponse](../../models/components/dtoscheduledtaskresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_tasks_scheduled_id_

Update a scheduled task by ID - Only enabled field can be changed (pause/resume)

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/tasks/scheduled/{id}" method="put" path="/tasks/scheduled/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.scheduled_tasks.put_tasks_scheduled_id_(id="<id>", enabled=True)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Scheduled Task ID                                                   |
| `enabled`                                                           | *bool*                                                              | :heavy_check_mark:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoScheduledTaskResponse](../../models/components/dtoscheduledtaskresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_tasks_scheduled_id_

Archive a scheduled task by ID (soft delete) - Sets status to archived and deletes from Temporal

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/tasks/scheduled/{id}" method="delete" path="/tasks/scheduled/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    flex_price.scheduled_tasks.delete_tasks_scheduled_id_(id="<id>")

    # Use the SDK ...

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Scheduled Task ID                                                   |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_tasks_scheduled_id_run

Trigger a force run export immediately for a scheduled task with optional custom time range

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/tasks/scheduled/{id}/run" method="post" path="/tasks/scheduled/{id}/run" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.scheduled_tasks.post_tasks_scheduled_id_run(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Scheduled Task ID                                                   |
| `end_time`                                                          | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `start_time`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoTriggerForceRunResponse](../../models/components/dtotriggerforcerunresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |