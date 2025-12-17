# DtoManualBalanceDebitRequest

## Example Usage

```typescript
import { DtoManualBalanceDebitRequest } from "@flexprice/sdk/models";

let value: DtoManualBalanceDebitRequest = {
  idempotencyKey: "<value>",
  transactionReason: "FREE_CREDIT_GRANT",
};
```

## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `credits`                                                            | *string*                                                             | :heavy_minus_sign:                                                   | credits is the number of credits to debit from the wallet            |
| `description`                                                        | *string*                                                             | :heavy_minus_sign:                                                   | description to add any specific details about the transaction        |
| `idempotencyKey`                                                     | *string*                                                             | :heavy_check_mark:                                                   | idempotency_key is a unique key for the transaction                  |
| `metadata`                                                           | Record<string, *string*>                                             | :heavy_minus_sign:                                                   | N/A                                                                  |
| `transactionReason`                                                  | [models.TypesTransactionReason](../models/typestransactionreason.md) | :heavy_check_mark:                                                   | N/A                                                                  |