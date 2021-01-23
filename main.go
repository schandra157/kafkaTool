package main

import (
	"fmt"
	"kafkaTool/ktools"
)

func main() {
	// load config from yaml file
	config, err := ktools.LoadBrokerConfig("config/kafka.yaml")
	if err != nil {
		fmt.Printf("error while loading the configuration : %v", err)
		return
	}
	if con := ktools.CreateConnection(config.BootstrapServers); con != nil {
		ktools.CreateTopics(*con, *config)
		//close connection once done
		defer (*con).Close()
	}
}
