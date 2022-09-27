//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package kafka_writer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type IKafkaWriter interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) (err error)
}
