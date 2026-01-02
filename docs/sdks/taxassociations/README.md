# TaxAssociations

## Overview

### Available Operations

* [get_taxes_associations](#get_taxes_associations) - List tax associations
* [post_taxes_associations](#post_taxes_associations) - Create Tax Association
* [get_taxes_associations_id_](#get_taxes_associations_id_) - Get Tax Association
* [put_taxes_associations_id_](#put_taxes_associations_id_) - Update tax association
* [delete_taxes_associations_id_](#delete_taxes_associations_id_) - Delete tax association

## get_taxes_associations

List tax associations

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/taxes/associations" method="get" path="/taxes/associations" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tax_associations.get_taxes_associations()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `entity_type`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Entity Type                                                         |
| `entity_id`                                                         | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Entity ID                                                           |
| `tax_rate_id`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Tax Rate ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoListTaxAssociationsResponse](../../models/components/dtolisttaxassociationsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_taxes_associations

Create a new tax association

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/taxes/associations" method="post" path="/taxes/associations" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tax_associations.post_taxes_associations(entity_id="<id>", entity_type="subscription", tax_rate_code="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `entity_id`                                                                            | *str*                                                                                  | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `entity_type`                                                                          | [components.TypesTaxRateEntityType](../../models/components/typestaxrateentitytype.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `tax_rate_code`                                                                        | *str*                                                                                  | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `auto_apply`                                                                           | *Optional[bool]*                                                                       | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `currency`                                                                             | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `metadata`                                                                             | Dict[str, *str*]                                                                       | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `priority`                                                                             | *Optional[int]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `retries`                                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                       | :heavy_minus_sign:                                                                     | Configuration to override the default retry behavior of the client.                    |

### Response

**[components.DtoTaxAssociationResponse](../../models/components/dtotaxassociationresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_taxes_associations_id_

Get a tax association by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/taxes/associations/{id}" method="get" path="/taxes/associations/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tax_associations.get_taxes_associations_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Tax Config ID                                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoTaxAssociationResponse](../../models/components/dtotaxassociationresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_taxes_associations_id_

Update a tax association by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/taxes/associations/{id}" method="put" path="/taxes/associations/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tax_associations.put_taxes_associations_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Tax Config ID                                                       |
| `auto_apply`                                                        | *Optional[bool]*                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `metadata`                                                          | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `priority`                                                          | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoTaxAssociationResponse](../../models/components/dtotaxassociationresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_taxes_associations_id_

Delete a tax association by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/taxes/associations/{id}" method="delete" path="/taxes/associations/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tax_associations.delete_taxes_associations_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Tax Config ID                                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoTaxAssociationResponse](../../models/components/dtotaxassociationresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |