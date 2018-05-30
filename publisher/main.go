package main

import (
	"fmt"
	"time"

	"github.com/fabric8-services/fabric8-nats/configuration"
	"github.com/fabric8-services/fabric8-nats/log"
	"github.com/nats-io/go-nats"
)

func main() {

	// loads the configuration
	config := configuration.New()

	// open a connection to the NATS server
	log.Infof("opening connection to '%s'...", config.GetBrokerURL())
	nc, err := nats.Connect(config.GetBrokerURL())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		nc.Close()
	}()

	log.Infof("connection opened: '%t'...", nc.IsConnected())

	// infinite loop of message publishing...
	count := 1
	subjects := config.GetSubjects()
	for {
		// block for a few seconds...
		<-time.After(3 * time.Second)
		msg := fmt.Sprintf("message #%d", count)
		for _, sub := range subjects {
			nc.Publish(sub, []byte(msg))
			log.Infof("published on subject '%s': '%s'", sub, msg)
		}
		nc.Flush()
		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		}
		count++
	}
}
