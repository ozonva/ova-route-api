package main

import (
	"fmt"
	"os"
	"os/signal"
	"ova_route_api/build"
	"ova_route_api/config"
	route "ova_route_api/internal/models"
	"syscall"
)

var Version = "development"

func main() {
	fmt.Println("Version:\t", Version)
	fmt.Println("build.Time:\t", build.Time)
	fmt.Println("build.User:\t", build.User)

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
