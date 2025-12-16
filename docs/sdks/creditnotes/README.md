# CreditNotes

## Overview

### Available Operations

* [list](#list) - List credit notes with filtering
* [create](#create) - Create a new credit note
* [finalize](#finalize) - Process a draft credit note
* [void](#void) - Void a credit note

## list

Lists credit notes with filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/creditnotes" method="get" path="/creditnotes" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.credit_notes.list()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `credit_note_ids`                                                                                 | List[*str*]                                                                                       | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `credit_note_status`                                                                              | List[[models.CreditNoteStatus](../../models/creditnotestatus.md)]                                 | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `credit_note_type`                                                                                | [Optional[models.CreditNoteType]](../../models/creditnotetype.md)                                 | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `end_time`                                                                                        | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `expand`                                                                                          | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `invoice_id`                                                                                      | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `limit`                                                                                           | *Optional[int]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `offset`                                                                                          | *Optional[int]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `order`                                                                                           | [Optional[models.GetCreditnotesQueryParamOrder]](../../models/getcreditnotesqueryparamorder.md)   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `sort`                                                                                            | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `start_time`                                                                                      | *Optional[str]*                                                                                   | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `status`                                                                                          | [Optional[models.GetCreditnotesQueryParamStatus]](../../models/getcreditnotesqueryparamstatus.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `retries`                                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                  | :heavy_minus_sign:                                                                                | Configuration to override the default retry behavior of the client.                               |

### Response

**[models.DtoListCreditNotesResponse](../../models/dtolistcreditnotesresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create

Creates a new credit note

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/creditnotes" method="post" path="/creditnotes" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.credit_notes.create(invoice_id="<id>", reason="BILLING_ERROR", process_credit_note=True)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `invoice_id`                                                                                          | *str*                                                                                                 | :heavy_check_mark:                                                                                    | invoice_id is the unique identifier of the invoice this credit note is applied to                     |
| `reason`                                                                                              | [models.TypesCreditNoteReason](../../models/typescreditnotereason.md)                                 | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `credit_note_number`                                                                                  | *Optional[str]*                                                                                       | :heavy_minus_sign:                                                                                    | credit_note_number is an optional human-readable identifier for the credit note                       |
| `idempotency_key`                                                                                     | *Optional[str]*                                                                                       | :heavy_minus_sign:                                                                                    | idempotency_key is an optional key used to prevent duplicate credit note creation                     |
| `line_items`                                                                                          | List[[models.DtoCreateCreditNoteLineItemRequest](../../models/dtocreatecreditnotelineitemrequest.md)] | :heavy_minus_sign:                                                                                    | line_items contains the individual line items that make up this credit note (minimum 1 required)      |
| `memo`                                                                                                | *Optional[str]*                                                                                       | :heavy_minus_sign:                                                                                    | memo is an optional free-text field for additional notes about the credit note                        |
| `metadata`                                                                                            | Dict[str, *str*]                                                                                      | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `process_credit_note`                                                                                 | *Optional[bool]*                                                                                      | :heavy_minus_sign:                                                                                    | process_credit_note is a flag to process the credit note after creation                               |
| `retries`                                                                                             | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                      | :heavy_minus_sign:                                                                                    | Configuration to override the default retry behavior of the client.                                   |

### Response

**[models.DtoCreditNoteResponse](../../models/dtocreditnoteresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## finalize

Processes a draft credit note

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/creditnotes/{id}/finalize" method="post" path="/creditnotes/{id}/finalize" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.credit_notes.finalize(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Credit note ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoCreditNoteResponse](../../models/dtocreditnoteresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## void

Voids a credit note

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/creditnotes/{id}/void" method="post" path="/creditnotes/{id}/void" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.credit_notes.void(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Credit note ID                                                      |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoCreditNoteResponse](../../models/dtocreditnoteresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 401, 403, 404           | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |