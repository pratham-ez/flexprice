# DtoCreateScheduledTaskRequest

## Example Usage

```typescript
import { DtoCreateScheduledTaskRequest } from "@flexprice/sdk/models/components";

let value: DtoCreateScheduledTaskRequest = {
  connectionId: "<id>",
  entityType: "credit_topups",
  interval: "custom",
  jobConfig: {},
};
```

## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `connectionId`                                                                                     | *string*                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `enabled`                                                                                          | *boolean*                                                                                          | :heavy_minus_sign:                                                                                 | N/A                                                                                                |
| `entityType`                                                                                       | [components.TypesScheduledTaskEntityType](../../models/components/typesscheduledtaskentitytype.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `interval`                                                                                         | [components.TypesScheduledTaskInterval](../../models/components/typesscheduledtaskinterval.md)     | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `jobConfig`                                                                                        | [components.TypesS3JobConfig](../../models/components/typess3jobconfig.md)                         | :heavy_check_mark:                                                                                 | N/A                                                                                                |