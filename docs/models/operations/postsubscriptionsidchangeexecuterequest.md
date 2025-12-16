# PostSubscriptionsIdChangeExecuteRequest

## Example Usage

```typescript
import { PostSubscriptionsIdChangeExecuteRequest } from "@flexprice/sdk/models/operations";

let value: PostSubscriptionsIdChangeExecuteRequest = {
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

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `id`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | Subscription ID                                                                     |
| `body`                                                                              | [models.DtoSubscriptionChangeRequest](../../models/dtosubscriptionchangerequest.md) | :heavy_check_mark:                                                                  | Subscription change request                                                         |