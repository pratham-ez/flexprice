# TypesPauseStatus

## Example Usage

```typescript
import { TypesPauseStatus } from "@flexprice/sdk/models";

let value: TypesPauseStatus = "none";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"none" | "active" | "scheduled" | "completed" | "cancelled" | Unrecognized<string>
```