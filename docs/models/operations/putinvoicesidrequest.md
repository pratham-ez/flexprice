# PutInvoicesIdRequest

## Example Usage

```typescript
import { PutInvoicesIdRequest } from "@flexprice/sdk/models/operations";

let value: PutInvoicesIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Invoice ID                                                                               |
| `body`                                                                                   | [components.DtoUpdateInvoiceRequest](../../models/components/dtoupdateinvoicerequest.md) | :heavy_check_mark:                                                                       | Invoice Update Request                                                                   |