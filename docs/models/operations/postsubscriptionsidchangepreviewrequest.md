# PostSubscriptionsIdChangePreviewRequest

## Example Usage

```typescript
import { PostSubscriptionsIdChangePreviewRequest } from "@flexprice/sdk/models/operations";

let value: PostSubscriptionsIdChangePreviewRequest = {
  id: "<id>",
  body: {
    billingCadence: "ONETIME",
    billingCycle: "anniversary",
    billingPeriod: "MONTHLY",
    prorationBehavior: "none",
    targetPlanId: "<id>",
  },
};
```

## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | Subscription ID                                                                                    |
| `body`                                                                                             | [components.DtoSubscriptionChangeRequest](../../models/components/dtosubscriptionchangerequest.md) | :heavy_check_mark:                                                                                 | Subscription change preview request                                                                |