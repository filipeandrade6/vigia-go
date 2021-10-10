package processador_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/processo"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/core/veiculo"
	"github.com/filipeandrade6/vigia-go/internal/database/tests"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/core/processador"

	"github.com/google/go-cmp/cmp"
)

// TODO teste de adicionar placa
// TODO teste de atualizar matchlist removendo e ver se aparece

func TestProcessador(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	cameraCore := camera.NewCore(log, db)
	processoCore := processo.NewCore(log, db)
	registroCore := registro.NewCore(log, db)
	veiculoCore := veiculo.NewCore(log, db)

	ctx := context.Background()

	t.Log("\tTestando Processador...............")
	{
		errChan := make(chan error)
		matchChan := make(chan string)
		stopChan := make(chan struct{})
		stoppedChan := make(chan struct{})

		ticker := time.NewTicker(4 * time.Second)

		np := processador.NewProcessador(
			"d03307d4-2b28-4c23-a004-3da25e5b8bb1",
			"/home/filipe",
			1,
			processoCore,
			cameraCore,
			registroCore,
			veiculoCore,
			errChan,
			matchChan,
			stopChan,
			stoppedChan,
		)

		err := np.AtualizarMatchList(ctx)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to update matchlist: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to update matchlist.", tests.Success)

		go np.Gerenciar()

		nProcesso := processo.NewProcesso{
			ServidorGravacaoID: "d03307d4-2b28-4c23-a004-3da25e5b8bb1",
			CameraID:           "d03307d4-2b28-4c23-a004-3da25e5b8ce3",
			Processador:        1,
			Adaptador:          1,
		}

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

		var registroMatch string
		select {
		case r := <-matchChan:
			registroMatch = r
		case <-ticker.C:
			t.Fatalf("\t%s\tShould NOT wait more than 5 seconds for match.", tests.Failed)
		}

		// time.Sleep(50 * time.Millisecond) // Wait for write in database

		matched, err := registroCore.QueryByID(ctx, registroMatch)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve registro by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve registro by ID.", tests.Success)

		_, err = veiculoCore.QueryByPlaca(ctx, matched.Placa)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve placa by registro: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve placa by registro.", tests.Success)

		if err = np.StartProcesso(ctx, prc.ProcessoID); !errors.As(err, &processador.ErrAlreadyStarted) {
			t.Fatalf("\t%s\tShould get already executing error: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould get already executing error.", tests.Success)

		nProcesso2 := processo.NewProcesso{
			ServidorGravacaoID: "d03307d4-2b28-4c23-a004-3da25e5b8bb1",
			CameraID:           "d03307d4-2b28-4c23-a004-3da25e5b8aa3",
			Processador:        1,
			Adaptador:          1,
		}

		prc2, err := processoCore.Create(ctx, nProcesso2)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to create another processo: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to create another processo.", tests.Success)

		err = np.StartProcesso(ctx, prc2.ProcessoID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to start another processo: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to start another processo.", tests.Success)

		prcList, err := np.ShowAllProcessos()
		if err != nil {
			t.Fatalf("\t%s\tShould be able to return the list of processos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to return the list of processos.", tests.Success)

		for i, p := range prcList {
			if p == false {
				t.Fatalf("\t%s\tProcesso[%q] should be running.", tests.Failed, i)
			}
		}
		t.Logf("\t%s\tShould all processos be running.", tests.Success)

		if len(prcList) != 2 {
			t.Fatalf("\t%s\tShould be two processos running.", tests.Failed)
		}
		t.Logf("\t%s\tShould be two processos running.", tests.Success)

		err = np.StopProcesso(prc.ProcessoID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to stop first processo: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be albe to stop first processo.", tests.Success)

		prcList, err = np.ShowAllProcessos()
		if err != nil {
			t.Fatalf("\t%s\tShould be able to return the list of processos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to return the list of processos.", tests.Success)

		prcPause := prcList[prc.ProcessoID]
		if prcPause {
			t.Fatalf("\t%s\tShould be paused the processo", tests.Failed)
		}
		t.Logf("\t%s\tShould be paused the processo.", tests.Success)

		if err = np.StopProcesso(prc.ProcessoID); !errors.As(err, &processador.ErrAlreadyStopped) {
			t.Fatalf("\t%s\tShould get already stopped error: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould get already stopped error.", tests.Success)

		if err = np.StopAllProcessos(); err != nil {
			t.Fatalf("\t%s\tShould be able to pause all processos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to pause all processos.", tests.Success)

		if err = np.StopProcesso("ffff07d4-2b28-4c23-a004-3da25e5b8bb1"); !errors.As(err, &processador.ErrProcessoNotFound) {
			t.Fatalf("\t%s\tShould get processo not found error: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould get processo not found error.", tests.Success)

		prcList, err = np.ShowAllProcessos()
		if err != nil {
			t.Fatalf("\t%s\tShould be able to return the list of processos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to return the list of processos.", tests.Success)

		for i, p := range prcList {
			if p == true {
				t.Fatalf("\t%s\tProcesso[%q] should be paused.", tests.Failed, i)
			}
		}
		t.Logf("\t%s\tShould all processos be paused.", tests.Success)

		if err = np.StartAllProcessos(ctx); err != nil {
			t.Fatalf("\t%s\tShould be able to start all processos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to start all processos.", tests.Success)

		prcList, err = np.ShowAllProcessos()
		if err != nil {
			t.Fatalf("\t%s\tShould be able to return the list of processos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to return the list of processos.", tests.Success)

		for i, p := range prcList {
			if p == false {
				t.Fatalf("\t%s\tProcesso[%q] should be running.", tests.Failed, i)
			}
		}
		t.Logf("\t%s\tShould all processos be running.", tests.Success)

		old := len(prcList)

		err = np.RemoveProcesso(prc.ProcessoID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to remove processo: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to remove processo.", tests.Success)

		prcList, err = np.ShowAllProcessos()
		if err != nil {
			t.Fatalf("\t%s\tShould be able to return the list of processos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to return the list of processos.", tests.Success)

		if (len(prcList) + 1) != old {
			t.Fatalf("\t%s\tShould be return the list of processos minus the removed processo.", tests.Failed)
		}
		t.Logf("\t%s\tShould be return the list of processos minus the removed processo:.", tests.Success)

		if _, ok := prcList[prc.ProcessoID]; ok {
			t.Fatalf("\t%s\tShould NOT be able to return the removed processo processoID[%q] from list of processos.", tests.Failed, prc.ProcessoID)
		}
		t.Logf("\t%s\tShould NOT be able to return the removed processo from list of processos.", tests.Success)

		err = np.RemoveAllProcessos()
		if err != nil {
			t.Fatalf("\t%s\tShould be able to remove all processos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to remove all processos.", tests.Success)

		prcList, err = np.ShowAllProcessos()
		if err != nil {
			t.Fatalf("\t%s\tShould be able to return the list of processos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to return the list of processos.", tests.Success)

		if len(prcList) != 0 {
			t.Fatalf("\t%s\tShould NOT be able to return any processos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to return any processos.", tests.Success)

		status := np.StatusHousekeeper()
		if !status {
			t.Fatalf("\t%s\tShould be running the housekeeper.", tests.Failed)
		}
		t.Logf("\t%s\tShould be running the housekeeper.", tests.Success)

		np.StopHousekeeper()

		status = np.StatusHousekeeper()
		if status {
			t.Fatalf("\t%s\tShould be stopped the housekeeper.", tests.Failed)
		}
		t.Logf("\t%s\tShould be stopped the housekeeper.", tests.Success)

		np.AtualizarHousekeeper(2)

		path, hours := np.GetArmazenamentoInfo()
		if path != "/home/filipe" || hours != 2 {
			t.Fatalf("\t%s\tShould get updated processador info.", tests.Failed)
		}
		t.Logf("\t%s\tShould be get updated processador info.", tests.Success)

		np.StopGerencia()

		select {
		case err := <-errChan:
			t.Fatalf("\t%s\tShould NOT get any error from channel: %s.", tests.Failed, err)
		case <-ticker.C:
			t.Logf("\t%s\tShould NOT get any error from channel.", tests.Success)
		}
	}
}
