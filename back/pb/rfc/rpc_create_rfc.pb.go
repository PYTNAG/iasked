// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: rfc/rpc_create_rfc.proto

package pb

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

type CreateRFCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int32  `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateRFCRequest) Reset() {
	*x = CreateRFCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfc_rpc_create_rfc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRFCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRFCRequest) ProtoMessage() {}

func (x *CreateRFCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rfc_rpc_create_rfc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRFCRequest.ProtoReflect.Descriptor instead.
func (*CreateRFCRequest) Descriptor() ([]byte, []int) {
	return file_rfc_rpc_create_rfc_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRFCRequest) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *CreateRFCRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateRFCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RfcId int32 `protobuf:"varint,1,opt,name=rfc_id,json=rfcId,proto3" json:"rfc_id,omitempty"`
}

func (x *CreateRFCResponse) Reset() {
	*x = CreateRFCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfc_rpc_create_rfc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRFCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRFCResponse) ProtoMessage() {}

func (x *CreateRFCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rfc_rpc_create_rfc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRFCResponse.ProtoReflect.Descriptor instead.
func (*CreateRFCResponse) Descriptor() ([]byte, []int) {
	return file_rfc_rpc_create_rfc_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRFCResponse) GetRfcId() int32 {
	if x != nil {
		return x.RfcId
	}
	return 0
}

var File_rfc_rpc_create_rfc_proto protoreflect.FileDescriptor

var file_rfc_rpc_create_rfc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x66, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x66, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x49,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x46, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x46, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x72, 0x66, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x72, 0x66, 0x63, 0x49, 0x64, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x59, 0x54, 0x4e, 0x41, 0x47, 0x2f, 0x69, 0x61, 0x73, 0x6b, 0x65,
	0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rfc_rpc_create_rfc_proto_rawDescOnce sync.Once
	file_rfc_rpc_create_rfc_proto_rawDescData = file_rfc_rpc_create_rfc_proto_rawDesc
)

func file_rfc_rpc_create_rfc_proto_rawDescGZIP() []byte {
	file_rfc_rpc_create_rfc_proto_rawDescOnce.Do(func() {
		file_rfc_rpc_create_rfc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rfc_rpc_create_rfc_proto_rawDescData)
	})
	return file_rfc_rpc_create_rfc_proto_rawDescData
}

var file_rfc_rpc_create_rfc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rfc_rpc_create_rfc_proto_goTypes = []interface{}{
	(*CreateRFCRequest)(nil),  // 0: pb.CreateRFCRequest
	(*CreateRFCResponse)(nil), // 1: pb.CreateRFCResponse
}
var file_rfc_rpc_create_rfc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rfc_rpc_create_rfc_proto_init() }
func file_rfc_rpc_create_rfc_proto_init() {
	if File_rfc_rpc_create_rfc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rfc_rpc_create_rfc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRFCRequest); i {
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
		file_rfc_rpc_create_rfc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRFCResponse); i {
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
			RawDescriptor: file_rfc_rpc_create_rfc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rfc_rpc_create_rfc_proto_goTypes,
		DependencyIndexes: file_rfc_rpc_create_rfc_proto_depIdxs,
		MessageInfos:      file_rfc_rpc_create_rfc_proto_msgTypes,
	}.Build()
	File_rfc_rpc_create_rfc_proto = out.File
	file_rfc_rpc_create_rfc_proto_rawDesc = nil
	file_rfc_rpc_create_rfc_proto_goTypes = nil
	file_rfc_rpc_create_rfc_proto_depIdxs = nil
}
