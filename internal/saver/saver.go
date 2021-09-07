package saver

import (
	"context"
	"ova-route-api/internal/flusher"
	"ova-route-api/internal/models"
	"sync"
	"time"
)

type Saver interface {
	Save(ctx context.Context, route models.Route) // заменить на свою сущность
	// Init()
	Close(ctx context.Context)
	BuffSize() uint
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(capacity uint, flusher flusher.Flusher) Saver {
	s := &saver{
		Mutex:   &sync.Mutex{},
		flusher: flusher,
		cap:     capacity,
		buff:    make([]models.Route, 0),
	}

	return s
}

type saver struct {
	*sync.Mutex
	once    sync.Once
	flusher flusher.Flusher
	cap     uint
	buff    []models.Route
}

func (s *saver) BuffSize() uint {
	s.Lock()
	defer s.Unlock()

	return uint(len(s.buff))
}

func (s *saver) Save(ctx context.Context, route models.Route) {
	// Стартанули автоматический сброс по таймеру
	s.once.Do(func() {
		go func(s *saver) {
			ticker := time.NewTicker(time.Second * 2)
			defer ticker.Stop()

			for {
				<-ticker.C
				s.flush(ctx)
			}
		}(s)
	})

	// Если достигли максимальной емкости
	if len(s.buff) == int(s.cap) {
		s.flush(ctx)
	}

	s.Lock()
	defer s.Unlock()
	s.buff = append(s.buff, route)
}

func (s *saver) Close(ctx context.Context) {
	s.flush(ctx)
}

func (s *saver) flush(ctx context.Context) {
	s.Lock()
	defer s.Unlock()

	s.buff = s.flusher.Flush(ctx, s.buff)
}
