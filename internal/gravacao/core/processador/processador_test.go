package processador_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/processo"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/database/tests"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/core/processador"
)

func TestProcessador(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	processoCore := processo.NewCore(log, db)
	registroCore := registro.NewCore(log, db)
	cameraCore := camera.NewCore(log, db)

	ctx := context.Background()

	t.Log("\tTestando Processador...............")
	{
		SupErrChan := make(chan error)

		go func() {
			for {
				fmt.Printf("\n\n%s\n\n", <-SupErrChan)
			}
		}()

		nProcesso := processo.NewProcesso{
			ServidorGravacaoID: "d03307d4-2b28-4c23-a004-3da25e5b8bb1",
			CameraID:           "d03307d4-2b28-4c23-a004-3da25e5b8ce3",
			Processador:        1,
			Adaptador:          1,
			Execucao:           false,
		}

		np := processador.NewProcessador("d03307d4-2b28-4c23-a004-3da25e5b8bb1", "/home/filipe", processoCore, cameraCore, registroCore, SupErrChan)
		go np.Gerenciar()

		prc, err := processoCore.Create(ctx, nProcesso)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("\nprocessoID: %s\n", prc.ProcessoID)

		err = np.StartProcesso(ctx, prc.ProcessoID)
		if err != nil {
			t.Fatal(err)
		}

		time.Sleep(time.Duration(5 * time.Second))

		nProcesso2 := nProcesso
		nProcesso2.CameraID = "d03307d4-2b28-4c23-a004-3da25e5b8aa3"

		prc2, err := processoCore.Create(ctx, nProcesso)
		if err != nil {
			t.Fatal(err)
		}

		err = np.StartProcesso(ctx, prc2.ProcessoID)
		if err != nil {
			t.Fatal(err)
		}

		np.StopProcesso(prc.ProcessoID)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("pausando o segundo")

		err = np.StopProcesso(prc2.ProcessoID)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("iniciando o primeiro")

		err = np.StartProcesso(ctx, prc.ProcessoID)
		if err != nil {
			t.Fatal(err)
		}

		err = np.StartProcesso(ctx, prc.ProcessoID)
		if err != nil {
			t.Fatal(err)
		}

		placas, _ := registroCore.Query(context.Background(), "", 1, 100)
		fmt.Println(placas)
	}
}
