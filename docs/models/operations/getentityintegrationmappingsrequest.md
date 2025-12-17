# GetEntityIntegrationMappingsRequest

## Example Usage

```typescript
import { GetEntityIntegrationMappingsRequest } from "@flexprice/sdk/models/operations";

let value: GetEntityIntegrationMappingsRequest = {};
```

## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `entityId`                                          | *string*                                            | :heavy_minus_sign:                                  | Filter by FlexPrice entity ID                       |
| `entityType`                                        | *string*                                            | :heavy_minus_sign:                                  | Filter by entity type                               |
| `providerType`                                      | *string*                                            | :heavy_minus_sign:                                  | Filter by provider type                             |
| `providerEntityId`                                  | *string*                                            | :heavy_minus_sign:                                  | Filter by provider entity ID                        |
| `limit`                                             | *number*                                            | :heavy_minus_sign:                                  | Number of results to return (default: 20, max: 100) |
| `offset`                                            | *number*                                            | :heavy_minus_sign:                                  | Pagination offset (default: 0)                      |