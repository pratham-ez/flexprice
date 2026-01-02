# DtoCreateFeatureRequest

## Example Usage

```typescript
import { DtoCreateFeatureRequest } from "@flexprice/sdk/models/components";

let value: DtoCreateFeatureRequest = {
  meter: {
    aggregation: {},
    eventName: "api_request",
    name: "API Usage Meter",
    resetUsage: "BILLING_PERIOD",
  },
  name: "<value>",
  type: "boolean",
};
```

## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `alertSettings`                                                                      | [components.TypesAlertSettings](../../models/components/typesalertsettings.md)       | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `description`                                                                        | *string*                                                                             | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `lookupKey`                                                                          | *string*                                                                             | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `metadata`                                                                           | Record<string, *string*>                                                             | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `meter`                                                                              | [components.DtoCreateMeterRequest](../../models/components/dtocreatemeterrequest.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `meterId`                                                                            | *string*                                                                             | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `name`                                                                               | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `type`                                                                               | [components.TypesFeatureType](../../models/components/typesfeaturetype.md)           | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `unitPlural`                                                                         | *string*                                                                             | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `unitSingular`                                                                       | *string*                                                                             | :heavy_minus_sign:                                                                   | N/A                                                                                  |