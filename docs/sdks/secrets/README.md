# Secrets

## Overview

### Available Operations

* [list_api_keys](#list_api_keys) - List API keys
* [create_api_key](#create_api_key) - Create a new API key
* [delete_api_key](#delete_api_key) - Delete an API key

## list_api_keys

Get a paginated list of API keys

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/secrets/api/keys" method="get" path="/secrets/api/keys" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.secrets.list_api_keys()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `limit`                                                             | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | Limit                                                               |
| `offset`                                                            | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | Offset                                                              |
| `status`                                                            | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Status (published/archived)                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoListSecretsResponse](../../models/dtolistsecretsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create_api_key

Create a new API key. Provide 'service_account_id' in body to create API key for a service account, otherwise creates for authenticated user.

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/secrets/api/keys" method="post" path="/secrets/api/keys" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.secrets.create_api_key(name="<value>", type_="integration")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `name`                                                              | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `type`                                                              | [models.TypesSecretType](../../models/typessecrettype.md)           | :heavy_check_mark:                                                  | N/A                                                                 |
| `expires_at`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `service_account_id`                                                | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoCreateAPIKeyResponse](../../models/dtocreateapikeyresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_api_key

Delete an API key by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/secrets/api/keys/{id}" method="delete" path="/secrets/api/keys/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    f_client.secrets.delete_api_key(id="<id>")

    # Use the SDK ...

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | API key ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 404                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |