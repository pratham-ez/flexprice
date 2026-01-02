# Payments

## Overview

### Available Operations

* [get_payments](#get_payments) - List payments
* [post_payments](#post_payments) - Create a new payment
* [get_payments_id_](#get_payments_id_) - Get a payment by ID
* [put_payments_id_](#put_payments_id_) - Update a payment
* [delete_payments_id_](#delete_payments_id_) - Delete a payment
* [post_payments_id_process](#post_payments_id_process) - Process a payment

## get_payments

List payments with the specified filter

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/payments" method="get" path="/payments" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.payments.get_payments()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `currency`                                                                             | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `destination_id`                                                                       | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `destination_type`                                                                     | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `end_time`                                                                             | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `expand`                                                                               | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `gateway_payment_id`                                                                   | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `gateway_tracking_id`                                                                  | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | For filtering by gateway tracking ID                                                   |
| `limit`                                                                                | *Optional[int]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `offset`                                                                               | *Optional[int]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `order`                                                                                | [Optional[operations.GetPaymentsOrder]](../../models/operations/getpaymentsorder.md)   | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `payment_gateway`                                                                      | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `payment_ids`                                                                          | List[*str*]                                                                            | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `payment_method_type`                                                                  | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `payment_status`                                                                       | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `sort`                                                                                 | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `start_time`                                                                           | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `status`                                                                               | [Optional[operations.GetPaymentsStatus]](../../models/operations/getpaymentsstatus.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `retries`                                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                       | :heavy_minus_sign:                                                                     | Configuration to override the default retry behavior of the client.                    |

### Response

**[components.DtoListPaymentsResponse](../../models/components/dtolistpaymentsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_payments

Create a new payment with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/payments" method="post" path="/payments" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.payments.post_payments(amount="910.96", currency="Costa Rican Colon", destination_id="<id>", destination_type="INVOICE", payment_method_type="CREDITS", process_payment=True, save_card_and_make_default=False)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `amount`                                                                                           | *str*                                                                                              | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `currency`                                                                                         | *str*                                                                                              | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `destination_id`                                                                                   | *str*                                                                                              | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `destination_type`                                                                                 | [components.TypesPaymentDestinationType](../../models/components/typespaymentdestinationtype.md)   | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `payment_method_type`                                                                              | [components.TypesPaymentMethodType](../../models/components/typespaymentmethodtype.md)             | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `cancel_url`                                                                                       | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `idempotency_key`                                                                                  | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `metadata`                                                                                         | Dict[str, *str*]                                                                                   | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `payment_gateway`                                                                                  | [Optional[components.TypesPaymentGatewayType]](../../models/components/typespaymentgatewaytype.md) | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `payment_method_id`                                                                                | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `process_payment`                                                                                  | *Optional[bool]*                                                                                   | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `save_card_and_make_default`                                                                       | *Optional[bool]*                                                                                   | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `success_url`                                                                                      | *Optional[str]*                                                                                    | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `retries`                                                                                          | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                   | :heavy_minus_sign:                                                                                 | Configuration to override the default retry behavior of the client.                                |

### Response

**[components.DtoPaymentResponse](../../models/components/dtopaymentresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_payments_id_

Get a payment by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/payments/{id}" method="get" path="/payments/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.payments.get_payments_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Payment ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPaymentResponse](../../models/components/dtopaymentresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_payments_id_

Update a payment with the specified configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/payments/{id}" method="put" path="/payments/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.payments.put_payments_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Payment ID                                                          |
| `error_message`                                                     | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `failed_at`                                                         | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `gateway_payment_id`                                                | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `metadata`                                                          | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `payment_gateway`                                                   | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `payment_method_id`                                                 | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `payment_status`                                                    | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `succeeded_at`                                                      | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPaymentResponse](../../models/components/dtopaymentresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## delete_payments_id_

Delete a payment

### Example Usage

<!-- UsageSnippet language="python" operationID="delete_/payments/{id}" method="delete" path="/payments/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.payments.delete_payments_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Payment ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_payments_id_process

Process a payment

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/payments/{id}/process" method="post" path="/payments/{id}/process" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.payments.post_payments_id_process(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Payment ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoPaymentResponse](../../models/components/dtopaymentresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |