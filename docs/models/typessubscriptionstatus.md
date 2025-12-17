# TypesSubscriptionStatus

## Example Usage

```typescript
import { TypesSubscriptionStatus } from "@flexprice/sdk/models";

let value: TypesSubscriptionStatus = "incomplete_expired";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"active" | "paused" | "cancelled" | "incomplete" | "incomplete_expired" | "past_due" | "trialing" | "unpaid" | "draft" | Unrecognized<string>
```