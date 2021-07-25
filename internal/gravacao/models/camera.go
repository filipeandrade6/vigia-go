package models

type Camera struct {
	ID             int
	Descricao      string
	IP             string
	Porta          int
	Canal          int
	UsuarioCamera  string
	SenhaCamera    string
	Cidade         string
	Geolocalizacao string
	Marca          string
	Modelo         string
	Informacao     string
}

// -- Table: public.cameras

// -- DROP TABLE public.cameras;

// CREATE TABLE IF NOT EXISTS public.cameras
// (
//     id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
//     descricao text COLLATE pg_catalog."default" NOT NULL,
//     ip text COLLATE pg_catalog."default" NOT NULL,
//     porta integer NOT NULL,
//     canal integer NOT NULL,
//     usuario_camera text COLLATE pg_catalog."default" NOT NULL,
//     senha_camera text COLLATE pg_catalog."default" NOT NULL,
//     cidade text COLLATE pg_catalog."default" NOT NULL,
//     geolocalizacao text COLLATE pg_catalog."default" NOT NULL,
//     marca text COLLATE pg_catalog."default" NOT NULL,
//     modelo text COLLATE pg_catalog."default" NOT NULL,
//     informacao text COLLATE pg_catalog."default" NOT NULL,
//     CONSTRAINT cameras_pkey PRIMARY KEY (id)
// )

// TABLESPACE pg_default;

// ALTER TABLE public.cameras
//     OWNER to postgres;
