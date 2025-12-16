# DtoSubscriptionChangeExecuteResponse

Response after successfully executing a subscription plan change

## Example Usage

```typescript
import { DtoSubscriptionChangeExecuteResponse } from "@flexprice/sdk/models";

let value: DtoSubscriptionChangeExecuteResponse = {
  invoice: {
    subscription: {
      latestInvoice: {
        subscription: {
          plan: {},
        },
      },
      plan: {},
    },
  },
};
```

## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `changeType`                                                                   | [models.TypesSubscriptionChangeType](../models/typessubscriptionchangetype.md) | :heavy_minus_sign:                                                             | N/A                                                                            |
| `creditGrants`                                                                 | [models.DtoCreditGrantResponse](../models/dtocreditgrantresponse.md)[]         | :heavy_minus_sign:                                                             | credit_grants contains any credit grants created for proration credits         |
| `effectiveDate`                                                                | *string*                                                                       | :heavy_minus_sign:                                                             | effective_date is when the change took effect                                  |
| `invoice`                                                                      | [models.DtoInvoiceResponse](../models/dtoinvoiceresponse.md)                   | :heavy_minus_sign:                                                             | N/A                                                                            |
| `metadata`                                                                     | Record<string, *string*>                                                       | :heavy_minus_sign:                                                             | metadata from the request                                                      |
| `newSubscription`                                                              | [models.DtoSubscriptionSummary](../models/dtosubscriptionsummary.md)           | :heavy_minus_sign:                                                             | N/A                                                                            |
| `oldSubscription`                                                              | [models.DtoSubscriptionSummary](../models/dtosubscriptionsummary.md)           | :heavy_minus_sign:                                                             | N/A                                                                            |
| `prorationApplied`                                                             | [models.DtoProrationDetails](../models/dtoprorationdetails.md)                 | :heavy_minus_sign:                                                             | N/A                                                                            |