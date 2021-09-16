package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("database called")
		// TODO colocar para mostrar as config atuais no gerencia
	},
}

var conectarCmd = &cobra.Command{
	Use:   "conectar",
	Short: "Conecta no banco de dados",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("database called")
		// TODO implementar
	},
}

var configurarCmd = &cobra.Command{
	Use:   "configurar",
	Short: "Configurar o banco de dados",
	Long: `Configurar o banco de dados em formato de URL. Exemplo:

postgres://user:password@localhost:5432/vigia?sslmode=disable .`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configurar called")
		// TODO implementar - vai receber em formato de URL
	},
}

func init() {
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.AddCommand(conectarCmd)
	databaseCmd.AddCommand(configurarCmd)
}
