package camera_test

import (
	"context"
	"errors"
	"testing"

	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/data/store/tests"

	"github.com/google/go-cmp/cmp"
)

func TestCamera(t *testing.T) {
	log, db, teardown := tests.New(t)
	t.Cleanup(teardown)

	core := camera.NewCore(log, db)

	ctx := context.Background()

	nc := camera.NewCamera{
		Descricao:  "camera testes 1",
		EnderecoIP: "1.2.3.4",
		Porta:      1234,
		Canal:      1,
		Usuario:    "admin",
		Senha:      "admin",
		Latitude:   "-12.4567",
		Longitude:  "-12.4567",
	}

	t.Log("\tGiven the need to work with Camera records.")
	{
		cam, err := core.Create(ctx, nc)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to create camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to create camera.", tests.Success)

		saved, err := core.QueryByID(ctx, cam.CameraID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve camera by ID: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve camera by ID.", tests.Success)

		if diff := cmp.Diff(cam, saved); diff != "" {
			t.Fatalf("\t%s\tShould get back the same camera. Diff:\n%s", tests.Failed, diff)
		}
		t.Logf("\t%s\tShould get back the same camera.", tests.Success)

		upd := camera.UpdateCamera{
			Descricao:  tests.StringPointer("udpated"),
			EnderecoIP: tests.StringPointer("123.123.210.210"),
			Porta:      tests.IntPointer(30),
			Canal:      tests.IntPointer(8),
			Usuario:    tests.StringPointer("manager"),
			Senha:      tests.StringPointer("manager"),
			Latitude:   tests.StringPointer("-23.4567"),
			Longitude:  tests.StringPointer("-23.4567"),
		}

		if err = core.Update(ctx, cam.CameraID, upd); err != nil {
			t.Fatalf("\t%s\tShould be able to update camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to update camera.", tests.Success)

		cams, err := core.Query(ctx, "", 1, 3)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve updated camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve updated camera.", tests.Success)

		want := cam
		want.Descricao = *upd.Descricao
		want.EnderecoIP = *upd.EnderecoIP
		want.Porta = *upd.Porta
		want.Canal = *upd.Canal
		want.Usuario = *upd.Usuario
		want.Senha = *upd.Senha
		want.Latitude = *upd.Latitude
		want.Longitude = *upd.Longitude

		var idx int
		for i, c := range cams {
			if c.CameraID == want.CameraID {
				idx = i
			}
		}
		if diff := cmp.Diff(want, cams[idx]); diff != "" {
			t.Fatalf("\t%s\tShould get back the same camera. Diff\n%s", tests.Failed, diff)
		}
		t.Logf("\t%s\tShould get back the same camera.", tests.Success)

		upd = camera.UpdateCamera{
			Porta: tests.IntPointer(54321),
		}

		if err = core.Update(ctx, cam.CameraID, upd); err != nil {
			t.Fatalf("\t%s\tShould be able to update just some fields of camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to update just some fields of camera.", tests.Success)

		saved, err = core.QueryByID(ctx, cam.CameraID)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve updated camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve updated camera.", tests.Success)

		t.Logf("%+v", saved)

		if saved.Porta != *upd.Porta {
			t.Fatalf("\t%s\tShould be able to see updated Porta field: got %q want %q.", tests.Failed, saved.Porta, *upd.Porta)
		}
		t.Logf("\t%s\tShould be able to see updated Porta field.", tests.Success)

		if err = core.Delete(ctx, cam.CameraID); err != nil {
			t.Fatalf("\t%s\tShould be able to delete camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to delete camera.", tests.Success)

		_, err = core.QueryByID(ctx, cam.CameraID)
		if !errors.Is(err, camera.ErrNotFound) {
			t.Fatalf("\t%s\tShould NOT be able to retrieve deleted camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould NOT be able to retrieve deleted camera.", tests.Success)
	}

	t.Log("\tGiven the need to page through Camera records.")
	{
		cam1, err := core.Query(ctx, "", 1, 1)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve camera for page 1: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve camera for page 1.", tests.Success)

		if len(cam1) != 1 {
			t.Fatalf("\t%s\tShould have a single servidor de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have a single servidor de gravacao.", tests.Success)

		cam2, err := core.Query(ctx, "", 2, 1)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve camera for page 2: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve camera for page 2.", tests.Success)

		if len(cam2) != 1 {
			t.Fatalf("\t%s\tShould have a single servidor de gravacao: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have a single servidor de gravacao.", tests.Success)

		if cam1[0].CameraID == cam2[0].CameraID {
			t.Logf("\t\tServidor1: %v", cam1[0].CameraID)
			t.Logf("\t\tServidor2: %v", cam2[0].CameraID)
			t.Fatalf("\t%s\tShould have different camera: %s.", tests.Failed, err)
		}
		t.Logf("\t%s\tShould have different camera.", tests.Success)
	}
}
