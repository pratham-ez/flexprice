# EntityIntegrationMappings

## Overview

### Available Operations

* [list](#list) - List entity integration mappings
* [create](#create) - Create entity integration mapping
* [get_by_id](#get_by_id) - Get entity integration mapping
* [delete](#delete) - Delete entity integration mapping

## list

Retrieve a list of entity integration mappings with optional filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/entity-integration-mappings" method="get" path="/entity-integration-mappings" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entity_integration_mappings.list()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `entity_id`                                                         | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by FlexPrice entity ID                                       |
| `entity_type`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by entity type                                               |
| `provider_type`                                                     | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by provider type                                             |
| `provider_entity_id`                                                | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by provider entity ID                                        |
| `limit`                                                             | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | Number of results to return (default: 20, max: 100)                 |
| `offset`                                                            | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | Pagination offset (default: 0)                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoListEntityIntegrationMappingsResponse](../../models/dtolistentityintegrationmappingsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create

Create a new entity integration mapping

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/entity-integration-mappings" method="post" path="/entity-integration-mappings" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entity_integration_mappings.create(entity_id="<id>", entity_type="addon", provider_entity_id="<id>", provider_type="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `entity_id`                                                                     | *str*                                                                           | :heavy_check_mark:                                                              | N/A                                                                             |
| `entity_type`                                                                   | [models.TypesIntegrationEntityType](../../models/typesintegrationentitytype.md) | :heavy_check_mark:                                                              | N/A                                                                             |
| `provider_entity_id`                                                            | *str*                                                                           | :heavy_check_mark:                                                              | N/A                                                                             |
| `provider_type`                                                                 | *str*                                                                           | :heavy_check_mark:                                                              | N/A                                                                             |
| `metadata`                                                                      | Dict[str, *Any*]                                                                | :heavy_minus_sign:                                                              | N/A                                                                             |
| `retries`                                                                       | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                | :heavy_minus_sign:                                                              | Configuration to override the default retry behavior of the client.             |

### Response

**[models.DtoEntityIntegrationMappingResponse](../../models/dtoentityintegrationmappingresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 409                | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_by_id

Retrieve a specific entity integration mapping by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/entity-integration-mappings/{id}" method="get" path="/entity-integration-mappings/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entity_integration_mappings.get_by_id(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Entity integration mapping ID                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoEntityIntegrationMappingResponse](../../models/dtoentityintegrationmappingresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 404                | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete

Delete an entity integration mapping

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/entity-integration-mappings/{id}" method="delete" path="/entity-integration-mappings/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    f_client.entity_integration_mappings.delete(id="<id>")

    # Use the SDK ...

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Entity integration mapping ID                                       |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 404                | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |