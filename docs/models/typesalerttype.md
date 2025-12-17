# TypesAlertType

## Example Usage

```typescript
import { TypesAlertType } from "@flexprice/sdk/models";

let value: TypesAlertType = "low_credit_balance";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"low_ongoing_balance" | "low_credit_balance" | "feature_wallet_balance" | Unrecognized<string>
```