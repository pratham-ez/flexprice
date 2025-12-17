# TypesNomodConnectionMetadata

## Example Usage

```typescript
import { TypesNomodConnectionMetadata } from "@flexprice/sdk/models";

let value: TypesNomodConnectionMetadata = {};
```

## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `apiKey`                                             | *string*                                             | :heavy_minus_sign:                                   | Nomod API Key (encrypted)                            |
| `webhookSecret`                                      | *string*                                             | :heavy_minus_sign:                                   | Basic Auth secret for webhooks (encrypted, optional) |