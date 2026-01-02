# PutCreditgrantsIdRequest

## Example Usage

```typescript
import { PutCreditgrantsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutCreditgrantsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | Credit Grant ID                                                                                  |
| `body`                                                                                           | [components.DtoUpdateCreditGrantRequest](../../models/components/dtoupdatecreditgrantrequest.md) | :heavy_check_mark:                                                                               | Credit Grant configuration                                                                       |