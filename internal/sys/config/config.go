package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func PrettyPrintConfig() string {
	return fmt.Sprintf(
		"VIGIA_GER_HOST: %s, VIGIA_GER_SERVER_CONN: %s, VIGIA_GER_SERVER_PORT: %d, VIGIA_GRA_HOST: %s, VIGIA_GRA_CONN: %s, VIGIA_GRA_PORT: %d, VIGIA_DB_HOST: %s, VIGIA_DB_USER: %s, VIGIA_DB_PASSWORD: %s, VIGIA_DB_NAME: %s, VIGIA_DB_MAXIDLECONNS: %d, VIGIA_DB_MAXOPENCONNS: %d, VIGIA_DB_SSLMODE: %s",
		viper.GetString("VIGIA_GER_HOST"),
		viper.GetString("VIGIA_GER_SERVER_CONN"),
		viper.GetInt("VIGIA_GER_SERVER_PORT"),
		viper.GetString("VIGIA_GRA_HOST"),
		viper.GetString("VIGIA_GRA_CONN"),
		viper.GetInt("VIGIA_GRA_PORT"),
		viper.GetString("VIGIA_DB_HOST"),
		viper.GetString("VIGIA_DB_USER"),
		viper.GetString("VIGIA_DB_PASSWORD"),
		viper.GetString("VIGIA_DB_NAME"),
		viper.GetInt("VIGIA_DB_MAXIDLECONNS"),
		viper.GetInt("VIGIA_DB_MAXOPENCONNS"),
		viper.GetString("VIGIA_DB_SSLMODE"),
	)
}
