# Creditgrants

## Overview

### Available Operations

* [Get](#get) - Get credit grants

## Get

Get credit grants with the specified filter

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/creditgrants" method="get" path="/creditgrants" -->
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

    res, err := s.Creditgrants.Get(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListCreditGrantsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |
| `endTime`                                                                                                 | **string*                                                                                                 | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `expand`                                                                                                  | **string*                                                                                                 | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `limit`                                                                                                   | **int64*                                                                                                  | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `offset`                                                                                                  | **int64*                                                                                                  | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `order`                                                                                                   | [*operations.GetCreditgrantsQueryParamOrder](../../models/operations/getcreditgrantsqueryparamorder.md)   | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `planIds`                                                                                                 | []*string*                                                                                                | :heavy_minus_sign:                                                                                        | Specific filters for credit grants                                                                        |
| `scope`                                                                                                   | [*operations.Scope](../../models/operations/scope.md)                                                     | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `sort`                                                                                                    | **string*                                                                                                 | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `startTime`                                                                                               | **string*                                                                                                 | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `status`                                                                                                  | [*operations.GetCreditgrantsQueryParamStatus](../../models/operations/getcreditgrantsqueryparamstatus.md) | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `subscriptionIds`                                                                                         | []*string*                                                                                                | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |

### Response

**[*operations.GetCreditgrantsResponse](../../models/operations/getcreditgrantsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.ErrorsErrorResponse | 400                           | application/json              |
| sdkerrors.ErrorsErrorResponse | 500                           | application/json              |
| sdkerrors.APIError            | 4XX, 5XX                      | \*/\*                         |