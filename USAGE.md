<!-- Start SDK Example Usage [usage] -->
```python
# Synchronous Example
from flexprice import FlexPrice


with FlexPrice(
    server_url="https://api.example.com",
    api_key_auth="<YOUR_API_KEY_HERE>",
) as flex_price:

    res = flex_price.addons.get_addons()

    # Handle response
    print(res)
```

</br>

The same SDK client can also be used to make asynchronous requests by importing asyncio.

```python
# Asynchronous Example
import asyncio
from flexprice import FlexPrice

async def main():

    async with FlexPrice(
        server_url="https://api.example.com",
        api_key_auth="<YOUR_API_KEY_HERE>",
    ) as flex_price:

        res = await flex_price.addons.get_addons_async()

        # Handle response
        print(res)

asyncio.run(main())
```
<!-- End SDK Example Usage [usage] -->