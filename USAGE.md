<!-- Start SDK Example Usage [usage] -->
```typescript
import { FlexPrice } from "@flexprice/sdk";

const flexPrice = new FlexPrice({
  serverURL: "https://api.example.com",
  apiKeyAuth: "<YOUR_API_KEY_HERE>",
});

async function run() {
  const result = await flexPrice.addons.getAddons({});

  console.log(result);
}

run();

```
<!-- End SDK Example Usage [usage] -->