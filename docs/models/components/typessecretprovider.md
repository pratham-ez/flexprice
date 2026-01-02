# TypesSecretProvider

## Example Usage

```typescript
import { TypesSecretProvider } from "@flexprice/sdk/models/components";

let value: TypesSecretProvider = "chargebee";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"flexprice" | "stripe" | "s3" | "hubspot" | "razorpay" | "chargebee" | "quickbooks" | "nomod" | Unrecognized<string>
```