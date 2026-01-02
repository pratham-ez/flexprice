# TypesBillingPeriod

## Example Usage

```typescript
import { TypesBillingPeriod } from "@flexprice/sdk/models/components";

let value: TypesBillingPeriod = "MONTHLY";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"MONTHLY" | "ANNUAL" | "WEEKLY" | "DAILY" | "QUARTERLY" | "HALF_YEARLY" | Unrecognized<string>
```