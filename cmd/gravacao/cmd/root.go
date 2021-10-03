package cmd

import (
	"fmt"
	"os"

	"github.com/filipeandrade6/vigia-go/internal/gravacao/config"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/core"
	"github.com/filipeandrade6/vigia-go/internal/sys/logger"

	"github.com/spf13/cobra"
)

var cfg config.Configuration

// TODO adicionar help
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gravacao",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		log, err := logger.New("GRAVACAO")
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

	rootCmd.Flags().StringVar(&cfg.Gravacao.Conn, "conn", cfg.Gravacao.Conn, "tipo de conexao do servidor gRPC")
	rootCmd.Flags().IntVar(&cfg.Gravacao.Port, "port", cfg.Gravacao.Port, "porta do servidor grpc")
	rootCmd.Flags().StringVar(&cfg.Gravacao.Armazenamento, "armazenamento", cfg.Gravacao.Armazenamento, "local de armazenamento")
	rootCmd.Flags().StringVar(&cfg.Gravacao.Housekeeper, "housekeeper", cfg.Gravacao.Housekeeper, "tempo de armazenamento (cron)")
}