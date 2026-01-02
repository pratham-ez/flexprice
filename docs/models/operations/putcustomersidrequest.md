# PutCustomersIdRequest

## Example Usage

```typescript
import { PutCustomersIdRequest } from "@flexprice/sdk/models/operations";

let value: PutCustomersIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `id`                                                                                       | *string*                                                                                   | :heavy_check_mark:                                                                         | Customer ID                                                                                |
| `body`                                                                                     | [components.DtoUpdateCustomerRequest](../../models/components/dtoupdatecustomerrequest.md) | :heavy_check_mark:                                                                         | Customer                                                                                   |