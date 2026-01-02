# TypesAddonStatus

## Example Usage

```typescript
import { TypesAddonStatus } from "@flexprice/sdk/models/components";

let value: TypesAddonStatus = "paused";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"active" | "cancelled" | "paused" | Unrecognized<string>
```