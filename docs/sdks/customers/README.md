# Customers

## Overview

### Available Operations

* [list](#list) - Get customers
* [create](#create) - Create a customer
* [get_by_lookup_key](#get_by_lookup_key) - Get a customer by lookup key
* [search](#search) - List customers by filter
* [get_usage_summary](#get_usage_summary) - Get customer usage summary
* [get_by_id](#get_by_id) - Get a customer
* [update](#update) - Update a customer
* [delete](#delete) - Delete a customer
* [get_entitlements](#get_entitlements) - Get customer entitlements
* [get_upcoming_grants](#get_upcoming_grants) - Get upcoming credit grant applications

## list

Get customers

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers" method="get" path="/customers" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.customers.list()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `customer_ids`                                                                                | List[*str*]                                                                                   | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `email`                                                                                       | *Optional[str]*                                                                               | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `end_time`                                                                                    | *Optional[str]*                                                                               | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `expand`                                                                                      | *Optional[str]*                                                                               | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `external_id`                                                                                 | *Optional[str]*                                                                               | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `external_ids`                                                                                | List[*str*]                                                                                   | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `limit`                                                                                       | *Optional[int]*                                                                               | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `offset`                                                                                      | *Optional[int]*                                                                               | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `order`                                                                                       | [Optional[models.GetCustomersQueryParamOrder]](../../models/getcustomersqueryparamorder.md)   | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `parent_customer_ids`                                                                         | List[*str*]                                                                                   | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `start_time`                                                                                  | *Optional[str]*                                                                               | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `status`                                                                                      | [Optional[models.GetCustomersQueryParamStatus]](../../models/getcustomersqueryparamstatus.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `retries`                                                                                     | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                              | :heavy_minus_sign:                                                                            | Configuration to override the default retry behavior of the client.                           |

### Response

**[models.DtoListCustomersResponse](../../models/dtolistcustomersresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create

Create a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/customers" method="post" path="/customers" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.customers.create(external_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                               | Type                                                                                                                                                                    | Required                                                                                                                                                                | Description                                                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `external_id`                                                                                                                                                           | *str*                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                      | external_id is the unique identifier from your system to reference this customer (required)                                                                             |
| `address_city`                                                                                                                                                          | *Optional[str]*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                      | address_city is the city name with maximum 100 characters                                                                                                               |
| `address_country`                                                                                                                                                       | *Optional[str]*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                      | address_country is the two-letter ISO 3166-1 alpha-2 country code                                                                                                       |
| `address_line1`                                                                                                                                                         | *Optional[str]*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                      | address_line1 is the primary address line with maximum 255 characters                                                                                                   |
| `address_line2`                                                                                                                                                         | *Optional[str]*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                      | address_line2 is the secondary address line with maximum 255 characters                                                                                                 |
| `address_postal_code`                                                                                                                                                   | *Optional[str]*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                      | address_postal_code is the ZIP code or postal code with maximum 20 characters                                                                                           |
| `address_state`                                                                                                                                                         | *Optional[str]*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                      | address_state is the state, province, or region name with maximum 100 characters                                                                                        |
| `email`                                                                                                                                                                 | *Optional[str]*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                      | email is the customer's email address and must be a valid email format if provided                                                                                      |
| `integration_entity_mapping`                                                                                                                                            | List[[models.DtoIntegrationEntityMapping](../../models/dtointegrationentitymapping.md)]                                                                                 | :heavy_minus_sign:                                                                                                                                                      | integration_entity_mapping contains provider integration mappings for this customer                                                                                     |
| `metadata`                                                                                                                                                              | Dict[str, *str*]                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                      | metadata contains additional key-value pairs for storing extra information                                                                                              |
| `name`                                                                                                                                                                  | *Optional[str]*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                      | name is the full name or company name of the customer                                                                                                                   |
| `parent_customer_external_id`                                                                                                                                           | *Optional[str]*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                      | parent_customer_external_id is the external ID of the parent customer from your system<br/>Exactly one of parent_customer_id or parent_customer_external_id may be provided |
| `parent_customer_id`                                                                                                                                                    | *Optional[str]*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                      | parent_customer_id is the internal FlexPrice ID of the parent customer                                                                                                  |
| `tax_rate_overrides`                                                                                                                                                    | List[[models.DtoTaxRateOverride](../../models/dtotaxrateoverride.md)]                                                                                                   | :heavy_minus_sign:                                                                                                                                                      | tax_rate_overrides contains tax rate configurations to be linked to this customer                                                                                       |
| `retries`                                                                                                                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                        | :heavy_minus_sign:                                                                                                                                                      | Configuration to override the default retry behavior of the client.                                                                                                     |

### Response

**[models.DtoCustomerResponse](../../models/dtocustomerresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_by_lookup_key

Get a customer by lookup key (external_id)

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/lookup/{lookup_key}" method="get" path="/customers/lookup/{lookup_key}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.customers.get_by_lookup_key(lookup_key="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `lookup_key`                                                        | *str*                                                               | :heavy_check_mark:                                                  | Customer Lookup Key (external_id)                                   |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoCustomerResponse](../../models/dtocustomerresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## search

List customers by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/customers/search" method="post" path="/customers/search" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.customers.search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `customer_ids`                                                                        | List[*str*]                                                                           | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `email`                                                                               | *Optional[str]*                                                                       | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `end_time`                                                                            | *Optional[str]*                                                                       | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `expand`                                                                              | *Optional[str]*                                                                       | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `external_id`                                                                         | *Optional[str]*                                                                       | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `external_ids`                                                                        | List[*str*]                                                                           | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `filters`                                                                             | List[[models.TypesFilterCondition](../../models/typesfiltercondition.md)]             | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `limit`                                                                               | *Optional[int]*                                                                       | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `offset`                                                                              | *Optional[int]*                                                                       | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `order`                                                                               | [Optional[models.TypesCustomerFilterOrder]](../../models/typescustomerfilterorder.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `parent_customer_ids`                                                                 | List[*str*]                                                                           | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `sort`                                                                                | List[[models.TypesSortCondition](../../models/typessortcondition.md)]                 | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `start_time`                                                                          | *Optional[str]*                                                                       | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `status`                                                                              | [Optional[models.TypesStatus]](../../models/typesstatus.md)                           | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `retries`                                                                             | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                      | :heavy_minus_sign:                                                                    | Configuration to override the default retry behavior of the client.                   |

### Response

**[models.DtoListCustomersResponse](../../models/dtolistcustomersresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_usage_summary

Get customer usage summary by customer_id or customer_lookup_key (external_customer_id)

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/usage" method="get" path="/customers/usage" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.customers.get_usage_summary()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `customer_id`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `customer_lookup_key`                                               | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `feature_ids`                                                       | List[*str*]                                                         | :heavy_minus_sign:                                                  | N/A                                                                 |
| `feature_lookup_keys`                                               | List[*str*]                                                         | :heavy_minus_sign:                                                  | N/A                                                                 |
| `subscription_ids`                                                  | List[*str*]                                                         | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoCustomerUsageSummaryResponse](../../models/dtocustomerusagesummaryresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_by_id

Get a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/{id}" method="get" path="/customers/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.customers.get_by_id(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Customer ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoCustomerResponse](../../models/dtocustomerresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update

Update a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/customers/{id}" method="put" path="/customers/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.customers.update(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `id`                                                                                                                                                                                                                                              | *str*                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                | Customer ID                                                                                                                                                                                                                                       |
| `address_city`                                                                                                                                                                                                                                    | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | address_city is the updated city name with maximum 100 characters                                                                                                                                                                                 |
| `address_country`                                                                                                                                                                                                                                 | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | address_country is the updated two-letter ISO 3166-1 alpha-2 country code                                                                                                                                                                         |
| `address_line1`                                                                                                                                                                                                                                   | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | address_line1 is the updated primary address line with maximum 255 characters                                                                                                                                                                     |
| `address_line2`                                                                                                                                                                                                                                   | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | address_line2 is the updated secondary address line with maximum 255 characters                                                                                                                                                                   |
| `address_postal_code`                                                                                                                                                                                                                             | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | address_postal_code is the updated postal code with maximum 20 characters                                                                                                                                                                         |
| `address_state`                                                                                                                                                                                                                                   | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | address_state is the updated state, province, or region name with maximum 100 characters                                                                                                                                                          |
| `email`                                                                                                                                                                                                                                           | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | email is the updated email address and must be a valid email format if provided                                                                                                                                                                   |
| `external_id`                                                                                                                                                                                                                                     | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | external_id is the updated external identifier for the customer                                                                                                                                                                                   |
| `integration_entity_mapping`                                                                                                                                                                                                                      | List[[models.DtoIntegrationEntityMapping](../../models/dtointegrationentitymapping.md)]                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                | integration_entity_mapping contains provider integration mappings for this customer                                                                                                                                                               |
| `metadata`                                                                                                                                                                                                                                        | Dict[str, *str*]                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                | metadata contains updated key-value pairs that will replace existing metadata                                                                                                                                                                     |
| `name`                                                                                                                                                                                                                                            | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | name is the updated name or company name for the customer                                                                                                                                                                                         |
| `parent_customer_external_id`                                                                                                                                                                                                                     | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | parent_customer_external_id is the external ID of the parent customer from your system<br/>Exactly one of parent_customer_id or parent_customer_external_id may be provided<br/>If you provide the external ID, the parent customer value will be ignored |
| `parent_customer_id`                                                                                                                                                                                                                              | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | parent_customer_id is the internal FlexPrice ID of the parent customer                                                                                                                                                                            |
| `retries`                                                                                                                                                                                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                | Configuration to override the default retry behavior of the client.                                                                                                                                                                               |

### Response

**[models.DtoCustomerResponse](../../models/dtocustomerresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete

Delete a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/customers/{id}" method="delete" path="/customers/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    f_client.customers.delete(id="<id>")

    # Use the SDK ...

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Customer ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_entitlements

Get customer entitlements

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/{id}/entitlements" method="get" path="/customers/{id}/entitlements" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.customers.get_entitlements(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Customer ID                                                         |
| `feature_ids`                                                       | List[*str*]                                                         | :heavy_minus_sign:                                                  | N/A                                                                 |
| `subscription_ids`                                                  | List[*str*]                                                         | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoCustomerEntitlementsResponse](../../models/dtocustomerentitlementsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_upcoming_grants

Get upcoming credit grant applications for a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/{id}/grants/upcoming" method="get" path="/customers/{id}/grants/upcoming" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.customers.get_upcoming_grants(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Customer ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoListCreditGrantApplicationsResponse](../../models/dtolistcreditgrantapplicationsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |