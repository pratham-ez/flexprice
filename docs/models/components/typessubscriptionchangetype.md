# TypesSubscriptionChangeType

## Example Usage

```typescript
import { TypesSubscriptionChangeType } from "@flexprice/sdk/models/components";

let value: TypesSubscriptionChangeType = "upgrade";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"upgrade" | "downgrade" | "lateral" | Unrecognized<string>
```