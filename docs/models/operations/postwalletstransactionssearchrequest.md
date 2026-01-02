# PostWalletsTransactionsSearchRequest

## Example Usage

```typescript
import { PostWalletsTransactionsSearchRequest } from "@flexprice/sdk/models/operations";

let value: PostWalletsTransactionsSearchRequest = {};
```

## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `expand`                                                                                           | *string*                                                                                           | :heavy_minus_sign:                                                                                 | Expand fields (e.g., customer,created_by_user,wallet)                                              |
| `body`                                                                                             | [components.TypesWalletTransactionFilter](../../models/components/typeswallettransactionfilter.md) | :heavy_minus_sign:                                                                                 | Filter                                                                                             |