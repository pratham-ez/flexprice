# Connections

## Overview

### Available Operations

* [list](#list) - Get connections
* [search](#search) - List connections by filter
* [get_by_id](#get_by_id) - Get a connection
* [update](#update) - Update a connection
* [delete](#delete) - Delete a connection

## list

Get a list of connections

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/connections" method="get" path="/connections" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.connections.list()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `connection_ids`                                                                                | List[*str*]                                                                                     | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `end_time`                                                                                      | *Optional[str]*                                                                                 | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `expand`                                                                                        | *Optional[str]*                                                                                 | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `limit`                                                                                         | *Optional[int]*                                                                                 | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `offset`                                                                                        | *Optional[int]*                                                                                 | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `order`                                                                                         | [Optional[models.GetConnectionsQueryParamOrder]](../../models/getconnectionsqueryparamorder.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `provider_type`                                                                                 | [Optional[models.ProviderType]](../../models/providertype.md)                                   | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `start_time`                                                                                    | *Optional[str]*                                                                                 | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `status`                                                                                        | [Optional[models.QueryParamStatus]](../../models/queryparamstatus.md)                           | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `retries`                                                                                       | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                | :heavy_minus_sign:                                                                              | Configuration to override the default retry behavior of the client.                             |

### Response

**[models.DtoListConnectionsResponse](../../models/dtolistconnectionsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## search

List connections by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/connections/search" method="post" path="/connections/search" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.connections.search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `connection_ids`                                                                          | List[*str*]                                                                               | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `end_time`                                                                                | *Optional[str]*                                                                           | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `expand`                                                                                  | *Optional[str]*                                                                           | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `filters`                                                                                 | List[[models.TypesFilterCondition](../../models/typesfiltercondition.md)]                 | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `limit`                                                                                   | *Optional[int]*                                                                           | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `offset`                                                                                  | *Optional[int]*                                                                           | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `order`                                                                                   | [Optional[models.TypesConnectionFilterOrder]](../../models/typesconnectionfilterorder.md) | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `provider_type`                                                                           | [Optional[models.TypesSecretProvider]](../../models/typessecretprovider.md)               | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `sort`                                                                                    | List[[models.TypesSortCondition](../../models/typessortcondition.md)]                     | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `start_time`                                                                              | *Optional[str]*                                                                           | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `status`                                                                                  | [Optional[models.TypesStatus]](../../models/typesstatus.md)                               | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `retries`                                                                                 | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                          | :heavy_minus_sign:                                                                        | Configuration to override the default retry behavior of the client.                       |

### Response

**[models.DtoListConnectionsResponse](../../models/dtolistconnectionsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_by_id

Get a connection by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/connections/{id}" method="get" path="/connections/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.connections.get_by_id(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Connection ID                                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoConnectionResponse](../../models/dtoconnectionresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update

Update a connection by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/connections/{id}" method="put" path="/connections/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.connections.update(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `id`                                                                                | *str*                                                                               | :heavy_check_mark:                                                                  | Connection ID                                                                       |
| `encrypted_secret_data`                                                             | [Optional[models.TypesConnectionMetadata]](../../models/typesconnectionmetadata.md) | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `metadata`                                                                          | Dict[str, *Any*]                                                                    | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `name`                                                                              | *Optional[str]*                                                                     | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `sync_config`                                                                       | [Optional[models.TypesSyncConfig]](../../models/typessyncconfig.md)                 | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `retries`                                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                    | :heavy_minus_sign:                                                                  | Configuration to override the default retry behavior of the client.                 |

### Response

**[models.DtoConnectionResponse](../../models/dtoconnectionresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete

Delete a connection by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/connections/{id}" method="delete" path="/connections/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    f_client.connections.delete(id="<id>")

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
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |