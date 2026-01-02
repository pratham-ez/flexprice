# TypesCreditNoteReason

## Example Usage

```typescript
import { TypesCreditNoteReason } from "@flexprice/sdk/models/components";

let value: TypesCreditNoteReason = "SERVICE_ISSUE";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"DUPLICATE" | "FRAUDULENT" | "ORDER_CHANGE" | "UNSATISFACTORY" | "SERVICE_ISSUE" | "BILLING_ERROR" | "SUBSCRIPTION_CANCELLATION" | Unrecognized<string>
```