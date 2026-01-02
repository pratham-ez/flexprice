# Entitlements

## Overview

### Available Operations

* [get_addons_id_entitlements](#get_addons_id_entitlements) - Get addon entitlements
* [get_entitlements](#get_entitlements) - Get entitlements
* [post_entitlements](#post_entitlements) - Create a new entitlement
* [post_entitlements_bulk](#post_entitlements_bulk) - Create multiple entitlements in bulk
* [post_entitlements_search](#post_entitlements_search) - List entitlements by filter
* [get_entitlements_id_](#get_entitlements_id_) - Get an entitlement by ID
* [put_entitlements_id_](#put_entitlements_id_) - Update an entitlement
* [delete_entitlements_id_](#delete_entitlements_id_) - Delete an entitlement
* [get_plans_id_entitlements](#get_plans_id_entitlements) - Get plan entitlements

## get_addons_id_entitlements

Get all entitlements for an addon

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/addons/{id}/entitlements" method="get" path="/addons/{id}/entitlements" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.entitlements.get_addons_id_entitlements(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Addon ID                                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoListEntitlementsResponse](../../models/components/dtolistentitlementsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_entitlements

Get entitlements with the specified filter

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/entitlements" method="get" path="/entitlements" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.entitlements.get_entitlements()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `end_time`                                                                                             | *Optional[str]*                                                                                        | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `entity_ids`                                                                                           | List[*str*]                                                                                            | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `entity_type`                                                                                          | [Optional[operations.GetEntitlementsEntityType]](../../models/operations/getentitlementsentitytype.md) | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `expand`                                                                                               | *Optional[str]*                                                                                        | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `feature_ids`                                                                                          | List[*str*]                                                                                            | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `feature_type`                                                                                         | [Optional[operations.FeatureType]](../../models/operations/featuretype.md)                             | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `is_enabled`                                                                                           | *Optional[bool]*                                                                                       | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `limit`                                                                                                | *Optional[int]*                                                                                        | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `offset`                                                                                               | *Optional[int]*                                                                                        | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `order`                                                                                                | [Optional[operations.GetEntitlementsOrder]](../../models/operations/getentitlementsorder.md)           | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `plan_ids`                                                                                             | List[*str*]                                                                                            | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `start_time`                                                                                           | *Optional[str]*                                                                                        | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `status`                                                                                               | [Optional[operations.GetEntitlementsStatus]](../../models/operations/getentitlementsstatus.md)         | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `retries`                                                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                       | :heavy_minus_sign:                                                                                     | Configuration to override the default retry behavior of the client.                                    |

### Response

**[components.DtoListEntitlementsResponse](../../models/components/dtolistentitlementsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_entitlements

Create a new entitlement with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/entitlements" method="post" path="/entitlements" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.entitlements.post_entitlements(feature_id="<id>", feature_type="boolean")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `feature_id`                                                                                                         | *str*                                                                                                                | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `feature_type`                                                                                                       | [components.TypesFeatureType](../../models/components/typesfeaturetype.md)                                           | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `end_date`                                                                                                           | *Optional[str]*                                                                                                      | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `entity_id`                                                                                                          | *Optional[str]*                                                                                                      | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `entity_type`                                                                                                        | [Optional[components.TypesEntitlementEntityType]](../../models/components/typesentitlemententitytype.md)             | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `is_enabled`                                                                                                         | *Optional[bool]*                                                                                                     | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `is_soft_limit`                                                                                                      | *Optional[bool]*                                                                                                     | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `parent_entitlement_id`                                                                                              | *Optional[str]*                                                                                                      | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `plan_id`                                                                                                            | *Optional[str]*                                                                                                      | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `start_date`                                                                                                         | *Optional[str]*                                                                                                      | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `static_value`                                                                                                       | *Optional[str]*                                                                                                      | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `usage_limit`                                                                                                        | *Optional[int]*                                                                                                      | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `usage_reset_period`                                                                                                 | [Optional[components.TypesEntitlementUsageResetPeriod]](../../models/components/typesentitlementusageresetperiod.md) | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `retries`                                                                                                            | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                     | :heavy_minus_sign:                                                                                                   | Configuration to override the default retry behavior of the client.                                                  |

### Response

**[components.DtoEntitlementResponse](../../models/components/dtoentitlementresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_entitlements_bulk

Create multiple entitlements with the specified configurations

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/entitlements/bulk" method="post" path="/entitlements/bulk" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.entitlements.post_entitlements_bulk(items=[])

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `items`                                                                                                | List[[components.DtoCreateEntitlementRequest](../../models/components/dtocreateentitlementrequest.md)] | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `retries`                                                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                       | :heavy_minus_sign:                                                                                     | Configuration to override the default retry behavior of the client.                                    |

### Response

**[components.DtoCreateBulkEntitlementResponse](../../models/components/dtocreatebulkentitlementresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_entitlements_search

List entitlements by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/entitlements/search" method="post" path="/entitlements/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.entitlements.post_entitlements_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `end_time`                                                                                                 | *Optional[str]*                                                                                            | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `entity_ids`                                                                                               | List[*str*]                                                                                                | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `entity_type`                                                                                              | [Optional[components.TypesEntitlementEntityType]](../../models/components/typesentitlemententitytype.md)   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `expand`                                                                                                   | *Optional[str]*                                                                                            | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `feature_ids`                                                                                              | List[*str*]                                                                                                | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `feature_type`                                                                                             | [Optional[components.TypesFeatureType]](../../models/components/typesfeaturetype.md)                       | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `filters`                                                                                                  | List[[components.TypesFilterCondition](../../models/components/typesfiltercondition.md)]                   | :heavy_minus_sign:                                                                                         | Specific filters for entitlements                                                                          |
| `is_enabled`                                                                                               | *Optional[bool]*                                                                                           | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `limit`                                                                                                    | *Optional[int]*                                                                                            | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `offset`                                                                                                   | *Optional[int]*                                                                                            | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `order`                                                                                                    | [Optional[components.TypesEntitlementFilterOrder]](../../models/components/typesentitlementfilterorder.md) | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `plan_ids`                                                                                                 | List[*str*]                                                                                                | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `sort`                                                                                                     | List[[components.TypesSortCondition](../../models/components/typessortcondition.md)]                       | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `start_time`                                                                                               | *Optional[str]*                                                                                            | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `status`                                                                                                   | [Optional[components.TypesStatus]](../../models/components/typesstatus.md)                                 | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `retries`                                                                                                  | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                           | :heavy_minus_sign:                                                                                         | Configuration to override the default retry behavior of the client.                                        |

### Response

**[components.DtoListEntitlementsResponse](../../models/components/dtolistentitlementsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_entitlements_id_

Get an entitlement by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/entitlements/{id}" method="get" path="/entitlements/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.entitlements.get_entitlements_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Entitlement ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoEntitlementResponse](../../models/components/dtoentitlementresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_entitlements_id_

Update an entitlement with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/entitlements/{id}" method="put" path="/entitlements/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.entitlements.put_entitlements_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `id`                                                                                                                 | *str*                                                                                                                | :heavy_check_mark:                                                                                                   | Entitlement ID                                                                                                       |
| `is_enabled`                                                                                                         | *Optional[bool]*                                                                                                     | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `is_soft_limit`                                                                                                      | *Optional[bool]*                                                                                                     | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `static_value`                                                                                                       | *Optional[str]*                                                                                                      | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `usage_limit`                                                                                                        | *Optional[int]*                                                                                                      | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `usage_reset_period`                                                                                                 | [Optional[components.TypesEntitlementUsageResetPeriod]](../../models/components/typesentitlementusageresetperiod.md) | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `retries`                                                                                                            | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                     | :heavy_minus_sign:                                                                                                   | Configuration to override the default retry behavior of the client.                                                  |

### Response

**[components.DtoEntitlementResponse](../../models/components/dtoentitlementresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_entitlements_id_

Delete an entitlement

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/entitlements/{id}" method="delete" path="/entitlements/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.entitlements.delete_entitlements_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Entitlement ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_plans_id_entitlements

Get all entitlements for a plan

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/plans/{id}/entitlements" method="get" path="/plans/{id}/entitlements" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.entitlements.get_plans_id_entitlements(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Plan ID                                                             |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPlanResponse](../../models/components/dtoplanresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |