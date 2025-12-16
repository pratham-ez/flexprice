# TaxRates

## Overview

### Available Operations

* [get_all](#get_all) - Get tax rates
* [create](#create) - Create a tax rate
* [get](#get) - Get a tax rate
* [update](#update) - Update a tax rate
* [delete](#delete) - Delete a tax rate

## get_all

Get tax rates

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/taxes/rates" method="get" path="/taxes/rates" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.tax_rates.get_all()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `end_time`                                                                                      | *Optional[str]*                                                                                 | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `expand`                                                                                        | *Optional[str]*                                                                                 | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `limit`                                                                                         | *Optional[int]*                                                                                 | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `offset`                                                                                        | *Optional[int]*                                                                                 | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `order`                                                                                         | [Optional[models.GetTaxesRatesQueryParamOrder]](../../models/gettaxesratesqueryparamorder.md)   | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `scope`                                                                                         | [Optional[models.QueryParamScope]](../../models/queryparamscope.md)                             | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `start_time`                                                                                    | *Optional[str]*                                                                                 | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `status`                                                                                        | [Optional[models.GetTaxesRatesQueryParamStatus]](../../models/gettaxesratesqueryparamstatus.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `taxrate_codes`                                                                                 | List[*str*]                                                                                     | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `taxrate_ids`                                                                                   | List[*str*]                                                                                     | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `retries`                                                                                       | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                | :heavy_minus_sign:                                                                              | Configuration to override the default retry behavior of the client.                             |

### Response

**[List[models.DtoTaxRateResponse]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create

Create a tax rate

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/taxes/rates" method="post" path="/taxes/rates" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.tax_rates.create(code="<value>", name="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `code`                                                                                | *str*                                                                                 | :heavy_check_mark:                                                                    | code is the unique alphanumeric case sensitive identifier for the tax rate (required) |
| `name`                                                                                | *str*                                                                                 | :heavy_check_mark:                                                                    | name is the human-readable name for the tax rate (required)                           |
| `description`                                                                         | *Optional[str]*                                                                       | :heavy_minus_sign:                                                                    | description is an optional text description providing details about the tax rate      |
| `fixed_value`                                                                         | *Optional[str]*                                                                       | :heavy_minus_sign:                                                                    | fixed_value is the fixed monetary amount when tax_rate_type is "fixed"                |
| `metadata`                                                                            | Dict[str, *str*]                                                                      | :heavy_minus_sign:                                                                    | metadata contains additional key-value pairs for storing extra information            |
| `percentage_value`                                                                    | *Optional[str]*                                                                       | :heavy_minus_sign:                                                                    | percentage_value is the percentage value (0-100) when tax_rate_type is "percentage"   |
| `scope`                                                                               | [Optional[models.TypesTaxRateScope]](../../models/typestaxratescope.md)               | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `tax_rate_type`                                                                       | [Optional[models.TypesTaxRateType]](../../models/typestaxratetype.md)                 | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `retries`                                                                             | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                      | :heavy_minus_sign:                                                                    | Configuration to override the default retry behavior of the client.                   |

### Response

**[models.DtoTaxRateResponse](../../models/dtotaxrateresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get

Get a tax rate

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/taxes/rates/{id}" method="get" path="/taxes/rates/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.tax_rates.get(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Tax rate ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoTaxRateResponse](../../models/dtotaxrateresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update

Update a tax rate

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/taxes/rates/{id}" method="put" path="/taxes/rates/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.tax_rates.update(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `id`                                                                          | *str*                                                                         | :heavy_check_mark:                                                            | Tax rate ID                                                                   |
| `code`                                                                        | *Optional[str]*                                                               | :heavy_minus_sign:                                                            | code is the updated unique alphanumeric identifier for the tax rate           |
| `description`                                                                 | *Optional[str]*                                                               | :heavy_minus_sign:                                                            | description is the updated text description for the tax rate                  |
| `metadata`                                                                    | Dict[str, *str*]                                                              | :heavy_minus_sign:                                                            | metadata contains updated key-value pairs that will replace existing metadata |
| `name`                                                                        | *Optional[str]*                                                               | :heavy_minus_sign:                                                            | name is the updated human-readable name for the tax rate                      |
| `tax_rate_status`                                                             | [Optional[models.TypesTaxRateStatus]](../../models/typestaxratestatus.md)     | :heavy_minus_sign:                                                            | N/A                                                                           |
| `retries`                                                                     | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)              | :heavy_minus_sign:                                                            | Configuration to override the default retry behavior of the client.           |

### Response

**[models.DtoTaxRateResponse](../../models/dtotaxrateresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete

Delete a tax rate

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/taxes/rates/{id}" method="delete" path="/taxes/rates/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    f_client.tax_rates.delete(id="<id>")

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
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |