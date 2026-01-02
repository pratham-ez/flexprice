# TypesCreditNoteStatus

## Example Usage

```typescript
import { TypesCreditNoteStatus } from "@flexprice/sdk/models/components";

let value: TypesCreditNoteStatus = "FINALIZED";
```

## Values

This is an open enum. Unrecognized values will be captured as the `Unrecognized<string>` branded type.

```typescript
"DRAFT" | "FINALIZED" | "VOIDED" | Unrecognized<string>
```