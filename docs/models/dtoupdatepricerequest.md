# DtoUpdatePriceRequest

## Example Usage

```typescript
import { DtoUpdatePriceRequest } from "@flexprice/sdk/models";

let value: DtoUpdatePriceRequest = {};
```

## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `amount`                                                                           | *string*                                                                           | :heavy_minus_sign:                                                                 | Amount is the new price amount that overrides the original price (optional)        |
| `billingModel`                                                                     | [models.TypesBillingModel](../models/typesbillingmodel.md)                         | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `description`                                                                      | *string*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `effectiveFrom`                                                                    | *string*                                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `groupId`                                                                          | *string*                                                                           | :heavy_minus_sign:                                                                 | GroupID is the id of the group to update the price in                              |
| `lookupKey`                                                                        | *string*                                                                           | :heavy_minus_sign:                                                                 | All price fields that can be updated<br/>Non-critical fields (can be updated directly) |
| `metadata`                                                                         | Record<string, *string*>                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `tierMode`                                                                         | [models.TypesBillingTier](../models/typesbillingtier.md)                           | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `tiers`                                                                            | [models.DtoCreatePriceTier](../models/dtocreatepricetier.md)[]                     | :heavy_minus_sign:                                                                 | Tiers determines the pricing tiers for this line item                              |
| `transformQuantity`                                                                | [models.PriceTransformQuantity](../models/pricetransformquantity.md)               | :heavy_minus_sign:                                                                 | N/A                                                                                |