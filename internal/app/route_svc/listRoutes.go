package routeSvc

import (
	"context"
	desc "ova_route_api/pkg/api/github.com/ozonva/ova-route-api/pkg/ova-route-api"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
)

func (a *RouteAPI) ListRoutes(ctx context.Context, req *empty.Empty) (*desc.ListRoutesResponse, error) {
	log.Info().Msg("ListRoutes request")
	return &desc.ListRoutesResponse{
		Items: []*desc.RouteResponse{
			{
				ID:        1,
				UserID:    1,
				RouteName: "name",
				Length:    1.5,
			},
		},
	}, nil
}
