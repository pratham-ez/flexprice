# TypesS3EncryptionType

## Example Usage

```typescript
import { TypesS3EncryptionType } from "@flexprice/sdk/models";

let value: TypesS3EncryptionType = "AES256";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"AES256" | "aws:kms" | "aws:kms:dsse" | Unrecognized<string>
```