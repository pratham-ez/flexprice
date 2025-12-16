# PutPlansIdRequest

## Example Usage

```typescript
import { PutPlansIdRequest } from "@flexprice/sdk/models/operations";

let value: PutPlansIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `id`                                                                | *string*                                                            | :heavy_check_mark:                                                  | Plan ID                                                             |
| `body`                                                              | [models.DtoUpdatePlanRequest](../../models/dtoupdateplanrequest.md) | :heavy_check_mark:                                                  | Plan update                                                         |