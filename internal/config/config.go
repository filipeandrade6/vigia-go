package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Web struct {
	APIHost         string        `mapstructure:"apihost"`
	DebugHost       string        `mapstructure:"debughost"`
	ReadTimeout     time.Duration `mapstructure:"readtimeout"`
	WriteTimeout    time.Duration `mapstructure:"writetimeout"`
	IdleTimeout     time.Duration `mapstructure:"idletimeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdowntimeout"`
}

type Auth struct {
	KeysFolder string `mapstructure:"keysfolder"`
	ActiveKID  string `mapstructure:"activekid"`
}

type DB struct {
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Name         string `mapstructure:"name"`
	MaxIdleConns int    `mapstructure:"maxidleconns"`
	MaxOpenConns int    `mapstructure:"maxopenconns"`
	DisableTLS   bool   `mapstructure:"disabletls"`
}

type Zipkin struct {
	ReporterURI string  `mapstructure:"reporteruri"`
	ServiceName string  `mapstructure:"servicename"`
	Probability float64 `mapstructure:"probability"`
}

type Configuration struct {
	Build  string
	Web    `mapstructure:"web"`
	Auth   `mapstructure:"auth"`
	DB     `mapstructure:"db"`
	Zipkin `mapstructure:"zipkin"`
}

func Load(build string) (Configuration, error) {
	viper.BindEnv("web.apihost", "WEB_APIHOST")
	viper.BindEnv("web.debughost", "WEB_DEBUGHOST")
	viper.BindEnv("web.readtimeout", "WEB_READTIMEOUT")
	viper.BindEnv("web.writetimeout", "WEB_WRITETIMEOUT")
	viper.BindEnv("web.idletimeout", "WEB_IDLETIMEOUT")
	viper.BindEnv("web.shutdowntimeout", "WEB_SHUTDOWNTIMEOUT")
	viper.BindEnv("auth.keysfolder", "AUTH_KEYSFOLDER")
	viper.BindEnv("auth.activekid", "AUTH_ACTIVEKID")
	viper.BindEnv("db.user", "DB_USER")
	viper.BindEnv("db.password", "DB_PASSWORD")
	viper.BindEnv("db.host", "DB_HOST")
	viper.BindEnv("db.name", "DB_NAME")
	viper.BindEnv("db.maxidleconns", "DB_MAXIDLECONNS")
	viper.BindEnv("db.maxopenconns", "DB_MAXOPENCONNS")
	viper.BindEnv("db.disabletls", "DB_DISABLETLS")
	viper.BindEnv("zipkin.reporteruri", "ZIPKIN_REPORTERURI")
	viper.BindEnv("zipkin.servicename", "ZIPKIN_SERVICENAME")
	viper.BindEnv("zipkin.probability", "ZIPKIN_PROBABILITY")

	viper.AutomaticEnv()

	cfg := Configuration{Build: build}

	if err := viper.Unmarshal(&cfg); err != nil {
		return Configuration{}, fmt.Errorf("unmarshalling: %w", err)
	}

	return cfg, nil
}
