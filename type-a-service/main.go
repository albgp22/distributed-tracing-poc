package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"distributedTracing/lib/myHttp"
	"distributedTracing/lib/tracing"

	"github.com/opentracing/opentracing-go"
)

const thisServiceName = "type-a-service"

func main() {
	tracer, closer := tracing.Init(thisServiceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) {
			span := tracing.StartSpanFromRequest(tracer, r)
			defer span.Finish()

			ctx := opentracing.ContextWithSpan(context.Background(), span)
			hostname := fmt.Sprintf("http://type-b-service-%d:8082/hello", rand.Intn(2))

			res, err := myHttp.PerformRequest(ctx, hostname)

			if err != nil {
				log.Fatalf("Error occurred when hellowing %s: %s", hostname, err)
			}
			w.Write([]byte(fmt.Sprintf("%s -> %s", thisServiceName, res)))
		},
	)
	log.Printf("Listening on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
