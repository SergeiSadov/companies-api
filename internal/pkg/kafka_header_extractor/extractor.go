package kafka_header_extractor

import "github.com/segmentio/kafka-go"

type IExtractor interface {
	Extract(msg kafka.Message) (value string, found bool)
}

type Extractor struct {
	header string
}

func NewExtractor(header string) *Extractor {
	return &Extractor{header: header}
}

func (e *Extractor) Extract(msg kafka.Message) (value string, found bool) {
	for _, h := range msg.Headers {
		if h.Key == e.header {
			return string(h.Value), true
		}
	}

	return
}
