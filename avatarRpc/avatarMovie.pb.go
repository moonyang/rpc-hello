// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: avatarMovie.proto

package avatar

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 请求体的结构体
type AvatarMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input []*Input `protobuf:"bytes,1,rep,name=input,proto3" json:"input,omitempty"`
}

func (x *AvatarMovieRequest) Reset() {
	*x = AvatarMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatarMovie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarMovieRequest) ProtoMessage() {}

func (x *AvatarMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_avatarMovie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarMovieRequest.ProtoReflect.Descriptor instead.
func (*AvatarMovieRequest) Descriptor() ([]byte, []int) {
	return file_avatarMovie_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarMovieRequest) GetInput() []*Input {
	if x != nil {
		return x.Input
	}
	return nil
}

// 响应的结构体
type AvatarMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output *Output `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *AvatarMovieResponse) Reset() {
	*x = AvatarMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatarMovie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarMovieResponse) ProtoMessage() {}

func (x *AvatarMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_avatarMovie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarMovieResponse.ProtoReflect.Descriptor instead.
func (*AvatarMovieResponse) Descriptor() ([]byte, []int) {
	return file_avatarMovie_proto_rawDescGZIP(), []int{1}
}

func (x *AvatarMovieResponse) GetOutput() *Output {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_avatarMovie_proto protoreflect.FileDescriptor

var file_avatarMovie_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x1a, 0x0b, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x12, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x22, 0x3d, 0x0a, 0x13, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x32, 0x5a, 0x0a, 0x0b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x4b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x12, 0x1a, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_avatarMovie_proto_rawDescOnce sync.Once
	file_avatarMovie_proto_rawDescData = file_avatarMovie_proto_rawDesc
)

func file_avatarMovie_proto_rawDescGZIP() []byte {
	file_avatarMovie_proto_rawDescOnce.Do(func() {
		file_avatarMovie_proto_rawDescData = protoimpl.X.CompressGZIP(file_avatarMovie_proto_rawDescData)
	})
	return file_avatarMovie_proto_rawDescData
}

var file_avatarMovie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_avatarMovie_proto_goTypes = []interface{}{
	(*AvatarMovieRequest)(nil),  // 0: avatar.AvatarMovieRequest
	(*AvatarMovieResponse)(nil), // 1: avatar.AvatarMovieResponse
	(*Input)(nil),               // 2: avatar.Input
	(*Output)(nil),              // 3: avatar.Output
}
var file_avatarMovie_proto_depIdxs = []int32{
	2, // 0: avatar.AvatarMovieRequest.input:type_name -> avatar.Input
	3, // 1: avatar.AvatarMovieResponse.output:type_name -> avatar.Output
	0, // 2: avatar.AvatarMovie.GetAvatarMovie:input_type -> avatar.AvatarMovieRequest
	1, // 3: avatar.AvatarMovie.GetAvatarMovie:output_type -> avatar.AvatarMovieResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_avatarMovie_proto_init() }
func file_avatarMovie_proto_init() {
	if File_avatarMovie_proto != nil {
		return
	}
	file_Input_proto_init()
	file_Output_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_avatarMovie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarMovieRequest); i {
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
		file_avatarMovie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarMovieResponse); i {
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
			RawDescriptor: file_avatarMovie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_avatarMovie_proto_goTypes,
		DependencyIndexes: file_avatarMovie_proto_depIdxs,
		MessageInfos:      file_avatarMovie_proto_msgTypes,
	}.Build()
	File_avatarMovie_proto = out.File
	file_avatarMovie_proto_rawDesc = nil
	file_avatarMovie_proto_goTypes = nil
	file_avatarMovie_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AvatarMovieClient is the client API for AvatarMovie service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AvatarMovieClient interface {
	GetAvatarMovie(ctx context.Context, in *AvatarMovieRequest, opts ...grpc.CallOption) (*AvatarMovieResponse, error)
}

type avatarMovieClient struct {
	cc grpc.ClientConnInterface
}

func NewAvatarMovieClient(cc grpc.ClientConnInterface) AvatarMovieClient {
	return &avatarMovieClient{cc}
}

func (c *avatarMovieClient) GetAvatarMovie(ctx context.Context, in *AvatarMovieRequest, opts ...grpc.CallOption) (*AvatarMovieResponse, error) {
	out := new(AvatarMovieResponse)
	err := c.cc.Invoke(ctx, "/avatar.AvatarMovie/GetAvatarMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvatarMovieServer is the server API for AvatarMovie service.
type AvatarMovieServer interface {
	GetAvatarMovie(context.Context, *AvatarMovieRequest) (*AvatarMovieResponse, error)
}

// UnimplementedAvatarMovieServer can be embedded to have forward compatible implementations.
type UnimplementedAvatarMovieServer struct {
}

func (*UnimplementedAvatarMovieServer) GetAvatarMovie(context.Context, *AvatarMovieRequest) (*AvatarMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvatarMovie not implemented")
}

func RegisterAvatarMovieServer(s *grpc.Server, srv AvatarMovieServer) {
	s.RegisterService(&_AvatarMovie_serviceDesc, srv)
}

func _AvatarMovie_GetAvatarMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvatarMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarMovieServer).GetAvatarMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/avatar.AvatarMovie/GetAvatarMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarMovieServer).GetAvatarMovie(ctx, req.(*AvatarMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AvatarMovie_serviceDesc = grpc.ServiceDesc{
	ServiceName: "avatar.AvatarMovie",
	HandlerType: (*AvatarMovieServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvatarMovie",
			Handler:    _AvatarMovie_GetAvatarMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "avatarMovie.proto",
}
