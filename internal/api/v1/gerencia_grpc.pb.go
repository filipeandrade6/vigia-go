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
	Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckRes, error)
	Match(ctx context.Context, in *MatchReq, opts ...grpc.CallOption) (*MatchRes, error)
	ErrorReport(ctx context.Context, in *ErrorReportReq, opts ...grpc.CallOption) (*ErrorReportRes, error)
}

type gerenciaClient struct {
	cc grpc.ClientConnInterface
}

func NewGerenciaClient(cc grpc.ClientConnInterface) GerenciaClient {
	return &gerenciaClient{cc}
}

func (c *gerenciaClient) Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckRes, error) {
	out := new(CheckRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) Match(ctx context.Context, in *MatchReq, opts ...grpc.CallOption) (*MatchRes, error) {
	out := new(MatchRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/Match", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerenciaClient) ErrorReport(ctx context.Context, in *ErrorReportReq, opts ...grpc.CallOption) (*ErrorReportRes, error) {
	out := new(ErrorReportRes)
	err := c.cc.Invoke(ctx, "/gerencia.Gerencia/ErrorReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GerenciaServer is the server API for Gerencia service.
// All implementations must embed UnimplementedGerenciaServer
// for forward compatibility
type GerenciaServer interface {
	Check(context.Context, *CheckReq) (*CheckRes, error)
	Match(context.Context, *MatchReq) (*MatchRes, error)
	ErrorReport(context.Context, *ErrorReportReq) (*ErrorReportRes, error)
	mustEmbedUnimplementedGerenciaServer()
}

// UnimplementedGerenciaServer must be embedded to have forward compatible implementations.
type UnimplementedGerenciaServer struct {
}

func (UnimplementedGerenciaServer) Check(context.Context, *CheckReq) (*CheckRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedGerenciaServer) Match(context.Context, *MatchReq) (*MatchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Match not implemented")
}
func (UnimplementedGerenciaServer) ErrorReport(context.Context, *ErrorReportReq) (*ErrorReportRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ErrorReport not implemented")
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

func _Gerencia_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).Check(ctx, req.(*CheckReq))
	}
	return interceptor(ctx, in, info, handler)
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

func _Gerencia_ErrorReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerenciaServer).ErrorReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerencia.Gerencia/ErrorReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerenciaServer).ErrorReport(ctx, req.(*ErrorReportReq))
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
			MethodName: "Check",
			Handler:    _Gerencia_Check_Handler,
		},
		{
			MethodName: "Match",
			Handler:    _Gerencia_Match_Handler,
		},
		{
			MethodName: "ErrorReport",
			Handler:    _Gerencia_ErrorReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gerencia.proto",
}
