# DtoCreateEntityIntegrationMappingRequest

## Example Usage

```typescript
import { DtoCreateEntityIntegrationMappingRequest } from "@flexprice/sdk/models/components";

let value: DtoCreateEntityIntegrationMappingRequest = {
  entityId: "<id>",
  entityType: "item",
  providerEntityId: "<id>",
  providerType: "<value>",
};
```

## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `entityId`                                                                                     | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `entityType`                                                                                   | [components.TypesIntegrationEntityType](../../models/components/typesintegrationentitytype.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `metadata`                                                                                     | Record<string, *any*>                                                                          | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `providerEntityId`                                                                             | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `providerType`                                                                                 | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |