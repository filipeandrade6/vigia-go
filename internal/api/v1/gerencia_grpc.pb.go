// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GerenciaClient is the client API for Gerencia service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GerenciaClient interface {
	// Gravacao requests
	Match(ctx context.Context, in *MatchReq, opts ...grpc.CallOption) (*MatchRes, error)
	// Gerencia Client requests
	Migrate(ctx context.Context, in *MigrateReq, opts ...grpc.CallOption) (*MigrateRes, error)
	CreateServidorGravacao(ctx context.Context, in *CreateServidorGravacaoReq, opts ...grpc.CallOption) (*CreateServidorGravacaoRes, error)
	ReadServidorGravacao(ctx context.Context, in *ReadServidorGravacaoReq, opts ...grpc.CallOption) (*ReadServidorGravacaoRes, error)
	UpdateServidorGravacao(ctx context.Context, in *UpdateServidorGravacaoReq, opts ...grpc.CallOption) (*UpdateServidorGravacaoRes, error)
	DeleteServidorGravacao(ctx context.Context, in *DeleteServidorGravacaoReq, opts ...grpc.CallOption) (*DeleteServidorGravacaoRes, error)
	CreateCamera(ctx context.Context, in *CreateCameraReq, opts ...grpc.CallOption) (*CreateCameraRes, error)
	ReadCamera(ctx context.Context, in *ReadCameraReq, opts ...grpc.CallOption) (*ReadCameraRes, error)
	UpdateCamera(ctx context.Context, in *UpdateCameraReq, opts ...grpc.CallOption) (*UpdateCameraRes, error)
	DeleteCamera(ctx context.Context, in *DeleteCameraReq, opts ...grpc.CallOption) (*DeleteCameraRes, error)
	CreateProcesso(ctx context.Context, in *CreateProcessoReq, opts ...grpc.CallOption) (*CreateProcessoRes, error)
	ReadProcesso(ctx context.Context, in *ReadProcessoReq, opts ...grpc.CallOption) (*ReadProcessoRes, error)
	UpdateProcesso(ctx context.Context, in *UpdateProcessoReq, opts ...grpc.CallOption) (*UpdateProcessoRes, error)
	DeleteProcesso(ctx context.Context, in *DeleteProcessoReq, opts ...grpc.CallOption) (*DeleteProcessoRes, error)
}

type gerenciaClient struct {
	cc grpc.ClientConnInterface
}

func NewGerenciaClient(cc grpc.ClientConnInterface) GerenciaClient {
	return &gerenciaClient{cc}
}

func (c *gerenciaClient) Match(ctx context.Context, in *MatchReq, opts ...grpc.CallOption) (*MatchRes, error) {
	out := new(MatchRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/Match", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) Migrate(ctx context.Context, in *MigrateReq, opts ...grpc.CallOption) (*MigrateRes, error) {
	out := new(MigrateRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/Migrate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) CreateServidorGravacao(ctx context.Context, in *CreateServidorGravacaoReq, opts ...grpc.CallOption) (*CreateServidorGravacaoRes, error) {
	out := new(CreateServidorGravacaoRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/CreateServidorGravacao", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) ReadServidorGravacao(ctx context.Context, in *ReadServidorGravacaoReq, opts ...grpc.CallOption) (*ReadServidorGravacaoRes, error) {
	out := new(ReadServidorGravacaoRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/ReadServidorGravacao", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) UpdateServidorGravacao(ctx context.Context, in *UpdateServidorGravacaoReq, opts ...grpc.CallOption) (*UpdateServidorGravacaoRes, error) {
	out := new(UpdateServidorGravacaoRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/UpdateServidorGravacao", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) DeleteServidorGravacao(ctx context.Context, in *DeleteServidorGravacaoReq, opts ...grpc.CallOption) (*DeleteServidorGravacaoRes, error) {
	out := new(DeleteServidorGravacaoRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/DeleteServidorGravacao", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) CreateCamera(ctx context.Context, in *CreateCameraReq, opts ...grpc.CallOption) (*CreateCameraRes, error) {
	out := new(CreateCameraRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/CreateCamera", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) ReadCamera(ctx context.Context, in *ReadCameraReq, opts ...grpc.CallOption) (*ReadCameraRes, error) {
	out := new(ReadCameraRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/ReadCamera", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) UpdateCamera(ctx context.Context, in *UpdateCameraReq, opts ...grpc.CallOption) (*UpdateCameraRes, error) {
	out := new(UpdateCameraRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/UpdateCamera", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) DeleteCamera(ctx context.Context, in *DeleteCameraReq, opts ...grpc.CallOption) (*DeleteCameraRes, error) {
	out := new(DeleteCameraRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/DeleteCamera", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) CreateProcesso(ctx context.Context, in *CreateProcessoReq, opts ...grpc.CallOption) (*CreateProcessoRes, error) {
	out := new(CreateProcessoRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/CreateProcesso", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) ReadProcesso(ctx context.Context, in *ReadProcessoReq, opts ...grpc.CallOption) (*ReadProcessoRes, error) {
	out := new(ReadProcessoRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/ReadProcesso", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) UpdateProcesso(ctx context.Context, in *UpdateProcessoReq, opts ...grpc.CallOption) (*UpdateProcessoRes, error) {
	out := new(UpdateProcessoRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/UpdateProcesso", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) DeleteProcesso(ctx context.Context, in *DeleteProcessoReq, opts ...grpc.CallOption) (*DeleteProcessoRes, error) {
	out := new(DeleteProcessoRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/DeleteProcesso", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GerenciaServer is the server API for Gerencia service.
// All implementations must embed UnimplementedGerenciaServer
// for forward compatibility
type GerenciaServer interface {
	// Gravacao requests
	Match(context.Context, *MatchReq) (*MatchRes, error)
	// Gerencia Client requests
	Migrate(context.Context, *MigrateReq) (*MigrateRes, error)
	CreateServidorGravacao(context.Context, *CreateServidorGravacaoReq) (*CreateServidorGravacaoRes, error)
	ReadServidorGravacao(context.Context, *ReadServidorGravacaoReq) (*ReadServidorGravacaoRes, error)
	UpdateServidorGravacao(context.Context, *UpdateServidorGravacaoReq) (*UpdateServidorGravacaoRes, error)
	DeleteServidorGravacao(context.Context, *DeleteServidorGravacaoReq) (*DeleteServidorGravacaoRes, error)
	CreateCamera(context.Context, *CreateCameraReq) (*CreateCameraRes, error)
	ReadCamera(context.Context, *ReadCameraReq) (*ReadCameraRes, error)
	UpdateCamera(context.Context, *UpdateCameraReq) (*UpdateCameraRes, error)
	DeleteCamera(context.Context, *DeleteCameraReq) (*DeleteCameraRes, error)
	CreateProcesso(context.Context, *CreateProcessoReq) (*CreateProcessoRes, error)
	ReadProcesso(context.Context, *ReadProcessoReq) (*ReadProcessoRes, error)
	UpdateProcesso(context.Context, *UpdateProcessoReq) (*UpdateProcessoRes, error)
	DeleteProcesso(context.Context, *DeleteProcessoReq) (*DeleteProcessoRes, error)
	mustEmbedUnimplementedGerenciaServer()
}

// UnimplementedGerenciaServer must be embedded to have forward compatible implementations.
type UnimplementedGerenciaServer struct {
}

func (UnimplementedGerenciaServer) Match(context.Context, *MatchReq) (*MatchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Match not implemented")
}
func (UnimplementedGerenciaServer) Migrate(context.Context, *MigrateReq) (*MigrateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Migrate not implemented")
}
func (UnimplementedGerenciaServer) CreateServidorGravacao(context.Context, *CreateServidorGravacaoReq) (*CreateServidorGravacaoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServidorGravacao not implemented")
}
func (UnimplementedGerenciaServer) ReadServidorGravacao(context.Context, *ReadServidorGravacaoReq) (*ReadServidorGravacaoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadServidorGravacao not implemented")
}
func (UnimplementedGerenciaServer) UpdateServidorGravacao(context.Context, *UpdateServidorGravacaoReq) (*UpdateServidorGravacaoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServidorGravacao not implemented")
}
func (UnimplementedGerenciaServer) DeleteServidorGravacao(context.Context, *DeleteServidorGravacaoReq) (*DeleteServidorGravacaoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServidorGravacao not implemented")
}
func (UnimplementedGerenciaServer) CreateCamera(context.Context, *CreateCameraReq) (*CreateCameraRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCamera not implemented")
}
func (UnimplementedGerenciaServer) ReadCamera(context.Context, *ReadCameraReq) (*ReadCameraRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCamera not implemented")
}
func (UnimplementedGerenciaServer) UpdateCamera(context.Context, *UpdateCameraReq) (*UpdateCameraRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCamera not implemented")
}
func (UnimplementedGerenciaServer) DeleteCamera(context.Context, *DeleteCameraReq) (*DeleteCameraRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCamera not implemented")
}
func (UnimplementedGerenciaServer) CreateProcesso(context.Context, *CreateProcessoReq) (*CreateProcessoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProcesso not implemented")
}
func (UnimplementedGerenciaServer) ReadProcesso(context.Context, *ReadProcessoReq) (*ReadProcessoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadProcesso not implemented")
}
func (UnimplementedGerenciaServer) UpdateProcesso(context.Context, *UpdateProcessoReq) (*UpdateProcessoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProcesso not implemented")
}
func (UnimplementedGerenciaServer) DeleteProcesso(context.Context, *DeleteProcessoReq) (*DeleteProcessoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProcesso not implemented")
}
func (UnimplementedGerenciaServer) mustEmbedUnimplementedGerenciaServer() {}

// UnsafeGerenciaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GerenciaServer will
// result in compilation errors.
type UnsafeGerenciaServer interface {
	mustEmbedUnimplementedGerenciaServer()
}

func RegisterGerenciaServer(s grpc.ServiceRegistrar, srv GerenciaServer) {
	s.RegisterService(&Gerencia_ServiceDesc, srv)
}

func _Gerencia_Match_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).Match(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/Match",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).Match(ctx, req.(*MatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_Migrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).Migrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/Migrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).Migrate(ctx, req.(*MigrateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_CreateServidorGravacao_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServidorGravacaoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).CreateServidorGravacao(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/CreateServidorGravacao",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).CreateServidorGravacao(ctx, req.(*CreateServidorGravacaoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_ReadServidorGravacao_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadServidorGravacaoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).ReadServidorGravacao(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/ReadServidorGravacao",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).ReadServidorGravacao(ctx, req.(*ReadServidorGravacaoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_UpdateServidorGravacao_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServidorGravacaoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).UpdateServidorGravacao(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/UpdateServidorGravacao",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).UpdateServidorGravacao(ctx, req.(*UpdateServidorGravacaoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_DeleteServidorGravacao_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServidorGravacaoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).DeleteServidorGravacao(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/DeleteServidorGravacao",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).DeleteServidorGravacao(ctx, req.(*DeleteServidorGravacaoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_CreateCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCameraReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).CreateCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/CreateCamera",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).CreateCamera(ctx, req.(*CreateCameraReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_ReadCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCameraReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).ReadCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/ReadCamera",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).ReadCamera(ctx, req.(*ReadCameraReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_UpdateCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCameraReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).UpdateCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/UpdateCamera",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).UpdateCamera(ctx, req.(*UpdateCameraReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_DeleteCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCameraReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).DeleteCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/DeleteCamera",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).DeleteCamera(ctx, req.(*DeleteCameraReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_CreateProcesso_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProcessoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).CreateProcesso(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/CreateProcesso",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).CreateProcesso(ctx, req.(*CreateProcessoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_ReadProcesso_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadProcessoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).ReadProcesso(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/ReadProcesso",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).ReadProcesso(ctx, req.(*ReadProcessoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_UpdateProcesso_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProcessoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).UpdateProcesso(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/UpdateProcesso",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).UpdateProcesso(ctx, req.(*UpdateProcessoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gerencia_DeleteProcesso_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProcessoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).DeleteProcesso(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/DeleteProcesso",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).DeleteProcesso(ctx, req.(*DeleteProcessoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Gerencia_ServiceDesc is the grpc.ServiceDesc for Gerencia service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gerencia_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gerencia.Gerencia",
	HandlerType: (*GerenciaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Match",
			Handler:    _Gerencia_Match_Handler,
		},
		{
			MethodName: "Migrate",
			Handler:    _Gerencia_Migrate_Handler,
		},
		{
			MethodName: "CreateServidorGravacao",
			Handler:    _Gerencia_CreateServidorGravacao_Handler,
		},
		{
			MethodName: "ReadServidorGravacao",
			Handler:    _Gerencia_ReadServidorGravacao_Handler,
		},
		{
			MethodName: "UpdateServidorGravacao",
			Handler:    _Gerencia_UpdateServidorGravacao_Handler,
		},
		{
			MethodName: "DeleteServidorGravacao",
			Handler:    _Gerencia_DeleteServidorGravacao_Handler,
		},
		{
			MethodName: "CreateCamera",
			Handler:    _Gerencia_CreateCamera_Handler,
		},
		{
			MethodName: "ReadCamera",
			Handler:    _Gerencia_ReadCamera_Handler,
		},
		{
			MethodName: "UpdateCamera",
			Handler:    _Gerencia_UpdateCamera_Handler,
		},
		{
			MethodName: "DeleteCamera",
			Handler:    _Gerencia_DeleteCamera_Handler,
		},
		{
			MethodName: "CreateProcesso",
			Handler:    _Gerencia_CreateProcesso_Handler,
		},
		{
			MethodName: "ReadProcesso",
			Handler:    _Gerencia_ReadProcesso_Handler,
		},
		{
			MethodName: "UpdateProcesso",
			Handler:    _Gerencia_UpdateProcesso_Handler,
		},
		{
			MethodName: "DeleteProcesso",
			Handler:    _Gerencia_DeleteProcesso_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gerencia.proto",
}
