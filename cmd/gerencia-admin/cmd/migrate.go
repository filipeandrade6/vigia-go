package cmd

import (
	"fmt"
	"os"

	"github.com/filipeandrade6/vigia-go/internal/gerencia-admin/config"
	"github.com/filipeandrade6/vigia-go/internal/gerencia-admin/core"
	"github.com/spf13/cobra"

	"github.com/filipeandrade6/vigia-go/internal/sys/logger"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log, err := logger.New("GRAVACAO")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer log.Sync()

		if err := core.Run(log, config.Configuration{}); err != nil {
			log.Errorw("startup", "ERROR", err)
			os.Exit(1)
		}
		fmt.Println("migrate called")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
