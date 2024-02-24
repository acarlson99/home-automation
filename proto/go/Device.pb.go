// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v3.6.1
// source: Device.proto

package _go

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

type AuthType_ServiceType int32

const (
	AuthType_GOVEE AuthType_ServiceType = 0
)

// Enum value maps for AuthType_ServiceType.
var (
	AuthType_ServiceType_name = map[int32]string{
		0: "GOVEE",
	}
	AuthType_ServiceType_value = map[string]int32{
		"GOVEE": 0,
	}
)

func (x AuthType_ServiceType) Enum() *AuthType_ServiceType {
	p := new(AuthType_ServiceType)
	*p = x
	return p
}

func (x AuthType_ServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthType_ServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_Device_proto_enumTypes[0].Descriptor()
}

func (AuthType_ServiceType) Type() protoreflect.EnumType {
	return &file_Device_proto_enumTypes[0]
}

func (x AuthType_ServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthType_ServiceType.Descriptor instead.
func (AuthType_ServiceType) EnumDescriptor() ([]byte, []int) {
	return file_Device_proto_rawDescGZIP(), []int{4, 0}
}

type ElgatoLight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Port string `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ElgatoLight) Reset() {
	*x = ElgatoLight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElgatoLight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElgatoLight) ProtoMessage() {}

func (x *ElgatoLight) ProtoReflect() protoreflect.Message {
	mi := &file_Device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElgatoLight.ProtoReflect.Descriptor instead.
func (*ElgatoLight) Descriptor() ([]byte, []int) {
	return file_Device_proto_rawDescGZIP(), []int{0}
}

func (x *ElgatoLight) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ElgatoLight) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ElgatoLight) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type GoveeLight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MacAddress string `protobuf:"bytes,1,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Model      string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GoveeLight) Reset() {
	*x = GoveeLight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoveeLight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoveeLight) ProtoMessage() {}

func (x *GoveeLight) ProtoReflect() protoreflect.Message {
	mi := &file_Device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoveeLight.ProtoReflect.Descriptor instead.
func (*GoveeLight) Descriptor() ([]byte, []int) {
	return file_Device_proto_rawDescGZIP(), []int{1}
}

func (x *GoveeLight) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *GoveeLight) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *GoveeLight) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SmartDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Device:
	//	*SmartDevice_ElgatoLight
	//	*SmartDevice_GoveeLight
	Device isSmartDevice_Device `protobuf_oneof:"Device"`
}

func (x *SmartDevice) Reset() {
	*x = SmartDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartDevice) ProtoMessage() {}

func (x *SmartDevice) ProtoReflect() protoreflect.Message {
	mi := &file_Device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartDevice.ProtoReflect.Descriptor instead.
func (*SmartDevice) Descriptor() ([]byte, []int) {
	return file_Device_proto_rawDescGZIP(), []int{2}
}

func (m *SmartDevice) GetDevice() isSmartDevice_Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (x *SmartDevice) GetElgatoLight() *ElgatoLight {
	if x, ok := x.GetDevice().(*SmartDevice_ElgatoLight); ok {
		return x.ElgatoLight
	}
	return nil
}

func (x *SmartDevice) GetGoveeLight() *GoveeLight {
	if x, ok := x.GetDevice().(*SmartDevice_GoveeLight); ok {
		return x.GoveeLight
	}
	return nil
}

type isSmartDevice_Device interface {
	isSmartDevice_Device()
}

type SmartDevice_ElgatoLight struct {
	ElgatoLight *ElgatoLight `protobuf:"bytes,1,opt,name=elgato_light,json=elgatoLight,proto3,oneof"`
}

type SmartDevice_GoveeLight struct {
	GoveeLight *GoveeLight `protobuf:"bytes,2,opt,name=govee_light,json=goveeLight,proto3,oneof"`
}

func (*SmartDevice_ElgatoLight) isSmartDevice_Device() {}

func (*SmartDevice_GoveeLight) isSmartDevice_Device() {}

type Devices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device  []*SmartDevice `protobuf:"bytes,1,rep,name=device,proto3" json:"device,omitempty"`
	ApiAuth []*AuthType    `protobuf:"bytes,2,rep,name=api_auth,json=apiAuth,proto3" json:"api_auth,omitempty"`
}

func (x *Devices) Reset() {
	*x = Devices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Devices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Devices) ProtoMessage() {}

func (x *Devices) ProtoReflect() protoreflect.Message {
	mi := &file_Device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Devices.ProtoReflect.Descriptor instead.
func (*Devices) Descriptor() ([]byte, []int) {
	return file_Device_proto_rawDescGZIP(), []int{3}
}

func (x *Devices) GetDevice() []*SmartDevice {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *Devices) GetApiAuth() []*AuthType {
	if x != nil {
		return x.ApiAuth
	}
	return nil
}

type AuthType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Service     AuthType_ServiceType `protobuf:"varint,3,opt,name=service,proto3,enum=github.com.acarlson99.devices.AuthType_ServiceType" json:"service,omitempty"`
	// Types that are assignable to Auth:
	//	*AuthType_ApiKey
	//	*AuthType_Header
	//	*AuthType_Bearer_
	Auth isAuthType_Auth `protobuf_oneof:"Auth"`
}

func (x *AuthType) Reset() {
	*x = AuthType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthType) ProtoMessage() {}

func (x *AuthType) ProtoReflect() protoreflect.Message {
	mi := &file_Device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthType.ProtoReflect.Descriptor instead.
func (*AuthType) Descriptor() ([]byte, []int) {
	return file_Device_proto_rawDescGZIP(), []int{4}
}

func (x *AuthType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AuthType) GetService() AuthType_ServiceType {
	if x != nil {
		return x.Service
	}
	return AuthType_GOVEE
}

func (m *AuthType) GetAuth() isAuthType_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (x *AuthType) GetApiKey() *AuthType_APIKey {
	if x, ok := x.GetAuth().(*AuthType_ApiKey); ok {
		return x.ApiKey
	}
	return nil
}

func (x *AuthType) GetHeader() *AuthType_APIKey {
	if x, ok := x.GetAuth().(*AuthType_Header); ok {
		return x.Header
	}
	return nil
}

func (x *AuthType) GetBearer() *AuthType_Bearer {
	if x, ok := x.GetAuth().(*AuthType_Bearer_); ok {
		return x.Bearer
	}
	return nil
}

type isAuthType_Auth interface {
	isAuthType_Auth()
}

type AuthType_ApiKey struct {
	ApiKey *AuthType_APIKey `protobuf:"bytes,4,opt,name=apiKey,proto3,oneof"`
}

type AuthType_Header struct {
	Header *AuthType_APIKey `protobuf:"bytes,5,opt,name=header,proto3,oneof"`
}

type AuthType_Bearer_ struct {
	Bearer *AuthType_Bearer `protobuf:"bytes,6,opt,name=bearer,proto3,oneof"`
}

func (*AuthType_ApiKey) isAuthType_Auth() {}

func (*AuthType_Header) isAuthType_Auth() {}

func (*AuthType_Bearer_) isAuthType_Auth() {}

type AuthType_Bearer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthType_Bearer) Reset() {
	*x = AuthType_Bearer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthType_Bearer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthType_Bearer) ProtoMessage() {}

func (x *AuthType_Bearer) ProtoReflect() protoreflect.Message {
	mi := &file_Device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthType_Bearer.ProtoReflect.Descriptor instead.
func (*AuthType_Bearer) Descriptor() ([]byte, []int) {
	return file_Device_proto_rawDescGZIP(), []int{4, 0}
}

func (x *AuthType_Bearer) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthType_APIKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuthType_APIKey) Reset() {
	*x = AuthType_APIKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthType_APIKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthType_APIKey) ProtoMessage() {}

func (x *AuthType_APIKey) ProtoReflect() protoreflect.Message {
	mi := &file_Device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthType_APIKey.ProtoReflect.Descriptor instead.
func (*AuthType_APIKey) Descriptor() ([]byte, []int) {
	return file_Device_proto_rawDescGZIP(), []int{4, 1}
}

func (x *AuthType_APIKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AuthType_APIKey) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AuthType_APIKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_Device_proto protoreflect.FileDescriptor

var file_Device_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x61, 0x72, 0x6c,
	0x73, 0x6f, 0x6e, 0x39, 0x39, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x47, 0x0a,
	0x0b, 0x45, 0x6c, 0x67, 0x61, 0x74, 0x6f, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x57, 0x0a, 0x0a, 0x47, 0x6f, 0x76, 0x65, 0x65, 0x4c,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xb6, 0x01, 0x0a, 0x0b, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4f, 0x0a, 0x0c, 0x65, 0x6c, 0x67, 0x61, 0x74, 0x6f, 0x5f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x61, 0x72, 0x6c, 0x73, 0x6f, 0x6e, 0x39, 0x39, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x6c, 0x67, 0x61, 0x74, 0x6f, 0x4c, 0x69, 0x67, 0x68,
	0x74, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6c, 0x67, 0x61, 0x74, 0x6f, 0x4c, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x4c, 0x0a, 0x0b, 0x67, 0x6f, 0x76, 0x65, 0x65, 0x5f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x61, 0x72, 0x6c, 0x73, 0x6f, 0x6e, 0x39, 0x39, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x76, 0x65, 0x65, 0x4c, 0x69, 0x67, 0x68, 0x74,
	0x48, 0x00, 0x52, 0x0a, 0x67, 0x6f, 0x76, 0x65, 0x65, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x08,
	0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x07, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x63, 0x61, 0x72, 0x6c, 0x73, 0x6f, 0x6e, 0x39, 0x39, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x61, 0x72, 0x6c, 0x73, 0x6f, 0x6e,
	0x39, 0x39, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x07, 0x61, 0x70, 0x69, 0x41, 0x75, 0x74, 0x68, 0x22, 0xf5, 0x03, 0x0a,
	0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4d, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x33, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63,
	0x61, 0x72, 0x6c, 0x73, 0x6f, 0x6e, 0x39, 0x39, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48,
	0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x61, 0x72,
	0x6c, 0x73, 0x6f, 0x6e, 0x39, 0x39, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x48, 0x00,
	0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x48, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x61, 0x72, 0x6c, 0x73, 0x6f, 0x6e, 0x39, 0x39,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x48, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x63, 0x61, 0x72, 0x6c, 0x73, 0x6f, 0x6e, 0x39, 0x39, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x1a, 0x1e, 0x0a, 0x06,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x44, 0x0a, 0x06,
	0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x18, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x4f, 0x56, 0x45, 0x45, 0x10, 0x00, 0x42, 0x06, 0x0a, 0x04,
	0x41, 0x75, 0x74, 0x68, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Device_proto_rawDescOnce sync.Once
	file_Device_proto_rawDescData = file_Device_proto_rawDesc
)

func file_Device_proto_rawDescGZIP() []byte {
	file_Device_proto_rawDescOnce.Do(func() {
		file_Device_proto_rawDescData = protoimpl.X.CompressGZIP(file_Device_proto_rawDescData)
	})
	return file_Device_proto_rawDescData
}

var file_Device_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Device_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Device_proto_goTypes = []interface{}{
	(AuthType_ServiceType)(0), // 0: github.com.acarlson99.devices.AuthType.ServiceType
	(*ElgatoLight)(nil),       // 1: github.com.acarlson99.devices.ElgatoLight
	(*GoveeLight)(nil),        // 2: github.com.acarlson99.devices.GoveeLight
	(*SmartDevice)(nil),       // 3: github.com.acarlson99.devices.SmartDevice
	(*Devices)(nil),           // 4: github.com.acarlson99.devices.Devices
	(*AuthType)(nil),          // 5: github.com.acarlson99.devices.AuthType
	(*AuthType_Bearer)(nil),   // 6: github.com.acarlson99.devices.AuthType.Bearer
	(*AuthType_APIKey)(nil),   // 7: github.com.acarlson99.devices.AuthType.APIKey
}
var file_Device_proto_depIdxs = []int32{
	1, // 0: github.com.acarlson99.devices.SmartDevice.elgato_light:type_name -> github.com.acarlson99.devices.ElgatoLight
	2, // 1: github.com.acarlson99.devices.SmartDevice.govee_light:type_name -> github.com.acarlson99.devices.GoveeLight
	3, // 2: github.com.acarlson99.devices.Devices.device:type_name -> github.com.acarlson99.devices.SmartDevice
	5, // 3: github.com.acarlson99.devices.Devices.api_auth:type_name -> github.com.acarlson99.devices.AuthType
	0, // 4: github.com.acarlson99.devices.AuthType.service:type_name -> github.com.acarlson99.devices.AuthType.ServiceType
	7, // 5: github.com.acarlson99.devices.AuthType.apiKey:type_name -> github.com.acarlson99.devices.AuthType.APIKey
	7, // 6: github.com.acarlson99.devices.AuthType.header:type_name -> github.com.acarlson99.devices.AuthType.APIKey
	6, // 7: github.com.acarlson99.devices.AuthType.bearer:type_name -> github.com.acarlson99.devices.AuthType.Bearer
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_Device_proto_init() }
func file_Device_proto_init() {
	if File_Device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElgatoLight); i {
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
		file_Device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoveeLight); i {
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
		file_Device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartDevice); i {
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
		file_Device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Devices); i {
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
		file_Device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthType); i {
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
		file_Device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthType_Bearer); i {
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
		file_Device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthType_APIKey); i {
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
	file_Device_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SmartDevice_ElgatoLight)(nil),
		(*SmartDevice_GoveeLight)(nil),
	}
	file_Device_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*AuthType_ApiKey)(nil),
		(*AuthType_Header)(nil),
		(*AuthType_Bearer_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Device_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Device_proto_goTypes,
		DependencyIndexes: file_Device_proto_depIdxs,
		EnumInfos:         file_Device_proto_enumTypes,
		MessageInfos:      file_Device_proto_msgTypes,
	}.Build()
	File_Device_proto = out.File
	file_Device_proto_rawDesc = nil
	file_Device_proto_goTypes = nil
	file_Device_proto_depIdxs = nil
}