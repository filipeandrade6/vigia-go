package cmd

import (
	"fmt"
	"os"

	"github.com/filipeandrade6/vigia-go/internal/gerencia-admin/config"
	"github.com/filipeandrade6/vigia-go/internal/gerencia-admin/core"
	"github.com/filipeandrade6/vigia-go/internal/sys/logger"

	"github.com/spf13/cobra"
)

var cfg config.Configuration

// TODO adicionar help
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gerencia-client",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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

	rootCmd.Flags().StringVar(&cfg.Service.Address, "service-endereco", cfg.Service.Address, "endereco do servico grpc")
	rootCmd.Flags().StringVar(&cfg.Service.Conn, "service-conn", cfg.Service.Conn, "tipo de conexao do servico grpc")
	rootCmd.Flags().IntVar(&cfg.Service.Port, "service-port", cfg.Service.Port, "porta do servidor grpc")
}
