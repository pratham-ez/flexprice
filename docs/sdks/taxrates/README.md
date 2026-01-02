# TaxRates

## Overview

### Available Operations

* [get_taxes_rates](#get_taxes_rates) - Get tax rates
* [post_taxes_rates](#post_taxes_rates) - Create a tax rate
* [get_taxes_rates_id_](#get_taxes_rates_id_) - Get a tax rate
* [put_taxes_rates_id_](#put_taxes_rates_id_) - Update a tax rate
* [delete_taxes_rates_id_](#delete_taxes_rates_id_) - Delete a tax rate

## get_taxes_rates

Get tax rates

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/taxes/rates" method="get" path="/taxes/rates" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tax_rates.get_taxes_rates()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `end_time`                                                                                 | *Optional[str]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `expand`                                                                                   | *Optional[str]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `limit`                                                                                    | *Optional[int]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `offset`                                                                                   | *Optional[int]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `order`                                                                                    | [Optional[operations.GetTaxesRatesOrder]](../../models/operations/gettaxesratesorder.md)   | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `scope`                                                                                    | [Optional[operations.GetTaxesRatesScope]](../../models/operations/gettaxesratesscope.md)   | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `start_time`                                                                               | *Optional[str]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `status`                                                                                   | [Optional[operations.GetTaxesRatesStatus]](../../models/operations/gettaxesratesstatus.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `taxrate_codes`                                                                            | List[*str*]                                                                                | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `taxrate_ids`                                                                              | List[*str*]                                                                                | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `retries`                                                                                  | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                           | :heavy_minus_sign:                                                                         | Configuration to override the default retry behavior of the client.                        |

### Response

**[List[components.DtoTaxRateResponse]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_taxes_rates

Create a tax rate

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/taxes/rates" method="post" path="/taxes/rates" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tax_rates.post_taxes_rates(code="<value>", name="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `code`                                                                                 | *str*                                                                                  | :heavy_check_mark:                                                                     | code is the unique alphanumeric case sensitive identifier for the tax rate (required)  |
| `name`                                                                                 | *str*                                                                                  | :heavy_check_mark:                                                                     | name is the human-readable name for the tax rate (required)                            |
| `description`                                                                          | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | description is an optional text description providing details about the tax rate       |
| `fixed_value`                                                                          | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | fixed_value is the fixed monetary amount when tax_rate_type is "fixed"                 |
| `metadata`                                                                             | Dict[str, *str*]                                                                       | :heavy_minus_sign:                                                                     | metadata contains additional key-value pairs for storing extra information             |
| `percentage_value`                                                                     | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | percentage_value is the percentage value (0-100) when tax_rate_type is "percentage"    |
| `scope`                                                                                | [Optional[components.TypesTaxRateScope]](../../models/components/typestaxratescope.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `tax_rate_type`                                                                        | [Optional[components.TypesTaxRateType]](../../models/components/typestaxratetype.md)   | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `retries`                                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                       | :heavy_minus_sign:                                                                     | Configuration to override the default retry behavior of the client.                    |

### Response

**[components.DtoTaxRateResponse](../../models/components/dtotaxrateresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_taxes_rates_id_

Get a tax rate

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/taxes/rates/{id}" method="get" path="/taxes/rates/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tax_rates.get_taxes_rates_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Tax rate ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoTaxRateResponse](../../models/components/dtotaxrateresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_taxes_rates_id_

Update a tax rate

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/taxes/rates/{id}" method="put" path="/taxes/rates/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tax_rates.put_taxes_rates_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `id`                                                                                     | *str*                                                                                    | :heavy_check_mark:                                                                       | Tax rate ID                                                                              |
| `code`                                                                                   | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | code is the updated unique alphanumeric identifier for the tax rate                      |
| `description`                                                                            | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | description is the updated text description for the tax rate                             |
| `metadata`                                                                               | Dict[str, *str*]                                                                         | :heavy_minus_sign:                                                                       | metadata contains updated key-value pairs that will replace existing metadata            |
| `name`                                                                                   | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | name is the updated human-readable name for the tax rate                                 |
| `tax_rate_status`                                                                        | [Optional[components.TypesTaxRateStatus]](../../models/components/typestaxratestatus.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `retries`                                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                         | :heavy_minus_sign:                                                                       | Configuration to override the default retry behavior of the client.                      |

### Response

**[components.DtoTaxRateResponse](../../models/components/dtotaxrateresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_taxes_rates_id_

Delete a tax rate

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/taxes/rates/{id}" method="delete" path="/taxes/rates/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    flex_price.tax_rates.delete_taxes_rates_id_(id="<id>")

    # Use the SDK ...

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Tax rate ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |