> ⚠️ **Work in Progress**: This library is in early development. I'm starting by implementing the Classic WoW API. Contributions welcome.

# Battle.net Go Client

A simple Go client library for interacting with the Battle.net API.

## Installation

```bash
go get github.com/kjhickman/bnet-client
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/kjhickman/bnet-client"
)

func main() {
	config := client.Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
	}

	client, err := client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	token, err := client.GetWoWToken(context.Background())
	if err != nil {
		log.Fatalf("Failed to get WoW token info: %v", err)
	}

	fmt.Printf("WoW Token Price: %d gold\n", token.Price)
}
```

## Examples

See the [examples](./examples/) directory for usage examples.

## License

This project is [MIT Licensed](LICENSE).