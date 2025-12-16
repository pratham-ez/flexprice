# PutWalletsIdRequest

## Example Usage

```typescript
import { PutWalletsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutWalletsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `id`                                                                    | *string*                                                                | :heavy_check_mark:                                                      | Wallet ID                                                               |
| `body`                                                                  | [models.DtoUpdateWalletRequest](../../models/dtoupdatewalletrequest.md) | :heavy_check_mark:                                                      | Update wallet request                                                   |