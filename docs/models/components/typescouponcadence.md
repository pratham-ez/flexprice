# TypesCouponCadence

## Example Usage

```typescript
import { TypesCouponCadence } from "@flexprice/sdk/models/components";

let value: TypesCouponCadence = "repeated";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"once" | "repeated" | "forever" | Unrecognized<string>
```