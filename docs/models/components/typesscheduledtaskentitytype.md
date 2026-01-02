# TypesScheduledTaskEntityType

## Example Usage

```typescript
import { TypesScheduledTaskEntityType } from "@flexprice/sdk/models/components";

let value: TypesScheduledTaskEntityType = "credit_topups";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"events" | "invoice" | "credit_topups" | "credit_usage" | Unrecognized<string>
```