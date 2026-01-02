# PutEnvironmentsIdRequest

## Example Usage

```typescript
import { PutEnvironmentsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutEnvironmentsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | Environment ID                                                                                   |
| `body`                                                                                           | [components.DtoUpdateEnvironmentRequest](../../models/components/dtoupdateenvironmentrequest.md) | :heavy_check_mark:                                                                               | Environment                                                                                      |