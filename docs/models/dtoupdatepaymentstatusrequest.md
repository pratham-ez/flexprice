# DtoUpdatePaymentStatusRequest

## Example Usage

```typescript
import { DtoUpdatePaymentStatusRequest } from "@flexprice/sdk/models";

let value: DtoUpdatePaymentStatusRequest = {
  paymentStatus: "INITIATED",
};
```

## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `amount`                                                     | *string*                                                     | :heavy_minus_sign:                                           | amount is the optional payment amount to record              |
| `paymentStatus`                                              | [models.TypesPaymentStatus](../models/typespaymentstatus.md) | :heavy_check_mark:                                           | N/A                                                          |