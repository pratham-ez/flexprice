# InvoiceType

invoice_type filters by the nature of the invoice (SUBSCRIPTION, ONE_OFF, or CREDIT)
Use this to separate recurring charges from one-time fees or credit adjustments

## Example Usage

```typescript
import { InvoiceType } from "@flexprice/sdk/models/operations";

let value: InvoiceType = "CREDIT";
```

## Values

```typescript
"SUBSCRIPTION" | "ONE_OFF" | "CREDIT"
```