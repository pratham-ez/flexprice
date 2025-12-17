# DtoListInvoicesResponse

## Example Usage

```typescript
import { DtoListInvoicesResponse } from "@flexprice/sdk/models";

let value: DtoListInvoicesResponse = {
  items: [
    {
      subscription: {
        latestInvoice: {
          subscription: {
            plan: {},
          },
        },
        plan: {},
      },
    },
  ],
};
```

## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `items`                                                                | [models.DtoInvoiceResponse](../models/dtoinvoiceresponse.md)[]         | :heavy_minus_sign:                                                     | N/A                                                                    |
| `pagination`                                                           | [models.TypesPaginationResponse](../models/typespaginationresponse.md) | :heavy_minus_sign:                                                     | N/A                                                                    |