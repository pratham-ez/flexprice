# DtoCancelSubscriptionResponse

## Example Usage

```typescript
import { DtoCancelSubscriptionResponse } from "@flexprice/sdk/models";

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

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `cancellationType`                                                     | [models.TypesCancellationType](../models/typescancellationtype.md)     | :heavy_minus_sign:                                                     | N/A                                                                    |
| `effectiveDate`                                                        | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `message`                                                              | *string*                                                               | :heavy_minus_sign:                                                     | Response metadata                                                      |
| `processedAt`                                                          | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `prorationDetails`                                                     | [models.DtoProrationDetail](../models/dtoprorationdetail.md)[]         | :heavy_minus_sign:                                                     | N/A                                                                    |
| `prorationInvoice`                                                     | [models.DtoInvoiceResponse](../models/dtoinvoiceresponse.md)           | :heavy_minus_sign:                                                     | N/A                                                                    |
| `reason`                                                               | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `status`                                                               | [models.TypesSubscriptionStatus](../models/typessubscriptionstatus.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `subscriptionId`                                                       | *string*                                                               | :heavy_minus_sign:                                                     | Basic cancellation info                                                |
| `totalCreditAmount`                                                    | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |