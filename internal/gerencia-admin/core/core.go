package core

import (
	"fmt"

	"github.com/filipeandrade6/vigia-go/internal/gravacao/client"
)

func Run() {
	gerenciaClient := client.NovoClientGerencia()
	if err := gerenciaClient.Migrate(); err != nil {
		fmt.Printf("deu ruim, %s", err)
	}
}
