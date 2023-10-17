// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: internal/uam/service/role/role.proto

package role

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

type RoleCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Level int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RoleCreateRequest) Reset() {
	*x = RoleCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_role_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreateRequest) ProtoMessage() {}

func (x *RoleCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_role_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreateRequest.ProtoReflect.Descriptor instead.
func (*RoleCreateRequest) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_role_role_proto_rawDescGZIP(), []int{0}
}

func (x *RoleCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleCreateRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *RoleCreateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RoleUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level int32  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RoleUpdateRequest) Reset() {
	*x = RoleUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_role_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUpdateRequest) ProtoMessage() {}

func (x *RoleUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_role_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUpdateRequest.ProtoReflect.Descriptor instead.
func (*RoleUpdateRequest) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_role_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleUpdateRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *RoleUpdateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RoleListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RoleListRequest) Reset() {
	*x = RoleListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_role_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleListRequest) ProtoMessage() {}

func (x *RoleListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_role_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleListRequest.ProtoReflect.Descriptor instead.
func (*RoleListRequest) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_role_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RoleIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RoleIdRequest) Reset() {
	*x = RoleIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_role_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleIdRequest) ProtoMessage() {}

func (x *RoleIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_role_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleIdRequest.ProtoReflect.Descriptor instead.
func (*RoleIdRequest) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_role_role_proto_rawDescGZIP(), []int{3}
}

func (x *RoleIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleIdRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level int32  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *RoleResponse) Reset() {
	*x = RoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_role_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleResponse) ProtoMessage() {}

func (x *RoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_role_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleResponse.ProtoReflect.Descriptor instead.
func (*RoleResponse) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_role_role_proto_rawDescGZIP(), []int{4}
}

func (x *RoleResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleResponse) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type RoleListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*RoleResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *RoleListResponse) Reset() {
	*x = RoleListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_role_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleListResponse) ProtoMessage() {}

func (x *RoleListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_role_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleListResponse.ProtoReflect.Descriptor instead.
func (*RoleListResponse) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_role_role_proto_rawDescGZIP(), []int{5}
}

func (x *RoleListResponse) GetItems() []*RoleResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_internal_uam_service_role_role_proto protoreflect.FileDescriptor

var file_internal_uam_service_role_role_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x61, 0x6d, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x53, 0x0a, 0x11,
	0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x63, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x27, 0x0a, 0x0f, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x35, 0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x48, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x22, 0x3c, 0x0a, 0x10, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x74,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x15, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_uam_service_role_role_proto_rawDescOnce sync.Once
	file_internal_uam_service_role_role_proto_rawDescData = file_internal_uam_service_role_role_proto_rawDesc
)

func file_internal_uam_service_role_role_proto_rawDescGZIP() []byte {
	file_internal_uam_service_role_role_proto_rawDescOnce.Do(func() {
		file_internal_uam_service_role_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_uam_service_role_role_proto_rawDescData)
	})
	return file_internal_uam_service_role_role_proto_rawDescData
}

var file_internal_uam_service_role_role_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_uam_service_role_role_proto_goTypes = []interface{}{
	(*RoleCreateRequest)(nil), // 0: role.RoleCreateRequest
	(*RoleUpdateRequest)(nil), // 1: role.RoleUpdateRequest
	(*RoleListRequest)(nil),   // 2: role.RoleListRequest
	(*RoleIdRequest)(nil),     // 3: role.RoleIdRequest
	(*RoleResponse)(nil),      // 4: role.RoleResponse
	(*RoleListResponse)(nil),  // 5: role.RoleListResponse
}
var file_internal_uam_service_role_role_proto_depIdxs = []int32{
	4, // 0: role.RoleListResponse.items:type_name -> role.RoleResponse
	2, // 1: role.Role.GetList:input_type -> role.RoleListRequest
	3, // 2: role.Role.GetById:input_type -> role.RoleIdRequest
	5, // 3: role.Role.GetList:output_type -> role.RoleListResponse
	4, // 4: role.Role.GetById:output_type -> role.RoleResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_uam_service_role_role_proto_init() }
func file_internal_uam_service_role_role_proto_init() {
	if File_internal_uam_service_role_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_uam_service_role_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleCreateRequest); i {
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
		file_internal_uam_service_role_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleUpdateRequest); i {
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
		file_internal_uam_service_role_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleListRequest); i {
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
		file_internal_uam_service_role_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleIdRequest); i {
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
		file_internal_uam_service_role_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleResponse); i {
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
		file_internal_uam_service_role_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleListResponse); i {
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
			RawDescriptor: file_internal_uam_service_role_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_uam_service_role_role_proto_goTypes,
		DependencyIndexes: file_internal_uam_service_role_role_proto_depIdxs,
		MessageInfos:      file_internal_uam_service_role_role_proto_msgTypes,
	}.Build()
	File_internal_uam_service_role_role_proto = out.File
	file_internal_uam_service_role_role_proto_rawDesc = nil
	file_internal_uam_service_role_role_proto_goTypes = nil
	file_internal_uam_service_role_role_proto_depIdxs = nil
}
