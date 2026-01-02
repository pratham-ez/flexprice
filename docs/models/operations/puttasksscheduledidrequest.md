# PutTasksScheduledIdRequest

## Example Usage

```typescript
import { PutTasksScheduledIdRequest } from "@flexprice/sdk/models/operations";

let value: PutTasksScheduledIdRequest = {
  id: "<id>",
  body: {
    enabled: false,
  },
};
```

## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `id`                                                                                                 | *string*                                                                                             | :heavy_check_mark:                                                                                   | Scheduled Task ID                                                                                    |
| `body`                                                                                               | [components.DtoUpdateScheduledTaskRequest](../../models/components/dtoupdatescheduledtaskrequest.md) | :heavy_check_mark:                                                                                   | Update request (enabled: true/false to pause/resume)                                                 |