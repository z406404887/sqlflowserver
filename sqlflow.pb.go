// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sqlflow.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// We take https://dev.mysql.com/doc/refman/8.0/en/data-types.html
// as the reference.  Because the real data is in text format as
// defined in message Column, we simplify type indicators here: INT
// represents all types of integer values in MySQL 8.0: TINYINT,
// INT, etc; FLOAT represent both fixed-point and floating-point and
// both 4-byte and 8-byte real values in MySQL 8.0.  DATE represents
// DATE, DATETIME, and TIMESTAMP.  STRING represents CHAR, VARCHAR,
// BINARY, VARBINARY, BLOB, TEXT.  More enum for other types like
// JSON and spatial in the future.
type FieldType int32

const (
	FieldType_INT    FieldType = 0
	FieldType_FLOAT  FieldType = 1
	FieldType_DATE   FieldType = 2
	FieldType_STRING FieldType = 3
)

var FieldType_name = map[int32]string{
	0: "INT",
	1: "FLOAT",
	2: "DATE",
	3: "STRING",
}

var FieldType_value = map[string]int32{
	"INT":    0,
	"FLOAT":  1,
	"DATE":   2,
	"STRING": 3,
}

func (x FieldType) String() string {
	return proto.EnumName(FieldType_name, int32(x))
}

func (FieldType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e9f9b081a5698ac, []int{0}
}

type RunRequest struct {
	Sql                  string   `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}
func (*RunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9f9b081a5698ac, []int{0}
}

func (m *RunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest.Unmarshal(m, b)
}
func (m *RunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest.Marshal(b, m, deterministic)
}
func (m *RunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest.Merge(m, src)
}
func (m *RunRequest) XXX_Size() int {
	return xxx_messageInfo_RunRequest.Size(m)
}
func (m *RunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest proto.InternalMessageInfo

func (m *RunRequest) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

type RunResponse struct {
	Message              []string             `protobuf:"bytes,1,rep,name=message,proto3" json:"message,omitempty"`
	Schema               map[string]FieldType `protobuf:"bytes,2,rep,name=schema,proto3" json:"schema,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=main.FieldType"`
	Columns              map[string]*Column   `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RunResponse) Reset()         { *m = RunResponse{} }
func (m *RunResponse) String() string { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()    {}
func (*RunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9f9b081a5698ac, []int{1}
}

func (m *RunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunResponse.Unmarshal(m, b)
}
func (m *RunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunResponse.Marshal(b, m, deterministic)
}
func (m *RunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunResponse.Merge(m, src)
}
func (m *RunResponse) XXX_Size() int {
	return xxx_messageInfo_RunResponse.Size(m)
}
func (m *RunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunResponse proto.InternalMessageInfo

func (m *RunResponse) GetMessage() []string {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *RunResponse) GetSchema() map[string]FieldType {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *RunResponse) GetColumns() map[string]*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

// Due to https://github.com/protocolbuffers/protobuf/issues/2388,
// we cannot write map<string, repeated string>; instead, we need to
// define repeated string as Column.
type Column struct {
	Data                 []string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9f9b081a5698ac, []int{2}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.FieldType", FieldType_name, FieldType_value)
	proto.RegisterType((*RunRequest)(nil), "main.RunRequest")
	proto.RegisterType((*RunResponse)(nil), "main.RunResponse")
	proto.RegisterMapType((map[string]*Column)(nil), "main.RunResponse.ColumnsEntry")
	proto.RegisterMapType((map[string]FieldType)(nil), "main.RunResponse.SchemaEntry")
	proto.RegisterType((*Column)(nil), "main.Column")
}

func init() { proto.RegisterFile("sqlflow.proto", fileDescriptor_6e9f9b081a5698ac) }

var fileDescriptor_6e9f9b081a5698ac = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x61, 0x6b, 0xe2, 0x40,
	0x10, 0x86, 0x6f, 0xb3, 0x9a, 0x5c, 0x26, 0xde, 0x5d, 0x6e, 0x3e, 0x05, 0x69, 0x25, 0x04, 0x0a,
	0xa1, 0x94, 0x50, 0x52, 0xa4, 0xd2, 0x6f, 0xd2, 0x6a, 0x6b, 0x11, 0x4b, 0xd7, 0xfc, 0x81, 0x54,
	0xb7, 0xad, 0xb8, 0x49, 0xd4, 0x4d, 0x2a, 0xf9, 0x53, 0xfd, 0x8d, 0xc5, 0xac, 0x96, 0x80, 0xfd,
	0x36, 0xbb, 0xf3, 0xbe, 0xcf, 0x0c, 0xef, 0xc0, 0x1f, 0xb9, 0x16, 0xaf, 0x22, 0xdb, 0x06, 0xab,
	0x4d, 0x96, 0x67, 0xd8, 0x48, 0xe2, 0x45, 0xea, 0x75, 0x00, 0x58, 0x91, 0x32, 0xbe, 0x2e, 0xb8,
	0xcc, 0xd1, 0x06, 0x2a, 0xd7, 0xc2, 0x21, 0x2e, 0xf1, 0x4d, 0xb6, 0x2b, 0xbd, 0x4f, 0x0d, 0xac,
	0x4a, 0x20, 0x57, 0x59, 0x2a, 0x39, 0x3a, 0x60, 0x24, 0x5c, 0xca, 0xf8, 0x8d, 0x3b, 0xc4, 0xa5,
	0xbe, 0xc9, 0x0e, 0x4f, 0xec, 0x82, 0x2e, 0x67, 0xef, 0x3c, 0x89, 0x1d, 0xcd, 0xa5, 0xbe, 0x15,
	0x9e, 0x06, 0xbb, 0x01, 0x41, 0xcd, 0x1c, 0x4c, 0xab, 0xfe, 0x20, 0xcd, 0x37, 0x25, 0xdb, 0x8b,
	0xb1, 0x07, 0xc6, 0x2c, 0x13, 0x45, 0x92, 0x4a, 0x87, 0x56, 0xbe, 0xce, 0xb1, 0xef, 0x56, 0x09,
	0x94, 0xf1, 0x20, 0x6f, 0x3f, 0x82, 0x55, 0x03, 0xee, 0x76, 0x5f, 0xf2, 0xf2, 0xb0, 0xfb, 0x92,
	0x97, 0x78, 0x06, 0xcd, 0x8f, 0x58, 0x14, 0xdc, 0xd1, 0x5c, 0xe2, 0xff, 0x0d, 0xff, 0x29, 0xf0,
	0x70, 0xc1, 0xc5, 0x3c, 0x2a, 0x57, 0x9c, 0xa9, 0xee, 0x8d, 0xd6, 0x23, 0xed, 0x07, 0x68, 0xd5,
	0x87, 0xfc, 0x00, 0xf3, 0xea, 0x30, 0x2b, 0x6c, 0x29, 0x98, 0x32, 0xd5, 0x48, 0xde, 0x09, 0xe8,
	0xea, 0x13, 0x11, 0x1a, 0xf3, 0x38, 0x8f, 0xf7, 0x39, 0x55, 0xf5, 0x79, 0x17, 0xcc, 0xef, 0xf9,
	0x68, 0x00, 0x1d, 0x4d, 0x22, 0xfb, 0x17, 0x9a, 0xd0, 0x1c, 0x8e, 0x9f, 0xfa, 0x91, 0x4d, 0xf0,
	0x37, 0x34, 0xee, 0xfa, 0xd1, 0xc0, 0xd6, 0x10, 0x40, 0x9f, 0x46, 0x6c, 0x34, 0xb9, 0xb7, 0x69,
	0x78, 0x0d, 0xc6, 0xf4, 0x79, 0x3c, 0x14, 0xd9, 0x16, 0x2f, 0x80, 0xb2, 0x22, 0x45, 0xbb, 0x96,
	0x52, 0x75, 0xbb, 0xf6, 0xff, 0xa3, 0xdc, 0x2e, 0xc9, 0x8b, 0x5e, 0xdd, 0xfa, 0xea, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x43, 0xcf, 0xe8, 0xd7, 0xfc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SQLFlowClient is the client API for SQLFlow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SQLFlowClient interface {
	// Run runs a SQL statement and returns a stream of RunResponse.  Some
	// SQL statements like `USE database` returns only a success message.
	// In such cases, the return stream contains only one RunResponse with
	// RunResponse.message set.  Some other SQL statements like `SELECT ...`
	// returns a table in addition to the status, and the table might be
	// big.  In such cases, Run returns a stream of RunResponse -- the
	// first one has RunResponse.schema set, and all RunReplies have
	// RunResponse.error and RunResponse.columns set -- each RunResponse doesn't
	// exceed the maximum message size
	// (https://stackoverflow.com/a/34186672/724872).
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (SQLFlow_RunClient, error)
}

type sQLFlowClient struct {
	cc *grpc.ClientConn
}

func NewSQLFlowClient(cc *grpc.ClientConn) SQLFlowClient {
	return &sQLFlowClient{cc}
}

func (c *sQLFlowClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (SQLFlow_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SQLFlow_serviceDesc.Streams[0], "/main.SQLFlow/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &sQLFlowRunClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SQLFlow_RunClient interface {
	Recv() (*RunResponse, error)
	grpc.ClientStream
}

type sQLFlowRunClient struct {
	grpc.ClientStream
}

func (x *sQLFlowRunClient) Recv() (*RunResponse, error) {
	m := new(RunResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SQLFlowServer is the server API for SQLFlow service.
type SQLFlowServer interface {
	// Run runs a SQL statement and returns a stream of RunResponse.  Some
	// SQL statements like `USE database` returns only a success message.
	// In such cases, the return stream contains only one RunResponse with
	// RunResponse.message set.  Some other SQL statements like `SELECT ...`
	// returns a table in addition to the status, and the table might be
	// big.  In such cases, Run returns a stream of RunResponse -- the
	// first one has RunResponse.schema set, and all RunReplies have
	// RunResponse.error and RunResponse.columns set -- each RunResponse doesn't
	// exceed the maximum message size
	// (https://stackoverflow.com/a/34186672/724872).
	Run(*RunRequest, SQLFlow_RunServer) error
}

func RegisterSQLFlowServer(s *grpc.Server, srv SQLFlowServer) {
	s.RegisterService(&_SQLFlow_serviceDesc, srv)
}

func _SQLFlow_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SQLFlowServer).Run(m, &sQLFlowRunServer{stream})
}

type SQLFlow_RunServer interface {
	Send(*RunResponse) error
	grpc.ServerStream
}

type sQLFlowRunServer struct {
	grpc.ServerStream
}

func (x *sQLFlowRunServer) Send(m *RunResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SQLFlow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.SQLFlow",
	HandlerType: (*SQLFlowServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _SQLFlow_Run_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sqlflow.proto",
}