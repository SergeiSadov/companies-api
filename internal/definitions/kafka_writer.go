package definitions

import (
	"time"

	"companies-api/internal/configs"

	"github.com/sarulabs/di"
	"github.com/segmentio/kafka-go"
)

const (
	CreateCompanyKafkaWriterDef = "create_company_kafka_writer"
	UpdateCompanyKafkaWriterDef = "update_company_kafka_writer"
	DeleteCompanyKafkaWriterDef = "delete_company_kafka_writer"
)

func GetCreateCompanyKafkaWriter() di.Def {
	return di.Def{
		Name:  CreateCompanyKafkaWriterDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(CfgDef).(configs.Config)

			return &kafka.Writer{
				Addr:         kafka.TCP(cfg.Kafka.Brokers...),
				Topic:        cfg.Kafka.CreateCompanyTopic,
				Balancer:     &kafka.LeastBytes{},
				BatchTimeout: time.Millisecond * 5,
			}, nil
		},
	}
}

func GetUpdateCompanyKafkaWriterDef() di.Def {
	return di.Def{
		Name:  UpdateCompanyKafkaWriterDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(CfgDef).(configs.Config)

			return &kafka.Writer{
				Addr:         kafka.TCP(cfg.Kafka.Brokers...),
				Topic:        cfg.Kafka.UpdateCompanyTopic,
				Balancer:     &kafka.LeastBytes{},
				BatchTimeout: time.Millisecond * 5,
			}, nil
		},
	}
}

func GetDeleteCompanyKafkaWriterDef() di.Def {
	return di.Def{
		Name:  DeleteCompanyKafkaWriterDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(CfgDef).(configs.Config)

			return &kafka.Writer{
				Addr:         kafka.TCP(cfg.Kafka.Brokers...),
				Topic:        cfg.Kafka.DeleteCompanyTopic,
				Balancer:     &kafka.LeastBytes{},
				BatchTimeout: time.Millisecond * 5,
			}, nil
		},
	}
}
