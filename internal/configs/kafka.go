package configs

type Kafka struct {
	Brokers            []string `envconfig:"KAFKA_BROKERS" required:"true" default:"localhost:9092"`
	CreateCompanyTopic string   `envconfig:"KAFKA_CREATE_COMPANY_TOPIC" required:"true" default:"company-created"`
	UpdateCompanyTopic string   `envconfig:"KAFKA_UPDATE_COMPANY_TOPIC" required:"true" default:"company-updated"`
	DeleteCompanyTopic string   `envconfig:"KAFKA_DELETE_COMPANY_TOPIC" required:"true" default:"company-deleted"`
}
