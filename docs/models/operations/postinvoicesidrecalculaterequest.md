# PostInvoicesIdRecalculateRequest

## Example Usage

```typescript
import { PostInvoicesIdRecalculateRequest } from "@flexprice/sdk/models/operations";

let value: PostInvoicesIdRecalculateRequest = {
  id: "<id>",
};
```

## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *string*                                                            | :heavy_check_mark:                                                  | Invoice ID                                                          |
| `finalize`                                                          | *boolean*                                                           | :heavy_minus_sign:                                                  | Whether to finalize the invoice after recalculation (default: true) |