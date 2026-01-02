# PutSubscriptionsLineitemsIdRequest

## Example Usage

```typescript
import { PutSubscriptionsLineitemsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutSubscriptionsLineitemsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                                                              | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                               | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | Line Item ID                                                                                                       |
| `body`                                                                                                             | [components.DtoUpdateSubscriptionLineItemRequest](../../models/components/dtoupdatesubscriptionlineitemrequest.md) | :heavy_check_mark:                                                                                                 | Update Line Item Request                                                                                           |