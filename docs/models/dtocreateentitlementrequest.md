# DtoCreateEntitlementRequest

## Example Usage

```typescript
import { DtoCreateEntitlementRequest } from "@flexprice/sdk/models";

let value: DtoCreateEntitlementRequest = {
  featureId: "<id>",
  featureType: "metered",
};
```

## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `entityId`                                                                               | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `entityType`                                                                             | [models.TypesEntitlementEntityType](../models/typesentitlemententitytype.md)             | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `featureId`                                                                              | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `featureType`                                                                            | [models.TypesFeatureType](../models/typesfeaturetype.md)                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `isEnabled`                                                                              | *boolean*                                                                                | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `isSoftLimit`                                                                            | *boolean*                                                                                | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `parentEntitlementId`                                                                    | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `planId`                                                                                 | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `staticValue`                                                                            | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `usageLimit`                                                                             | *number*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `usageResetPeriod`                                                                       | [models.TypesEntitlementUsageResetPeriod](../models/typesentitlementusageresetperiod.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |