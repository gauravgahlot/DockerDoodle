// Code generated by protoc-gen-go. DO NOT EDIT.
// source: docker-host.proto

package pb

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

type DockerHost struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ColorCode            string   `protobuf:"bytes,3,opt,name=colorCode,proto3" json:"colorCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerHost) Reset()         { *m = DockerHost{} }
func (m *DockerHost) String() string { return proto.CompactTextString(m) }
func (*DockerHost) ProtoMessage()    {}
func (*DockerHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_1358bea4b2532a29, []int{0}
}

func (m *DockerHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerHost.Unmarshal(m, b)
}
func (m *DockerHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerHost.Marshal(b, m, deterministic)
}
func (m *DockerHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerHost.Merge(m, src)
}
func (m *DockerHost) XXX_Size() int {
	return xxx_messageInfo_DockerHost.Size(m)
}
func (m *DockerHost) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerHost.DiscardUnknown(m)
}

var xxx_messageInfo_DockerHost proto.InternalMessageInfo

func (m *DockerHost) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DockerHost) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DockerHost) GetColorCode() string {
	if m != nil {
		return m.ColorCode
	}
	return ""
}

type HostContainerCount struct {
	Containers           map[string]int32 `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *HostContainerCount) Reset()         { *m = HostContainerCount{} }
func (m *HostContainerCount) String() string { return proto.CompactTextString(m) }
func (*HostContainerCount) ProtoMessage()    {}
func (*HostContainerCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_1358bea4b2532a29, []int{1}
}

func (m *HostContainerCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostContainerCount.Unmarshal(m, b)
}
func (m *HostContainerCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostContainerCount.Marshal(b, m, deterministic)
}
func (m *HostContainerCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostContainerCount.Merge(m, src)
}
func (m *HostContainerCount) XXX_Size() int {
	return xxx_messageInfo_HostContainerCount.Size(m)
}
func (m *HostContainerCount) XXX_DiscardUnknown() {
	xxx_messageInfo_HostContainerCount.DiscardUnknown(m)
}

var xxx_messageInfo_HostContainerCount proto.InternalMessageInfo

func (m *HostContainerCount) GetContainers() map[string]int32 {
	if m != nil {
		return m.Containers
	}
	return nil
}

type GetContainersCountRequest struct {
	Hosts                []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	All                  bool     `protobuf:"varint,2,opt,name=all,proto3" json:"all,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContainersCountRequest) Reset()         { *m = GetContainersCountRequest{} }
func (m *GetContainersCountRequest) String() string { return proto.CompactTextString(m) }
func (*GetContainersCountRequest) ProtoMessage()    {}
func (*GetContainersCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1358bea4b2532a29, []int{2}
}

func (m *GetContainersCountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContainersCountRequest.Unmarshal(m, b)
}
func (m *GetContainersCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContainersCountRequest.Marshal(b, m, deterministic)
}
func (m *GetContainersCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContainersCountRequest.Merge(m, src)
}
func (m *GetContainersCountRequest) XXX_Size() int {
	return xxx_messageInfo_GetContainersCountRequest.Size(m)
}
func (m *GetContainersCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContainersCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetContainersCountRequest proto.InternalMessageInfo

func (m *GetContainersCountRequest) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *GetContainersCountRequest) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

type GetContainersCountResponse struct {
	HostContainers       []*HostContainerCount `protobuf:"bytes,1,rep,name=hostContainers,proto3" json:"hostContainers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetContainersCountResponse) Reset()         { *m = GetContainersCountResponse{} }
func (m *GetContainersCountResponse) String() string { return proto.CompactTextString(m) }
func (*GetContainersCountResponse) ProtoMessage()    {}
func (*GetContainersCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1358bea4b2532a29, []int{3}
}

func (m *GetContainersCountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContainersCountResponse.Unmarshal(m, b)
}
func (m *GetContainersCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContainersCountResponse.Marshal(b, m, deterministic)
}
func (m *GetContainersCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContainersCountResponse.Merge(m, src)
}
func (m *GetContainersCountResponse) XXX_Size() int {
	return xxx_messageInfo_GetContainersCountResponse.Size(m)
}
func (m *GetContainersCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContainersCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetContainersCountResponse proto.InternalMessageInfo

func (m *GetContainersCountResponse) GetHostContainers() []*HostContainerCount {
	if m != nil {
		return m.HostContainers
	}
	return nil
}

type GetContainersRequest struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContainersRequest) Reset()         { *m = GetContainersRequest{} }
func (m *GetContainersRequest) String() string { return proto.CompactTextString(m) }
func (*GetContainersRequest) ProtoMessage()    {}
func (*GetContainersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1358bea4b2532a29, []int{4}
}

func (m *GetContainersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContainersRequest.Unmarshal(m, b)
}
func (m *GetContainersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContainersRequest.Marshal(b, m, deterministic)
}
func (m *GetContainersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContainersRequest.Merge(m, src)
}
func (m *GetContainersRequest) XXX_Size() int {
	return xxx_messageInfo_GetContainersRequest.Size(m)
}
func (m *GetContainersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContainersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetContainersRequest proto.InternalMessageInfo

func (m *GetContainersRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type GetContainersResponse struct {
	Containers           []*Container `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetContainersResponse) Reset()         { *m = GetContainersResponse{} }
func (m *GetContainersResponse) String() string { return proto.CompactTextString(m) }
func (*GetContainersResponse) ProtoMessage()    {}
func (*GetContainersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1358bea4b2532a29, []int{5}
}

func (m *GetContainersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContainersResponse.Unmarshal(m, b)
}
func (m *GetContainersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContainersResponse.Marshal(b, m, deterministic)
}
func (m *GetContainersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContainersResponse.Merge(m, src)
}
func (m *GetContainersResponse) XXX_Size() int {
	return xxx_messageInfo_GetContainersResponse.Size(m)
}
func (m *GetContainersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContainersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetContainersResponse proto.InternalMessageInfo

func (m *GetContainersResponse) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

type GetStatsRequest struct {
	Containers           map[string]int32 `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Host                 string           `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetStatsRequest) Reset()         { *m = GetStatsRequest{} }
func (m *GetStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatsRequest) ProtoMessage()    {}
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1358bea4b2532a29, []int{6}
}

func (m *GetStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsRequest.Unmarshal(m, b)
}
func (m *GetStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsRequest.Marshal(b, m, deterministic)
}
func (m *GetStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsRequest.Merge(m, src)
}
func (m *GetStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatsRequest.Size(m)
}
func (m *GetStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsRequest proto.InternalMessageInfo

func (m *GetStatsRequest) GetContainers() map[string]int32 {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *GetStatsRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type GetStatsReponse struct {
	Stats                map[int32]float32 `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetStatsReponse) Reset()         { *m = GetStatsReponse{} }
func (m *GetStatsReponse) String() string { return proto.CompactTextString(m) }
func (*GetStatsReponse) ProtoMessage()    {}
func (*GetStatsReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1358bea4b2532a29, []int{7}
}

func (m *GetStatsReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsReponse.Unmarshal(m, b)
}
func (m *GetStatsReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsReponse.Marshal(b, m, deterministic)
}
func (m *GetStatsReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsReponse.Merge(m, src)
}
func (m *GetStatsReponse) XXX_Size() int {
	return xxx_messageInfo_GetStatsReponse.Size(m)
}
func (m *GetStatsReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsReponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsReponse proto.InternalMessageInfo

func (m *GetStatsReponse) GetStats() map[int32]float32 {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*DockerHost)(nil), "DockerHost")
	proto.RegisterType((*HostContainerCount)(nil), "HostContainerCount")
	proto.RegisterMapType((map[string]int32)(nil), "HostContainerCount.ContainersEntry")
	proto.RegisterType((*GetContainersCountRequest)(nil), "GetContainersCountRequest")
	proto.RegisterType((*GetContainersCountResponse)(nil), "GetContainersCountResponse")
	proto.RegisterType((*GetContainersRequest)(nil), "GetContainersRequest")
	proto.RegisterType((*GetContainersResponse)(nil), "GetContainersResponse")
	proto.RegisterType((*GetStatsRequest)(nil), "GetStatsRequest")
	proto.RegisterMapType((map[string]int32)(nil), "GetStatsRequest.ContainersEntry")
	proto.RegisterType((*GetStatsReponse)(nil), "GetStatsReponse")
	proto.RegisterMapType((map[int32]float32)(nil), "GetStatsReponse.StatsEntry")
}

func init() { proto.RegisterFile("docker-host.proto", fileDescriptor_1358bea4b2532a29) }

var fileDescriptor_1358bea4b2532a29 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0xae, 0xd2, 0x40,
	0x14, 0xce, 0x14, 0x6a, 0x2e, 0xc7, 0x78, 0xe1, 0x8e, 0x60, 0xea, 0xe0, 0x82, 0xd4, 0x0d, 0x21,
	0x71, 0x82, 0xb8, 0x21, 0x1a, 0x8d, 0xb1, 0x1a, 0xdc, 0xe8, 0xa2, 0xac, 0x74, 0x57, 0xca, 0x24,
	0x34, 0xd4, 0x4e, 0xed, 0x0c, 0x24, 0x6c, 0x7c, 0x0e, 0x9f, 0xc0, 0x27, 0xf2, 0x81, 0xcc, 0xcc,
	0xf4, 0x8f, 0x52, 0xd9, 0xdc, 0xdd, 0x39, 0x67, 0xbe, 0x73, 0xbe, 0xef, 0xfc, 0xb4, 0x70, 0xb7,
	0xe5, 0xe1, 0x9e, 0x65, 0x2f, 0x76, 0x5c, 0x48, 0x9a, 0x66, 0x5c, 0x72, 0xd2, 0x0f, 0x79, 0x22,
	0x83, 0x28, 0x61, 0x99, 0x09, 0xb8, 0x5f, 0x01, 0x3e, 0x6a, 0xd4, 0x67, 0x2e, 0x24, 0xbe, 0x05,
	0x2b, 0x4a, 0x1d, 0x34, 0x41, 0xd3, 0x9e, 0x6f, 0x45, 0x29, 0xc6, 0xd0, 0x4d, 0x82, 0x1f, 0xcc,
	0xb1, 0x74, 0x44, 0xdb, 0xf8, 0x19, 0xf4, 0x42, 0x1e, 0xf3, 0xcc, 0xe3, 0x5b, 0xe6, 0x74, 0xf4,
	0x43, 0x15, 0x70, 0x7f, 0x23, 0xc0, 0xaa, 0x94, 0x57, 0xf0, 0x78, 0xfc, 0x90, 0x48, 0xec, 0x01,
	0x94, 0xcc, 0xc2, 0x41, 0x93, 0xce, 0xf4, 0xe1, 0xe2, 0x39, 0xbd, 0x04, 0xd2, 0xd2, 0x15, 0x9f,
	0x12, 0x99, 0x9d, 0xfc, 0x5a, 0x1a, 0x79, 0x0b, 0xfd, 0xc6, 0x33, 0x1e, 0x40, 0x67, 0xcf, 0x4e,
	0xb9, 0x62, 0x65, 0xe2, 0x21, 0xd8, 0xc7, 0x20, 0x3e, 0x18, 0xcd, 0xb6, 0x6f, 0x9c, 0xd7, 0xd6,
	0x12, 0xb9, 0x1e, 0x3c, 0x5d, 0xb1, 0x8a, 0x4f, 0x68, 0x42, 0x9f, 0xfd, 0x3c, 0x30, 0x21, 0x55,
	0x9a, 0x1a, 0x93, 0xd1, 0xd6, 0xf3, 0x8d, 0xa3, 0xca, 0x07, 0x71, 0xac, 0x4b, 0xdd, 0xf8, 0xca,
	0x74, 0xbf, 0x01, 0x69, 0x2b, 0x22, 0x52, 0x9e, 0x08, 0x86, 0xdf, 0xc0, 0xed, 0xae, 0xde, 0x53,
	0xd1, 0xea, 0xe3, 0x96, 0x56, 0xfd, 0x06, 0xd4, 0x9d, 0xc1, 0xf0, 0xac, 0x74, 0x21, 0x0d, 0x43,
	0x57, 0x21, 0xf3, 0x26, 0xb5, 0xed, 0x7a, 0x30, 0x6a, 0x60, 0x73, 0x05, 0xb3, 0x96, 0x41, 0x43,
	0x35, 0xd5, 0xfa, 0x3c, 0xdd, 0x3f, 0x08, 0xfa, 0x2b, 0x26, 0xd7, 0x32, 0x90, 0x25, 0xd9, 0xfb,
	0x96, 0xfc, 0x09, 0x6d, 0xa0, 0xae, 0x6d, 0xa9, 0x94, 0x6b, 0x55, 0x72, 0xef, 0xbb, 0xb9, 0x5f,
	0x75, 0x9d, 0xa6, 0xcf, 0x97, 0x60, 0x0b, 0xe5, 0xe7, 0x12, 0xc7, 0xb4, 0x01, 0xa0, 0xda, 0x31,
	0xea, 0x0c, 0x92, 0x2c, 0x01, 0xaa, 0x60, 0x9d, 0xdf, 0x6e, 0xe1, 0xb7, 0x6a, 0xfc, 0x8b, 0xbf,
	0x08, 0xee, 0xaa, 0xaf, 0x64, 0xcd, 0xb2, 0x63, 0x14, 0x32, 0xfc, 0x05, 0xf0, 0xe5, 0x29, 0x60,
	0x42, 0xff, 0x7b, 0x64, 0x64, 0x4c, 0xaf, 0xdc, 0xce, 0x3b, 0x78, 0x74, 0xf6, 0x8a, 0x47, 0xb4,
	0xed, 0x1c, 0xc8, 0x13, 0xda, 0xbe, 0xf9, 0x39, 0xdc, 0x14, 0x33, 0xc0, 0x83, 0xe6, 0xc6, 0xc8,
	0xa0, 0x39, 0xa0, 0x39, 0xfa, 0xd0, 0xfd, 0x6e, 0xa5, 0x9b, 0xcd, 0x03, 0xfd, 0x23, 0x78, 0xf5,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x59, 0x29, 0xd8, 0x2e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DockerHostServiceClient is the client API for DockerHostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DockerHostServiceClient interface {
	GetContainersCount(ctx context.Context, in *GetContainersCountRequest, opts ...grpc.CallOption) (*GetContainersCountResponse, error)
	GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*GetContainersResponse, error)
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (DockerHostService_GetStatsClient, error)
}

type dockerHostServiceClient struct {
	cc *grpc.ClientConn
}

func NewDockerHostServiceClient(cc *grpc.ClientConn) DockerHostServiceClient {
	return &dockerHostServiceClient{cc}
}

func (c *dockerHostServiceClient) GetContainersCount(ctx context.Context, in *GetContainersCountRequest, opts ...grpc.CallOption) (*GetContainersCountResponse, error) {
	out := new(GetContainersCountResponse)
	err := c.cc.Invoke(ctx, "/DockerHostService/GetContainersCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerHostServiceClient) GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*GetContainersResponse, error) {
	out := new(GetContainersResponse)
	err := c.cc.Invoke(ctx, "/DockerHostService/GetContainers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerHostServiceClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (DockerHostService_GetStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DockerHostService_serviceDesc.Streams[0], "/DockerHostService/GetStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &dockerHostServiceGetStatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DockerHostService_GetStatsClient interface {
	Recv() (*GetStatsReponse, error)
	grpc.ClientStream
}

type dockerHostServiceGetStatsClient struct {
	grpc.ClientStream
}

func (x *dockerHostServiceGetStatsClient) Recv() (*GetStatsReponse, error) {
	m := new(GetStatsReponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DockerHostServiceServer is the server API for DockerHostService service.
type DockerHostServiceServer interface {
	GetContainersCount(context.Context, *GetContainersCountRequest) (*GetContainersCountResponse, error)
	GetContainers(context.Context, *GetContainersRequest) (*GetContainersResponse, error)
	GetStats(*GetStatsRequest, DockerHostService_GetStatsServer) error
}

// UnimplementedDockerHostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDockerHostServiceServer struct {
}

func (*UnimplementedDockerHostServiceServer) GetContainersCount(ctx context.Context, req *GetContainersCountRequest) (*GetContainersCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainersCount not implemented")
}
func (*UnimplementedDockerHostServiceServer) GetContainers(ctx context.Context, req *GetContainersRequest) (*GetContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainers not implemented")
}
func (*UnimplementedDockerHostServiceServer) GetStats(req *GetStatsRequest, srv DockerHostService_GetStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}

func RegisterDockerHostServiceServer(s *grpc.Server, srv DockerHostServiceServer) {
	s.RegisterService(&_DockerHostService_serviceDesc, srv)
}

func _DockerHostService_GetContainersCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainersCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerHostServiceServer).GetContainersCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DockerHostService/GetContainersCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerHostServiceServer).GetContainersCount(ctx, req.(*GetContainersCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerHostService_GetContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerHostServiceServer).GetContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DockerHostService/GetContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerHostServiceServer).GetContainers(ctx, req.(*GetContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerHostService_GetStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DockerHostServiceServer).GetStats(m, &dockerHostServiceGetStatsServer{stream})
}

type DockerHostService_GetStatsServer interface {
	Send(*GetStatsReponse) error
	grpc.ServerStream
}

type dockerHostServiceGetStatsServer struct {
	grpc.ServerStream
}

func (x *dockerHostServiceGetStatsServer) Send(m *GetStatsReponse) error {
	return x.ServerStream.SendMsg(m)
}

var _DockerHostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DockerHostService",
	HandlerType: (*DockerHostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContainersCount",
			Handler:    _DockerHostService_GetContainersCount_Handler,
		},
		{
			MethodName: "GetContainers",
			Handler:    _DockerHostService_GetContainers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStats",
			Handler:       _DockerHostService_GetStats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "docker-host.proto",
}
