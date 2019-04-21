package config

import (
	"sync"

	"github.com/spf13/viper"
)

type (
	// ImmutableConfig return implementation of methods in config
	ImmutableConfig interface {
		GetPort() int
		GetNATSHost() string
		GetAWSRegion() string
		GetAWSAccessKey() string
		GetAWSSecretAccessKey() string
		GetAWSSessionToken() string
		GetAWSSESMailFrom() string
		GetMywishesSecretKey() string
	}

	im struct {
		Port               int    `mapstructure:"PORT"`
		NATSHost           string `mapstructure:"NATS_HOST"`
		AWSRegion          string `mapstructure:"AWS_REGION"`
		AWSAccessKey       string `mapstructure:"AWS_ACCESS_KEY"`
		AWSSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
		AWSSessionToken    string `mapstructure:"AWS_SESSION_TOKEN"`
		AWSSESMailFrom     string `mapstructure:"AWS_SES_MAIL_FROM"`
		MywishesSecretKey  string `mapstructure:"MYWISHES_SECRET_KEY"`
	}
)

func (i *im) GetPort() int {
	return i.Port
}

func (i *im) GetNATSHost() string {
	return i.NATSHost
}

func (i *im) GetAWSRegion() string {
	return i.AWSRegion
}

func (i *im) GetAWSAccessKey() string {
	return i.AWSAccessKey
}

func (i *im) GetAWSSecretAccessKey() string {
	return i.AWSSecretAccessKey
}

func (i *im) GetAWSSessionToken() string {
	return i.AWSSessionToken
}

func (i *im) GetAWSSESMailFrom() string {
	return i.AWSSESMailFrom
}

func (i *im) GetMywishesSecretKey() string {
	return i.MywishesSecretKey
}

var (
	imOnce    sync.Once
	myEnv     map[string]string
	immutable im
)

// NewImmutableConfig is a factory that return implementation of ImmutableConfig
func NewImmutableConfig() ImmutableConfig {
	imOnce.Do(func() {
		v := viper.New()
		if myEnv["APP_ENV"] == "staging" {
			v.SetConfigName("app.config.staging")
		} else if myEnv["APP_ENV"] == "production" {
			v.SetConfigName("app.config.prod")
		} else {
			v.SetConfigName("app.config.dev")
		}

		v.AddConfigPath(".")
		v.SetEnvPrefix("SES")
		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			panic("[SES-CONF] failed reading env configuration")
		}

		v.Unmarshal(&immutable)
	})
	return &immutable
}
