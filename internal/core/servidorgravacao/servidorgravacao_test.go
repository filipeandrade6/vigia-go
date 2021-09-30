package servidorgravacao_test

import (
	"context"
	"errors"
	"testing"

	"github.com/filipeandrade6/vigia-go/internal/core/servidorgravacao"
	"github.com/filipeandrade6/vigia-go/internal/data/store/tests"
	"github.com/google/go-cmp/cmp"
)

func TestServidorGravacao(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	core := servidorgravacao.NewCore(log, db)

	ctx := context.Background()

	ns := servidorgravacao.NewServidorGravacao{
		EnderecoIP: "10.20.30.40",
		Porta:      5001,
	}

	t.Log("\tGiven the need to work with Servidores de Gravacao records.")
	{
		sv, err := core.Create(ctx, ns)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to create servidor de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to create servidor de gravacao.", tests.Success)

		saved, err := core.QueryByID(ctx, sv.ServidorGravacaoID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve servidor de gravacao by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve servidor de gravacao by ID.", tests.Success)

		if diff := cmp.Diff(sv, saved); diff != "" {
			t.Fatalf("\t%s\tShould get back the same servidor de gravacao. Diff:\n%s", tests.Failed, diff)
		}
		t.Logf("\t%s\tShould get back the same servidor de gravacao.", tests.Success)

		upd := servidorgravacao.UpdateServidorGravacao{
			EnderecoIP: tests.StringPointer("60.50.30.20"),
			Porta:      tests.IntPointer(2727),
		}

		if err = core.Update(ctx, sv.ServidorGravacaoID, upd); err != nil {
			t.Fatalf("\t%s\tShould be able to update servidor de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to update servidor de gravacao.", tests.Success)

		svs, err := core.Query(ctx, "", 1, 3)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve updated servidor de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve updated servidor de gravacao.", tests.Success)

		want := sv
		want.EnderecoIP = *upd.EnderecoIP
		want.Porta = *upd.Porta

		var idx int
		for i, s := range svs {
			if s.EnderecoIP == want.EnderecoIP {
				idx = i
			}
		}
		if diff := cmp.Diff(want, svs[idx]); diff != "" {
			t.Fatalf("\t%s\tShould get back the same servidor de gravacao. Diff:\n%s", tests.Failed, diff)
		}
		t.Logf("\t%s\tShould get back the same servidor de gravacao.", tests.Success)

		upd = servidorgravacao.UpdateServidorGravacao{
			Porta: tests.IntPointer(4343),
		}

		if err = core.Update(ctx, sv.ServidorGravacaoID, upd); err != nil {
			t.Fatalf("\t%s\tShould be able to update just some fields of servidor de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to update just some fields of servidor de gravacao.", tests.Success)

		saved, err = core.QueryByID(ctx, sv.ServidorGravacaoID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve servidor de gravacao by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve servidor de gravacao by ID.", tests.Success)

		if saved.Porta != *upd.Porta {
			t.Fatalf("\t%s\tShould be able to see updated Porta field: got %q want %q.", tests.Failed, saved.Porta, *upd.Porta)
		}
		t.Logf("\t%s\tShould be able to see updated Porta field.", tests.Success)

		if err = core.Delete(ctx, sv.ServidorGravacaoID); err != nil {
			t.Fatalf("\t%s\tShould be able to delete servidor de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to delete servidor de gravacao.", tests.Success)

		_, err = core.QueryByID(ctx, sv.ServidorGravacaoID)
		if !errors.Is(err, servidorgravacao.ErrNotFound) {
			t.Fatalf("\t%s\tShould NOT be able to retrieve servidor de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to retrieve servidor de gravacao.", tests.Success)
	}

	t.Log("\tGiven the need to page through Servidores de Gravacao records.")
	{
		sv1, err := core.Query(ctx, "", 1, 1)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve servidores de gravacao for page 1: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve servidores de gravacao for page 1.", tests.Success)

		if len(sv1) != 1 {
			t.Fatalf("\t%s\tShould have a single servidor de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have a single servidor de gravacao.", tests.Success)

		sv2, err := core.Query(ctx, "", 2, 1)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve servidores de gravacao for page 2: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve servidores de gravacao for page 2.", tests.Success)

		if len(sv2) != 1 {
			t.Fatalf("\t%s\tShould have a single servidor de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have a single servidor de gravacao.", tests.Success)

		if sv1[0].ServidorGravacaoID == sv2[0].ServidorGravacaoID {
			t.Logf("\t\tServidor1: %v", sv1[0].ServidorGravacaoID)
			t.Logf("\t\tServidor2: %v", sv2[0].ServidorGravacaoID)
			t.Fatalf("\t%s\tShould have different servidores de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have different servidores de gravacao.", tests.Success)
	}
}
