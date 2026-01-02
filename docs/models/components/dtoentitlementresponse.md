# DtoEntitlementResponse

## Example Usage

```typescript
import { DtoEntitlementResponse } from "@flexprice/sdk/models/components";

let value: DtoEntitlementResponse = {
  addon: {
    prices: [
      {
        meter: {
          createdAt: "2024-03-20T15:04:05Z",
          eventName: "api_request",
          id: "550e8400-e29b-41d4-a716-446655440000",
          name: "API Usage Meter",
          status: "published",
          tenantId: "tenant123",
          updatedAt: "2024-03-20T15:04:05Z",
        },
        plan: {},
      },
    ],
  },
  feature: {
    meter: {
      createdAt: "2024-03-20T15:04:05Z",
      eventName: "api_request",
      id: "550e8400-e29b-41d4-a716-446655440000",
      name: "API Usage Meter",
      status: "published",
      tenantId: "tenant123",
      updatedAt: "2024-03-20T15:04:05Z",
    },
  },
  plan: {},
};
```

## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `addon`                                                                                                    | [components.DtoAddonResponse](../../models/components/dtoaddonresponse.md)                                 | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `createdAt`                                                                                                | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `createdBy`                                                                                                | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `displayOrder`                                                                                             | *number*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `endDate`                                                                                                  | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `entityId`                                                                                                 | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `entityType`                                                                                               | [components.TypesEntitlementEntityType](../../models/components/typesentitlemententitytype.md)             | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `environmentId`                                                                                            | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `feature`                                                                                                  | [components.DtoFeatureResponse](../../models/components/dtofeatureresponse.md)                             | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `featureId`                                                                                                | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `featureType`                                                                                              | [components.TypesFeatureType](../../models/components/typesfeaturetype.md)                                 | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `id`                                                                                                       | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `isEnabled`                                                                                                | *boolean*                                                                                                  | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `isSoftLimit`                                                                                              | *boolean*                                                                                                  | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `parentEntitlementId`                                                                                      | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `plan`                                                                                                     | [components.DtoPlanResponse](../../models/components/dtoplanresponse.md)                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `planId`                                                                                                   | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | TODO: Remove this once we have a proper entitlement entity type                                            |
| `startDate`                                                                                                | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `staticValue`                                                                                              | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `status`                                                                                                   | [components.TypesStatus](../../models/components/typesstatus.md)                                           | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `tenantId`                                                                                                 | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `updatedAt`                                                                                                | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `updatedBy`                                                                                                | *string*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `usageLimit`                                                                                               | *number*                                                                                                   | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |
| `usageResetPeriod`                                                                                         | [components.TypesEntitlementUsageResetPeriod](../../models/components/typesentitlementusageresetperiod.md) | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |