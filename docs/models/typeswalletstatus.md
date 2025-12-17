# TypesWalletStatus

## Example Usage

```typescript
import { TypesWalletStatus } from "@flexprice/sdk/models";

let value: TypesWalletStatus = "closed";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"active" | "frozen" | "closed" | Unrecognized<string>
```