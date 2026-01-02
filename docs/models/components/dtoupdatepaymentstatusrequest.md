# DtoUpdatePaymentStatusRequest

## Example Usage

```typescript
import { DtoUpdatePaymentStatusRequest } from "@flexprice/sdk/models/components";

let value: DtoUpdatePaymentStatusRequest = {
  paymentStatus: "INITIATED",
};
```

## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `amount`                                                                       | *string*                                                                       | :heavy_minus_sign:                                                             | amount is the optional payment amount to record                                |
| `paymentStatus`                                                                | [components.TypesPaymentStatus](../../models/components/typespaymentstatus.md) | :heavy_check_mark:                                                             | N/A                                                                            |