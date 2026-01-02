# PutTaxesAssociationsIdRequest

## Example Usage

```typescript
import { PutTaxesAssociationsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutTaxesAssociationsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                   | *string*                                                                                               | :heavy_check_mark:                                                                                     | Tax Config ID                                                                                          |
| `body`                                                                                                 | [components.DtoTaxAssociationUpdateRequest](../../models/components/dtotaxassociationupdaterequest.md) | :heavy_check_mark:                                                                                     | Tax Config Request                                                                                     |