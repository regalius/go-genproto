// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/customer_client_link_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [CustomerClientLinkService.GetCustomerClientLink][google.ads.googleads.v2.services.CustomerClientLinkService.GetCustomerClientLink].
type GetCustomerClientLinkRequest struct {
	// The resource name of the customer client link to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerClientLinkRequest) Reset()         { *m = GetCustomerClientLinkRequest{} }
func (m *GetCustomerClientLinkRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerClientLinkRequest) ProtoMessage()    {}
func (*GetCustomerClientLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f0870b920ea667, []int{0}
}

func (m *GetCustomerClientLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerClientLinkRequest.Unmarshal(m, b)
}
func (m *GetCustomerClientLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerClientLinkRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerClientLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerClientLinkRequest.Merge(m, src)
}
func (m *GetCustomerClientLinkRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerClientLinkRequest.Size(m)
}
func (m *GetCustomerClientLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerClientLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerClientLinkRequest proto.InternalMessageInfo

func (m *GetCustomerClientLinkRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CustomerClientLinkService.MutateCustomerClientLink][google.ads.googleads.v2.services.CustomerClientLinkService.MutateCustomerClientLink].
type MutateCustomerClientLinkRequest struct {
	// The ID of the customer whose customer link are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform on the individual CustomerClientLink.
	Operation            *CustomerClientLinkOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateCustomerClientLinkRequest) Reset()         { *m = MutateCustomerClientLinkRequest{} }
func (m *MutateCustomerClientLinkRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerClientLinkRequest) ProtoMessage()    {}
func (*MutateCustomerClientLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f0870b920ea667, []int{1}
}

func (m *MutateCustomerClientLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerClientLinkRequest.Unmarshal(m, b)
}
func (m *MutateCustomerClientLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerClientLinkRequest.Marshal(b, m, deterministic)
}
func (m *MutateCustomerClientLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerClientLinkRequest.Merge(m, src)
}
func (m *MutateCustomerClientLinkRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerClientLinkRequest.Size(m)
}
func (m *MutateCustomerClientLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerClientLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerClientLinkRequest proto.InternalMessageInfo

func (m *MutateCustomerClientLinkRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerClientLinkRequest) GetOperation() *CustomerClientLinkOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// A single operation (create, update) on a CustomerClientLink.
type CustomerClientLinkOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CustomerClientLinkOperation_Create
	//	*CustomerClientLinkOperation_Update
	Operation            isCustomerClientLinkOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CustomerClientLinkOperation) Reset()         { *m = CustomerClientLinkOperation{} }
func (m *CustomerClientLinkOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerClientLinkOperation) ProtoMessage()    {}
func (*CustomerClientLinkOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f0870b920ea667, []int{2}
}

func (m *CustomerClientLinkOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerClientLinkOperation.Unmarshal(m, b)
}
func (m *CustomerClientLinkOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerClientLinkOperation.Marshal(b, m, deterministic)
}
func (m *CustomerClientLinkOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerClientLinkOperation.Merge(m, src)
}
func (m *CustomerClientLinkOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerClientLinkOperation.Size(m)
}
func (m *CustomerClientLinkOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerClientLinkOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerClientLinkOperation proto.InternalMessageInfo

func (m *CustomerClientLinkOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCustomerClientLinkOperation_Operation interface {
	isCustomerClientLinkOperation_Operation()
}

type CustomerClientLinkOperation_Create struct {
	Create *resources.CustomerClientLink `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerClientLinkOperation_Update struct {
	Update *resources.CustomerClientLink `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

func (*CustomerClientLinkOperation_Create) isCustomerClientLinkOperation_Operation() {}

func (*CustomerClientLinkOperation_Update) isCustomerClientLinkOperation_Operation() {}

func (m *CustomerClientLinkOperation) GetOperation() isCustomerClientLinkOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CustomerClientLinkOperation) GetCreate() *resources.CustomerClientLink {
	if x, ok := m.GetOperation().(*CustomerClientLinkOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CustomerClientLinkOperation) GetUpdate() *resources.CustomerClientLink {
	if x, ok := m.GetOperation().(*CustomerClientLinkOperation_Update); ok {
		return x.Update
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomerClientLinkOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomerClientLinkOperation_Create)(nil),
		(*CustomerClientLinkOperation_Update)(nil),
	}
}

// Response message for a CustomerClientLink mutate.
type MutateCustomerClientLinkResponse struct {
	// A result that identifies the resource affected by the mutate request.
	Result               *MutateCustomerClientLinkResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MutateCustomerClientLinkResponse) Reset()         { *m = MutateCustomerClientLinkResponse{} }
func (m *MutateCustomerClientLinkResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerClientLinkResponse) ProtoMessage()    {}
func (*MutateCustomerClientLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f0870b920ea667, []int{3}
}

func (m *MutateCustomerClientLinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerClientLinkResponse.Unmarshal(m, b)
}
func (m *MutateCustomerClientLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerClientLinkResponse.Marshal(b, m, deterministic)
}
func (m *MutateCustomerClientLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerClientLinkResponse.Merge(m, src)
}
func (m *MutateCustomerClientLinkResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerClientLinkResponse.Size(m)
}
func (m *MutateCustomerClientLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerClientLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerClientLinkResponse proto.InternalMessageInfo

func (m *MutateCustomerClientLinkResponse) GetResult() *MutateCustomerClientLinkResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// The result for a single customer client link mutate.
type MutateCustomerClientLinkResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerClientLinkResult) Reset()         { *m = MutateCustomerClientLinkResult{} }
func (m *MutateCustomerClientLinkResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerClientLinkResult) ProtoMessage()    {}
func (*MutateCustomerClientLinkResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f0870b920ea667, []int{4}
}

func (m *MutateCustomerClientLinkResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerClientLinkResult.Unmarshal(m, b)
}
func (m *MutateCustomerClientLinkResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerClientLinkResult.Marshal(b, m, deterministic)
}
func (m *MutateCustomerClientLinkResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerClientLinkResult.Merge(m, src)
}
func (m *MutateCustomerClientLinkResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerClientLinkResult.Size(m)
}
func (m *MutateCustomerClientLinkResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerClientLinkResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerClientLinkResult proto.InternalMessageInfo

func (m *MutateCustomerClientLinkResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerClientLinkRequest)(nil), "google.ads.googleads.v2.services.GetCustomerClientLinkRequest")
	proto.RegisterType((*MutateCustomerClientLinkRequest)(nil), "google.ads.googleads.v2.services.MutateCustomerClientLinkRequest")
	proto.RegisterType((*CustomerClientLinkOperation)(nil), "google.ads.googleads.v2.services.CustomerClientLinkOperation")
	proto.RegisterType((*MutateCustomerClientLinkResponse)(nil), "google.ads.googleads.v2.services.MutateCustomerClientLinkResponse")
	proto.RegisterType((*MutateCustomerClientLinkResult)(nil), "google.ads.googleads.v2.services.MutateCustomerClientLinkResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/customer_client_link_service.proto", fileDescriptor_75f0870b920ea667)
}

var fileDescriptor_75f0870b920ea667 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xc1, 0x6b, 0xd4, 0x4e,
	0x14, 0xc7, 0x7f, 0x49, 0x7f, 0x14, 0x3a, 0xab, 0x97, 0x80, 0xb8, 0x6e, 0x4b, 0xbb, 0xc4, 0x1e,
	0xca, 0x1e, 0x66, 0x20, 0x52, 0x94, 0xd4, 0x15, 0xb3, 0x8b, 0xb6, 0x82, 0xb5, 0x25, 0xc2, 0x22,
	0xba, 0xb0, 0x4c, 0x93, 0x69, 0x08, 0x9b, 0x64, 0x62, 0x66, 0xb2, 0x97, 0xda, 0x8b, 0xe0, 0xc1,
	0xb3, 0x7f, 0x80, 0xe0, 0xd1, 0x3f, 0xa5, 0xe0, 0xc9, 0x7f, 0x41, 0x3c, 0x78, 0xf3, 0x3f, 0x90,
	0x64, 0x66, 0xb2, 0x2b, 0xdd, 0xec, 0x4a, 0x7b, 0x7b, 0x3b, 0xf3, 0xde, 0xe7, 0xbd, 0xef, 0x7b,
	0xf3, 0x36, 0xa0, 0x1f, 0x50, 0x1a, 0x44, 0x04, 0x61, 0x9f, 0x21, 0x61, 0x16, 0xd6, 0xc4, 0x42,
	0x8c, 0x64, 0x93, 0xd0, 0x23, 0x0c, 0x79, 0x39, 0xe3, 0x34, 0x26, 0xd9, 0xc8, 0x8b, 0x42, 0x92,
	0xf0, 0x51, 0x14, 0x26, 0xe3, 0x91, 0xbc, 0x85, 0x69, 0x46, 0x39, 0x35, 0xda, 0x22, 0x12, 0x62,
	0x9f, 0xc1, 0x0a, 0x02, 0x27, 0x16, 0x54, 0x90, 0xd6, 0xc3, 0xba, 0x34, 0x19, 0x61, 0x34, 0xcf,
	0xea, 0xf2, 0x08, 0x7e, 0x6b, 0x43, 0x45, 0xa7, 0x21, 0xc2, 0x49, 0x42, 0x39, 0xe6, 0x21, 0x4d,
	0x98, 0xbc, 0x95, 0xd9, 0x51, 0xf9, 0xeb, 0x24, 0x3f, 0x45, 0xa7, 0x21, 0x89, 0xfc, 0x51, 0x8c,
	0x99, 0x8a, 0xbf, 0x3d, 0x13, 0x2f, 0xe8, 0xe2, 0xc2, 0xec, 0x83, 0x8d, 0x7d, 0xc2, 0xfb, 0x32,
	0x73, 0xbf, 0xbc, 0x7a, 0x1e, 0x26, 0x63, 0x97, 0xbc, 0xcd, 0x09, 0xe3, 0xc6, 0x5d, 0x70, 0x53,
	0x15, 0x38, 0x4a, 0x70, 0x4c, 0x9a, 0x5a, 0x5b, 0xdb, 0x59, 0x73, 0x6f, 0xa8, 0xc3, 0x17, 0x38,
	0x26, 0xe6, 0x67, 0x0d, 0x6c, 0x1d, 0xe6, 0x1c, 0x73, 0x52, 0x0f, 0xda, 0x02, 0x8d, 0x4a, 0x5f,
	0xe8, 0x4b, 0x0c, 0x50, 0x47, 0xcf, 0x7c, 0xe3, 0x0d, 0x58, 0xa3, 0x29, 0xc9, 0x4a, 0x61, 0x4d,
	0xbd, 0xad, 0xed, 0x34, 0xac, 0x2e, 0x5c, 0xd6, 0x56, 0x78, 0x39, 0xe1, 0x91, 0x82, 0xb8, 0x53,
	0x9e, 0xf9, 0x51, 0x07, 0xeb, 0x0b, 0x5c, 0x8d, 0x3d, 0xd0, 0xc8, 0x53, 0x1f, 0x73, 0x52, 0x36,
	0xad, 0xf9, 0x7f, 0x99, 0xbe, 0xa5, 0xd2, 0xab, 0xbe, 0xc2, 0xa7, 0x45, 0x5f, 0x0f, 0x31, 0x1b,
	0xbb, 0x40, 0xb8, 0x17, 0xb6, 0x71, 0x04, 0x56, 0xbd, 0x8c, 0x60, 0x2e, 0x9a, 0xd3, 0xb0, 0x76,
	0x6b, 0xcb, 0xae, 0x66, 0x3d, 0xa7, 0xee, 0x83, 0xff, 0x5c, 0x89, 0x29, 0x80, 0x02, 0x2f, 0xfb,
	0x70, 0x75, 0xa0, 0xc0, 0xf4, 0x1a, 0x33, 0xbd, 0x35, 0xdf, 0x81, 0x76, 0xfd, 0xb0, 0x58, 0x4a,
	0x13, 0x46, 0x8c, 0x57, 0x60, 0x35, 0x23, 0x2c, 0x8f, 0xb8, 0x94, 0xf4, 0x78, 0xf9, 0x24, 0x16,
	0x30, 0xf3, 0x88, 0xbb, 0x92, 0x67, 0x3e, 0x01, 0x9b, 0x8b, 0x3d, 0xff, 0xe9, 0xc9, 0x59, 0xbf,
	0x57, 0xc0, 0x9d, 0xcb, 0x84, 0x97, 0xa2, 0x18, 0xe3, 0x9b, 0x06, 0x6e, 0xcd, 0x7d, 0xd6, 0xc6,
	0xa3, 0xe5, 0x42, 0x16, 0xed, 0x43, 0xeb, 0x6a, 0xa3, 0x30, 0xbb, 0xef, 0xbf, 0xff, 0xf8, 0xa4,
	0xdf, 0x37, 0x76, 0x8b, 0x8d, 0x3f, 0xfb, 0x4b, 0x5e, 0x57, 0xed, 0x00, 0x43, 0x9d, 0xea, 0x2f,
	0x60, 0x1a, 0xca, 0x50, 0xe7, 0xdc, 0xf8, 0xa9, 0x81, 0x66, 0x5d, 0xd7, 0x0c, 0xe7, 0x3a, 0xb3,
	0x11, 0xaa, 0x7a, 0xd7, 0x1a, 0x6f, 0xf9, 0x64, 0xcc, 0x7e, 0x29, 0xb1, 0x6b, 0x3e, 0x28, 0x24,
	0x4e, 0x35, 0x9d, 0xcd, 0x6c, 0x7d, 0xb7, 0x73, 0x3e, 0x4f, 0xa1, 0x1d, 0x97, 0x6c, 0x5b, 0xeb,
	0xb4, 0xd6, 0x2f, 0x9c, 0xe6, 0x34, 0xbf, 0xb4, 0xd2, 0x90, 0x41, 0x8f, 0xc6, 0xbd, 0x0f, 0x3a,
	0xd8, 0xf6, 0x68, 0xbc, 0xb4, 0xd6, 0xde, 0x66, 0xed, 0xcb, 0x38, 0x2e, 0x36, 0xf9, 0x58, 0x7b,
	0x7d, 0x20, 0x19, 0x01, 0x8d, 0x70, 0x12, 0x40, 0x9a, 0x05, 0x28, 0x20, 0x49, 0xb9, 0xe7, 0x68,
	0x9a, 0xb5, 0xfe, 0x9b, 0xb0, 0xa7, 0x8c, 0x2f, 0xfa, 0xca, 0xbe, 0xe3, 0x7c, 0xd5, 0xdb, 0xfb,
	0x02, 0xe8, 0xf8, 0x0c, 0x0a, 0xb3, 0xb0, 0x06, 0x16, 0x94, 0x89, 0xd9, 0x85, 0x72, 0x19, 0x3a,
	0x3e, 0x1b, 0x56, 0x2e, 0xc3, 0x81, 0x35, 0x54, 0x2e, 0xbf, 0xf4, 0x6d, 0x71, 0x6e, 0xdb, 0x8e,
	0xcf, 0x6c, 0xbb, 0x72, 0xb2, 0xed, 0x81, 0x65, 0xdb, 0xca, 0xed, 0x64, 0xb5, 0xac, 0xf3, 0xde,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x5b, 0xb9, 0xb9, 0xba, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerClientLinkServiceClient is the client API for CustomerClientLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClientLinkServiceClient interface {
	// Returns the requested CustomerClientLink in full detail.
	GetCustomerClientLink(ctx context.Context, in *GetCustomerClientLinkRequest, opts ...grpc.CallOption) (*resources.CustomerClientLink, error)
	// Creates or updates a customer client link. Operation statuses are returned.
	MutateCustomerClientLink(ctx context.Context, in *MutateCustomerClientLinkRequest, opts ...grpc.CallOption) (*MutateCustomerClientLinkResponse, error)
}

type customerClientLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClientLinkServiceClient(cc grpc.ClientConnInterface) CustomerClientLinkServiceClient {
	return &customerClientLinkServiceClient{cc}
}

func (c *customerClientLinkServiceClient) GetCustomerClientLink(ctx context.Context, in *GetCustomerClientLinkRequest, opts ...grpc.CallOption) (*resources.CustomerClientLink, error) {
	out := new(resources.CustomerClientLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CustomerClientLinkService/GetCustomerClientLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClientLinkServiceClient) MutateCustomerClientLink(ctx context.Context, in *MutateCustomerClientLinkRequest, opts ...grpc.CallOption) (*MutateCustomerClientLinkResponse, error) {
	out := new(MutateCustomerClientLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CustomerClientLinkService/MutateCustomerClientLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientLinkServiceServer is the server API for CustomerClientLinkService service.
type CustomerClientLinkServiceServer interface {
	// Returns the requested CustomerClientLink in full detail.
	GetCustomerClientLink(context.Context, *GetCustomerClientLinkRequest) (*resources.CustomerClientLink, error)
	// Creates or updates a customer client link. Operation statuses are returned.
	MutateCustomerClientLink(context.Context, *MutateCustomerClientLinkRequest) (*MutateCustomerClientLinkResponse, error)
}

// UnimplementedCustomerClientLinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerClientLinkServiceServer struct {
}

func (*UnimplementedCustomerClientLinkServiceServer) GetCustomerClientLink(ctx context.Context, req *GetCustomerClientLinkRequest) (*resources.CustomerClientLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerClientLink not implemented")
}
func (*UnimplementedCustomerClientLinkServiceServer) MutateCustomerClientLink(ctx context.Context, req *MutateCustomerClientLinkRequest) (*MutateCustomerClientLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerClientLink not implemented")
}

func RegisterCustomerClientLinkServiceServer(s *grpc.Server, srv CustomerClientLinkServiceServer) {
	s.RegisterService(&_CustomerClientLinkService_serviceDesc, srv)
}

func _CustomerClientLinkService_GetCustomerClientLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerClientLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientLinkServiceServer).GetCustomerClientLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CustomerClientLinkService/GetCustomerClientLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientLinkServiceServer).GetCustomerClientLink(ctx, req.(*GetCustomerClientLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerClientLinkService_MutateCustomerClientLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerClientLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientLinkServiceServer).MutateCustomerClientLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CustomerClientLinkService/MutateCustomerClientLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientLinkServiceServer).MutateCustomerClientLink(ctx, req.(*MutateCustomerClientLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerClientLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CustomerClientLinkService",
	HandlerType: (*CustomerClientLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerClientLink",
			Handler:    _CustomerClientLinkService_GetCustomerClientLink_Handler,
		},
		{
			MethodName: "MutateCustomerClientLink",
			Handler:    _CustomerClientLinkService_MutateCustomerClientLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/customer_client_link_service.proto",
}
