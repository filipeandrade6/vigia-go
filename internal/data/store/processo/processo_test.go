package processo_test

import (
	"context"
	"errors"
	"testing"

	"github.com/filipeandrade6/vigia-go/internal/data/store/processo"
	"github.com/filipeandrade6/vigia-go/internal/data/store/tests"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"

	"github.com/google/go-cmp/cmp"
)

func TestProcesso(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	processoStore := processo.NewStore(log, db)

	ctx := context.Background()

	claimsAdmin := auth.Claims{Roles: []string{auth.RoleAdmin}}
	claimsManager := auth.Claims{Roles: []string{auth.RoleManager}}
	claimsUser := auth.Claims{Roles: []string{auth.RoleUser}}

	// TODO ver a necessidade de seed de processo no DB
	// TODO criar as restrições de adição - somente uma câmera com mesmo IP etc.

	p := processo.Processo{
		ServidorGravacaoID: "d03307d4-2b28-4c23-a004-3da25e5b8bb1", // seeded from migration.
		CameraID: "d03307d4-2b28-4c23-a004-3da25e5b8aa3", // seeded from migration.
		Processador: 1,
		Adaptador: 1,
		Execucao: false,
	}

	t.Log("\tGiven the need to work with Processo records.")
	{
		processoID, err := processoStore.Create(ctx, claimsAdmin, c)
		if err != nil {
			t.Fatalf("\t%s\tAdmin should be able to create camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to create camera.", tests.Success)

		c.EnderecoIP = "2.3.4.5"

		if _, err = processoStore.Create(ctx, claimsManager, c); err != nil {
			t.Fatalf("\t%s\tManager should be able to create camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should be able to create camera.", tests.Success)

		if _, err = processoStore.Create(ctx, claimsUser, c); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to create camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT beable to create camera.", tests.Success)
