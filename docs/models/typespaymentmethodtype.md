# TypesPaymentMethodType

## Example Usage

```typescript
import { TypesPaymentMethodType } from "@flexprice/sdk/models";

let value: TypesPaymentMethodType = "CARD";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"CARD" | "ACH" | "OFFLINE" | "CREDITS" | "PAYMENT_LINK" | Unrecognized<string>
```