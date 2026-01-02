# DtoGetUsageAnalyticsResponse

## Example Usage

```typescript
import { DtoGetUsageAnalyticsResponse } from "@flexprice/sdk/models/components";

let value: DtoGetUsageAnalyticsResponse = {
  items: [
    {
      price: {
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
    },
  ],
};
```

## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `currency`                                                                           | *string*                                                                             | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `items`                                                                              | [components.DtoUsageAnalyticItem](../../models/components/dtousageanalyticitem.md)[] | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `totalCost`                                                                          | *string*                                                                             | :heavy_minus_sign:                                                                   | N/A                                                                                  |