# DtoAddonResponse

## Example Usage

```typescript
import { DtoAddonResponse } from "@flexprice/sdk/models/components";

let value: DtoAddonResponse = {
  entitlements: [
    {
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
    },
  ],
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
};
```

## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `createdAt`                                                                              | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `createdBy`                                                                              | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `description`                                                                            | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `entitlements`                                                                           | [components.DtoEntitlementResponse](../../models/components/dtoentitlementresponse.md)[] | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `environmentId`                                                                          | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `id`                                                                                     | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `lookupKey`                                                                              | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `metadata`                                                                               | Record<string, *any*>                                                                    | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `name`                                                                                   | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `prices`                                                                                 | [components.DtoPriceResponse](../../models/components/dtopriceresponse.md)[]             | :heavy_minus_sign:                                                                       | Optional expanded fields                                                                 |
| `status`                                                                                 | [components.TypesStatus](../../models/components/typesstatus.md)                         | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `tenantId`                                                                               | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `type`                                                                                   | [components.TypesAddonType](../../models/components/typesaddontype.md)                   | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `updatedAt`                                                                              | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `updatedBy`                                                                              | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |