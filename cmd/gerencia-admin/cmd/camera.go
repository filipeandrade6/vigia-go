package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cameraCmd = &cobra.Command{
	Use:   "camera",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("camera called")
	},
}

var adicionarCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Args:  cobra.ExactArgs(7),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adicionar camera called")
	},
}

var listarCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listar camera called")
	},
}

// TODO - criar comando interativo...
var atualizarCmd = &cobra.Command{
	Use:       "atualizar",
	Short:     "Atualiza a informções de uma câmera existente",
	Args:      cobra.ExactValidArgs(8),
	ValidArgs: []string{"camera_id", "descricao", "endereco_ip", "porta", "canal", "usuario", "senha", "geolocalizacao"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("atualizar camera called")
		fmt.Println(strings.Join(args, " "))
	},
}

var removerCmd = &cobra.Command{
	Use:   "remover [camera_id]",
	Short: "Remove a câmera do banco de dados e dos servidores de gravação",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remover camera called")
	},
}

func init() {
	rootCmd.AddCommand(cameraCmd)
	cameraCmd.AddCommand(adicionarCmd)
	cameraCmd.AddCommand(listarCmd)
	cameraCmd.AddCommand(atualizarCmd)
	cameraCmd.AddCommand(removerCmd)
}
