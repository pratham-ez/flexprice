# Entitlements

## Overview

### Available Operations

* [get_for_addon](#get_for_addon) - Get addon entitlements
* [list](#list) - Get entitlements
* [create](#create) - Create a new entitlement
* [bulk_create](#bulk_create) - Create multiple entitlements in bulk
* [filter](#filter) - List entitlements by filter
* [get](#get) - Get an entitlement by ID
* [update](#update) - Update an entitlement
* [delete](#delete) - Delete an entitlement

## get_for_addon

Get all entitlements for an addon

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/addons/{id}/entitlements" method="get" path="/addons/{id}/entitlements" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entitlements.get_for_addon(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Addon ID                                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoListEntitlementsResponse](../../models/dtolistentitlementsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## list

Get entitlements with the specified filter

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/entitlements" method="get" path="/entitlements" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entitlements.list()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `request`                                                               | [models.GetEntitlementsRequest](../../models/getentitlementsrequest.md) | :heavy_check_mark:                                                      | The request object to use for the request.                              |
| `retries`                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)        | :heavy_minus_sign:                                                      | Configuration to override the default retry behavior of the client.     |

### Response

**[models.DtoListEntitlementsResponse](../../models/dtolistentitlementsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create

Create a new entitlement with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/entitlements" method="post" path="/entitlements" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entitlements.create(feature_id="<id>", feature_type="boolean")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `feature_id`                                                                                          | *str*                                                                                                 | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `feature_type`                                                                                        | [models.TypesFeatureType](../../models/typesfeaturetype.md)                                           | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `entity_id`                                                                                           | *Optional[str]*                                                                                       | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `entity_type`                                                                                         | [Optional[models.TypesEntitlementEntityType]](../../models/typesentitlemententitytype.md)             | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `is_enabled`                                                                                          | *Optional[bool]*                                                                                      | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `is_soft_limit`                                                                                       | *Optional[bool]*                                                                                      | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `parent_entitlement_id`                                                                               | *Optional[str]*                                                                                       | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `plan_id`                                                                                             | *Optional[str]*                                                                                       | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `static_value`                                                                                        | *Optional[str]*                                                                                       | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `usage_limit`                                                                                         | *Optional[int]*                                                                                       | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `usage_reset_period`                                                                                  | [Optional[models.TypesEntitlementUsageResetPeriod]](../../models/typesentitlementusageresetperiod.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `retries`                                                                                             | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                      | :heavy_minus_sign:                                                                                    | Configuration to override the default retry behavior of the client.                                   |

### Response

**[models.DtoEntitlementResponse](../../models/dtoentitlementresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## bulk_create

Create multiple entitlements with the specified configurations

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/entitlements/bulk" method="post" path="/entitlements/bulk" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entitlements.bulk_create(items=[])

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `items`                                                                                 | List[[models.DtoCreateEntitlementRequest](../../models/dtocreateentitlementrequest.md)] | :heavy_check_mark:                                                                      | N/A                                                                                     |
| `retries`                                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                        | :heavy_minus_sign:                                                                      | Configuration to override the default retry behavior of the client.                     |

### Response

**[models.DtoCreateBulkEntitlementResponse](../../models/dtocreatebulkentitlementresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## filter

List entitlements by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/entitlements/search" method="post" path="/entitlements/search" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entitlements.filter()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `end_time`                                                                                  | *Optional[str]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `entity_ids`                                                                                | List[*str*]                                                                                 | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `entity_type`                                                                               | [Optional[models.TypesEntitlementEntityType]](../../models/typesentitlemententitytype.md)   | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `expand`                                                                                    | *Optional[str]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `feature_ids`                                                                               | List[*str*]                                                                                 | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `feature_type`                                                                              | [Optional[models.TypesFeatureType]](../../models/typesfeaturetype.md)                       | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `filters`                                                                                   | List[[models.TypesFilterCondition](../../models/typesfiltercondition.md)]                   | :heavy_minus_sign:                                                                          | Specific filters for entitlements                                                           |
| `is_enabled`                                                                                | *Optional[bool]*                                                                            | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `limit`                                                                                     | *Optional[int]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `offset`                                                                                    | *Optional[int]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `order`                                                                                     | [Optional[models.TypesEntitlementFilterOrder]](../../models/typesentitlementfilterorder.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `plan_ids`                                                                                  | List[*str*]                                                                                 | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `sort`                                                                                      | List[[models.TypesSortCondition](../../models/typessortcondition.md)]                       | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `start_time`                                                                                | *Optional[str]*                                                                             | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `status`                                                                                    | [Optional[models.TypesStatus]](../../models/typesstatus.md)                                 | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `retries`                                                                                   | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                            | :heavy_minus_sign:                                                                          | Configuration to override the default retry behavior of the client.                         |

### Response

**[models.DtoListEntitlementsResponse](../../models/dtolistentitlementsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get

Get an entitlement by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/entitlements/{id}" method="get" path="/entitlements/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entitlements.get(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Entitlement ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoEntitlementResponse](../../models/dtoentitlementresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update

Update an entitlement with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/entitlements/{id}" method="put" path="/entitlements/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entitlements.update(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `id`                                                                                                  | *str*                                                                                                 | :heavy_check_mark:                                                                                    | Entitlement ID                                                                                        |
| `is_enabled`                                                                                          | *Optional[bool]*                                                                                      | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `is_soft_limit`                                                                                       | *Optional[bool]*                                                                                      | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `static_value`                                                                                        | *Optional[str]*                                                                                       | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `usage_limit`                                                                                         | *Optional[int]*                                                                                       | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `usage_reset_period`                                                                                  | [Optional[models.TypesEntitlementUsageResetPeriod]](../../models/typesentitlementusageresetperiod.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `retries`                                                                                             | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                      | :heavy_minus_sign:                                                                                    | Configuration to override the default retry behavior of the client.                                   |

### Response

**[models.DtoEntitlementResponse](../../models/dtoentitlementresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete

Delete an entitlement

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/entitlements/{id}" method="delete" path="/entitlements/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.entitlements.delete(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Entitlement ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, models.GinH]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |