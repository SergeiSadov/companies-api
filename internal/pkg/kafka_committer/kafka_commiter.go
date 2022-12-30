//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package kafka_committer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type IKafkaCommitter interface {
	CommitMessages(ctx context.Context, msgs ...kafka.Message) error
}
