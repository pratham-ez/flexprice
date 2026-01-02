# TypesTransactionStatus

## Example Usage

```typescript
import { TypesTransactionStatus } from "@flexprice/sdk/models/components";

let value: TypesTransactionStatus = "failed";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"pending" | "completed" | "failed" | Unrecognized<string>
```