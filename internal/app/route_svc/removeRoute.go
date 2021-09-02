package routeSvc

import (
	"context"
	desc "ova_route_api/pkg/api/github.com/ozonva/ova-route-api/pkg/ova-route-api"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
)

func (a *RouteAPI) RemoveRoute(ctx context.Context, req *desc.CreateRouteRequest) (*empty.Empty, error) {
	log.Info().Msg("RemoveRoute request")
	return &emptypb.Empty{}, nil
}
