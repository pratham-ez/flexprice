# PriceUnits

## Overview

### Available Operations

* [get_prices_units](#get_prices_units) - List price units
* [post_prices_units](#post_prices_units) - Create a new price unit
* [get_prices_units_code_code_](#get_prices_units_code_code_) - Get a price unit by code
* [post_prices_units_search](#post_prices_units_search) - List price units by filter
* [get_prices_units_id_](#get_prices_units_id_) - Get a price unit by ID
* [put_prices_units_id_](#put_prices_units_id_) - Update a price unit
* [delete_prices_units_id_](#delete_prices_units_id_) - Delete a price unit

## get_prices_units

Get a paginated list of price units with optional filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/prices/units" method="get" path="/prices/units" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.price_units.get_prices_units()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `status`                                                            | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by status                                                    |
| `limit`                                                             | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | Limit number of results                                             |
| `offset`                                                            | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | Offset for pagination                                               |
| `sort`                                                              | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Sort field                                                          |
| `order`                                                             | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Sort order (asc/desc)                                               |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoListPriceUnitsResponse](../../models/components/dtolistpriceunitsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_prices_units

Create a new price unit with the provided details

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/prices/units" method="post" path="/prices/units" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.price_units.post_prices_units(base_currency="<value>", code="<value>", conversion_rate="<value>", name="<value>", symbol="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `base_currency`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | *str*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | base_currency  is the currency that the price unit is based on                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| `code`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | *str*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `conversion_rate`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | *str*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | ConversionRate defines the exchange rate from this price unit to the base currency.<br/>This rate is used to convert amounts in the custom price unit to the base currency for storage and billing.<br/><br/>Conversion formula:<br/>  price_unit_amount * conversion_rate = base_currency_amount<br/><br/>Example:<br/>  If conversion_rate = "0.01" and base_currency = "usd":<br/>  100 price_unit tokens * 0.01 = 1.00 USD<br/><br/>Note: Rounding precision is determined by the base currency (e.g., USD uses 2 decimal places, JPY uses 0). |
| `name`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | *str*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `symbol`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | *str*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `metadata`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Dict[str, *str*]                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `retries`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Configuration to override the default retry behavior of the client.                                                                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[components.DtoCreatePriceUnitResponse](../../models/components/dtocreatepriceunitresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_prices_units_code_code_

Get a price unit by code

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/prices/units/code/{code}" method="get" path="/prices/units/code/{code}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.price_units.get_prices_units_code_code_(code="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `code`                                                              | *str*                                                               | :heavy_check_mark:                                                  | Price unit code                                                     |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPriceUnitResponse](../../models/components/dtopriceunitresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_prices_units_search

List price units by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/prices/units/search" method="post" path="/prices/units/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.price_units.post_prices_units_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `expand`                                                                   | *Optional[str]*                                                            | :heavy_minus_sign:                                                         | N/A                                                                        |
| `limit`                                                                    | *Optional[int]*                                                            | :heavy_minus_sign:                                                         | N/A                                                                        |
| `offset`                                                                   | *Optional[int]*                                                            | :heavy_minus_sign:                                                         | N/A                                                                        |
| `order`                                                                    | *Optional[str]*                                                            | :heavy_minus_sign:                                                         | N/A                                                                        |
| `sort`                                                                     | *Optional[str]*                                                            | :heavy_minus_sign:                                                         | N/A                                                                        |
| `status`                                                                   | [Optional[components.TypesStatus]](../../models/components/typesstatus.md) | :heavy_minus_sign:                                                         | N/A                                                                        |
| `retries`                                                                  | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)           | :heavy_minus_sign:                                                         | Configuration to override the default retry behavior of the client.        |

### Response

**[components.DtoListPriceUnitsResponse](../../models/components/dtolistpriceunitsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_prices_units_id_

Get a price unit by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/prices/units/{id}" method="get" path="/prices/units/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.price_units.get_prices_units_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Price unit ID                                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPriceUnitResponse](../../models/components/dtopriceunitresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_prices_units_id_

Update an existing price unit with the provided details. Only name and metadata can be updated.

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/prices/units/{id}" method="put" path="/prices/units/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.price_units.put_prices_units_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Price unit ID                                                       |
| `metadata`                                                          | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `name`                                                              | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPriceUnitResponse](../../models/components/dtopriceunitresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_prices_units_id_

Delete an existing price unit.

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/prices/units/{id}" method="delete" path="/prices/units/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.price_units.delete_prices_units_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Price unit ID                                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |