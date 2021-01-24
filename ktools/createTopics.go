package ktools

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func buildTopicConfig(details KafkaBroker) []kafka.TopicSpecification {
	var topicConfigs []kafka.TopicSpecification
	for _, topic := range details.Topics {
		topicInternal := make(map[string]string)
		topicInternal["retention.ms"] = topic.RetentionMs
		topicInternal["compression.type"] = topic.Compression

		topicCfg := kafka.TopicSpecification{
			Topic:             topic.Name,
			NumPartitions:     details.NumPartitions,
			ReplicationFactor: details.ReplicationFactor,
			Config:            topicInternal,
		}
		topicConfigs = append(topicConfigs, topicCfg)
	}
	return topicConfigs
}

func alterConfig(topicConfigs []kafka.TopicSpecification) []kafka.ConfigResource {
	alterTopicsCgf := make([]kafka.ConfigResource, 0, 0)
	for _, topic := range topicConfigs {
		var cfg []kafka.ConfigEntry
		for k, v := range topic.Config {
			cfg = append(cfg, kafka.ConfigEntry{Name: k, Value: v})
		}
		a := kafka.ConfigResource{
			Type:   kafka.ResourceTopic,
			Name:   topic.Topic,
			Config: cfg,
		}
		alterTopicsCgf = append(alterTopicsCgf, a)
	}
	return alterTopicsCgf
}

//CreateTopics will create Topic if not already exist , If exist will modify the existing topic configuration
func CreateTopics(ctx context.Context, admin *kafka.AdminClient, details KafkaBroker) {

	// create topic configs
	topicConfigs := buildTopicConfig(details)
	results, _ := admin.CreateTopics(ctx, topicConfigs)

	fmt.Println("result topic creation", results)

	failedTopics := make(map[string]bool)
	for _, res := range results {
		if res.Error.Code() == kafka.ErrTopicAlreadyExists {
			failedTopics[res.Topic] = true
		}
	}

	if len(failedTopics) > 0 {
		alterTopicConfigs := make([]kafka.TopicSpecification, 0, 0)
		for _, topic := range topicConfigs {
			if _, ok := failedTopics[topic.Topic]; ok {
				alterTopicConfigs = append(alterTopicConfigs, topic)
			}
		}

		alterTopicsCgf := alterConfig(alterTopicConfigs)
		res, err := admin.AlterConfigs(ctx, alterTopicsCgf)
		fmt.Println("result topic alteration", res, err)
	}

}
