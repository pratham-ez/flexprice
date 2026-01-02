# Rbac

## Overview

### Available Operations

* [get_rbac_roles](#get_rbac_roles) - List all RBAC roles
* [get_rbac_roles_id_](#get_rbac_roles_id_) - Get a specific RBAC role

## get_rbac_roles

Returns all available roles with their permissions, names, and descriptions

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/rbac/roles" method="get" path="/rbac/roles" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.rbac.get_rbac_roles()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, Any]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_rbac_roles_id_

Returns details of a specific role including permissions, name, and description

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/rbac/roles/{id}" method="get" path="/rbac/roles/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.rbac.get_rbac_roles_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Role ID                                                             |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, Any]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |