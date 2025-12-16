# Costs

## Overview

### Available Operations

* [create](#create) - Create a new costsheet
* [get_active](#get_active) - Get active costsheet for tenant
* [get_analytics](#get_analytics) - Get combined revenue and cost analytics
* [search](#search) - List costsheets by filter
* [get_by_id](#get_by_id) - Get a costsheet by ID
* [update](#update) - Update a costsheet
* [delete](#delete) - Delete a costsheet

## create

Create a new costsheet with the specified name

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/costs" method="post" path="/costs" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.costs.create(name="<value>")

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

**[models.DtoCreateCostsheetResponse](../../models/dtocreatecostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 409                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_active

Get the active costsheet for the current tenant

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/costs/active" method="get" path="/costs/active" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.costs.get_active()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoCostsheetResponse](../../models/dtocostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 404                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_analytics

Retrieve combined analytics with ROI, margin, and detailed breakdowns. If start_time and end_time are not provided, defaults to last 7 days.

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/costs/analytics" method="post" path="/costs/analytics" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.costs.get_analytics()

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

**[models.DtoGetDetailedCostAnalyticsResponse](../../models/dtogetdetailedcostanalyticsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## search

List costsheet records by filter with POST body

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/costs/search" method="post" path="/costs/search" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.costs.search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `costsheet_i_ds`                                                              | List[*str*]                                                                   | :heavy_minus_sign:                                                            | CostsheetIDs allows filtering by specific costsheet IDs                       |
| `environment_id`                                                              | *Optional[str]*                                                               | :heavy_minus_sign:                                                            | EnvironmentID filters by specific environment ID                              |
| `filters`                                                                     | List[[models.TypesFilterCondition](../../models/typesfiltercondition.md)]     | :heavy_minus_sign:                                                            | Filters contains custom filtering conditions                                  |
| `lookup_key`                                                                  | *Optional[str]*                                                               | :heavy_minus_sign:                                                            | LookupKey filters by lookup key                                               |
| `name`                                                                        | *Optional[str]*                                                               | :heavy_minus_sign:                                                            | Name filters by costsheet name                                                |
| `query_filter`                                                                | [Optional[models.TypesQueryFilter]](../../models/typesqueryfilter.md)         | :heavy_minus_sign:                                                            | N/A                                                                           |
| `sort`                                                                        | List[[models.TypesSortCondition](../../models/typessortcondition.md)]         | :heavy_minus_sign:                                                            | Sort specifies result ordering preferences                                    |
| `status`                                                                      | [Optional[models.TypesStatus]](../../models/typesstatus.md)                   | :heavy_minus_sign:                                                            | N/A                                                                           |
| `tenant_id`                                                                   | *Optional[str]*                                                               | :heavy_minus_sign:                                                            | TenantID filters by specific tenant ID                                        |
| `time_range_filter`                                                           | [Optional[models.TypesTimeRangeFilter]](../../models/typestimerangefilter.md) | :heavy_minus_sign:                                                            | N/A                                                                           |
| `retries`                                                                     | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)              | :heavy_minus_sign:                                                            | Configuration to override the default retry behavior of the client.           |

### Response

**[models.DtoListCostsheetResponse](../../models/dtolistcostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_by_id

Get a costsheet by ID with optional price expansion

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/costs/{id}" method="get" path="/costs/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.costs.get_by_id(id="<id>")

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

**[models.DtoGetCostsheetResponse](../../models/dtogetcostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update

Update a costsheet with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/costs/{id}" method="put" path="/costs/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.costs.update(id="<id>")

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

**[models.DtoUpdateCostsheetResponse](../../models/dtoupdatecostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404, 409                | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete

Soft delete a costsheet by setting its status to deleted

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/costs/{id}" method="delete" path="/costs/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.costs.delete(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Costsheet ID                                                        |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoDeleteCostsheetResponse](../../models/dtodeletecostsheetresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |