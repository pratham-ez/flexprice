# Events

## Overview

### Available Operations

* [ingest](#ingest) - Ingest event
* [get_analytics](#get_analytics) - Get usage analytics
* [bulk_ingest](#bulk_ingest) - Bulk Ingest events
* [huggingface_inference](#huggingface_inference) - Get hugging face inference data
* [get_monitoring_data](#get_monitoring_data) - Get monitoring data
* [query](#query) - List raw events
* [get_usage_stats](#get_usage_stats) - Get usage statistics
* [get_usage_by_meter](#get_usage_by_meter) - Get usage by meter

## ingest

Ingest a new event into the system

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events" method="post" path="/events" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.events.ingest(event_name="api_request", external_customer_id="customer456", customer_id="customer456", event_id="event123", properties={
        ""response_status"": "200}",
        "{"request_size"": "100",
    }, source="api", timestamp="2024-03-20T15:04:05Z")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `event_name`                                                        | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 | api_request                                                         |
| `external_customer_id`                                              | *str*                                                               | :heavy_check_mark:                                                  | N/A                                                                 | customer456                                                         |
| `customer_id`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 | customer456                                                         |
| `event_id`                                                          | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 | event123                                                            |
| `properties`                                                        | Dict[str, *str*]                                                    | :heavy_minus_sign:                                                  | Handled separately for dynamic columns                              | {<br/>"\"response_status\"": "200}",<br/>"{\"request_size\"": "100"<br/>} |
| `source`                                                            | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 | api                                                                 |
| `timestamp`                                                         | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Handled separately due to parsing                                   | 2024-03-20T15:04:05Z                                                |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |                                                                     |

### Response

**[Dict[str, str]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_analytics

Retrieve comprehensive usage analytics with filtering, grouping, and time-series data

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/analytics" method="post" path="/events/analytics" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.events.get_analytics(request=open("example.file", "rb"))

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `request`                                                                                 | [Union[bytes, IO[bytes], io.BufferedReader]](../../models/dtogetusageanalyticsrequest.md) | :heavy_check_mark:                                                                        | The request object to use for the request.                                                |
| `retries`                                                                                 | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                          | :heavy_minus_sign:                                                                        | Configuration to override the default retry behavior of the client.                       |

### Response

**[models.DtoGetUsageAnalyticsResponse](../../models/dtogetusageanalyticsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## bulk_ingest

Ingest bulk events into the system

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/bulk" method="post" path="/events/bulk" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.events.bulk_ingest(events=[])

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `events`                                                                    | List[[models.DtoIngestEventRequest](../../models/dtoingesteventrequest.md)] | :heavy_check_mark:                                                          | N/A                                                                         |
| `retries`                                                                   | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)            | :heavy_minus_sign:                                                          | Configuration to override the default retry behavior of the client.         |

### Response

**[Dict[str, str]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## huggingface_inference

Retrieve hugging face inference data for events

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/huggingface-inference" method="post" path="/events/huggingface-inference" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.events.huggingface_inference()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoGetHuggingFaceBillingDataResponse](../../models/dtogethuggingfacebillingdataresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_monitoring_data

Retrieve monitoring data for events including consumer lag and event metrics (last 24 hours by default)

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/events/monitoring" method="get" path="/events/monitoring" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.events.get_monitoring_data()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `window_size`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Window size for time series data (e.g., 'HOUR', 'DAY') - optional   |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoGetMonitoringDataResponse](../../models/dtogetmonitoringdataresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## query

Retrieve raw events with pagination and filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/query" method="post" path="/events/query" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.events.query(request=open("example.file", "rb"))

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `request`                                                                         | [Union[bytes, IO[bytes], io.BufferedReader]](../../models/dtogeteventsrequest.md) | :heavy_check_mark:                                                                | The request object to use for the request.                                        |
| `retries`                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                  | :heavy_minus_sign:                                                                | Configuration to override the default retry behavior of the client.               |

### Response

**[models.DtoGetEventsResponse](../../models/dtogeteventsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_usage_stats

Retrieve aggregated usage statistics for events

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/usage" method="post" path="/events/usage" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.events.get_usage_stats(request=open("example.file", "rb"))

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [Union[bytes, IO[bytes], io.BufferedReader]](../../models/dtogetusagerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `retries`                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                 | :heavy_minus_sign:                                                               | Configuration to override the default retry behavior of the client.              |

### Response

**[models.DtoGetUsageResponse](../../models/dtogetusageresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_usage_by_meter

Retrieve aggregated usage statistics using meter configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/usage/meter" method="post" path="/events/usage/meter" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.events.get_usage_by_meter(request=open("example.file", "rb"))

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `request`                                                                               | [Union[bytes, IO[bytes], io.BufferedReader]](../../models/dtogetusagebymeterrequest.md) | :heavy_check_mark:                                                                      | The request object to use for the request.                                              |
| `retries`                                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                        | :heavy_minus_sign:                                                                      | Configuration to override the default retry behavior of the client.                     |

### Response

**[models.DtoGetUsageResponse](../../models/dtogetusageresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |