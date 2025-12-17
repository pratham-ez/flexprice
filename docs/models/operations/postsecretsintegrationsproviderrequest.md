# PostSecretsIntegrationsProviderRequest

## Example Usage

```typescript
import { PostSecretsIntegrationsProviderRequest } from "@flexprice/sdk/models/operations";

let value: PostSecretsIntegrationsProviderRequest = {
  provider: "<value>",
  body: {
    credentials: {
      "key": "<value>",
      "key1": "<value>",
      "key2": "<value>",
    },
    name: "<value>",
    provider: "quickbooks",
  },
};
```

## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `provider`                                                                        | *string*                                                                          | :heavy_check_mark:                                                                | Integration provider                                                              |
| `body`                                                                            | [models.DtoCreateIntegrationRequest](../../models/dtocreateintegrationrequest.md) | :heavy_check_mark:                                                                | Integration creation request                                                      |