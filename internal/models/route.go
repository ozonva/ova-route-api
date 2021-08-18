package route

import "errors"

var id uint64

type Route struct {
	ID        uint64  // Первичный ключь
	UserID    uint64  // Обязательный .. вторичный ключь для связи с пользователем
	RouteName string  // Название маршрута
	Length    float64 // Протяженность маршрута
}

func New(userID uint64, routeName string, length float64) (Route, error) {
	id += 1 // костылик, генерим уникальные id
	if length < 0 {
		return Route{}, errors.New("route lenght must be greater than zero")
	}

	route := Route{
		ID:        id,
		UserID:    userID,
		RouteName: routeName,
		Length:    length,
	}

	return route, nil
}

func (r *Route) RouteTime(route Route, speed float64) (float64, error) {
	if speed <= 0 {
		return 0, errors.New("speed must be greater than zero")
	}

	return route.Length / speed, nil
}
