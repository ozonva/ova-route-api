package flusher

import (
	"context"
	"ova-route-api/internal/models"
	"ova-route-api/internal/repository"
	"ova-route-api/internal/utils"

	"github.com/opentracing/opentracing-go"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(ctx context.Context, routes []models.Route) []models.Route
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(chunkSize int, entityRepo repository.Repo) Flusher {
	return &flusher{
		tracer:     opentracing.GlobalTracer(),
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	tracer     opentracing.Tracer
	chunkSize  int
	entityRepo repository.Repo
}

func (f flusher) Flush(ctx context.Context, routes []models.Route) []models.Route {
	var resp []models.Route

	bulks, err := utils.SplitToBulks(routes, uint(f.chunkSize))
	if err != nil {
		// handle error
		return routes
	}

	span := opentracing.SpanFromContext(ctx)

	for _, v := range bulks {
		childSpan := f.tracer.StartSpan(
			"child",
			opentracing.Tag{Key: "BulkCount", Value: len(v)},
			opentracing.ChildOf(span.Context()),
		)
		err = f.entityRepo.AddRoutes(v)

		if err != nil {
			// handle error
			resp = append(resp, v...)
		}

		childSpan.Finish()
	}

	span.Finish()

	return resp
}
