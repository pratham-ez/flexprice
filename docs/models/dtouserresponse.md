# DtoUserResponse

## Example Usage

```typescript
import { DtoUserResponse } from "@flexprice/sdk/models";

let value: DtoUserResponse = {};
```

## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `email`                                                    | *string*                                                   | :heavy_minus_sign:                                         | Empty for service accounts                                 |
| `id`                                                       | *string*                                                   | :heavy_minus_sign:                                         | N/A                                                        |
| `roles`                                                    | *string*[]                                                 | :heavy_minus_sign:                                         | N/A                                                        |
| `tenant`                                                   | [models.DtoTenantResponse](../models/dtotenantresponse.md) | :heavy_minus_sign:                                         | N/A                                                        |
| `type`                                                     | [models.TypesUserType](../models/typesusertype.md)         | :heavy_minus_sign:                                         | N/A                                                        |