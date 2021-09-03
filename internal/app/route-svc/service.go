package routesvc

import (
	"context"
	"ova-route-api/internal/flusher"
	"ova-route-api/internal/models"
	"ova-route-api/internal/repository"
	"ova-route-api/internal/saver"

	desc "ova-route-api/pkg/api/github.com/ozonva/ova-route-api/pkg/ova-route-api"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog"
)

type RouteAPI struct {
	desc.RouteServer
	logger     zerolog.Logger
	repository repository.Repo
	saver      saver.Saver
}

func NewRouteAPI(logger zerolog.Logger, repository repository.Repo) desc.RouteServer {
	return &RouteAPI{
		logger:     logger,
		repository: repository,
		saver:      saver.NewSaver(1, flusher.NewFlusher(1, repository)),
	}
}

func (api *RouteAPI) CreateRoute(ctx context.Context, req *desc.CreateRouteRequest) (*empty.Empty, error) {
	api.logger.Info().Msg("CreateConference request")

	route := models.Route{
		UserID:    req.UserID,
		RouteName: req.RouteName,
		Length:    req.Length,
	}

	api.saver.Save(route)

	return &emptypb.Empty{}, nil
}

func (api *RouteAPI) DescribeRoute(ctx context.Context, req *desc.DescribeRouteRequest) (*desc.RouteResponse, error) {
	api.logger.Info().Msg("DescribeRoute request")

	route := models.Route{
		UserID:    req.UserID,
		RouteName: req.RouteName,
		Length:    req.Length,
	}

	route, err := api.repository.DescribeRoute(route)
	if err != nil {
		// Handle
	}

	resp := desc.RouteResponse{
		ID:        route.ID,
		UserID:    route.UserID,
		RouteName: route.RouteName,
		Length:    route.Length,
	}

	return &resp, nil
}

func (api *RouteAPI) ListRoutes(ctx context.Context, req *desc.ListRoutesRequest) (*desc.ListRoutesResponse, error) {
	api.logger.Info().Msg("ListRoutes request")

	routes, err := api.repository.ListRoutes(req.Limit, req.Offset)
	if err != nil {
		// Handle
	}

	resp := desc.ListRoutesResponse{}
	for _, v := range routes {
		item := desc.RouteResponse{
			ID:        v.ID,
			UserID:    v.UserID,
			RouteName: v.RouteName,
			Length:    v.Length,
		}
		resp.Items = append(resp.Items, &item)
	}

	return &resp, nil
}

func (api *RouteAPI) RemoveRoute(ctx context.Context, req *desc.RemoveRouteRequest) (*empty.Empty, error) {
	api.logger.Info().Msg("RemoveRoute request")

	err := api.repository.RemoveRoute(req.ID)
	if err != nil {
		// Handle
	}

	return &emptypb.Empty{}, nil
}
