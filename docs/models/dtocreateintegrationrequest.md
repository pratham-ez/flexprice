# DtoCreateIntegrationRequest

## Example Usage

```typescript
import { DtoCreateIntegrationRequest } from "@flexprice/sdk/models";

let value: DtoCreateIntegrationRequest = {
  credentials: {},
  name: "<value>",
  provider: "stripe",
};
```

## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `credentials`                                                  | Record<string, *string*>                                       | :heavy_check_mark:                                             | N/A                                                            |
| `name`                                                         | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `provider`                                                     | [models.TypesSecretProvider](../models/typessecretprovider.md) | :heavy_check_mark:                                             | N/A                                                            |