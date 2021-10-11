package processador_test

import (
	"context"
	"testing"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/core/veiculo"
	"github.com/filipeandrade6/vigia-go/internal/database/tests"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/service/processador" // TODO trocar aqui
)

// TODO teste de adicionar placa
// TODO teste de atualizar matchlist removendo e ver se aparece

func TestProcessador(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	// cameraCore := camera.NewCore(log, db)
	// processoCore := processo.NewCore(log, db)
	registroCore := registro.NewCore(log, db)
	veiculoCore := veiculo.NewCore(log, db)

	ctx := context.Background()

	t.Log("\tTestando Processador...............")
	{
		errChan := make(chan error)
		matchChan := make(chan string)
		stopChan := make(chan struct{})
		stoppedChan := make(chan struct{})

		ticker := time.NewTicker(3 * time.Second)

		np := processador.New(
			registroCore,
			"d03307d4-2b28-4c23-a004-3da25e5b8bb1",
			"/home/filipe",
			1,
			errChan,
			matchChan,
			stopChan,
			stoppedChan,
		)

		veiculos, err := veiculoCore.QueryAll(ctx)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to query veiculos: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to query veiculos.", tests.Success)

		var veiculosList []string

		for _, v := range veiculos {
			veiculosList = append(veiculosList, v.Placa)
		}

		err = np.AtualizarMatchList(veiculosList)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to update matchlist: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to update matchlist.", tests.Success)

		go np.Start()

		nProcesso := []processador.Processo{
			{
				ProcessoID:  "d03307d4-2b28-4c23-a004-3da32e5b8bb1",
				EnderecoIP:  "10.20.30.40",
				Porta:       1,
				Canal:       1,
				Usuario:     "admin",
				Senha:       "admin",
				Processador: 0,
			},
		}

		np.StartProcessos(nProcesso)

		prcs := np.ListProcessos()
		if len(prcs) != 1 || prcs[0] != nProcesso[0].ProcessoID {
			t.Fatalf("\t%s\tShould be able to retrieve only the started processo.", tests.Failed)
		}
		t.Logf("\t%s\tShould be able to retrieve only the started processo.", tests.Success)

		nProcesso = append(nProcesso, processador.Processo{
			ProcessoID:  "d03307d4-2b28-4c23-a004-3da32e5b8a61",
			EnderecoIP:  "11.21.31.41",
			Porta:       1,
			Canal:       1,
			Usuario:     "admin",
			Senha:       "admin",
			Processador: 0,
		})

		np.StartProcessos(nProcesso)

		prcs = np.ListProcessos()
		if len(prcs) != 2 {
			t.Fatalf("\t%s\tShould be able to retrieve only the started processos.", tests.Failed)
		}
		t.Logf("\t%s\tShould be able to retrieve only the started processos.", tests.Success)

		var registroMatch string
		select {
		case r := <-matchChan:
			registroMatch = r
		case <-ticker.C:
			t.Fatalf("\t%s\tShould NOT wait more than 5 seconds for match.", tests.Failed)
		}

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

		path, hours := np.GetServidorInfo()
		if path != "/home/filipe" || hours != 2 {
			t.Fatalf("\t%s\tShould get updated processador info.", tests.Failed)
		}
		t.Logf("\t%s\tShould be get updated processador info.", tests.Success)

		// np.StopGerencia()

		select {
		case err := <-errChan:
			t.Fatalf("\t%s\tShould NOT get any error from channel: %s.", tests.Failed, err)
		case <-ticker.C:
			t.Logf("\t%s\tShould NOT get any error from channel.", tests.Success)
		}

		// TODO vai receber erro de already executing
	}
}
