# Costs

## Overview

### Available Operations

* [post_costs](#post_costs) - Create a new costsheet
* [get_costs_active](#get_costs_active) - Get active costsheet for tenant
* [post_costs_analytics](#post_costs_analytics) - Get combined revenue and cost analytics
* [post_costs_analytics_v2](#post_costs_analytics_v2) - Get combined revenue and cost analytics
* [post_costs_search](#post_costs_search) - List costsheets by filter
* [get_costs_id_](#get_costs_id_) - Get a costsheet by ID
* [put_costs_id_](#put_costs_id_) - Update a costsheet
* [delete_costs_id_](#delete_costs_id_) - Delete a costsheet

## post_costs

Create a new costsheet with the specified name

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/costs" method="post" path="/costs" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.costs.post_costs(name="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `name`                                                              | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `description`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `lookup_key`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `metadata`                                                          | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCreateCostsheetResponse](../../models/components/dtocreatecostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 409                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_costs_active

Get the active costsheet for the current tenant

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/costs/active" method="get" path="/costs/active" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.costs.get_costs_active()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCostsheetResponse](../../models/components/dtocostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 404                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_costs_analytics

Retrieve combined analytics with ROI, margin, and detailed breakdowns. If start_time and end_time are not provided, defaults to last 7 days.

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/costs/analytics" method="post" path="/costs/analytics" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.costs.post_costs_analytics()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `end_time`                                                             | *Optional[str]*                                                        | :heavy_minus_sign:                                                     | N/A                                                                    |
| `expand`                                                               | List[*str*]                                                            | :heavy_minus_sign:                                                     | Expand options - specify which entities to expand                      |
| `external_customer_id`                                                 | *Optional[str]*                                                        | :heavy_minus_sign:                                                     | Optional - for specific customer                                       |
| `feature_ids`                                                          | List[*str*]                                                            | :heavy_minus_sign:                                                     | Additional filters                                                     |
| `limit`                                                                | *Optional[int]*                                                        | :heavy_minus_sign:                                                     | Pagination                                                             |
| `offset`                                                               | *Optional[int]*                                                        | :heavy_minus_sign:                                                     | N/A                                                                    |
| `start_time`                                                           | *Optional[str]*                                                        | :heavy_minus_sign:                                                     | Time range fields (optional - defaults to last 7 days if not provided) |
| `retries`                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)       | :heavy_minus_sign:                                                     | Configuration to override the default retry behavior of the client.    |

### Response

**[components.DtoGetDetailedCostAnalyticsResponse](../../models/components/dtogetdetailedcostanalyticsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_costs_analytics_v2

Retrieve combined analytics with ROI, margin, and detailed breakdowns. If start_time and end_time are not provided, defaults to last 7 days.

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/costs/analytics-v2" method="post" path="/costs/analytics-v2" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.costs.post_costs_analytics_v2()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `end_time`                                                             | *Optional[str]*                                                        | :heavy_minus_sign:                                                     | N/A                                                                    |
| `expand`                                                               | List[*str*]                                                            | :heavy_minus_sign:                                                     | Expand options - specify which entities to expand                      |
| `external_customer_id`                                                 | *Optional[str]*                                                        | :heavy_minus_sign:                                                     | Optional - for specific customer                                       |
| `feature_ids`                                                          | List[*str*]                                                            | :heavy_minus_sign:                                                     | Additional filters                                                     |
| `limit`                                                                | *Optional[int]*                                                        | :heavy_minus_sign:                                                     | Pagination                                                             |
| `offset`                                                               | *Optional[int]*                                                        | :heavy_minus_sign:                                                     | N/A                                                                    |
| `start_time`                                                           | *Optional[str]*                                                        | :heavy_minus_sign:                                                     | Time range fields (optional - defaults to last 7 days if not provided) |
| `retries`                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)       | :heavy_minus_sign:                                                     | Configuration to override the default retry behavior of the client.    |

### Response

**[components.DtoGetDetailedCostAnalyticsResponse](../../models/components/dtogetdetailedcostanalyticsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_costs_search

List costsheet records by filter with POST body

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/costs/search" method="post" path="/costs/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.costs.post_costs_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `costsheet_i_ds`                                                                             | List[*str*]                                                                                  | :heavy_minus_sign:                                                                           | CostsheetIDs allows filtering by specific costsheet IDs                                      |
| `environment_id`                                                                             | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | EnvironmentID filters by specific environment ID                                             |
| `filters`                                                                                    | List[[components.TypesFilterCondition](../../models/components/typesfiltercondition.md)]     | :heavy_minus_sign:                                                                           | Filters contains custom filtering conditions                                                 |
| `lookup_key`                                                                                 | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | LookupKey filters by lookup key                                                              |
| `name`                                                                                       | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | Name filters by costsheet name                                                               |
| `query_filter`                                                                               | [Optional[components.TypesQueryFilter]](../../models/components/typesqueryfilter.md)         | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `sort`                                                                                       | List[[components.TypesSortCondition](../../models/components/typessortcondition.md)]         | :heavy_minus_sign:                                                                           | Sort specifies result ordering preferences                                                   |
| `status`                                                                                     | [Optional[components.TypesStatus]](../../models/components/typesstatus.md)                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `tenant_id`                                                                                  | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | TenantID filters by specific tenant ID                                                       |
| `time_range_filter`                                                                          | [Optional[components.TypesTimeRangeFilter]](../../models/components/typestimerangefilter.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `retries`                                                                                    | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                             | :heavy_minus_sign:                                                                           | Configuration to override the default retry behavior of the client.                          |

### Response

**[components.DtoListCostsheetResponse](../../models/components/dtolistcostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_costs_id_

Get a costsheet by ID with optional price expansion

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/costs/{id}" method="get" path="/costs/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.costs.get_costs_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Costsheet ID                                                        |
| `expand`                                                            | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Comma-separated list of fields to expand (e.g., 'prices')           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoGetCostsheetResponse](../../models/components/dtogetcostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_costs_id_

Update a costsheet with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/costs/{id}" method="put" path="/costs/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.costs.put_costs_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Costsheet ID                                                        |
| `description`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `lookup_key`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `metadata`                                                          | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `name`                                                              | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoUpdateCostsheetResponse](../../models/components/dtoupdatecostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404, 409                | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_costs_id_

Soft delete a costsheet by setting its status to deleted

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/costs/{id}" method="delete" path="/costs/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.costs.delete_costs_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Costsheet ID                                                        |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoDeleteCostsheetResponse](../../models/components/dtodeletecostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |