# DeletePricesIdRequest

## Example Usage

```typescript
import { DeletePricesIdRequest } from "@flexprice/sdk/models/operations";

let value: DeletePricesIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | Price ID                                                                             |
| `body`                                                                               | [components.DtoDeletePriceRequest](../../models/components/dtodeletepricerequest.md) | :heavy_check_mark:                                                                   | Delete Price Request                                                                 |