# PostWebhooksRazorpayTenantIdEnvironmentIdRequest

## Example Usage

```typescript
import { PostWebhooksRazorpayTenantIdEnvironmentIdRequest } from "@flexprice/sdk/models/operations";

let value: PostWebhooksRazorpayTenantIdEnvironmentIdRequest = {
  tenantId: "<id>",
  environmentId: "<id>",
  xRazorpaySignature: "<value>",
};
```

## Fields

| Field                      | Type                       | Required                   | Description                |
| -------------------------- | -------------------------- | -------------------------- | -------------------------- |
| `tenantId`                 | *string*                   | :heavy_check_mark:         | Tenant ID                  |
| `environmentId`            | *string*                   | :heavy_check_mark:         | Environment ID             |
| `xRazorpaySignature`       | *string*                   | :heavy_check_mark:         | Razorpay webhook signature |