# PostWebhooksQuickbooksTenantIdEnvironmentIdRequest

## Example Usage

```typescript
import { PostWebhooksQuickbooksTenantIdEnvironmentIdRequest } from "@flexprice/sdk/models/operations";

let value: PostWebhooksQuickbooksTenantIdEnvironmentIdRequest = {
  tenantId: "<id>",
  environmentId: "<id>",
};
```

## Fields

| Field                        | Type                         | Required                     | Description                  |
| ---------------------------- | ---------------------------- | ---------------------------- | ---------------------------- |
| `tenantId`                   | *string*                     | :heavy_check_mark:           | Tenant ID                    |
| `environmentId`              | *string*                     | :heavy_check_mark:           | Environment ID               |
| `intuitSignature`            | *string*                     | :heavy_minus_sign:           | QuickBooks webhook signature |