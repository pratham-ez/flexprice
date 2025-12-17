# TypesFeatureFilter

## Example Usage

```typescript
import { TypesFeatureFilter } from "@flexprice/sdk/models";

let value: TypesFeatureFilter = {};
```

## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `endTime`                                                              | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `expand`                                                               | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `featureIds`                                                           | *string*[]                                                             | :heavy_minus_sign:                                                     | Feature specific filters                                               |
| `filters`                                                              | [models.TypesFilterCondition](../models/typesfiltercondition.md)[]     | :heavy_minus_sign:                                                     | filters allows complex filtering based on multiple fields              |
| `limit`                                                                | *number*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `lookupKey`                                                            | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `lookupKeys`                                                           | *string*[]                                                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `meterIds`                                                             | *string*[]                                                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `nameContains`                                                         | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `offset`                                                               | *number*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `order`                                                                | [models.TypesFeatureFilterOrder](../models/typesfeaturefilterorder.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `sort`                                                                 | [models.TypesSortCondition](../models/typessortcondition.md)[]         | :heavy_minus_sign:                                                     | N/A                                                                    |
| `startTime`                                                            | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `status`                                                               | [models.TypesStatus](../models/typesstatus.md)                         | :heavy_minus_sign:                                                     | N/A                                                                    |