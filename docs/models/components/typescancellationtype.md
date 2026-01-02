# TypesCancellationType

## Example Usage

```typescript
import { TypesCancellationType } from "@flexprice/sdk/models/components";

let value: TypesCancellationType = "immediate";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"immediate" | "end_of_period" | Unrecognized<string>
```