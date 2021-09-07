package broker

import (
	"context"
)

type EventType string

const (
	RouteCreated EventType = "RouteCreated"
	RouteUpdated EventType = "RouteUpdated"
	RouteRemoved EventType = "RouteRemoved"
)

type Producer interface {
	Produce(ctx context.Context, event EventType) error
}
