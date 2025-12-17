# DtoTopUpWalletResponse

## Example Usage

```typescript
import { DtoTopUpWalletResponse } from "@flexprice/sdk/models";

let value: DtoTopUpWalletResponse = {};
```

## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `invoiceId`                                                                      | *string*                                                                         | :heavy_minus_sign:                                                               | Invoice ID if an invoice was created (only for PURCHASED_CREDIT_INVOICED)        |
| `wallet`                                                                         | [models.DtoWalletResponse](../models/dtowalletresponse.md)                       | :heavy_minus_sign:                                                               | N/A                                                                              |
| `walletTransaction`                                                              | [models.DtoWalletTransactionResponse](../models/dtowallettransactionresponse.md) | :heavy_minus_sign:                                                               | N/A                                                                              |