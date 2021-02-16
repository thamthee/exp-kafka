package configs

import (
	"io"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		IsDebug bool `mapstructure:"is-debug"`
	} `mapstructure:"app"`

	Web struct {
		APIHost         string        `mapstructure:"api-host"`
		ReadTimeout     time.Duration `mapstructure:"read-timeout"`
		WriteTimeout    time.Duration `mapstructure:"write-timeout"`
		ShutdownTimeout time.Duration `mapstructure:"shutdown-timeout"`
	} `mapstructure:"web"`

	Kafka struct {
		Brokers  []string `mapstructure:"brokers"`
		ClientID string   `mapstructure:"client-id"`
	} `mapstructure:"kafka"`
}

func New(file io.Reader, fileType string) (Config, error) {
	viper.SetConfigType(fileType)
	if err := viper.ReadConfig(file); err != nil {
		return Config{}, errors.Wrap(err, "read config")
	}

	viper.AutomaticEnv()

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		return Config{}, errors.Wrap(err, "unable to decode config")
	}

	if url, ok := viper.Get("kafkaURL").(string); ok {
		conf.Kafka.Brokers = []string{url}
	}

	return conf, nil
}
