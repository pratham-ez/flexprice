# TypesFileType

## Example Usage

```typescript
import { TypesFileType } from "@flexprice/sdk/models/components";

let value: TypesFileType = "JSON";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"CSV" | "JSON" | Unrecognized<string>
```