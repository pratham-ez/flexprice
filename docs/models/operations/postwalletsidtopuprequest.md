# PostWalletsIdTopUpRequest

## Example Usage

```typescript
import { PostWalletsIdTopUpRequest } from "@flexprice/sdk/models/operations";

let value: PostWalletsIdTopUpRequest = {
  id: "<id>",
  body: {
    transactionReason: "MANUAL_BALANCE_DEBIT",
  },
};
```

## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `id`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | Wallet ID                                                             |
| `body`                                                                | [models.DtoTopUpWalletRequest](../../models/dtotopupwalletrequest.md) | :heavy_check_mark:                                                    | Top up request                                                        |