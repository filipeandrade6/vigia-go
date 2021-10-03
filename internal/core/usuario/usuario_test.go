package usuario_test

import (
	"context"
	"errors"
	"testing"

	"github.com/filipeandrade6/vigia-go/internal/core/usuario"
	"github.com/filipeandrade6/vigia-go/internal/database/tests"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/google/go-cmp/cmp"
)

// TODO teste com email repetido

func TestUsuario(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	core := usuario.NewCore(log, db)

	ctx := context.Background()

	t.Log("\tGiven the need to work with Usuario records.")
	{
		nu := usuario.NewUsuario{
			Email:  "filipe@vigia.com.br",
			Funcao: []string{"ADMIN"},
			Senha:  "password",
		}

		usr, err := core.Create(ctx, nu)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to create usuario: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to create usuario.", tests.Success)

		saved, err := core.QueryByID(ctx, usr.UsuarioID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve usuario by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve usuario by ID.", tests.Success)

		if diff := cmp.Diff(usr, saved); diff != "" {
			t.Fatalf("\t%s\tShould get back the same usuario. Diff:\n%s", tests.Failed, diff)
		}
		t.Logf("\t%s\tShould get back the same usuario.", tests.Success)

		nu2 := usuario.NewUsuario{
			Email:  "filipe@vigia.com.br",
			Funcao: []string{"ADMIN"},
			Senha:  "password",
		}

		_, err = core.Create(ctx, nu2)
		if !errors.Is(err, usuario.ErrEmailAlreadyExists) {
			t.Fatalf("\t%s\tShould NOT be able to create usuario with an already existing Email: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to create usuario with an already existing Email.", tests.Success)

		upd := usuario.UpdateUsuario{
			UsuarioID: usr.UsuarioID,
			Email:     &wrapperspb.StringValue{Value: "filipe@vigia2.com.br"},
		}

		if err = core.Update(ctx, upd); err != nil {
			t.Fatalf("\t%s\tShould be able to update usuario: %s.", tests.Failed, err) // TODO dando erro aqui de email ja existe, mas não existe
		}
		t.Logf("\t%s\tShould be able to update usuario.", tests.Success)

		upd2 := usuario.UpdateUsuario{
			UsuarioID: usr.UsuarioID,
			Email:     &wrapperspb.StringValue{Value: "filipe@vigia2.com.br"},
		}

		if err = core.Update(ctx, upd2); !errors.Is(err, usuario.ErrEmailAlreadyExists) {
			t.Fatalf("\t%s\tShould NOT be able to update usuario Email to an already existing Email: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to update usuario Email to an already existing Email.", tests.Success)

		saved, err = core.QueryByEmail(ctx, upd.Email.GetValue())
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve usuario by Email: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve usuario by Email.", tests.Success)

		if saved.Email != upd.Email.GetValue() {
			t.Errorf("\t%s\tShould be able to see updates to Email.", tests.Failed)
			t.Logf("\t\tGot: %v", saved.Email)
			t.Logf("\t\tExp: %v", upd.Email.GetValue())
		} else {
			t.Logf("\t%s\tShould be able to see updates to Email.", tests.Success)
		}

		if err := core.Delete(ctx, usr.UsuarioID); err != nil {
			t.Fatalf("\t%s\tShould be able to delete usuario: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to delete usuario.", tests.Success)

		_, err = core.QueryByID(ctx, usr.UsuarioID)
		if !errors.Is(err, usuario.ErrNotFound) {
			t.Fatalf("\t%s\tShould NOT be able to retrieve usuario: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to retrieve usuario.", tests.Success)
	}

	t.Log("\tGiven the need to page through Usuario records.")
	{
		users1, err := core.Query(ctx, "", 1, 1)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve users for page 1: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve users for page 1.", tests.Success)

		if len(users1) != 1 {
			t.Fatalf("\t%s\tShould have a single user : %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have a single user.", tests.Success)

		users2, err := core.Query(ctx, "", 2, 1)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve users for page 2: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve users for page 2.", tests.Success)

		if len(users2) != 1 {
			t.Fatalf("\t%s\tShould have a single user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have a single user.", tests.Success)

		if users1[0].UsuarioID == users2[0].UsuarioID {
			t.Logf("\t\tUser1: %v", users1[0].UsuarioID)
			t.Logf("\t\tUser2: %v", users2[0].UsuarioID)
			t.Fatalf("\t%s\tShould have different users: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have different users.", tests.Success)
	}
}
