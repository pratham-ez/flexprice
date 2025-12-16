# PutPricesIdRequest

## Example Usage

```typescript
import { PutPricesIdRequest } from "@flexprice/sdk/models/operations";

let value: PutPricesIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `id`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | Price ID                                                              |
| `body`                                                                | [models.DtoUpdatePriceRequest](../../models/dtoupdatepricerequest.md) | :heavy_check_mark:                                                    | Price configuration                                                   |