# DtoCreateTaskRequest

## Example Usage

```typescript
import { DtoCreateTaskRequest } from "@flexprice/sdk/models";

let value: DtoCreateTaskRequest = {
  entityType: "CUSTOMERS",
  fileType: "JSON",
  fileUrl: "https://necessary-wombat.net/",
  taskType: "IMPORT",
};
```

## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `entityType`                                           | [models.TypesEntityType](../models/typesentitytype.md) | :heavy_check_mark:                                     | N/A                                                    |
| `fileName`                                             | *string*                                               | :heavy_minus_sign:                                     | N/A                                                    |
| `fileType`                                             | [models.TypesFileType](../models/typesfiletype.md)     | :heavy_check_mark:                                     | N/A                                                    |
| `fileUrl`                                              | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `metadata`                                             | Record<string, *any*>                                  | :heavy_minus_sign:                                     | N/A                                                    |
| `taskType`                                             | [models.TypesTaskType](../models/typestasktype.md)     | :heavy_check_mark:                                     | N/A                                                    |