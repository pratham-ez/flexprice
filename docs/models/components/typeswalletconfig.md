# TypesWalletConfig

## Example Usage

```typescript
import { TypesWalletConfig } from "@flexprice/sdk/models/components";

let value: TypesWalletConfig = {};
```

## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `allowedPriceTypes`                                                                                              | [components.TypesWalletConfigPriceType](../../models/components/typeswalletconfigpricetype.md)[]                 | :heavy_minus_sign:                                                                                               | AllowedPriceTypes is a list of price types that are allowed for the wallet<br/>nil means all price types are allowed |