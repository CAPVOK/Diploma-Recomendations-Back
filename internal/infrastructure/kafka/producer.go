package kafka

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"time"

	"github.com/segmentio/kafka-go"
)

type kafkaProducer struct {
	writer *kafka.Writer
	logger *zap.Logger
	topic  string
}

type IKafkaProducer interface {
	Send(ctx context.Context, key string, value interface{}) error
}

func NewKafkaProducer(broker, topic string, logger *zap.Logger) IKafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &kafkaProducer{
		writer: writer,
		topic:  topic,
		logger: logger.Named("kafka-producer"),
	}
}

func (p *kafkaProducer) Send(ctx context.Context, key string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		p.logger.Error("marshal value", zap.Error(err))
		return err
	}

	p.logger.Info("Sending kafka message", zap.String("key", key), zap.String("value", string(bytes)))

	msg := kafka.Message{
		Key:   []byte(key),
		Value: bytes,
		Time:  time.Now(),
	}

	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		p.logger.Error("write messages", zap.Error(err))
		return err
	}

	p.logger.Info("message sent", zap.String("key", key))

	return nil
}
