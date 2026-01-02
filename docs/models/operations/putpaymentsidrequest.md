# PutPaymentsIdRequest

## Example Usage

```typescript
import { PutPaymentsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutPaymentsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Payment ID                                                                               |
| `body`                                                                                   | [components.DtoUpdatePaymentRequest](../../models/components/dtoupdatepaymentrequest.md) | :heavy_check_mark:                                                                       | Payment configuration                                                                    |