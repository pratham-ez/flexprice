# TypesSubscriptionStatus

## Example Usage

```typescript
import { TypesSubscriptionStatus } from "@flexprice/sdk/models/components";

let value: TypesSubscriptionStatus = "incomplete";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"active" | "paused" | "cancelled" | "incomplete" | "trialing" | "draft" | Unrecognized<string>
```