package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Service struct {
	Address string `mapstructure:"address"`
	Conn    string `mapstructure:"conn"`
	Port    int    `mapstructure:"port"`
}

type Configuration struct {
	Build   string
	Service Service `mapstructure:"service"`
}

func ParseConfig(build string) (Configuration, error) {
	viper.SetDefault("service.address", "gerencia")
	viper.SetDefault("service.conn", "tcp")
	viper.SetDefault("service.port", "12346")

	viper.BindEnv("service.host", "VIGIA_GER_ADDRESS")
	viper.BindEnv("service.conn", "VIGIA_GER_SERVER_CONN")
	viper.BindEnv("service.port", "VIGIA_GER_SERVER_PORT")

	viper.AutomaticEnv()

	cfg := Configuration{Build: build}

	if err := viper.Unmarshal(&cfg); err != nil {
		return Configuration{}, fmt.Errorf("unmsarshalling config: %w", err)
	}

	return cfg, nil
}
