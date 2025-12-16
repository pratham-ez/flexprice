# GetEntityIntegrationMappingsRequest


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `EntityID`                                          | **string*                                           | :heavy_minus_sign:                                  | Filter by FlexPrice entity ID                       |
| `EntityType`                                        | **string*                                           | :heavy_minus_sign:                                  | Filter by entity type                               |
| `ProviderType`                                      | **string*                                           | :heavy_minus_sign:                                  | Filter by provider type                             |
| `ProviderEntityID`                                  | **string*                                           | :heavy_minus_sign:                                  | Filter by provider entity ID                        |
| `Limit`                                             | **int64*                                            | :heavy_minus_sign:                                  | Number of results to return (default: 20, max: 100) |
| `Offset`                                            | **int64*                                            | :heavy_minus_sign:                                  | Pagination offset (default: 0)                      |