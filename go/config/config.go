package config

import "github.com/spf13/viper"

var (
	Config AppConfig
)

type AppConfig struct {
	Messaging     MessagingConfig     `yaml:"messaging"`
	DataIngestion DataIngestionConfig `yaml:"data ingestion"`
	Api           ApiConfig           `yaml:"api"`
}

type MessagingConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DataIngestionConfig struct {
	QueueName string `yaml:"queue name"`
}

type ApiConfig struct {
	QueueName string `yaml:"queue name"`
}

func LoadConfig() error {
	viper.SetConfigFile("./config/config.yml") // TODO: add more customization
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&Config)
}
