# Connections

## Overview

### Available Operations

* [get_connections](#get_connections) - Get connections
* [post_connections_search](#post_connections_search) - List connections by filter
* [get_connections_id_](#get_connections_id_) - Get a connection
* [put_connections_id_](#put_connections_id_) - Update a connection
* [delete_connections_id_](#delete_connections_id_) - Delete a connection

## get_connections

Get a list of connections

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/connections" method="get" path="/connections" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.connections.get_connections()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `connection_ids`                                                                             | List[*str*]                                                                                  | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `end_time`                                                                                   | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `expand`                                                                                     | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `limit`                                                                                      | *Optional[int]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `offset`                                                                                     | *Optional[int]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `order`                                                                                      | [Optional[operations.GetConnectionsOrder]](../../models/operations/getconnectionsorder.md)   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `provider_type`                                                                              | [Optional[operations.ProviderType]](../../models/operations/providertype.md)                 | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `start_time`                                                                                 | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `status`                                                                                     | [Optional[operations.GetConnectionsStatus]](../../models/operations/getconnectionsstatus.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `retries`                                                                                    | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                             | :heavy_minus_sign:                                                                           | Configuration to override the default retry behavior of the client.                          |

### Response

**[components.DtoListConnectionsResponse](../../models/components/dtolistconnectionsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_connections_search

List connections by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/connections/search" method="post" path="/connections/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.connections.post_connections_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `connection_ids`                                                                                         | List[*str*]                                                                                              | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `end_time`                                                                                               | *Optional[str]*                                                                                          | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `expand`                                                                                                 | *Optional[str]*                                                                                          | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `filters`                                                                                                | List[[components.TypesFilterCondition](../../models/components/typesfiltercondition.md)]                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `limit`                                                                                                  | *Optional[int]*                                                                                          | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `offset`                                                                                                 | *Optional[int]*                                                                                          | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `order`                                                                                                  | [Optional[components.TypesConnectionFilterOrder]](../../models/components/typesconnectionfilterorder.md) | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `provider_type`                                                                                          | [Optional[components.TypesSecretProvider]](../../models/components/typessecretprovider.md)               | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `sort`                                                                                                   | List[[components.TypesSortCondition](../../models/components/typessortcondition.md)]                     | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `start_time`                                                                                             | *Optional[str]*                                                                                          | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `status`                                                                                                 | [Optional[components.TypesStatus]](../../models/components/typesstatus.md)                               | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `retries`                                                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                         | :heavy_minus_sign:                                                                                       | Configuration to override the default retry behavior of the client.                                      |

### Response

**[components.DtoListConnectionsResponse](../../models/components/dtolistconnectionsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_connections_id_

Get a connection by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/connections/{id}" method="get" path="/connections/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.connections.get_connections_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Connection ID                                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoConnectionResponse](../../models/components/dtoconnectionresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_connections_id_

Update a connection by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/connections/{id}" method="put" path="/connections/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.connections.put_connections_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `id`                                                                                               | *str*                                                                                              | :heavy_check_mark:                                                                                 | Connection ID                                                                                      |
| `encrypted_secret_data`                                                                            | [Optional[components.TypesConnectionMetadata]](../../models/components/typesconnectionmetadata.md) | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `metadata`                                                                                         | Dict[str, *Any*]                                                                                   | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `name`                                                                                             | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `sync_config`                                                                                      | [Optional[components.TypesSyncConfig]](../../models/components/typessyncconfig.md)                 | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `retries`                                                                                          | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                   | :heavy_minus_sign:                                                                                 | Configuration to override the default retry behavior of the client.                                |

### Response

**[components.DtoConnectionResponse](../../models/components/dtoconnectionresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_connections_id_

Delete a connection by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/connections/{id}" method="delete" path="/connections/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    flex_price.connections.delete_connections_id_(id="<id>")

    # Use the SDK ...

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Connection ID                                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |