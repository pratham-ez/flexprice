# TypesBillingModel

## Example Usage

```typescript
import { TypesBillingModel } from "@flexprice/sdk/models";

let value: TypesBillingModel = "FLAT_FEE";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"FLAT_FEE" | "PACKAGE" | "TIERED" | Unrecognized<string>
```