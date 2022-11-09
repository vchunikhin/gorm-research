package config

import "github.com/spf13/viper"

const (
	configType = "env"
	configName = "app"
)

type Config struct {
	DBHost         string `mapstructure:"MYSQL_HOST"`
	DBUserName     string `mapstructure:"MYSQL_USER"`
	DBUserPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBName         string `mapstructure:"MYSQL_DB"`
	DBPort         string `mapstructure:"MYSQL_PORT"`
	ServerPort     string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
