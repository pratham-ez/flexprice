# Users

## Overview

### Available Operations

* [create_service_account](#create_service_account) - Create service account
* [get_with_me](#get_with_me) - Get user info
* [search](#search) - List service accounts with filters

## create_service_account

Create a new service account with required roles. Only service accounts can be created via this endpoint.

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/users" method="post" path="/users" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.users.create_service_account(roles=[], type_="service_account")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `roles`                                                             | List[*str*]                                                         | :heavy_check_mark:                                                  | Roles are required                                                  |
| `type`                                                              | [models.TypesUserType](../../models/typesusertype.md)               | :heavy_check_mark:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoUserResponse](../../models/dtouserresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_with_me

Get the current user's information

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/users/me" method="get" path="/users/me" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.users.get_with_me()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoUserResponse](../../models/dtouserresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 401                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## search

Search and filter service accounts by type, roles, etc.

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/users/search" method="post" path="/users/search" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.users.search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `end_time`                                                                    | *Optional[str]*                                                               | :heavy_minus_sign:                                                            | N/A                                                                           |
| `expand`                                                                      | *Optional[str]*                                                               | :heavy_minus_sign:                                                            | N/A                                                                           |
| `filters`                                                                     | List[[models.TypesFilterCondition](../../models/typesfiltercondition.md)]     | :heavy_minus_sign:                                                            | filters allows complex filtering based on multiple fields                     |
| `limit`                                                                       | *Optional[int]*                                                               | :heavy_minus_sign:                                                            | N/A                                                                           |
| `offset`                                                                      | *Optional[int]*                                                               | :heavy_minus_sign:                                                            | N/A                                                                           |
| `order`                                                                       | [Optional[models.TypesUserFilterOrder]](../../models/typesuserfilterorder.md) | :heavy_minus_sign:                                                            | N/A                                                                           |
| `roles`                                                                       | List[*str*]                                                                   | :heavy_minus_sign:                                                            | N/A                                                                           |
| `sort`                                                                        | List[[models.TypesSortCondition](../../models/typessortcondition.md)]         | :heavy_minus_sign:                                                            | N/A                                                                           |
| `start_time`                                                                  | *Optional[str]*                                                               | :heavy_minus_sign:                                                            | N/A                                                                           |
| `status`                                                                      | [Optional[models.TypesStatus]](../../models/typesstatus.md)                   | :heavy_minus_sign:                                                            | N/A                                                                           |
| `type`                                                                        | [Optional[models.TypesUserType]](../../models/typesusertype.md)               | :heavy_minus_sign:                                                            | N/A                                                                           |
| `user_ids`                                                                    | List[*str*]                                                                   | :heavy_minus_sign:                                                            | Specific filters for users                                                    |
| `retries`                                                                     | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)              | :heavy_minus_sign:                                                            | Configuration to override the default retry behavior of the client.           |

### Response

**[models.DtoListUsersResponse](../../models/dtolistusersresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |