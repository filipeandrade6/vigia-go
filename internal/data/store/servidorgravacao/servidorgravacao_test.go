package servidorgravacao_test

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/filipeandrade6/vigia-go/internal/data/store/servidorgravacao"
// 	"github.com/filipeandrade6/vigia-go/internal/data/store/tests"
// 	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
// 	"github.com/filipeandrade6/vigia-go/internal/sys/database"
// 	"github.com/google/go-cmp/cmp"
// )

// func TestServidorGravacao(t *testing.T) {
// 	log, db, teardown := tests.New(t)
// 	t.Cleanup(teardown)

// 	servidorGravacaoStore := servidorgravacao.NewStore(log, db)

// 	ctx := context.Background()

// 	claimsAdmin := auth.Claims{Roles: []string{auth.RoleAdmin}}
// 	claimsManager := auth.Claims{Roles: []string{auth.RoleManager}}
// 	claimsUser := auth.Claims{Roles: []string{auth.RoleUser}}

// 	s := servidorgravacao.ServidorGravacao{
// 		EnderecoIP: "10.20.30.40",
// 		Porta:      5001,
// 		Host:       "sv1",
// 	}

// 	t.Log("\tGiven the need to work with Servidores de Gravacao records.")
// 	{
// 		svID, err := servidorGravacaoStore.Create(ctx, claimsAdmin, s)
// 		if err != nil {
// 			t.Fatalf("\t%s\tAdmin should be able to create servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tAdmin should be able to create servidor de gravacao.", tests.Success)

// 		if _, err := servidorGravacaoStore.Create(ctx, claimsAdmin, s); err == nil {
// 			t.Fatalf("\t%s\tShould be able to create servidor de gravacao with existing endereco_ip:porta: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tShould be able to create sservidor de gravacao with existing endereco_ip:porta.", tests.Success)

// 		if _, err := servidorGravacaoStore.Create(ctx, claimsManager, s); !errors.As(err, &database.ErrForbidden) {
// 			t.Fatalf("\t%s\tManager should NOT be able to create servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tManager should NOT be able to create servidor de gravacao.", tests.Success)

// 		if _, err := servidorGravacaoStore.Create(ctx, claimsUser, s); !errors.As(err, &database.ErrForbidden) {
// 			t.Fatalf("\t%s\tUser should NOT be able to create servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tUser should NOT be able to create servidor de gravacao.", tests.Success)

// 		// ---

// 		sv, err := servidorGravacaoStore.QueryByID(ctx, claimsAdmin, svID)
// 		if err != nil {
// 			t.Fatalf("\t%s\tAdmin should be able to retrieve servidor de gravacao by ID: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tAdmin should be able to retrieve servidor de gravacao by ID.", tests.Success)

// 		if _, err := servidorGravacaoStore.QueryByID(ctx, claimsManager, svID); !errors.As(err, &database.ErrForbidden) {
// 			t.Fatalf("\t%s\tManager should NOT be able to retrieve servidor de gravacao by ID: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tManager should NOT be able to retrieve servidor de gravacao by ID.", tests.Success)

// 		if _, err := servidorGravacaoStore.QueryByID(ctx, claimsUser, svID); !errors.As(err, &database.ErrForbidden) {
// 			t.Fatalf("\t%s\tUser should NOT be able to retrieve servidor de gravacao by ID: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tUser should NOT be able to retreive servidor de gravacao by ID.", tests.Success)

// 		if _, err := servidorGravacaoStore.QueryByID(ctx, claimsAdmin, "bad ID"); !errors.As(err, &database.ErrInvalidID) {
// 			t.Logf("\t%s\tShould NOT be able to retrieve servidor de gravacao by bad ID: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tShould NOT be able to retrieve retrieve servidor de gravacao by bad ID", tests.Success)

// 		// ---

// 		svs, err := servidorGravacaoStore.Query(ctx, claimsAdmin, "random query", 1, 1)
// 		if len(svs) != 0 {
// 			t.Fatalf("\t%s\tShould NOT return any servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tShould NOT return any servidor de gravacao.", tests.Success)

// 		// ---

