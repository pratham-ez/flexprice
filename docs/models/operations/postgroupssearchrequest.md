# PostGroupsSearchRequest

## Example Usage

```typescript
import { PostGroupsSearchRequest } from "@flexprice/sdk/models/operations";

let value: PostGroupsSearchRequest = {};
```

## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `entityType`                                    | *string*                                        | :heavy_minus_sign:                              | Filter by entity type (e.g., 'price')           |
| `name`                                          | *string*                                        | :heavy_minus_sign:                              | Filter by group name (contains search)          |
| `lookupKey`                                     | *string*                                        | :heavy_minus_sign:                              | Filter by lookup key (exact match)              |
| `limit`                                         | *number*                                        | :heavy_minus_sign:                              | Number of items to return (default: 20)         |
| `offset`                                        | *number*                                        | :heavy_minus_sign:                              | Number of items to skip (default: 0)            |
| `sortBy`                                        | *string*                                        | :heavy_minus_sign:                              | Field to sort by (name, created_at, updated_at) |
| `sortOrder`                                     | *string*                                        | :heavy_minus_sign:                              | Sort order (asc, desc)                          |