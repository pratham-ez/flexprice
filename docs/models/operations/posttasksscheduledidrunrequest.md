# PostTasksScheduledIdRunRequest

## Example Usage

```typescript
import { PostTasksScheduledIdRunRequest } from "@flexprice/sdk/models/operations";

let value: PostTasksScheduledIdRunRequest = {
  id: "<id>",
};
```

## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `id`                                                                          | *string*                                                                      | :heavy_check_mark:                                                            | Scheduled Task ID                                                             |
| `body`                                                                        | [models.DtoTriggerForceRunRequest](../../models/dtotriggerforcerunrequest.md) | :heavy_minus_sign:                                                            | Optional start and end time for custom range                                  |