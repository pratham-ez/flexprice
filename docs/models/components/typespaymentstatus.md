# TypesPaymentStatus

## Example Usage

```typescript
import { TypesPaymentStatus } from "@flexprice/sdk/models/components";

let value: TypesPaymentStatus = "INITIATED";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"INITIATED" | "PENDING" | "PROCESSING" | "SUCCEEDED" | "OVERPAID" | "FAILED" | "REFUNDED" | "PARTIALLY_REFUNDED" | Unrecognized<string>
```