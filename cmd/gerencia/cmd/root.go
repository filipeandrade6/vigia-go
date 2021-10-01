package cmd

import (
	"fmt"
	"os"

	"github.com/filipeandrade6/vigia-go/internal/gerencia/config"
	"github.com/filipeandrade6/vigia-go/internal/gerencia/core"
	"github.com/filipeandrade6/vigia-go/internal/sys/logger"

	"github.com/spf13/cobra"
)

var cfg config.Configuration

var rootCmd = &cobra.Command{
	Use:   "gerencia",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		log, err := logger.New("GERENCIA")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer log.Sync()

		if err := core.Run(log, cfg); err != nil {
			log.Errorw("startup", "ERROR", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	var err error
	cfg, err = config.ParseConfig("cli")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.Flags().StringVar(&cfg.Auth.Directory, "auth-directory", cfg.Auth.Directory, "diretorio onde se encontra as credenciais")
	rootCmd.Flags().StringVar(&cfg.Auth.ActiveKID, "auth-activekid", cfg.Auth.ActiveKID, "nome da chave privada")
	rootCmd.Flags().StringVar(&cfg.Database.Host, "db-host", cfg.Database.Host, "host do banco de dados")
	rootCmd.Flags().StringVar(&cfg.Database.User, "db-user", cfg.Database.User, "usuario do banco de dados")
	rootCmd.Flags().StringVar(&cfg.Database.Password, "db-password", cfg.Database.Password, "senha do banco de dados")
	rootCmd.Flags().StringVar(&cfg.Database.Name, "db-name", cfg.Database.Name, "nome do banco de dados")
	rootCmd.Flags().IntVar(&cfg.Database.MaxIDLEConns, "db-maxidleconns", cfg.Database.MaxIDLEConns, "numero maximo de conexoes ociosas")
	rootCmd.Flags().IntVar(&cfg.Database.MaxOpenConns, "db-maxopenconns", cfg.Database.MaxOpenConns, "numero maximo de conexoes abertas")
	rootCmd.Flags().StringVar(&cfg.Database.SSLMode, "db-sslmode", cfg.Database.SSLMode, "modo SSL de conexao com o banco de dados")
	rootCmd.Flags().StringVar(&cfg.Service.Address, "service-endereco", cfg.Service.Address, "endereco do servico grpc")
	rootCmd.Flags().StringVar(&cfg.Service.Conn, "service-conn", cfg.Service.Conn, "tipo de conexao do servico grpc")
	rootCmd.Flags().IntVar(&cfg.Service.Port, "service-port", cfg.Service.Port, "porta do servidor grpc")
}
