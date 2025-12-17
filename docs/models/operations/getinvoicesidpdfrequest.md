# GetInvoicesIdPdfRequest

## Example Usage

```typescript
import { GetInvoicesIdPdfRequest } from "@flexprice/sdk/models/operations";

let value: GetInvoicesIdPdfRequest = {
  id: "<id>",
};
```

## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `id`                                        | *string*                                    | :heavy_check_mark:                          | Invoice ID                                  |
| `url`                                       | *boolean*                                   | :heavy_minus_sign:                          | Return presigned URL from s3 instead of PDF |