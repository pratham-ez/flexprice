# PutTaxesRatesIdRequest

## Example Usage

```typescript
import { PutTaxesRatesIdRequest } from "@flexprice/sdk/models/operations";

let value: PutTaxesRatesIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Tax rate ID                                                                              |
| `body`                                                                                   | [components.DtoUpdateTaxRateRequest](../../models/components/dtoupdatetaxraterequest.md) | :heavy_check_mark:                                                                       | Tax rate to update                                                                       |