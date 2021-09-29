package usuario_test

import (
	"context"
	"errors"
	"testing"

	"github.com/filipeandrade6/vigia-go/internal/data/store/tests"
	"github.com/filipeandrade6/vigia-go/internal/data/store/usuario"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/filipeandrade6/vigia-go/internal/sys/validate"

	"github.com/google/go-cmp/cmp"
)

func TestUsuario(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	usuarioStore := usuario.NewStore(log, db)

	ctx := context.Background()

	claimsAdmin := auth.Claims{Roles: []string{auth.RoleAdmin}}
	claimsManager := auth.Claims{Roles: []string{auth.RoleManager}}
	claimsUser := auth.Claims{Roles: []string{auth.RoleUser}}

	t.Log("\tGiven the need to work with User records.")
	{
		u := usuario.Usuario{
			Email:  "filipe@teste.com",
			Funcao: []string{auth.RoleAdmin},
			Senha:  "secret",
		}

		usuarioID, err := usuarioStore.Create(ctx, claimsAdmin, u)
		if err != nil {
			t.Fatalf("\t%s\tAdmin should be able to create user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to create user.", tests.Success)

		if _, err = usuarioStore.Create(ctx, claimsManager, u); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tManager should NOT be able to create user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should NOT be able to create user.", tests.Success)

		if _, err = usuarioStore.Create(ctx, claimsUser, u); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to create user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT beable to create user.", tests.Success)

		// ---

		usr, err := usuarioStore.QueryByID(ctx, claimsAdmin, usuarioID)
		if err != nil {
			t.Fatalf("\t%s\tAdmin should be able to retrieve user by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to retrieve user by ID.", tests.Success)

		if _, err = usuarioStore.QueryByID(ctx, claimsManager, usuarioID); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tManager should NOT be able to retrieve user by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should NOT be able to retrieve user by ID.", tests.Success)

		if _, err = usuarioStore.QueryByID(ctx, claimsUser, usuarioID); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to retrieve user by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT be able to retrieve user by ID.", tests.Success)

		if _, err := usuarioStore.QueryByID(ctx, claimsAdmin, "bad ID"); !errors.As(err, &database.ErrInvalidID) {
			t.Logf("\t%s\tShould NOT be able to retrieve user by bad ID: %s", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to retrieve user by bad ID", tests.Success)

		// ---

		usrs, err := usuarioStore.Query(ctx, claimsAdmin, "query aleatoria", 1, 1)
		if len(usrs) != 0 {
			t.Fatalf("\t%s\tShould NOT return any user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT return any user.", tests.Success)

		if _, err = usuarioStore.Query(ctx, claimsManager, "", 1, 1); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tManager should NOT be able to query users: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should NOT be able to query users.", tests.Success)

		if _, err = usuarioStore.Query(ctx, claimsUser, "", 1, 1); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to query users: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT be able to query users.", tests.Success)

		// ---

		if diff := cmp.Diff(usuarioID, usr.UsuarioID); diff != "" {
			t.Fatalf("\t%s\tShould get back the same user. Diff:\n%s", tests.Failed, diff)
		}
		t.Logf("\t%s\tShould get back the same user.", tests.Success)

		// ---

		u.UsuarioID = usuarioID
		u.Email = "filipe@teste2.com"
		u.Senha = "secret2"

		if err = usuarioStore.Update(ctx, claimsAdmin, u); err != nil {
			t.Fatalf("\t%s\tAdmin should be able to update user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to update user.", tests.Success)

		if err = usuarioStore.Update(ctx, claimsManager, u); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tManager should NOT be able to update user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should NOT be able to update user.", tests.Success)

		if err = usuarioStore.Update(ctx, claimsUser, u); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to update user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT be able to update user.", tests.Success)

		// ---

		uuidTest := validate.GenerateID()
		claimsAdmin.Subject = uuidTest

		if err = usuarioStore.Delete(ctx, claimsAdmin, uuidTest); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tAdmin should NOT be able to delete itself: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should NOT be able to delete itself.", tests.Success)

		if err = usuarioStore.Delete(ctx, claimsAdmin, u.UsuarioID); err != nil {
			t.Fatalf("\t%s\tAdmin should be able to delete user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to delete user.", tests.Success)

		if err = usuarioStore.Delete(ctx, claimsManager, u.UsuarioID); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tManager should NOT be able to delete user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should NOT be able to delete user.", tests.Success)

		if err = usuarioStore.Delete(ctx, claimsUser, u.UsuarioID); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to delete user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT be able to delete user.", tests.Success)

		// ---

		if _, err = usuarioStore.QueryByID(ctx, claimsAdmin, u.UsuarioID); !errors.As(err, &database.ErrNotFound) {
			t.Fatalf("\t%s\tShould NOT be able to retrieve user: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to retrieve user.", tests.Success)
	}

	t.Log("\tGiven the need to page through User records.")
	{
		users1, err := usuarioStore.Query(ctx, claimsAdmin, "", 1, 1)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve users for page 1: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve users for page 1.", tests.Success)

		if len(users1) != 1 {
			t.Fatalf("\t%s\tShould have a single user : %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have a single user.", tests.Success)

		users2, err := usuarioStore.Query(ctx, claimsAdmin, "", 2, 1)
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
