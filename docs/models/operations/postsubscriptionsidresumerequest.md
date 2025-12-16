# PostSubscriptionsIdResumeRequest

## Example Usage

```typescript
import { PostSubscriptionsIdResumeRequest } from "@flexprice/sdk/models/operations";

let value: PostSubscriptionsIdResumeRequest = {
  id: "<id>",
  body: {
    resumeMode: "auto",
  },
};
```

## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `id`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | Subscription ID                                                                     |
| `body`                                                                              | [models.DtoResumeSubscriptionRequest](../../models/dtoresumesubscriptionrequest.md) | :heavy_check_mark:                                                                  | Resume subscription request                                                         |