# DtoListSubscriptionsResponse

## Example Usage

```typescript
import { DtoListSubscriptionsResponse } from "@flexprice/sdk/models";

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

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `items`                                                                  | [models.DtoSubscriptionResponse](../models/dtosubscriptionresponse.md)[] | :heavy_minus_sign:                                                       | N/A                                                                      |
| `pagination`                                                             | [models.TypesPaginationResponse](../models/typespaginationresponse.md)   | :heavy_minus_sign:                                                       | N/A                                                                      |