# detach [![Go Reference](https://pkg.go.dev/badge/github.com/fxrlv/detach.svg)](https://pkg.go.dev/github.com/fxrlv/detach)

```go
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/fxrlv/detach"
)

func main() {
	main, cancel := signal.NotifyContext(
		context.Background(), os.Interrupt,
	)
	defer cancel()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// For incoming server requests, the context is canceled when the
		// client's connection closes, the request is canceled (with HTTP/2),
		// or when the ServeHTTP method returns.
		ctx := r.Context()

		ctx = detach.WithCancel(ctx, main)
		go func() {
			// Will be cancelled by signal.
			<-ctx.Done()

			// Values are associated with http.Request context.
			server := ctx.Value(http.ServerContextKey).(*http.Server)
			server.Close()
		}()
	})

	http.ListenAndServe(":8080", nil)
}
```
