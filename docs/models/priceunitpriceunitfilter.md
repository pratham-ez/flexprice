# PriceunitPriceUnitFilter

## Example Usage

```typescript
import { PriceunitPriceUnitFilter } from "@flexprice/sdk/models";

let value: PriceunitPriceUnitFilter = {};
```

## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `environmentId`                                                    | *string*                                                           | :heavy_minus_sign:                                                 | EnvironmentID filters by specific environment ID                   |
| `filters`                                                          | [models.TypesFilterCondition](../models/typesfiltercondition.md)[] | :heavy_minus_sign:                                                 | Filters allows complex filtering based on multiple fields          |
| `queryFilter`                                                      | [models.TypesQueryFilter](../models/typesqueryfilter.md)           | :heavy_minus_sign:                                                 | N/A                                                                |
| `sort`                                                             | [models.TypesSortCondition](../models/typessortcondition.md)[]     | :heavy_minus_sign:                                                 | Sort allows sorting by multiple fields                             |
| `status`                                                           | [models.TypesStatus](../models/typesstatus.md)                     | :heavy_minus_sign:                                                 | N/A                                                                |
| `tenantId`                                                         | *string*                                                           | :heavy_minus_sign:                                                 | TenantID filters by specific tenant ID                             |
| `timeRangeFilter`                                                  | [models.TypesTimeRangeFilter](../models/typestimerangefilter.md)   | :heavy_minus_sign:                                                 | N/A                                                                |