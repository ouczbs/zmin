// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: WrapMessage.proto

package zpb

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

type WrapMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content  []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`    // 协议二进制
	Request  int32  `protobuf:"varint,2,opt,name=request,proto3" json:"request,omitempty"`   //请求码
	Response int32  `protobuf:"varint,3,opt,name=response,proto3" json:"response,omitempty"` //响应码
	Code     int32  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`         //返回码
}

func (x *WrapMessage) Reset() {
	*x = WrapMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WrapMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapMessage) ProtoMessage() {}

func (x *WrapMessage) ProtoReflect() protoreflect.Message {
	mi := &file_WrapMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapMessage.ProtoReflect.Descriptor instead.
func (*WrapMessage) Descriptor() ([]byte, []int) {
	return file_WrapMessage_proto_rawDescGZIP(), []int{0}
}

func (x *WrapMessage) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *WrapMessage) GetRequest() int32 {
	if x != nil {
		return x.Request
	}
	return 0
}

func (x *WrapMessage) GetResponse() int32 {
	if x != nil {
		return x.Response
	}
	return 0
}

func (x *WrapMessage) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_WrapMessage_proto protoreflect.FileDescriptor

var file_WrapMessage_proto_rawDesc = []byte{
	0x0a, 0x11, 0x57, 0x72, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x7a, 0x70, 0x62, 0x22, 0x71, 0x0a, 0x0b, 0x57, 0x72, 0x61, 0x70,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x3b, 0x7a, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WrapMessage_proto_rawDescOnce sync.Once
	file_WrapMessage_proto_rawDescData = file_WrapMessage_proto_rawDesc
)

func file_WrapMessage_proto_rawDescGZIP() []byte {
	file_WrapMessage_proto_rawDescOnce.Do(func() {
		file_WrapMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_WrapMessage_proto_rawDescData)
	})
	return file_WrapMessage_proto_rawDescData
}

var file_WrapMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WrapMessage_proto_goTypes = []interface{}{
	(*WrapMessage)(nil), // 0: zpb.WrapMessage
}
var file_WrapMessage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WrapMessage_proto_init() }
func file_WrapMessage_proto_init() {
	if File_WrapMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WrapMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapMessage); i {
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
			RawDescriptor: file_WrapMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WrapMessage_proto_goTypes,
		DependencyIndexes: file_WrapMessage_proto_depIdxs,
		MessageInfos:      file_WrapMessage_proto_msgTypes,
	}.Build()
	File_WrapMessage_proto = out.File
	file_WrapMessage_proto_rawDesc = nil
	file_WrapMessage_proto_goTypes = nil
	file_WrapMessage_proto_depIdxs = nil
}