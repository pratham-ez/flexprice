# DtoCancelSubscriptionRequest

## Example Usage

```typescript
import { DtoCancelSubscriptionRequest } from "@flexprice/sdk/models";

let value: DtoCancelSubscriptionRequest = {
  cancellationType: "immediate",
};
```

## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `cancellationType`                                                   | [models.TypesCancellationType](../models/typescancellationtype.md)   | :heavy_check_mark:                                                   | N/A                                                                  |
| `prorationBehavior`                                                  | [models.TypesProrationBehavior](../models/typesprorationbehavior.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `reason`                                                             | *string*                                                             | :heavy_minus_sign:                                                   | Reason for cancellation (for audit and business intelligence)        |