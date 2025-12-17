# TypesPauseMode

## Example Usage

```typescript
import { TypesPauseMode } from "@flexprice/sdk/models";

let value: TypesPauseMode = "scheduled";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"immediate" | "scheduled" | "period_end" | Unrecognized<string>
```