# Prices

## Overview

### Available Operations

* [get_prices](#get_prices) - Get prices
* [post_prices](#post_prices) - Create a new price
* [post_prices_bulk](#post_prices_bulk) - Create multiple prices in bulk
* [get_prices_id_](#get_prices_id_) - Get a price by ID
* [put_prices_id_](#put_prices_id_) - Update a price
* [delete_prices_id_](#delete_prices_id_) - Delete a price

## get_prices

Get prices with the specified filter

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/prices" method="get" path="/prices" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.prices.get_prices(allow_expired_prices=False)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `allow_expired_prices`                                                                     | *Optional[bool]*                                                                           | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `end_time`                                                                                 | *Optional[str]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `entity_ids`                                                                               | List[*str*]                                                                                | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `entity_type`                                                                              | [Optional[operations.GetPricesEntityType]](../../models/operations/getpricesentitytype.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `expand`                                                                                   | *Optional[str]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `limit`                                                                                    | *Optional[int]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `meter_ids`                                                                                | List[*str*]                                                                                | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `offset`                                                                                   | *Optional[int]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `order`                                                                                    | [Optional[operations.GetPricesOrder]](../../models/operations/getpricesorder.md)           | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `parent_price_id`                                                                          | *Optional[str]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `plan_ids`                                                                                 | List[*str*]                                                                                | :heavy_minus_sign:                                                                         | Price override filtering fields                                                            |
| `price_ids`                                                                                | List[*str*]                                                                                | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `sort`                                                                                     | *Optional[str]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `start_date_lt`                                                                            | *Optional[str]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `start_time`                                                                               | *Optional[str]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `status`                                                                                   | [Optional[operations.GetPricesStatus]](../../models/operations/getpricesstatus.md)         | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `subscription_id`                                                                          | *Optional[str]*                                                                            | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `retries`                                                                                  | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                           | :heavy_minus_sign:                                                                         | Configuration to override the default retry behavior of the client.                        |

### Response

**[components.DtoListPricesResponse](../../models/components/dtolistpricesresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_prices

Create a new price with the specified configuration. Supports both regular and price unit configurations.

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/prices" method="post" path="/prices" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.prices.post_prices(billing_cadence="RECURRING", billing_model="TIERED", billing_period="DAILY", currency="Iceland Krona", entity_id="<id>", entity_type="COSTSHEET", invoice_cadence="ADVANCE", price_unit_type="CUSTOM", type_="USAGE")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `billing_cadence`                                                                                | [components.TypesBillingCadence](../../models/components/typesbillingcadence.md)                 | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `billing_model`                                                                                  | [components.TypesBillingModel](../../models/components/typesbillingmodel.md)                     | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `billing_period`                                                                                 | [components.TypesBillingPeriod](../../models/components/typesbillingperiod.md)                   | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `currency`                                                                                       | *str*                                                                                            | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `entity_id`                                                                                      | *str*                                                                                            | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `entity_type`                                                                                    | [components.TypesPriceEntityType](../../models/components/typespriceentitytype.md)               | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `invoice_cadence`                                                                                | [components.TypesInvoiceCadence](../../models/components/typesinvoicecadence.md)                 | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `price_unit_type`                                                                                | [components.TypesPriceUnitType](../../models/components/typespriceunittype.md)                   | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `type`                                                                                           | [components.TypesPriceType](../../models/components/typespricetype.md)                           | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `amount`                                                                                         | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `billing_period_count`                                                                           | *Optional[int]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `description`                                                                                    | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `display_name`                                                                                   | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `end_date`                                                                                       | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `filter_values`                                                                                  | Dict[str, List[*str*]]                                                                           | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `group_id`                                                                                       | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | GroupID is the id of the group to add the price to                                               |
| `lookup_key`                                                                                     | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `metadata`                                                                                       | Dict[str, *str*]                                                                                 | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `meter_id`                                                                                       | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `min_quantity`                                                                                   | *Optional[int]*                                                                                  | :heavy_minus_sign:                                                                               | MinQuantity is the minimum quantity of the price                                                 |
| `price_unit_config`                                                                              | [Optional[components.DtoPriceUnitConfig]](../../models/components/dtopriceunitconfig.md)         | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `start_date`                                                                                     | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `tier_mode`                                                                                      | [Optional[components.TypesBillingTier]](../../models/components/typesbillingtier.md)             | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `tiers`                                                                                          | List[[components.DtoCreatePriceTier](../../models/components/dtocreatepricetier.md)]             | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `transform_quantity`                                                                             | [Optional[components.PriceTransformQuantity]](../../models/components/pricetransformquantity.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `trial_period`                                                                                   | *Optional[int]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `retries`                                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                 | :heavy_minus_sign:                                                                               | Configuration to override the default retry behavior of the client.                              |

### Response

**[components.DtoPriceResponse](../../models/components/dtopriceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_prices_bulk

Create multiple prices with the specified configurations. Supports both regular and price unit configurations.

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/prices/bulk" method="post" path="/prices/bulk" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.prices.post_prices_bulk(items=[])

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `items`                                                                                    | List[[components.DtoCreatePriceRequest](../../models/components/dtocreatepricerequest.md)] | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `retries`                                                                                  | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                           | :heavy_minus_sign:                                                                         | Configuration to override the default retry behavior of the client.                        |

### Response

**[components.DtoCreateBulkPriceResponse](../../models/components/dtocreatebulkpriceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_prices_id_

Get a price by ID with expanded meter and price unit information

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/prices/{id}" method="get" path="/prices/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.prices.get_prices_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Price ID                                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPriceResponse](../../models/components/dtopriceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_prices_id_

Update a price with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/prices/{id}" method="put" path="/prices/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.prices.put_prices_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `id`                                                                                             | *str*                                                                                            | :heavy_check_mark:                                                                               | Price ID                                                                                         |
| `amount`                                                                                         | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | Amount is the new price amount that overrides the original price (optional)                      |
| `billing_model`                                                                                  | [Optional[components.TypesBillingModel]](../../models/components/typesbillingmodel.md)           | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `description`                                                                                    | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `effective_from`                                                                                 | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `group_id`                                                                                       | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | GroupID is the id of the group to update the price in                                            |
| `lookup_key`                                                                                     | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | All price fields that can be updated<br/>Non-critical fields (can be updated directly)           |
| `metadata`                                                                                       | Dict[str, *str*]                                                                                 | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `tier_mode`                                                                                      | [Optional[components.TypesBillingTier]](../../models/components/typesbillingtier.md)             | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `tiers`                                                                                          | List[[components.DtoCreatePriceTier](../../models/components/dtocreatepricetier.md)]             | :heavy_minus_sign:                                                                               | Tiers determines the pricing tiers for this line item                                            |
| `transform_quantity`                                                                             | [Optional[components.PriceTransformQuantity]](../../models/components/pricetransformquantity.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `retries`                                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                 | :heavy_minus_sign:                                                                               | Configuration to override the default retry behavior of the client.                              |

### Response

**[components.DtoPriceResponse](../../models/components/dtopriceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_prices_id_

Delete a price

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/prices/{id}" method="delete" path="/prices/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.prices.delete_prices_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Price ID                                                            |
| `end_date`                                                          | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |