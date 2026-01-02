# Events

## Overview

### Available Operations

* [post_events](#post_events) - Ingest event
* [post_events_analytics](#post_events_analytics) - Get usage analytics
* [post_events_bulk](#post_events_bulk) - Bulk Ingest events
* [post_events_huggingface_inference](#post_events_huggingface_inference) - Get hugging face inference data
* [get_events_monitoring](#get_events_monitoring) - Get monitoring data
* [post_events_query](#post_events_query) - List raw events
* [post_events_usage](#post_events_usage) - Get usage statistics
* [post_events_usage_meter](#post_events_usage_meter) - Get usage by meter

## post_events

Ingest a new event into the system

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events" method="post" path="/events" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.events.post_events(event_name="api_request", external_customer_id="customer456", customer_id="customer456", event_id="event123", properties={
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
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_events_analytics

Retrieve comprehensive usage analytics with filtering, grouping, and time-series data

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/analytics" method="post" path="/events/analytics" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.events.post_events_analytics(request=open("example.file", "rb"))

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `request`                                                                                 | [Union[bytes, IO[bytes], io.BufferedReader]](../../models/dtogetusageanalyticsrequest.md) | :heavy_check_mark:                                                                        | The request object to use for the request.                                                |
| `retries`                                                                                 | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                          | :heavy_minus_sign:                                                                        | Configuration to override the default retry behavior of the client.                       |

### Response

**[components.DtoGetUsageAnalyticsResponse](../../models/components/dtogetusageanalyticsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_events_bulk

Ingest bulk events into the system

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/bulk" method="post" path="/events/bulk" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.events.post_events_bulk(events=[])

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `events`                                                                                   | List[[components.DtoIngestEventRequest](../../models/components/dtoingesteventrequest.md)] | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `retries`                                                                                  | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                           | :heavy_minus_sign:                                                                         | Configuration to override the default retry behavior of the client.                        |

### Response

**[Dict[str, str]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_events_huggingface_inference

Retrieve hugging face inference data for events

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/huggingface-inference" method="post" path="/events/huggingface-inference" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.events.post_events_huggingface_inference()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoGetHuggingFaceBillingDataResponse](../../models/components/dtogethuggingfacebillingdataresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_events_monitoring

Retrieve monitoring data for events including consumer lag and event metrics (last 24 hours by default)

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/events/monitoring" method="get" path="/events/monitoring" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.events.get_events_monitoring()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `window_size`                                                       | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Window size for time series data (e.g., 'HOUR', 'DAY') - optional   |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoGetMonitoringDataResponse](../../models/components/dtogetmonitoringdataresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_events_query

Retrieve raw events with pagination and filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/query" method="post" path="/events/query" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.events.post_events_query(request=open("example.file", "rb"))

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `request`                                                                         | [Union[bytes, IO[bytes], io.BufferedReader]](../../models/dtogeteventsrequest.md) | :heavy_check_mark:                                                                | The request object to use for the request.                                        |
| `retries`                                                                         | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                  | :heavy_minus_sign:                                                                | Configuration to override the default retry behavior of the client.               |

### Response

**[components.DtoGetEventsResponse](../../models/components/dtogeteventsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_events_usage

Retrieve aggregated usage statistics for events

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/usage" method="post" path="/events/usage" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.events.post_events_usage(request=open("example.file", "rb"))

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [Union[bytes, IO[bytes], io.BufferedReader]](../../models/dtogetusagerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `retries`                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                 | :heavy_minus_sign:                                                               | Configuration to override the default retry behavior of the client.              |

### Response

**[components.DtoGetUsageResponse](../../models/components/dtogetusageresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_events_usage_meter

Retrieve aggregated usage statistics using meter configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/events/usage/meter" method="post" path="/events/usage/meter" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.events.post_events_usage_meter(request=open("example.file", "rb"))

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `request`                                                                               | [Union[bytes, IO[bytes], io.BufferedReader]](../../models/dtogetusagebymeterrequest.md) | :heavy_check_mark:                                                                      | The request object to use for the request.                                              |
| `retries`                                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                        | :heavy_minus_sign:                                                                      | Configuration to override the default retry behavior of the client.                     |

### Response

**[components.DtoGetUsageResponse](../../models/components/dtogetusageresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |