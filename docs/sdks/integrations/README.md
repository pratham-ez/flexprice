# Integrations

## Overview

### Available Operations

* [get_secrets_integrations_by_provider_provider_](#get_secrets_integrations_by_provider_provider_) - Get integration details
* [post_secrets_integrations_create_provider_](#post_secrets_integrations_create_provider_) - Create or update an integration
* [get_secrets_integrations_linked](#get_secrets_integrations_linked) - List linked integrations
* [delete_secrets_integrations_id_](#delete_secrets_integrations_id_) - Delete an integration

## get_secrets_integrations_by_provider_provider_

Get details of a specific integration

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/secrets/integrations/by-provider/{provider}" method="get" path="/secrets/integrations/by-provider/{provider}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.integrations.get_secrets_integrations_by_provider_provider_(provider="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `provider`                                                          | *str*                                                               | :heavy_check_mark:                                                  | Integration provider                                                |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSecretResponse](../../models/components/dtosecretresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 404                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_secrets_integrations_create_provider_

Create or update integration credentials

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/secrets/integrations/create/{provider}" method="post" path="/secrets/integrations/create/{provider}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.integrations.post_secrets_integrations_create_provider_(provider_param="<value>", credentials={
        "key": "<value>",
    }, name="<value>", provider="chargebee")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `provider_param`                                                                 | *str*                                                                            | :heavy_check_mark:                                                               | Integration provider                                                             |
| `credentials`                                                                    | Dict[str, *str*]                                                                 | :heavy_check_mark:                                                               | N/A                                                                              |
| `name`                                                                           | *str*                                                                            | :heavy_check_mark:                                                               | N/A                                                                              |
| `provider`                                                                       | [components.TypesSecretProvider](../../models/components/typessecretprovider.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `retries`                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                 | :heavy_minus_sign:                                                               | Configuration to override the default retry behavior of the client.              |

### Response

**[components.DtoSecretResponse](../../models/components/dtosecretresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_secrets_integrations_linked

Get a list of unique providers which have a valid linked integration secret

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/secrets/integrations/linked" method="get" path="/secrets/integrations/linked" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.integrations.get_secrets_integrations_linked()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoLinkedIntegrationsResponse](../../models/components/dtolinkedintegrationsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_secrets_integrations_id_

Delete integration credentials

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/secrets/integrations/{id}" method="delete" path="/secrets/integrations/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    flex_price.integrations.delete_secrets_integrations_id_(id="<id>")

    # Use the SDK ...

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Integration ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 404                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |