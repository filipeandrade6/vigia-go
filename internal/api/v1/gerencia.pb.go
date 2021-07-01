// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: gerencia.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Camera struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID no banco de dados
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Camera) Reset() {
	*x = Camera{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gerencia_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Camera) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Camera) ProtoMessage() {}

func (x *Camera) ProtoReflect() protoreflect.Message {
	mi := &file_gerencia_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Camera.ProtoReflect.Descriptor instead.
func (*Camera) Descriptor() ([]byte, []int) {
	return file_gerencia_proto_rawDescGZIP(), []int{0}
}

func (x *Camera) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Camera a ser processada/modificada
	Camera      *Camera             `protobuf:"bytes,1,opt,name=camera,proto3" json:"camera,omitempty"` // Processador a ser acoplado na câmera ou que irá substituir
	Processador *Params_Processador `protobuf:"bytes,2,opt,name=processador,proto3" json:"processador,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gerencia_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_gerencia_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_gerencia_proto_rawDescGZIP(), []int{1}
}

func (x *Params) GetCamera() *Camera {
	if x != nil {
		return x.Camera
	}
	return nil
}

func (x *Params) GetProcessador() *Params_Processador {
	if x != nil {
		return x.Processador
	}
	return nil
}

type StatusArmazenamento struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Discos []*StatusArmazenamento_Disco `protobuf:"bytes,1,rep,name=discos,proto3" json:"discos,omitempty"`
}

func (x *StatusArmazenamento) Reset() {
	*x = StatusArmazenamento{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gerencia_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusArmazenamento) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusArmazenamento) ProtoMessage() {}

func (x *StatusArmazenamento) ProtoReflect() protoreflect.Message {
	mi := &file_gerencia_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusArmazenamento.ProtoReflect.Descriptor instead.
func (*StatusArmazenamento) Descriptor() ([]byte, []int) {
	return file_gerencia_proto_rawDescGZIP(), []int{2}
}

func (x *StatusArmazenamento) GetDiscos() []*StatusArmazenamento_Disco {
	if x != nil {
		return x.Discos
	}
	return nil
}

type ArmazenamentoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ArmazenamentoParams) Reset() {
	*x = ArmazenamentoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gerencia_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArmazenamentoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArmazenamentoParams) ProtoMessage() {}

func (x *ArmazenamentoParams) ProtoReflect() protoreflect.Message {
	mi := &file_gerencia_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArmazenamentoParams.ProtoReflect.Descriptor instead.
func (*ArmazenamentoParams) Descriptor() ([]byte, []int) {
	return file_gerencia_proto_rawDescGZIP(), []int{3}
}

type StatusIniciar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusIniciar) Reset() {
	*x = StatusIniciar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gerencia_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusIniciar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusIniciar) ProtoMessage() {}

func (x *StatusIniciar) ProtoReflect() protoreflect.Message {
	mi := &file_gerencia_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusIniciar.ProtoReflect.Descriptor instead.
func (*StatusIniciar) Descriptor() ([]byte, []int) {
	return file_gerencia_proto_rawDescGZIP(), []int{4}
}

type StatusAlterar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusAlterar) Reset() {
	*x = StatusAlterar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gerencia_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusAlterar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAlterar) ProtoMessage() {}

func (x *StatusAlterar) ProtoReflect() protoreflect.Message {
	mi := &file_gerencia_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAlterar.ProtoReflect.Descriptor instead.
func (*StatusAlterar) Descriptor() ([]byte, []int) {
	return file_gerencia_proto_rawDescGZIP(), []int{5}
}

type StatusParar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusParar) Reset() {
	*x = StatusParar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gerencia_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusParar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusParar) ProtoMessage() {}

func (x *StatusParar) ProtoReflect() protoreflect.Message {
	mi := &file_gerencia_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusParar.ProtoReflect.Descriptor instead.
func (*StatusParar) Descriptor() ([]byte, []int) {
	return file_gerencia_proto_rawDescGZIP(), []int{6}
}

type StatusRemover struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRemover) Reset() {
	*x = StatusRemover{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gerencia_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRemover) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRemover) ProtoMessage() {}

func (x *StatusRemover) ProtoReflect() protoreflect.Message {
	mi := &file_gerencia_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRemover.ProtoReflect.Descriptor instead.
func (*StatusRemover) Descriptor() ([]byte, []int) {
	return file_gerencia_proto_rawDescGZIP(), []int{7}
}

type Params_Processador struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Processador no baco de dados
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Params_Processador) Reset() {
	*x = Params_Processador{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gerencia_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params_Processador) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params_Processador) ProtoMessage() {}

func (x *Params_Processador) ProtoReflect() protoreflect.Message {
	mi := &file_gerencia_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params_Processador.ProtoReflect.Descriptor instead.
func (*Params_Processador) Descriptor() ([]byte, []int) {
	return file_gerencia_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Params_Processador) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StatusArmazenamento_Disco struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disponivel int32  `protobuf:"varint,1,opt,name=disponivel,proto3" json:"disponivel,omitempty"`
	Alocado    int32  `protobuf:"varint,2,opt,name=alocado,proto3" json:"alocado,omitempty"`
	Utilizado  int32  `protobuf:"varint,3,opt,name=utilizado,proto3" json:"utilizado,omitempty"`
	Caminho    string `protobuf:"bytes,4,opt,name=caminho,proto3" json:"caminho,omitempty"`
}

func (x *StatusArmazenamento_Disco) Reset() {
	*x = StatusArmazenamento_Disco{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gerencia_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusArmazenamento_Disco) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusArmazenamento_Disco) ProtoMessage() {}

func (x *StatusArmazenamento_Disco) ProtoReflect() protoreflect.Message {
	mi := &file_gerencia_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusArmazenamento_Disco.ProtoReflect.Descriptor instead.
func (*StatusArmazenamento_Disco) Descriptor() ([]byte, []int) {
	return file_gerencia_proto_rawDescGZIP(), []int{2, 0}
}

func (x *StatusArmazenamento_Disco) GetDisponivel() int32 {
	if x != nil {
		return x.Disponivel
	}
	return 0
}

func (x *StatusArmazenamento_Disco) GetAlocado() int32 {
	if x != nil {
		return x.Alocado
	}
	return 0
}

func (x *StatusArmazenamento_Disco) GetUtilizado() int32 {
	if x != nil {
		return x.Utilizado
	}
	return 0
}

func (x *StatusArmazenamento_Disco) GetCaminho() string {
	if x != nil {
		return x.Caminho
	}
	return ""
}

var File_gerencia_proto protoreflect.FileDescriptor

var file_gerencia_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x22, 0x18, 0x0a, 0x06, 0x43, 0x61,
	0x6d, 0x65, 0x72, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x28, 0x0a, 0x06, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x52, 0x06, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x12, 0x3e, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x1a, 0x1d, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xcd, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x41, 0x72, 0x6d, 0x61, 0x7a, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f,
	0x12, 0x3b, 0x0a, 0x06, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x41, 0x72, 0x6d, 0x61, 0x7a, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x52, 0x06, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x73, 0x1a, 0x79, 0x0a,
	0x05, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x6e,
	0x69, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x70,
	0x6f, 0x6e, 0x69, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x6f, 0x63, 0x61, 0x64,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x6c, 0x6f, 0x63, 0x61, 0x64, 0x6f,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x64, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x61, 0x6d, 0x69, 0x6e, 0x68, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x61, 0x6d, 0x69, 0x6e, 0x68, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x72, 0x6d, 0x61,
	0x7a, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x72,
	0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x61,
	0x72, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x72, 0x61, 0x72,
	0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x72, 0x32, 0xe5, 0x02, 0x0a, 0x08, 0x47, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x12, 0x3e,
	0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x12, 0x10, 0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x72, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x0f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x12, 0x10, 0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x72, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x0d, 0x70, 0x61, 0x72, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x12,
	0x10, 0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x50, 0x61, 0x72, 0x61, 0x72, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0f, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x12, 0x10, 0x2e,
	0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x1a,
	0x17, 0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x72, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x1b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x41, 0x72, 0x6d, 0x61, 0x7a, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x6f, 0x47, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x12, 0x1d, 0x2e, 0x67, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x69, 0x61, 0x2e, 0x41, 0x72, 0x6d, 0x61, 0x7a, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1d, 0x2e, 0x67, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x69, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x72, 0x6d, 0x61, 0x7a, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gerencia_proto_rawDescOnce sync.Once
	file_gerencia_proto_rawDescData = file_gerencia_proto_rawDesc
)

func file_gerencia_proto_rawDescGZIP() []byte {
	file_gerencia_proto_rawDescOnce.Do(func() {
		file_gerencia_proto_rawDescData = protoimpl.X.CompressGZIP(file_gerencia_proto_rawDescData)
	})
	return file_gerencia_proto_rawDescData
}

var file_gerencia_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gerencia_proto_goTypes = []interface{}{
	(*Camera)(nil),                    // 0: gerencia.Camera
	(*Params)(nil),                    // 1: gerencia.Params
	(*StatusArmazenamento)(nil),       // 2: gerencia.StatusArmazenamento
	(*ArmazenamentoParams)(nil),       // 3: gerencia.ArmazenamentoParams
	(*StatusIniciar)(nil),             // 4: gerencia.StatusIniciar
	(*StatusAlterar)(nil),             // 5: gerencia.StatusAlterar
	(*StatusParar)(nil),               // 6: gerencia.StatusParar
	(*StatusRemover)(nil),             // 7: gerencia.StatusRemover
	(*Params_Processador)(nil),        // 8: gerencia.Params.Processador
	(*StatusArmazenamento_Disco)(nil), // 9: gerencia.StatusArmazenamento.Disco
}
var file_gerencia_proto_depIdxs = []int32{
	0, // 0: gerencia.Params.camera:type_name -> gerencia.Camera
	8, // 1: gerencia.Params.processador:type_name -> gerencia.Params.Processador
	9, // 2: gerencia.StatusArmazenamento.discos:type_name -> gerencia.StatusArmazenamento.Disco
	1, // 3: gerencia.Gerencia.iniciarProcesso:input_type -> gerencia.Params
	1, // 4: gerencia.Gerencia.alterarProcesso:input_type -> gerencia.Params
	0, // 5: gerencia.Gerencia.pararProcesso:input_type -> gerencia.Camera
	0, // 6: gerencia.Gerencia.removerProcesso:input_type -> gerencia.Camera
	3, // 7: gerencia.Gerencia.statusArmazenamentoGravacao:input_type -> gerencia.ArmazenamentoParams
	4, // 8: gerencia.Gerencia.iniciarProcesso:output_type -> gerencia.StatusIniciar
	5, // 9: gerencia.Gerencia.alterarProcesso:output_type -> gerencia.StatusAlterar
	6, // 10: gerencia.Gerencia.pararProcesso:output_type -> gerencia.StatusParar
	7, // 11: gerencia.Gerencia.removerProcesso:output_type -> gerencia.StatusRemover
	2, // 12: gerencia.Gerencia.statusArmazenamentoGravacao:output_type -> gerencia.StatusArmazenamento
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gerencia_proto_init() }
func file_gerencia_proto_init() {
	if File_gerencia_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gerencia_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Camera); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gerencia_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gerencia_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusArmazenamento); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gerencia_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArmazenamentoParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gerencia_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusIniciar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gerencia_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusAlterar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gerencia_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusParar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gerencia_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRemover); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gerencia_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params_Processador); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gerencia_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusArmazenamento_Disco); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gerencia_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gerencia_proto_goTypes,
		DependencyIndexes: file_gerencia_proto_depIdxs,
		MessageInfos:      file_gerencia_proto_msgTypes,
	}.Build()
	File_gerencia_proto = out.File
	file_gerencia_proto_rawDesc = nil
	file_gerencia_proto_goTypes = nil
	file_gerencia_proto_depIdxs = nil
}
