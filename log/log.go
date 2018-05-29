package log

import (
	"fmt"

	"github.com/fabric8-services/fabric8-nats/configuration"
	"github.com/sirupsen/logrus"
)

var config configuration.Config

func init() {
	config = configuration.New()

}

// Infof displays the msg with optional args at the `info` level,
// preceeded by the name of the pod in which the program is running.
func Infof(msg string, args ...interface{}) {
	logrus.Infof("[%s] %s", config.GetPodName(), fmt.Sprintf(msg, args...))
}

// Fatal displays the given err at the `fatal` level,
// preceeded by the name of the pod in which the program is running.
func Fatal(err error) {
	logrus.Fatalf("[%s] %v", config.GetPodName(), err)
}

// Warn displays the given err at the `warn` level,
// preceeded by the name of the pod in which the program is running.
func Warn(msg string) {
	logrus.Warn("[%s] %s", config.GetPodName(), msg)
}
