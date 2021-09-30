package usuario_test

import (
	"context"
	"errors"
	"testing"

	"github.com/filipeandrade6/vigia-go/internal/core/usuario"
	"github.com/filipeandrade6/vigia-go/internal/data/store/tests"

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
			t.Fatalf("\t%s\tShould be able to create user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to create user.", tests.Success)

		saved, err := core.QueryByID(ctx, usr.UsuarioID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve user by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve user by ID.", tests.Success)

		if diff := cmp.Diff(usr, saved); diff != "" {
			t.Fatalf("\t%s\tShould get back the same usuario. Diff:\n%s", tests.Failed, diff)
		}
		t.Logf("\t%s\tShould get back the same usuario.", tests.Success)

		upd := usuario.UpdateUsuario{
			Email: tests.StringPointer("filipe@vigia2.com.br"),
		}

		if err = core.Update(ctx, usr.UsuarioID, upd); err != nil {
			t.Fatalf("\t%s\tShould be able to update user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to update user.", tests.Success)

		saved, err = core.QueryByEmail(ctx, *upd.Email)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve user by Email: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve user by Email.", tests.Success)

		if saved.Email != *upd.Email {
			t.Errorf("\t%s\tShould be able to see updates to Email.", tests.Failed)
			t.Logf("\t\tGot: %v", saved.Email)
			t.Logf("\t\tExp: %v", *upd.Email)
		} else {
			t.Logf("\t%s\tShould be able to see updates to Email.", tests.Success)
		}

		if err := core.Delete(ctx, usr.UsuarioID); err != nil {
			t.Fatalf("\t%s\tShould be able to delete usuario: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to delete usuario.", tests.Success)

		_, err = core.QueryByID(ctx, usr.UsuarioID)
		if !errors.Is(err, usuario.ErrNotFound) {
			t.Fatalf("\t%s\tShould NOT be able to retrieve user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to retrieve user.", tests.Success)
	}

	t.Log("\tGiven the need to page through User records.")
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
