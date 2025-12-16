# TaxAssociations

## Overview

### Available Operations

* [list](#list) - List tax associations
* [create](#create) - Create Tax Association
* [get_by_id](#get_by_id) - Get Tax Association
* [update](#update) - Update tax association
* [delete](#delete) - Delete tax association

## list

List tax associations

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/taxes/associations" method="get" path="/taxes/associations" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.tax_associations.list()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `auto_apply`                                                                                      | *Optional[bool]*                                                                                  | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `currency`                                                                                        | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `end_time`                                                                                        | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `entity_id`                                                                                       | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `entity_type`                                                                                     | [Optional[models.TypesTaxRateEntityType]](../../models/typestaxrateentitytype.md)                 | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `expand`                                                                                          | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `limit`                                                                                           | *Optional[int]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `offset`                                                                                          | *Optional[int]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `order`                                                                                           | [Optional[models.TypesTaxAssociationFilterOrder]](../../models/typestaxassociationfilterorder.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `sort`                                                                                            | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `start_time`                                                                                      | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `status`                                                                                          | [Optional[models.TypesStatus]](../../models/typesstatus.md)                                       | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `tax_association_ids`                                                                             | List[*str*]                                                                                       | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `tax_rate_ids`                                                                                    | List[*str*]                                                                                       | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `retries`                                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                  | :heavy_minus_sign:                                                                                | Configuration to override the default retry behavior of the client.                               |

### Response

**[models.DtoListTaxAssociationsResponse](../../models/dtolisttaxassociationsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create

Create a new tax association

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/taxes/associations" method="post" path="/taxes/associations" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.tax_associations.create(entity_id="<id>", entity_type="subscription", tax_rate_code="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `entity_id`                                                             | *str*                                                                   | :heavy_check_mark:                                                      | N/A                                                                     |
| `entity_type`                                                           | [models.TypesTaxRateEntityType](../../models/typestaxrateentitytype.md) | :heavy_check_mark:                                                      | N/A                                                                     |
| `tax_rate_code`                                                         | *str*                                                                   | :heavy_check_mark:                                                      | N/A                                                                     |
| `auto_apply`                                                            | *Optional[bool]*                                                        | :heavy_minus_sign:                                                      | N/A                                                                     |
| `currency`                                                              | *Optional[str]*                                                         | :heavy_minus_sign:                                                      | N/A                                                                     |
| `metadata`                                                              | Dict[str, *str*]                                                        | :heavy_minus_sign:                                                      | N/A                                                                     |
| `priority`                                                              | *Optional[int]*                                                         | :heavy_minus_sign:                                                      | N/A                                                                     |
| `retries`                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)        | :heavy_minus_sign:                                                      | Configuration to override the default retry behavior of the client.     |

### Response

**[models.DtoTaxAssociationResponse](../../models/dtotaxassociationresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_by_id

Get a tax association by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/taxes/associations/{id}" method="get" path="/taxes/associations/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.tax_associations.get_by_id(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Tax Config ID                                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoTaxAssociationResponse](../../models/dtotaxassociationresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update

Update a tax association by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/taxes/associations/{id}" method="put" path="/taxes/associations/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.tax_associations.update(id="<id>")

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

**[models.DtoTaxAssociationResponse](../../models/dtotaxassociationresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete

Delete a tax association by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/taxes/associations/{id}" method="delete" path="/taxes/associations/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.tax_associations.delete(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Tax Config ID                                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoTaxAssociationResponse](../../models/dtotaxassociationresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |