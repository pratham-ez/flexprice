# Prices

## Overview

### Available Operations

* [list](#list) - Get prices
* [create](#create) - Create a new price
* [bulk_create](#bulk_create) - Create multiple prices in bulk
* [get_by_id](#get_by_id) - Get a price by ID
* [update](#update) - Update a price
* [delete](#delete) - Delete a price

## list

Get prices with the specified filter

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/prices" method="get" path="/prices" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.prices.list(request={})

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `request`                                                           | [models.GetPricesRequest](../../models/getpricesrequest.md)         | :heavy_check_mark:                                                  | The request object to use for the request.                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoListPricesResponse](../../models/dtolistpricesresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create

Create a new price with the specified configuration. Supports both regular and price unit configurations.

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/prices" method="post" path="/prices" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.prices.create(billing_cadence="RECURRING", billing_model="TIERED", billing_period="DAILY", currency="Iceland Krona", entity_id="<id>", entity_type="COSTSHEET", invoice_cadence="ADVANCE", price_unit_type="CUSTOM", type_="USAGE")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `billing_cadence`                                                                 | [models.TypesBillingCadence](../../models/typesbillingcadence.md)                 | :heavy_check_mark:                                                                | N/A                                                                               |
| `billing_model`                                                                   | [models.TypesBillingModel](../../models/typesbillingmodel.md)                     | :heavy_check_mark:                                                                | N/A                                                                               |
| `billing_period`                                                                  | [models.TypesBillingPeriod](../../models/typesbillingperiod.md)                   | :heavy_check_mark:                                                                | N/A                                                                               |
| `currency`                                                                        | *str*                                                                             | :heavy_check_mark:                                                                | N/A                                                                               |
| `entity_id`                                                                       | *str*                                                                             | :heavy_check_mark:                                                                | N/A                                                                               |
| `entity_type`                                                                     | [models.TypesPriceEntityType](../../models/typespriceentitytype.md)               | :heavy_check_mark:                                                                | N/A                                                                               |
| `invoice_cadence`                                                                 | [models.TypesInvoiceCadence](../../models/typesinvoicecadence.md)                 | :heavy_check_mark:                                                                | N/A                                                                               |
| `price_unit_type`                                                                 | [models.TypesPriceUnitType](../../models/typespriceunittype.md)                   | :heavy_check_mark:                                                                | N/A                                                                               |
| `type`                                                                            | [models.TypesPriceType](../../models/typespricetype.md)                           | :heavy_check_mark:                                                                | N/A                                                                               |
| `amount`                                                                          | *Optional[str]*                                                                   | :heavy_minus_sign:                                                                | N/A                                                                               |
| `billing_period_count`                                                            | *Optional[int]*                                                                   | :heavy_minus_sign:                                                                | N/A                                                                               |
| `description`                                                                     | *Optional[str]*                                                                   | :heavy_minus_sign:                                                                | N/A                                                                               |
| `display_name`                                                                    | *Optional[str]*                                                                   | :heavy_minus_sign:                                                                | N/A                                                                               |
| `end_date`                                                                        | *Optional[str]*                                                                   | :heavy_minus_sign:                                                                | N/A                                                                               |
| `filter_values`                                                                   | Dict[str, List[*str*]]                                                            | :heavy_minus_sign:                                                                | N/A                                                                               |
| `group_id`                                                                        | *Optional[str]*                                                                   | :heavy_minus_sign:                                                                | GroupID is the id of the group to add the price to                                |
| `lookup_key`                                                                      | *Optional[str]*                                                                   | :heavy_minus_sign:                                                                | N/A                                                                               |
| `metadata`                                                                        | Dict[str, *str*]                                                                  | :heavy_minus_sign:                                                                | N/A                                                                               |
| `meter_id`                                                                        | *Optional[str]*                                                                   | :heavy_minus_sign:                                                                | N/A                                                                               |
| `min_quantity`                                                                    | *Optional[int]*                                                                   | :heavy_minus_sign:                                                                | MinQuantity is the minimum quantity of the price                                  |
| `price_unit_config`                                                               | [Optional[models.DtoPriceUnitConfig]](../../models/dtopriceunitconfig.md)         | :heavy_minus_sign:                                                                | N/A                                                                               |
| `start_date`                                                                      | *Optional[str]*                                                                   | :heavy_minus_sign:                                                                | N/A                                                                               |
| `tier_mode`                                                                       | [Optional[models.TypesBillingTier]](../../models/typesbillingtier.md)             | :heavy_minus_sign:                                                                | N/A                                                                               |
| `tiers`                                                                           | List[[models.DtoCreatePriceTier](../../models/dtocreatepricetier.md)]             | :heavy_minus_sign:                                                                | N/A                                                                               |
| `transform_quantity`                                                              | [Optional[models.PriceTransformQuantity]](../../models/pricetransformquantity.md) | :heavy_minus_sign:                                                                | N/A                                                                               |
| `trial_period`                                                                    | *Optional[int]*                                                                   | :heavy_minus_sign:                                                                | N/A                                                                               |
| `retries`                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                  | :heavy_minus_sign:                                                                | Configuration to override the default retry behavior of the client.               |

### Response

**[models.DtoPriceResponse](../../models/dtopriceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## bulk_create

Create multiple prices with the specified configurations. Supports both regular and price unit configurations.

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/prices/bulk" method="post" path="/prices/bulk" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.prices.bulk_create(items=[])

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `items`                                                                     | List[[models.DtoCreatePriceRequest](../../models/dtocreatepricerequest.md)] | :heavy_check_mark:                                                          | N/A                                                                         |
| `retries`                                                                   | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)            | :heavy_minus_sign:                                                          | Configuration to override the default retry behavior of the client.         |

### Response

**[models.DtoCreateBulkPriceResponse](../../models/dtocreatebulkpriceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_by_id

Get a price by ID with expanded meter and price unit information

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/prices/{id}" method="get" path="/prices/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.prices.get_by_id(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Price ID                                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoPriceResponse](../../models/dtopriceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update

Update a price with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/prices/{id}" method="put" path="/prices/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.prices.update(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `id`                                                                               | *str*                                                                              | :heavy_check_mark:                                                                 | Price ID                                                                           |
| `amount`                                                                           | *Optional[str]*                                                                    | :heavy_minus_sign:                                                                 | Amount is the new price amount that overrides the original price (optional)        |
| `billing_model`                                                                    | [Optional[models.TypesBillingModel]](../../models/typesbillingmodel.md)            | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `description`                                                                      | *Optional[str]*                                                                    | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `effective_from`                                                                   | *Optional[str]*                                                                    | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `group_id`                                                                         | *Optional[str]*                                                                    | :heavy_minus_sign:                                                                 | GroupID is the id of the group to update the price in                              |
| `lookup_key`                                                                       | *Optional[str]*                                                                    | :heavy_minus_sign:                                                                 | All price fields that can be updated<br/>Non-critical fields (can be updated directly) |
| `metadata`                                                                         | Dict[str, *str*]                                                                   | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `tier_mode`                                                                        | [Optional[models.TypesBillingTier]](../../models/typesbillingtier.md)              | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `tiers`                                                                            | List[[models.DtoCreatePriceTier](../../models/dtocreatepricetier.md)]              | :heavy_minus_sign:                                                                 | Tiers determines the pricing tiers for this line item                              |
| `transform_quantity`                                                               | [Optional[models.PriceTransformQuantity]](../../models/pricetransformquantity.md)  | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `retries`                                                                          | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                   | :heavy_minus_sign:                                                                 | Configuration to override the default retry behavior of the client.                |

### Response

**[models.DtoPriceResponse](../../models/dtopriceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete

Delete a price

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/prices/{id}" method="delete" path="/prices/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.prices.delete(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Price ID                                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, models.GinH]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |