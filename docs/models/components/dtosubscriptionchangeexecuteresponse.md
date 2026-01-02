# DtoSubscriptionChangeExecuteResponse

Response after successfully executing a subscription plan change

## Example Usage

```typescript
import { DtoSubscriptionChangeExecuteResponse } from "@flexprice/sdk/models/components";

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

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `changeType`                                                                                     | [components.TypesSubscriptionChangeType](../../models/components/typessubscriptionchangetype.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `creditGrants`                                                                                   | [components.DtoCreditGrantResponse](../../models/components/dtocreditgrantresponse.md)[]         | :heavy_minus_sign:                                                                               | credit_grants contains any credit grants created for proration credits                           |
| `effectiveDate`                                                                                  | *string*                                                                                         | :heavy_minus_sign:                                                                               | effective_date is when the change took effect                                                    |
| `invoice`                                                                                        | [components.DtoInvoiceResponse](../../models/components/dtoinvoiceresponse.md)                   | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `metadata`                                                                                       | Record<string, *string*>                                                                         | :heavy_minus_sign:                                                                               | metadata from the request                                                                        |
| `newSubscription`                                                                                | [components.DtoSubscriptionSummary](../../models/components/dtosubscriptionsummary.md)           | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `oldSubscription`                                                                                | [components.DtoSubscriptionSummary](../../models/components/dtosubscriptionsummary.md)           | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `prorationApplied`                                                                               | [components.DtoProrationDetails](../../models/components/dtoprorationdetails.md)                 | :heavy_minus_sign:                                                                               | N/A                                                                                              |