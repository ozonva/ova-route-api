package kafka

import (
	"context"
	"ova-route-api/internal/broker"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
)

// the topic and broker address are initialized as constants
// const (
// 	topic         = "message-log"
// 	brokerAddress = "localhost:9092"
// )

type producer struct {
	loggger zerolog.Logger
	counter uint64
	writer  *kafka.Writer
}

func NewProducer(topic, brokerAddress string, logger zerolog.Logger) broker.Producer {
	p := producer{
		loggger: logger,
		counter: 0,
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:      []string{brokerAddress},
			Topic:        topic,
			BatchSize:    10,
			BatchTimeout: 2 * time.Second,
		}),
	}

	return p
}

func (p producer) Produce(ctx context.Context, event broker.EventType) error {
	err := p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(int(p.counter))),
		Value: []byte(event),
	})
	if err != nil {
		p.loggger.Error().Msgf("producer could not write message. err: %v ", err.Error())
		return err
	}

	return nil
}
