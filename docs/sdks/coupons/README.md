# Coupons

## Overview

### Available Operations

* [get_coupons](#get_coupons) - List coupons with filtering
* [post_coupons](#post_coupons) - Create a new coupon
* [get_coupons_id_](#get_coupons_id_) - Get a coupon by ID
* [put_coupons_id_](#put_coupons_id_) - Update a coupon
* [delete_coupons_id_](#delete_coupons_id_) - Delete a coupon

## get_coupons

Lists coupons with filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/coupons" method="get" path="/coupons" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.coupons.get_coupons()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `coupon_ids`                                                                         | List[*str*]                                                                          | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `expand`                                                                             | *Optional[str]*                                                                      | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `limit`                                                                              | *Optional[int]*                                                                      | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `offset`                                                                             | *Optional[int]*                                                                      | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `order`                                                                              | [Optional[operations.GetCouponsOrder]](../../models/operations/getcouponsorder.md)   | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `status`                                                                             | [Optional[operations.GetCouponsStatus]](../../models/operations/getcouponsstatus.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `retries`                                                                            | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                     | :heavy_minus_sign:                                                                   | Configuration to override the default retry behavior of the client.                  |

### Response

**[components.DtoListCouponsResponse](../../models/components/dtolistcouponsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_coupons

Creates a new coupon

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/coupons" method="post" path="/coupons" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.coupons.post_coupons(cadence="forever", name="<value>", type_="percentage")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `cadence`                                                                      | [components.TypesCouponCadence](../../models/components/typescouponcadence.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `name`                                                                         | *str*                                                                          | :heavy_check_mark:                                                             | N/A                                                                            |
| `type`                                                                         | [components.TypesCouponType](../../models/components/typescoupontype.md)       | :heavy_check_mark:                                                             | N/A                                                                            |
| `amount_off`                                                                   | *Optional[str]*                                                                | :heavy_minus_sign:                                                             | N/A                                                                            |
| `currency`                                                                     | *Optional[str]*                                                                | :heavy_minus_sign:                                                             | N/A                                                                            |
| `duration_in_periods`                                                          | *Optional[int]*                                                                | :heavy_minus_sign:                                                             | N/A                                                                            |
| `max_redemptions`                                                              | *Optional[int]*                                                                | :heavy_minus_sign:                                                             | N/A                                                                            |
| `metadata`                                                                     | Dict[str, *str*]                                                               | :heavy_minus_sign:                                                             | N/A                                                                            |
| `percentage_off`                                                               | *Optional[str]*                                                                | :heavy_minus_sign:                                                             | N/A                                                                            |
| `redeem_after`                                                                 | *Optional[str]*                                                                | :heavy_minus_sign:                                                             | N/A                                                                            |
| `redeem_before`                                                                | *Optional[str]*                                                                | :heavy_minus_sign:                                                             | N/A                                                                            |
| `rules`                                                                        | Dict[str, *Any*]                                                               | :heavy_minus_sign:                                                             | N/A                                                                            |
| `retries`                                                                      | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)               | :heavy_minus_sign:                                                             | Configuration to override the default retry behavior of the client.            |

### Response

**[components.DtoCouponResponse](../../models/components/dtocouponresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_coupons_id_

Retrieves a coupon by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/coupons/{id}" method="get" path="/coupons/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.coupons.get_coupons_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Coupon ID                                                           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCouponResponse](../../models/components/dtocouponresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_coupons_id_

Updates an existing coupon

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/coupons/{id}" method="put" path="/coupons/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.coupons.put_coupons_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Coupon ID                                                           |
| `metadata`                                                          | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `name`                                                              | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCouponResponse](../../models/components/dtocouponresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_coupons_id_

Deletes a coupon

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/coupons/{id}" method="delete" path="/coupons/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.coupons.delete_coupons_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Coupon ID                                                           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, str]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |