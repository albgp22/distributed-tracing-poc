package myHttp

import (
	"context"
	"distributedTracing/lib/tracing"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"io/ioutil"
	"net/http"
)

func PerformRequest(ctx context.Context, url string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "send-request-type-a")
	defer span.Finish()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	if err := tracing.Inject(span, req); err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}
	return string(body), nil
}
