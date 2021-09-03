package flusher

import (
	"ova-route-api/internal/models"
	"ova-route-api/internal/repository"
	"ova-route-api/internal/utils"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(entities []models.Route) []models.Route
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(chunkSize int, entityRepo repository.Repo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize  int
	entityRepo repository.Repo
}

func (f flusher) Flush(routes []models.Route) []models.Route {
	var resp []models.Route

	bulks, err := utils.SplitToBulks(routes, uint(f.chunkSize))
	if err != nil {
		// handle error
		return routes
	}

	for _, v := range bulks {
		err = f.entityRepo.AddRoutes(v)
		if err != nil {
			// handle error
			resp = append(resp, v...)
		}
	}

	return resp
}
