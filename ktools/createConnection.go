package ktools

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//CreateConnection creates kafka connection
func CreateConnection(brkrAddress string) (*kafka.AdminClient, error) {
	return kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": brkrAddress})
}
