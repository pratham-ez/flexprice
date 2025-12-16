# TypesEntityType

## Example Usage

```typescript
import { TypesEntityType } from "@flexprice/sdk/models";

let value: TypesEntityType = "PRICES";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"EVENTS" | "PRICES" | "CUSTOMERS" | Unrecognized<string>
```