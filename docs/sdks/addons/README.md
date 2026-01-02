# Addons

## Overview

### Available Operations

* [get_addons](#get_addons) - List addons
* [post_addons](#post_addons) - Create addon
* [get_addons_lookup_lookup_key_](#get_addons_lookup_lookup_key_) - Get addon by lookup key
* [post_addons_search](#post_addons_search) - List addons by filter
* [get_addons_id_](#get_addons_id_) - Get addon
* [put_addons_id_](#put_addons_id_) - Update addon
* [delete_addons_id_](#delete_addons_id_) - Delete addon

## get_addons

Get addons with optional filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/addons" method="get" path="/addons" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.addons.get_addons()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `addon_ids`                                                                        | List[*str*]                                                                        | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `addon_type`                                                                       | [Optional[operations.AddonType]](../../models/operations/addontype.md)             | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `end_time`                                                                         | *Optional[str]*                                                                    | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `expand`                                                                           | *Optional[str]*                                                                    | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `limit`                                                                            | *Optional[int]*                                                                    | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `lookup_keys`                                                                      | List[*str*]                                                                        | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `offset`                                                                           | *Optional[int]*                                                                    | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `order`                                                                            | [Optional[operations.GetAddonsOrder]](../../models/operations/getaddonsorder.md)   | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `start_time`                                                                       | *Optional[str]*                                                                    | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `status`                                                                           | [Optional[operations.GetAddonsStatus]](../../models/operations/getaddonsstatus.md) | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `retries`                                                                          | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                   | :heavy_minus_sign:                                                                 | Configuration to override the default retry behavior of the client.                |

### Response

**[components.DtoListAddonsResponse](../../models/components/dtolistaddonsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_addons

Create a new addon

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/addons" method="post" path="/addons" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.addons.post_addons(lookup_key="<value>", name="<value>", type_="onetime")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `lookup_key`                                                           | *str*                                                                  | :heavy_check_mark:                                                     | N/A                                                                    |
| `name`                                                                 | *str*                                                                  | :heavy_check_mark:                                                     | N/A                                                                    |
| `type`                                                                 | [components.TypesAddonType](../../models/components/typesaddontype.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `description`                                                          | *Optional[str]*                                                        | :heavy_minus_sign:                                                     | N/A                                                                    |
| `metadata`                                                             | Dict[str, *Any*]                                                       | :heavy_minus_sign:                                                     | N/A                                                                    |
| `retries`                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)       | :heavy_minus_sign:                                                     | Configuration to override the default retry behavior of the client.    |

### Response

**[components.DtoCreateAddonResponse](../../models/components/dtocreateaddonresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_addons_lookup_lookup_key_

Get an addon by lookup key

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/addons/lookup/{lookup_key}" method="get" path="/addons/lookup/{lookup_key}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.addons.get_addons_lookup_lookup_key_(lookup_key="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `lookup_key`                                                        | *str*                                                               | :heavy_check_mark:                                                  | Addon Lookup Key                                                    |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoAddonResponse](../../models/components/dtoaddonresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_addons_search

List addons by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/addons/search" method="post" path="/addons/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.addons.post_addons_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `addon_ids`                                                                                    | List[*str*]                                                                                    | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `addon_type`                                                                                   | [Optional[components.TypesAddonType]](../../models/components/typesaddontype.md)               | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `end_time`                                                                                     | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `expand`                                                                                       | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `filters`                                                                                      | List[[components.TypesFilterCondition](../../models/components/typesfiltercondition.md)]       | :heavy_minus_sign:                                                                             | filters allows complex filtering based on multiple fields                                      |
| `limit`                                                                                        | *Optional[int]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `lookup_keys`                                                                                  | List[*str*]                                                                                    | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `offset`                                                                                       | *Optional[int]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `order`                                                                                        | [Optional[components.TypesAddonFilterOrder]](../../models/components/typesaddonfilterorder.md) | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `sort`                                                                                         | List[[components.TypesSortCondition](../../models/components/typessortcondition.md)]           | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `start_time`                                                                                   | *Optional[str]*                                                                                | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `status`                                                                                       | [Optional[components.TypesStatus]](../../models/components/typesstatus.md)                     | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `retries`                                                                                      | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                               | :heavy_minus_sign:                                                                             | Configuration to override the default retry behavior of the client.                            |

### Response

**[components.DtoListAddonsResponse](../../models/components/dtolistaddonsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_addons_id_

Get an addon by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/addons/{id}" method="get" path="/addons/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.addons.get_addons_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Addon ID                                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoAddonResponse](../../models/components/dtoaddonresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_addons_id_

Update an existing addon

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/addons/{id}" method="put" path="/addons/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.addons.put_addons_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Addon ID                                                            |
| `description`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `metadata`                                                          | Dict[str, *Any*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `name`                                                              | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoAddonResponse](../../models/components/dtoaddonresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_addons_id_

Delete an addon

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/addons/{id}" method="delete" path="/addons/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.addons.delete_addons_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Addon ID                                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |