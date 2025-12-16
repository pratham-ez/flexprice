# TypesPriceType

## Example Usage

```typescript
import { TypesPriceType } from "@flexprice/sdk/models";

let value: TypesPriceType = "USAGE";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"USAGE" | "FIXED" | Unrecognized<string>
```