package servidorgravacao_test

import (
	"context"
	"errors"
	"testing"

	"github.com/filipeandrade6/vigia-go/internal/data/store/servidorgravacao"
	"github.com/filipeandrade6/vigia-go/internal/data/store/tests"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"
)

func TestServidorGravacao(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	servidorGravacaoStore := servidorgravacao.NewStore(log, db)

	ctx := context.Background()

	claimsAdmin := auth.Claims{Roles: []string{auth.RoleAdmin}}
	claimsManager := auth.Claims{Roles: []string{auth.RoleManager}}
	claimsUser := auth.Claims{Roles: []string{auth.RoleUser}}

	s := servidorgravacao.ServidorGravacao{
		EnderecoIP: "10.20.30.40",
		Porta:      5001,
		Host:       "sv1",
	}

	t.Log("\tGiven the need to work with Servidores Gravacao records.")
	{
		svID, err := servidorGravacaoStore.Create(ctx, claimsAdmin, s)
		if err != nil {
			t.Fatalf("\t%s\tAdmin should be able to create servidor de gravacao: %s", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to create servidor de gravacao.", tests.Success)

		if _, err := servidorGravacaoStore.Create(ctx, claimsManager, s); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tManager should NOT be able to create servidor de gravacao: %s", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should NOT be able to create servidor de gravacao.", tests.Success)

		if _, err := servidorGravacaoStore.Create(ctx, claimsUser, s); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to create servidor de gravacao: %s", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT be able to create servidor de gravacao.", tests.Success)

		// ---

		sv, err := servidorGravacaoStore.QueryByID(ctx, claimsAdmin, svID)
		if err != nil {
			t.Fatalf("\t%s\tAdmin should be able to retrieve servidor de gravacao by ID: %s", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to retrieve servidor de gravacao by ID.", tests.Success)

		if _, err := servidorGravacaoStore.QueryByID(ctx, claimsManager, svID); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tManager should NOT be able to retrieve servidor de gravacao by ID: %s", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should NOT be able to retrieve servidor de gravacao by ID.", tests.Success)

		if _, err := servidorGravacaoStore.QueryByID(ctx, claimsUser, svID); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to retrieve servidor de gravacao by ID: %s", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT be able to retreive servidor de gravacao by ID.", tests.Success)

		if _, err := servidorGravacaoStore.QueryByID(ctx, claimsAdmin, "bad ID"); !errors.As(err, &database.ErrInvalidID) {
			t.Logf("\t%s\tShould NOT be able to retrieve servidor de gravacao by bad ID: %s", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to retrieve retrieve servidor de gravacao by bad ID", tests.Success)

		// ---

		svs, err := servidorGravacaoStore.QueryByID(ctx)
	}
}
