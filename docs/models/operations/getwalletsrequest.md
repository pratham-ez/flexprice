# GetWalletsRequest

## Example Usage

```typescript
import { GetWalletsRequest } from "@flexprice/sdk/models/operations";

let value: GetWalletsRequest = {};
```

## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `alertEnabled`                                                             | *boolean*                                                                  | :heavy_minus_sign:                                                         | N/A                                                                        |
| `expand`                                                                   | *string*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `limit`                                                                    | *number*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `offset`                                                                   | *number*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `order`                                                                    | [operations.GetWalletsOrder](../../models/operations/getwalletsorder.md)   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `sort`                                                                     | *string*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `status`                                                                   | [operations.GetWalletsStatus](../../models/operations/getwalletsstatus.md) | :heavy_minus_sign:                                                         | N/A                                                                        |
| `walletIds`                                                                | *string*[]                                                                 | :heavy_minus_sign:                                                         | N/A                                                                        |