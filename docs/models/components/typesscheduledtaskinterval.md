# TypesScheduledTaskInterval

## Example Usage

```typescript
import { TypesScheduledTaskInterval } from "@flexprice/sdk/models/components";

let value: TypesScheduledTaskInterval = "daily";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"15MIN" | "custom" | "hourly" | "daily" | Unrecognized<string>
```