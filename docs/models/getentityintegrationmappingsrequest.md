# GetEntityIntegrationMappingsRequest


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `entity_id`                                         | *Optional[str]*                                     | :heavy_minus_sign:                                  | Filter by FlexPrice entity ID                       |
| `entity_type`                                       | *Optional[str]*                                     | :heavy_minus_sign:                                  | Filter by entity type                               |
| `provider_type`                                     | *Optional[str]*                                     | :heavy_minus_sign:                                  | Filter by provider type                             |
| `provider_entity_id`                                | *Optional[str]*                                     | :heavy_minus_sign:                                  | Filter by provider entity ID                        |
| `limit`                                             | *Optional[int]*                                     | :heavy_minus_sign:                                  | Number of results to return (default: 20, max: 100) |
| `offset`                                            | *Optional[int]*                                     | :heavy_minus_sign:                                  | Pagination offset (default: 0)                      |