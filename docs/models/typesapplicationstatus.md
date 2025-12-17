# TypesApplicationStatus

## Example Usage

```typescript
import { TypesApplicationStatus } from "@flexprice/sdk/models";

let value: TypesApplicationStatus = "skipped";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"applied" | "failed" | "pending" | "skipped" | "cancelled" | Unrecognized<string>
```