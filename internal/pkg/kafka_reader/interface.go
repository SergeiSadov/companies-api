//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package kafka_reader

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type IKafkaReader interface {
	FetchMessage(ctx context.Context) (kafka.Message, error)
}
