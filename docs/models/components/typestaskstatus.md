# TypesTaskStatus

## Example Usage

```typescript
import { TypesTaskStatus } from "@flexprice/sdk/models/components";

let value: TypesTaskStatus = "FAILED";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"PENDING" | "PROCESSING" | "COMPLETED" | "FAILED" | Unrecognized<string>
```