package configuration

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	// Constants for viper variable names. Will be used to set
	// default values as well as to get each value

	varBrokerURL = "broker.url"
	varPodName   = "pod.name"
	varSubjects  = "subjects"
)

// Config encapsulates the Viper configuration registry which stores the
// configuration data in-memory.
type Config struct {
	v *viper.Viper
}

// New creates a configuration reader object using a configurable configuration
// file path.
func New() Config {
	c := Config{
		v: viper.New(),
	}
	c.v.SetEnvPrefix("F8")
	c.v.AutomaticEnv()
	c.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	c.v.SetTypeByDefaultValue(true)
	c.setConfigDefaults()
	return c
}

func (c *Config) setConfigDefaults() {
	c.v.SetDefault(varPodName, "localhost")
	c.v.SetDefault(varSubjects, "test")
}

// GetBrokerURL returns URL of the broker to connect to, to publish and subscribe to messages
func (c *Config) GetBrokerURL() string {
	return c.v.GetString(varBrokerURL)
}

// GetPodName returns the name of the pod that runs the program
func (c *Config) GetPodName() string {
	return c.v.GetString(varPodName)
}

// GetSubjects returns the subject to publish/subscribe
func (c *Config) GetSubjects() []string {
	subjects := c.v.GetString(varSubjects)
	return strings.Split(subjects, ",")
}
