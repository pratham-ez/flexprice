# TypesTaxRateEntityType

## Example Usage

```typescript
import { TypesTaxRateEntityType } from "@flexprice/sdk/models/components";

let value: TypesTaxRateEntityType = "invoice";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"customer" | "subscription" | "invoice" | "tenant" | Unrecognized<string>
```