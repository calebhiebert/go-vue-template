package main

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func SetupTracing() (tracer opentracing.Tracer, closer io.Closer, err error) {
	cfg := config.Configuration{
		ServiceName: "go-vue-template",

		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},

		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 15 * time.Second,
		},
	}

	tracer, closer, err = cfg.NewTracer(config.Logger(jaeger.StdLogger))

	opentracing.SetGlobalTracer(tracer)

	return
}
