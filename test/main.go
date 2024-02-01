package main

import (
	"time"

	"github.com/eonianmonk/tollbooth_fasthttp"

	"github.com/didip/tollbooth"

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

	// convert 1 req per hour to max req per second
	max := converReqLimitToLimitPerSeconds(time.Hour, 1)
	// Create a limiter struct.
	limiter := tollbooth.NewLimiter(max, nil)
	limiter.SetTokenBucketExpirationTTL(time.Hour * 2)
	fasthttp.ListenAndServe(":4444", tollbooth_fasthttp.LimitHandler(requestHandler, limiter))
}

// converts number of allowed requests per timespan to allowed requests per second
func converReqLimitToLimitPerSeconds(timespan time.Duration, allowedRequests int) float64 {
	return float64(allowedRequests) / float64(timespan/time.Second)
}

func helloHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Hello, World!"))
}
