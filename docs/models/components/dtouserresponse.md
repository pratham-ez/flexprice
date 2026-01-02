# DtoUserResponse

## Example Usage

```typescript
import { DtoUserResponse } from "@flexprice/sdk/models/components";

let value: DtoUserResponse = {};
```

## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `email`                                                                      | *string*                                                                     | :heavy_minus_sign:                                                           | Empty for service accounts                                                   |
| `id`                                                                         | *string*                                                                     | :heavy_minus_sign:                                                           | N/A                                                                          |
| `roles`                                                                      | *string*[]                                                                   | :heavy_minus_sign:                                                           | N/A                                                                          |
| `tenant`                                                                     | [components.DtoTenantResponse](../../models/components/dtotenantresponse.md) | :heavy_minus_sign:                                                           | N/A                                                                          |
| `type`                                                                       | [components.TypesUserType](../../models/components/typesusertype.md)         | :heavy_minus_sign:                                                           | N/A                                                                          |