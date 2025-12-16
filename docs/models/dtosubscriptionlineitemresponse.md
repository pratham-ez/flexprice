# DtoSubscriptionLineItemResponse

## Example Usage

```typescript
import { DtoSubscriptionLineItemResponse } from "@flexprice/sdk/models";

let value: DtoSubscriptionLineItemResponse = {
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
};
```

## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `billingPeriod`                                                                                | [models.TypesBillingPeriod](../models/typesbillingperiod.md)                                   | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `createdAt`                                                                                    | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `createdBy`                                                                                    | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `currency`                                                                                     | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `customerId`                                                                                   | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `displayName`                                                                                  | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `endDate`                                                                                      | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `entityId`                                                                                     | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `entityType`                                                                                   | [models.TypesSubscriptionLineItemEntityType](../models/typessubscriptionlineitementitytype.md) | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `environmentId`                                                                                | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `id`                                                                                           | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `invoiceCadence`                                                                               | [models.TypesInvoiceCadence](../models/typesinvoicecadence.md)                                 | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `metadata`                                                                                     | Record<string, *string*>                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `meterDisplayName`                                                                             | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `meterId`                                                                                      | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `planDisplayName`                                                                              | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `price`                                                                                        | [models.DtoPriceResponse](../models/dtopriceresponse.md)                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `priceId`                                                                                      | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `priceType`                                                                                    | [models.TypesPriceType](../models/typespricetype.md)                                           | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `priceUnit`                                                                                    | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `priceUnitId`                                                                                  | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `quantity`                                                                                     | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `startDate`                                                                                    | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `status`                                                                                       | [models.TypesStatus](../models/typesstatus.md)                                                 | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `subscriptionId`                                                                               | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `subscriptionPhaseId`                                                                          | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `tenantId`                                                                                     | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `trialPeriod`                                                                                  | *number*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `updatedAt`                                                                                    | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `updatedBy`                                                                                    | *string*                                                                                       | :heavy_minus_sign:                                                                             | N/A                                                                                            |