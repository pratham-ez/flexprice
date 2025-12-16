# TypesChargebeeConnectionMetadata


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `APIKey`                                                       | **string*                                                      | :heavy_minus_sign:                                             | Chargebee API key (encrypted)                                  |
| `Site`                                                         | **string*                                                      | :heavy_minus_sign:                                             | Chargebee site name (not encrypted)                            |
| `WebhookPassword`                                              | **string*                                                      | :heavy_minus_sign:                                             | Basic Auth password for webhooks (encrypted)                   |
| `WebhookSecret`                                                | **string*                                                      | :heavy_minus_sign:                                             | Chargebee Webhook Secret (encrypted, optional, NOT USED in v2) |
| `WebhookUsername`                                              | **string*                                                      | :heavy_minus_sign:                                             | Basic Auth username for webhooks (encrypted)                   |