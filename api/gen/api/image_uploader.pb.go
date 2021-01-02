// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: image_uploader.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ImageUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to File:
	//	*ImageUploadRequest_FileMeta_
	//	*ImageUploadRequest_Data
	File isImageUploadRequest_File `protobuf_oneof:"file"`
}

func (x *ImageUploadRequest) Reset() {
	*x = ImageUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_uploader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUploadRequest) ProtoMessage() {}

func (x *ImageUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_uploader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUploadRequest.ProtoReflect.Descriptor instead.
func (*ImageUploadRequest) Descriptor() ([]byte, []int) {
	return file_image_uploader_proto_rawDescGZIP(), []int{0}
}

func (m *ImageUploadRequest) GetFile() isImageUploadRequest_File {
	if m != nil {
		return m.File
	}
	return nil
}

func (x *ImageUploadRequest) GetFileMeta() *ImageUploadRequest_FileMeta {
	if x, ok := x.GetFile().(*ImageUploadRequest_FileMeta_); ok {
		return x.FileMeta
	}
	return nil
}

func (x *ImageUploadRequest) GetData() []byte {
	if x, ok := x.GetFile().(*ImageUploadRequest_Data); ok {
		return x.Data
	}
	return nil
}

type isImageUploadRequest_File interface {
	isImageUploadRequest_File()
}

type ImageUploadRequest_FileMeta_ struct {
	FileMeta *ImageUploadRequest_FileMeta `protobuf:"bytes,1,opt,name=file_meta,json=fileMeta,proto3,oneof"`
}

type ImageUploadRequest_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*ImageUploadRequest_FileMeta_) isImageUploadRequest_File() {}

func (*ImageUploadRequest_Data) isImageUploadRequest_File() {}

type ImageUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Size        int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Filename    string `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *ImageUploadResponse) Reset() {
	*x = ImageUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_uploader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUploadResponse) ProtoMessage() {}

func (x *ImageUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_uploader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUploadResponse.ProtoReflect.Descriptor instead.
func (*ImageUploadResponse) Descriptor() ([]byte, []int) {
	return file_image_uploader_proto_rawDescGZIP(), []int{1}
}

func (x *ImageUploadResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ImageUploadResponse) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ImageUploadResponse) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *ImageUploadResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type ImageUploadRequest_FileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *ImageUploadRequest_FileMeta) Reset() {
	*x = ImageUploadRequest_FileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_uploader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUploadRequest_FileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUploadRequest_FileMeta) ProtoMessage() {}

func (x *ImageUploadRequest_FileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_image_uploader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUploadRequest_FileMeta.ProtoReflect.Descriptor instead.
func (*ImageUploadRequest_FileMeta) Descriptor() ([]byte, []int) {
	return file_image_uploader_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ImageUploadRequest_FileMeta) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

var File_image_uploader_proto protoreflect.FileDescriptor

var file_image_uploader_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x22, 0xa6, 0x01, 0x0a, 0x12, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x26, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0x7c, 0x0a, 0x13, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x6b, 0x0a,
	0x12, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x22, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x67, 0x65,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_uploader_proto_rawDescOnce sync.Once
	file_image_uploader_proto_rawDescData = file_image_uploader_proto_rawDesc
)

func file_image_uploader_proto_rawDescGZIP() []byte {
	file_image_uploader_proto_rawDescOnce.Do(func() {
		file_image_uploader_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_uploader_proto_rawDescData)
	})
	return file_image_uploader_proto_rawDescData
}

var file_image_uploader_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_image_uploader_proto_goTypes = []interface{}{
	(*ImageUploadRequest)(nil),          // 0: image.uploader.ImageUploadRequest
	(*ImageUploadResponse)(nil),         // 1: image.uploader.ImageUploadResponse
	(*ImageUploadRequest_FileMeta)(nil), // 2: image.uploader.ImageUploadRequest.FileMeta
}
var file_image_uploader_proto_depIdxs = []int32{
	2, // 0: image.uploader.ImageUploadRequest.file_meta:type_name -> image.uploader.ImageUploadRequest.FileMeta
	0, // 1: image.uploader.ImageUploadService.Upload:input_type -> image.uploader.ImageUploadRequest
	1, // 2: image.uploader.ImageUploadService.Upload:output_type -> image.uploader.ImageUploadResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_image_uploader_proto_init() }
func file_image_uploader_proto_init() {
	if File_image_uploader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_uploader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUploadRequest); i {
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
		file_image_uploader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUploadResponse); i {
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
		file_image_uploader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUploadRequest_FileMeta); i {
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
	file_image_uploader_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ImageUploadRequest_FileMeta_)(nil),
		(*ImageUploadRequest_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_image_uploader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_uploader_proto_goTypes,
		DependencyIndexes: file_image_uploader_proto_depIdxs,
		MessageInfos:      file_image_uploader_proto_msgTypes,
	}.Build()
	File_image_uploader_proto = out.File
	file_image_uploader_proto_rawDesc = nil
	file_image_uploader_proto_goTypes = nil
	file_image_uploader_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImageUploadServiceClient is the client API for ImageUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageUploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (ImageUploadService_UploadClient, error)
}

type imageUploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageUploadServiceClient(cc grpc.ClientConnInterface) ImageUploadServiceClient {
	return &imageUploadServiceClient{cc}
}

func (c *imageUploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (ImageUploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ImageUploadService_serviceDesc.Streams[0], "/image.uploader.ImageUploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageUploadServiceUploadClient{stream}
	return x, nil
}

type ImageUploadService_UploadClient interface {
	Send(*ImageUploadRequest) error
	CloseAndRecv() (*ImageUploadResponse, error)
	grpc.ClientStream
}

type imageUploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *imageUploadServiceUploadClient) Send(m *ImageUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageUploadServiceUploadClient) CloseAndRecv() (*ImageUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ImageUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageUploadServiceServer is the server API for ImageUploadService service.
type ImageUploadServiceServer interface {
	Upload(ImageUploadService_UploadServer) error
}

// UnimplementedImageUploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedImageUploadServiceServer struct {
}

func (*UnimplementedImageUploadServiceServer) Upload(ImageUploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterImageUploadServiceServer(s *grpc.Server, srv ImageUploadServiceServer) {
	s.RegisterService(&_ImageUploadService_serviceDesc, srv)
}

func _ImageUploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageUploadServiceServer).Upload(&imageUploadServiceUploadServer{stream})
}

type ImageUploadService_UploadServer interface {
	SendAndClose(*ImageUploadResponse) error
	Recv() (*ImageUploadRequest, error)
	grpc.ServerStream
}

type imageUploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *imageUploadServiceUploadServer) SendAndClose(m *ImageUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageUploadServiceUploadServer) Recv() (*ImageUploadRequest, error) {
	m := new(ImageUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ImageUploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "image.uploader.ImageUploadService",
	HandlerType: (*ImageUploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _ImageUploadService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "image_uploader.proto",
}