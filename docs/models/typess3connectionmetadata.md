# TypesS3ConnectionMetadata

## Example Usage

```typescript
import { TypesS3ConnectionMetadata } from "@flexprice/sdk/models";

let value: TypesS3ConnectionMetadata = {};
```

## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `awsAccessKeyId`                                        | *string*                                                | :heavy_minus_sign:                                      | AWS access key (encrypted)                              |
| `awsSecretAccessKey`                                    | *string*                                                | :heavy_minus_sign:                                      | AWS secret access key (encrypted)                       |
| `awsSessionToken`                                       | *string*                                                | :heavy_minus_sign:                                      | AWS session token for temporary credentials (encrypted) |