package usuario_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/data/store/tests"
	"github.com/filipeandrade6/vigia-go/internal/data/store/usuario"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/go-cmp/cmp"
)

func TestUsuario(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown) // TODO verificar se o teardown funciona...

	usuarioStore := usuario.NewStore(log, db)

	t.Log("Given the need to work with User records.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen handling a single User.", testID)
		{
			ctx := context.Background()

			u := usuario.Usuario{
				Email:  "filipe@teste.com",
				Funcao: []string{auth.RoleAdmin},
				Senha:  "secret",
			}

			claims := auth.Claims{
				StandardClaims: jwt.StandardClaims{
					Issuer:    "service project",
					ExpiresAt: time.Now().Add(time.Hour).Unix(),
					IssuedAt:  time.Now().UTC().Unix(),
				},
				Roles: []string{auth.RoleAdmin},
			}

			usr, err := usuarioStore.Create(ctx, claims, u)
			if err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to create user : %s.", tests.Failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to create user.", tests.Success, testID)

			saved, err := usuarioStore.QueryByID(ctx, claims, usr)
			if err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to retrieve user by ID: %s.", tests.Failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to retrieve user by ID.", tests.Success, testID)

			if diff := cmp.Diff(usr, saved); diff != "" {
				t.Fatalf("\t%s\tTest %d:\tShould get back the same user. Diff:\n%s", tests.Failed, testID, diff)
			}
			t.Logf("\t%s\tTest %d:\tShould get back the same user.", tests.Success, testID)

			u.Email = "filipe@teste2.com"
			u.Senha = "secret2"

			if err := usuarioStore.Update(ctx, claims, u); err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to update user : %s.", tests.Failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to update user.", tests.Success, testID)

			if err := usuarioStore.Delete(ctx, claims, u.UsuarioID); err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to delete user : %s.", tests.Failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to delete user.", tests.Success, testID)

			_, err = usuarioStore.QueryByID(ctx, claims, u.UsuarioID)
			if !errors.As(err, &database.ErrNotFound) {
				t.Fatalf("\t%s\tTest %d:\tShould NOT be able to retrieve user : %s.", tests.Failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould NOT be able to retrieve user.", tests.Success, testID)
		}
	}
}
