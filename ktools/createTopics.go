package ktools

import (
	"log"

	"github.com/Shopify/sarama"
)

//CreateTopics will create Topic if not already exist , If exist will modify the existing topic configuration
func CreateTopics(admin sarama.ClusterAdmin, details KafkaBroker) {
	for _, topic := range details.Topics {
		err := admin.CreateTopic(topic.Name, &sarama.TopicDetail{
			NumPartitions:     int32(details.NumPartitions),
			ReplicationFactor: int16(details.ReplicationFactor),
			ConfigEntries:     topic.Config,
		}, false)

		if err != nil {
			log.Printf("%v \n", err)
			err = admin.AlterConfig(sarama.TopicResource, topic.Name, topic.Config, false)
			if err != nil {
				log.Printf("Error while updating topic configuration: %v", err.Error())
			} else {
				log.Printf("Topic updated successfully:%v", topic.Name)
			}
		} else {
			log.Printf("Topic Created successfully:%v", topic.Name)
		}
	}
}
