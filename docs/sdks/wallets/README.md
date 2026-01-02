# Wallets

## Overview

### Available Operations

* [get_customers_wallets](#get_customers_wallets) - Get Customer Wallets
* [get_customers_id_wallets](#get_customers_id_wallets) - Get wallets by customer ID
* [get_wallets](#get_wallets) - List wallets
* [post_wallets](#post_wallets) - Create a new wallet
* [post_wallets_search](#post_wallets_search) - List wallets by filter
* [post_wallets_transactions_search](#post_wallets_transactions_search) - List wallet transactions by filter
* [get_wallets_id_](#get_wallets_id_) - Get wallet by ID
* [put_wallets_id_](#put_wallets_id_) - Update a wallet
* [get_wallets_id_balance_real_time](#get_wallets_id_balance_real_time) - Get wallet balance
* [post_wallets_id_terminate](#post_wallets_id_terminate) - Terminate a wallet
* [post_wallets_id_top_up](#post_wallets_id_top_up) - Top up wallet
* [get_wallets_id_transactions](#get_wallets_id_transactions) - Get wallet transactions

## get_customers_wallets

Get all wallets for a customer by lookup key or id

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/wallets" method="get" path="/customers/wallets" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.get_customers_wallets(include_real_time_balance=False)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `include_real_time_balance`                                         | *Optional[bool]*                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `lookup_key`                                                        | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[List[components.DtoWalletResponse]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_customers_id_wallets

Get all wallets for a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/{id}/wallets" method="get" path="/customers/{id}/wallets" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.get_customers_id_wallets(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Customer ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[List[components.DtoWalletResponse]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_wallets

List wallets with optional filtering

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/wallets" method="get" path="/wallets" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.get_wallets()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `alert_enabled`                                                                      | *Optional[bool]*                                                                     | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `expand`                                                                             | *Optional[str]*                                                                      | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `limit`                                                                              | *Optional[int]*                                                                      | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `offset`                                                                             | *Optional[int]*                                                                      | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `order`                                                                              | [Optional[operations.GetWalletsOrder]](../../models/operations/getwalletsorder.md)   | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `sort`                                                                               | *Optional[str]*                                                                      | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `status`                                                                             | [Optional[operations.GetWalletsStatus]](../../models/operations/getwalletsstatus.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `wallet_ids`                                                                         | List[*str*]                                                                          | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `retries`                                                                            | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                     | :heavy_minus_sign:                                                                   | Configuration to override the default retry behavior of the client.                  |

### Response

**[components.TypesListResponseDtoWalletResponse](../../models/components/typeslistresponsedtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_wallets

Create a new wallet for a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/wallets" method="post" path="/wallets" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.post_wallets(currency="Belize Dollar", conversion_rate="1", initial_credits_to_load="0")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `currency`                                                                                                                                                                                                                                                               | *str*                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                      |
| `alert_config`                                                                                                                                                                                                                                                           | [Optional[components.DtoAlertConfig]](../../models/components/dtoalertconfig.md)                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                      |
| `alert_enabled`                                                                                                                                                                                                                                                          | *Optional[bool]*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                       | alert_enabled is the flag to enable alerts for the wallet<br/>defaults to true, can be explicitly set to false to disable alerts                                                                                                                                         |
| `auto_topup`                                                                                                                                                                                                                                                             | [Optional[components.TypesAutoTopup]](../../models/components/typesautotopup.md)                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                      |
| `config`                                                                                                                                                                                                                                                                 | [Optional[components.TypesWalletConfig]](../../models/components/typeswalletconfig.md)                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                      |
| `conversion_rate`                                                                                                                                                                                                                                                        | *Optional[str]*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                       | amount in the currency =  number of credits * conversion_rate<br/>ex if conversion_rate is 1, then 1 USD = 1 credit<br/>ex if conversion_rate is 2, then 1 USD = 0.5 credits<br/>ex if conversion_rate is 0.5, then 1 USD = 2 credits                                    |
| `customer_id`                                                                                                                                                                                                                                                            | *Optional[str]*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                      |
| `description`                                                                                                                                                                                                                                                            | *Optional[str]*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                      |
| `external_customer_id`                                                                                                                                                                                                                                                   | *Optional[str]*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                       | external_customer_id is the customer id in the external system                                                                                                                                                                                                           |
| `initial_credits_expiry_date_utc`                                                                                                                                                                                                                                        | *Optional[str]*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                       | initial_credits_expiry_date_utc is the expiry date in UTC timezone (optional to set nil means no expiry)<br/>ex 2025-01-01 00:00:00 UTC                                                                                                                                  |
| `initial_credits_to_load`                                                                                                                                                                                                                                                | *Optional[str]*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                       | initial_credits_to_load is the number of credits to load to the wallet<br/>if not provided, the wallet will be created with 0 balance<br/>NOTE: this is not the amount in the currency, but the number of credits                                                        |
| `initial_credits_to_load_expiry_date`                                                                                                                                                                                                                                    | *Optional[int]*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                       | initial_credits_to_load_expiry_date YYYYMMDD format in UTC timezone (optional to set nil means no expiry)<br/>for ex 20250101 means the credits will expire on 2025-01-01 00:00:00 UTC<br/>hence they will be available for use until 2024-12-31 23:59:59 UTC            |
| `metadata`                                                                                                                                                                                                                                                               | Dict[str, *str*]                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                      |
| `name`                                                                                                                                                                                                                                                                   | *Optional[str]*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                      |
| `price_unit`                                                                                                                                                                                                                                                             | *Optional[str]*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                       | price_unit is the code of the price unit to use for wallet creation<br/>If provided, the price unit will be used to set the currency and conversion rate of the wallet:<br/>- currency: set to price unit's base_currency<br/>- conversion_rate: set to price unit's conversion_rate |
| `wallet_type`                                                                                                                                                                                                                                                            | [Optional[components.TypesWalletType]](../../models/components/typeswallettype.md)                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                      |
| `retries`                                                                                                                                                                                                                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                       | Configuration to override the default retry behavior of the client.                                                                                                                                                                                                      |

### Response

**[components.DtoWalletResponse](../../models/components/dtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_wallets_search

List wallets by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/wallets/search" method="post" path="/wallets/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.post_wallets_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [components.TypesWalletFilter](../../models/components/typeswalletfilter.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `retries`                                                                    | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)             | :heavy_minus_sign:                                                           | Configuration to override the default retry behavior of the client.          |

### Response

**[components.TypesListResponseDtoWalletResponse](../../models/components/typeslistresponsedtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_wallets_transactions_search

List wallet transactions by filter

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/wallets/transactions/search" method="post" path="/wallets/transactions/search" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.post_wallets_transactions_search()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `expand_param`                                                                                                         | *Optional[str]*                                                                                                        | :heavy_minus_sign:                                                                                                     | Expand fields (e.g., customer,created_by_user,wallet)                                                                  |
| `created_by`                                                                                                           | *Optional[str]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `credits_available_gt`                                                                                                 | *Optional[float]*                                                                                                      | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `end_time`                                                                                                             | *Optional[str]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `expand`                                                                                                               | *Optional[str]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `expiry_date_after`                                                                                                    | *Optional[str]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `expiry_date_before`                                                                                                   | *Optional[str]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `filters`                                                                                                              | List[[components.TypesFilterCondition](../../models/components/typesfiltercondition.md)]                               | :heavy_minus_sign:                                                                                                     | filters allows complex filtering based on multiple fields                                                              |
| `id`                                                                                                                   | *Optional[str]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `limit`                                                                                                                | *Optional[int]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `offset`                                                                                                               | *Optional[int]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `order`                                                                                                                | [Optional[components.TypesWalletTransactionFilterOrder]](../../models/components/typeswallettransactionfilterorder.md) | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `priority`                                                                                                             | *Optional[int]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `reference_id`                                                                                                         | *Optional[str]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `reference_type`                                                                                                       | *Optional[str]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `sort`                                                                                                                 | List[[components.TypesSortCondition](../../models/components/typessortcondition.md)]                                   | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `start_time`                                                                                                           | *Optional[str]*                                                                                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `status`                                                                                                               | [Optional[components.TypesStatus]](../../models/components/typesstatus.md)                                             | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `transaction_reason`                                                                                                   | [Optional[components.TypesTransactionReason]](../../models/components/typestransactionreason.md)                       | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `transaction_status`                                                                                                   | [Optional[components.TypesTransactionStatus]](../../models/components/typestransactionstatus.md)                       | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `type`                                                                                                                 | [Optional[components.TypesTransactionType]](../../models/components/typestransactiontype.md)                           | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `retries`                                                                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                       | :heavy_minus_sign:                                                                                                     | Configuration to override the default retry behavior of the client.                                                    |

### Response

**[components.DtoListWalletTransactionsResponse](../../models/components/dtolistwallettransactionsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_wallets_id_

Get a wallet by its ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/wallets/{id}" method="get" path="/wallets/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.get_wallets_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Wallet ID                                                           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoWalletResponse](../../models/components/dtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## put_wallets_id_

Update a wallet's details including auto top-up configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/wallets/{id}" method="put" path="/wallets/{id}" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.put_wallets_id_(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `id`                                                                                   | *str*                                                                                  | :heavy_check_mark:                                                                     | Wallet ID                                                                              |
| `alert_config`                                                                         | [Optional[components.DtoAlertConfig]](../../models/components/dtoalertconfig.md)       | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `alert_enabled`                                                                        | *Optional[bool]*                                                                       | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `auto_topup`                                                                           | [Optional[components.TypesAutoTopup]](../../models/components/typesautotopup.md)       | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `config`                                                                               | [Optional[components.TypesWalletConfig]](../../models/components/typeswalletconfig.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `description`                                                                          | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `metadata`                                                                             | Dict[str, *str*]                                                                       | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `name`                                                                                 | *Optional[str]*                                                                        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `retries`                                                                              | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                       | :heavy_minus_sign:                                                                     | Configuration to override the default retry behavior of the client.                    |

### Response

**[components.DtoWalletResponse](../../models/components/dtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_wallets_id_balance_real_time

Get real-time balance of a wallet

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/wallets/{id}/balance/real-time" method="get" path="/wallets/{id}/balance/real-time" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.get_wallets_id_balance_real_time(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Wallet ID                                                           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoWalletBalanceResponse](../../models/components/dtowalletbalanceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_wallets_id_terminate

Terminates a wallet by closing it and debiting remaining balance

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/wallets/{id}/terminate" method="post" path="/wallets/{id}/terminate" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.post_wallets_id_terminate(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Wallet ID                                                           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[components.DtoWalletResponse](../../models/components/dtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## post_wallets_id_top_up

Add credits to a wallet

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/wallets/{id}/top-up" method="post" path="/wallets/{id}/top-up" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.post_wallets_id_top_up(id="<id>", transaction_reason="FREE_CREDIT_GRANT")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `id`                                                                                                                                                                                                                                                                                                                                                                                     | *str*                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | Wallet ID                                                                                                                                                                                                                                                                                                                                                                                |
| `transaction_reason`                                                                                                                                                                                                                                                                                                                                                                     | [components.TypesTransactionReason](../../models/components/typestransactionreason.md)                                                                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                                                                                                                                      |
| `amount`                                                                                                                                                                                                                                                                                                                                                                                 | *Optional[str]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | amount is the amount in the currency of the wallet to be added<br/>NOTE: this is not the number of credits to add, but the amount in the currency<br/>amount = credits_to_add * conversion_rate<br/>if both amount and credits_to_add are provided, amount will be ignored<br/>ex if the wallet has a conversion_rate of 2 then adding an amount of<br/>10 USD in the wallet wil add 5 credits in the wallet |
| `credits_to_add`                                                                                                                                                                                                                                                                                                                                                                         | *Optional[str]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | credits_to_add is the number of credits to add to the wallet                                                                                                                                                                                                                                                                                                                             |
| `description`                                                                                                                                                                                                                                                                                                                                                                            | *Optional[str]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | description to add any specific details about the transaction                                                                                                                                                                                                                                                                                                                            |
| `expiry_date_utc`                                                                                                                                                                                                                                                                                                                                                                        | *Optional[str]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | expiry_date_utc is the expiry date in UTC timezone<br/>ex 2025-01-01 00:00:00 UTC                                                                                                                                                                                                                                                                                                        |
| `idempotency_key`                                                                                                                                                                                                                                                                                                                                                                        | *Optional[str]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | idempotency_key is a unique key for the transaction                                                                                                                                                                                                                                                                                                                                      |
| `metadata`                                                                                                                                                                                                                                                                                                                                                                               | Dict[str, *str*]                                                                                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                                                                                                                                      |
| `priority`                                                                                                                                                                                                                                                                                                                                                                               | *Optional[int]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | priority is the priority of the transaction<br/>lower number means higher priority<br/>default is nil which means no priority at all                                                                                                                                                                                                                                                     |
| `retries`                                                                                                                                                                                                                                                                                                                                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | Configuration to override the default retry behavior of the client.                                                                                                                                                                                                                                                                                                                      |

### Response

**[components.DtoTopUpWalletResponse](../../models/components/dtotopupwalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_wallets_id_transactions

Get transactions for a wallet with pagination

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/wallets/{id}/transactions" method="get" path="/wallets/{id}/transactions" -->
```python
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.wallets.get_wallets_id_transactions(id_path_parameter="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `id_path_parameter`                                                                                              | *str*                                                                                                            | :heavy_check_mark:                                                                                               | Wallet ID                                                                                                        |
| `created_by`                                                                                                     | *Optional[str]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `credits_available_gt`                                                                                           | *Optional[float]*                                                                                                | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `end_time`                                                                                                       | *Optional[str]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `expand`                                                                                                         | *Optional[str]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `expiry_date_after`                                                                                              | *Optional[str]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `expiry_date_before`                                                                                             | *Optional[str]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `id_query_parameter`                                                                                             | *Optional[str]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `limit`                                                                                                          | *Optional[int]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `offset`                                                                                                         | *Optional[int]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `order`                                                                                                          | [Optional[operations.GetWalletsIDTransactionsOrder]](../../models/operations/getwalletsidtransactionsorder.md)   | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `priority`                                                                                                       | *Optional[int]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `reference_id`                                                                                                   | *Optional[str]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `reference_type`                                                                                                 | *Optional[str]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `start_time`                                                                                                     | *Optional[str]*                                                                                                  | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `status`                                                                                                         | [Optional[operations.GetWalletsIDTransactionsStatus]](../../models/operations/getwalletsidtransactionsstatus.md) | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `transaction_reason`                                                                                             | [Optional[operations.TransactionReason]](../../models/operations/transactionreason.md)                           | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `transaction_status`                                                                                             | [Optional[operations.TransactionStatus]](../../models/operations/transactionstatus.md)                           | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `type`                                                                                                           | [Optional[operations.Type]](../../models/operations/type.md)                                                     | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `retries`                                                                                                        | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                 | :heavy_minus_sign:                                                                                               | Configuration to override the default retry behavior of the client.                                              |

### Response

**[components.DtoListWalletTransactionsResponse](../../models/components/dtolistwallettransactionsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |