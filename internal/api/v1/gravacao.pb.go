// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.14.0
// source: gravacao.proto

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

type InfoProcessosResp_Processo_StatusProcesso int32

const (
	InfoProcessosResp_Processo_PARADO     InfoProcessosResp_Processo_StatusProcesso = 0
	InfoProcessosResp_Processo_EXECUTANDO InfoProcessosResp_Processo_StatusProcesso = 1
)

// Enum value maps for InfoProcessosResp_Processo_StatusProcesso.
var (
	InfoProcessosResp_Processo_StatusProcesso_name = map[int32]string{
		0: "PARADO",
		1: "EXECUTANDO",
	}
	InfoProcessosResp_Processo_StatusProcesso_value = map[string]int32{
		"PARADO":     0,
		"EXECUTANDO": 1,
	}
)

func (x InfoProcessosResp_Processo_StatusProcesso) Enum() *InfoProcessosResp_Processo_StatusProcesso {
	p := new(InfoProcessosResp_Processo_StatusProcesso)
	*p = x
	return p
}

func (x InfoProcessosResp_Processo_StatusProcesso) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InfoProcessosResp_Processo_StatusProcesso) Descriptor() protoreflect.EnumDescriptor {
	return file_gravacao_proto_enumTypes[0].Descriptor()
}

func (InfoProcessosResp_Processo_StatusProcesso) Type() protoreflect.EnumType {
	return &file_gravacao_proto_enumTypes[0]
}

func (x InfoProcessosResp_Processo_StatusProcesso) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InfoProcessosResp_Processo_StatusProcesso.Descriptor instead.
func (InfoProcessosResp_Processo_StatusProcesso) EnumDescriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{3, 0, 0}
}

type ConfigurarProcessoReq_Acao int32

const (
	ConfigurarProcessoReq_CONFIGURAR ConfigurarProcessoReq_Acao = 0
	ConfigurarProcessoReq_INICIAR    ConfigurarProcessoReq_Acao = 1
	ConfigurarProcessoReq_PARAR      ConfigurarProcessoReq_Acao = 2
	ConfigurarProcessoReq_REMOVER    ConfigurarProcessoReq_Acao = 3
	ConfigurarProcessoReq_INFO       ConfigurarProcessoReq_Acao = 4
)

// Enum value maps for ConfigurarProcessoReq_Acao.
var (
	ConfigurarProcessoReq_Acao_name = map[int32]string{
		0: "CONFIGURAR",
		1: "INICIAR",
		2: "PARAR",
		3: "REMOVER",
		4: "INFO",
	}
	ConfigurarProcessoReq_Acao_value = map[string]int32{
		"CONFIGURAR": 0,
		"INICIAR":    1,
		"PARAR":      2,
		"REMOVER":    3,
		"INFO":       4,
	}
)

func (x ConfigurarProcessoReq_Acao) Enum() *ConfigurarProcessoReq_Acao {
	p := new(ConfigurarProcessoReq_Acao)
	*p = x
	return p
}

func (x ConfigurarProcessoReq_Acao) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigurarProcessoReq_Acao) Descriptor() protoreflect.EnumDescriptor {
	return file_gravacao_proto_enumTypes[1].Descriptor()
}

func (ConfigurarProcessoReq_Acao) Type() protoreflect.EnumType {
	return &file_gravacao_proto_enumTypes[1]
}

func (x ConfigurarProcessoReq_Acao) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigurarProcessoReq_Acao.Descriptor instead.
func (ConfigurarProcessoReq_Acao) EnumDescriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{4, 0}
}

type ConfigurarProcessoResp_Status int32

const (
	ConfigurarProcessoResp_INEXSISTENTE ConfigurarProcessoResp_Status = 0
	ConfigurarProcessoResp_EXECUTANDO   ConfigurarProcessoResp_Status = 1
	ConfigurarProcessoResp_PARADO       ConfigurarProcessoResp_Status = 2
)

// Enum value maps for ConfigurarProcessoResp_Status.
var (
	ConfigurarProcessoResp_Status_name = map[int32]string{
		0: "INEXSISTENTE",
		1: "EXECUTANDO",
		2: "PARADO",
	}
	ConfigurarProcessoResp_Status_value = map[string]int32{
		"INEXSISTENTE": 0,
		"EXECUTANDO":   1,
		"PARADO":       2,
	}
)

func (x ConfigurarProcessoResp_Status) Enum() *ConfigurarProcessoResp_Status {
	p := new(ConfigurarProcessoResp_Status)
	*p = x
	return p
}

func (x ConfigurarProcessoResp_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigurarProcessoResp_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_gravacao_proto_enumTypes[2].Descriptor()
}

func (ConfigurarProcessoResp_Status) Type() protoreflect.EnumType {
	return &file_gravacao_proto_enumTypes[2]
}

func (x ConfigurarProcessoResp_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigurarProcessoResp_Status.Descriptor instead.
func (ConfigurarProcessoResp_Status) EnumDescriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{5, 0}
}

type IniciarProcessamentoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IniciarProcessamentoReq) Reset() {
	*x = IniciarProcessamentoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gravacao_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IniciarProcessamentoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IniciarProcessamentoReq) ProtoMessage() {}

func (x *IniciarProcessamentoReq) ProtoReflect() protoreflect.Message {
	mi := &file_gravacao_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IniciarProcessamentoReq.ProtoReflect.Descriptor instead.
func (*IniciarProcessamentoReq) Descriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{0}
}

type IniciarProcessamentoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *IniciarProcessamentoResp) Reset() {
	*x = IniciarProcessamentoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gravacao_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IniciarProcessamentoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IniciarProcessamentoResp) ProtoMessage() {}

func (x *IniciarProcessamentoResp) ProtoReflect() protoreflect.Message {
	mi := &file_gravacao_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IniciarProcessamentoResp.ProtoReflect.Descriptor instead.
func (*IniciarProcessamentoResp) Descriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{1}
}

func (x *IniciarProcessamentoResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type InfoProcessosReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoProcessosReq) Reset() {
	*x = InfoProcessosReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gravacao_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoProcessosReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoProcessosReq) ProtoMessage() {}

func (x *InfoProcessosReq) ProtoReflect() protoreflect.Message {
	mi := &file_gravacao_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoProcessosReq.ProtoReflect.Descriptor instead.
func (*InfoProcessosReq) Descriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{2}
}

type InfoProcessosResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processos []*InfoProcessosResp_Processo `protobuf:"bytes,1,rep,name=processos,proto3" json:"processos,omitempty"`
}

func (x *InfoProcessosResp) Reset() {
	*x = InfoProcessosResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gravacao_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoProcessosResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoProcessosResp) ProtoMessage() {}

func (x *InfoProcessosResp) ProtoReflect() protoreflect.Message {
	mi := &file_gravacao_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoProcessosResp.ProtoReflect.Descriptor instead.
func (*InfoProcessosResp) Descriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{3}
}

func (x *InfoProcessosResp) GetProcessos() []*InfoProcessosResp_Processo {
	if x != nil {
		return x.Processos
	}
	return nil
}

type ConfigurarProcessoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acao               ConfigurarProcessoReq_Acao `protobuf:"varint,1,opt,name=acao,proto3,enum=gravacao.ConfigurarProcessoReq_Acao" json:"acao,omitempty"`
	Id                 int32                      `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	CameraId           int32                      `protobuf:"varint,3,opt,name=camera_id,json=cameraId,proto3" json:"camera_id,omitempty"`
	ProcessadorCaminho string                     `protobuf:"bytes,4,opt,name=processador_caminho,json=processadorCaminho,proto3" json:"processador_caminho,omitempty"`
}

func (x *ConfigurarProcessoReq) Reset() {
	*x = ConfigurarProcessoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gravacao_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurarProcessoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurarProcessoReq) ProtoMessage() {}

func (x *ConfigurarProcessoReq) ProtoReflect() protoreflect.Message {
	mi := &file_gravacao_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurarProcessoReq.ProtoReflect.Descriptor instead.
func (*ConfigurarProcessoReq) Descriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigurarProcessoReq) GetAcao() ConfigurarProcessoReq_Acao {
	if x != nil {
		return x.Acao
	}
	return ConfigurarProcessoReq_CONFIGURAR
}

func (x *ConfigurarProcessoReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConfigurarProcessoReq) GetCameraId() int32 {
	if x != nil {
		return x.CameraId
	}
	return 0
}

func (x *ConfigurarProcessoReq) GetProcessadorCaminho() string {
	if x != nil {
		return x.ProcessadorCaminho
	}
	return ""
}

type ConfigurarProcessoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status           ConfigurarProcessoResp_Status `protobuf:"varint,1,opt,name=status,proto3,enum=gravacao.ConfigurarProcessoResp_Status" json:"status,omitempty"`
	Armazenamento    string                        `protobuf:"bytes,2,opt,name=Armazenamento,proto3" json:"Armazenamento,omitempty"`
	ContadorCapturas int32                         `protobuf:"varint,3,opt,name=contador_capturas,json=contadorCapturas,proto3" json:"contador_capturas,omitempty"`
}

func (x *ConfigurarProcessoResp) Reset() {
	*x = ConfigurarProcessoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gravacao_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurarProcessoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurarProcessoResp) ProtoMessage() {}

func (x *ConfigurarProcessoResp) ProtoReflect() protoreflect.Message {
	mi := &file_gravacao_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurarProcessoResp.ProtoReflect.Descriptor instead.
func (*ConfigurarProcessoResp) Descriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigurarProcessoResp) GetStatus() ConfigurarProcessoResp_Status {
	if x != nil {
		return x.Status
	}
	return ConfigurarProcessoResp_INEXSISTENTE
}

func (x *ConfigurarProcessoResp) GetArmazenamento() string {
	if x != nil {
		return x.Armazenamento
	}
	return ""
}

func (x *ConfigurarProcessoResp) GetContadorCapturas() int32 {
	if x != nil {
		return x.ContadorCapturas
	}
	return 0
}

type AtualizarListaVeiculosReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Veiculo []*AtualizarListaVeiculosReq_Veiculo `protobuf:"bytes,1,rep,name=veiculo,proto3" json:"veiculo,omitempty"`
}

func (x *AtualizarListaVeiculosReq) Reset() {
	*x = AtualizarListaVeiculosReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gravacao_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtualizarListaVeiculosReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtualizarListaVeiculosReq) ProtoMessage() {}

func (x *AtualizarListaVeiculosReq) ProtoReflect() protoreflect.Message {
	mi := &file_gravacao_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtualizarListaVeiculosReq.ProtoReflect.Descriptor instead.
func (*AtualizarListaVeiculosReq) Descriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{6}
}

func (x *AtualizarListaVeiculosReq) GetVeiculo() []*AtualizarListaVeiculosReq_Veiculo {
	if x != nil {
		return x.Veiculo
	}
	return nil
}

type AtualizarListaVeiculosResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AtualizarListaVeiculosResp) Reset() {
	*x = AtualizarListaVeiculosResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gravacao_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtualizarListaVeiculosResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtualizarListaVeiculosResp) ProtoMessage() {}

func (x *AtualizarListaVeiculosResp) ProtoReflect() protoreflect.Message {
	mi := &file_gravacao_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtualizarListaVeiculosResp.ProtoReflect.Descriptor instead.
func (*AtualizarListaVeiculosResp) Descriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{7}
}

type InfoProcessosResp_Processo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int32                                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CameraId           int32                                     `protobuf:"varint,2,opt,name=camera_id,json=cameraId,proto3" json:"camera_id,omitempty"`
	ProcessadorCaminho string                                    `protobuf:"bytes,3,opt,name=processador_caminho,json=processadorCaminho,proto3" json:"processador_caminho,omitempty"`
	Status             InfoProcessosResp_Processo_StatusProcesso `protobuf:"varint,4,opt,name=status,proto3,enum=gravacao.InfoProcessosResp_Processo_StatusProcesso" json:"status,omitempty"`
}

func (x *InfoProcessosResp_Processo) Reset() {
	*x = InfoProcessosResp_Processo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gravacao_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoProcessosResp_Processo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoProcessosResp_Processo) ProtoMessage() {}

func (x *InfoProcessosResp_Processo) ProtoReflect() protoreflect.Message {
	mi := &file_gravacao_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoProcessosResp_Processo.ProtoReflect.Descriptor instead.
func (*InfoProcessosResp_Processo) Descriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{3, 0}
}

func (x *InfoProcessosResp_Processo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InfoProcessosResp_Processo) GetCameraId() int32 {
	if x != nil {
		return x.CameraId
	}
	return 0
}

func (x *InfoProcessosResp_Processo) GetProcessadorCaminho() string {
	if x != nil {
		return x.ProcessadorCaminho
	}
	return ""
}

func (x *InfoProcessosResp_Processo) GetStatus() InfoProcessosResp_Processo_StatusProcesso {
	if x != nil {
		return x.Status
	}
	return InfoProcessosResp_Processo_PARADO
}

type AtualizarListaVeiculosReq_Veiculo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Placa  string `protobuf:"bytes,1,opt,name=placa,proto3" json:"placa,omitempty"`
	Cor    string `protobuf:"bytes,2,opt,name=cor,proto3" json:"cor,omitempty"`
	Marca  string `protobuf:"bytes,3,opt,name=marca,proto3" json:"marca,omitempty"`
	Modelo string `protobuf:"bytes,4,opt,name=modelo,proto3" json:"modelo,omitempty"`
	Tipo   string `protobuf:"bytes,5,opt,name=tipo,proto3" json:"tipo,omitempty"`
}

func (x *AtualizarListaVeiculosReq_Veiculo) Reset() {
	*x = AtualizarListaVeiculosReq_Veiculo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gravacao_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtualizarListaVeiculosReq_Veiculo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtualizarListaVeiculosReq_Veiculo) ProtoMessage() {}

func (x *AtualizarListaVeiculosReq_Veiculo) ProtoReflect() protoreflect.Message {
	mi := &file_gravacao_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtualizarListaVeiculosReq_Veiculo.ProtoReflect.Descriptor instead.
func (*AtualizarListaVeiculosReq_Veiculo) Descriptor() ([]byte, []int) {
	return file_gravacao_proto_rawDescGZIP(), []int{6, 0}
}

func (x *AtualizarListaVeiculosReq_Veiculo) GetPlaca() string {
	if x != nil {
		return x.Placa
	}
	return ""
}

func (x *AtualizarListaVeiculosReq_Veiculo) GetCor() string {
	if x != nil {
		return x.Cor
	}
	return ""
}

func (x *AtualizarListaVeiculosReq_Veiculo) GetMarca() string {
	if x != nil {
		return x.Marca
	}
	return ""
}

func (x *AtualizarListaVeiculosReq_Veiculo) GetModelo() string {
	if x != nil {
		return x.Modelo
	}
	return ""
}

func (x *AtualizarListaVeiculosReq_Veiculo) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

var File_gravacao_proto protoreflect.FileDescriptor

var file_gravacao_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x49, 0x6e,
	0x69, 0x63, 0x69, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x32, 0x0a, 0x18, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x72,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x49, 0x6e, 0x66,
	0x6f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x22, 0xbd, 0x02,
	0x0a, 0x11, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61,
	0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x73, 0x1a, 0xe3, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49,
	0x64, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72,
	0x5f, 0x63, 0x61, 0x6d, 0x69, 0x6e, 0x68, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x43, 0x61, 0x6d, 0x69, 0x6e,
	0x68, 0x6f, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x2c, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x52, 0x41, 0x44, 0x4f, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x41, 0x4e, 0x44, 0x4f, 0x10, 0x01, 0x22, 0xf6, 0x01,
	0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x38, 0x0a, 0x04, 0x61, 0x63, 0x61, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x52, 0x65, 0x71, 0x2e, 0x41, 0x63, 0x61, 0x6f, 0x52, 0x04, 0x61, 0x63, 0x61,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x12, 0x2f,
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x5f, 0x63, 0x61,
	0x6d, 0x69, 0x6e, 0x68, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x43, 0x61, 0x6d, 0x69, 0x6e, 0x68, 0x6f, 0x22,
	0x45, 0x0a, 0x04, 0x41, 0x63, 0x61, 0x6f, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x55, 0x52, 0x41, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x49, 0x43, 0x49,
	0x41, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x52, 0x41, 0x52, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x04, 0x22, 0xe4, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x72, 0x6d, 0x61, 0x7a, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x72, 0x6d, 0x61, 0x7a,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x64, 0x6f, 0x72, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x61, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x64, 0x6f, 0x72, 0x43, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x61, 0x73, 0x22, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x45, 0x58, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54, 0x45, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x41, 0x4e, 0x44, 0x4f, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x52, 0x41, 0x44, 0x4f, 0x10, 0x02, 0x22, 0xd7, 0x01,
	0x0a, 0x19, 0x41, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x61,
	0x56, 0x65, 0x69, 0x63, 0x75, 0x6c, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x12, 0x45, 0x0a, 0x07, 0x76,
	0x65, 0x69, 0x63, 0x75, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67,
	0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x41, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x56, 0x65, 0x69, 0x63, 0x75, 0x6c, 0x6f, 0x73, 0x52, 0x65,
	0x71, 0x2e, 0x56, 0x65, 0x69, 0x63, 0x75, 0x6c, 0x6f, 0x52, 0x07, 0x76, 0x65, 0x69, 0x63, 0x75,
	0x6c, 0x6f, 0x1a, 0x73, 0x0a, 0x07, 0x56, 0x65, 0x69, 0x63, 0x75, 0x6c, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x6c, 0x61, 0x63, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c,
	0x61, 0x63, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x72, 0x63, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x61, 0x72, 0x63, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x22, 0x1c, 0x0a, 0x1a, 0x41, 0x74, 0x75, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x56, 0x65, 0x69, 0x63, 0x75, 0x6c, 0x6f,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x32, 0xf9, 0x02, 0x0a, 0x08, 0x47, 0x72, 0x61, 0x76, 0x61, 0x63,
	0x61, 0x6f, 0x12, 0x4a, 0x0a, 0x0d, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x73, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x1b, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x59,
	0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x12, 0x1f, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x16, 0x41, 0x74, 0x75,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x56, 0x65, 0x69, 0x63, 0x75,
	0x6c, 0x6f, 0x73, 0x12, 0x23, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x41,
	0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x56, 0x65, 0x69,
	0x63, 0x75, 0x6c, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61,
	0x63, 0x61, 0x6f, 0x2e, 0x41, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x61, 0x56, 0x65, 0x69, 0x63, 0x75, 0x6c, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x5f, 0x0a, 0x14, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x61,
	0x63, 0x61, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x67, 0x72,
	0x61, 0x76, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x72, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gravacao_proto_rawDescOnce sync.Once
	file_gravacao_proto_rawDescData = file_gravacao_proto_rawDesc
)

func file_gravacao_proto_rawDescGZIP() []byte {
	file_gravacao_proto_rawDescOnce.Do(func() {
		file_gravacao_proto_rawDescData = protoimpl.X.CompressGZIP(file_gravacao_proto_rawDescData)
	})
	return file_gravacao_proto_rawDescData
}

var file_gravacao_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_gravacao_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gravacao_proto_goTypes = []interface{}{
	(InfoProcessosResp_Processo_StatusProcesso)(0), // 0: gravacao.InfoProcessosResp.Processo.StatusProcesso
	(ConfigurarProcessoReq_Acao)(0),                // 1: gravacao.ConfigurarProcessoReq.Acao
	(ConfigurarProcessoResp_Status)(0),             // 2: gravacao.ConfigurarProcessoResp.Status
	(*IniciarProcessamentoReq)(nil),                // 3: gravacao.IniciarProcessamentoReq
	(*IniciarProcessamentoResp)(nil),               // 4: gravacao.IniciarProcessamentoResp
	(*InfoProcessosReq)(nil),                       // 5: gravacao.InfoProcessosReq
	(*InfoProcessosResp)(nil),                      // 6: gravacao.InfoProcessosResp
	(*ConfigurarProcessoReq)(nil),                  // 7: gravacao.ConfigurarProcessoReq
	(*ConfigurarProcessoResp)(nil),                 // 8: gravacao.ConfigurarProcessoResp
	(*AtualizarListaVeiculosReq)(nil),              // 9: gravacao.AtualizarListaVeiculosReq
	(*AtualizarListaVeiculosResp)(nil),             // 10: gravacao.AtualizarListaVeiculosResp
	(*InfoProcessosResp_Processo)(nil),             // 11: gravacao.InfoProcessosResp.Processo
	(*AtualizarListaVeiculosReq_Veiculo)(nil),      // 12: gravacao.AtualizarListaVeiculosReq.Veiculo
}
var file_gravacao_proto_depIdxs = []int32{
	11, // 0: gravacao.InfoProcessosResp.processos:type_name -> gravacao.InfoProcessosResp.Processo
	1,  // 1: gravacao.ConfigurarProcessoReq.acao:type_name -> gravacao.ConfigurarProcessoReq.Acao
	2,  // 2: gravacao.ConfigurarProcessoResp.status:type_name -> gravacao.ConfigurarProcessoResp.Status
	12, // 3: gravacao.AtualizarListaVeiculosReq.veiculo:type_name -> gravacao.AtualizarListaVeiculosReq.Veiculo
	0,  // 4: gravacao.InfoProcessosResp.Processo.status:type_name -> gravacao.InfoProcessosResp.Processo.StatusProcesso
	5,  // 5: gravacao.Gravacao.InfoProcessos:input_type -> gravacao.InfoProcessosReq
	7,  // 6: gravacao.Gravacao.ConfigurarProcesso:input_type -> gravacao.ConfigurarProcessoReq
	9,  // 7: gravacao.Gravacao.AtualizarListaVeiculos:input_type -> gravacao.AtualizarListaVeiculosReq
	3,  // 8: gravacao.Gravacao.IniciarProcessamento:input_type -> gravacao.IniciarProcessamentoReq
	6,  // 9: gravacao.Gravacao.InfoProcessos:output_type -> gravacao.InfoProcessosResp
	8,  // 10: gravacao.Gravacao.ConfigurarProcesso:output_type -> gravacao.ConfigurarProcessoResp
	10, // 11: gravacao.Gravacao.AtualizarListaVeiculos:output_type -> gravacao.AtualizarListaVeiculosResp
	4,  // 12: gravacao.Gravacao.IniciarProcessamento:output_type -> gravacao.IniciarProcessamentoResp
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_gravacao_proto_init() }
func file_gravacao_proto_init() {
	if File_gravacao_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gravacao_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IniciarProcessamentoReq); i {
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
		file_gravacao_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IniciarProcessamentoResp); i {
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
		file_gravacao_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoProcessosReq); i {
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
		file_gravacao_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoProcessosResp); i {
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
		file_gravacao_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurarProcessoReq); i {
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
		file_gravacao_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurarProcessoResp); i {
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
		file_gravacao_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtualizarListaVeiculosReq); i {
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
		file_gravacao_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtualizarListaVeiculosResp); i {
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
		file_gravacao_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoProcessosResp_Processo); i {
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
		file_gravacao_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtualizarListaVeiculosReq_Veiculo); i {
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
			RawDescriptor: file_gravacao_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gravacao_proto_goTypes,
		DependencyIndexes: file_gravacao_proto_depIdxs,
		EnumInfos:         file_gravacao_proto_enumTypes,
		MessageInfos:      file_gravacao_proto_msgTypes,
	}.Build()
	File_gravacao_proto = out.File
	file_gravacao_proto_rawDesc = nil
	file_gravacao_proto_goTypes = nil
	file_gravacao_proto_depIdxs = nil
}
