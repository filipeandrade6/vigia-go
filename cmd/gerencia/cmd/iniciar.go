package cmd

import (
	"fmt"
	"os"

	"github.com/filipeandrade6/vigia-go/internal/gerencia/core"
	"github.com/filipeandrade6/vigia-go/internal/logger"
	"github.com/spf13/cobra"
)

// iniciarCmd represents the iniciar command
var iniciarCmd = &cobra.Command{
	Use:   "iniciar",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		log, err := logger.New("GERENCIA")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer log.Sync()

		if err := core.Run(log); err != nil {
			log.Errorw("startup", "ERROR", err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(iniciarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// iniciarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// iniciarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
