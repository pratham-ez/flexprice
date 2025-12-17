# PostSubscriptionsIdActivateRequest

## Example Usage

```typescript
import { PostSubscriptionsIdActivateRequest } from "@flexprice/sdk/models/operations";

let value: PostSubscriptionsIdActivateRequest = {
  id: "<id>",
  body: {
    startDate: "<value>",
  },
};
```

## Fields

| Field                                                                                             | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `id`                                                                                              | *string*                                                                                          | :heavy_check_mark:                                                                                | Subscription ID                                                                                   |
| `body`                                                                                            | [models.DtoActivateDraftSubscriptionRequest](../../models/dtoactivatedraftsubscriptionrequest.md) | :heavy_check_mark:                                                                                | Activate Draft Subscription Request                                                               |