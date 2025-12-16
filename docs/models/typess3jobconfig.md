# TypesS3JobConfig

## Example Usage

```typescript
import { TypesS3JobConfig } from "@flexprice/sdk/models";

let value: TypesS3JobConfig = {};
```

## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `bucket`                                                                       | *string*                                                                       | :heavy_minus_sign:                                                             | S3 bucket name                                                                 |
| `compression`                                                                  | [models.TypesS3CompressionType](../models/typess3compressiontype.md)           | :heavy_minus_sign:                                                             | N/A                                                                            |
| `encryption`                                                                   | [models.TypesS3EncryptionType](../models/typess3encryptiontype.md)             | :heavy_minus_sign:                                                             | N/A                                                                            |
| `endpointUrl`                                                                  | *string*                                                                       | :heavy_minus_sign:                                                             | Custom S3 endpoint URL (e.g., "http://minio:9000" for MinIO)                   |
| `keyPrefix`                                                                    | *string*                                                                       | :heavy_minus_sign:                                                             | Optional prefix for S3 keys (e.g., "flexprice-exports/")                       |
| `region`                                                                       | *string*                                                                       | :heavy_minus_sign:                                                             | AWS region (e.g., "us-west-2")                                                 |
| `usePathStyle`                                                                 | *boolean*                                                                      | :heavy_minus_sign:                                                             | Use path-style addressing instead of virtual-hosted-style (required for MinIO) |