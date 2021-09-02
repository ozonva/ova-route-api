package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"ova_route_api/build"
	"ova_route_api/config"
	api "ova_route_api/internal/app/route_svc"
	route "ova_route_api/internal/models"
	"syscall"

	desc "ova_route_api/pkg/api/github.com/ozonva/ova-route-api/pkg/ova-route-api"

	"google.golang.org/grpc"
)

var Version = "development"

const (
	grpcPort           = ":8082"
	grpcServerEndpoint = "localhost:8082"
)

func main() {
	fmt.Println("Version:\t", Version)
	fmt.Println("build.Time:\t", build.Time)
	fmt.Println("build.User:\t", build.User)

	if err := run(); err != nil {
		log.Fatal(err)
	}

	cfg := config.Getconfig()
	fmt.Printf("Пользователи: %v. Группы: %v.\n", cfg.Users, cfg.Groups)

	route, err := route.New(1, "Маршрут для бега", 4)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("Создан маршрут: %v. ID пользователя: %v. Протяженность: %v км\n", route.RouteName, route.UserID, route.Length)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("closing")
}

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterRouteServer(s, api.NewRouteAPI())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
