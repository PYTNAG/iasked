// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: comment/rpc_count_rfc_comments.proto

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

type CountRFCCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RfcId int32 `protobuf:"varint,1,opt,name=rfc_id,json=rfcId,proto3" json:"rfc_id,omitempty"`
}

func (x *CountRFCCommentsRequest) Reset() {
	*x = CountRFCCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_count_rfc_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountRFCCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountRFCCommentsRequest) ProtoMessage() {}

func (x *CountRFCCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_count_rfc_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountRFCCommentsRequest.ProtoReflect.Descriptor instead.
func (*CountRFCCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comment_rpc_count_rfc_comments_proto_rawDescGZIP(), []int{0}
}

func (x *CountRFCCommentsRequest) GetRfcId() int32 {
	if x != nil {
		return x.RfcId
	}
	return 0
}

type CountRFCCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentsCount int32 `protobuf:"varint,1,opt,name=comments_count,json=commentsCount,proto3" json:"comments_count,omitempty"`
}

func (x *CountRFCCommentsResponse) Reset() {
	*x = CountRFCCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_count_rfc_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountRFCCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountRFCCommentsResponse) ProtoMessage() {}

func (x *CountRFCCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_count_rfc_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountRFCCommentsResponse.ProtoReflect.Descriptor instead.
func (*CountRFCCommentsResponse) Descriptor() ([]byte, []int) {
	return file_comment_rpc_count_rfc_comments_proto_rawDescGZIP(), []int{1}
}

func (x *CountRFCCommentsResponse) GetCommentsCount() int32 {
	if x != nil {
		return x.CommentsCount
	}
	return 0
}

var File_comment_rpc_count_rfc_comments_proto protoreflect.FileDescriptor

var file_comment_rpc_count_rfc_comments_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x72, 0x66, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x30, 0x0a, 0x17, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x46, 0x43, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x66, 0x63, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x66, 0x63, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x18,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x46, 0x43, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x59,
	0x54, 0x4e, 0x41, 0x47, 0x2f, 0x69, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_rpc_count_rfc_comments_proto_rawDescOnce sync.Once
	file_comment_rpc_count_rfc_comments_proto_rawDescData = file_comment_rpc_count_rfc_comments_proto_rawDesc
)

func file_comment_rpc_count_rfc_comments_proto_rawDescGZIP() []byte {
	file_comment_rpc_count_rfc_comments_proto_rawDescOnce.Do(func() {
		file_comment_rpc_count_rfc_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_rpc_count_rfc_comments_proto_rawDescData)
	})
	return file_comment_rpc_count_rfc_comments_proto_rawDescData
}

var file_comment_rpc_count_rfc_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_comment_rpc_count_rfc_comments_proto_goTypes = []interface{}{
	(*CountRFCCommentsRequest)(nil),  // 0: pb.CountRFCCommentsRequest
	(*CountRFCCommentsResponse)(nil), // 1: pb.CountRFCCommentsResponse
}
var file_comment_rpc_count_rfc_comments_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comment_rpc_count_rfc_comments_proto_init() }
func file_comment_rpc_count_rfc_comments_proto_init() {
	if File_comment_rpc_count_rfc_comments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_rpc_count_rfc_comments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountRFCCommentsRequest); i {
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
		file_comment_rpc_count_rfc_comments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountRFCCommentsResponse); i {
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
			RawDescriptor: file_comment_rpc_count_rfc_comments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_comment_rpc_count_rfc_comments_proto_goTypes,
		DependencyIndexes: file_comment_rpc_count_rfc_comments_proto_depIdxs,
		MessageInfos:      file_comment_rpc_count_rfc_comments_proto_msgTypes,
	}.Build()
	File_comment_rpc_count_rfc_comments_proto = out.File
	file_comment_rpc_count_rfc_comments_proto_rawDesc = nil
	file_comment_rpc_count_rfc_comments_proto_goTypes = nil
	file_comment_rpc_count_rfc_comments_proto_depIdxs = nil
}
