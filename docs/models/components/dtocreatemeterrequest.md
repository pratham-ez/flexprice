# DtoCreateMeterRequest

## Example Usage

```typescript
import { DtoCreateMeterRequest } from "@flexprice/sdk/models/components";

let value: DtoCreateMeterRequest = {
  aggregation: {},
  eventName: "api_request",
  name: "API Usage Meter",
  resetUsage: "BILLING_PERIOD",
};
```

## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `aggregation`                                                              | [components.MeterAggregation](../../models/components/meteraggregation.md) | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `eventName`                                                                | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        | api_request                                                                |
| `filters`                                                                  | [components.MeterFilter](../../models/components/meterfilter.md)[]         | :heavy_minus_sign:                                                         | N/A                                                                        |                                                                            |
| `name`                                                                     | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        | API Usage Meter                                                            |
| `resetUsage`                                                               | [components.TypesResetUsage](../../models/components/typesresetusage.md)   | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |