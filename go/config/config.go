package config

import "github.com/spf13/viper"

var (
	Config AppConfig
)

type AppConfig struct {
	Messaging      MessagingConfig      `yaml:"messaging"`
	DataIngestion  DataIngestionConfig  `yaml:"dataingestion"`
	Transformation TransformationConfig `yaml:"transformation"`
	Api            ApiConfig            `yaml:"api"`
	Tmdb           TMBDConfig           `yaml"tmdb"`
}

type MessagingConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DataIngestionConfig struct {
	QueueName string `yaml:"queuename"`
}

type TransformationConfig struct {
	QueueName string `yaml:"queuename"`
}

type ApiConfig struct {
	QueueName string `yaml:"queuename"`
}

type TMBDConfig struct {
	ApiKey string `yaml:"apikey"`
}

func LoadConfig() error {
	viper.SetConfigFile("./config/config.yml") // TODO: add more customization
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&Config)
}
