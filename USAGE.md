<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	gosdk "github.com/flexprice/go-sdk"
	"github.com/flexprice/go-sdk/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := gosdk.New(
		"https://api.example.com",
		gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Addons.GetAddons(ctx, operations.GetAddonsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->