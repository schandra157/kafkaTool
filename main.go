package main

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if con, err1 := ktools.CreateConnection(config.BootstrapServers); err1 == nil {
		ktools.CreateTopics(ctx, con, *config)
		defer con.Close()
	} else {
		fmt.Println("err", err1)
	}

}
