package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Web struct {
	APIHost         string        `mapstructure:"APIHOST"`
	DebugHost       string        `mapstructure:"DEBUGHOST"`
	ReadTimeout     time.Duration `mapstructure:"READTIMEOUT"`
	WriteTimeout    time.Duration `mapstructure:"WRITETIMEOUT"`
	IdleTimeout     time.Duration `mapstructure:"IDLETIMEOUT"`
	ShutdownTimeout time.Duration `mapstructure:"SHUTDOWNTIMEOUT"`
}

type Auth struct {
	KeysFolder string `mapstructure:"KEYSFOLDER"`
	ActiveKID  string `mapstructure:"ACTIVEKID"`
}

type DB struct {
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
	Host         string `mapstructure:"HOST"`
	Name         string `mapstructure:"NAME"`
	MaxIdleConns int    `mapstructure:"MAXIDLECONNS"`
	MaxOpenConns int    `mapstructure:"MAXOPENCONNS"`
	DisableTLS   bool   `mapstructure:"DISABLETLS"`
}

type Zipkin struct {
	ReporterURI string  `mapstructure:"REPORTERURI"`
	ServiceName string  `mapstructure:"SERVICENAME"`
	Probability float64 `mapstructure:"PROBABILITY"`
}

type Configuration struct {
	Build  string
	Web    `mapstructure:"WEB"`
	Auth   `mapstructure:"AUTH"`
	DB     `mapstructure:"DB"`
	Zipkin `mapstructure:"ZIPKIN"`
}

func LoadDefault(build string) (Configuration, error) {
	viper.BindEnv("WEB.APIHOST", "WEB_APIHOST")
	viper.BindEnv("WEB.DEBUGHOST", "WEB_DEBUGHOST")
	viper.BindEnv("WEB.READTIMEOUT", "WEB_READTIMEOUT")
	viper.BindEnv("WEB.WRITETIMEOUT", "WEB_WRITETIMEOUT")
	viper.BindEnv("WEB.IDLETIMEOUT", "WEB_IDLETIMEOUT")
	viper.BindEnv("WEB.SHUTDOWNTIMEOUT", "WEB_SHUTDOWNTIMEOUT")
	viper.BindEnv("AUTH.KEYSFOLDER", "AUTH_KEYSFOLDER")
	viper.BindEnv("AUTH.ACTIVEKID", "AUTH_ACTIVEKID")
	viper.BindEnv("DB.USER", "DB_USER")
	viper.BindEnv("DB.PASSWORD", "DB_PASSWORD")
	viper.BindEnv("DB.HOST", "DB_HOST")
	viper.BindEnv("DB.NAME", "DB_NAME")
	viper.BindEnv("DB.MAXIDLECONNS", "DB_MAXIDLECONNS")
	viper.BindEnv("DB.MAXOPENCONNS", "DB_MAXOPENCONNS")
	viper.BindEnv("DB.DISABLETLS", "DB_DISABLETLS")
	viper.BindEnv("ZIPKIN.REPORTERURI", "ZIPKIN_REPORTERURI")
	viper.BindEnv("ZIPKIN.SERVICENAME", "ZIPKIN_SERVICENAME")
	viper.BindEnv("ZIPKIN.PROBABILITY", "ZIPKIN_PROBABILITY")

	viper.AutomaticEnv()

	cfg := Configuration{Build: build}

	if err := viper.Unmarshal(&cfg); err != nil {
		return Configuration{}, fmt.Errorf("unmarshalling: %w", err)
	}

	return cfg, nil
}
