# DtoOverrideLineItemRequest

## Example Usage

```typescript
import { DtoOverrideLineItemRequest } from "@flexprice/sdk/models";

let value: DtoOverrideLineItemRequest = {
  priceId: "<id>",
};
```

## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `amount`                                                                    | *string*                                                                    | :heavy_minus_sign:                                                          | Amount is the new price amount that overrides the original price (optional) |
| `billingModel`                                                              | [models.TypesBillingModel](../models/typesbillingmodel.md)                  | :heavy_minus_sign:                                                          | N/A                                                                         |
| `priceId`                                                                   | *string*                                                                    | :heavy_check_mark:                                                          | PriceID references the plan price to override                               |
| `quantity`                                                                  | *string*                                                                    | :heavy_minus_sign:                                                          | Quantity for this line item (optional)                                      |
| `tierMode`                                                                  | [models.TypesBillingTier](../models/typesbillingtier.md)                    | :heavy_minus_sign:                                                          | N/A                                                                         |
| `tiers`                                                                     | [models.DtoCreatePriceTier](../models/dtocreatepricetier.md)[]              | :heavy_minus_sign:                                                          | Tiers determines the pricing tiers for this line item                       |
| `transformQuantity`                                                         | [models.PriceTransformQuantity](../models/pricetransformquantity.md)        | :heavy_minus_sign:                                                          | N/A                                                                         |