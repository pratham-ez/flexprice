# CreditNotes

## Overview

### Available Operations

* [List](#list) - List credit notes with filtering
* [Create](#create) - Create a new credit note
* [Finalize](#finalize) - Process a draft credit note
* [Void](#void) - Void a credit note

## List

Lists credit notes with filtering

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/creditnotes" method="get" path="/creditnotes" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.CreditNotes.List(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListCreditNotesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |
| `creditNoteIds`                                                                                         | []*string*                                                                                              | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `creditNoteStatus`                                                                                      | [][operations.CreditNoteStatus](../../models/operations/creditnotestatus.md)                            | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `creditNoteType`                                                                                        | [*operations.CreditNoteType](../../models/operations/creditnotetype.md)                                 | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `endTime`                                                                                               | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `expand`                                                                                                | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `invoiceID`                                                                                             | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `limit`                                                                                                 | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `offset`                                                                                                | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `order`                                                                                                 | [*operations.GetCreditnotesQueryParamOrder](../../models/operations/getcreditnotesqueryparamorder.md)   | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `sort`                                                                                                  | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `startTime`                                                                                             | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `status`                                                                                                | [*operations.GetCreditnotesQueryParamStatus](../../models/operations/getcreditnotesqueryparamstatus.md) | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |

### Response

**[*operations.GetCreditnotesResponse](../../models/operations/getcreditnotesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 401, 403, 404            | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Creates a new credit note

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/creditnotes" method="post" path="/creditnotes" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.CreditNotes.Create(ctx, components.DtoCreateCreditNoteRequest{
        InvoiceID: "<id>",
        Reason: components.TypesCreditNoteReasonBillingError,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.DtoCreateCreditNoteRequest](../../models/components/dtocreatecreditnoterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PostCreditnotesResponse](../../models/operations/postcreditnotesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 401, 403, 404            | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Finalize

Processes a draft credit note

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/creditnotes/{id}/finalize" method="post" path="/creditnotes/{id}/finalize" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.CreditNotes.Finalize(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Credit note ID                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostCreditnotesIDFinalizeResponse](../../models/operations/postcreditnotesidfinalizeresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 401, 403, 404            | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Void

Voids a credit note

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/creditnotes/{id}/void" method="post" path="/creditnotes/{id}/void" -->
```go
package main

import(
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := gosdk.New(
        "https://api.example.com",
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.CreditNotes.Void(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Credit note ID                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostCreditnotesIDVoidResponse](../../models/operations/postcreditnotesidvoidresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400, 401, 403, 404            | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |