# Webhooks

## Overview

### Available Operations

* [postWebhooksChargebeeTenantIdEnvironmentId](#postwebhookschargebeetenantidenvironmentid) - Handle Chargebee webhook events
* [postWebhooksHubspotTenantIdEnvironmentId](#postwebhookshubspottenantidenvironmentid) - Handle HubSpot webhook events
* [postWebhooksNomodTenantIdEnvironmentId](#postwebhooksnomodtenantidenvironmentid) - Handle Nomod webhook events
* [postWebhooksQuickbooksTenantIdEnvironmentId](#postwebhooksquickbookstenantidenvironmentid) - Handle QuickBooks webhook events
* [postWebhooksRazorpayTenantIdEnvironmentId](#postwebhooksrazorpaytenantidenvironmentid) - Handle Razorpay webhook events
* [postWebhooksStripeTenantIdEnvironmentId](#postwebhooksstripetenantidenvironmentid) - Handle Stripe webhook events

## postWebhooksChargebeeTenantIdEnvironmentId

Process incoming Chargebee webhook events for payment status updates

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/webhooks/chargebee/{tenant_id}/{environment_id}" method="post" path="/webhooks/chargebee/{tenant_id}/{environment_id}" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.webhooks.postWebhooksChargebeeTenantIdEnvironmentId("<id>", "<id>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { webhooksPostWebhooksChargebeeTenantIdEnvironmentId } from "@flexprice/sdk/funcs/webhooksPostWebhooksChargebeeTenantIdEnvironmentId.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await webhooksPostWebhooksChargebeeTenantIdEnvironmentId(flexPrice, "<id>", "<id>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("webhooksPostWebhooksChargebeeTenantIdEnvironmentId failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `tenantId`                                                                                                                                                                     | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Tenant ID                                                                                                                                                                      |
| `environmentId`                                                                                                                                                                | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Environment ID                                                                                                                                                                 |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[{ [k: string]: any }](../../models/.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postWebhooksHubspotTenantIdEnvironmentId

Process incoming HubSpot webhook events for deal closed won and customer creation

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/webhooks/hubspot/{tenant_id}/{environment_id}" method="post" path="/webhooks/hubspot/{tenant_id}/{environment_id}" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.webhooks.postWebhooksHubspotTenantIdEnvironmentId("<id>", "<id>", "<value>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { webhooksPostWebhooksHubspotTenantIdEnvironmentId } from "@flexprice/sdk/funcs/webhooksPostWebhooksHubspotTenantIdEnvironmentId.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await webhooksPostWebhooksHubspotTenantIdEnvironmentId(flexPrice, "<id>", "<id>", "<value>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("webhooksPostWebhooksHubspotTenantIdEnvironmentId failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `tenantId`                                                                                                                                                                     | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Tenant ID                                                                                                                                                                      |
| `environmentId`                                                                                                                                                                | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Environment ID                                                                                                                                                                 |
| `xHubSpotSignatureV3`                                                                                                                                                          | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | HubSpot webhook signature                                                                                                                                                      |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[{ [k: string]: any }](../../models/.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postWebhooksNomodTenantIdEnvironmentId

Process incoming Nomod webhook events for payment and invoice payments

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/webhooks/nomod/{tenant_id}/{environment_id}" method="post" path="/webhooks/nomod/{tenant_id}/{environment_id}" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.webhooks.postWebhooksNomodTenantIdEnvironmentId("<id>", "<id>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { webhooksPostWebhooksNomodTenantIdEnvironmentId } from "@flexprice/sdk/funcs/webhooksPostWebhooksNomodTenantIdEnvironmentId.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await webhooksPostWebhooksNomodTenantIdEnvironmentId(flexPrice, "<id>", "<id>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("webhooksPostWebhooksNomodTenantIdEnvironmentId failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `tenantId`                                                                                                                                                                     | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Tenant ID                                                                                                                                                                      |
| `environmentId`                                                                                                                                                                | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Environment ID                                                                                                                                                                 |
| `xApiKey`                                                                                                                                                                      | *string*                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                             | Nomod webhook secret (if configured)                                                                                                                                           |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[{ [k: string]: any }](../../models/.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postWebhooksQuickbooksTenantIdEnvironmentId

Process incoming QuickBooks webhook events for payment sync

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/webhooks/quickbooks/{tenant_id}/{environment_id}" method="post" path="/webhooks/quickbooks/{tenant_id}/{environment_id}" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.webhooks.postWebhooksQuickbooksTenantIdEnvironmentId("<id>", "<id>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { webhooksPostWebhooksQuickbooksTenantIdEnvironmentId } from "@flexprice/sdk/funcs/webhooksPostWebhooksQuickbooksTenantIdEnvironmentId.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await webhooksPostWebhooksQuickbooksTenantIdEnvironmentId(flexPrice, "<id>", "<id>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("webhooksPostWebhooksQuickbooksTenantIdEnvironmentId failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `tenantId`                                                                                                                                                                     | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Tenant ID                                                                                                                                                                      |
| `environmentId`                                                                                                                                                                | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Environment ID                                                                                                                                                                 |
| `intuitSignature`                                                                                                                                                              | *string*                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                             | QuickBooks webhook signature                                                                                                                                                   |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[{ [k: string]: any }](../../models/.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postWebhooksRazorpayTenantIdEnvironmentId

Process incoming Razorpay webhook events for payment capture and failure

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/webhooks/razorpay/{tenant_id}/{environment_id}" method="post" path="/webhooks/razorpay/{tenant_id}/{environment_id}" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.webhooks.postWebhooksRazorpayTenantIdEnvironmentId("<id>", "<id>", "<value>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { webhooksPostWebhooksRazorpayTenantIdEnvironmentId } from "@flexprice/sdk/funcs/webhooksPostWebhooksRazorpayTenantIdEnvironmentId.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await webhooksPostWebhooksRazorpayTenantIdEnvironmentId(flexPrice, "<id>", "<id>", "<value>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("webhooksPostWebhooksRazorpayTenantIdEnvironmentId failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `tenantId`                                                                                                                                                                     | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Tenant ID                                                                                                                                                                      |
| `environmentId`                                                                                                                                                                | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Environment ID                                                                                                                                                                 |
| `xRazorpaySignature`                                                                                                                                                           | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Razorpay webhook signature                                                                                                                                                     |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[{ [k: string]: any }](../../models/.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |

## postWebhooksStripeTenantIdEnvironmentId

Process incoming Stripe webhook events for payment status updates and customer creation

### Example Usage

<!-- UsageSnippet language="typescript" operationID="post_/webhooks/stripe/{tenant_id}/{environment_id}" method="post" path="/webhooks/stripe/{tenant_id}/{environment_id}" -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.webhooks.postWebhooksStripeTenantIdEnvironmentId("<id>", "<id>", "<value>");

  console.log(result);
}

run();
```

### Standalone function

The standalone function version of this method:

```typescript
import { FlexPriceCore } from "@flexprice/sdk/core.js";
import { webhooksPostWebhooksStripeTenantIdEnvironmentId } from "@flexprice/sdk/funcs/webhooksPostWebhooksStripeTenantIdEnvironmentId.js";

// Use `FlexPriceCore` for best tree-shaking performance.
// You can create one instance of it to use across an application.
const flexPrice = new FlexPriceCore({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const res = await webhooksPostWebhooksStripeTenantIdEnvironmentId(flexPrice, "<id>", "<id>", "<value>");
  if (res.ok) {
    const { value: result } = res;
    console.log(result);
  } else {
    console.log("webhooksPostWebhooksStripeTenantIdEnvironmentId failed:", res.error);
  }
}

run();
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `tenantId`                                                                                                                                                                     | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Tenant ID                                                                                                                                                                      |
| `environmentId`                                                                                                                                                                | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Environment ID                                                                                                                                                                 |
| `stripeSignature`                                                                                                                                                              | *string*                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                             | Stripe webhook signature                                                                                                                                                       |
| `options`                                                                                                                                                                      | RequestOptions                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                             | Used to set various options for making HTTP requests.                                                                                                                          |
| `options.fetchOptions`                                                                                                                                                         | [RequestInit](https://developer.mozilla.org/en-US/docs/Web/API/Request/Request#options)                                                                                        | :heavy_minus_sign:                                                                                                                                                             | Options that are passed to the underlying HTTP request. This can be used to inject extra headers for examples. All `Request` options, except `method` and `body`, are allowed. |
| `options.retries`                                                                                                                                                              | [RetryConfig](../../lib/utils/retryconfig.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                             | Enables retrying HTTP requests under certain failure conditions.                                                                                                               |

### Response

**Promise\<[{ [k: string]: any }](../../models/.md)\>**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| errors.FlexPriceDefaultError | 4XX, 5XX                     | \*/\*                        |