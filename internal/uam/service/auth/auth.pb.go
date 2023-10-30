// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: internal/uam/service/auth/auth.proto

package auth

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

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Roles []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *AuthCredential) Reset() {
	*x = AuthCredential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCredential) ProtoMessage() {}

func (x *AuthCredential) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthCredential.ProtoReflect.Descriptor instead.
func (*AuthCredential) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthCredential) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthCredential) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string          `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Data *AuthCredential `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AuthResponse) GetData() *AuthCredential {
	if x != nil {
		return x.Data
	}
	return nil
}

type AuthVerifyTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthVerifyTokenRequest) Reset() {
	*x = AuthVerifyTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthVerifyTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthVerifyTokenRequest) ProtoMessage() {}

func (x *AuthVerifyTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthVerifyTokenRequest.ProtoReflect.Descriptor instead.
func (*AuthVerifyTokenRequest) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthVerifyTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthVerifyTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AuthVerifyTokenResponse) Reset() {
	*x = AuthVerifyTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_uam_service_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthVerifyTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthVerifyTokenResponse) ProtoMessage() {}

func (x *AuthVerifyTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_uam_service_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthVerifyTokenResponse.ProtoReflect.Descriptor instead.
func (*AuthVerifyTokenResponse) Descriptor() ([]byte, []int) {
	return file_internal_uam_service_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *AuthVerifyTokenResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AuthVerifyTokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_uam_service_auth_auth_proto protoreflect.FileDescriptor

var file_internal_uam_service_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x61, 0x6d, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x5b, 0x0a, 0x0b,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3c, 0x0a, 0x0e, 0x41, 0x75, 0x74,
	0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2e, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x47, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x82,
	0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_uam_service_auth_auth_proto_rawDescOnce sync.Once
	file_internal_uam_service_auth_auth_proto_rawDescData = file_internal_uam_service_auth_auth_proto_rawDesc
)

func file_internal_uam_service_auth_auth_proto_rawDescGZIP() []byte {
	file_internal_uam_service_auth_auth_proto_rawDescOnce.Do(func() {
		file_internal_uam_service_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_uam_service_auth_auth_proto_rawDescData)
	})
	return file_internal_uam_service_auth_auth_proto_rawDescData
}

var file_internal_uam_service_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_uam_service_auth_auth_proto_goTypes = []interface{}{
	(*AuthRequest)(nil),             // 0: auth.AuthRequest
	(*AuthCredential)(nil),          // 1: auth.AuthCredential
	(*AuthResponse)(nil),            // 2: auth.AuthResponse
	(*AuthVerifyTokenRequest)(nil),  // 3: auth.AuthVerifyTokenRequest
	(*AuthVerifyTokenResponse)(nil), // 4: auth.AuthVerifyTokenResponse
}
var file_internal_uam_service_auth_auth_proto_depIdxs = []int32{
	1, // 0: auth.AuthResponse.data:type_name -> auth.AuthCredential
	0, // 1: auth.Auth.Login:input_type -> auth.AuthRequest
	3, // 2: auth.Auth.VerifyToken:input_type -> auth.AuthVerifyTokenRequest
	2, // 3: auth.Auth.Login:output_type -> auth.AuthResponse
	4, // 4: auth.Auth.VerifyToken:output_type -> auth.AuthVerifyTokenResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_uam_service_auth_auth_proto_init() }
func file_internal_uam_service_auth_auth_proto_init() {
	if File_internal_uam_service_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_uam_service_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_internal_uam_service_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthCredential); i {
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
		file_internal_uam_service_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_internal_uam_service_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthVerifyTokenRequest); i {
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
		file_internal_uam_service_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthVerifyTokenResponse); i {
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
			RawDescriptor: file_internal_uam_service_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_uam_service_auth_auth_proto_goTypes,
		DependencyIndexes: file_internal_uam_service_auth_auth_proto_depIdxs,
		MessageInfos:      file_internal_uam_service_auth_auth_proto_msgTypes,
	}.Build()
	File_internal_uam_service_auth_auth_proto = out.File
	file_internal_uam_service_auth_auth_proto_rawDesc = nil
	file_internal_uam_service_auth_auth_proto_goTypes = nil
	file_internal_uam_service_auth_auth_proto_depIdxs = nil
}
