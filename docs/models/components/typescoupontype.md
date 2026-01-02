# TypesCouponType

## Example Usage

```typescript
import { TypesCouponType } from "@flexprice/sdk/models/components";

let value: TypesCouponType = "fixed";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"fixed" | "percentage" | Unrecognized<string>
```