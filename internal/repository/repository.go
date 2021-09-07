package repository

import "ova-route-api/internal/models"

// Repo - интерфейс хранилища для сущности Entity
type Repo interface {
	AddRoute(route models.Route) (models.Route, error)
	AddRoutes(routes []models.Route) error
	DescribeRoute(route models.Route) (models.Route, error)
	ListRoutes(limit, offset uint64) ([]models.Route, error)
	RemoveRoute(routeID uint64) error
}
