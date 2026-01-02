# AlertLogs

## Overview

### Available Operations

* [post_alert_search](#post_alert_search) - List alert logs by filter

## post_alert_search

List alert logs by filter with optional expand for customer, wallet, and feature

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/alert/search" method="post" path="/alert/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.alert_logs.post_alert_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `alert_status`                                                                                       | [Optional[components.TypesAlertState]](../../models/components/typesalertstate.md)                   | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `alert_type`                                                                                         | [Optional[components.TypesAlertType]](../../models/components/typesalerttype.md)                     | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `customer_id`                                                                                        | *Optional[str]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `end_time`                                                                                           | *Optional[str]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `entity_id`                                                                                          | *Optional[str]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `entity_type`                                                                                        | [Optional[components.TypesAlertEntityType]](../../models/components/typesalertentitytype.md)         | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `expand`                                                                                             | *Optional[str]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `filters`                                                                                            | List[[components.TypesFilterCondition](../../models/components/typesfiltercondition.md)]             | :heavy_minus_sign:                                                                                   | filters allows complex filtering based on multiple fields                                            |
| `limit`                                                                                              | *Optional[int]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `offset`                                                                                             | *Optional[int]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `order`                                                                                              | [Optional[components.TypesAlertLogFilterOrder]](../../models/components/typesalertlogfilterorder.md) | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `sort`                                                                                               | List[[components.TypesSortCondition](../../models/components/typessortcondition.md)]                 | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `start_time`                                                                                         | *Optional[str]*                                                                                      | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `status`                                                                                             | [Optional[components.TypesStatus]](../../models/components/typesstatus.md)                           | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `retries`                                                                                            | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                     | :heavy_minus_sign:                                                                                   | Configuration to override the default retry behavior of the client.                                  |

### Response

**[components.DtoListAlertLogsResponse](../../models/components/dtolistalertlogsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |