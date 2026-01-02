# TypesRazorpayConnectionMetadata

## Example Usage

```typescript
import { TypesRazorpayConnectionMetadata } from "@flexprice/sdk/models/components";

let value: TypesRazorpayConnectionMetadata = {};
```

## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `keyId`                                       | *string*                                      | :heavy_minus_sign:                            | Razorpay Key ID (encrypted)                   |
| `secretKey`                                   | *string*                                      | :heavy_minus_sign:                            | Razorpay Secret Key (encrypted)               |
| `webhookSecret`                               | *string*                                      | :heavy_minus_sign:                            | Razorpay Webhook Secret (encrypted, optional) |