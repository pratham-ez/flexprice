# Features

## Overview

### Available Operations

* [get_features](#get_features) - List features
* [post_features](#post_features) - Create a new feature
* [post_features_search](#post_features_search) - List features by filter
* [get_features_id_](#get_features_id_) - Get a feature by ID
* [put_features_id_](#put_features_id_) - Update a feature
* [delete_features_id_](#delete_features_id_) - Delete a feature

## get_features

List features with optional filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/features" method="get" path="/features" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.features.get_features()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `end_time`                                                                             | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `expand`                                                                               | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `feature_ids`                                                                          | List[*str*]                                                                            | :heavy_minus_sign:                                                                     | Feature specific filters                                                               |
| `limit`                                                                                | *Optional[int]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `lookup_key`                                                                           | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `lookup_keys`                                                                          | List[*str*]                                                                            | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `meter_ids`                                                                            | List[*str*]                                                                            | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `name_contains`                                                                        | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `offset`                                                                               | *Optional[int]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `order`                                                                                | [Optional[operations.GetFeaturesOrder]](../../models/operations/getfeaturesorder.md)   | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `start_time`                                                                           | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `status`                                                                               | [Optional[operations.GetFeaturesStatus]](../../models/operations/getfeaturesstatus.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `retries`                                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                       | :heavy_minus_sign:                                                                     | Configuration to override the default retry behavior of the client.                    |

### Response

**[components.DtoListFeaturesResponse](../../models/components/dtolistfeaturesresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_features

Create a new feature

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/features" method="post" path="/features" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.features.post_features(name="<value>", type_="metered", meter={
        "aggregation": {},
        "event_name": "api_request",
        "name": "API Usage Meter",
        "reset_usage": "NEVER",
    })

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `name`                                                                                         | *str*                                                                                          | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `type`                                                                                         | [components.TypesFeatureType](../../models/components/typesfeaturetype.md)                     | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `alert_settings`                                                                               | [Optional[components.TypesAlertSettings]](../../models/components/typesalertsettings.md)       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `description`                                                                                  | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `lookup_key`                                                                                   | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `metadata`                                                                                     | Dict[str, *str*]                                                                               | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `meter`                                                                                        | [Optional[components.DtoCreateMeterRequest]](../../models/components/dtocreatemeterrequest.md) | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `meter_id`                                                                                     | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `unit_plural`                                                                                  | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `unit_singular`                                                                                | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `retries`                                                                                      | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                               | :heavy_minus_sign:                                                                             | Configuration to override the default retry behavior of the client.                            |

### Response

**[components.DtoFeatureResponse](../../models/components/dtofeatureresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_features_search

List features by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/features/search" method="post" path="/features/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.features.post_features_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `end_time`                                                                                         | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `expand`                                                                                           | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `feature_ids`                                                                                      | List[*str*]                                                                                        | :heavy_minus_sign:                                                                                 | Feature specific filters                                                                           |
| `filters`                                                                                          | List[[components.TypesFilterCondition](../../models/components/typesfiltercondition.md)]           | :heavy_minus_sign:                                                                                 | filters allows complex filtering based on multiple fields                                          |
| `limit`                                                                                            | *Optional[int]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `lookup_key`                                                                                       | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `lookup_keys`                                                                                      | List[*str*]                                                                                        | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `meter_ids`                                                                                        | List[*str*]                                                                                        | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `name_contains`                                                                                    | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `offset`                                                                                           | *Optional[int]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `order`                                                                                            | [Optional[components.TypesFeatureFilterOrder]](../../models/components/typesfeaturefilterorder.md) | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `sort`                                                                                             | List[[components.TypesSortCondition](../../models/components/typessortcondition.md)]               | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `start_time`                                                                                       | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `status`                                                                                           | [Optional[components.TypesStatus]](../../models/components/typesstatus.md)                         | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `retries`                                                                                          | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                   | :heavy_minus_sign:                                                                                 | Configuration to override the default retry behavior of the client.                                |

### Response

**[components.DtoListFeaturesResponse](../../models/components/dtolistfeaturesresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_features_id_

Get a feature by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/features/{id}" method="get" path="/features/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.features.get_features_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Feature ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoFeatureResponse](../../models/components/dtofeatureresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_features_id_

Update a feature by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/features/{id}" method="put" path="/features/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.features.put_features_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `id`                                                                                     | *str*                                                                                    | :heavy_check_mark:                                                                       | Feature ID                                                                               |
| `alert_settings`                                                                         | [Optional[components.TypesAlertSettings]](../../models/components/typesalertsettings.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `description`                                                                            | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `filters`                                                                                | List[[components.MeterFilter](../../models/components/meterfilter.md)]                   | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `metadata`                                                                               | Dict[str, *str*]                                                                         | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `name`                                                                                   | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `unit_plural`                                                                            | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `unit_singular`                                                                          | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `retries`                                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                         | :heavy_minus_sign:                                                                       | Configuration to override the default retry behavior of the client.                      |

### Response

**[components.DtoFeatureResponse](../../models/components/dtofeatureresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_features_id_

Delete a feature by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/features/{id}" method="delete" path="/features/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.features.delete_features_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Feature ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |