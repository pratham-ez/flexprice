# TypesFeatureType

## Example Usage

```typescript
import { TypesFeatureType } from "@flexprice/sdk/models";

let value: TypesFeatureType = "metered";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"metered" | "boolean" | "static" | Unrecognized<string>
```