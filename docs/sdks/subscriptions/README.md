# Subscriptions

## Overview

### Available Operations

* [getSubscriptions](#getsubscriptions) - List subscriptions
* [postSubscriptions](#postsubscriptions) - Create subscription
* [postSubscriptionsAddon](#postsubscriptionsaddon) - Add addon to subscription
* [deleteSubscriptionsAddon](#deletesubscriptionsaddon) - Remove addon from subscription
* [putSubscriptionsLineitemsId](#putsubscriptionslineitemsid) - Update subscription line item
* [deleteSubscriptionsLineitemsId](#deletesubscriptionslineitemsid) - Delete subscription line item
* [postSubscriptionsSearch](#postsubscriptionssearch) - List subscriptions by filter
* [postSubscriptionsUsage](#postsubscriptionsusage) - Get usage by subscription
* [getSubscriptionsId](#getsubscriptionsid) - Get subscription
* [postSubscriptionsIdActivate](#postsubscriptionsidactivate) - Activate draft subscription
* [getSubscriptionsIdAddonsAssociations](#getsubscriptionsidaddonsassociations) - Get active addon associations
* [postSubscriptionsIdCancel](#postsubscriptionsidcancel) - Cancel subscription
* [postSubscriptionsIdChangeExecute](#postsubscriptionsidchangeexecute) - Execute subscription plan change
* [postSubscriptionsIdChangePreview](#postsubscriptionsidchangepreview) - Preview subscription plan change
* [getSubscriptionsIdEntitlements](#getsubscriptionsidentitlements) - Get subscription entitlements
* [getSubscriptionsIdGrantsUpcoming](#getsubscriptionsidgrantsupcoming) - Get upcoming credit grant applications
* [postSubscriptionsIdPause](#postsubscriptionsidpause) - Pause a subscription
* [getSubscriptionsIdPauses](#getsubscriptionsidpauses) - List all pauses for a subscription
* [postSubscriptionsIdResume](#postsubscriptionsidresume) - Resume a paused subscription

## getSubscriptions

Get subscriptions with optional filtering

### Example Usage

<!-- UsageSnippet language="typescript" operationID="get_/subscriptions" method="get" path="/subscriptions" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.getSubscriptions({});

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsGetSubscriptions } from "@flexprice/sdk/funcs/subscriptionsGetSubscriptions.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsGetSubscriptions(flexPrice, {});
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsGetSubscriptions failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [operations.GetSubscriptionsRequest](../../models/operations/getsubscriptionsrequest.md)                                                                                       | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoListSubscriptionsResponse](../../models/components/dtolistsubscriptionsresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postSubscriptions

Create a new subscription

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/subscriptions" method="post" path="/subscriptions" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.postSubscriptions({
    billingCadence: "RECURRING",
    billingPeriod: "HALF_YEARLY",
    currency: "Nakfa",
    planId: "<id>",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPostSubscriptions } from "@flexprice/sdk/funcs/subscriptionsPostSubscriptions.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPostSubscriptions(flexPrice, {
    billingCadence: "RECURRING",
    billingPeriod: "HALF_YEARLY",
    currency: "Nakfa",
    planId: "<id>",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPostSubscriptions failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [components.DtoCreateSubscriptionRequest](../../models/components/dtocreatesubscriptionrequest.md)                                                                             | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSubscriptionResponse](../../models/components/dtosubscriptionresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postSubscriptionsAddon

Add an addon to a subscription

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/subscriptions/addon" method="post" path="/subscriptions/addon" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.postSubscriptionsAddon({
    addonId: "<id>",
    subscriptionId: "<id>",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPostSubscriptionsAddon } from "@flexprice/sdk/funcs/subscriptionsPostSubscriptionsAddon.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPostSubscriptionsAddon(flexPrice, {
    addonId: "<id>",
    subscriptionId: "<id>",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPostSubscriptionsAddon failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [components.DtoAddAddonRequest](../../models/components/dtoaddaddonrequest.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoAddonAssociationResponse](../../models/components/dtoaddonassociationresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## deleteSubscriptionsAddon

Remove an addon from a subscription

### Example Usage

<!-- UsageSnippet language="typescript" operationID="delete_/subscriptions/addon" method="delete" path="/subscriptions/addon" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.deleteSubscriptionsAddon({
    addonAssociationId: "<id>",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsDeleteSubscriptionsAddon } from "@flexprice/sdk/funcs/subscriptionsDeleteSubscriptionsAddon.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsDeleteSubscriptionsAddon(flexPrice, {
    addonAssociationId: "<id>",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsDeleteSubscriptionsAddon failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [components.DtoRemoveAddonRequest](../../models/components/dtoremoveaddonrequest.md)                                                                                           | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSuccessResponse](../../models/components/dtosuccessresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## putSubscriptionsLineitemsId

Update a subscription line item by terminating the existing one and creating a new one

### Example Usage

<!-- UsageSnippet language="typescript" operationID="put_/subscriptions/lineitems/{id}" method="put" path="/subscriptions/lineitems/{id}" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.putSubscriptionsLineitemsId("<id>", {});

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPutSubscriptionsLineitemsId } from "@flexprice/sdk/funcs/subscriptionsPutSubscriptionsLineitemsId.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPutSubscriptionsLineitemsId(flexPrice, "<id>", {});
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPutSubscriptionsLineitemsId failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Line Item ID                                                                                                                                                                   |
| `body`                                                                                                                                                                         | [components.DtoUpdateSubscriptionLineItemRequest](../../models/components/dtoupdatesubscriptionlineitemrequest.md)                                                             | :heavy_check_mark:                                                                                                                                                             | Update Line Item Request                                                                                                                                                       |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSubscriptionLineItemResponse](../../models/components/dtosubscriptionlineitemresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## deleteSubscriptionsLineitemsId

Delete a subscription line item by setting its end date

### Example Usage

<!-- UsageSnippet language="typescript" operationID="delete_/subscriptions/lineitems/{id}" method="delete" path="/subscriptions/lineitems/{id}" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.deleteSubscriptionsLineitemsId("<id>", {});

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsDeleteSubscriptionsLineitemsId } from "@flexprice/sdk/funcs/subscriptionsDeleteSubscriptionsLineitemsId.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsDeleteSubscriptionsLineitemsId(flexPrice, "<id>", {});
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsDeleteSubscriptionsLineitemsId failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Line Item ID                                                                                                                                                                   |
| `body`                                                                                                                                                                         | [components.DtoDeleteSubscriptionLineItemRequest](../../models/components/dtodeletesubscriptionlineitemrequest.md)                                                             | :heavy_check_mark:                                                                                                                                                             | Delete Line Item Request                                                                                                                                                       |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSubscriptionLineItemResponse](../../models/components/dtosubscriptionlineitemresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postSubscriptionsSearch

List subscriptions by filter

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/subscriptions/search" method="post" path="/subscriptions/search" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.postSubscriptionsSearch({});

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPostSubscriptionsSearch } from "@flexprice/sdk/funcs/subscriptionsPostSubscriptionsSearch.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPostSubscriptionsSearch(flexPrice, {});
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPostSubscriptionsSearch failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [components.TypesSubscriptionFilter](../../models/components/typessubscriptionfilter.md)                                                                                       | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoListSubscriptionsResponse](../../models/components/dtolistsubscriptionsresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postSubscriptionsUsage

Get usage for a subscription

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/subscriptions/usage" method="post" path="/subscriptions/usage" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.postSubscriptionsUsage({
    endTime: "2024-03-20T00:00:00Z",
    lifetimeUsage: false,
    startTime: "2024-03-13T00:00:00Z",
    subscriptionId: "123",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPostSubscriptionsUsage } from "@flexprice/sdk/funcs/subscriptionsPostSubscriptionsUsage.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPostSubscriptionsUsage(flexPrice, {
    endTime: "2024-03-20T00:00:00Z",
    lifetimeUsage: false,
    startTime: "2024-03-13T00:00:00Z",
    subscriptionId: "123",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPostSubscriptionsUsage failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                                                                                      | [components.DtoGetUsageBySubscriptionRequest](../../models/components/dtogetusagebysubscriptionrequest.md)                                                                     | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoGetUsageBySubscriptionResponse](../../models/components/dtogetusagebysubscriptionresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## getSubscriptionsId

Get a subscription by ID

### Example Usage

<!-- UsageSnippet language="typescript" operationID="get_/subscriptions/{id}" method="get" path="/subscriptions/{id}" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.getSubscriptionsId("<id>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsGetSubscriptionsId } from "@flexprice/sdk/funcs/subscriptionsGetSubscriptionsId.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsGetSubscriptionsId(flexPrice, "<id>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsGetSubscriptionsId failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSubscriptionResponse](../../models/components/dtosubscriptionresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postSubscriptionsIdActivate

Activate a draft subscription with a new start date

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/subscriptions/{id}/activate" method="post" path="/subscriptions/{id}/activate" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.postSubscriptionsIdActivate("<id>", {
    startDate: "<value>",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPostSubscriptionsIdActivate } from "@flexprice/sdk/funcs/subscriptionsPostSubscriptionsIdActivate.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPostSubscriptionsIdActivate(flexPrice, "<id>", {
    startDate: "<value>",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPostSubscriptionsIdActivate failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `body`                                                                                                                                                                         | [components.DtoActivateDraftSubscriptionRequest](../../models/components/dtoactivatedraftsubscriptionrequest.md)                                                               | :heavy_check_mark:                                                                                                                                                             | Activate Draft Subscription Request                                                                                                                                            |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSubscriptionResponse](../../models/components/dtosubscriptionresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## getSubscriptionsIdAddonsAssociations

Get active addon associations for a subscription

### Example Usage

<!-- UsageSnippet language="typescript" operationID="get_/subscriptions/{id}/addons/associations" method="get" path="/subscriptions/{id}/addons/associations" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.getSubscriptionsIdAddonsAssociations("<id>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsGetSubscriptionsIdAddonsAssociations } from "@flexprice/sdk/funcs/subscriptionsGetSubscriptionsIdAddonsAssociations.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsGetSubscriptionsIdAddonsAssociations(flexPrice, "<id>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsGetSubscriptionsIdAddonsAssociations failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoAddonAssociationResponse[]](../../models/.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postSubscriptionsIdCancel

Cancel a subscription with enhanced proration support

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/subscriptions/{id}/cancel" method="post" path="/subscriptions/{id}/cancel" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.postSubscriptionsIdCancel("<id>", {
    cancellationType: "end_of_period",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPostSubscriptionsIdCancel } from "@flexprice/sdk/funcs/subscriptionsPostSubscriptionsIdCancel.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPostSubscriptionsIdCancel(flexPrice, "<id>", {
    cancellationType: "end_of_period",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPostSubscriptionsIdCancel failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `body`                                                                                                                                                                         | [components.DtoCancelSubscriptionRequest](../../models/components/dtocancelsubscriptionrequest.md)                                                                             | :heavy_check_mark:                                                                                                                                                             | Cancel Subscription Request                                                                                                                                                    |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoCancelSubscriptionResponse](../../models/components/dtocancelsubscriptionresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400                          | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postSubscriptionsIdChangeExecute

Execute a subscription plan change, including proration and invoice generation

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/subscriptions/{id}/change/execute" method="post" path="/subscriptions/{id}/change/execute" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.postSubscriptionsIdChangeExecute("<id>", {
    billingCadence: "ONETIME",
    billingCycle: "anniversary",
    billingPeriod: "HALF_YEARLY",
    prorationBehavior: "none",
    targetPlanId: "<id>",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPostSubscriptionsIdChangeExecute } from "@flexprice/sdk/funcs/subscriptionsPostSubscriptionsIdChangeExecute.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPostSubscriptionsIdChangeExecute(flexPrice, "<id>", {
    billingCadence: "ONETIME",
    billingCycle: "anniversary",
    billingPeriod: "HALF_YEARLY",
    prorationBehavior: "none",
    targetPlanId: "<id>",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPostSubscriptionsIdChangeExecute failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `body`                                                                                                                                                                         | [components.DtoSubscriptionChangeRequest](../../models/components/dtosubscriptionchangerequest.md)                                                                             | :heavy_check_mark:                                                                                                                                                             | Subscription change request                                                                                                                                                    |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSubscriptionChangeExecuteResponse](../../models/components/dtosubscriptionchangeexecuteresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postSubscriptionsIdChangePreview

Preview the impact of changing a subscription's plan, including proration calculations

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/subscriptions/{id}/change/preview" method="post" path="/subscriptions/{id}/change/preview" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.postSubscriptionsIdChangePreview("<id>", {
    billingCadence: "RECURRING",
    billingCycle: "anniversary",
    billingPeriod: "WEEKLY",
    prorationBehavior: "create_prorations",
    targetPlanId: "<id>",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPostSubscriptionsIdChangePreview } from "@flexprice/sdk/funcs/subscriptionsPostSubscriptionsIdChangePreview.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPostSubscriptionsIdChangePreview(flexPrice, "<id>", {
    billingCadence: "RECURRING",
    billingCycle: "anniversary",
    billingPeriod: "WEEKLY",
    prorationBehavior: "create_prorations",
    targetPlanId: "<id>",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPostSubscriptionsIdChangePreview failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `body`                                                                                                                                                                         | [components.DtoSubscriptionChangeRequest](../../models/components/dtosubscriptionchangerequest.md)                                                                             | :heavy_check_mark:                                                                                                                                                             | Subscription change preview request                                                                                                                                            |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSubscriptionChangePreviewResponse](../../models/components/dtosubscriptionchangepreviewresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## getSubscriptionsIdEntitlements

Get all entitlements for a subscription

### Example Usage

<!-- UsageSnippet language="typescript" operationID="get_/subscriptions/{id}/entitlements" method="get" path="/subscriptions/{id}/entitlements" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.getSubscriptionsIdEntitlements("<id>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsGetSubscriptionsIdEntitlements } from "@flexprice/sdk/funcs/subscriptionsGetSubscriptionsIdEntitlements.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsGetSubscriptionsIdEntitlements(flexPrice, "<id>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsGetSubscriptionsIdEntitlements failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `featureIds`                                                                                                                                                                   | *string*[]                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                             | Feature IDs to filter by                                                                                                                                                       |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSubscriptionEntitlementsResponse](../../models/components/dtosubscriptionentitlementsresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## getSubscriptionsIdGrantsUpcoming

Get upcoming credit grant applications for a subscription

### Example Usage

<!-- UsageSnippet language="typescript" operationID="get_/subscriptions/{id}/grants/upcoming" method="get" path="/subscriptions/{id}/grants/upcoming" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.getSubscriptionsIdGrantsUpcoming("<id>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsGetSubscriptionsIdGrantsUpcoming } from "@flexprice/sdk/funcs/subscriptionsGetSubscriptionsIdGrantsUpcoming.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsGetSubscriptionsIdGrantsUpcoming(flexPrice, "<id>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsGetSubscriptionsIdGrantsUpcoming failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoListCreditGrantApplicationsResponse](../../models/components/dtolistcreditgrantapplicationsresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postSubscriptionsIdPause

Pause a subscription with the specified parameters

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/subscriptions/{id}/pause" method="post" path="/subscriptions/{id}/pause" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.postSubscriptionsIdPause("<id>", {
    pauseMode: "period_end",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPostSubscriptionsIdPause } from "@flexprice/sdk/funcs/subscriptionsPostSubscriptionsIdPause.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPostSubscriptionsIdPause(flexPrice, "<id>", {
    pauseMode: "period_end",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPostSubscriptionsIdPause failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `body`                                                                                                                                                                         | [components.DtoPauseSubscriptionRequest](../../models/components/dtopausesubscriptionrequest.md)                                                                               | :heavy_check_mark:                                                                                                                                                             | Pause subscription request                                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSubscriptionPauseResponse](../../models/components/dtosubscriptionpauseresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## getSubscriptionsIdPauses

List all pauses for a subscription

### Example Usage

<!-- UsageSnippet language="typescript" operationID="get_/subscriptions/{id}/pauses" method="get" path="/subscriptions/{id}/pauses" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.getSubscriptionsIdPauses("<id>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsGetSubscriptionsIdPauses } from "@flexprice/sdk/funcs/subscriptionsGetSubscriptionsIdPauses.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsGetSubscriptionsIdPauses(flexPrice, "<id>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsGetSubscriptionsIdPauses failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoListSubscriptionPausesResponse[]](../../models/.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postSubscriptionsIdResume

Resume a paused subscription with the specified parameters

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/subscriptions/{id}/resume" method="post" path="/subscriptions/{id}/resume" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.subscriptions.postSubscriptionsIdResume("<id>", {
    resumeMode: "scheduled",
  });

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { subscriptionsPostSubscriptionsIdResume } from "@flexprice/sdk/funcs/subscriptionsPostSubscriptionsIdResume.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await subscriptionsPostSubscriptionsIdResume(flexPrice, "<id>", {
    resumeMode: "scheduled",
  });
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("subscriptionsPostSubscriptionsIdResume failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`                                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Subscription ID                                                                                                                                                                |
| `body`                                                                                                                                                                         | [components.DtoResumeSubscriptionRequest](../../models/components/dtoresumesubscriptionrequest.md)                                                                             | :heavy_check_mark:                                                                                                                                                             | Resume subscription request                                                                                                                                                    |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[components.DtoSubscriptionPauseResponse](../../models/components/dtosubscriptionpauseresponse.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.ErrorsErrorResponse   | 400, 404                     | application/json             |
| errors.ErrorsErrorResponse   | 500                          | application/json             |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |