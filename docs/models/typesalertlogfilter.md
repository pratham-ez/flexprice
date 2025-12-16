# TypesAlertLogFilter

## Example Usage

```typescript
import { TypesAlertLogFilter } from "@flexprice/sdk/models";

let value: TypesAlertLogFilter = {};
```

## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `alertStatus`                                                            | [models.TypesAlertState](../models/typesalertstate.md)                   | :heavy_minus_sign:                                                       | N/A                                                                      |
| `alertType`                                                              | [models.TypesAlertType](../models/typesalerttype.md)                     | :heavy_minus_sign:                                                       | N/A                                                                      |
| `customerId`                                                             | *string*                                                                 | :heavy_minus_sign:                                                       | N/A                                                                      |
| `endTime`                                                                | *string*                                                                 | :heavy_minus_sign:                                                       | N/A                                                                      |
| `entityId`                                                               | *string*                                                                 | :heavy_minus_sign:                                                       | N/A                                                                      |
| `entityType`                                                             | [models.TypesAlertEntityType](../models/typesalertentitytype.md)         | :heavy_minus_sign:                                                       | N/A                                                                      |
| `expand`                                                                 | *string*                                                                 | :heavy_minus_sign:                                                       | N/A                                                                      |
| `filters`                                                                | [models.TypesFilterCondition](../models/typesfiltercondition.md)[]       | :heavy_minus_sign:                                                       | filters allows complex filtering based on multiple fields                |
| `limit`                                                                  | *number*                                                                 | :heavy_minus_sign:                                                       | N/A                                                                      |
| `offset`                                                                 | *number*                                                                 | :heavy_minus_sign:                                                       | N/A                                                                      |
| `order`                                                                  | [models.TypesAlertLogFilterOrder](../models/typesalertlogfilterorder.md) | :heavy_minus_sign:                                                       | N/A                                                                      |
| `sort`                                                                   | [models.TypesSortCondition](../models/typessortcondition.md)[]           | :heavy_minus_sign:                                                       | N/A                                                                      |
| `startTime`                                                              | *string*                                                                 | :heavy_minus_sign:                                                       | N/A                                                                      |
| `status`                                                                 | [models.TypesStatus](../models/typesstatus.md)                           | :heavy_minus_sign:                                                       | N/A                                                                      |