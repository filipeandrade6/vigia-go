// Package product contains product related CRUD functionality.
package camera

import (
	"context"
	"fmt"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Store manages the set of API's for product access.
type Store struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

// NewStore constructs a product store for api access.
func NewStore(log *zap.SugaredLogger, db *sqlx.DB) Store {
	return Store{
		log: log,
		db:  db,
	}
}

// Create adds a Camera to the database. It returns the created Camera with
// fields like ID and DateCreated populated.
func (s Store) Create(ctx context.Context, cam Camera, now time.Time) (Camera, error) {
	// TODO validate
	// if err := validate.Check(cam); err != nil {
	// 	return Camera{}, fmt.Errorf("validating data: %w", err)
	// }

	c := Camera{
		Nome:            cam.Nome,
		IP:              cam.IP,
		Porta:           cam.Porta,
		Canal:           cam.Canal,
		Usuario:         cam.Usuario,
		Senha:           cam.Senha,
		Regiao:          cam.Regiao,
		Geolocalizacao:  cam.Geolocalizacao,
		Marca:           cam.Marca,
		Modelo:          cam.Modelo,
		Informacao:      cam.Informacao,
		DataCriacao:     now,
		DataAtualizacao: now,
	}

	const q = `
	INSERT INTO cameras
		(camera_id, nome, ip, porta, canal, usuario, senha, regiao, geolocalizacao, marca, modelo, informacao)
	VALUES
		(:camera_id, :nome, :ip, :porta, :canal, :usuario, :senha, :regiao, :geolocalizacao, :marca, :modelo, :informacao)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, c); err != nil {
		return Camera{}, fmt.Errorf("inserting camera: %w", err)
	}

	return cam, nil
}

// TODO UpdateCamera struct?
// Update modifies data about a Camera. It will error if the specified ID is
// invalid or does not reference an existing Product.
// func (s Store) Update(ctx context.Context, claims auth.Claims, productID string, up UpdateProduct, now time.Time) error {
// func (s Store) Update(ctx context.Context, cameraID string, up Camera, now time.Time) error {
// 	if err := validate.CheckID(cameraID); err != nil {
// 		return database.ErrInvalidID
// 	}
// 	if err := validate.Check(up); err != nil {
// 		return fmt.Errorf("validating data: %w", err)
// 	}

// 	cam, err := s.QueryByID(ctx, cameraID)
// 	if err != nil {
// 		return fmt.Errorf("updating camera cameraID[%s]: %w", cameraID, err)
// 	}

// 	// // If you are not an admin and looking to retrieve someone elses product.
// 	// if !claims.Authorized(auth.RoleAdmin) && prd.UserID != claims.Subject {
// 	// 	return database.ErrForbidden
// 	// }

// 	if up.Nome != nil {
// 		cam.Name = up.Nome
// 	}
// 	if up.IP != nil {
// 		cam.IP = up.IP
// 	}

// 	cam.DataAtualizacao = now

// 	const q = `
// 	UPDATE
// 		products
// 	SET
// 		"name" = :name,
// 		"cost" = :cost,
// 		"quantity" = :quantity,
// 		"date_updated" = :date_updated
// 	WHERE
// 		product_id = :product_id`

// 	if err := database.NamedExecContext(ctx, s.log, s.db, q, cam); err != nil {
// 		return fmt.Errorf("updating product ID[%s]: %w", cam.ID, err)
// 	}

// 	return nil
// }

func (s Store) QueryByID(ctx context.Context, cameraID string) (Camera, error) {
	// TODO validate
	// if err := validate.CheckID(cameraID); err != nil {
	// 	return Camera{}, database.ErrInvalidID
	// }

	data := struct {
		CameraID string `db: camera_id`
	}{
		CameraID: cameraID,
	}

	// TODO retornoar leituras, consumo de banda, tempo de atividade, etc.
	const q = `
	SELECT
		*
	FROM
		cameras
	WHERE
		camera_id = :camera_id`

	var cam Camera
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &cam); err != nil {
		if err == database.ErrNotFound {
			return Camera{}, database.ErrNotFound
		}
		return Camera{}, fmt.Errorf("selecting camera ID[%q]: %w", data.CameraID, err)
	}

	return cam, nil
}

// // Delete removes the product identified by a given ID.
// func (s Store) Delete(ctx context.Context, claims auth.Claims, productID string) error {
// 	if err := validate.CheckID(productID); err != nil {
// 		return database.ErrInvalidID
// 	}

// 	// If you are not an admin.
// 	if !claims.Authorized(auth.RoleAdmin) {
// 		return database.ErrForbidden
// 	}

// 	data := struct {
// 		ProductID string `db:"product_id"`
// 	}{
// 		ProductID: productID,
// 	}

// 	const q = `
// 	DELETE FROM
// 		products
// 	WHERE
// 		product_id = :product_id`

// 	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
// 		return fmt.Errorf("deleting product ID[%s]: %w", data.ProductID, err)
// 	}

// 	return nil
// }

// // Query gets all Products from the database.
// func (s Store) Query(ctx context.Context, pageNumber int, rowsPerPage int) ([]Product, error) {
// 	data := struct {
// 		Offset      int `db:"offset"`
// 		RowsPerPage int `db:"rows_per_page"`
// 	}{
// 		Offset:      (pageNumber - 1) * rowsPerPage,
// 		RowsPerPage: rowsPerPage,
// 	}

// 	const q = `
// 	SELECT
// 		p.*,
// 		COALESCE(SUM(s.quantity) ,0) AS sold,
// 		COALESCE(SUM(s.paid), 0) AS revenue
// 	FROM
// 		products AS p
// 	LEFT JOIN
// 		sales AS s ON p.product_id = s.product_id
// 	GROUP BY
// 		p.product_id
// 	ORDER BY
// 		user_id
// 	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

// 	var products []
// 	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &products); err != nil {
// 		if err == database.ErrNotFound {
// 			return nil, database.ErrNotFound
// 		}
// 		return nil, fmt.Errorf("selecting products: %w", err)
// 	}

// 	return products, nil
// }

// // QueryByID finds the product identified by a given ID.
// func (s Store) QueryByID(ctx context.Context, productID string) (Product, error) {
// 	if err := validate.CheckID(productID); err != nil {
// 		return Product{}, database.ErrInvalidID
// 	}

// 	data := struct {
// 		ProductID string `db:"product_id"`
// 	}{
// 		ProductID: productID,
// 	}

// 	const q = `
// 	SELECT
// 		p.*,
// 		COALESCE(SUM(s.quantity), 0) AS sold,
// 		COALESCE(SUM(s.paid), 0) AS revenue
// 	FROM
// 		products AS p
// 	LEFT JOIN
// 		sales AS s ON p.product_id = s.product_id
// 	WHERE
// 		p.product_id = :product_id
// 	GROUP BY
// 		p.product_id`

// 	var prd Product
// 	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &prd); err != nil {
// 		if err == database.ErrNotFound {
// 			return Product{}, database.ErrNotFound
// 		}
// 		return Product{}, fmt.Errorf("selecting product ID[%q]: %w", data.ProductID, err)
// 	}

// 	return prd, nil
// }

// // QueryByUserID finds the product identified by a given User ID.
// func (s Store) QueryByUserID(ctx context.Context, userID string) ([]Product, error) {
// 	if err := validate.CheckID(userID); err != nil {
// 		return nil, database.ErrInvalidID
// 	}

// 	data := struct {
// 		UserID string `db:"user_id"`
// 	}{
// 		UserID: userID,
// 	}

// 	const q = `
// 	SELECT
// 		p.*,
// 		COALESCE(SUM(s.quantity), 0) AS sold,
// 		COALESCE(SUM(s.paid), 0) AS revenue
// 	FROM
// 		products AS p
// 	LEFT JOIN
// 		sales AS s ON p.product_id = s.product_id
// 	WHERE
// 		p.user_id = :user_id
// 	GROUP BY
// 		p.product_id`

// 	var products []Product
// 	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &products); err != nil {
// 		if err == database.ErrNotFound {
// 			return nil, database.ErrNotFound
// 		}
// 		return nil, fmt.Errorf("selecting products: %w", err)
// 	}

// 	return products, nil
// }

// TODO aqui so para demonstrar
// func (s *ServidorGravacao) ToProtobuf() *pb.RegistrarServidorGravacaoReq {
// 	return &pb.RegistrarServidorGravacaoReq{
// 		Id:         s.ID,
// 		EnderecoIp: s.EnderecoIP,
// 		Porta:      s.Porta,
// 	}
// }

// func (s *ServidorGravacao) FromProtobuf(sv *pb.RegistrarServidorGravacaoResp) {
// 	s.ID = sv.GetId()
// }
