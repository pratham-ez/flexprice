# TypesHubSpotConnectionMetadata

## Example Usage

```typescript
import { TypesHubSpotConnectionMetadata } from "@flexprice/sdk/models";

let value: TypesHubSpotConnectionMetadata = {};
```

## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `accessToken`                                                  | *string*                                                       | :heavy_minus_sign:                                             | Private App Access Token (encrypted)                           |
| `appId`                                                        | *string*                                                       | :heavy_minus_sign:                                             | HubSpot App ID (optional, not encrypted)                       |
| `clientSecret`                                                 | *string*                                                       | :heavy_minus_sign:                                             | Private App Client Secret for webhook verification (encrypted) |