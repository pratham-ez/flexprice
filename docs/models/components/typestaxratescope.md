# TypesTaxRateScope

## Example Usage

```typescript
import { TypesTaxRateScope } from "@flexprice/sdk/models/components";

let value: TypesTaxRateScope = "ONETIME";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"INTERNAL" | "EXTERNAL" | "ONETIME" | Unrecognized<string>
```