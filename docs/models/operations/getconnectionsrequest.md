# GetConnectionsRequest

## Example Usage

```typescript
import { GetConnectionsRequest } from "@flexprice/sdk/models/operations";

let value: GetConnectionsRequest = {};
```

## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `connectionIds`                                                            | *string*[]                                                                 | :heavy_minus_sign:                                                         | N/A                                                                        |
| `endTime`                                                                  | *string*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `expand`                                                                   | *string*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `limit`                                                                    | *number*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `offset`                                                                   | *number*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `order`                                                                    | [operations.QueryParamOrder](../../models/operations/queryparamorder.md)   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `providerType`                                                             | [operations.ProviderType](../../models/operations/providertype.md)         | :heavy_minus_sign:                                                         | N/A                                                                        |
| `startTime`                                                                | *string*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `status`                                                                   | [operations.QueryParamStatus](../../models/operations/queryparamstatus.md) | :heavy_minus_sign:                                                         | N/A                                                                        |