# DtoManualBalanceDebitRequest


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `credits`                                                            | *Optional[str]*                                                      | :heavy_minus_sign:                                                   | credits is the number of credits to debit from the wallet            |
| `description`                                                        | *Optional[str]*                                                      | :heavy_minus_sign:                                                   | description to add any specific details about the transaction        |
| `idempotency_key`                                                    | *str*                                                                | :heavy_check_mark:                                                   | idempotency_key is a unique key for the transaction                  |
| `metadata`                                                           | Dict[str, *str*]                                                     | :heavy_minus_sign:                                                   | N/A                                                                  |
| `transaction_reason`                                                 | [models.TypesTransactionReason](../models/typestransactionreason.md) | :heavy_check_mark:                                                   | N/A                                                                  |