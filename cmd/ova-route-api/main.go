package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"ova-route-api/build"
	"ova-route-api/config"
	"ova-route-api/internal/repository/pgrepository"
	"syscall"

	api "ova-route-api/internal/app/route-svc"

	desc "ova-route-api/pkg/api/github.com/ozonva/ova-route-api/pkg/ova-route-api"

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

	// Build the layers of the service "onion" from the inside out.
	repository := pgrepository.New(logger)

	// Putting each component into its own block is mostly for aesthetics: it
	// clearly demarcates the scope in which each listener/socket may be used.
	var g run.Group
	{
		listen, err := net.Listen("tcp", cfg.GRPCAdr)
		if err != nil {
			logger.Fatal().Msgf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		desc.RegisterRouteServer(s, api.NewRouteAPI(logger, repository))

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

// func run() error {
// 	listen, err := net.Listen("tcp", grpcPort)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	desc.RegisterLecture6DemoServer(s, api.NewLecture6DemoAPI())
// 	if err := s.Serve(listen); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// 	return nil
// }
