# PutConnectionsIdRequest

## Example Usage

```typescript
import { PutConnectionsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutConnectionsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `id`                                                                            | *string*                                                                        | :heavy_check_mark:                                                              | Connection ID                                                                   |
| `body`                                                                          | [models.DtoUpdateConnectionRequest](../../models/dtoupdateconnectionrequest.md) | :heavy_check_mark:                                                              | Connection                                                                      |