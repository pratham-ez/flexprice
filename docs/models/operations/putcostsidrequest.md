# PutCostsIdRequest

## Example Usage

```typescript
import { PutCostsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutCostsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `id`                                                                          | *string*                                                                      | :heavy_check_mark:                                                            | Costsheet ID                                                                  |
| `body`                                                                        | [models.DtoUpdateCostsheetRequest](../../models/dtoupdatecostsheetrequest.md) | :heavy_check_mark:                                                            | Costsheet configuration                                                       |