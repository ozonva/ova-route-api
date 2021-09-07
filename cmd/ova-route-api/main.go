package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"ova-route-api/build"
	"ova-route-api/config"
	"ova-route-api/internal/repository/pgrepository"
	"syscall"

	api "ova-route-api/internal/app/route-svc"

	desc "ova-route-api/pkg/api/github.com/ozonva/ova-route-api/pkg/ova-route-api"

	broker "ova-route-api/internal/broker/kafka"

	"github.com/prometheus/client_golang/prometheus"

	opentracing "github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/oklog/run"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var Version = "development"

func main() {
	fmt.Println("Version:\t", Version)
	fmt.Println("build.Time:\t", build.Time)
	fmt.Println("build.User:\t", build.User)

	// Create a single logger, which we'll use and give to other components.
	var logger zerolog.Logger
	{
		logger = log.Level(zerolog.DebugLevel)
		logger = log.Output(os.Stderr)
		logger = log.With().Timestamp().Logger()
		logger = log.With().Caller().Logger()
	}

	logger.Log().Msg("service started")
	defer logger.Log().Msg("service ended")

	// Read connfig file
	cfg := config.Getconfig()

	// Create the (sparse) metrics we'll use in the service. They, too, are
	// dependencies that we pass to components that use them.
	var callCount *prometheus.CounterVec
	{
		callCount = prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ova",
			Subsystem: "route_api",
			Name:      "call_count",
			Help:      "Total count of success call",
		}, []string{"method"})
	}

	prometheus.MustRegister(callCount)

	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	jcfg := jaegercfg.Configuration{
		ServiceName: "ova-route-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Initialize tracer
	tracer, closer, err := jcfg.NewTracer()
	if err != nil {
		logger.Log().Msgf("Initialize tracer error: %v", err)
	}
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	// Build the layers of the service "onion" from the inside out.
	var (
		repository = pgrepository.New(logger)
		broker     = broker.NewProducer(cfg.KafkaTopic, cfg.KafkaAdDress, logger)
		routeSVC   = api.NewRouteAPI(logger, repository, broker, callCount)
	)

	// Putting each component into its own block is mostly for aesthetics: it
	// clearly demarcates the scope in which each listener/socket may be used.
	var g run.Group
	{
		debugListener, err := net.Listen("tcp", cfg.MetrikAddr)
		if err != nil {
			logger.Log().Msgf("transport: %v,  during: %v, err: %v", "metrik/HTTP", "Listen", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log().Msgf("transport: %v,  addr: %v", "metrik/HTTP", cfg.MetrikAddr)
			return http.Serve(debugListener, http.DefaultServeMux)
		}, func(error) {
			debugListener.Close()
		})
	}
	{
		listen, err := net.Listen("tcp", cfg.GRPCAddr)
		if err != nil {
			logger.Fatal().Msgf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		desc.RegisterRouteServer(s, routeSVC)

		g.Add(func() error {
			defer logger.Fatal().Msgf("failed to serve: %v", err)
			return s.Serve(listen)
		}, func(error) {
			listen.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	logger.Log().Msgf("The group was terminated with %v", g.Run())
}
