// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry/collector.proto

package telemetry

import (
	context "context"
	fmt "fmt"
	_ "github.com/airmap/interfaces/src/go"
	system "github.com/airmap/interfaces/src/go/system"
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

// Update wraps types used in the exchange of telemetry updates with a collector.
type Update struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Update) Reset()         { *m = Update{} }
func (m *Update) String() string { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()    {}
func (*Update) Descriptor() ([]byte, []int) {
	return fileDescriptor_42f4f365628b62e0, []int{0}
}

func (m *Update) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update.Unmarshal(m, b)
}
func (m *Update) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update.Marshal(b, m, deterministic)
}
func (m *Update) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update.Merge(m, src)
}
func (m *Update) XXX_Size() int {
	return xxx_messageInfo_Update.Size(m)
}
func (m *Update) XXX_DiscardUnknown() {
	xxx_messageInfo_Update.DiscardUnknown(m)
}

var xxx_messageInfo_Update proto.InternalMessageInfo

// FromProvider wraps messages being sent from a provider to a collector.
type Update_FromProvider struct {
	// Types that are valid to be assigned to Details:
	//	*Update_FromProvider_Status
	//	*Update_FromProvider_Report
	Details              isUpdate_FromProvider_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Update_FromProvider) Reset()         { *m = Update_FromProvider{} }
func (m *Update_FromProvider) String() string { return proto.CompactTextString(m) }
func (*Update_FromProvider) ProtoMessage()    {}
func (*Update_FromProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_42f4f365628b62e0, []int{0, 0}
}

func (m *Update_FromProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update_FromProvider.Unmarshal(m, b)
}
func (m *Update_FromProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update_FromProvider.Marshal(b, m, deterministic)
}
func (m *Update_FromProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_FromProvider.Merge(m, src)
}
func (m *Update_FromProvider) XXX_Size() int {
	return xxx_messageInfo_Update_FromProvider.Size(m)
}
func (m *Update_FromProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_Update_FromProvider.DiscardUnknown(m)
}

var xxx_messageInfo_Update_FromProvider proto.InternalMessageInfo

type isUpdate_FromProvider_Details interface {
	isUpdate_FromProvider_Details()
}

type Update_FromProvider_Status struct {
	Status *system.Status `protobuf:"bytes,1,opt,name=status,proto3,oneof"`
}

type Update_FromProvider_Report struct {
	Report *Report `protobuf:"bytes,2,opt,name=report,proto3,oneof"`
}

func (*Update_FromProvider_Status) isUpdate_FromProvider_Details() {}

func (*Update_FromProvider_Report) isUpdate_FromProvider_Details() {}

func (m *Update_FromProvider) GetDetails() isUpdate_FromProvider_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Update_FromProvider) GetStatus() *system.Status {
	if x, ok := m.GetDetails().(*Update_FromProvider_Status); ok {
		return x.Status
	}
	return nil
}

func (m *Update_FromProvider) GetReport() *Report {
	if x, ok := m.GetDetails().(*Update_FromProvider_Report); ok {
		return x.Report
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Update_FromProvider) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Update_FromProvider_Status)(nil),
		(*Update_FromProvider_Report)(nil),
	}
}

// ToProvider wraps messages being sent from a collector back to a provider.
type Update_ToProvider struct {
	// Types that are valid to be assigned to Details:
	//	*Update_ToProvider_Status
	//	*Update_ToProvider_Ack
	Details              isUpdate_ToProvider_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Update_ToProvider) Reset()         { *m = Update_ToProvider{} }
func (m *Update_ToProvider) String() string { return proto.CompactTextString(m) }
func (*Update_ToProvider) ProtoMessage()    {}
func (*Update_ToProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_42f4f365628b62e0, []int{0, 1}
}

func (m *Update_ToProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update_ToProvider.Unmarshal(m, b)
}
func (m *Update_ToProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update_ToProvider.Marshal(b, m, deterministic)
}
func (m *Update_ToProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_ToProvider.Merge(m, src)
}
func (m *Update_ToProvider) XXX_Size() int {
	return xxx_messageInfo_Update_ToProvider.Size(m)
}
func (m *Update_ToProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_Update_ToProvider.DiscardUnknown(m)
}

