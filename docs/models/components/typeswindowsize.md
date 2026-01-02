# TypesWindowSize

## Example Usage

```typescript
import { TypesWindowSize } from "@flexprice/sdk/models/components";

let value: TypesWindowSize = "6HOUR";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"MINUTE" | "15MIN" | "30MIN" | "HOUR" | "3HOUR" | "6HOUR" | "12HOUR" | "DAY" | "WEEK" | "MONTH" | Unrecognized<string>
```