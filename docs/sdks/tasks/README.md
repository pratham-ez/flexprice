# Tasks

## Overview

### Available Operations

* [get_tasks](#get_tasks) - List tasks
* [post_tasks](#post_tasks) - Create a new task
* [get_tasks_result](#get_tasks_result) - Get task processing result
* [get_tasks_id_](#get_tasks_id_) - Get a task
* [put_tasks_id_status](#put_tasks_id_status) - Update task status

## get_tasks

List tasks with optional filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/tasks" method="get" path="/tasks" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tasks.get_tasks()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `created_by`                                                                             | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `end_time`                                                                               | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `entity_type`                                                                            | [Optional[operations.GetTasksEntityType]](../../models/operations/gettasksentitytype.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `expand`                                                                                 | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `limit`                                                                                  | *Optional[int]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `offset`                                                                                 | *Optional[int]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `order`                                                                                  | [Optional[operations.GetTasksOrder]](../../models/operations/gettasksorder.md)           | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `scheduled_task_id`                                                                      | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `sort`                                                                                   | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `start_time`                                                                             | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `status`                                                                                 | [Optional[operations.GetTasksStatus]](../../models/operations/gettasksstatus.md)         | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `task_status`                                                                            | [Optional[operations.TaskStatus]](../../models/operations/taskstatus.md)                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `task_type`                                                                              | [Optional[operations.TaskType]](../../models/operations/tasktype.md)                     | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `retries`                                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                         | :heavy_minus_sign:                                                                       | Configuration to override the default retry behavior of the client.                      |

### Response

**[components.DtoListTasksResponse](../../models/components/dtolisttasksresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_tasks

Create a new task for processing files asynchronously

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/tasks" method="post" path="/tasks" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tasks.post_tasks(entity_type="CUSTOMERS", file_type="JSON", file_url="https://juicy-fundraising.biz/", task_type="IMPORT")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `entity_type`                                                            | [components.TypesEntityType](../../models/components/typesentitytype.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `file_type`                                                              | [components.TypesFileType](../../models/components/typesfiletype.md)     | :heavy_check_mark:                                                       | N/A                                                                      |
| `file_url`                                                               | *str*                                                                    | :heavy_check_mark:                                                       | N/A                                                                      |
| `task_type`                                                              | [components.TypesTaskType](../../models/components/typestasktype.md)     | :heavy_check_mark:                                                       | N/A                                                                      |
| `file_name`                                                              | *Optional[str]*                                                          | :heavy_minus_sign:                                                       | N/A                                                                      |
| `metadata`                                                               | Dict[str, *Any*]                                                         | :heavy_minus_sign:                                                       | N/A                                                                      |
| `retries`                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)         | :heavy_minus_sign:                                                       | Configuration to override the default retry behavior of the client.      |

### Response

**[components.DtoTaskResponse](../../models/components/dtotaskresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_tasks_result

Get the result of a task processing workflow

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/tasks/result" method="get" path="/tasks/result" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tasks.get_tasks_result(workflow_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `workflow_id`                                                       | *str*                                                               | :heavy_check_mark:                                                  | Workflow ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.ModelsTemporalWorkflowResult](../../models/components/modelstemporalworkflowresult.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_tasks_id_

Get a task by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/tasks/{id}" method="get" path="/tasks/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tasks.get_tasks_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Task ID                                                             |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoTaskResponse](../../models/components/dtotaskresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_tasks_id_status

Update a task's status

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/tasks/{id}/status" method="put" path="/tasks/{id}/status" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tasks.put_tasks_id_status(id="<id>", task_status="COMPLETED")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `id`                                                                     | *str*                                                                    | :heavy_check_mark:                                                       | Task ID                                                                  |
| `task_status`                                                            | [components.TypesTaskStatus](../../models/components/typestaskstatus.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `retries`                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)         | :heavy_minus_sign:                                                       | Configuration to override the default retry behavior of the client.      |

### Response

**[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |