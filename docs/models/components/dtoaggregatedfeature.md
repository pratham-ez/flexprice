# DtoAggregatedFeature

## Example Usage

```typescript
import { DtoAggregatedFeature } from "@flexprice/sdk/models/components";

let value: DtoAggregatedFeature = {
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
};
```

## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `entitlement`                                                                              | [components.DtoAggregatedEntitlement](../../models/components/dtoaggregatedentitlement.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `feature`                                                                                  | [components.DtoFeatureResponse](../../models/components/dtofeatureresponse.md)             | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `sources`                                                                                  | [components.DtoEntitlementSource](../../models/components/dtoentitlementsource.md)[]       | :heavy_minus_sign:                                                                         | N/A                                                                                        |