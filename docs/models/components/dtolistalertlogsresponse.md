# DtoListAlertLogsResponse

## Example Usage

```typescript
import { DtoListAlertLogsResponse } from "@flexprice/sdk/models/components";

let value: DtoListAlertLogsResponse = {
  items: [
    {
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
    },
  ],
};
```

## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `items`                                                                                  | [components.DtoAlertLogResponse](../../models/components/dtoalertlogresponse.md)[]       | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `pagination`                                                                             | [components.TypesPaginationResponse](../../models/components/typespaginationresponse.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |