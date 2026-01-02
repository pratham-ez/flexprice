# DtoOverrideLineItemRequest

## Example Usage

```typescript
import { DtoOverrideLineItemRequest } from "@flexprice/sdk/models/components";

let value: DtoOverrideLineItemRequest = {
  priceId: "<id>",
};
```

## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `amount`                                                                               | *string*                                                                               | :heavy_minus_sign:                                                                     | Amount is the new price amount that overrides the original price (optional)            |
| `billingModel`                                                                         | [components.TypesBillingModel](../../models/components/typesbillingmodel.md)           | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `priceId`                                                                              | *string*                                                                               | :heavy_check_mark:                                                                     | PriceID references the plan price to override                                          |
| `quantity`                                                                             | *string*                                                                               | :heavy_minus_sign:                                                                     | Quantity for this line item (optional)                                                 |
| `tierMode`                                                                             | [components.TypesBillingTier](../../models/components/typesbillingtier.md)             | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `tiers`                                                                                | [components.DtoCreatePriceTier](../../models/components/dtocreatepricetier.md)[]       | :heavy_minus_sign:                                                                     | Tiers determines the pricing tiers for this line item                                  |
| `transformQuantity`                                                                    | [components.PriceTransformQuantity](../../models/components/pricetransformquantity.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |