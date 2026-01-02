# DtoCancelSubscriptionResponse

## Example Usage

```typescript
import { DtoCancelSubscriptionResponse } from "@flexprice/sdk/models/components";

let value: DtoCancelSubscriptionResponse = {
  prorationInvoice: {
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

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `cancellationType`                                                                       | [components.TypesCancellationType](../../models/components/typescancellationtype.md)     | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `effectiveDate`                                                                          | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `message`                                                                                | *string*                                                                                 | :heavy_minus_sign:                                                                       | Response metadata                                                                        |
| `processedAt`                                                                            | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `prorationDetails`                                                                       | [components.DtoProrationDetail](../../models/components/dtoprorationdetail.md)[]         | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `prorationInvoice`                                                                       | [components.DtoInvoiceResponse](../../models/components/dtoinvoiceresponse.md)           | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `reason`                                                                                 | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `status`                                                                                 | [components.TypesSubscriptionStatus](../../models/components/typessubscriptionstatus.md) | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `subscriptionId`                                                                         | *string*                                                                                 | :heavy_minus_sign:                                                                       | Basic cancellation info                                                                  |
| `totalCreditAmount`                                                                      | *string*                                                                                 | :heavy_minus_sign:                                                                       | N/A                                                                                      |