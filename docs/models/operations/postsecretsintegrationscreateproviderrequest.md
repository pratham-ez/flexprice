# PostSecretsIntegrationsCreateProviderRequest

## Example Usage

```typescript
import { PostSecretsIntegrationsCreateProviderRequest } from "@flexprice/sdk/models/operations";

let value: PostSecretsIntegrationsCreateProviderRequest = {
  provider: "<value>",
  body: {
    credentials: {
      "key": "<value>",
      "key1": "<value>",
      "key2": "<value>",
    },
    name: "<value>",
    provider: "chargebee",
  },
};
```

## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `provider`                                                                                       | *string*                                                                                         | :heavy_check_mark:                                                                               | Integration provider                                                                             |
| `body`                                                                                           | [components.DtoCreateIntegrationRequest](../../models/components/dtocreateintegrationrequest.md) | :heavy_check_mark:                                                                               | Integration creation request                                                                     |