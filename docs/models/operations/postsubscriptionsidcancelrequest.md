# PostSubscriptionsIdCancelRequest

## Example Usage

```typescript
import { PostSubscriptionsIdCancelRequest } from "@flexprice/sdk/models/operations";

let value: PostSubscriptionsIdCancelRequest = {
  id: "<id>",
  body: {
    cancellationType: "end_of_period",
  },
};
```

## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `id`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | Subscription ID                                                                     |
| `body`                                                                              | [models.DtoCancelSubscriptionRequest](../../models/dtocancelsubscriptionrequest.md) | :heavy_check_mark:                                                                  | Cancel Subscription Request                                                         |