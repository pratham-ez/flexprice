# PutPricesUnitsIdRequest

## Example Usage

```typescript
import { PutPricesUnitsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutPricesUnitsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `id`                                                                          | *string*                                                                      | :heavy_check_mark:                                                            | Price unit ID                                                                 |
| `body`                                                                        | [models.DtoUpdatePriceUnitRequest](../../models/dtoupdatepriceunitrequest.md) | :heavy_check_mark:                                                            | Price unit details to update                                                  |