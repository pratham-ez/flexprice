# Customers

## Overview

### Available Operations

* [get_customers](#get_customers) - Get customers
* [post_customers](#post_customers) - Create a customer
* [get_customers_external_external_id_](#get_customers_external_external_id_) - Get a customer by external id
* [post_customers_search](#post_customers_search) - List customers by filter
* [get_customers_usage](#get_customers_usage) - Get customer usage summary
* [get_customers_id_](#get_customers_id_) - Get a customer
* [put_customers_id_](#put_customers_id_) - Update a customer
* [delete_customers_id_](#delete_customers_id_) - Delete a customer
* [get_customers_id_entitlements](#get_customers_id_entitlements) - Get customer entitlements
* [get_customers_id_grants_upcoming](#get_customers_id_grants_upcoming) - Get upcoming credit grant applications

## get_customers

Get customers

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers" method="get" path="/customers" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.customers.get_customers()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `customer_ids`                                                                           | List[*str*]                                                                              | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `email`                                                                                  | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `end_time`                                                                               | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `expand`                                                                                 | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `external_id`                                                                            | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `external_ids`                                                                           | List[*str*]                                                                              | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `limit`                                                                                  | *Optional[int]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `offset`                                                                                 | *Optional[int]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `order`                                                                                  | [Optional[operations.GetCustomersOrder]](../../models/operations/getcustomersorder.md)   | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `parent_customer_ids`                                                                    | List[*str*]                                                                              | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `start_time`                                                                             | *Optional[str]*                                                                          | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `status`                                                                                 | [Optional[operations.GetCustomersStatus]](../../models/operations/getcustomersstatus.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `retries`                                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                         | :heavy_minus_sign:                                                                       | Configuration to override the default retry behavior of the client.                      |

### Response

**[components.DtoListCustomersResponse](../../models/components/dtolistcustomersresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_customers

Create a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/customers" method="post" path="/customers" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.customers.post_customers(external_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                                                     | Type                                                                                                                                                                                                          | Required                                                                                                                                                                                                      | Description                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `external_id`                                                                                                                                                                                                 | *str*                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                            | external_id is the unique identifier from your system to reference this customer (required)                                                                                                                   |
| `address_city`                                                                                                                                                                                                | *Optional[str]*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                            | address_city is the city name with maximum 100 characters                                                                                                                                                     |
| `address_country`                                                                                                                                                                                             | *Optional[str]*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                            | address_country is the two-letter ISO 3166-1 alpha-2 country code                                                                                                                                             |
| `address_line1`                                                                                                                                                                                               | *Optional[str]*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                            | address_line1 is the primary address line with maximum 255 characters                                                                                                                                         |
| `address_line2`                                                                                                                                                                                               | *Optional[str]*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                            | address_line2 is the secondary address line with maximum 255 characters                                                                                                                                       |
| `address_postal_code`                                                                                                                                                                                         | *Optional[str]*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                            | address_postal_code is the ZIP code or postal code with maximum 20 characters                                                                                                                                 |
| `address_state`                                                                                                                                                                                               | *Optional[str]*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                            | address_state is the state, province, or region name with maximum 100 characters                                                                                                                              |
| `email`                                                                                                                                                                                                       | *Optional[str]*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                            | email is the customer's email address and must be a valid email format if provided                                                                                                                            |
| `integration_entity_mapping`                                                                                                                                                                                  | List[[components.DtoIntegrationEntityMapping](../../models/components/dtointegrationentitymapping.md)]                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                            | integration_entity_mapping contains provider integration mappings for this customer                                                                                                                           |
| `metadata`                                                                                                                                                                                                    | Dict[str, *str*]                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                            | metadata contains additional key-value pairs for storing extra information                                                                                                                                    |
| `name`                                                                                                                                                                                                        | *Optional[str]*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                            | name is the full name or company name of the customer                                                                                                                                                         |
| `parent_customer_external_id`                                                                                                                                                                                 | *Optional[str]*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                            | parent_customer_external_id is the external ID of the parent customer from your system<br/>Exactly one of parent_customer_id or parent_customer_external_id may be provided                                   |
| `parent_customer_id`                                                                                                                                                                                          | *Optional[str]*                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                            | parent_customer_id is the internal FlexPrice ID of the parent customer                                                                                                                                        |
| `skip_onboarding_workflow`                                                                                                                                                                                    | *Optional[bool]*                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                            | skip_onboarding_workflow when true, prevents the customer onboarding workflow from being triggered<br/>This is used internally when a customer is created via a workflow to prevent infinite loops<br/>Default: false |
| `tax_rate_overrides`                                                                                                                                                                                          | List[[components.DtoTaxRateOverride](../../models/components/dtotaxrateoverride.md)]                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                            | tax_rate_overrides contains tax rate configurations to be linked to this customer                                                                                                                             |
| `retries`                                                                                                                                                                                                     | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                            | Configuration to override the default retry behavior of the client.                                                                                                                                           |

### Response

**[components.DtoCustomerResponse](../../models/components/dtocustomerresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_customers_external_external_id_

Get a customer by external id

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/external/{external_id}" method="get" path="/customers/external/{external_id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.customers.get_customers_external_external_id_(external_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `external_id`                                                       | *str*                                                               | :heavy_check_mark:                                                  | Customer External ID                                                |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCustomerResponse](../../models/components/dtocustomerresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_customers_search

List customers by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/customers/search" method="post" path="/customers/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.customers.post_customers_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `customer_ids`                                                                                       | List[*str*]                                                                                          | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `email`                                                                                              | *Optional[str]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `end_time`                                                                                           | *Optional[str]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `expand`                                                                                             | *Optional[str]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `external_id`                                                                                        | *Optional[str]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `external_ids`                                                                                       | List[*str*]                                                                                          | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `filters`                                                                                            | List[[components.TypesFilterCondition](../../models/components/typesfiltercondition.md)]             | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `limit`                                                                                              | *Optional[int]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `offset`                                                                                             | *Optional[int]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `order`                                                                                              | [Optional[components.TypesCustomerFilterOrder]](../../models/components/typescustomerfilterorder.md) | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `parent_customer_ids`                                                                                | List[*str*]                                                                                          | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `sort`                                                                                               | List[[components.TypesSortCondition](../../models/components/typessortcondition.md)]                 | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `start_time`                                                                                         | *Optional[str]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `status`                                                                                             | [Optional[components.TypesStatus]](../../models/components/typesstatus.md)                           | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `retries`                                                                                            | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                     | :heavy_minus_sign:                                                                                   | Configuration to override the default retry behavior of the client.                                  |

### Response

**[components.DtoListCustomersResponse](../../models/components/dtolistcustomersresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_customers_usage

Get customer usage summary by customer_id or customer_lookup_key (external_customer_id)

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/usage" method="get" path="/customers/usage" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.customers.get_customers_usage()

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

**[components.DtoCustomerUsageSummaryResponse](../../models/components/dtocustomerusagesummaryresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_customers_id_

Get a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/{id}" method="get" path="/customers/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.customers.get_customers_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Customer ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCustomerResponse](../../models/components/dtocustomerresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_customers_id_

Update a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/customers/{id}" method="put" path="/customers/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.customers.put_customers_id_(id="<id>")

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
| `integration_entity_mapping`                                                                                                                                                                                                                      | List[[components.DtoIntegrationEntityMapping](../../models/components/dtointegrationentitymapping.md)]                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                | integration_entity_mapping contains provider integration mappings for this customer                                                                                                                                                               |
| `metadata`                                                                                                                                                                                                                                        | Dict[str, *str*]                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                | metadata contains updated key-value pairs that will replace existing metadata                                                                                                                                                                     |
| `name`                                                                                                                                                                                                                                            | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | name is the updated name or company name for the customer                                                                                                                                                                                         |
| `parent_customer_external_id`                                                                                                                                                                                                                     | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | parent_customer_external_id is the external ID of the parent customer from your system<br/>Exactly one of parent_customer_id or parent_customer_external_id may be provided<br/>If you provide the external ID, the parent customer value will be ignored |
| `parent_customer_id`                                                                                                                                                                                                                              | *Optional[str]*                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                | parent_customer_id is the internal FlexPrice ID of the parent customer                                                                                                                                                                            |
| `retries`                                                                                                                                                                                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                | Configuration to override the default retry behavior of the client.                                                                                                                                                                               |

### Response

**[components.DtoCustomerResponse](../../models/components/dtocustomerresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_customers_id_

Delete a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/customers/{id}" method="delete" path="/customers/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    flex_price.customers.delete_customers_id_(id="<id>")

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
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_customers_id_entitlements

Get customer entitlements

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/{id}/entitlements" method="get" path="/customers/{id}/entitlements" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.customers.get_customers_id_entitlements(id="<id>")

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

**[components.DtoCustomerEntitlementsResponse](../../models/components/dtocustomerentitlementsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_customers_id_grants_upcoming

Get upcoming credit grant applications for a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/{id}/grants/upcoming" method="get" path="/customers/{id}/grants/upcoming" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.customers.get_customers_id_grants_upcoming(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Customer ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoListCreditGrantApplicationsResponse](../../models/components/dtolistcreditgrantapplicationsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |