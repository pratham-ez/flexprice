# TypesChargebeeConnectionMetadata

## Example Usage

```typescript
import { TypesChargebeeConnectionMetadata } from "@flexprice/sdk/models";

let value: TypesChargebeeConnectionMetadata = {};
```

## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `apiKey`                                                       | *string*                                                       | :heavy_minus_sign:                                             | Chargebee API key (encrypted)                                  |
| `site`                                                         | *string*                                                       | :heavy_minus_sign:                                             | Chargebee site name (not encrypted)                            |
| `webhookPassword`                                              | *string*                                                       | :heavy_minus_sign:                                             | Basic Auth password for webhooks (encrypted)                   |
| `webhookSecret`                                                | *string*                                                       | :heavy_minus_sign:                                             | Chargebee Webhook Secret (encrypted, optional, NOT USED in v2) |
| `webhookUsername`                                              | *string*                                                       | :heavy_minus_sign:                                             | Basic Auth username for webhooks (encrypted)                   |