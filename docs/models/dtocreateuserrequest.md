# DtoCreateUserRequest

## Example Usage

```typescript
import { DtoCreateUserRequest } from "@flexprice/sdk/models";

let value: DtoCreateUserRequest = {
  roles: [
    "<value 1>",
    "<value 2>",
    "<value 3>",
  ],
  type: "service_account",
};
```

## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `roles`                                            | *string*[]                                         | :heavy_check_mark:                                 | Roles are required                                 |
| `type`                                             | [models.TypesUserType](../models/typesusertype.md) | :heavy_check_mark:                                 | N/A                                                |