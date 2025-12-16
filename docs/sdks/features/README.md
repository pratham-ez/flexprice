# Features

## Overview

### Available Operations

* [list](#list) - List features
* [create](#create) - Create a new feature
* [search](#search) - List features by filter
* [get_by_id](#get_by_id) - Get a feature by ID
* [update](#update) - Update a feature
* [delete](#delete) - Delete a feature

## list

List features with optional filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/features" method="get" path="/features" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.features.list()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `end_time`                                                                                  | *Optional[str]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `expand`                                                                                    | *Optional[str]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `feature_ids`                                                                               | List[*str*]                                                                                 | :heavy_minus_sign:                                                                          | Feature specific filters                                                                    |
| `limit`                                                                                     | *Optional[int]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `lookup_key`                                                                                | *Optional[str]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `lookup_keys`                                                                               | List[*str*]                                                                                 | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `meter_ids`                                                                                 | List[*str*]                                                                                 | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `name_contains`                                                                             | *Optional[str]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `offset`                                                                                    | *Optional[int]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `order`                                                                                     | [Optional[models.GetFeaturesQueryParamOrder]](../../models/getfeaturesqueryparamorder.md)   | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `start_time`                                                                                | *Optional[str]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `status`                                                                                    | [Optional[models.GetFeaturesQueryParamStatus]](../../models/getfeaturesqueryparamstatus.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `retries`                                                                                   | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                            | :heavy_minus_sign:                                                                          | Configuration to override the default retry behavior of the client.                         |

### Response

**[models.DtoListFeaturesResponse](../../models/dtolistfeaturesresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create

Create a new feature

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/features" method="post" path="/features" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.features.create(name="<value>", type_="metered", meter={
        "aggregation": {},
        "event_name": "api_request",
        "name": "API Usage Meter",
        "reset_usage": "NEVER",
    })

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `name`                                                                          | *str*                                                                           | :heavy_check_mark:                                                              | N/A                                                                             |
| `type`                                                                          | [models.TypesFeatureType](../../models/typesfeaturetype.md)                     | :heavy_check_mark:                                                              | N/A                                                                             |
| `alert_settings`                                                                | [Optional[models.TypesAlertSettings]](../../models/typesalertsettings.md)       | :heavy_minus_sign:                                                              | N/A                                                                             |
| `description`                                                                   | *Optional[str]*                                                                 | :heavy_minus_sign:                                                              | N/A                                                                             |
| `lookup_key`                                                                    | *Optional[str]*                                                                 | :heavy_minus_sign:                                                              | N/A                                                                             |
| `metadata`                                                                      | Dict[str, *str*]                                                                | :heavy_minus_sign:                                                              | N/A                                                                             |
| `meter`                                                                         | [Optional[models.DtoCreateMeterRequest]](../../models/dtocreatemeterrequest.md) | :heavy_minus_sign:                                                              | N/A                                                                             |
| `meter_id`                                                                      | *Optional[str]*                                                                 | :heavy_minus_sign:                                                              | N/A                                                                             |
| `unit_plural`                                                                   | *Optional[str]*                                                                 | :heavy_minus_sign:                                                              | N/A                                                                             |
| `unit_singular`                                                                 | *Optional[str]*                                                                 | :heavy_minus_sign:                                                              | N/A                                                                             |
| `retries`                                                                       | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                | :heavy_minus_sign:                                                              | Configuration to override the default retry behavior of the client.             |

### Response

**[models.DtoFeatureResponse](../../models/dtofeatureresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## search

List features by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/features/search" method="post" path="/features/search" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.features.search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `end_time`                                                                          | *Optional[str]*                                                                     | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `expand`                                                                            | *Optional[str]*                                                                     | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `feature_ids`                                                                       | List[*str*]                                                                         | :heavy_minus_sign:                                                                  | Feature specific filters                                                            |
| `filters`                                                                           | List[[models.TypesFilterCondition](../../models/typesfiltercondition.md)]           | :heavy_minus_sign:                                                                  | filters allows complex filtering based on multiple fields                           |
| `limit`                                                                             | *Optional[int]*                                                                     | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `lookup_key`                                                                        | *Optional[str]*                                                                     | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `lookup_keys`                                                                       | List[*str*]                                                                         | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `meter_ids`                                                                         | List[*str*]                                                                         | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `name_contains`                                                                     | *Optional[str]*                                                                     | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `offset`                                                                            | *Optional[int]*                                                                     | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `order`                                                                             | [Optional[models.TypesFeatureFilterOrder]](../../models/typesfeaturefilterorder.md) | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `sort`                                                                              | List[[models.TypesSortCondition](../../models/typessortcondition.md)]               | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `start_time`                                                                        | *Optional[str]*                                                                     | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `status`                                                                            | [Optional[models.TypesStatus]](../../models/typesstatus.md)                         | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `retries`                                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                    | :heavy_minus_sign:                                                                  | Configuration to override the default retry behavior of the client.                 |

### Response

**[models.DtoListFeaturesResponse](../../models/dtolistfeaturesresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_by_id

Get a feature by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/features/{id}" method="get" path="/features/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.features.get_by_id(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Feature ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoFeatureResponse](../../models/dtofeatureresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update

Update a feature by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/features/{id}" method="put" path="/features/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.features.update(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `id`                                                                      | *str*                                                                     | :heavy_check_mark:                                                        | Feature ID                                                                |
| `alert_settings`                                                          | [Optional[models.TypesAlertSettings]](../../models/typesalertsettings.md) | :heavy_minus_sign:                                                        | N/A                                                                       |
| `description`                                                             | *Optional[str]*                                                           | :heavy_minus_sign:                                                        | N/A                                                                       |
| `filters`                                                                 | List[[models.MeterFilter](../../models/meterfilter.md)]                   | :heavy_minus_sign:                                                        | N/A                                                                       |
| `metadata`                                                                | Dict[str, *str*]                                                          | :heavy_minus_sign:                                                        | N/A                                                                       |
| `name`                                                                    | *Optional[str]*                                                           | :heavy_minus_sign:                                                        | N/A                                                                       |
| `unit_plural`                                                             | *Optional[str]*                                                           | :heavy_minus_sign:                                                        | N/A                                                                       |
| `unit_singular`                                                           | *Optional[str]*                                                           | :heavy_minus_sign:                                                        | N/A                                                                       |
| `retries`                                                                 | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)          | :heavy_minus_sign:                                                        | Configuration to override the default retry behavior of the client.       |

### Response

**[models.DtoFeatureResponse](../../models/dtofeatureresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete

Delete a feature by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/features/{id}" method="delete" path="/features/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.features.delete(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Feature ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, models.GinH]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |