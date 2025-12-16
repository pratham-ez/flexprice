# TypesUserFilter

## Example Usage

```typescript
import { TypesUserFilter } from "@flexprice/sdk/models";

let value: TypesUserFilter = {};
```

## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `endTime`                                                          | *string*                                                           | :heavy_minus_sign:                                                 | N/A                                                                |
| `expand`                                                           | *string*                                                           | :heavy_minus_sign:                                                 | N/A                                                                |
| `filters`                                                          | [models.TypesFilterCondition](../models/typesfiltercondition.md)[] | :heavy_minus_sign:                                                 | filters allows complex filtering based on multiple fields          |
| `limit`                                                            | *number*                                                           | :heavy_minus_sign:                                                 | N/A                                                                |
| `offset`                                                           | *number*                                                           | :heavy_minus_sign:                                                 | N/A                                                                |
| `order`                                                            | [models.TypesUserFilterOrder](../models/typesuserfilterorder.md)   | :heavy_minus_sign:                                                 | N/A                                                                |
| `roles`                                                            | *string*[]                                                         | :heavy_minus_sign:                                                 | N/A                                                                |
| `sort`                                                             | [models.TypesSortCondition](../models/typessortcondition.md)[]     | :heavy_minus_sign:                                                 | N/A                                                                |
| `startTime`                                                        | *string*                                                           | :heavy_minus_sign:                                                 | N/A                                                                |
| `status`                                                           | [models.TypesStatus](../models/typesstatus.md)                     | :heavy_minus_sign:                                                 | N/A                                                                |
| `type`                                                             | [models.TypesUserType](../models/typesusertype.md)                 | :heavy_minus_sign:                                                 | N/A                                                                |
| `userIds`                                                          | *string*[]                                                         | :heavy_minus_sign:                                                 | Specific filters for users                                         |