# Webhooks

## Overview

### Available Operations

* [post_webhooks_chargebee_tenant_id_environment_id_](#post_webhooks_chargebee_tenant_id_environment_id_) - Handle Chargebee webhook events
* [post_webhooks_hubspot_tenant_id_environment_id_](#post_webhooks_hubspot_tenant_id_environment_id_) - Handle HubSpot webhook events
* [post_webhooks_nomod_tenant_id_environment_id_](#post_webhooks_nomod_tenant_id_environment_id_) - Handle Nomod webhook events
* [post_webhooks_quickbooks_tenant_id_environment_id_](#post_webhooks_quickbooks_tenant_id_environment_id_) - Handle QuickBooks webhook events
* [post_webhooks_razorpay_tenant_id_environment_id_](#post_webhooks_razorpay_tenant_id_environment_id_) - Handle Razorpay webhook events
* [post_webhooks_stripe_tenant_id_environment_id_](#post_webhooks_stripe_tenant_id_environment_id_) - Handle Stripe webhook events

## post_webhooks_chargebee_tenant_id_environment_id_

Process incoming Chargebee webhook events for payment status updates

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/chargebee/{tenant_id}/{environment_id}" method="post" path="/webhooks/chargebee/{tenant_id}/{environment_id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.webhooks.post_webhooks_chargebee_tenant_id_environment_id_(tenant_id="<id>", environment_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `tenant_id`                                                         | *str*                                                               | :heavy_check_mark:                                                  | Tenant ID                                                           |
| `environment_id`                                                    | *str*                                                               | :heavy_check_mark:                                                  | Environment ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, Any]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_webhooks_hubspot_tenant_id_environment_id_

Process incoming HubSpot webhook events for deal closed won and customer creation

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/hubspot/{tenant_id}/{environment_id}" method="post" path="/webhooks/hubspot/{tenant_id}/{environment_id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.webhooks.post_webhooks_hubspot_tenant_id_environment_id_(tenant_id="<id>", environment_id="<id>", x_hub_spot_signature_v3="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `tenant_id`                                                         | *str*                                                               | :heavy_check_mark:                                                  | Tenant ID                                                           |
| `environment_id`                                                    | *str*                                                               | :heavy_check_mark:                                                  | Environment ID                                                      |
| `x_hub_spot_signature_v3`                                           | *str*                                                               | :heavy_check_mark:                                                  | HubSpot webhook signature                                           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, Any]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_webhooks_nomod_tenant_id_environment_id_

Process incoming Nomod webhook events for payment and invoice payments

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/nomod/{tenant_id}/{environment_id}" method="post" path="/webhooks/nomod/{tenant_id}/{environment_id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.webhooks.post_webhooks_nomod_tenant_id_environment_id_(tenant_id="<id>", environment_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `tenant_id`                                                         | *str*                                                               | :heavy_check_mark:                                                  | Tenant ID                                                           |
| `environment_id`                                                    | *str*                                                               | :heavy_check_mark:                                                  | Environment ID                                                      |
| `x_api_key`                                                         | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Nomod webhook secret (if configured)                                |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, Any]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_webhooks_quickbooks_tenant_id_environment_id_

Process incoming QuickBooks webhook events for payment sync

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/quickbooks/{tenant_id}/{environment_id}" method="post" path="/webhooks/quickbooks/{tenant_id}/{environment_id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.webhooks.post_webhooks_quickbooks_tenant_id_environment_id_(tenant_id="<id>", environment_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `tenant_id`                                                         | *str*                                                               | :heavy_check_mark:                                                  | Tenant ID                                                           |
| `environment_id`                                                    | *str*                                                               | :heavy_check_mark:                                                  | Environment ID                                                      |
| `intuit_signature`                                                  | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | QuickBooks webhook signature                                        |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, Any]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_webhooks_razorpay_tenant_id_environment_id_

Process incoming Razorpay webhook events for payment capture and failure

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/razorpay/{tenant_id}/{environment_id}" method="post" path="/webhooks/razorpay/{tenant_id}/{environment_id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.webhooks.post_webhooks_razorpay_tenant_id_environment_id_(tenant_id="<id>", environment_id="<id>", x_razorpay_signature="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `tenant_id`                                                         | *str*                                                               | :heavy_check_mark:                                                  | Tenant ID                                                           |
| `environment_id`                                                    | *str*                                                               | :heavy_check_mark:                                                  | Environment ID                                                      |
| `x_razorpay_signature`                                              | *str*                                                               | :heavy_check_mark:                                                  | Razorpay webhook signature                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, Any]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_webhooks_stripe_tenant_id_environment_id_

Process incoming Stripe webhook events for payment status updates and customer creation

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/stripe/{tenant_id}/{environment_id}" method="post" path="/webhooks/stripe/{tenant_id}/{environment_id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.webhooks.post_webhooks_stripe_tenant_id_environment_id_(tenant_id="<id>", environment_id="<id>", stripe_signature="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `tenant_id`                                                         | *str*                                                               | :heavy_check_mark:                                                  | Tenant ID                                                           |
| `environment_id`                                                    | *str*                                                               | :heavy_check_mark:                                                  | Environment ID                                                      |
| `stripe_signature`                                                  | *str*                                                               | :heavy_check_mark:                                                  | Stripe webhook signature                                            |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, Any]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |