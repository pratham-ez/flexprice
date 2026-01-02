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

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `id`                                                                                           | *string*                                                                                       | :heavy_check_mark:                                                                             | Connection ID                                                                                  |
| `body`                                                                                         | [components.DtoUpdateConnectionRequest](../../models/components/dtoupdateconnectionrequest.md) | :heavy_check_mark:                                                                             | Connection                                                                                     |