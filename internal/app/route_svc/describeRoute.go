package routeSvc

import (
	"context"
	desc "ova_route_api/pkg/api/github.com/ozonva/ova-route-api/pkg/ova-route-api"

	"github.com/rs/zerolog/log"
)

func (a *RouteAPI) DescribeRoute(ctx context.Context, req *desc.CreateRouteRequest) (*desc.RouteResponse, error) {
	log.Info().Msg("DescribeRoute request")
	return &desc.RouteResponse{
		ID:        1,
		UserID:    req.UserID,
		RouteName: req.GetRouteName(),
		Length:    req.Length,
	}, nil
}
