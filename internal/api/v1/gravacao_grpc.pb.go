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

// GravacaoClient is the client API for Gravacao service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GravacaoClient interface {
	InfoProcessos(ctx context.Context, in *InfoProcessosReq, opts ...grpc.CallOption) (*InfoProcessosResp, error)
	ConfigurarProcesso(ctx context.Context, in *ConfigurarProcessoReq, opts ...grpc.CallOption) (*ConfigurarProcessoResp, error)
	AtualizarListaVeiculos(ctx context.Context, in *AtualizarListaVeiculosReq, opts ...grpc.CallOption) (*AtualizarListaVeiculosResp, error)
	IniciarProcessamento(ctx context.Context, in *IniciarProcessamentoReq, opts ...grpc.CallOption) (*IniciarProcessamentoResp, error)
}

type gravacaoClient struct {
	cc grpc.ClientConnInterface
}

func NewGravacaoClient(cc grpc.ClientConnInterface) GravacaoClient {
	return &gravacaoClient{cc}
}

func (c *gravacaoClient) InfoProcessos(ctx context.Context, in *InfoProcessosReq, opts ...grpc.CallOption) (*InfoProcessosResp, error) {
	out := new(InfoProcessosResp)
	err := c.cc.Invoke(ctx, "/gravacao.Gravacao/InfoProcessos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gravacaoClient) ConfigurarProcesso(ctx context.Context, in *ConfigurarProcessoReq, opts ...grpc.CallOption) (*ConfigurarProcessoResp, error) {
	out := new(ConfigurarProcessoResp)
	err := c.cc.Invoke(ctx, "/gravacao.Gravacao/ConfigurarProcesso", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gravacaoClient) AtualizarListaVeiculos(ctx context.Context, in *AtualizarListaVeiculosReq, opts ...grpc.CallOption) (*AtualizarListaVeiculosResp, error) {
	out := new(AtualizarListaVeiculosResp)
	err := c.cc.Invoke(ctx, "/gravacao.Gravacao/AtualizarListaVeiculos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gravacaoClient) IniciarProcessamento(ctx context.Context, in *IniciarProcessamentoReq, opts ...grpc.CallOption) (*IniciarProcessamentoResp, error) {
	out := new(IniciarProcessamentoResp)
	err := c.cc.Invoke(ctx, "/gravacao.Gravacao/IniciarProcessamento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GravacaoServer is the server API for Gravacao service.
// All implementations must embed UnimplementedGravacaoServer
// for forward compatibility
type GravacaoServer interface {
	InfoProcessos(context.Context, *InfoProcessosReq) (*InfoProcessosResp, error)
	ConfigurarProcesso(context.Context, *ConfigurarProcessoReq) (*ConfigurarProcessoResp, error)
	AtualizarListaVeiculos(context.Context, *AtualizarListaVeiculosReq) (*AtualizarListaVeiculosResp, error)
	IniciarProcessamento(context.Context, *IniciarProcessamentoReq) (*IniciarProcessamentoResp, error)
	mustEmbedUnimplementedGravacaoServer()
}

// UnimplementedGravacaoServer must be embedded to have forward compatible implementations.
type UnimplementedGravacaoServer struct {
}

func (UnimplementedGravacaoServer) InfoProcessos(context.Context, *InfoProcessosReq) (*InfoProcessosResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoProcessos not implemented")
}
func (UnimplementedGravacaoServer) ConfigurarProcesso(context.Context, *ConfigurarProcessoReq) (*ConfigurarProcessoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigurarProcesso not implemented")
}
func (UnimplementedGravacaoServer) AtualizarListaVeiculos(context.Context, *AtualizarListaVeiculosReq) (*AtualizarListaVeiculosResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtualizarListaVeiculos not implemented")
}
func (UnimplementedGravacaoServer) IniciarProcessamento(context.Context, *IniciarProcessamentoReq) (*IniciarProcessamentoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IniciarProcessamento not implemented")
}
func (UnimplementedGravacaoServer) mustEmbedUnimplementedGravacaoServer() {}

// UnsafeGravacaoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GravacaoServer will
// result in compilation errors.
type UnsafeGravacaoServer interface {
	mustEmbedUnimplementedGravacaoServer()
}

func RegisterGravacaoServer(s grpc.ServiceRegistrar, srv GravacaoServer) {
	s.RegisterService(&Gravacao_ServiceDesc, srv)
}

func _Gravacao_InfoProcessos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoProcessosReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GravacaoServer).InfoProcessos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gravacao.Gravacao/InfoProcessos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GravacaoServer).InfoProcessos(ctx, req.(*InfoProcessosReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gravacao_ConfigurarProcesso_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigurarProcessoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GravacaoServer).ConfigurarProcesso(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gravacao.Gravacao/ConfigurarProcesso",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GravacaoServer).ConfigurarProcesso(ctx, req.(*ConfigurarProcessoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gravacao_AtualizarListaVeiculos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtualizarListaVeiculosReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GravacaoServer).AtualizarListaVeiculos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gravacao.Gravacao/AtualizarListaVeiculos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GravacaoServer).AtualizarListaVeiculos(ctx, req.(*AtualizarListaVeiculosReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gravacao_IniciarProcessamento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IniciarProcessamentoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GravacaoServer).IniciarProcessamento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gravacao.Gravacao/IniciarProcessamento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GravacaoServer).IniciarProcessamento(ctx, req.(*IniciarProcessamentoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Gravacao_ServiceDesc is the grpc.ServiceDesc for Gravacao service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gravacao_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gravacao.Gravacao",
	HandlerType: (*GravacaoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InfoProcessos",
			Handler:    _Gravacao_InfoProcessos_Handler,
		},
		{
			MethodName: "ConfigurarProcesso",
			Handler:    _Gravacao_ConfigurarProcesso_Handler,
		},
		{
			MethodName: "AtualizarListaVeiculos",
			Handler:    _Gravacao_AtualizarListaVeiculos_Handler,
		},
		{
			MethodName: "IniciarProcessamento",
			Handler:    _Gravacao_IniciarProcessamento_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gravacao.proto",
}
