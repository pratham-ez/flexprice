# DtoListCreditNotesResponse

## Example Usage

```typescript
import { DtoListCreditNotesResponse } from "@flexprice/sdk/models";

let value: DtoListCreditNotesResponse = {
  items: [
    {
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
| `items`                                                                | [models.DtoCreditNoteResponse](../models/dtocreditnoteresponse.md)[]   | :heavy_minus_sign:                                                     | N/A                                                                    |
| `pagination`                                                           | [models.TypesPaginationResponse](../models/typespaginationresponse.md) | :heavy_minus_sign:                                                     | N/A                                                                    |