# DtoUpdateSubscriptionLineItemRequest

## Example Usage

```typescript
import { DtoUpdateSubscriptionLineItemRequest } from "@flexprice/sdk/models";

let value: DtoUpdateSubscriptionLineItemRequest = {};
```

## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `amount`                                                                    | *string*                                                                    | :heavy_minus_sign:                                                          | Amount is the new price amount that overrides the original price            |
| `billingModel`                                                              | [models.TypesBillingModel](../models/typesbillingmodel.md)                  | :heavy_minus_sign:                                                          | N/A                                                                         |
| `effectiveFrom`                                                             | *string*                                                                    | :heavy_minus_sign:                                                          | EffectiveFrom for the existing line item (if not provided, defaults to now) |
| `metadata`                                                                  | Record<string, *string*>                                                    | :heavy_minus_sign:                                                          | Metadata for the new line item                                              |
| `tierMode`                                                                  | [models.TypesBillingTier](../models/typesbillingtier.md)                    | :heavy_minus_sign:                                                          | N/A                                                                         |
| `tiers`                                                                     | [models.DtoCreatePriceTier](../models/dtocreatepricetier.md)[]              | :heavy_minus_sign:                                                          | Tiers determines the pricing tiers for this line item                       |
| `transformQuantity`                                                         | [models.PriceTransformQuantity](../models/pricetransformquantity.md)        | :heavy_minus_sign:                                                          | N/A                                                                         |