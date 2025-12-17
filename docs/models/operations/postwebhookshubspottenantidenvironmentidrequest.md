# PostWebhooksHubspotTenantIdEnvironmentIdRequest

## Example Usage

```typescript
import { PostWebhooksHubspotTenantIdEnvironmentIdRequest } from "@flexprice/sdk/models/operations";

let value: PostWebhooksHubspotTenantIdEnvironmentIdRequest = {
  tenantId: "<id>",
  environmentId: "<id>",
  xHubSpotSignatureV3: "<value>",
};
```

## Fields

| Field                     | Type                      | Required                  | Description               |
| ------------------------- | ------------------------- | ------------------------- | ------------------------- |
| `tenantId`                | *string*                  | :heavy_check_mark:        | Tenant ID                 |
| `environmentId`           | *string*                  | :heavy_check_mark:        | Environment ID            |
| `xHubSpotSignatureV3`     | *string*                  | :heavy_check_mark:        | HubSpot webhook signature |