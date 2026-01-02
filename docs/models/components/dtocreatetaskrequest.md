# DtoCreateTaskRequest

## Example Usage

```typescript
import { DtoCreateTaskRequest } from "@flexprice/sdk/models/components";

let value: DtoCreateTaskRequest = {
  entityType: "CUSTOMERS",
  fileType: "JSON",
  fileUrl: "https://necessary-wombat.net/",
  taskType: "IMPORT",
};
```

## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `entityType`                                                             | [components.TypesEntityType](../../models/components/typesentitytype.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `fileName`                                                               | *string*                                                                 | :heavy_minus_sign:                                                       | N/A                                                                      |
| `fileType`                                                               | [components.TypesFileType](../../models/components/typesfiletype.md)     | :heavy_check_mark:                                                       | N/A                                                                      |
| `fileUrl`                                                                | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `metadata`                                                               | Record<string, *any*>                                                    | :heavy_minus_sign:                                                       | N/A                                                                      |
| `taskType`                                                               | [components.TypesTaskType](../../models/components/typestasktype.md)     | :heavy_check_mark:                                                       | N/A                                                                      |