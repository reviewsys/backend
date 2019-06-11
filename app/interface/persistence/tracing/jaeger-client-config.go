package tracing

import (
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	cfg "github.com/reviewsys/backend/config"
	"github.com/sirupsen/logrus"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func NewJaeger(config cfg.Config) (opentracing.Tracer, io.Closer, error) {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		// parsing errors might happen here, such as when we get a string where we expect a number
		logrus.Printf("Could not parse Jaeger env vars: %s", err)
		return nil, nil, err
	}
	serviceName := config.GetString("app.id")

	cfg.ServiceName = serviceName

	return cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
}
