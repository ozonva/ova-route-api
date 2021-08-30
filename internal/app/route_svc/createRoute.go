package routeSvc

import (
	"context"
	desc "ova_route_api/pkg/api/github.com/ozonva/ova-route-api/pkg/ova-route-api"

	"github.com/rs/zerolog/log"
)

func (a *RouteAPI) CreateRoute(ctx context.Context, req *desc.CreateRouteRequest) (*desc.RouteResponse, error) {
	log.Info().Msg("CreateConference request")
	return &desc.RouteResponse{
		ID:        1,
		UserID:    1,
		RouteName: "name",
		Length:    1.5,
	}, nil
}
