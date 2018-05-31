package main

import (
	"fmt"
	"time"

	"github.com/fabric8-services/fabric8-nats/configuration"
	"github.com/fabric8-services/fabric8-nats/log"
	"github.com/nats-io/go-nats-streaming"
)

func main() {

	// loads the configuration
	config := configuration.New()

	// open a connection to the NATS server
	log.Infof("opening connection to cluster '%s' on %s'...", config.GetClusterID(), config.GetBrokerURL())
	sc, err := stan.Connect(config.GetClusterID(), config.GetClientID(), stan.NatsURL(config.GetBrokerURL()))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at '%s'", err, config.GetBrokerURL())
	}
	defer sc.Close()

	log.Infof("connection established with server at '%s'", config.GetBrokerURL())

	// infinite loop of message publishing...
	count := 1
	subjects := config.GetSubjects()
	for {
		// block for a few seconds...
		<-time.After(3 * time.Second)
		msg := fmt.Sprintf("message #%d", count)
		for _, sub := range subjects {
			sc.Publish(sub, []byte(msg))
			log.Infof("published on subject '%s': '%s'", sub, msg)
		}
		count++
	}
}