var xxx_messageInfo_Update_ToProvider proto.InternalMessageInfo

type isUpdate_ToProvider_Details interface {
	isUpdate_ToProvider_Details()
}

type Update_ToProvider_Status struct {
	Status *system.Status `protobuf:"bytes,1,opt,name=status,proto3,oneof"`
}

type Update_ToProvider_Ack struct {
	Ack *system.Ack `protobuf:"bytes,2,opt,name=ack,proto3,oneof"`
}

func (*Update_ToProvider_Status) isUpdate_ToProvider_Details() {}

func (*Update_ToProvider_Ack) isUpdate_ToProvider_Details() {}

func (m *Update_ToProvider) GetDetails() isUpdate_ToProvider_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Update_ToProvider) GetStatus() *system.Status {
	if x, ok := m.GetDetails().(*Update_ToProvider_Status); ok {
		return x.Status
	}
	return nil
}

func (m *Update_ToProvider) GetAck() *system.Ack {
	if x, ok := m.GetDetails().(*Update_ToProvider_Ack); ok {
		return x.Ack
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Update_ToProvider) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Update_ToProvider_Status)(nil),
		(*Update_ToProvider_Ack)(nil),
	}
}

func init() {
	proto.RegisterType((*Update)(nil), "telemetry.Update")
	proto.RegisterType((*Update_FromProvider)(nil), "telemetry.Update.FromProvider")
	proto.RegisterType((*Update_ToProvider)(nil), "telemetry.Update.ToProvider")
}

func init() { proto.RegisterFile("telemetry/collector.proto", fileDescriptor_42f4f365628b62e0) }

var fileDescriptor_42f4f365628b62e0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x18, 0xc4, 0x1b, 0x85, 0x48, 0xb7, 0xa2, 0x35, 0x82, 0xd4, 0x20, 0x2a, 0x9e, 0x02, 0xe2, 0xae,
	0xd4, 0x27, 0xb0, 0x05, 0xe9, 0x51, 0x52, 0xbd, 0x78, 0x72, 0xb3, 0xf9, 0xac, 0x4b, 0xfe, 0x7c,
	0x61, 0xf7, 0xab, 0xd8, 0x97, 0xf5, 0x59, 0xa4, 0xd9, 0x24, 0x0d, 0x78, 0xf3, 0x3a, 0xbf, 0x61,
	0x66, 0x18, 0x76, 0x4e, 0x90, 0x43, 0x01, 0x64, 0x36, 0x42, 0x61, 0x9e, 0x83, 0x22, 0x34, 0xbc,
	0x32, 0x48, 0x18, 0x0c, 0x3b, 0x14, 0x8e, 0xe1, 0x9b, 0xa0, 0xb4, 0x1a, 0x4b, 0xeb, 0x60, 0x78,
	0x6a, 0x37, 0x96, 0xa0, 0x10, 0x96, 0x24, 0xad, 0x5b, 0x71, 0xdc, 0x88, 0x52, 0x65, 0x8d, 0x72,
	0xb6, 0x8b, 0x37, 0x50, 0xa1, 0x21, 0xa7, 0xdf, 0xfc, 0x78, 0xcc, 0x7f, 0xad, 0x52, 0x49, 0x10,
	0x56, 0xec, 0xf0, 0xc9, 0x60, 0xf1, 0x6c, 0xf0, 0x4b, 0xa7, 0x60, 0x82, 0x88, 0xf9, 0x2e, 0x74,
	0xe2, 0x5d, 0x7b, 0xd1, 0x68, 0x7a, 0xc4, 0x5d, 0x2a, 0x5f, 0xd6, 0xea, 0x62, 0x10, 0x37, 0x3c,
	0xb8, 0x65, 0xbe, 0x0b, 0x9d, 0xec, 0xd5, 0xce, 0x13, 0xde, 0xb5, 0xf1, 0xb8, 0x06, 0x5b, 0xb3,
	0xb3, 0xcc, 0x86, 0xec, 0x20, 0x05, 0x92, 0x3a, 0xb7, 0x61, 0xc2, 0xd8, 0x0b, 0xfe, 0xa3, 0xef,
	0x8a, 0xed, 0x4b, 0x95, 0x35, 0x65, 0xa3, 0xd6, 0xf6, 0xa8, 0xb2, 0xc5, 0x20, 0xde, 0x92, 0x5e,
	0xc7, 0xf4, 0x9d, 0x0d, 0xe7, 0xed, 0x9f, 0xc1, 0x92, 0x1d, 0xcf, 0xb1, 0x2c, 0x41, 0x51, 0xd7,
	0x7a, 0xd9, 0xdb, 0xea, 0x8e, 0xe0, 0xfd, 0x17, 0xc2, 0x8b, 0xbf, 0x7c, 0xb7, 0x39, 0xf2, 0xee,
	0xbd, 0x99, 0x78, 0xbb, 0x5b, 0x69, 0xfa, 0x5c, 0x27, 0x5c, 0x61, 0x21, 0xa4, 0x36, 0x85, 0xac,
	0x84, 0x2e, 0x09, 0xcc, 0x87, 0x54, 0x60, 0x85, 0x35, 0x4a, 0xac, 0x50, 0x74, 0x31, 0x89, 0x5f,
	0x5f, 0xff, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x8d, 0xac, 0x44, 0xf3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CollectorClient is the client API for Collector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectorClient interface {
	// ConnectProvider connects a stream of updates from a  provider to a collector
	ConnectProvider(ctx context.Context, opts ...grpc.CallOption) (Collector_ConnectProviderClient, error)
}

