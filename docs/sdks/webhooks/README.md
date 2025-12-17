# Webhooks

## Overview

### Available Operations

* [process_chargebee](#process_chargebee) - Handle Chargebee webhook events
* [handle_hubspot](#handle_hubspot) - Handle HubSpot webhook events
* [process_nomod](#process_nomod) - Handle Nomod webhook events
* [handle_quickbooks](#handle_quickbooks) - Handle QuickBooks webhook events
* [handle_razorpay](#handle_razorpay) - Handle Razorpay webhook events
* [process_stripe](#process_stripe) - Handle Stripe webhook events

## process_chargebee

Process incoming Chargebee webhook events for payment status updates

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/chargebee/{tenant_id}/{environment_id}" method="post" path="/webhooks/chargebee/{tenant_id}/{environment_id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.webhooks.process_chargebee(tenant_id="<id>", environment_id="<id>")

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
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## handle_hubspot

Process incoming HubSpot webhook events for deal closed won and customer creation

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/hubspot/{tenant_id}/{environment_id}" method="post" path="/webhooks/hubspot/{tenant_id}/{environment_id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.webhooks.handle_hubspot(tenant_id="<id>", environment_id="<id>", x_hub_spot_signature_v3="<value>")

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
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## process_nomod

Process incoming Nomod webhook events for payment and invoice payments

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/nomod/{tenant_id}/{environment_id}" method="post" path="/webhooks/nomod/{tenant_id}/{environment_id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.webhooks.process_nomod(tenant_id="<id>", environment_id="<id>")

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
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## handle_quickbooks

Process incoming QuickBooks webhook events for payment sync

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/quickbooks/{tenant_id}/{environment_id}" method="post" path="/webhooks/quickbooks/{tenant_id}/{environment_id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.webhooks.handle_quickbooks(tenant_id="<id>", environment_id="<id>")

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
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## handle_razorpay

Process incoming Razorpay webhook events for payment capture and failure

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/razorpay/{tenant_id}/{environment_id}" method="post" path="/webhooks/razorpay/{tenant_id}/{environment_id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.webhooks.handle_razorpay(tenant_id="<id>", environment_id="<id>", x_razorpay_signature="<value>")

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
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## process_stripe

Process incoming Stripe webhook events for payment status updates and customer creation

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/webhooks/stripe/{tenant_id}/{environment_id}" method="post" path="/webhooks/stripe/{tenant_id}/{environment_id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.webhooks.process_stripe(tenant_id="<id>", environment_id="<id>", stripe_signature="<value>")

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
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |