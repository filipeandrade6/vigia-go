package camera_test

import (
	"context"
	"errors"
	"testing"

	"github.com/filipeandrade6/vigia-go/internal/data/store/camera"
	"github.com/filipeandrade6/vigia-go/internal/data/store/tests"
	"github.com/filipeandrade6/vigia-go/internal/sys/auth"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"

	"github.com/google/go-cmp/cmp"
)

func TestCamera(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	cameraStore := camera.NewStore(log, db)

	ctx := context.Background()

	claimsAdmin := auth.Claims{Roles: []string{auth.RoleAdmin}}
	claimsManager := auth.Claims{Roles: []string{auth.RoleManager}}
	claimsUser := auth.Claims{Roles: []string{auth.RoleUser}}

	c := camera.Camera{
		Descricao:      "camera testes 1",
		EnderecoIP:     "1.2.3.4",
		Porta:          1234,
		Canal:          1,
		Usuario:        "admin",
		Senha:          "admin",
		Geolocalizacao: "-12.3456, -12.3456",
	}

	t.Log("\tGiven the need to work with Camera records.")
	{
		cameraID, err := cameraStore.Create(ctx, claimsAdmin, c)
		if err != nil {
			t.Fatalf("\t%s\tAdmin should be able to create camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to create camera.", tests.Success)

		c.EnderecoIP = "2.3.4.5"

		if _, err = cameraStore.Create(ctx, claimsManager, c); err != nil {
			t.Fatalf("\t%s\tManager should be able to create camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should be able to create camera.", tests.Success)

		if _, err = cameraStore.Create(ctx, claimsUser, c); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to create camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT beable to create camera.", tests.Success)

		// ---

		cam, err := cameraStore.QueryByID(ctx, claimsAdmin, cameraID)
		if err != nil {
			t.Fatalf("\t%s\tAdmin should be able to retrieve camera by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to retrieve camera by ID.", tests.Success)

		if _, err = cameraStore.QueryByID(ctx, claimsManager, cameraID); err != nil {
			t.Fatalf("\t%s\tManager should be able to retrieve camera by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should be able to retrieve camera by ID.", tests.Success)

		cam2, err := cameraStore.QueryByID(ctx, claimsUser, cameraID)
		if err != nil || cam2.EnderecoIP != "" || cam2.Porta != 0 || cam2.Canal != 0 || cam2.Usuario != "" || cam2.Senha != "" {
			t.Fatalf("\t%s\tUser should be able to retrieve camera by ID without critical info: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should be able to retrieve camera by ID without critical info.", tests.Success)

		if _, err := cameraStore.QueryByID(ctx, claimsAdmin, "bad ID"); !errors.As(err, &database.ErrInvalidID) {
			t.Logf("\t%s\tShould NOT be able to retrieve camera by bad ID: %s", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to retrieve camera by bad ID", tests.Success)

		// ---

		cams, err := cameraStore.Query(ctx, claimsAdmin, "random query", 1, 1)
		if len(cams) != 0 {
			t.Fatalf("\t%s\tShould NOT return any camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT return any camera.", tests.Success)

		// ---

		if diff := cmp.Diff(cameraID, cam.CameraID); diff != "" {
			t.Fatalf("\t%s\tShould get back the same camera. Diff:\n%s", tests.Failed, diff)
		}
		t.Logf("\t%s\tShould get back the same camera.", tests.Success)

		// ---

		c.CameraID = cameraID
		c.EnderecoIP = "210.210.210.210"
		c.Porta = 5678
		c.Canal = 2
		c.Usuario = "manager"
		c.Senha = "manager"
		c.Geolocalizacao = "-23.4567, -23.4567"

		if err = cameraStore.Update(ctx, claimsAdmin, c); err != nil {
			t.Fatalf("\t%s\tAdmin should be able to update camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to update camera.", tests.Success)

		if err = cameraStore.Update(ctx, claimsManager, c); err != nil {
			t.Fatalf("\t%s\tManager should be able to update camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should be able to update camera.", tests.Success)

		if err = cameraStore.Update(ctx, claimsUser, c); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to update camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT be able to update camera.", tests.Success)

		// ---

		if err = cameraStore.Delete(ctx, claimsAdmin, c.CameraID); err != nil {
			t.Fatalf("\t%s\tAdmin should be able to delete camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tAdmin should be able to delete camera.", tests.Success)

		if err = cameraStore.Delete(ctx, claimsManager, c.CameraID); err != nil { // TODO e pra dar erro?
			t.Fatalf("\t%s\tManager should be able to delete camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tManager should be able to delete camera.", tests.Success)

		if err = cameraStore.Delete(ctx, claimsUser, c.CameraID); !errors.As(err, &database.ErrForbidden) {
			t.Fatalf("\t%s\tUser should NOT be able to delete camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tUser should NOT be able to delete camera.", tests.Success)

		// ---

		if _, err = cameraStore.QueryByID(ctx, claimsAdmin, c.CameraID); !errors.As(err, &database.ErrNotFound) {
			t.Fatalf("\t%s\tShould NOT be able to retrieve camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to retrieve camera.", tests.Success)
	}

	t.Log("\tGiven the need to page through Camera records.")
	{
		camera1, err := cameraStore.Query(ctx, claimsAdmin, "", 1, 1)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve cameras for page 1: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve cameras for page 1.", tests.Success)

		if len(camera1) != 1 {
			t.Fatalf("\t%s\tShould have a single camera : %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have a single camera.", tests.Success)

		camera2, err := cameraStore.Query(ctx, claimsAdmin, "", 2, 1)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve cameras for page 2: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve cameras for page 2.", tests.Success)

		if len(camera2) != 1 {
			t.Fatalf("\t%s\tShould have a single camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have a single camera.", tests.Success)

		if camera1[0].CameraID == camera2[0].CameraID {
			t.Logf("\t\tCamera1: %v", camera1[0].CameraID)
			t.Logf("\t\tCamera2: %v", camera2[0].CameraID)
			t.Fatalf("\t%s\tShould have different cameras: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have different cameras.", tests.Success)
	}
}
