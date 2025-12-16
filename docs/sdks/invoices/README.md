# Invoices

## Overview

### Available Operations

* [get_summary_by_customer_id](#get_summary_by_customer_id) - Get a customer invoice summary
* [list](#list) - List invoices
* [create](#create) - Create a new one off invoice
* [preview](#preview) - Get a preview invoice
* [search](#search) - List invoices by filter
* [get_by_id](#get_by_id) - Get an invoice by ID
* [update](#update) - Update an invoice
* [trigger_comms](#trigger_comms) - Trigger communication webhook for an invoice
* [finalize](#finalize) - Finalize an invoice
* [update_payment_status](#update_payment_status) - Update invoice payment status
* [initiate_payment](#initiate_payment) - Attempt payment for an invoice
* [get_pdf](#get_pdf) - Get PDF for an invoice
* [recalculate](#recalculate) - Recalculate invoice totals and line items
* [void](#void) - Void an invoice

## get_summary_by_customer_id

Get a customer invoice summary

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/{id}/invoices/summary" method="get" path="/customers/{id}/invoices/summary" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.get_summary_by_customer_id(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Customer ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoCustomerMultiCurrencyInvoiceSummary](../../models/dtocustomermulticurrencyinvoicesummary.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## list

List invoices with optional filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/invoices" method="get" path="/invoices" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.list()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `request`                                                           | [models.GetInvoicesRequest](../../models/getinvoicesrequest.md)     | :heavy_check_mark:                                                  | The request object to use for the request.                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoListInvoicesResponse](../../models/dtolistinvoicesresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create

Create a new one off invoice with the provided details

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/invoices" method="post" path="/invoices" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.create(amount_due="<value>", currency="Lilangeni", customer_id="<id>", subtotal="<value>", total="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                         | Type                                                                                                                                                                              | Required                                                                                                                                                                          | Description                                                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `amount_due`                                                                                                                                                                      | *str*                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                | amount_due is the total amount that needs to be paid for this invoice                                                                                                             |
| `currency`                                                                                                                                                                        | *str*                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                | currency is the three-letter ISO currency code (e.g., USD, EUR) for the invoice                                                                                                   |
| `customer_id`                                                                                                                                                                     | *str*                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                | customer_id is the unique identifier of the customer this invoice belongs to                                                                                                      |
| `subtotal`                                                                                                                                                                        | *str*                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                | subtotal is the amount before taxes and discounts are applied                                                                                                                     |
| `total`                                                                                                                                                                           | *str*                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                | total is the total amount of the invoice including taxes and discounts                                                                                                            |
| `amount_paid`                                                                                                                                                                     | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | amount_paid is the amount that has been paid towards this invoice                                                                                                                 |
| `billing_period`                                                                                                                                                                  | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | billing_period is the period this invoice covers (e.g., "monthly", "yearly")                                                                                                      |
| `billing_reason`                                                                                                                                                                  | [Optional[models.TypesInvoiceBillingReason]](../../models/typesinvoicebillingreason.md)                                                                                           | :heavy_minus_sign:                                                                                                                                                                | N/A                                                                                                                                                                               |
| `coupons`                                                                                                                                                                         | List[*str*]                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                | coupons                                                                                                                                                                           |
| `description`                                                                                                                                                                     | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | description is an optional text description of the invoice                                                                                                                        |
| `due_date`                                                                                                                                                                        | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | due_date is the date by which payment is expected                                                                                                                                 |
| `environment_id`                                                                                                                                                                  | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | environment_id is the unique identifier of the environment this invoice belongs to                                                                                                |
| `idempotency_key`                                                                                                                                                                 | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | idempotency_key is an optional key used to prevent duplicate invoice creation                                                                                                     |
| `invoice_coupons`                                                                                                                                                                 | List[[models.DtoInvoiceCoupon](../../models/dtoinvoicecoupon.md)]                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                | Invoice Coupns                                                                                                                                                                    |
| `invoice_number`                                                                                                                                                                  | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | invoice_number is an optional human-readable identifier for the invoice                                                                                                           |
| `invoice_pdf_url`                                                                                                                                                                 | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | invoice_pdf_url is the URL where customers can download the PDF version of this invoice                                                                                           |
| `invoice_status`                                                                                                                                                                  | [Optional[models.TypesInvoiceStatus]](../../models/typesinvoicestatus.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                                | N/A                                                                                                                                                                               |
| `invoice_type`                                                                                                                                                                    | [Optional[models.TypesInvoiceType]](../../models/typesinvoicetype.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                                | N/A                                                                                                                                                                               |
| `line_item_coupons`                                                                                                                                                               | List[[models.DtoInvoiceLineItemCoupon](../../models/dtoinvoicelineitemcoupon.md)]                                                                                                 | :heavy_minus_sign:                                                                                                                                                                | Invoice Line Item Coupons                                                                                                                                                         |
| `line_items`                                                                                                                                                                      | List[[models.DtoCreateInvoiceLineItemRequest](../../models/dtocreateinvoicelineitemrequest.md)]                                                                                   | :heavy_minus_sign:                                                                                                                                                                | line_items contains the individual items that make up this invoice                                                                                                                |
| `metadata`                                                                                                                                                                        | Dict[str, *str*]                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                | N/A                                                                                                                                                                               |
| `payment_status`                                                                                                                                                                  | [Optional[models.TypesPaymentStatus]](../../models/typespaymentstatus.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                                                | N/A                                                                                                                                                                               |
| `period_end`                                                                                                                                                                      | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | period_end is the end date of the billing period                                                                                                                                  |
| `period_start`                                                                                                                                                                    | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | period_start is the start date of the billing period                                                                                                                              |
| `prepared_tax_rates`                                                                                                                                                              | List[[models.DtoTaxRateResponse](../../models/dtotaxrateresponse.md)]                                                                                                             | :heavy_minus_sign:                                                                                                                                                                | prepared_tax_rates contains the tax rates pre-resolved by the caller (e.g., billing service)<br/>These are applied at invoice level by the invoice service without further resolution |
| `subscription_id`                                                                                                                                                                 | *Optional[str]*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                | subscription_id is the optional unique identifier of the subscription associated with this invoice                                                                                |
| `tax_rate_overrides`                                                                                                                                                              | List[[models.DtoTaxRateOverride](../../models/dtotaxrateoverride.md)]                                                                                                             | :heavy_minus_sign:                                                                                                                                                                | tax_rate_overrides is the tax rate overrides to be applied to the invoice                                                                                                         |
| `tax_rates`                                                                                                                                                                       | List[*str*]                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                | tax_rates                                                                                                                                                                         |
| `retries`                                                                                                                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                | Configuration to override the default retry behavior of the client.                                                                                                               |

### Response

**[models.DtoInvoiceResponse](../../models/dtoinvoiceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## preview

Get a preview invoice

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/invoices/preview" method="post" path="/invoices/preview" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.preview(subscription_id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `subscription_id`                                                                   | *str*                                                                               | :heavy_check_mark:                                                                  | subscription_id is the unique identifier of the subscription to preview invoice for |
| `period_end`                                                                        | *Optional[str]*                                                                     | :heavy_minus_sign:                                                                  | period_end is the optional end date of the period to preview                        |
| `period_start`                                                                      | *Optional[str]*                                                                     | :heavy_minus_sign:                                                                  | period_start is the optional start date of the period to preview                    |
| `retries`                                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                    | :heavy_minus_sign:                                                                  | Configuration to override the default retry behavior of the client.                 |

### Response

**[models.DtoInvoiceResponse](../../models/dtoinvoiceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## search

List invoices by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/invoices/search" method="post" path="/invoices/search" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `amount_due_gt`                                                                                                                                                                  | *Optional[float]*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                               | amount_due_gt filters invoices with a total amount due greater than the specified value<br/>Useful for finding invoices above a certain threshold or identifying high-value invoices |
| `amount_remaining_gt`                                                                                                                                                            | *Optional[float]*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                               | amount_remaining_gt filters invoices with an outstanding balance greater than the specified value<br/>Useful for finding invoices that still have significant unpaid amounts     |
| `customer_id`                                                                                                                                                                    | *Optional[str]*                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                               | customer_id filters invoices for a specific customer using FlexPrice's internal customer ID<br/>This is the ID returned by FlexPrice when creating or retrieving customers       |
| `end_time`                                                                                                                                                                       | *Optional[str]*                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                               | N/A                                                                                                                                                                              |
| `expand`                                                                                                                                                                         | *Optional[str]*                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                               | N/A                                                                                                                                                                              |
| `external_customer_id`                                                                                                                                                           | *Optional[str]*                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                               | external_customer_id filters invoices for a customer using your system's customer identifier<br/>This is the ID you provided when creating the customer in FlexPrice             |
| `filters`                                                                                                                                                                        | List[[models.TypesFilterCondition](../../models/typesfiltercondition.md)]                                                                                                        | :heavy_minus_sign:                                                                                                                                                               | N/A                                                                                                                                                                              |
| `invoice_ids`                                                                                                                                                                    | List[*str*]                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                               | invoice_ids restricts results to invoices with the specified IDs<br/>Use this to retrieve specific invoices when you know their exact identifiers                                |
| `invoice_status`                                                                                                                                                                 | List[[models.TypesInvoiceStatus](../../models/typesinvoicestatus.md)]                                                                                                            | :heavy_minus_sign:                                                                                                                                                               | invoice_status filters by the current state of invoices in their lifecycle<br/>Multiple statuses can be specified to include invoices in any of the listed states                |
| `invoice_type`                                                                                                                                                                   | [Optional[models.TypesInvoiceType]](../../models/typesinvoicetype.md)                                                                                                            | :heavy_minus_sign:                                                                                                                                                               | N/A                                                                                                                                                                              |
| `limit`                                                                                                                                                                          | *Optional[int]*                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                               | N/A                                                                                                                                                                              |
| `offset`                                                                                                                                                                         | *Optional[int]*                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                               | N/A                                                                                                                                                                              |
| `order`                                                                                                                                                                          | [Optional[models.TypesInvoiceFilterOrder]](../../models/typesinvoicefilterorder.md)                                                                                              | :heavy_minus_sign:                                                                                                                                                               | N/A                                                                                                                                                                              |
| `payment_status`                                                                                                                                                                 | List[[models.TypesPaymentStatus](../../models/typespaymentstatus.md)]                                                                                                            | :heavy_minus_sign:                                                                                                                                                               | payment_status filters by the payment state of invoices<br/>Multiple statuses can be specified to include invoices with any of the listed payment states                         |
| `skip_line_items`                                                                                                                                                                | *Optional[bool]*                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                               | SkipLineItems if true, will not include line items in the response                                                                                                               |
| `sort`                                                                                                                                                                           | List[[models.TypesSortCondition](../../models/typessortcondition.md)]                                                                                                            | :heavy_minus_sign:                                                                                                                                                               | N/A                                                                                                                                                                              |
| `start_time`                                                                                                                                                                     | *Optional[str]*                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                               | N/A                                                                                                                                                                              |
| `status`                                                                                                                                                                         | [Optional[models.TypesStatus]](../../models/typesstatus.md)                                                                                                                      | :heavy_minus_sign:                                                                                                                                                               | N/A                                                                                                                                                                              |
| `subscription_id`                                                                                                                                                                | *Optional[str]*                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                               | subscription_id filters invoices generated for a specific subscription<br/>Only returns invoices that were created as part of the specified subscription's billing               |
| `retries`                                                                                                                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                               | Configuration to override the default retry behavior of the client.                                                                                                              |

### Response

**[models.DtoListInvoicesResponse](../../models/dtolistinvoicesresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_by_id

Get detailed information about an invoice

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/invoices/{id}" method="get" path="/invoices/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.get_by_id(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `id`                                                                                    | *str*                                                                                   | :heavy_check_mark:                                                                      | Invoice ID                                                                              |
| `expand_by_source`                                                                      | *Optional[bool]*                                                                        | :heavy_minus_sign:                                                                      | Include source-level price breakdown for usage line items (legacy)                      |
| `group_by`                                                                              | List[*str*]                                                                             | :heavy_minus_sign:                                                                      | Group usage breakdown by specified fields (e.g., source, feature_id, properties.org_id) |
| `retries`                                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                        | :heavy_minus_sign:                                                                      | Configuration to override the default retry behavior of the client.                     |

### Response

**[models.DtoInvoiceResponse](../../models/dtoinvoiceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 404                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update

Update invoice details like PDF URL and due date.

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/invoices/{id}" method="put" path="/invoices/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.update(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `id`                                                                                    | *str*                                                                                   | :heavy_check_mark:                                                                      | Invoice ID                                                                              |
| `due_date`                                                                              | *Optional[str]*                                                                         | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `invoice_pdf_url`                                                                       | *Optional[str]*                                                                         | :heavy_minus_sign:                                                                      | invoice_pdf_url is the URL where customers can download the PDF version of this invoice |
| `metadata`                                                                              | Dict[str, *str*]                                                                        | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `retries`                                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                        | :heavy_minus_sign:                                                                      | Configuration to override the default retry behavior of the client.                     |

### Response

**[models.DtoInvoiceResponse](../../models/dtoinvoiceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## trigger_comms

Triggers a communication webhook event containing all information about the invoice

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/invoices/{id}/comms/trigger" method="post" path="/invoices/{id}/comms/trigger" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.trigger_comms(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Invoice ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, models.GinH]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## finalize

Finalize a draft invoice

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/invoices/{id}/finalize" method="post" path="/invoices/{id}/finalize" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.finalize(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Invoice ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, models.GinH]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update_payment_status

Update the payment status of an invoice

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/invoices/{id}/payment" method="put" path="/invoices/{id}/payment" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.update_payment_status(id="<id>", payment_status="REFUNDED")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Invoice ID                                                          |
| `payment_status`                                                    | [models.TypesPaymentStatus](../../models/typespaymentstatus.md)     | :heavy_check_mark:                                                  | N/A                                                                 |
| `amount`                                                            | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | amount is the optional payment amount to record                     |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoInvoiceResponse](../../models/dtoinvoiceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## initiate_payment

Attempt to pay an invoice using customer's available wallets

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/invoices/{id}/payment/attempt" method="post" path="/invoices/{id}/payment/attempt" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.initiate_payment(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Invoice ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, models.GinH]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_pdf

Retrieve the PDF document for a specific invoice by its ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/invoices/{id}/pdf" method="get" path="/invoices/{id}/pdf" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.get_pdf(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Invoice ID                                                          |
| `url`                                                               | *Optional[bool]*                                                    | :heavy_minus_sign:                                                  | Return presigned URL from s3 instead of PDF                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[httpx.Response](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## recalculate

Recalculate totals and line items for a draft invoice, useful when subscription line items or usage data has changed

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/invoices/{id}/recalculate" method="post" path="/invoices/{id}/recalculate" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.recalculate(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Invoice ID                                                          |
| `finalize`                                                          | *Optional[bool]*                                                    | :heavy_minus_sign:                                                  | Whether to finalize the invoice after recalculation (default: true) |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoInvoiceResponse](../../models/dtoinvoiceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## void

Void an invoice that hasn't been paid

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/invoices/{id}/void" method="post" path="/invoices/{id}/void" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.invoices.void(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Invoice ID                                                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Dict[str, models.GinH]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |