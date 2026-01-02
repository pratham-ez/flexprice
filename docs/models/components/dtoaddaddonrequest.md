# DtoAddAddonRequest

## Example Usage

```typescript
import { DtoAddAddonRequest } from "@flexprice/sdk/models/components";

let value: DtoAddAddonRequest = {
  addonId: "<id>",
  subscriptionId: "<id>",
};
```

## Fields

| Field                 | Type                  | Required              | Description           |
| --------------------- | --------------------- | --------------------- | --------------------- |
| `addonId`             | *string*              | :heavy_check_mark:    | N/A                   |
| `metadata`            | Record<string, *any*> | :heavy_minus_sign:    | N/A                   |
| `startDate`           | *string*              | :heavy_minus_sign:    | N/A                   |
| `subscriptionId`      | *string*              | :heavy_check_mark:    | N/A                   |