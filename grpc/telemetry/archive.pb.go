// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry/archive.proto

package telemetry

import (
	context "context"
	fmt "fmt"
	_ "github.com/airmap/interfaces/grpc"
	ids "github.com/airmap/interfaces/grpc/ids"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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

// OperationQuery models input parameters for querying historial operation telemetry.
type OperationQuery struct {
	Operation            *ids.Operation       `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Start                *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	SampleRate           *duration.Duration   `protobuf:"bytes,4,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OperationQuery) Reset()         { *m = OperationQuery{} }
func (m *OperationQuery) String() string { return proto.CompactTextString(m) }
func (*OperationQuery) ProtoMessage()    {}
func (*OperationQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b397bcda53d493bb, []int{0}
}

func (m *OperationQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationQuery.Unmarshal(m, b)
}
func (m *OperationQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationQuery.Marshal(b, m, deterministic)
}
func (m *OperationQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationQuery.Merge(m, src)
}
func (m *OperationQuery) XXX_Size() int {
	return xxx_messageInfo_OperationQuery.Size(m)
}
func (m *OperationQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationQuery.DiscardUnknown(m)
}

var xxx_messageInfo_OperationQuery proto.InternalMessageInfo

func (m *OperationQuery) GetOperation() *ids.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *OperationQuery) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *OperationQuery) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *OperationQuery) GetSampleRate() *duration.Duration {
	if m != nil {
		return m.SampleRate
	}
	return nil
}

func init() {
	proto.RegisterType((*OperationQuery)(nil), "telemetry.OperationQuery")
}

func init() { proto.RegisterFile("telemetry/archive.proto", fileDescriptor_b397bcda53d493bb) }

var fileDescriptor_b397bcda53d493bb = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x99, 0xb6, 0xdf, 0xa7, 0x4d, 0xb1, 0x68, 0x40, 0x9d, 0x0e, 0x54, 0xc5, 0x95, 0xa0,
	0x26, 0x52, 0xc1, 0xbd, 0x45, 0x74, 0x29, 0x0e, 0xae, 0xdc, 0x48, 0x3a, 0x73, 0x3b, 0x0d, 0x74,
	0x26, 0x21, 0xb9, 0x23, 0x76, 0xe7, 0xda, 0x07, 0xd1, 0xf7, 0xf3, 0x09, 0xa4, 0x49, 0x67, 0xc6,
	0x3f, 0x1b, 0x97, 0xb9, 0xe7, 0x77, 0x2e, 0xe7, 0x9e, 0x90, 0x5d, 0x84, 0x39, 0xe4, 0x80, 0x66,
	0xc1, 0x85, 0x49, 0x66, 0xf2, 0x09, 0x98, 0x36, 0x0a, 0x15, 0xed, 0xd6, 0x42, 0xb4, 0x9f, 0x29,
	0x95, 0xcd, 0x81, 0x3b, 0x61, 0x52, 0x4e, 0x39, 0xca, 0x1c, 0x2c, 0x8a, 0x5c, 0x7b, 0x36, 0xda,
	0xfb, 0x09, 0xa4, 0xa5, 0x11, 0x28, 0x55, 0xb1, 0xd2, 0x37, 0xe1, 0x19, 0xa1, 0xb0, 0x52, 0x15,
	0x76, 0x35, 0xd9, 0x90, 0xa9, 0xe5, 0x32, 0xad, 0x9e, 0x3b, 0x4d, 0x0a, 0x03, 0x5a, 0x19, 0xf4,
	0xf3, 0xc3, 0x8f, 0x80, 0xf4, 0x6f, 0x35, 0xf8, 0x65, 0x77, 0x25, 0x98, 0x05, 0x3d, 0x21, 0x5d,
	0x55, 0x4d, 0xc2, 0xe0, 0x20, 0x38, 0xea, 0x8d, 0xfa, 0x6c, 0xb9, 0xa9, 0xe6, 0xe2, 0x06, 0xa0,
	0x17, 0xe4, 0x9f, 0x45, 0x61, 0x30, 0x6c, 0x39, 0x32, 0x62, 0x3e, 0x29, 0xab, 0x92, 0xb2, 0xfb,
	0xea, 0x94, 0x71, 0xe7, 0xe5, 0x6d, 0x18, 0xc4, 0x1e, 0xa7, 0x23, 0xd2, 0x86, 0x22, 0x0d, 0xdb,
	0x7f, 0x74, 0x2d, 0x61, 0x7a, 0x4d, 0x7a, 0x56, 0xe4, 0x7a, 0x0e, 0x8f, 0x46, 0x20, 0x84, 0x1d,
	0xe7, 0x1d, 0xfc, 0xf2, 0x5e, 0xad, 0xba, 0x19, 0x93, 0xa5, 0xf5, 0xf5, 0x7d, 0xd8, 0x5a, 0x0f,
	0x62, 0xe2, 0x9d, 0xb1, 0x40, 0x18, 0xc5, 0x64, 0xed, 0xd2, 0x7f, 0x05, 0xbd, 0x21, 0xdb, 0xee,
	0xea, 0xe6, 0x36, 0xd7, 0x8e, 0xa5, 0x03, 0x56, 0x37, 0xc6, 0xbe, 0x17, 0x14, 0x6d, 0x7d, 0x91,
	0x3c, 0x7e, 0x16, 0x8c, 0x4f, 0x1f, 0x8e, 0x33, 0x89, 0xb3, 0x72, 0xc2, 0x12, 0x95, 0x73, 0x21,
	0x4d, 0x2e, 0x34, 0x97, 0x05, 0x82, 0x99, 0x8a, 0x04, 0x2c, 0xcf, 0x8c, 0x4e, 0x78, 0xed, 0x9b,
	0xfc, 0x77, 0x69, 0xcf, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x14, 0x1b, 0x12, 0x1e, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArchiveClient is the client API for Archive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArchiveClient interface {
	QueryOperationReports(ctx context.Context, in *OperationQuery, opts ...grpc.CallOption) (Archive_QueryOperationReportsClient, error)
}

