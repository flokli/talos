// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: resource/secrets/secrets.proto

package secrets

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CertAndKeyPEM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cert []byte `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	Key  []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CertAndKeyPEM) Reset() {
	*x = CertAndKeyPEM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_secrets_secrets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertAndKeyPEM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertAndKeyPEM) ProtoMessage() {}

func (x *CertAndKeyPEM) ProtoReflect() protoreflect.Message {
	mi := &file_resource_secrets_secrets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertAndKeyPEM.ProtoReflect.Descriptor instead.
func (*CertAndKeyPEM) Descriptor() ([]byte, []int) {
	return file_resource_secrets_secrets_proto_rawDescGZIP(), []int{0}
}

func (x *CertAndKeyPEM) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *CertAndKeyPEM) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

// APISpec describes secrets.API.
type APISpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaPem  []byte         `protobuf:"bytes,1,opt,name=ca_pem,json=caPem,proto3" json:"ca_pem,omitempty"`
	Server *CertAndKeyPEM `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Client *CertAndKeyPEM `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *APISpec) Reset() {
	*x = APISpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_secrets_secrets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APISpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APISpec) ProtoMessage() {}

func (x *APISpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_secrets_secrets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APISpec.ProtoReflect.Descriptor instead.
func (*APISpec) Descriptor() ([]byte, []int) {
	return file_resource_secrets_secrets_proto_rawDescGZIP(), []int{1}
}

func (x *APISpec) GetCaPem() []byte {
	if x != nil {
		return x.CaPem
	}
	return nil
}

func (x *APISpec) GetServer() *CertAndKeyPEM {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *APISpec) GetClient() *CertAndKeyPEM {
	if x != nil {
		return x.Client
	}
	return nil
}

var File_resource_secrets_secrets_proto protoreflect.FileDescriptor

var file_resource_secrets_secrets_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x22, 0x35, 0x0a, 0x0d, 0x43, 0x65, 0x72, 0x74, 0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79,
	0x50, 0x45, 0x4d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x92, 0x01, 0x0a, 0x07, 0x41, 0x50,
	0x49, 0x53, 0x70, 0x65, 0x63, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x5f, 0x70, 0x65, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x61, 0x50, 0x65, 0x6d, 0x12, 0x37, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e,
	0x43, 0x65, 0x72, 0x74, 0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x50, 0x45, 0x4d, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x41, 0x6e, 0x64,
	0x4b, 0x65, 0x79, 0x50, 0x45, 0x4d, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x43,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6c,
	0x6f, 0x73, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_secrets_secrets_proto_rawDescOnce sync.Once
	file_resource_secrets_secrets_proto_rawDescData = file_resource_secrets_secrets_proto_rawDesc
)

func file_resource_secrets_secrets_proto_rawDescGZIP() []byte {
	file_resource_secrets_secrets_proto_rawDescOnce.Do(func() {
		file_resource_secrets_secrets_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_secrets_secrets_proto_rawDescData)
	})
	return file_resource_secrets_secrets_proto_rawDescData
}

var (
	file_resource_secrets_secrets_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_resource_secrets_secrets_proto_goTypes  = []interface{}{
		(*CertAndKeyPEM)(nil), // 0: resource.secrets.CertAndKeyPEM
		(*APISpec)(nil),       // 1: resource.secrets.APISpec
	}
)

var file_resource_secrets_secrets_proto_depIdxs = []int32{
	0, // 0: resource.secrets.APISpec.server:type_name -> resource.secrets.CertAndKeyPEM
	0, // 1: resource.secrets.APISpec.client:type_name -> resource.secrets.CertAndKeyPEM
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_resource_secrets_secrets_proto_init() }
func file_resource_secrets_secrets_proto_init() {
	if File_resource_secrets_secrets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_secrets_secrets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertAndKeyPEM); i {
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
		file_resource_secrets_secrets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APISpec); i {
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
			RawDescriptor: file_resource_secrets_secrets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_secrets_secrets_proto_goTypes,
		DependencyIndexes: file_resource_secrets_secrets_proto_depIdxs,
		MessageInfos:      file_resource_secrets_secrets_proto_msgTypes,
	}.Build()
	File_resource_secrets_secrets_proto = out.File
	file_resource_secrets_secrets_proto_rawDesc = nil
	file_resource_secrets_secrets_proto_goTypes = nil
	file_resource_secrets_secrets_proto_depIdxs = nil
}
