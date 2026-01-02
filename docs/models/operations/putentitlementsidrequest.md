# PutEntitlementsIdRequest

## Example Usage

```typescript
import { PutEntitlementsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutEntitlementsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | Entitlement ID                                                                                   |
| `body`                                                                                           | [components.DtoUpdateEntitlementRequest](../../models/components/dtoupdateentitlementrequest.md) | :heavy_check_mark:                                                                               | Entitlement configuration                                                                        |