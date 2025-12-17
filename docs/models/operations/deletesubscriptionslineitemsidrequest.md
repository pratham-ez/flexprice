# DeleteSubscriptionsLineitemsIdRequest

## Example Usage

```typescript
import { DeleteSubscriptionsLineitemsIdRequest } from "@flexprice/sdk/models/operations";

let value: DeleteSubscriptionsLineitemsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `id`                                                                                                | *string*                                                                                            | :heavy_check_mark:                                                                                  | Line Item ID                                                                                        |
| `body`                                                                                              | [models.DtoDeleteSubscriptionLineItemRequest](../../models/dtodeletesubscriptionlineitemrequest.md) | :heavy_check_mark:                                                                                  | Delete Line Item Request                                                                            |