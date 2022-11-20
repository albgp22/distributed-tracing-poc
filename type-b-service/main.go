package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"distributedTracing/lib/tracing"

	"github.com/opentracing/opentracing-go"
)

var thisServiceName = ""

func main() {
	ins := os.Getenv("INSTANCE_NO")
	thisServiceName = fmt.Sprintf("type-b-service-%s", ins)

	tracer, closer := tracing.Init(thisServiceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) {
			span := tracing.StartSpanFromRequest(tracer, r)
			defer span.Finish()

			w.Write([]byte(thisServiceName))
		},
	)
	log.Printf("Listening on localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
