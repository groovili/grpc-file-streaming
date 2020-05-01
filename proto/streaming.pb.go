// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/streaming.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type File struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_5556cc946f1d51e2, []int{0}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type FileRequest struct {
	Filename             string   `protobuf:"bytes,1,opt,name=Filename,proto3" json:"Filename,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	Data                 *File    `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5556cc946f1d51e2, []int{1}
}

func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequest.Unmarshal(m, b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
}
func (m *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(m, src)
}
func (m *FileRequest) XXX_Size() int {
	return xxx_messageInfo_FileRequest.Size(m)
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

func (m *FileRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *FileRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileRequest) GetData() *File {
	if m != nil {
		return m.Data
	}
	return nil
}

type FileResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileResponse) Reset()         { *m = FileResponse{} }
func (m *FileResponse) String() string { return proto.CompactTextString(m) }
func (*FileResponse) ProtoMessage()    {}
func (*FileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5556cc946f1d51e2, []int{2}
}

func (m *FileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileResponse.Unmarshal(m, b)
}
func (m *FileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileResponse.Marshal(b, m, deterministic)
}
func (m *FileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileResponse.Merge(m, src)
}
func (m *FileResponse) XXX_Size() int {
	return xxx_messageInfo_FileResponse.Size(m)
}
func (m *FileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileResponse proto.InternalMessageInfo

func (m *FileResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *FileResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*File)(nil), "proto.File")
	proto.RegisterType((*FileRequest)(nil), "proto.FileRequest")
	proto.RegisterType((*FileResponse)(nil), "proto.FileResponse")
}

func init() {
	proto.RegisterFile("proto/streaming.proto", fileDescriptor_5556cc946f1d51e2)
}

var fileDescriptor_5556cc946f1d51e2 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xcd, 0xcc, 0x4b, 0xd7, 0x03, 0xf3, 0x85, 0x58,
	0xc1, 0x94, 0x92, 0x02, 0x17, 0x8b, 0x5b, 0x66, 0x4e, 0xaa, 0x90, 0x04, 0x17, 0xbb, 0x73, 0x7e,
	0x5e, 0x49, 0x6a, 0x5e, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x8c, 0xab, 0x14, 0xc7,
	0xc5, 0x0d, 0x52, 0x11, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc5, 0xc5, 0x01, 0xe2,
	0xe6, 0x25, 0xe6, 0xa6, 0x82, 0x55, 0x72, 0x06, 0xc1, 0xf9, 0x42, 0x42, 0x5c, 0x2c, 0xc1, 0x99,
	0x55, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x60, 0xb6, 0x90, 0x3c, 0x17, 0x8b, 0x4b,
	0x62, 0x49, 0xa2, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0x37, 0xc4, 0x76, 0x3d, 0xb0, 0x89,
	0x60, 0x09, 0x25, 0x1b, 0x2e, 0x1e, 0x88, 0xf9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0x60, 0x97, 0x04,
	0x97, 0x26, 0x27, 0xa7, 0x16, 0x17, 0x83, 0xcd, 0xe7, 0x08, 0x82, 0x71, 0xb1, 0x19, 0x6f, 0xe4,
	0xc5, 0xc5, 0x0b, 0xd2, 0x1d, 0x0c, 0xf3, 0x9d, 0x90, 0x25, 0x17, 0x57, 0x70, 0x6a, 0x5e, 0x0a,
	0x44, 0x40, 0x48, 0x08, 0xd9, 0x3e, 0x88, 0x0f, 0xa4, 0x84, 0x51, 0xc4, 0x20, 0xb6, 0x2a, 0x31,
	0x68, 0x30, 0x3a, 0xb1, 0x47, 0x41, 0x02, 0x25, 0x89, 0x0d, 0x4c, 0x19, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x82, 0xd3, 0x47, 0xf9, 0x3b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileStreamingClient is the client API for FileStreaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileStreamingClient interface {
	SendStream(ctx context.Context, opts ...grpc.CallOption) (FileStreaming_SendStreamClient, error)
}

type fileStreamingClient struct {
	cc grpc.ClientConnInterface
}

func NewFileStreamingClient(cc grpc.ClientConnInterface) FileStreamingClient {
	return &fileStreamingClient{cc}
}

func (c *fileStreamingClient) SendStream(ctx context.Context, opts ...grpc.CallOption) (FileStreaming_SendStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileStreaming_serviceDesc.Streams[0], "/proto.FileStreaming/SendStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileStreamingSendStreamClient{stream}
	return x, nil
}

type FileStreaming_SendStreamClient interface {
	Send(*FileRequest) error
	CloseAndRecv() (*FileResponse, error)
	grpc.ClientStream
}

type fileStreamingSendStreamClient struct {
	grpc.ClientStream
}

func (x *fileStreamingSendStreamClient) Send(m *FileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileStreamingSendStreamClient) CloseAndRecv() (*FileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileStreamingServer is the server API for FileStreaming service.
type FileStreamingServer interface {
	SendStream(FileStreaming_SendStreamServer) error
}

// UnimplementedFileStreamingServer can be embedded to have forward compatible implementations.
type UnimplementedFileStreamingServer struct {
}

func (*UnimplementedFileStreamingServer) SendStream(srv FileStreaming_SendStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SendStream not implemented")
}

func RegisterFileStreamingServer(s *grpc.Server, srv FileStreamingServer) {
	s.RegisterService(&_FileStreaming_serviceDesc, srv)
}

func _FileStreaming_SendStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileStreamingServer).SendStream(&fileStreamingSendStreamServer{stream})
}

type FileStreaming_SendStreamServer interface {
	SendAndClose(*FileResponse) error
	Recv() (*FileRequest, error)
	grpc.ServerStream
}

type fileStreamingSendStreamServer struct {
	grpc.ServerStream
}

func (x *fileStreamingSendStreamServer) SendAndClose(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileStreamingSendStreamServer) Recv() (*FileRequest, error) {
	m := new(FileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileStreaming_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FileStreaming",
	HandlerType: (*FileStreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendStream",
			Handler:       _FileStreaming_SendStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/streaming.proto",
}
