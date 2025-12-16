# TypesAlertCondition

## Example Usage

```typescript
import { TypesAlertCondition } from "@flexprice/sdk/models";

let value: TypesAlertCondition = "above";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"above" | "below" | Unrecognized<string>
```