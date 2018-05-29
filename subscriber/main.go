package main

import (
	"runtime"

	"github.com/fabric8-services/fabric8-nats/configuration"
	"github.com/fabric8-services/fabric8-nats/log"
	"github.com/nats-io/go-nats"
)

func main() {

	// loads the configuration
	config := configuration.New()

	go func() {
		// open a connection to the NATS server
		log.Infof("opening connection to '%s'...", config.GetBrokerURL())
		nc, err := nats.Connect(config.GetBrokerURL())
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			log.Warn("closing the connection...")
			nc.Close()
		}()
		log.Infof("connection opened: '%t'...", nc.IsConnected())

		_, err = nc.Subscribe(config.GetSubject(), func(msg *nats.Msg) {
			// Handle the message
			log.Infof("received message on '%s': '%s'", msg.Subject, string(msg.Data))
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("listening on '%s'...", config.GetSubject())
		done := make(chan bool)
		// block to keep the connection and the subscription alive
		<-done
	}()
	// exit main rountine, but keep other routines alive
	runtime.Goexit()
}
