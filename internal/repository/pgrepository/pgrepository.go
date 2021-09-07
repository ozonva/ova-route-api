package pgrepository

import (
	"context"
	"database/sql"
	"ova-route-api/internal/models"
	"time"

	"github.com/rs/zerolog"

	_ "github.com/jackc/pgx/stdlib"
)

type repository struct {
	db     *sql.DB
	logger zerolog.Logger
}

func New(logger zerolog.Logger) repository {
	// TODO вынсти в конфиг
	dsn := "postgres://ozon_user:secret@localhost:49153/ozon?sslmode=disable"
	// dsn := fmt.Sprintf("user=ozon_user dbname=ozon sslmode=disable password=secret port=49154",
	// "ozon_user", )
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Fatal().Msgf("failed to load driver: %v", err)
	}

	// err = db.PingContext(ctx)
	err = db.Ping()
	if err != nil {
		logger.Fatal().Msgf("failed to connect to db: %v", err)
	}

	// Макс. число открытых соединений от этого процесса
	db.SetMaxOpenConns(10)
	// Макс. число открытых неиспользуемых соединений
	db.SetMaxIdleConns(1)
	// Макс. время жизни одного подключения
	db.SetConnMaxLifetime(60 * time.Minute)

	return repository{
		db:     db,
		logger: logger,
	}
}

func (repo repository) AddRoute(route models.Route) (models.Route, error) {
	query := "INSERT INTO routes (user_id, route_name, length) VALUES ($1, $2, $3) RETURNING id"
	err := repo.db.QueryRow(query, route.UserID, route.RouteName, route.Length).Scan(&route.ID)
	if err != nil {
		repo.logger.Error().Msgf("err create route: %v", err)
		return models.Route{}, err
	}

	return route, nil
}

func (repo repository) AddRoutes(routes []models.Route) error {
	ctx := context.TODO()
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		repo.logger.Error().Msgf("cant access connection: %v", err)
		return err
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "INSERT INTO routes (user_id, route_name, length) VALUES ($1, $2, $3)")
	if err != nil {
		repo.logger.Error().Msgf("err prepare request: %v", err)
		return err
	}
	defer stmt.Close()

	for _, route := range routes {
		_, err := stmt.Exec(route.UserID, route.RouteName, route.Length)
		if err != nil {
			repo.logger.Error().Msgf("err exec request: %v", err)
			return err
		}
	}

	return nil
}

func (repo repository) DescribeRoute(route models.Route) (models.Route, error) {
	query := "UPDATE routes SET user_id = $1, route_name = $2, length = $3"
	// _, err := db.ExecContext(ctx, query, , "new year", "watch")
	// repo.logger.Error().Msgf("UserID: %v, RouteName: %v, Length: %v", route.UserID, route.RouteName, route.Length)
	_, err := repo.db.Exec(query, route.UserID, route.RouteName, route.Length)
	if err != nil {
		repo.logger.Error().Msgf("err update route: %v", err)
		return models.Route{}, err
	}

	return route, nil
}

func (repo repository) ListRoutes(limit, offset uint64) ([]models.Route, error) {
	var resp []models.Route
	// var rowsCount uint64
	// err := repo.db.QueryRow("SELECT count(*) FROM routes").Scan(&rowsCount)
	// if err != nil {
	// 	repo.logger.Error().Msgf("err getting route: %v", err)
	// 	return []models.Route{}, err
	// }

	rows, err := repo.db.Query(`SELECT * FROM routes order by 
		id desc limit $1 offset $2`, limit, offset)
	if err != nil {
		repo.logger.Error().Msgf("err getting route: %v", err)
		return []models.Route{}, err
	}
	defer rows.Close()

	for rows.Next() {
		route := models.Route{}
		err = rows.Scan(&route.ID, &route.UserID, &route.RouteName, &route.Length)
		if err != nil {
			repo.logger.Error().Msgf("err read row: %v", err)
			continue
		}
		resp = append(resp, route)
	}

	return resp, nil
}

func (repo repository) RemoveRoute(routeID uint64) error {
	_, err := repo.db.Exec("DELETE FROM routes WHERE id = $1", routeID)
	if err != nil {
		if err != nil {
			repo.logger.Error().Msgf("err delete row: %v", err)
			return err
		}
	}
	return nil
}
