# CostsheetFilter

## Example Usage

```typescript
import { CostsheetFilter } from "@flexprice/sdk/models";

let value: CostsheetFilter = {};
```

## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `costsheetIDs`                                                     | *string*[]                                                         | :heavy_minus_sign:                                                 | CostsheetIDs allows filtering by specific costsheet IDs            |
| `environmentID`                                                    | *string*                                                           | :heavy_minus_sign:                                                 | EnvironmentID filters by specific environment ID                   |
| `filters`                                                          | [models.TypesFilterCondition](../models/typesfiltercondition.md)[] | :heavy_minus_sign:                                                 | Filters contains custom filtering conditions                       |
| `lookupKey`                                                        | *string*                                                           | :heavy_minus_sign:                                                 | LookupKey filters by lookup key                                    |
| `name`                                                             | *string*                                                           | :heavy_minus_sign:                                                 | Name filters by costsheet name                                     |
| `queryFilter`                                                      | [models.TypesQueryFilter](../models/typesqueryfilter.md)           | :heavy_minus_sign:                                                 | N/A                                                                |
| `sort`                                                             | [models.TypesSortCondition](../models/typessortcondition.md)[]     | :heavy_minus_sign:                                                 | Sort specifies result ordering preferences                         |
| `status`                                                           | [models.TypesStatus](../models/typesstatus.md)                     | :heavy_minus_sign:                                                 | N/A                                                                |
| `tenantID`                                                         | *string*                                                           | :heavy_minus_sign:                                                 | TenantID filters by specific tenant ID                             |
| `timeRangeFilter`                                                  | [models.TypesTimeRangeFilter](../models/typestimerangefilter.md)   | :heavy_minus_sign:                                                 | N/A                                                                |