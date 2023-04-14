// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: LoginMessage.proto

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

// @message
type LoginAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Code     int32  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *LoginAccount) Reset() {
	*x = LoginAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LoginMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAccount) ProtoMessage() {}

func (x *LoginAccount) ProtoReflect() protoreflect.Message {
	mi := &file_LoginMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAccount.ProtoReflect.Descriptor instead.
func (*LoginAccount) Descriptor() ([]byte, []int) {
	return file_LoginMessage_proto_rawDescGZIP(), []int{0}
}

func (x *LoginAccount) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *LoginAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginAccount) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId  int32  `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	ServerId int32  `protobuf:"varint,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Player   string `protobuf:"bytes,4,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LoginMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_LoginMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_LoginMessage_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Role) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Role) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *Role) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

// @message
type LoginAccountAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Code     int32  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *LoginAccountAck) Reset() {
	*x = LoginAccountAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LoginMessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAccountAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAccountAck) ProtoMessage() {}

func (x *LoginAccountAck) ProtoReflect() protoreflect.Message {
	mi := &file_LoginMessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAccountAck.ProtoReflect.Descriptor instead.
func (*LoginAccountAck) Descriptor() ([]byte, []int) {
	return file_LoginMessage_proto_rawDescGZIP(), []int{2}
}

func (x *LoginAccountAck) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *LoginAccountAck) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginAccountAck) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginAccountAck) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

// @message
type RegisterAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Code     int32  `protobuf:"varint,6,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *RegisterAccount) Reset() {
	*x = RegisterAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LoginMessage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAccount) ProtoMessage() {}

func (x *RegisterAccount) ProtoReflect() protoreflect.Message {
	mi := &file_LoginMessage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAccount.ProtoReflect.Descriptor instead.
func (*RegisterAccount) Descriptor() ([]byte, []int) {
	return file_LoginMessage_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterAccount) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *RegisterAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterAccount) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegisterAccount) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterAccount) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_LoginMessage_proto protoreflect.FileDescriptor

var file_LoginMessage_proto_rawDesc = []byte{
	0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x7a, 0x70, 0x62, 0x22, 0x66, 0x0a, 0x0c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x66, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x69, 0x0a, 0x0f, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x3b, 0x7a, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LoginMessage_proto_rawDescOnce sync.Once
	file_LoginMessage_proto_rawDescData = file_LoginMessage_proto_rawDesc
)

func file_LoginMessage_proto_rawDescGZIP() []byte {
	file_LoginMessage_proto_rawDescOnce.Do(func() {
		file_LoginMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_LoginMessage_proto_rawDescData)
	})
	return file_LoginMessage_proto_rawDescData
}

var file_LoginMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_LoginMessage_proto_goTypes = []interface{}{
	(*LoginAccount)(nil),    // 0: zpb.LoginAccount
	(*Role)(nil),            // 1: zpb.Role
	(*LoginAccountAck)(nil), // 2: zpb.LoginAccountAck
	(*RegisterAccount)(nil), // 3: zpb.RegisterAccount
}
var file_LoginMessage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LoginMessage_proto_init() }
func file_LoginMessage_proto_init() {
	if File_LoginMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LoginMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAccount); i {
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
		file_LoginMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_LoginMessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAccountAck); i {
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
		file_LoginMessage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAccount); i {
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
			RawDescriptor: file_LoginMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LoginMessage_proto_goTypes,
		DependencyIndexes: file_LoginMessage_proto_depIdxs,
		MessageInfos:      file_LoginMessage_proto_msgTypes,
	}.Build()
	File_LoginMessage_proto = out.File
	file_LoginMessage_proto_rawDesc = nil
	file_LoginMessage_proto_goTypes = nil
	file_LoginMessage_proto_depIdxs = nil
}
