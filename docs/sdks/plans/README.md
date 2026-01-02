# Plans

## Overview

### Available Operations

* [get_plans](#get_plans) - Get plans
* [post_plans](#post_plans) - Create a new plan
* [post_plans_search](#post_plans_search) - List plans by filter
* [get_plans_id_](#get_plans_id_) - Get a plan
* [put_plans_id_](#put_plans_id_) - Update a plan
* [delete_plans_id_](#delete_plans_id_) - Delete a plan
* [post_plans_id_sync_subscriptions](#post_plans_id_sync_subscriptions) - Synchronize plan prices

## get_plans

Get plans with optional filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/plans" method="get" path="/plans" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.plans.get_plans()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `end_time`                                                                       | *Optional[str]*                                                                  | :heavy_minus_sign:                                                               | N/A                                                                              |
| `expand`                                                                         | *Optional[str]*                                                                  | :heavy_minus_sign:                                                               | N/A                                                                              |
| `limit`                                                                          | *Optional[int]*                                                                  | :heavy_minus_sign:                                                               | N/A                                                                              |
| `lookup_key`                                                                     | *Optional[str]*                                                                  | :heavy_minus_sign:                                                               | N/A                                                                              |
| `offset`                                                                         | *Optional[int]*                                                                  | :heavy_minus_sign:                                                               | N/A                                                                              |
| `order`                                                                          | [Optional[operations.GetPlansOrder]](../../models/operations/getplansorder.md)   | :heavy_minus_sign:                                                               | N/A                                                                              |
| `plan_ids`                                                                       | List[*str*]                                                                      | :heavy_minus_sign:                                                               | N/A                                                                              |
| `start_time`                                                                     | *Optional[str]*                                                                  | :heavy_minus_sign:                                                               | N/A                                                                              |
| `status`                                                                         | [Optional[operations.GetPlansStatus]](../../models/operations/getplansstatus.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `retries`                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                 | :heavy_minus_sign:                                                               | Configuration to override the default retry behavior of the client.              |

### Response

**[components.DtoListPlansResponse](../../models/components/dtolistplansresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_plans

Create a new plan with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/plans" method="post" path="/plans" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.plans.post_plans(name="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `name`                                                              | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `description`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `display_order`                                                     | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `lookup_key`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `metadata`                                                          | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPlanResponse](../../models/components/dtoplanresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_plans_search

List plans by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/plans/search" method="post" path="/plans/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.plans.post_plans_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `end_time`                                                                                   | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `expand`                                                                                     | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `filters`                                                                                    | List[[components.TypesFilterCondition](../../models/components/typesfiltercondition.md)]     | :heavy_minus_sign:                                                                           | filters allows complex filtering based on multiple fields                                    |
| `limit`                                                                                      | *Optional[int]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `lookup_key`                                                                                 | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `offset`                                                                                     | *Optional[int]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `order`                                                                                      | [Optional[components.TypesPlanFilterOrder]](../../models/components/typesplanfilterorder.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `plan_ids`                                                                                   | List[*str*]                                                                                  | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `sort`                                                                                       | List[[components.TypesSortCondition](../../models/components/typessortcondition.md)]         | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `start_time`                                                                                 | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `status`                                                                                     | [Optional[components.TypesStatus]](../../models/components/typesstatus.md)                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `retries`                                                                                    | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                             | :heavy_minus_sign:                                                                           | Configuration to override the default retry behavior of the client.                          |

### Response

**[components.DtoListPlansResponse](../../models/components/dtolistplansresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_plans_id_

Get a plan by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/plans/{id}" method="get" path="/plans/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.plans.get_plans_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Plan ID                                                             |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPlanResponse](../../models/components/dtoplanresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_plans_id_

Update a plan by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/plans/{id}" method="put" path="/plans/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.plans.put_plans_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Plan ID                                                             |
| `description`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `display_order`                                                     | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `lookup_key`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `metadata`                                                          | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `name`                                                              | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPlanResponse](../../models/components/dtoplanresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_plans_id_

Delete a plan by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/plans/{id}" method="delete" path="/plans/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.plans.delete_plans_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Plan ID                                                             |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_plans_id_sync_subscriptions

Synchronize current plan prices with all existing active subscriptions

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/plans/{id}/sync/subscriptions" method="post" path="/plans/{id}/sync/subscriptions" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.plans.post_plans_id_sync_subscriptions(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Plan ID                                                             |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.ModelsTemporalWorkflowResult](../../models/components/modelstemporalworkflowresult.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404, 422                | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |