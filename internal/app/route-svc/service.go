package routesvc

import (
	"context"
	"ova-route-api/internal/broker"
	"ova-route-api/internal/flusher"
	"ova-route-api/internal/models"
	"ova-route-api/internal/repository"
	"ova-route-api/internal/saver"

	desc "ova-route-api/pkg/api/github.com/ozonva/ova-route-api/pkg/ova-route-api"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
)

type RouteAPI struct {
	desc.RouteServer
	tracer     opentracing.Tracer
	logger     zerolog.Logger
	repository repository.Repo
	callCount  *prometheus.CounterVec
	producer   broker.Producer
	saver      saver.Saver
}

func NewRouteAPI(
	logger zerolog.Logger,
	repository repository.Repo,
	producer broker.Producer,
	callCount *prometheus.CounterVec,
) desc.RouteServer {
	return &RouteAPI{
		tracer:     opentracing.GlobalTracer(),
		logger:     logger,
		repository: repository,
		callCount:  callCount,
		producer:   producer,
		saver:      saver.NewSaver(1, flusher.NewFlusher(1, repository)),
	}
}

// CREATE
func (api *RouteAPI) CreateRoute(ctx context.Context, req *desc.CreateRouteRequest) (*empty.Empty, error) {
	api.logger.Info().Msg("CreateRoute request")
	api.callCount.WithLabelValues("CREATE").Add(1)
	api.producer.Produce(ctx, broker.RouteCreated)

	route := models.Route{
		UserID:    req.UserID,
		RouteName: req.RouteName,
		Length:    req.Length,
	}

	route, err := api.repository.AddRoute(route)
	if err != nil {
		// Handle
	}

	return &emptypb.Empty{}, nil
}

// MULTI CREATE
func (api *RouteAPI) MultiCreateRoute(ctx context.Context, req *desc.MultiCreateRouteRequest) (*empty.Empty, error) {
	api.logger.Info().Msg("MultiCreateRoute request")

	parentSpan := api.tracer.StartSpan("parent")
	ctx = opentracing.ContextWithSpan(ctx, parentSpan)

	for _, route := range req.Items {
		route := models.Route{
			UserID:    route.UserID,
			RouteName: route.RouteName,
			Length:    route.Length,
		}

		api.saver.Save(ctx, route)
	}

	return &emptypb.Empty{}, nil
}

// UPDATE
func (api *RouteAPI) DescribeRoute(ctx context.Context, req *desc.DescribeRouteRequest) (*desc.RouteResponse, error) {
	api.logger.Info().Msg("DescribeRoute request")
	api.callCount.WithLabelValues("UPDATE").Add(1)
	api.producer.Produce(ctx, broker.RouteUpdated)

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

// DELETE
func (api *RouteAPI) RemoveRoute(ctx context.Context, req *desc.RemoveRouteRequest) (*empty.Empty, error) {
	api.logger.Info().Msg("RemoveRoute request")
	api.callCount.WithLabelValues("DELETE").Add(1)
	api.producer.Produce(ctx, broker.RouteRemoved)

	err := api.repository.RemoveRoute(req.ID)
	if err != nil {
		// Handle
	}

	return &emptypb.Empty{}, nil
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
