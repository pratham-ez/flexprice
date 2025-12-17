# TypesAlertState

## Example Usage

```typescript
import { TypesAlertState } from "@flexprice/sdk/models";

let value: TypesAlertState = "in_alarm";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"ok" | "info" | "warning" | "in_alarm" | Unrecognized<string>
```