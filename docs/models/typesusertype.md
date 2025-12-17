# TypesUserType

## Example Usage

```typescript
import { TypesUserType } from "@flexprice/sdk/models";

let value: TypesUserType = "user";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"user" | "service_account" | Unrecognized<string>
```