type archiveClient struct {
	cc *grpc.ClientConn
}

func NewArchiveClient(cc *grpc.ClientConn) ArchiveClient {
	return &archiveClient{cc}
}

func (c *archiveClient) QueryOperationReports(ctx context.Context, in *OperationQuery, opts ...grpc.CallOption) (Archive_QueryOperationReportsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Archive_serviceDesc.Streams[0], "/telemetry.Archive/QueryOperationReports", opts...)
	if err != nil {
		return nil, err
	}
	x := &archiveQueryOperationReportsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Archive_QueryOperationReportsClient interface {
	Recv() (*Report, error)
	grpc.ClientStream
}

type archiveQueryOperationReportsClient struct {
	grpc.ClientStream
}

func (x *archiveQueryOperationReportsClient) Recv() (*Report, error) {
	m := new(Report)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArchiveServer is the server API for Archive service.
type ArchiveServer interface {
	QueryOperationReports(*OperationQuery, Archive_QueryOperationReportsServer) error
}

func RegisterArchiveServer(s *grpc.Server, srv ArchiveServer) {
	s.RegisterService(&_Archive_serviceDesc, srv)
}

func _Archive_QueryOperationReports_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OperationQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArchiveServer).QueryOperationReports(m, &archiveQueryOperationReportsServer{stream})
}

type Archive_QueryOperationReportsServer interface {
	Send(*Report) error
	grpc.ServerStream
}

type archiveQueryOperationReportsServer struct {
	grpc.ServerStream
}

func (x *archiveQueryOperationReportsServer) Send(m *Report) error {
	return x.ServerStream.SendMsg(m)
}

var _Archive_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telemetry.Archive",
	HandlerType: (*ArchiveServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryOperationReports",
			Handler:       _Archive_QueryOperationReports_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "telemetry/archive.proto",
}