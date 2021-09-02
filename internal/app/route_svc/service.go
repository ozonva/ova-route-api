package routeSvc

import (
	desc "ova_route_api/pkg/api/github.com/ozonva/ova-route-api/pkg/ova-route-api"
)

type RouteAPI struct {
	desc.RouteServer
}

func NewRouteAPI() desc.RouteServer {
	return &RouteAPI{}
}
