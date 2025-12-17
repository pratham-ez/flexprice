# DtoCreateScheduledTaskRequest

## Example Usage

```typescript
import { DtoCreateScheduledTaskRequest } from "@flexprice/sdk/models";

let value: DtoCreateScheduledTaskRequest = {
  connectionId: "<id>",
  entityType: "invoice",
  interval: "hourly",
  jobConfig: {},
};
```

## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `connectionId`                                                                   | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              |
| `enabled`                                                                        | *boolean*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |
| `entityType`                                                                     | [models.TypesScheduledTaskEntityType](../models/typesscheduledtaskentitytype.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `interval`                                                                       | [models.TypesScheduledTaskInterval](../models/typesscheduledtaskinterval.md)     | :heavy_check_mark:                                                               | N/A                                                                              |
| `jobConfig`                                                                      | [models.TypesS3JobConfig](../models/typess3jobconfig.md)                         | :heavy_check_mark:                                                               | N/A                                                                              |