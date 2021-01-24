package ktools

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//TopicDesc holds topic Related Information
type TopicDesc struct {
	Name        string `yaml:"name"`
	RetentionMs string `yaml:"retentionMs"`
	Compression string `yaml:"compression"`
}

//KafkaBroker holds broker Info and topic details of kafka
type KafkaBroker struct {
	BootstrapServers  string `yaml:"bootstrapServers"`
	ReplicationFactor int    `yaml:"replication"`
	NumPartitions     int    `yaml:"partitions"`
	Topics            []TopicDesc
}

//LoadBrokerConfig loads and return broker config from filepath
func LoadBrokerConfig(filePath string) (*KafkaBroker, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	cfg := KafkaBroker{}
	if err = yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
