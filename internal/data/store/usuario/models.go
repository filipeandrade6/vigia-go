package usuario

import (
	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
)

type Usuario struct {
	UsuarioID string   `db:"usuario_id"`
	Email     string   `db:"email"`
	Funcao    []string `db:"funcao"`
	SenhaHash string   `db:senha_hash`
}

func (u Usuario) ToProto() *pb.Usuario {
	return &pb.Usuario{
		UsuarioId: u.UsuarioID,
		Email:     u.Email,
		Funcao:    u.Funcao,
		SenhaHash: u.SenhaHash,
	}
}

func FromProto(u *pb.Usuario) Usuario {
	return Usuario{
		UsuarioID: u.GetUsuarioId(),
		Email:     u.GetEmail(),
		Funcao:    u.GetFuncao(),
		SenhaHash: u.GetSenhaHash(),
	}
}

type Usuarios []Usuario

func (u Usuarios) ToProto() []*pb.Usuario {
	var usuarios []*pb.Usuario

	for _, usuario := range u {
		usuarios = append(usuarios, usuario.ToProto())
	}

	return usuarios
}

func UsuariosFromProto(u []*pb.Usuario) Usuarios {
	var usuarios Usuarios

	for _, usuario := range u {
		usuarios = append(usuarios, FromProto(usuario))
	}

	return usuarios
}
