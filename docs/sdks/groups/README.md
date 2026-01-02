# Groups

## Overview

### Available Operations

* [post_groups](#post_groups) - Create a group
* [post_groups_search](#post_groups_search) - Get groups
* [get_groups_id_](#get_groups_id_) - Get a group
* [delete_groups_id_](#delete_groups_id_) - Delete a group

## post_groups

Create a new group for organizing entities (prices, plans, customers, etc.)

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/groups" method="post" path="/groups" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.groups.post_groups(entity_type="<value>", lookup_key="<value>", name="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `entity_type`                                                       | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `lookup_key`                                                        | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `name`                                                              | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoGroupResponse](../../models/components/dtogroupresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_groups_search

Get groups with optional filtering via query parameters

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/groups/search" method="post" path="/groups/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.groups.post_groups_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `entity_type`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by entity type (e.g., 'price')                               |
| `name`                                                              | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by group name (contains search)                              |
| `lookup_key`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Filter by lookup key (exact match)                                  |
| `limit`                                                             | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | Number of items to return (default: 20)                             |
| `offset`                                                            | *Optional[int]*                                                     | :heavy_minus_sign:                                                  | Number of items to skip (default: 0)                                |
| `sort_by`                                                           | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Field to sort by (name, created_at, updated_at)                     |
| `sort_order`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Sort order (asc, desc)                                              |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoListGroupsResponse](../../models/components/dtolistgroupsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_groups_id_

Get a group by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/groups/{id}" method="get" path="/groups/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.groups.get_groups_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Group ID                                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoGroupResponse](../../models/components/dtogroupresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_groups_id_

Delete a group and remove all entity associations

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/groups/{id}" method="delete" path="/groups/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    flex_price.groups.delete_groups_id_(id="<id>")

    # Use the SDK ...

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Group ID                                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |