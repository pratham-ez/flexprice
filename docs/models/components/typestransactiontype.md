# TypesTransactionType

## Example Usage

```typescript
import { TypesTransactionType } from "@flexprice/sdk/models/components";

let value: TypesTransactionType = "credit";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"credit" | "debit" | Unrecognized<string>
```