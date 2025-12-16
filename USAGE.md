<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"log"
)

func main() {
	ctx := context.Background()

	s := gosdk.New(
		"https://api.example.com",
		gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Addons.List(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.DtoListAddonsResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->