// 		if diff := cmp.Diff(svID, sv.ServidorGravacaoID); diff != "" {
// 			t.Fatalf("\t%s\tShould get back the same servidor de gravacao. Diff:\n%s", tests.Failed, diff)
// 		}
// 		t.Logf("\t%s\tShould get back the same servidor de gravacao.", tests.Success)

// 		// ---

// 		s.ServidorGravacaoID = svID
// 		s.EnderecoIP = "11.22.33.44"
// 		s.Porta = 1005
// 		s.Host = "sv2"

// 		if err = servidorGravacaoStore.Update(ctx, claimsAdmin, s); err != nil {
// 			t.Fatalf("\t%s\tAdmin should be able to update servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tAdmin should be able to update servidor de gravacao.", tests.Success)

// 		if err = servidorGravacaoStore.Update(ctx, claimsManager, s); !errors.As(err, &database.ErrForbidden) {
// 			t.Fatalf("\t%s\tManager should NOT be able to update servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tManager should be able to update servidor de gravacao.", tests.Success)

// 		if err = servidorGravacaoStore.Update(ctx, claimsUser, s); !errors.As(err, &database.ErrForbidden) {
// 			t.Fatalf("\t%s\tUser should NOT be able to update servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tUser should NOT be able to update servidor de gravacao.", tests.Success)

// 		// ---

// 		if err = servidorGravacaoStore.Delete(ctx, claimsAdmin, s.ServidorGravacaoID); err != nil {
// 			t.Fatalf("\t%s\tAdmin should be able to delete servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tAdmin should be able to delete servidor de gravacao.", tests.Success)

// 		if err = servidorGravacaoStore.Delete(ctx, claimsManager, s.ServidorGravacaoID); !errors.As(err, &database.ErrForbidden) { // TODO e pra dar erro?
// 			t.Fatalf("\t%s\tManager should NOT be able to delete servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tManager should NOT be able to delete servidor de gravacao.", tests.Success)

// 		if err = servidorGravacaoStore.Delete(ctx, claimsUser, s.ServidorGravacaoID); !errors.As(err, &database.ErrForbidden) {
// 			t.Fatalf("\t%s\tUser should NOT be able to delete servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tUser should NOT be able to delete servidor de gravacao.", tests.Success)

// 		// ---

// 		if _, err = servidorGravacaoStore.QueryByID(ctx, claimsAdmin, s.ServidorGravacaoID); !errors.As(err, &database.ErrNotFound) {
// 			t.Fatalf("\t%s\tShould NOT be able to retrieve servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tShould NOT be able to retrieve servidor de gravacao.", tests.Success)
// 	}

// 	t.Log("\tGiven the need to page through Servidores de Gravacao records.")
// 	{
// 		sv1, err := servidorGravacaoStore.Query(ctx, claimsAdmin, "", 1, 1)
// 		if err != nil {
// 			t.Fatalf("\t%s\tShould be able to retrieve servidores de gravacao for page 1: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tShould be able to retrieve servidores de gravacao for page 1.", tests.Success)

// 		if len(sv1) != 1 {
// 			t.Fatalf("\t%s\tShould have a single servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tShould have a single servidor de gravacao.", tests.Success)

// 		sv2, err := servidorGravacaoStore.Query(ctx, claimsAdmin, "", 2, 1)
// 		if err != nil {
// 			t.Fatalf("\t%s\tShould be able to retrieve servidores de gravacao for page 2: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tShould be able to retrieve servidores de gravacao for page 2.", tests.Success)

// 		if len(sv2) != 1 {
// 			t.Fatalf("\t%s\tShould have a single servidor de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tShould have a single servidor de gravacao.", tests.Success)

// 		if sv1[0].ServidorGravacaoID == sv2[0].ServidorGravacaoID {
// 			t.Logf("\t\tServidor1: %v", sv1[0].ServidorGravacaoID)
// 			t.Logf("\t\tServidor2: %v", sv2[0].ServidorGravacaoID)
// 			t.Fatalf("\t%s\tShould have different servidores de gravacao: %s.", tests.Failed, err)
// 		}
// 		t.Logf("\t%s\tShould have different servidores de gravacao.", tests.Success)
// 	}
// }
