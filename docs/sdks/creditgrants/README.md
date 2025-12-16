# Creditgrants

## Overview

### Available Operations

* [get](#get) - Get credit grants

## get

Get credit grants with the specified filter

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/creditgrants" method="get" path="/creditgrants" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.creditgrants.get()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `end_time`                                                                                          | *Optional[str]*                                                                                     | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `expand`                                                                                            | *Optional[str]*                                                                                     | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `limit`                                                                                             | *Optional[int]*                                                                                     | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `offset`                                                                                            | *Optional[int]*                                                                                     | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `order`                                                                                             | [Optional[models.GetCreditgrantsQueryParamOrder]](../../models/getcreditgrantsqueryparamorder.md)   | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `plan_ids`                                                                                          | List[*str*]                                                                                         | :heavy_minus_sign:                                                                                  | Specific filters for credit grants                                                                  |
| `scope`                                                                                             | [Optional[models.Scope]](../../models/scope.md)                                                     | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `sort`                                                                                              | *Optional[str]*                                                                                     | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `start_time`                                                                                        | *Optional[str]*                                                                                     | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `status`                                                                                            | [Optional[models.GetCreditgrantsQueryParamStatus]](../../models/getcreditgrantsqueryparamstatus.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `subscription_ids`                                                                                  | List[*str*]                                                                                         | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `retries`                                                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                    | :heavy_minus_sign:                                                                                  | Configuration to override the default retry behavior of the client.                                 |

### Response

**[models.DtoListCreditGrantsResponse](../../models/dtolistcreditgrantsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |