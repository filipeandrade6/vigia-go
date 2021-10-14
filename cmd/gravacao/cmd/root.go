package cmd

import (
	"fmt"
	"os"

	// "path/filepath"

	"github.com/filipeandrade6/vigia-go/internal/gravacao/config"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/core"
	"github.com/filipeandrade6/vigia-go/internal/sys/logger"

	"github.com/spf13/cobra"
)

var cfg config.Configuration

var rootCmd = &cobra.Command{
	Use:   "gravacao",
	Short: "Servico de gravacao",
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

	rootCmd.Flags().IntVar(&cfg.Gravacao.Port, "porta", cfg.Gravacao.Port, "porta para o servico gRPC")
}
