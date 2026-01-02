# TypesQuickBooksConnectionMetadata

## Example Usage

```typescript
import { TypesQuickBooksConnectionMetadata } from "@flexprice/sdk/models/components";

let value: TypesQuickBooksConnectionMetadata = {};
```

## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `accessToken`                                                                        | *string*                                                                             | :heavy_minus_sign:                                                                   | Managed internally - set after auth code exchange or token refresh                   |
| `authCode`                                                                           | *string*                                                                             | :heavy_minus_sign:                                                                   | Optional - for initial setup via auth code (will be cleared after token exchange)    |
| `clientId`                                                                           | *string*                                                                             | :heavy_minus_sign:                                                                   | Required for initial connection setup                                                |
| `clientSecret`                                                                       | *string*                                                                             | :heavy_minus_sign:                                                                   | OAuth Client Secret (encrypted)                                                      |
| `environment`                                                                        | *string*                                                                             | :heavy_minus_sign:                                                                   | "sandbox" or "production"                                                            |
| `incomeAccountId`                                                                    | *string*                                                                             | :heavy_minus_sign:                                                                   | Optional configuration                                                               |
| `oauthSessionData`                                                                   | *string*                                                                             | :heavy_minus_sign:                                                                   | Temporary OAuth session data (only used during OAuth flow, cleared after completion) |
| `realmId`                                                                            | *string*                                                                             | :heavy_minus_sign:                                                                   | QuickBooks Company ID (not encrypted)                                                |
| `redirectUri`                                                                        | *string*                                                                             | :heavy_minus_sign:                                                                   | OAuth Redirect URI (temporary)                                                       |
| `refreshToken`                                                                       | *string*                                                                             | :heavy_minus_sign:                                                                   | OAuth Refresh Token (encrypted)                                                      |
| `webhookVerifierToken`                                                               | *string*                                                                             | :heavy_minus_sign:                                                                   | Webhook security                                                                     |