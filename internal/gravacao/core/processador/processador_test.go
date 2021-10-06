package processador_test

import (
	"context"
	"testing"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/processo"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/core/veiculo"
	"github.com/filipeandrade6/vigia-go/internal/database/tests"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/core/processador"

	"github.com/google/go-cmp/cmp"
)

func TestProcessador(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	processoCore := processo.NewCore(log, db)
	registroCore := registro.NewCore(log, db)
	cameraCore := camera.NewCore(log, db)
	veiculoCore := veiculo.NewCore(log, db)

	ctx := context.Background()

	t.Log("\tTestando Processador...............")
	{
		errChan := make(chan error)
		matchChan := make(chan string)

		nProcesso := processo.NewProcesso{
			ServidorGravacaoID: "d03307d4-2b28-4c23-a004-3da25e5b8bb1",
			CameraID:           "d03307d4-2b28-4c23-a004-3da25e5b8ce3",
			Processador:        1,
			Adaptador:          1,
			Execucao:           false,
		}

		np := processador.NewProcessador("d03307d4-2b28-4c23-a004-3da25e5b8bb1", "/home/filipe", processoCore, cameraCore, registroCore, errChan, matchChan)
		go np.Gerenciar()

		prc, err := processoCore.Create(ctx, nProcesso)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to create processo: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to create processo.", tests.Success)

		saved, err := processoCore.QueryByID(ctx, prc.ProcessoID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve processo by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve processo by ID.", tests.Success)

		if diff := cmp.Diff(prc, saved); diff != "" {
			t.Fatalf("\t%s\tShould get back the same processo. Diff:\n%s", tests.Failed, diff)
		}
		t.Logf("\t%s\tShould get back the same processo.", tests.Success)

		err = np.StartProcesso(ctx, prc.ProcessoID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to start processo: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to start processo.", tests.Success)

		registroMatch := <-matchChan

		regSaved, err := registroCore.Query(ctx, prc.ProcessoID, 1, 1)
		if len(regSaved) <= 0 {
			t.Fatalf("\t%s\tShould have received registro by processoID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have received registro by processoID", tests.Success)

		err = np.StartProcesso(ctx, prc.ProcessoID) // TODO criar erros e colocar errors.Is(ssxxx,xxxx)
		if err != nil {
			t.Fatalf("\t%s\tShould get already executing error: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould get already executing error.", tests.Success)

		err = np.StopProcesso(ctx, prc.ProcessoID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to stop processo: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be albe to stop processo.", tests.Success)

		matched, err := registroCore.QueryByID(ctx, registroMatch)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve registro by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve registro by ID.", tests.Success)

		_, err = veiculoCore.QueryByPlaca(ctx, matched.Placa)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve placa by registro: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve registro by registro.", tests.Success)

		err = np.StopProcesso(ctx, prc.ProcessoID) // TODO criar erros e colocar errors.Is(ssxxx,xxxx)
		if err != nil {
			t.Fatalf("\t%s\tShould get already stopped error: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould get already stopped error.", tests.Success)

		select {
		case <-errChan:
			t.Fatalf("\t%s\tShould NOT get error: %s.", tests.Failed, err)
		}
	}
}
