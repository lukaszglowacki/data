package cfg

import (
	"github.com/spf13/viper"
)

func Init(name string) (*viper.Viper, error) {
	c := viper.New()
	// always check for an environment variable on `viper.Get()` calls
	c.AutomaticEnv()

	// Look for the {name}.(yml|yaml) in the following paths, in order of precedence
	c.SetConfigName(name)

	c.AddConfigPath("/etc/config/")
	c.AddConfigPath("./config/")

	if err := c.ReadInConfig(); err != nil {
		// Not the best to encapsulate err, but adding context so the error is easier to debug if Config is malformed.
		return nil, err
	}
	return c, nil
}
