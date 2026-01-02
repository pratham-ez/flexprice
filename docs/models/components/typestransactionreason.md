# TypesTransactionReason

## Example Usage

```typescript
import { TypesTransactionReason } from "@flexprice/sdk/models/components";

let value: TypesTransactionReason = "WALLET_TERMINATION";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"INVOICE_PAYMENT" | "FREE_CREDIT_GRANT" | "SUBSCRIPTION_CREDIT_GRANT" | "PURCHASED_CREDIT_INVOICED" | "PURCHASED_CREDIT_DIRECT" | "CREDIT_NOTE" | "CREDIT_EXPIRED" | "WALLET_TERMINATION" | "MANUAL_BALANCE_DEBIT" | Unrecognized<string>
```