package core

import (
	"fmt"

	"github.com/filipeandrade6/vigia-go/internal/gravacao/client"
)

func Run() {
	fmt.Println("chegou aqui antes de criar o client de genrecia")
	gerenciaClient := client.NovoClientGerencia()
	fmt.Println("chegou aqui")
	if err := gerenciaClient.Migrate(); err != nil {
		fmt.Printf("deu ruim, %s", err)
	}
}
