# TypesBillingCadence

## Example Usage

```typescript
import { TypesBillingCadence } from "@flexprice/sdk/models/components";

let value: TypesBillingCadence = "RECURRING";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"RECURRING" | "ONETIME" | Unrecognized<string>
```