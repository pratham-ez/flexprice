# TypesChargebeeConnectionMetadata


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `api_key`                                                      | *Optional[str]*                                                | :heavy_minus_sign:                                             | Chargebee API key (encrypted)                                  |
| `site`                                                         | *Optional[str]*                                                | :heavy_minus_sign:                                             | Chargebee site name (not encrypted)                            |
| `webhook_password`                                             | *Optional[str]*                                                | :heavy_minus_sign:                                             | Basic Auth password for webhooks (encrypted)                   |
| `webhook_secret`                                               | *Optional[str]*                                                | :heavy_minus_sign:                                             | Chargebee Webhook Secret (encrypted, optional, NOT USED in v2) |
| `webhook_username`                                             | *Optional[str]*                                                | :heavy_minus_sign:                                             | Basic Auth username for webhooks (encrypted)                   |