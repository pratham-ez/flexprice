# GetPricesUnitsRequest

## Example Usage

```typescript
import { GetPricesUnitsRequest } from "@flexprice/sdk/models/operations";

let value: GetPricesUnitsRequest = {};
```

## Fields

| Field                   | Type                    | Required                | Description             |
| ----------------------- | ----------------------- | ----------------------- | ----------------------- |
| `status`                | *string*                | :heavy_minus_sign:      | Filter by status        |
| `limit`                 | *number*                | :heavy_minus_sign:      | Limit number of results |
| `offset`                | *number*                | :heavy_minus_sign:      | Offset for pagination   |
| `sort`                  | *string*                | :heavy_minus_sign:      | Sort field              |
| `order`                 | *string*                | :heavy_minus_sign:      | Sort order (asc/desc)   |