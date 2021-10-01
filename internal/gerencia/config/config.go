package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Auth struct {
	Directory string `mapstructure:"directory"`
	ActiveKID string `mapstructure:"activekid"`
}

type Database struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Name         string `mapstructure:"name"`
	MaxIDLEConns int    `mapstructure:"maxidleconns"`
	MaxOpenConns int    `mapstructure:"maxopenconns"`
	SSLMode      string `mapstructure:"sslmode"`
}

type Service struct {
	Address string `mapstructure:"address"`
	Conn    string `mapstructure:"conn"`
	Port    int    `mapstructure:"port"`
}

type Configuration struct {
	Build    string
	Auth     Auth     `mapstructure:"auth"`
	Database Database `mapstructure:"database"`
	Service  Service  `mapstructure:"service"`
}

func ParseConfig(build string) (Configuration, error) {
	viper.SetDefault("auth.directory", "deployments/keys")
	viper.SetDefault("auth.activekid", "bcc18baa-7830-4cfc-8f96-8a26ede5d81f")
	viper.SetDefault("database.host", "dev-postgres:5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "secret")
	viper.SetDefault("database.name", "vigia")
	viper.SetDefault("database.maxidleconns", "0")
	viper.SetDefault("database.maxopenconns", "0")
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("service.address", "gerencia")
	viper.SetDefault("service.conn", "tcp")
	viper.SetDefault("service.port", "12346")

	viper.BindEnv("auth.directory", "VIGIA_AUTH_DIR")
	viper.BindEnv("auth.activekid", "VIGIA_AUTH_ACTIVEKID")
	viper.BindEnv("database.host", "VIGIA_DB_HOST")
	viper.BindEnv("database.user", "VIGIA_DB_USER")
	viper.BindEnv("database.password", "VIGIA_DB_PASSWORD")
	viper.BindEnv("database.name", "VIGIA_DB_NAME")
	viper.BindEnv("database.maxidleconns", "VIGIA_DB_MAXIDLECONNS")
	viper.BindEnv("database.maxopenconns", "VIGIA_DB_MAXOPENCONNS")
	viper.BindEnv("database.sslmode", "VIGIA_DB_SSLMODE")
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
