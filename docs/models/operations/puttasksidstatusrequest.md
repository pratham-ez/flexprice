# PutTasksIdStatusRequest

## Example Usage

```typescript
import { PutTasksIdStatusRequest } from "@flexprice/sdk/models/operations";

let value: PutTasksIdStatusRequest = {
  id: "<id>",
  body: {
    taskStatus: "PENDING",
  },
};
```

## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `id`                                                                            | *string*                                                                        | :heavy_check_mark:                                                              | Task ID                                                                         |
| `body`                                                                          | [models.DtoUpdateTaskStatusRequest](../../models/dtoupdatetaskstatusrequest.md) | :heavy_check_mark:                                                              | Status update                                                                   |