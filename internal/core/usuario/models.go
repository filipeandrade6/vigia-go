package usuario

import (
	"unsafe"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/core/usuario/db"
)

type Usuario struct {
	UsuarioID string   `json:"usuario_id"`
	Email     string   `json:"email"`
	Funcao    []string `json:"funcao"`
	Senha     string   `json:"senha"`
}

type NewUsuario struct {
	Email  string   `json:"email" validate:"required,email"`
	Funcao []string `json:"funcao" validate:"required"`
	Senha  string   `json:"senha" validate:"required"`
}

type UpdateUsuario struct {
	Email  *string  `json:"email" validate:"omitempty,email"`
	Funcao []string `json:"funcao"`
	Senha  *string  `json:"senha"`
}

// =============================================================================

func toUsuario(dbUsr db.Usuario) Usuario {
	pu := (*Usuario)(unsafe.Pointer(&dbUsr))
	return *pu
}

func toUsuarioSlice(dbUsrs []db.Usuario) []Usuario {
	usuarios := make([]Usuario, len(dbUsrs))
	for i, dbUsr := range dbUsrs {
		usuarios[i] = toUsuario(dbUsr)
	}
	return usuarios
}

// =============================================================================

func (u Usuario) ToProto() *pb.Usuario {
	return &pb.Usuario{
		UsuarioId: u.UsuarioID,
		Email:     u.Email,
		Funcao:    u.Funcao,
		Senha:     u.Senha,
	}
}

func FromProto(u *pb.Usuario) Usuario {
	return Usuario{
		UsuarioID: u.GetUsuarioId(),
		Email:     u.GetEmail(),
		Funcao:    u.GetFuncao(),
		Senha:     u.GetSenha(),
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
