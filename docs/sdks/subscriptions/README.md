# Subscriptions

## Overview

### Available Operations

* [get_subscriptions](#get_subscriptions) - List subscriptions
* [post_subscriptions](#post_subscriptions) - Create subscription
* [post_subscriptions_addon](#post_subscriptions_addon) - Add addon to subscription
* [delete_subscriptions_addon](#delete_subscriptions_addon) - Remove addon from subscription
* [put_subscriptions_lineitems_id_](#put_subscriptions_lineitems_id_) - Update subscription line item
* [delete_subscriptions_lineitems_id_](#delete_subscriptions_lineitems_id_) - Delete subscription line item
* [post_subscriptions_search](#post_subscriptions_search) - List subscriptions by filter
* [post_subscriptions_usage](#post_subscriptions_usage) - Get usage by subscription
* [get_subscriptions_id_](#get_subscriptions_id_) - Get subscription
* [post_subscriptions_id_activate](#post_subscriptions_id_activate) - Activate draft subscription
* [get_subscriptions_id_addons_associations](#get_subscriptions_id_addons_associations) - Get active addon associations
* [post_subscriptions_id_cancel](#post_subscriptions_id_cancel) - Cancel subscription
* [post_subscriptions_id_change_execute](#post_subscriptions_id_change_execute) - Execute subscription plan change
* [post_subscriptions_id_change_preview](#post_subscriptions_id_change_preview) - Preview subscription plan change
* [get_subscriptions_id_entitlements](#get_subscriptions_id_entitlements) - Get subscription entitlements
* [get_subscriptions_id_grants_upcoming](#get_subscriptions_id_grants_upcoming) - Get upcoming credit grant applications
* [post_subscriptions_id_pause](#post_subscriptions_id_pause) - Pause a subscription
* [get_subscriptions_id_pauses](#get_subscriptions_id_pauses) - List all pauses for a subscription
* [post_subscriptions_id_resume](#post_subscriptions_id_resume) - Resume a paused subscription

## get_subscriptions

Get subscriptions with optional filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/subscriptions" method="get" path="/subscriptions" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.get_subscriptions()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `active_at`                                                                                      | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | ActiveAt filters subscriptions that are active at the given time                                 |
| `billing_cadence`                                                                                | List[[operations.BillingCadence](../../models/operations/billingcadence.md)]                     | :heavy_minus_sign:                                                                               | BillingCadence filters by billing cadence                                                        |
| `billing_period`                                                                                 | List[[operations.BillingPeriod](../../models/operations/billingperiod.md)]                       | :heavy_minus_sign:                                                                               | BillingPeriod filters by billing period                                                          |
| `customer_id`                                                                                    | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | CustomerID filters by customer ID                                                                |
| `end_time`                                                                                       | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `expand`                                                                                         | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `external_customer_id`                                                                           | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | ExternalCustomerID filters by external customer ID                                               |
| `invoicing_customer_ids`                                                                         | List[*str*]                                                                                      | :heavy_minus_sign:                                                                               | InvoicingCustomerIDs filters by invoicing customer ID                                            |
| `limit`                                                                                          | *Optional[int]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `offset`                                                                                         | *Optional[int]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `order`                                                                                          | [Optional[operations.GetSubscriptionsOrder]](../../models/operations/getsubscriptionsorder.md)   | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `plan_id`                                                                                        | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | PlanID filters by plan ID                                                                        |
| `start_time`                                                                                     | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `status`                                                                                         | [Optional[operations.GetSubscriptionsStatus]](../../models/operations/getsubscriptionsstatus.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `subscription_ids`                                                                               | List[*str*]                                                                                      | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `subscription_status`                                                                            | List[[operations.SubscriptionStatus](../../models/operations/subscriptionstatus.md)]             | :heavy_minus_sign:                                                                               | SubscriptionStatus filters by subscription status                                                |
| `with_line_items`                                                                                | *Optional[bool]*                                                                                 | :heavy_minus_sign:                                                                               | WithLineItems includes line items in the response                                                |
| `retries`                                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                 | :heavy_minus_sign:                                                                               | Configuration to override the default retry behavior of the client.                              |

### Response

**[components.DtoListSubscriptionsResponse](../../models/components/dtolistsubscriptionsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_subscriptions

Create a new subscription

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/subscriptions" method="post" path="/subscriptions" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.post_subscriptions(billing_cadence="RECURRING", billing_period="HALF_YEARLY", currency="Nakfa", plan_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                         | Type                                                                                                                                              | Required                                                                                                                                          | Description                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| `billing_cadence`                                                                                                                                 | [components.TypesBillingCadence](../../models/components/typesbillingcadence.md)                                                                  | :heavy_check_mark:                                                                                                                                | N/A                                                                                                                                               |
| `billing_period`                                                                                                                                  | [components.TypesBillingPeriod](../../models/components/typesbillingperiod.md)                                                                    | :heavy_check_mark:                                                                                                                                | N/A                                                                                                                                               |
| `currency`                                                                                                                                        | *str*                                                                                                                                             | :heavy_check_mark:                                                                                                                                | N/A                                                                                                                                               |
| `plan_id`                                                                                                                                         | *str*                                                                                                                                             | :heavy_check_mark:                                                                                                                                | N/A                                                                                                                                               |
| `addons`                                                                                                                                          | List[[components.DtoAddAddonToSubscriptionRequest](../../models/components/dtoaddaddontosubscriptionrequest.md)]                                  | :heavy_minus_sign:                                                                                                                                | Addons represents addons to be added to the subscription during creation                                                                          |
| `billing_cycle`                                                                                                                                   | [Optional[components.TypesBillingCycle]](../../models/components/typesbillingcycle.md)                                                            | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `billing_period_count`                                                                                                                            | *Optional[int]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `collection_method`                                                                                                                               | [Optional[components.TypesCollectionMethod]](../../models/components/typescollectionmethod.md)                                                    | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `commitment_amount`                                                                                                                               | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | CommitmentAmount is the minimum amount a customer commits to paying for a billing period                                                          |
| `coupons`                                                                                                                                         | List[*str*]                                                                                                                                       | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `credit_grants`                                                                                                                                   | List[[components.DtoCreateCreditGrantRequest](../../models/components/dtocreatecreditgrantrequest.md)]                                            | :heavy_minus_sign:                                                                                                                                | Credit grants to be applied when subscription is created                                                                                          |
| `customer_id`                                                                                                                                     | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | customer_id is the flexprice customer id<br/>and it is prioritized over external_customer_id in case both are provided.                           |
| `customer_timezone`                                                                                                                               | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | Timezone of the customer.<br/>If not set, the default value is UTC.                                                                               |
| `enable_true_up`                                                                                                                                  | *Optional[bool]*                                                                                                                                  | :heavy_minus_sign:                                                                                                                                | Enable Commitment True Up Fee                                                                                                                     |
| `end_date`                                                                                                                                        | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `external_customer_id`                                                                                                                            | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | external_customer_id is the customer id in your DB<br/>and must be same as what you provided as external_id while creating the customer in flexprice. |
| `gateway_payment_method_id`                                                                                                                       | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `invoice_billing`                                                                                                                                 | [Optional[components.TypesInvoiceBilling]](../../models/components/typesinvoicebilling.md)                                                        | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `line_item_commitments`                                                                                                                           | Dict[str, [components.DtoLineItemCommitmentConfig](../../models/components/dtolineitemcommitmentconfig.md)]                                       | :heavy_minus_sign:                                                                                                                                | LineItemCommitments allows setting commitment configuration per line item (keyed by price_id)                                                     |
| `line_item_coupons`                                                                                                                               | Dict[str, List[*str*]]                                                                                                                            | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `lookup_key`                                                                                                                                      | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `metadata`                                                                                                                                        | Dict[str, *str*]                                                                                                                                  | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `overage_factor`                                                                                                                                  | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | OverageFactor is a multiplier applied to usage beyond the commitment amount                                                                       |
| `override_entitlements`                                                                                                                           | List[[components.DtoOverrideEntitlementRequest](../../models/components/dtooverrideentitlementrequest.md)]                                        | :heavy_minus_sign:                                                                                                                                | OverrideEntitlements allows customizing specific entitlements for this subscription                                                               |
| `override_line_items`                                                                                                                             | List[[components.DtoOverrideLineItemRequest](../../models/components/dtooverridelineitemrequest.md)]                                              | :heavy_minus_sign:                                                                                                                                | OverrideLineItems allows customizing specific prices for this subscription                                                                        |
| `payment_behavior`                                                                                                                                | [Optional[components.TypesPaymentBehavior]](../../models/components/typespaymentbehavior.md)                                                      | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `phases`                                                                                                                                          | List[[components.DtoSubscriptionPhaseCreateRequest](../../models/components/dtosubscriptionphasecreaterequest.md)]                                | :heavy_minus_sign:                                                                                                                                | Phases represents subscription phases to be created with the subscription                                                                         |
| `proration_behavior`                                                                                                                              | [Optional[components.TypesProrationBehavior]](../../models/components/typesprorationbehavior.md)                                                  | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `start_date`                                                                                                                                      | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `subscription_status`                                                                                                                             | [Optional[components.TypesSubscriptionStatus]](../../models/components/typessubscriptionstatus.md)                                                | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `tax_rate_overrides`                                                                                                                              | List[[components.DtoTaxRateOverride](../../models/components/dtotaxrateoverride.md)]                                                              | :heavy_minus_sign:                                                                                                                                | tax_rate_overrides is the tax rate overrides	to be applied to the subscription                                                                    |
| `trial_end`                                                                                                                                       | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `trial_start`                                                                                                                                     | *Optional[str]*                                                                                                                                   | :heavy_minus_sign:                                                                                                                                | N/A                                                                                                                                               |
| `retries`                                                                                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                  | :heavy_minus_sign:                                                                                                                                | Configuration to override the default retry behavior of the client.                                                                               |

### Response

**[components.DtoSubscriptionResponse](../../models/components/dtosubscriptionresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_subscriptions_addon

Add an addon to a subscription

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/subscriptions/addon" method="post" path="/subscriptions/addon" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.post_subscriptions_addon(addon_id="<id>", subscription_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `addon_id`                                                          | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `subscription_id`                                                   | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `metadata`                                                          | Dict[str, *Any*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `start_date`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoAddonAssociationResponse](../../models/components/dtoaddonassociationresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_subscriptions_addon

Remove an addon from a subscription

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/subscriptions/addon" method="delete" path="/subscriptions/addon" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.delete_subscriptions_addon(addon_association_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `addon_association_id`                                              | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `reason`                                                            | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_subscriptions_lineitems_id_

Update a subscription line item by terminating the existing one and creating a new one

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/subscriptions/lineitems/{id}" method="put" path="/subscriptions/lineitems/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.put_subscriptions_lineitems_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `id`                                                                                             | *str*                                                                                            | :heavy_check_mark:                                                                               | Line Item ID                                                                                     |
| `amount`                                                                                         | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | Amount is the new price amount that overrides the original price                                 |
| `billing_model`                                                                                  | [Optional[components.TypesBillingModel]](../../models/components/typesbillingmodel.md)           | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `commitment_amount`                                                                              | *Optional[float]*                                                                                | :heavy_minus_sign:                                                                               | Commitment fields                                                                                |
| `commitment_overage_factor`                                                                      | *Optional[float]*                                                                                | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `commitment_quantity`                                                                            | *Optional[float]*                                                                                | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `commitment_true_up_enabled`                                                                     | *Optional[bool]*                                                                                 | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `commitment_type`                                                                                | [Optional[components.TypesCommitmentType]](../../models/components/typescommitmenttype.md)       | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `commitment_windowed`                                                                            | *Optional[bool]*                                                                                 | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `effective_from`                                                                                 | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | EffectiveFrom for the existing line item (if not provided, defaults to now)                      |
| `metadata`                                                                                       | Dict[str, *str*]                                                                                 | :heavy_minus_sign:                                                                               | Metadata for the new line item                                                                   |
| `tier_mode`                                                                                      | [Optional[components.TypesBillingTier]](../../models/components/typesbillingtier.md)             | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `tiers`                                                                                          | List[[components.DtoCreatePriceTier](../../models/components/dtocreatepricetier.md)]             | :heavy_minus_sign:                                                                               | Tiers determines the pricing tiers for this line item                                            |
| `transform_quantity`                                                                             | [Optional[components.PriceTransformQuantity]](../../models/components/pricetransformquantity.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `retries`                                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                 | :heavy_minus_sign:                                                                               | Configuration to override the default retry behavior of the client.                              |

### Response

**[components.DtoSubscriptionLineItemResponse](../../models/components/dtosubscriptionlineitemresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_subscriptions_lineitems_id_

Delete a subscription line item by setting its end date

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/subscriptions/lineitems/{id}" method="delete" path="/subscriptions/lineitems/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.delete_subscriptions_lineitems_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Line Item ID                                                        |
| `effective_from`                                                    | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSubscriptionLineItemResponse](../../models/components/dtosubscriptionlineitemresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_subscriptions_search

List subscriptions by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/subscriptions/search" method="post" path="/subscriptions/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.post_subscriptions_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `active_at`                                                                                                  | *Optional[str]*                                                                                              | :heavy_minus_sign:                                                                                           | ActiveAt filters subscriptions that are active at the given time                                             |
| `billing_cadence`                                                                                            | List[[components.TypesBillingCadence](../../models/components/typesbillingcadence.md)]                       | :heavy_minus_sign:                                                                                           | BillingCadence filters by billing cadence                                                                    |
| `billing_period`                                                                                             | List[[components.TypesBillingPeriod](../../models/components/typesbillingperiod.md)]                         | :heavy_minus_sign:                                                                                           | BillingPeriod filters by billing period                                                                      |
| `customer_id`                                                                                                | *Optional[str]*                                                                                              | :heavy_minus_sign:                                                                                           | CustomerID filters by customer ID                                                                            |
| `end_time`                                                                                                   | *Optional[str]*                                                                                              | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `expand`                                                                                                     | *Optional[str]*                                                                                              | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `external_customer_id`                                                                                       | *Optional[str]*                                                                                              | :heavy_minus_sign:                                                                                           | ExternalCustomerID filters by external customer ID                                                           |
| `filters`                                                                                                    | List[[components.TypesFilterCondition](../../models/components/typesfiltercondition.md)]                     | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `invoicing_customer_ids`                                                                                     | List[*str*]                                                                                                  | :heavy_minus_sign:                                                                                           | InvoicingCustomerIDs filters by invoicing customer ID                                                        |
| `limit`                                                                                                      | *Optional[int]*                                                                                              | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `offset`                                                                                                     | *Optional[int]*                                                                                              | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `order`                                                                                                      | [Optional[components.TypesSubscriptionFilterOrder]](../../models/components/typessubscriptionfilterorder.md) | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `plan_id`                                                                                                    | *Optional[str]*                                                                                              | :heavy_minus_sign:                                                                                           | PlanID filters by plan ID                                                                                    |
| `sort`                                                                                                       | List[[components.TypesSortCondition](../../models/components/typessortcondition.md)]                         | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `start_time`                                                                                                 | *Optional[str]*                                                                                              | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `status`                                                                                                     | [Optional[components.TypesStatus]](../../models/components/typesstatus.md)                                   | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `subscription_ids`                                                                                           | List[*str*]                                                                                                  | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `subscription_status`                                                                                        | List[[components.TypesSubscriptionStatus](../../models/components/typessubscriptionstatus.md)]               | :heavy_minus_sign:                                                                                           | SubscriptionStatus filters by subscription status                                                            |
| `with_line_items`                                                                                            | *Optional[bool]*                                                                                             | :heavy_minus_sign:                                                                                           | WithLineItems includes line items in the response                                                            |
| `retries`                                                                                                    | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                             | :heavy_minus_sign:                                                                                           | Configuration to override the default retry behavior of the client.                                          |

### Response

**[components.DtoListSubscriptionsResponse](../../models/components/dtolistsubscriptionsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_subscriptions_usage

Get usage for a subscription

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/subscriptions/usage" method="post" path="/subscriptions/usage" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.post_subscriptions_usage(subscription_id="123", end_time="2024-03-20T00:00:00Z", lifetime_usage=False, start_time="2024-03-13T00:00:00Z")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `subscription_id`                                                   | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 | 123                                                                 |
| `end_time`                                                          | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 | 2024-03-20T00:00:00Z                                                |
| `lifetime_usage`                                                    | *Optional[bool]*                                                    | :heavy_minus_sign:                                                  | N/A                                                                 | false                                                               |
| `start_time`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 | 2024-03-13T00:00:00Z                                                |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |                                                                     |

### Response

**[components.DtoGetUsageBySubscriptionResponse](../../models/components/dtogetusagebysubscriptionresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_subscriptions_id_

Get a subscription by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/subscriptions/{id}" method="get" path="/subscriptions/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.get_subscriptions_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Subscription ID                                                     |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSubscriptionResponse](../../models/components/dtosubscriptionresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_subscriptions_id_activate

Activate a draft subscription with a new start date

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/subscriptions/{id}/activate" method="post" path="/subscriptions/{id}/activate" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.post_subscriptions_id_activate(id="<id>", start_date="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `id`                                                                  | *str*                                                                 | :heavy_check_mark:                                                    | Subscription ID                                                       |
| `start_date`                                                          | *str*                                                                 | :heavy_check_mark:                                                    | start_date is the new start date for the subscription when activating |
| `retries`                                                             | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)      | :heavy_minus_sign:                                                    | Configuration to override the default retry behavior of the client.   |

### Response

**[components.DtoSubscriptionResponse](../../models/components/dtosubscriptionresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_subscriptions_id_addons_associations

Get active addon associations for a subscription

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/subscriptions/{id}/addons/associations" method="get" path="/subscriptions/{id}/addons/associations" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.get_subscriptions_id_addons_associations(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Subscription ID                                                     |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[List[components.DtoAddonAssociationResponse]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_subscriptions_id_cancel

Cancel a subscription with enhanced proration support

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/subscriptions/{id}/cancel" method="post" path="/subscriptions/{id}/cancel" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.post_subscriptions_id_cancel(id="<id>", cancellation_type="end_of_period")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `id`                                                                                             | *str*                                                                                            | :heavy_check_mark:                                                                               | Subscription ID                                                                                  |
| `cancellation_type`                                                                              | [components.TypesCancellationType](../../models/components/typescancellationtype.md)             | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `proration_behavior`                                                                             | [Optional[components.TypesProrationBehavior]](../../models/components/typesprorationbehavior.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `reason`                                                                                         | *Optional[str]*                                                                                  | :heavy_minus_sign:                                                                               | Reason for cancellation (for audit and business intelligence)                                    |
| `retries`                                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                 | :heavy_minus_sign:                                                                               | Configuration to override the default retry behavior of the client.                              |

### Response

**[components.DtoCancelSubscriptionResponse](../../models/components/dtocancelsubscriptionresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_subscriptions_id_change_execute

Execute a subscription plan change, including proration and invoice generation

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/subscriptions/{id}/change/execute" method="post" path="/subscriptions/{id}/change/execute" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.post_subscriptions_id_change_execute(id="<id>", billing_cadence="ONETIME", billing_cycle="anniversary", billing_period="HALF_YEARLY", proration_behavior="none", target_plan_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `id`                                                                                   | *str*                                                                                  | :heavy_check_mark:                                                                     | Subscription ID                                                                        |
| `billing_cadence`                                                                      | [components.TypesBillingCadence](../../models/components/typesbillingcadence.md)       | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `billing_cycle`                                                                        | [components.TypesBillingCycle](../../models/components/typesbillingcycle.md)           | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `billing_period`                                                                       | [components.TypesBillingPeriod](../../models/components/typesbillingperiod.md)         | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `proration_behavior`                                                                   | [components.TypesProrationBehavior](../../models/components/typesprorationbehavior.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `target_plan_id`                                                                       | *str*                                                                                  | :heavy_check_mark:                                                                     | target_plan_id is the ID of the new plan to change to (required)                       |
| `billing_period_count`                                                                 | *Optional[int]*                                                                        | :heavy_minus_sign:                                                                     | billing_period_count is the billing period count for the new subscription              |
| `metadata`                                                                             | Dict[str, *str*]                                                                       | :heavy_minus_sign:                                                                     | metadata contains additional key-value pairs for storing extra information             |
| `retries`                                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                       | :heavy_minus_sign:                                                                     | Configuration to override the default retry behavior of the client.                    |

### Response

**[components.DtoSubscriptionChangeExecuteResponse](../../models/components/dtosubscriptionchangeexecuteresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_subscriptions_id_change_preview

Preview the impact of changing a subscription's plan, including proration calculations

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/subscriptions/{id}/change/preview" method="post" path="/subscriptions/{id}/change/preview" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.post_subscriptions_id_change_preview(id="<id>", billing_cadence="RECURRING", billing_cycle="anniversary", billing_period="WEEKLY", proration_behavior="create_prorations", target_plan_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `id`                                                                                   | *str*                                                                                  | :heavy_check_mark:                                                                     | Subscription ID                                                                        |
| `billing_cadence`                                                                      | [components.TypesBillingCadence](../../models/components/typesbillingcadence.md)       | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `billing_cycle`                                                                        | [components.TypesBillingCycle](../../models/components/typesbillingcycle.md)           | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `billing_period`                                                                       | [components.TypesBillingPeriod](../../models/components/typesbillingperiod.md)         | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `proration_behavior`                                                                   | [components.TypesProrationBehavior](../../models/components/typesprorationbehavior.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `target_plan_id`                                                                       | *str*                                                                                  | :heavy_check_mark:                                                                     | target_plan_id is the ID of the new plan to change to (required)                       |
| `billing_period_count`                                                                 | *Optional[int]*                                                                        | :heavy_minus_sign:                                                                     | billing_period_count is the billing period count for the new subscription              |
| `metadata`                                                                             | Dict[str, *str*]                                                                       | :heavy_minus_sign:                                                                     | metadata contains additional key-value pairs for storing extra information             |
| `retries`                                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                       | :heavy_minus_sign:                                                                     | Configuration to override the default retry behavior of the client.                    |

### Response

**[components.DtoSubscriptionChangePreviewResponse](../../models/components/dtosubscriptionchangepreviewresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_subscriptions_id_entitlements

Get all entitlements for a subscription

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/subscriptions/{id}/entitlements" method="get" path="/subscriptions/{id}/entitlements" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.get_subscriptions_id_entitlements(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Subscription ID                                                     |
| `feature_ids`                                                       | List[*str*]                                                         | :heavy_minus_sign:                                                  | Feature IDs to filter by                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSubscriptionEntitlementsResponse](../../models/components/dtosubscriptionentitlementsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_subscriptions_id_grants_upcoming

Get upcoming credit grant applications for a subscription

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/subscriptions/{id}/grants/upcoming" method="get" path="/subscriptions/{id}/grants/upcoming" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.get_subscriptions_id_grants_upcoming(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Subscription ID                                                     |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoListCreditGrantApplicationsResponse](../../models/components/dtolistcreditgrantapplicationsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_subscriptions_id_pause

Pause a subscription with the specified parameters

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/subscriptions/{id}/pause" method="post" path="/subscriptions/{id}/pause" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.post_subscriptions_id_pause(id="<id>", pause_mode="period_end")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                                         | Type                                                                                                                                                                                              | Required                                                                                                                                                                                          | Description                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `id`                                                                                                                                                                                              | *str*                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                | Subscription ID                                                                                                                                                                                   |
| `pause_mode`                                                                                                                                                                                      | [components.TypesPauseMode](../../models/components/typespausemode.md)                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                | N/A                                                                                                                                                                                               |
| `dry_run`                                                                                                                                                                                         | *Optional[bool]*                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                | Whether to perform a dry run<br/>@Description If true, validates the request and shows impact without actually pausing the subscription<br/>@Example false                                        |
| `metadata`                                                                                                                                                                                        | Dict[str, *str*]                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                | Additional metadata as key-value pairs<br/>@Description Optional metadata for storing additional information about the pause<br/>@Example {"requested_by": "customer", "channel": "support_ticket"} |
| `pause_days`                                                                                                                                                                                      | *Optional[int]*                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                | Duration of the pause in days<br/>@Description Number of days to pause the subscription. Cannot be used together with pause_end. Must be greater than 0<br/>@Example 30                           |
| `pause_end`                                                                                                                                                                                       | *Optional[str]*                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                | End date for the subscription pause<br/>@Description ISO 8601 timestamp when the pause should end. Cannot be used together with pause_days. Must be after pause_start<br/>@Example "2024-02-15T00:00:00Z" |
| `pause_start`                                                                                                                                                                                     | *Optional[str]*                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                | Start date for the subscription pause<br/>@Description ISO 8601 timestamp when the pause should begin. Required when pause_mode is "scheduled"<br/>@Example "2024-01-15T00:00:00Z"                |
| `reason`                                                                                                                                                                                          | *Optional[str]*                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                | Reason for pausing the subscription<br/>@Description Optional reason for the pause. Maximum 255 characters<br/>@Example "Customer requested temporary suspension"                                 |
| `retries`                                                                                                                                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                | Configuration to override the default retry behavior of the client.                                                                                                                               |

### Response

**[components.DtoSubscriptionPauseResponse](../../models/components/dtosubscriptionpauseresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_subscriptions_id_pauses

List all pauses for a subscription

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/subscriptions/{id}/pauses" method="get" path="/subscriptions/{id}/pauses" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.get_subscriptions_id_pauses(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Subscription ID                                                     |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[List[components.DtoListSubscriptionPausesResponse]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_subscriptions_id_resume

Resume a paused subscription with the specified parameters

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/subscriptions/{id}/resume" method="post" path="/subscriptions/{id}/resume" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.subscriptions.post_subscriptions_id_resume(id="<id>", resume_mode="scheduled")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                                        | Type                                                                                                                                                                                             | Required                                                                                                                                                                                         | Description                                                                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                                             | *str*                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                               | Subscription ID                                                                                                                                                                                  |
| `resume_mode`                                                                                                                                                                                    | [components.TypesResumeMode](../../models/components/typesresumemode.md)                                                                                                                         | :heavy_check_mark:                                                                                                                                                                               | N/A                                                                                                                                                                                              |
| `dry_run`                                                                                                                                                                                        | *Optional[bool]*                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                               | Whether to perform a dry run<br/>@Description If true, validates the request and shows impact without actually resuming the subscription<br/>@Example false                                      |
| `metadata`                                                                                                                                                                                       | Dict[str, *str*]                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                               | Additional metadata as key-value pairs<br/>@Description Optional metadata for storing additional information about the resume operation<br/>@Example {"resumed_by": "admin", "reason": "issue_resolved"} |
| `retries`                                                                                                                                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                               | Configuration to override the default retry behavior of the client.                                                                                                                              |

### Response

**[components.DtoSubscriptionPauseResponse](../../models/components/dtosubscriptionpauseresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |