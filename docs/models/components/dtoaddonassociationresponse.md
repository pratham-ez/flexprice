# DtoAddonAssociationResponse

## Example Usage

```typescript
import { DtoAddonAssociationResponse } from "@flexprice/sdk/models/components";

let value: DtoAddonAssociationResponse = {
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
  subscription: {
    latestInvoice: {
      subscription: {
        plan: {},
      },
    },
    plan: {},
  },
};
```

## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `addon`                                                                                                  | [components.DtoAddonResponse](../../models/components/dtoaddonresponse.md)                               | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `addonId`                                                                                                | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `addonStatus`                                                                                            | [components.TypesAddonStatus](../../models/components/typesaddonstatus.md)                               | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `cancellationReason`                                                                                     | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `cancelledAt`                                                                                            | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `createdAt`                                                                                              | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `createdBy`                                                                                              | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `endDate`                                                                                                | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `entityId`                                                                                               | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `entityType`                                                                                             | [components.TypesAddonAssociationEntityType](../../models/components/typesaddonassociationentitytype.md) | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `environmentId`                                                                                          | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `id`                                                                                                     | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `metadata`                                                                                               | Record<string, *any*>                                                                                    | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `startDate`                                                                                              | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `status`                                                                                                 | [components.TypesStatus](../../models/components/typesstatus.md)                                         | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `subscription`                                                                                           | [components.DtoSubscriptionResponse](../../models/components/dtosubscriptionresponse.md)                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `tenantId`                                                                                               | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `updatedAt`                                                                                              | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `updatedBy`                                                                                              | *string*                                                                                                 | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |