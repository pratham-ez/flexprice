# DtoCreateCreditGrantRequest

## Example Usage

```typescript
import { DtoCreateCreditGrantRequest } from "@flexprice/sdk/models";

let value: DtoCreateCreditGrantRequest = {
  cadence: "RECURRING",
  credits: "<value>",
  name: "<value>",
  scope: "PLAN",
};
```

## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `cadence`                                                                                    | [models.TypesCreditGrantCadence](../models/typescreditgrantcadence.md)                       | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `credits`                                                                                    | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `expirationDuration`                                                                         | *number*                                                                                     | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `expirationDurationUnit`                                                                     | [models.TypesCreditGrantExpiryDurationUnit](../models/typescreditgrantexpirydurationunit.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `expirationType`                                                                             | [models.TypesCreditGrantExpiryType](../models/typescreditgrantexpirytype.md)                 | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `metadata`                                                                                   | Record<string, *string*>                                                                     | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `name`                                                                                       | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `period`                                                                                     | [models.TypesCreditGrantPeriod](../models/typescreditgrantperiod.md)                         | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `periodCount`                                                                                | *number*                                                                                     | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `planId`                                                                                     | *string*                                                                                     | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `priority`                                                                                   | *number*                                                                                     | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `scope`                                                                                      | [models.TypesCreditGrantScope](../models/typescreditgrantscope.md)                           | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `subscriptionId`                                                                             | *string*                                                                                     | :heavy_minus_sign:                                                                           | N/A                                                                                          |