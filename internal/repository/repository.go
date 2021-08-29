package repository

import "ova_route_api/internal/models"

// Repo - интерфейс хранилища для сущности Entity
type Repo interface {
	AddEntities(entities []models.Route) error
	ListEntities(limit, offset uint64) ([]models.Route, error)
	DescribeEntity(entityId uint64) (*models.Route, error)
}
