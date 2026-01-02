# PostWebhooksStripeTenantIdEnvironmentIdRequest

## Example Usage

```typescript
import { PostWebhooksStripeTenantIdEnvironmentIdRequest } from "@flexprice/sdk/models/operations";

let value: PostWebhooksStripeTenantIdEnvironmentIdRequest = {
  tenantId: "<id>",
  environmentId: "<id>",
  stripeSignature: "<value>",
};
```

## Fields

| Field                    | Type                     | Required                 | Description              |
| ------------------------ | ------------------------ | ------------------------ | ------------------------ |
| `tenantId`               | *string*                 | :heavy_check_mark:       | Tenant ID                |
| `environmentId`          | *string*                 | :heavy_check_mark:       | Environment ID           |
| `stripeSignature`        | *string*                 | :heavy_check_mark:       | Stripe webhook signature |