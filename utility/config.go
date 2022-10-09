package utility

import "github.com/spf13/viper"

type Config struct {
	DATABASE_URL  string `mapstructure:"DATABASE_URL"`
	DATABASE_NAME string `mapstructure:"DATABASE_NAME"`
	PORT          string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// The "viper.AutomaticEnv()" line below is so that Viper can successfully override the values it reads from
	// the config file with environment variables. That's very convenient when we want to deploy the application to
	// different environments such as staging or production in the future.
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal((&config))
	return
}
