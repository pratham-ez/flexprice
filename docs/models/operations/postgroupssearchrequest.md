# PostGroupsSearchRequest


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `EntityType`                                    | **string*                                       | :heavy_minus_sign:                              | Filter by entity type (e.g., 'price')           |
| `Name`                                          | **string*                                       | :heavy_minus_sign:                              | Filter by group name (contains search)          |
| `LookupKey`                                     | **string*                                       | :heavy_minus_sign:                              | Filter by lookup key (exact match)              |
| `Limit`                                         | **int64*                                        | :heavy_minus_sign:                              | Number of items to return (default: 20)         |
| `Offset`                                        | **int64*                                        | :heavy_minus_sign:                              | Number of items to skip (default: 0)            |
| `SortBy`                                        | **string*                                       | :heavy_minus_sign:                              | Field to sort by (name, created_at, updated_at) |
| `SortOrder`                                     | **string*                                       | :heavy_minus_sign:                              | Sort order (asc, desc)                          |