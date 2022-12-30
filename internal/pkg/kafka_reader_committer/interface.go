//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package kafka_reader_committer

import (
	"companies-api/internal/pkg/kafka_committer"
	"companies-api/internal/pkg/kafka_reader"
)

type IKafkaReaderCommitter interface {
	kafka_committer.IKafkaCommitter
	kafka_reader.IKafkaReader
}
