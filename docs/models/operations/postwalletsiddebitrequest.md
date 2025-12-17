# PostWalletsIdDebitRequest

## Example Usage

```typescript
import { PostWalletsIdDebitRequest } from "@flexprice/sdk/models/operations";

let value: PostWalletsIdDebitRequest = {
  id: "<id>",
  body: {
    idempotencyKey: "<value>",
    transactionReason: "SUBSCRIPTION_CREDIT_GRANT",
  },
};
```

## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `id`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | Wallet ID                                                                           |
| `body`                                                                              | [models.DtoManualBalanceDebitRequest](../../models/dtomanualbalancedebitrequest.md) | :heavy_check_mark:                                                                  | Debit wallet request                                                                |