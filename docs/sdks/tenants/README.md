# Tenants

## Overview

### Available Operations

* [get_tenant_billing](#get_tenant_billing) - Get billing usage for the current tenant
* [post_tenants](#post_tenants) - Create a new tenant
* [put_tenants_update](#put_tenants_update) - Update a tenant
* [get_tenants_id_](#get_tenants_id_) - Get tenant by ID

## get_tenant_billing

Get the subscription and usage details for the current tenant

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/tenant/billing" method="get" path="/tenant/billing" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tenants.get_tenant_billing()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoTenantBillingUsage](../../models/components/dtotenantbillingusage.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_tenants

Create a new tenant

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/tenants" method="post" path="/tenants" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tenants.post_tenants(name="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `name`                                                                                             | *str*                                                                                              | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `billing_details`                                                                                  | [Optional[components.DtoTenantBillingDetails]](../../models/components/dtotenantbillingdetails.md) | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `retries`                                                                                          | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                   | :heavy_minus_sign:                                                                                 | Configuration to override the default retry behavior of the client.                                |

### Response

**[components.DtoTenantResponse](../../models/components/dtotenantresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_tenants_update

Update a tenant's details including name and billing information

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/tenants/update" method="put" path="/tenants/update" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tenants.put_tenants_update()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `billing_details`                                                                                  | [Optional[components.DtoTenantBillingDetails]](../../models/components/dtotenantbillingdetails.md) | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `metadata`                                                                                         | Dict[str, *str*]                                                                                   | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `name`                                                                                             | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `retries`                                                                                          | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                   | :heavy_minus_sign:                                                                                 | Configuration to override the default retry behavior of the client.                                |

### Response

**[components.DtoTenantResponse](../../models/components/dtotenantresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_tenants_id_

Get tenant by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/tenants/{id}" method="get" path="/tenants/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.tenants.get_tenants_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Tenant ID                                                           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoTenantResponse](../../models/components/dtotenantresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 404                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |