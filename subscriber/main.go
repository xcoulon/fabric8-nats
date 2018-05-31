package main

import (
	"runtime"

	"github.com/fabric8-services/fabric8-nats/configuration"
	"github.com/fabric8-services/fabric8-nats/log"
	"github.com/nats-io/go-nats-streaming"
)

func main() {

	// loads the configuration
	config := configuration.New()

	go func() {
		// open a connection to the NATS server
		log.Infof("opening connection to cluster '%s' on %s'...", config.GetClusterID(), config.GetBrokerURL())
		sc, err := stan.Connect(config.GetClusterID(), config.GetClientID(), stan.NatsURL(config.GetBrokerURL()))
		if err != nil {
			log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at '%s'", err, config.GetBrokerURL())
		}
		defer sc.Close()

		log.Infof("connection established with server at '%s'", config.GetBrokerURL())

		subjects := config.GetSubjects()

		for _, sub := range subjects {
			_, err = sc.QueueSubscribe(sub, config.GetQueueGroup(), func(msg *stan.Msg) {
				// Handle the message
				log.Infof("received message with subject '%s': '%s'", msg.Subject, string(msg.Data))
			}, stan.DurableName(config.GetDurableName()), stan.DeliverAllAvailable())
			if err != nil {
				log.Fatal(err)
			}
			log.Infof("listening on '%s'...", sub)
		}
		done := make(chan bool)
		// block to keep the connection and the subscription alive
		<-done
	}()
	// exit main rountine, but keep other routines alive
	runtime.Goexit()
}
