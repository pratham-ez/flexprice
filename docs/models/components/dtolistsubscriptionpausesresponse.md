# DtoListSubscriptionPausesResponse

Response object for listing subscription pauses with total count

## Example Usage

```typescript
import { DtoListSubscriptionPausesResponse } from "@flexprice/sdk/models/components";

let value: DtoListSubscriptionPausesResponse = {};
```

## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `items`                                                                                              | [components.DtoSubscriptionPauseResponse](../../models/components/dtosubscriptionpauseresponse.md)[] | :heavy_minus_sign:                                                                                   | List of subscription pause objects<br/>@Description Array of subscription pauses                     |
| `total`                                                                                              | *number*                                                                                             | :heavy_minus_sign:                                                                                   | Total number of pauses<br/>@Description Total count of subscription pauses in the response           |