// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: app/extra/config.proto

package extra

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain        string       `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	SpineAddress  string       `protobuf:"bytes,2,opt,name=spineAddress,proto3" json:"spineAddress,omitempty"`
	Authenticator string       `protobuf:"bytes,3,opt,name=authenticator,proto3" json:"authenticator,omitempty"`
	AccessToken   string       `protobuf:"bytes,4,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	Panel         *PanelConfig `protobuf:"bytes,5,opt,name=panel,proto3" json:"panel,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_extra_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_extra_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_extra_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Config) GetSpineAddress() string {
	if x != nil {
		return x.SpineAddress
	}
	return ""
}

func (x *Config) GetAuthenticator() string {
	if x != nil {
		return x.Authenticator
	}
	return ""
}

func (x *Config) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Config) GetPanel() *PanelConfig {
	if x != nil {
		return x.Panel
	}
	return nil
}

type PanelConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseUrl   string `protobuf:"bytes,1,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
	NodeId    uint32 `protobuf:"varint,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	SecretKey string `protobuf:"bytes,3,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
}

func (x *PanelConfig) Reset() {
	*x = PanelConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_extra_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanelConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanelConfig) ProtoMessage() {}

func (x *PanelConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_extra_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanelConfig.ProtoReflect.Descriptor instead.
func (*PanelConfig) Descriptor() ([]byte, []int) {
	return file_app_extra_config_proto_rawDescGZIP(), []int{1}
}

func (x *PanelConfig) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *PanelConfig) GetNodeId() uint32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *PanelConfig) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

var File_app_extra_config_proto protoreflect.FileDescriptor

var file_app_extra_config_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x2f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0xbf, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73,
	0x70, 0x69, 0x6e, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x70, 0x69, 0x6e, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x31, 0x0a, 0x05, 0x70, 0x61, 0x6e, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x05, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x22, 0x5d, 0x0a, 0x0b, 0x50, 0x61,
	0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73,
	0x65, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x32, 0x0a, 0x12, 0x63, 0x6f, 0x6d,
	0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x50,
	0x01, 0x5a, 0x09, 0x61, 0x70, 0x70, 0x2f, 0x65, 0x78, 0x74, 0x72, 0x61, 0xaa, 0x02, 0x0e, 0x58,
	0x72, 0x61, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_extra_config_proto_rawDescOnce sync.Once
	file_app_extra_config_proto_rawDescData = file_app_extra_config_proto_rawDesc
)

func file_app_extra_config_proto_rawDescGZIP() []byte {
	file_app_extra_config_proto_rawDescOnce.Do(func() {
		file_app_extra_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_extra_config_proto_rawDescData)
	})
	return file_app_extra_config_proto_rawDescData
}

var file_app_extra_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_extra_config_proto_goTypes = []interface{}{
	(*Config)(nil),      // 0: xray.app.extra.Config
	(*PanelConfig)(nil), // 1: xray.app.extra.PanelConfig
}
var file_app_extra_config_proto_depIdxs = []int32{
	1, // 0: xray.app.extra.Config.panel:type_name -> xray.app.extra.PanelConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_extra_config_proto_init() }
func file_app_extra_config_proto_init() {
	if File_app_extra_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_extra_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_app_extra_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanelConfig); i {
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
			RawDescriptor: file_app_extra_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_extra_config_proto_goTypes,
		DependencyIndexes: file_app_extra_config_proto_depIdxs,
		MessageInfos:      file_app_extra_config_proto_msgTypes,
	}.Build()
	File_app_extra_config_proto = out.File
	file_app_extra_config_proto_rawDesc = nil
	file_app_extra_config_proto_goTypes = nil
	file_app_extra_config_proto_depIdxs = nil
}
