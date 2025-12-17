# DtoCreatePriceUnitRequest

## Example Usage

```typescript
import { DtoCreatePriceUnitRequest } from "@flexprice/sdk/models";

let value: DtoCreatePriceUnitRequest = {
  baseCurrency: "<value>",
  code: "<value>",
  conversionRate: "<value>",
  name: "<value>",
  symbol: "<value>",
};
```

## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `baseCurrency`     | *string*           | :heavy_check_mark: | N/A                |
| `code`             | *string*           | :heavy_check_mark: | N/A                |
| `conversionRate`   | *string*           | :heavy_check_mark: | N/A                |
| `name`             | *string*           | :heavy_check_mark: | N/A                |
| `precision`        | *number*           | :heavy_minus_sign: | N/A                |
| `symbol`           | *string*           | :heavy_check_mark: | N/A                |