type collectorClient struct {
	cc *grpc.ClientConn
}

func NewCollectorClient(cc *grpc.ClientConn) CollectorClient {
	return &collectorClient{cc}
}

func (c *collectorClient) ConnectProvider(ctx context.Context, opts ...grpc.CallOption) (Collector_ConnectProviderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Collector_serviceDesc.Streams[0], "/telemetry.Collector/ConnectProvider", opts...)
	if err != nil {
		return nil, err
	}
	x := &collectorConnectProviderClient{stream}
	return x, nil
}

type Collector_ConnectProviderClient interface {
	Send(*Update_FromProvider) error
	Recv() (*Update_ToProvider, error)
	grpc.ClientStream
}

type collectorConnectProviderClient struct {
	grpc.ClientStream
}

func (x *collectorConnectProviderClient) Send(m *Update_FromProvider) error {
	return x.ClientStream.SendMsg(m)
}

func (x *collectorConnectProviderClient) Recv() (*Update_ToProvider, error) {
	m := new(Update_ToProvider)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CollectorServer is the server API for Collector service.
type CollectorServer interface {
	// ConnectProvider connects a stream of updates from a  provider to a collector
	ConnectProvider(Collector_ConnectProviderServer) error
}

// UnimplementedCollectorServer can be embedded to have forward compatible implementations.
type UnimplementedCollectorServer struct {
}

func (*UnimplementedCollectorServer) ConnectProvider(srv Collector_ConnectProviderServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectProvider not implemented")
}

func RegisterCollectorServer(s *grpc.Server, srv CollectorServer) {
	s.RegisterService(&_Collector_serviceDesc, srv)
}

func _Collector_ConnectProvider_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CollectorServer).ConnectProvider(&collectorConnectProviderServer{stream})
}

type Collector_ConnectProviderServer interface {
	Send(*Update_ToProvider) error
	Recv() (*Update_FromProvider, error)
	grpc.ServerStream
}

type collectorConnectProviderServer struct {
	grpc.ServerStream
}

func (x *collectorConnectProviderServer) Send(m *Update_ToProvider) error {
	return x.ServerStream.SendMsg(m)
}

func (x *collectorConnectProviderServer) Recv() (*Update_FromProvider, error) {
	m := new(Update_FromProvider)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Collector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telemetry.Collector",
	HandlerType: (*CollectorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectProvider",
			Handler:       _Collector_ConnectProvider_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "telemetry/collector.proto",
}
