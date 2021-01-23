package ktools

import (
	"log"

	"github.com/Shopify/sarama"
)

//CreateConnection creates kafka connection
func CreateConnection(brkrAddress string) *sarama.ClusterAdmin {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	admin, err := sarama.NewClusterAdmin([]string{brkrAddress}, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
		return nil
	}
	return &admin
}
