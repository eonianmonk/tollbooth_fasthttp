## tollbooth_fasthttp

[Fasthttp](https://github.com/valyala/fasthttp) middleware for rate limiting HTTP requests.


## Version:

This shim uses `v7` API.


## Five Minutes Tutorial

```
package main

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_fasthttp"
	"github.com/valyala/fasthttp"
)

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/hello":
			helloHandler(ctx)
		default:
			ctx.Error("Unsupporterd path", fasthttp.StatusNotFound)
		}
	}

	// Create a limiter struct.
	limiter := tollbooth.NewLimiter(1, time.Second)

	fasthttp.ListenAndServe(":4444", tollbooth_fasthttp.LimitHandler(requestHandler, limiter))
}

func helloHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Hello, World!"))
}
```