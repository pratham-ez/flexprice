# PutCouponsIdRequest

## Example Usage

```typescript
import { PutCouponsIdRequest } from "@flexprice/sdk/models/operations";

let value: PutCouponsIdRequest = {
  id: "<id>",
  body: {},
};
```

## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `id`                                                                    | *string*                                                                | :heavy_check_mark:                                                      | Coupon ID                                                               |
| `body`                                                                  | [models.DtoUpdateCouponRequest](../../models/dtoupdatecouponrequest.md) | :heavy_check_mark:                                                      | Coupon update request                                                   |