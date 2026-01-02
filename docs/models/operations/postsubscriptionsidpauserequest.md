# PostSubscriptionsIdPauseRequest

## Example Usage

```typescript
import { PostSubscriptionsIdPauseRequest } from "@flexprice/sdk/models/operations";

let value: PostSubscriptionsIdPauseRequest = {
  id: "<id>",
  body: {
    pauseMode: "period_end",
  },
};
```

## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | Subscription ID                                                                                  |
| `body`                                                                                           | [components.DtoPauseSubscriptionRequest](../../models/components/dtopausesubscriptionrequest.md) | :heavy_check_mark:                                                                               | Pause subscription request                                                                       |