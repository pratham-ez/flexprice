# DtoTopUpWalletResponse

## Example Usage

```typescript
import { DtoTopUpWalletResponse } from "@flexprice/sdk/models/components";

let value: DtoTopUpWalletResponse = {};
```

## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `invoiceId`                                                                                        | *string*                                                                                           | :heavy_minus_sign:                                                                                 | Invoice ID if an invoice was created (only for PURCHASED_CREDIT_INVOICED)                          |
| `wallet`                                                                                           | [components.DtoWalletResponse](../../models/components/dtowalletresponse.md)                       | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `walletTransaction`                                                                                | [components.DtoWalletTransactionResponse](../../models/components/dtowallettransactionresponse.md) | :heavy_minus_sign:                                                                                 | N/A                                                                                                |