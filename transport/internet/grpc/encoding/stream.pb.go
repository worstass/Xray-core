// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: transport/internet/grpc/encoding/stream.proto

package encoding

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Hunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Hunk) Reset() {
	*x = Hunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_grpc_encoding_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hunk) ProtoMessage() {}

func (x *Hunk) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_grpc_encoding_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hunk.ProtoReflect.Descriptor instead.
func (*Hunk) Descriptor() ([]byte, []int) {
	return file_transport_internet_grpc_encoding_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Hunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_transport_internet_grpc_encoding_stream_proto protoreflect.FileDescriptor

var file_transport_internet_grpc_encoding_stream_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x25, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x1a, 0x0a, 0x04, 0x48, 0x75, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x72, 0x0a, 0x0b, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x63, 0x0a, 0x03, 0x54, 0x75, 0x6e, 0x12, 0x2b, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x48, 0x75, 0x6e, 0x6b, 0x1a, 0x2b, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x75,
	0x6e, 0x6b, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c, 0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_internet_grpc_encoding_stream_proto_rawDescOnce sync.Once
	file_transport_internet_grpc_encoding_stream_proto_rawDescData = file_transport_internet_grpc_encoding_stream_proto_rawDesc
)

func file_transport_internet_grpc_encoding_stream_proto_rawDescGZIP() []byte {
	file_transport_internet_grpc_encoding_stream_proto_rawDescOnce.Do(func() {
		file_transport_internet_grpc_encoding_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_internet_grpc_encoding_stream_proto_rawDescData)
	})
	return file_transport_internet_grpc_encoding_stream_proto_rawDescData
}

var file_transport_internet_grpc_encoding_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transport_internet_grpc_encoding_stream_proto_goTypes = []interface{}{
	(*Hunk)(nil), // 0: xray.transport.internet.grpc.encoding.Hunk
}
var file_transport_internet_grpc_encoding_stream_proto_depIdxs = []int32{
	0, // 0: xray.transport.internet.grpc.encoding.GRPCService.Tun:input_type -> xray.transport.internet.grpc.encoding.Hunk
	0, // 1: xray.transport.internet.grpc.encoding.GRPCService.Tun:output_type -> xray.transport.internet.grpc.encoding.Hunk
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_internet_grpc_encoding_stream_proto_init() }
func file_transport_internet_grpc_encoding_stream_proto_init() {
	if File_transport_internet_grpc_encoding_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_internet_grpc_encoding_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hunk); i {
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
			RawDescriptor: file_transport_internet_grpc_encoding_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_internet_grpc_encoding_stream_proto_goTypes,
		DependencyIndexes: file_transport_internet_grpc_encoding_stream_proto_depIdxs,
		MessageInfos:      file_transport_internet_grpc_encoding_stream_proto_msgTypes,
	}.Build()
	File_transport_internet_grpc_encoding_stream_proto = out.File
	file_transport_internet_grpc_encoding_stream_proto_rawDesc = nil
	file_transport_internet_grpc_encoding_stream_proto_goTypes = nil
	file_transport_internet_grpc_encoding_stream_proto_depIdxs = nil
}
