# CreditNotes

## Overview

### Available Operations

* [get_creditnotes](#get_creditnotes) - List credit notes with filtering
* [post_creditnotes](#post_creditnotes) - Create a new credit note
* [get_creditnotes_id_](#get_creditnotes_id_) - Get a credit note by ID
* [post_creditnotes_id_finalize](#post_creditnotes_id_finalize) - Process a draft credit note
* [post_creditnotes_id_void](#post_creditnotes_id_void) - Void a credit note

## get_creditnotes

Lists credit notes with filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/creditnotes" method="get" path="/creditnotes" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_notes.get_creditnotes()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `credit_note_ids`                                                                            | List[*str*]                                                                                  | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `credit_note_status`                                                                         | List[[operations.CreditNoteStatus](../../models/operations/creditnotestatus.md)]             | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `credit_note_type`                                                                           | [Optional[operations.CreditNoteType]](../../models/operations/creditnotetype.md)             | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `end_time`                                                                                   | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `expand`                                                                                     | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `invoice_id`                                                                                 | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `limit`                                                                                      | *Optional[int]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `offset`                                                                                     | *Optional[int]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `order`                                                                                      | [Optional[operations.GetCreditnotesOrder]](../../models/operations/getcreditnotesorder.md)   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `sort`                                                                                       | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `start_time`                                                                                 | *Optional[str]*                                                                              | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `status`                                                                                     | [Optional[operations.GetCreditnotesStatus]](../../models/operations/getcreditnotesstatus.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `retries`                                                                                    | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                             | :heavy_minus_sign:                                                                           | Configuration to override the default retry behavior of the client.                          |

### Response

**[components.DtoListCreditNotesResponse](../../models/components/dtolistcreditnotesresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_creditnotes

Creates a new credit note

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/creditnotes" method="post" path="/creditnotes" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_notes.post_creditnotes(invoice_id="<id>", reason="BILLING_ERROR", process_credit_note=True)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `invoice_id`                                                                                                         | *str*                                                                                                                | :heavy_check_mark:                                                                                                   | invoice_id is the unique identifier of the invoice this credit note is applied to                                    |
| `reason`                                                                                                             | [components.TypesCreditNoteReason](../../models/components/typescreditnotereason.md)                                 | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `credit_note_number`                                                                                                 | *Optional[str]*                                                                                                      | :heavy_minus_sign:                                                                                                   | credit_note_number is an optional human-readable identifier for the credit note                                      |
| `idempotency_key`                                                                                                    | *Optional[str]*                                                                                                      | :heavy_minus_sign:                                                                                                   | idempotency_key is an optional key used to prevent duplicate credit note creation                                    |
| `line_items`                                                                                                         | List[[components.DtoCreateCreditNoteLineItemRequest](../../models/components/dtocreatecreditnotelineitemrequest.md)] | :heavy_minus_sign:                                                                                                   | line_items contains the individual line items that make up this credit note (minimum 1 required)                     |
| `memo`                                                                                                               | *Optional[str]*                                                                                                      | :heavy_minus_sign:                                                                                                   | memo is an optional free-text field for additional notes about the credit note                                       |
| `metadata`                                                                                                           | Dict[str, *str*]                                                                                                     | :heavy_minus_sign:                                                                                                   | N/A                                                                                                                  |
| `process_credit_note`                                                                                                | *Optional[bool]*                                                                                                     | :heavy_minus_sign:                                                                                                   | process_credit_note is a flag to process the credit note after creation                                              |
| `retries`                                                                                                            | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                     | :heavy_minus_sign:                                                                                                   | Configuration to override the default retry behavior of the client.                                                  |

### Response

**[components.DtoCreditNoteResponse](../../models/components/dtocreditnoteresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_creditnotes_id_

Retrieves a credit note by ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/creditnotes/{id}" method="get" path="/creditnotes/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_notes.get_creditnotes_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Credit note ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCreditNoteResponse](../../models/components/dtocreditnoteresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_creditnotes_id_finalize

Processes a draft credit note

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/creditnotes/{id}/finalize" method="post" path="/creditnotes/{id}/finalize" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_notes.post_creditnotes_id_finalize(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Credit note ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCreditNoteResponse](../../models/components/dtocreditnoteresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_creditnotes_id_void

Voids a credit note

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/creditnotes/{id}/void" method="post" path="/creditnotes/{id}/void" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.credit_notes.post_creditnotes_id_void(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Credit note ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoCreditNoteResponse](../../models/components/dtocreditnoteresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |