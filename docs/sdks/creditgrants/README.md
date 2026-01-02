# CreditGrants

## Overview

### Available Operations

* [get_creditgrants](#get_creditgrants) - Get credit grants
* [post_creditgrants](#post_creditgrants) - Create a new credit grant
* [get_creditgrants_id_](#get_creditgrants_id_) - Get a credit grant by ID
* [put_creditgrants_id_](#put_creditgrants_id_) - Update a credit grant
* [delete_creditgrants_id_](#delete_creditgrants_id_) - Delete a credit grant
* [get_plans_id_creditgrants](#get_plans_id_creditgrants) - Get plan credit grants

## get_creditgrants

Get credit grants with the specified filter

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/creditgrants" method="get" path="/creditgrants" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_grants.get_creditgrants()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `end_time`                                                                                     | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `expand`                                                                                       | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `limit`                                                                                        | *Optional[int]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `offset`                                                                                       | *Optional[int]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `order`                                                                                        | [Optional[operations.GetCreditgrantsOrder]](../../models/operations/getcreditgrantsorder.md)   | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `plan_ids`                                                                                     | List[*str*]                                                                                    | :heavy_minus_sign:                                                                             | Specific filters for credit grants                                                             |
| `scope`                                                                                        | [Optional[operations.GetCreditgrantsScope]](../../models/operations/getcreditgrantsscope.md)   | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `sort`                                                                                         | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `start_time`                                                                                   | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `status`                                                                                       | [Optional[operations.GetCreditgrantsStatus]](../../models/operations/getcreditgrantsstatus.md) | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `subscription_ids`                                                                             | List[*str*]                                                                                    | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `retries`                                                                                      | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                               | :heavy_minus_sign:                                                                             | Configuration to override the default retry behavior of the client.                            |

### Response

**[components.DtoListCreditGrantsResponse](../../models/components/dtolistcreditgrantsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_creditgrants

Create a new credit grant with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/creditgrants" method="post" path="/creditgrants" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_grants.post_creditgrants(cadence="RECURRING", credits="<value>", name="<value>", scope="PLAN")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `cadence`                                                                                                                | [components.TypesCreditGrantCadence](../../models/components/typescreditgrantcadence.md)                                 | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `credits`                                                                                                                | *str*                                                                                                                    | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `name`                                                                                                                   | *str*                                                                                                                    | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `scope`                                                                                                                  | [components.TypesCreditGrantScope](../../models/components/typescreditgrantscope.md)                                     | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `expiration_duration`                                                                                                    | *Optional[int]*                                                                                                          | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `expiration_duration_unit`                                                                                               | [Optional[components.TypesCreditGrantExpiryDurationUnit]](../../models/components/typescreditgrantexpirydurationunit.md) | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `expiration_type`                                                                                                        | [Optional[components.TypesCreditGrantExpiryType]](../../models/components/typescreditgrantexpirytype.md)                 | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `metadata`                                                                                                               | Dict[str, *str*]                                                                                                         | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `period`                                                                                                                 | [Optional[components.TypesCreditGrantPeriod]](../../models/components/typescreditgrantperiod.md)                         | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `period_count`                                                                                                           | *Optional[int]*                                                                                                          | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `plan_id`                                                                                                                | *Optional[str]*                                                                                                          | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `priority`                                                                                                               | *Optional[int]*                                                                                                          | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `subscription_id`                                                                                                        | *Optional[str]*                                                                                                          | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `retries`                                                                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                         | :heavy_minus_sign:                                                                                                       | Configuration to override the default retry behavior of the client.                                                      |

### Response

**[components.DtoCreditGrantResponse](../../models/components/dtocreditgrantresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_creditgrants_id_

Get a credit grant by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/creditgrants/{id}" method="get" path="/creditgrants/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_grants.get_creditgrants_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Credit Grant ID                                                     |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCreditGrantResponse](../../models/components/dtocreditgrantresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_creditgrants_id_

Update a credit grant with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/creditgrants/{id}" method="put" path="/creditgrants/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_grants.put_creditgrants_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Credit Grant ID                                                     |
| `metadata`                                                          | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `name`                                                              | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCreditGrantResponse](../../models/components/dtocreditgrantresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_creditgrants_id_

Delete a credit grant

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/creditgrants/{id}" method="delete" path="/creditgrants/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_grants.delete_creditgrants_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Credit Grant ID                                                     |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_plans_id_creditgrants

Get all credit grants for a plan

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/plans/{id}/creditgrants" method="get" path="/plans/{id}/creditgrants" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_grants.get_plans_id_creditgrants(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Plan ID                                                             |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoListCreditGrantsResponse](../../models/components/dtolistcreditgrantsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |