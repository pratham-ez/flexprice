# TypesPriceEntityType

## Example Usage

```typescript
import { TypesPriceEntityType } from "@flexprice/sdk/models";

let value: TypesPriceEntityType = "PRICE";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"PLAN" | "SUBSCRIPTION" | "ADDON" | "PRICE" | "COSTSHEET" | Unrecognized<string>
```