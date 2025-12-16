# Wallets

## Overview

### Available Operations

* [list_customer_wallets](#list_customer_wallets) - Get Customer Wallets
* [get_by_customer_id](#get_by_customer_id) - Get wallets by customer ID
* [create](#create) - Create a new wallet
* [get](#get) - Get wallet by ID
* [update](#update) - Update a wallet
* [get_balance_real_time](#get_balance_real_time) - Get wallet balance
* [debit](#debit) - Debit a wallet
* [terminate](#terminate) - Terminate a wallet
* [top_up](#top_up) - Top up wallet
* [get_transactions](#get_transactions) - Get wallet transactions

## list_customer_wallets

Get all wallets for a customer by lookup key or id

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/wallets" method="get" path="/customers/wallets" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.wallets.list_customer_wallets(include_real_time_balance=False)

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

**[List[models.DtoWalletResponse]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_by_customer_id

Get all wallets for a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/customers/{id}/wallets" method="get" path="/customers/{id}/wallets" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.wallets.get_by_customer_id(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Customer ID                                                         |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[List[models.DtoWalletResponse]](../../models/.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## create

Create a new wallet for a customer

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/wallets" method="post" path="/wallets" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.wallets.create(currency="Belize Dollar", conversion_rate="1", initial_credits_to_load="0")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                                                                                             | Type                                                                                                                                                                                                                                                  | Required                                                                                                                                                                                                                                              | Description                                                                                                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `currency`                                                                                                                                                                                                                                            | *str*                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `alert_config`                                                                                                                                                                                                                                        | [Optional[models.DtoAlertConfig]](../../models/dtoalertconfig.md)                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `alert_enabled`                                                                                                                                                                                                                                       | *Optional[bool]*                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                    | alert_enabled is the flag to enable alerts for the wallet<br/>defaults to true, can be explicitly set to false to disable alerts                                                                                                                      |
| `auto_topup_amount`                                                                                                                                                                                                                                   | *Optional[str]*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `auto_topup_min_balance`                                                                                                                                                                                                                              | *Optional[str]*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `auto_topup_trigger`                                                                                                                                                                                                                                  | [Optional[models.TypesAutoTopupTrigger]](../../models/typesautotopuptrigger.md)                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `config`                                                                                                                                                                                                                                              | [Optional[models.TypesWalletConfig]](../../models/typeswalletconfig.md)                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `conversion_rate`                                                                                                                                                                                                                                     | *Optional[str]*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | amount in the currency =  number of credits * conversion_rate<br/>ex if conversion_rate is 1, then 1 USD = 1 credit<br/>ex if conversion_rate is 2, then 1 USD = 0.5 credits<br/>ex if conversion_rate is 0.5, then 1 USD = 2 credits                 |
| `customer_id`                                                                                                                                                                                                                                         | *Optional[str]*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `description`                                                                                                                                                                                                                                         | *Optional[str]*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `external_customer_id`                                                                                                                                                                                                                                | *Optional[str]*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | external_customer_id is the customer id in the external system                                                                                                                                                                                        |
| `initial_credits_expiry_date_utc`                                                                                                                                                                                                                     | *Optional[str]*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | initial_credits_expiry_date_utc is the expiry date in UTC timezone (optional to set nil means no expiry)<br/>ex 2025-01-01 00:00:00 UTC                                                                                                               |
| `initial_credits_to_load`                                                                                                                                                                                                                             | *Optional[str]*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | initial_credits_to_load is the number of credits to load to the wallet<br/>if not provided, the wallet will be created with 0 balance<br/>NOTE: this is not the amount in the currency, but the number of credits                                     |
| `initial_credits_to_load_expiry_date`                                                                                                                                                                                                                 | *Optional[int]*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | initial_credits_to_load_expiry_date YYYYMMDD format in UTC timezone (optional to set nil means no expiry)<br/>for ex 20250101 means the credits will expire on 2025-01-01 00:00:00 UTC<br/>hence they will be available for use until 2024-12-31 23:59:59 UTC |
| `metadata`                                                                                                                                                                                                                                            | Dict[str, *str*]                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `name`                                                                                                                                                                                                                                                | *Optional[str]*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `wallet_type`                                                                                                                                                                                                                                         | [Optional[models.TypesWalletType]](../../models/typeswallettype.md)                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                    | N/A                                                                                                                                                                                                                                                   |
| `retries`                                                                                                                                                                                                                                             | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                    | Configuration to override the default retry behavior of the client.                                                                                                                                                                                   |

### Response

**[models.DtoWalletResponse](../../models/dtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get

Get a wallet by its ID

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/wallets/{id}" method="get" path="/wallets/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.wallets.get(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Wallet ID                                                           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoWalletResponse](../../models/dtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## update

Update a wallet's details including auto top-up configuration

### Example Usage

<!-- UsageSnippet language="python" operationID="put_/wallets/{id}" method="put" path="/wallets/{id}" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.wallets.update(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `id`                                                                            | *str*                                                                           | :heavy_check_mark:                                                              | Wallet ID                                                                       |
| `alert_config`                                                                  | [Optional[models.DtoAlertConfig]](../../models/dtoalertconfig.md)               | :heavy_minus_sign:                                                              | N/A                                                                             |
| `alert_enabled`                                                                 | *Optional[bool]*                                                                | :heavy_minus_sign:                                                              | N/A                                                                             |
| `auto_topup_amount`                                                             | *Optional[str]*                                                                 | :heavy_minus_sign:                                                              | N/A                                                                             |
| `auto_topup_min_balance`                                                        | *Optional[str]*                                                                 | :heavy_minus_sign:                                                              | N/A                                                                             |
| `auto_topup_trigger`                                                            | [Optional[models.TypesAutoTopupTrigger]](../../models/typesautotopuptrigger.md) | :heavy_minus_sign:                                                              | N/A                                                                             |
| `config`                                                                        | [Optional[models.TypesWalletConfig]](../../models/typeswalletconfig.md)         | :heavy_minus_sign:                                                              | N/A                                                                             |
| `description`                                                                   | *Optional[str]*                                                                 | :heavy_minus_sign:                                                              | N/A                                                                             |
| `metadata`                                                                      | Dict[str, *str*]                                                                | :heavy_minus_sign:                                                              | N/A                                                                             |
| `name`                                                                          | *Optional[str]*                                                                 | :heavy_minus_sign:                                                              | N/A                                                                             |
| `retries`                                                                       | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                | :heavy_minus_sign:                                                              | Configuration to override the default retry behavior of the client.             |

### Response

**[models.DtoWalletResponse](../../models/dtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_balance_real_time

Get real-time balance of a wallet

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/wallets/{id}/balance/real-time" method="get" path="/wallets/{id}/balance/real-time" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.wallets.get_balance_real_time(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Wallet ID                                                           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoWalletBalanceResponse](../../models/dtowalletbalanceresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## debit

Debit a wallet by debiting credits from a wallet

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/wallets/{id}/debit" method="post" path="/wallets/{id}/debit" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.wallets.debit(id="<id>", idempotency_key="<value>", transaction_reason="SUBSCRIPTION_CREDIT_GRANT")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `id`                                                                    | *str*                                                                   | :heavy_check_mark:                                                      | Wallet ID                                                               |
| `idempotency_key`                                                       | *str*                                                                   | :heavy_check_mark:                                                      | idempotency_key is a unique key for the transaction                     |
| `transaction_reason`                                                    | [models.TypesTransactionReason](../../models/typestransactionreason.md) | :heavy_check_mark:                                                      | N/A                                                                     |
| `credits`                                                               | *Optional[str]*                                                         | :heavy_minus_sign:                                                      | credits is the number of credits to debit from the wallet               |
| `description`                                                           | *Optional[str]*                                                         | :heavy_minus_sign:                                                      | description to add any specific details about the transaction           |
| `metadata`                                                              | Dict[str, *str*]                                                        | :heavy_minus_sign:                                                      | N/A                                                                     |
| `retries`                                                               | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)        | :heavy_minus_sign:                                                      | Configuration to override the default retry behavior of the client.     |

### Response

**[models.DtoWalletResponse](../../models/dtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## terminate

Terminates a wallet by closing it and debiting remaining balance

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/wallets/{id}/terminate" method="post" path="/wallets/{id}/terminate" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.wallets.terminate(id="<id>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *str*                                                               | :heavy_check_mark:                                                  | Wallet ID                                                           |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.DtoWalletResponse](../../models/dtowalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## top_up

Add credits to a wallet

### Example Usage

<!-- UsageSnippet language="python" operationID="post_/wallets/{id}/top-up" method="post" path="/wallets/{id}/top-up" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.wallets.top_up(id="<id>", transaction_reason="FREE_CREDIT_GRANT")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `id`                                                                                                                                                                                                                                                                                                                                                                                     | *str*                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | Wallet ID                                                                                                                                                                                                                                                                                                                                                                                |
| `transaction_reason`                                                                                                                                                                                                                                                                                                                                                                     | [models.TypesTransactionReason](../../models/typestransactionreason.md)                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                                                                                                                                      |
| `amount`                                                                                                                                                                                                                                                                                                                                                                                 | *Optional[str]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | amount is the amount in the currency of the wallet to be added<br/>NOTE: this is not the number of credits to add, but the amount in the currency<br/>amount = credits_to_add * conversion_rate<br/>if both amount and credits_to_add are provided, amount will be ignored<br/>ex if the wallet has a conversion_rate of 2 then adding an amount of<br/>10 USD in the wallet wil add 5 credits in the wallet |
| `credits_to_add`                                                                                                                                                                                                                                                                                                                                                                         | *Optional[str]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | credits_to_add is the number of credits to add to the wallet                                                                                                                                                                                                                                                                                                                             |
| `description`                                                                                                                                                                                                                                                                                                                                                                            | *Optional[str]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | description to add any specific details about the transaction                                                                                                                                                                                                                                                                                                                            |
| `expiry_date_utc`                                                                                                                                                                                                                                                                                                                                                                        | *Optional[str]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | expiry_date_utc is the expiry date in UTC timezone<br/>ex 2025-01-01 00:00:00 UTC                                                                                                                                                                                                                                                                                                        |
| `idempotency_key`                                                                                                                                                                                                                                                                                                                                                                        | *Optional[str]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | idempotency_key is a unique key for the transaction                                                                                                                                                                                                                                                                                                                                      |
| `metadata`                                                                                                                                                                                                                                                                                                                                                                               | Dict[str, *str*]                                                                                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                                                                                                                                      |
| `priority`                                                                                                                                                                                                                                                                                                                                                                               | *Optional[int]*                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | priority is the priority of the transaction<br/>lower number means higher priority<br/>default is nil which means no priority at all                                                                                                                                                                                                                                                     |
| `retries`                                                                                                                                                                                                                                                                                                                                                                                | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | Configuration to override the default retry behavior of the client.                                                                                                                                                                                                                                                                                                                      |

### Response

**[models.DtoTopUpWalletResponse](../../models/dtotopupwalletresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## get_transactions

Get transactions for a wallet with pagination

### Example Usage

<!-- UsageSnippet language="python" operationID="get_/wallets/{id}/transactions" method="get" path="/wallets/{id}/transactions" -->
```python
from flexprice import Flexprice


with Flexprice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as f_client:

    res = f_client.wallets.get_transactions(id_path_parameter="<value>")

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                                                                             | Type                                                                                                                  | Required                                                                                                              | Description                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `id_path_parameter`                                                                                                   | *str*                                                                                                                 | :heavy_check_mark:                                                                                                    | Wallet ID                                                                                                             |
| `credits_available_gt`                                                                                                | *Optional[float]*                                                                                                     | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `end_time`                                                                                                            | *Optional[str]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `expand`                                                                                                              | *Optional[str]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `expiry_date_after`                                                                                                   | *Optional[str]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `expiry_date_before`                                                                                                  | *Optional[str]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `id_query_parameter`                                                                                                  | *Optional[str]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `limit`                                                                                                               | *Optional[int]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `offset`                                                                                                              | *Optional[int]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `order`                                                                                                               | [Optional[models.GetWalletsIDTransactionsQueryParamOrder]](../../models/getwalletsidtransactionsqueryparamorder.md)   | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `priority`                                                                                                            | *Optional[int]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `reference_id`                                                                                                        | *Optional[str]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `reference_type`                                                                                                      | *Optional[str]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `sort`                                                                                                                | *Optional[str]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `start_time`                                                                                                          | *Optional[str]*                                                                                                       | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `status`                                                                                                              | [Optional[models.GetWalletsIDTransactionsQueryParamStatus]](../../models/getwalletsidtransactionsqueryparamstatus.md) | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `transaction_reason`                                                                                                  | [Optional[models.TransactionReason]](../../models/transactionreason.md)                                               | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `transaction_status`                                                                                                  | [Optional[models.TransactionStatus]](../../models/transactionstatus.md)                                               | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `type`                                                                                                                | [Optional[models.Type]](../../models/type.md)                                                                         | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `retries`                                                                                                             | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)                                                      | :heavy_minus_sign:                                                                                                    | Configuration to override the default retry behavior of the client.                                                   |

### Response

**[models.DtoListWalletTransactionsResponse](../../models/dtolistwallettransactionsresponse.md)**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexpriceDefaultError | 4XX, 5XX                     | \*/\*                        |