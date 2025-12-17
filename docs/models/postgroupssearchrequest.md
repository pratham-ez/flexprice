# PostGroupsSearchRequest


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `entity_type`                                   | *Optional[str]*                                 | :heavy_minus_sign:                              | Filter by entity type (e.g., 'price')           |
| `name`                                          | *Optional[str]*                                 | :heavy_minus_sign:                              | Filter by group name (contains search)          |
| `lookup_key`                                    | *Optional[str]*                                 | :heavy_minus_sign:                              | Filter by lookup key (exact match)              |
| `limit`                                         | *Optional[int]*                                 | :heavy_minus_sign:                              | Number of items to return (default: 20)         |
| `offset`                                        | *Optional[int]*                                 | :heavy_minus_sign:                              | Number of items to skip (default: 0)            |
| `sort_by`                                       | *Optional[str]*                                 | :heavy_minus_sign:                              | Field to sort by (name, created_at, updated_at) |
| `sort_order`                                    | *Optional[str]*                                 | :heavy_minus_sign:                              | Sort order (asc, desc)                          |