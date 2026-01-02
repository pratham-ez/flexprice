# DtoListSubscriptionsResponse

## Example Usage

```typescript
import { DtoListSubscriptionsResponse } from "@flexprice/sdk/models/components";

let value: DtoListSubscriptionsResponse = {
  items: [
    {
      latestInvoice: {
        subscription: {
          plan: {},
        },
      },
      plan: {},
    },
  ],
};
```

## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `items`                                                                                    | [components.DtoSubscriptionResponse](../../models/components/dtosubscriptionresponse.md)[] | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `pagination`                                                                               | [components.TypesPaginationResponse](../../models/components/typespaginationresponse.md)   | :heavy_minus_sign:                                                                         | N/A                                                                                        |