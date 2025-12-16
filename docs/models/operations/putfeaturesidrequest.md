# PutFeaturesIdRequest

## Example Usage

```typescript
import { PutFeaturesIdRequest } from "@flexprice/sdk/models/operations";

let value: PutFeaturesIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `id`                                                                      | *string*                                                                  | :heavy_check_mark:                                                        | Feature ID                                                                |
| `body`                                                                    | [models.DtoUpdateFeatureRequest](../../models/dtoupdatefeaturerequest.md) | :heavy_check_mark:                                                        | Feature update data                                                       |