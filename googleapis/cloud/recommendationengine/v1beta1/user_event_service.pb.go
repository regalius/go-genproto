// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recommendationengine/v1beta1/user_event_service.proto

package recommendationengine

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	_ "google.golang.org/genproto/googleapis/type/date"
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

// Request message for PurgeUserEvents method.
type PurgeUserEventsRequest struct {
	// Required. The resource name of the event_store under which the events are
	// created. The format is
	// "projects/${projectId}/locations/global/catalogs/${catalogId}/eventStores/${eventStoreId}"
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The filter string to specify the events to be deleted. Empty
	// string filter is not allowed. This filter can also be used with
	// ListUserEvents API to list events that will be deleted. The eligible fields
	// for filtering are:
	// * eventType - UserEvent.eventType field of type string.
	// * eventTime - in ISO 8601 "zulu" format.
	// * visitorId - field of type string. Specifying this will delete all events
	// associated with a visitor.
	// * userId - field of type string. Specifying this will delete all events
	// associated with a user.
	// Example 1: Deleting all events in a time range.
	// `eventTime > "2012-04-23T18:25:43.511Z" eventTime <
	// "2012-04-23T18:30:43.511Z"`
	// Example 2: Deleting specific eventType in time range.
	// `eventTime > "2012-04-23T18:25:43.511Z" eventType = "detail-page-view"`
	// Example 3: Deleting all events for a specific visitor
	// `visitorId = visitor1024`
	// The filtering fields are assumed to have an implicit AND.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. The default value is false. Override this flag to true to
	// actually perform the purge. If the field is not set to true, a sampling of
	// events to be deleted will be returned.
	Force                bool     `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurgeUserEventsRequest) Reset()         { *m = PurgeUserEventsRequest{} }
func (m *PurgeUserEventsRequest) String() string { return proto.CompactTextString(m) }
func (*PurgeUserEventsRequest) ProtoMessage()    {}
func (*PurgeUserEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2db17806dba696a, []int{0}
}

func (m *PurgeUserEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeUserEventsRequest.Unmarshal(m, b)
}
func (m *PurgeUserEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeUserEventsRequest.Marshal(b, m, deterministic)
}
func (m *PurgeUserEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeUserEventsRequest.Merge(m, src)
}
func (m *PurgeUserEventsRequest) XXX_Size() int {
	return xxx_messageInfo_PurgeUserEventsRequest.Size(m)
}
func (m *PurgeUserEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeUserEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeUserEventsRequest proto.InternalMessageInfo

func (m *PurgeUserEventsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *PurgeUserEventsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *PurgeUserEventsRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

// Metadata related to the progress of the PurgeUserEvents operation.
// This will be returned by the google.longrunning.Operation.metadata field.
type PurgeUserEventsMetadata struct {
	// The ID of the request / operation.
	OperationName string `protobuf:"bytes,1,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	// Operation create time.
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PurgeUserEventsMetadata) Reset()         { *m = PurgeUserEventsMetadata{} }
func (m *PurgeUserEventsMetadata) String() string { return proto.CompactTextString(m) }
func (*PurgeUserEventsMetadata) ProtoMessage()    {}
func (*PurgeUserEventsMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2db17806dba696a, []int{1}
}

func (m *PurgeUserEventsMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeUserEventsMetadata.Unmarshal(m, b)
}
func (m *PurgeUserEventsMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeUserEventsMetadata.Marshal(b, m, deterministic)
}
func (m *PurgeUserEventsMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeUserEventsMetadata.Merge(m, src)
}
func (m *PurgeUserEventsMetadata) XXX_Size() int {
	return xxx_messageInfo_PurgeUserEventsMetadata.Size(m)
}
func (m *PurgeUserEventsMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeUserEventsMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeUserEventsMetadata proto.InternalMessageInfo

func (m *PurgeUserEventsMetadata) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *PurgeUserEventsMetadata) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

// Response of the PurgeUserEventsRequest. If the long running operation is
// successfully done, then this message is returned by the
// google.longrunning.Operations.response field.
type PurgeUserEventsResponse struct {
	// The total count of events purged as a result of the operation.
	PurgedEventsCount int64 `protobuf:"varint,1,opt,name=purged_events_count,json=purgedEventsCount,proto3" json:"purged_events_count,omitempty"`
	// A sampling of events deleted (or will be deleted) depending on the `force`
	// property in the request. Max of 500 items will be returned.
	UserEventsSample     []*UserEvent `protobuf:"bytes,2,rep,name=user_events_sample,json=userEventsSample,proto3" json:"user_events_sample,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PurgeUserEventsResponse) Reset()         { *m = PurgeUserEventsResponse{} }
func (m *PurgeUserEventsResponse) String() string { return proto.CompactTextString(m) }
func (*PurgeUserEventsResponse) ProtoMessage()    {}
func (*PurgeUserEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2db17806dba696a, []int{2}
}

func (m *PurgeUserEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeUserEventsResponse.Unmarshal(m, b)
}
func (m *PurgeUserEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeUserEventsResponse.Marshal(b, m, deterministic)
}
func (m *PurgeUserEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeUserEventsResponse.Merge(m, src)
}
func (m *PurgeUserEventsResponse) XXX_Size() int {
	return xxx_messageInfo_PurgeUserEventsResponse.Size(m)
}
func (m *PurgeUserEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeUserEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeUserEventsResponse proto.InternalMessageInfo

func (m *PurgeUserEventsResponse) GetPurgedEventsCount() int64 {
	if m != nil {
		return m.PurgedEventsCount
	}
	return 0
}

func (m *PurgeUserEventsResponse) GetUserEventsSample() []*UserEvent {
	if m != nil {
		return m.UserEventsSample
	}
	return nil
}

// Request message for WriteUserEvent method.
type WriteUserEventRequest struct {
	// Required. The parent eventStore resource name, such as
	// "projects/1234/locations/global/catalogs/default_catalog/eventStores/default_event_store".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. User event to write.
	UserEvent            *UserEvent `protobuf:"bytes,2,opt,name=user_event,json=userEvent,proto3" json:"user_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WriteUserEventRequest) Reset()         { *m = WriteUserEventRequest{} }
func (m *WriteUserEventRequest) String() string { return proto.CompactTextString(m) }
func (*WriteUserEventRequest) ProtoMessage()    {}
func (*WriteUserEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2db17806dba696a, []int{3}
}

func (m *WriteUserEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteUserEventRequest.Unmarshal(m, b)
}
func (m *WriteUserEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteUserEventRequest.Marshal(b, m, deterministic)
}
func (m *WriteUserEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteUserEventRequest.Merge(m, src)
}
func (m *WriteUserEventRequest) XXX_Size() int {
	return xxx_messageInfo_WriteUserEventRequest.Size(m)
}
func (m *WriteUserEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteUserEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteUserEventRequest proto.InternalMessageInfo

func (m *WriteUserEventRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *WriteUserEventRequest) GetUserEvent() *UserEvent {
	if m != nil {
		return m.UserEvent
	}
	return nil
}

// Request message for CollectUserEvent method.
type CollectUserEventRequest struct {
	// Required. The parent eventStore name, such as
	// "projects/1234/locations/global/catalogs/default_catalog/eventStores/default_event_store".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. URL encoded UserEvent proto.
	UserEvent string `protobuf:"bytes,2,opt,name=user_event,json=userEvent,proto3" json:"user_event,omitempty"`
	// The url including cgi-parameters but excluding the hash fragment. The URL
	// must be truncated to 1.5K bytes to conservatively be under the 2K bytes.
	// This is often more useful than the referer url, because many browsers only
	// send the domain for 3rd party requests.
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// The event timestamp in milliseconds. This prevents browser caching of
	// otherwise identical get requests. The name is abbreviated to reduce the
	// payload bytes.
	Ets                  int64    `protobuf:"varint,4,opt,name=ets,proto3" json:"ets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectUserEventRequest) Reset()         { *m = CollectUserEventRequest{} }
func (m *CollectUserEventRequest) String() string { return proto.CompactTextString(m) }
func (*CollectUserEventRequest) ProtoMessage()    {}
func (*CollectUserEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2db17806dba696a, []int{4}
}

func (m *CollectUserEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectUserEventRequest.Unmarshal(m, b)
}
func (m *CollectUserEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectUserEventRequest.Marshal(b, m, deterministic)
}
func (m *CollectUserEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectUserEventRequest.Merge(m, src)
}
func (m *CollectUserEventRequest) XXX_Size() int {
	return xxx_messageInfo_CollectUserEventRequest.Size(m)
}
func (m *CollectUserEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectUserEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectUserEventRequest proto.InternalMessageInfo

func (m *CollectUserEventRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CollectUserEventRequest) GetUserEvent() string {
	if m != nil {
		return m.UserEvent
	}
	return ""
}

func (m *CollectUserEventRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *CollectUserEventRequest) GetEts() int64 {
	if m != nil {
		return m.Ets
	}
	return 0
}

// Request message for ListUserEvents method.
type ListUserEventsRequest struct {
	// Required. The parent eventStore resource name, such as
	// "projects/*/locations/*/catalogs/default_catalog/eventStores/default_event_store".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Maximum number of results to return per page. If zero, the
	// service will choose a reasonable default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The previous ListUserEventsResponse.next_page_token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. Filtering expression to specify restrictions over
	// returned events. This is a sequence of terms, where each term applies some
	// kind of a restriction to the returned user events. Use this expression to
	// restrict results to a specific time range, or filter events by eventType.
	//    eg: eventTime > "2012-04-23T18:25:43.511Z" eventsMissingCatalogItems
	//    eventTime<"2012-04-23T18:25:43.511Z" eventType=search
	//
	//   We expect only 3 types of fields:
	//
	//    * eventTime: this can be specified a maximum of 2 times, once with a
	//      less than operator and once with a greater than operator. The
	//      eventTime restrict should result in one contiguous valid eventTime
	//      range.
	//
	//    * eventType: only 1 eventType restriction can be specified.
	//
	//    * eventsMissingCatalogItems: specififying this will restrict results
	//      to events for which catalog items were not found in the catalog. The
	//      default behavior is to return only those events for which catalog
	//      items were found.
	//
	//   Some examples of valid filters expressions:
	//
	//   * Example 1: eventTime > "2012-04-23T18:25:43.511Z"
	//             eventTime < "2012-04-23T18:30:43.511Z"
	//   * Example 2: eventTime > "2012-04-23T18:25:43.511Z"
	//             eventType = detail-page-view
	//   * Example 3: eventsMissingCatalogItems
	//             eventType = search eventTime < "2018-04-23T18:30:43.511Z"
	//   * Example 4: eventTime > "2012-04-23T18:25:43.511Z"
	//   * Example 5: eventType = search
	//   * Example 6: eventsMissingCatalogItems
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserEventsRequest) Reset()         { *m = ListUserEventsRequest{} }
func (m *ListUserEventsRequest) String() string { return proto.CompactTextString(m) }
func (*ListUserEventsRequest) ProtoMessage()    {}
func (*ListUserEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2db17806dba696a, []int{5}
}

func (m *ListUserEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserEventsRequest.Unmarshal(m, b)
}
func (m *ListUserEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserEventsRequest.Marshal(b, m, deterministic)
}
func (m *ListUserEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserEventsRequest.Merge(m, src)
}
func (m *ListUserEventsRequest) XXX_Size() int {
	return xxx_messageInfo_ListUserEventsRequest.Size(m)
}
func (m *ListUserEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserEventsRequest proto.InternalMessageInfo

func (m *ListUserEventsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListUserEventsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListUserEventsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListUserEventsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// Response message for ListUserEvents method.
type ListUserEventsResponse struct {
	// The user events.
	UserEvents []*UserEvent `protobuf:"bytes,1,rep,name=user_events,json=userEvents,proto3" json:"user_events,omitempty"`
	// If empty, the list is complete. If nonempty, the token to pass to the next
	// request's ListUserEvents.page_token.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserEventsResponse) Reset()         { *m = ListUserEventsResponse{} }
func (m *ListUserEventsResponse) String() string { return proto.CompactTextString(m) }
func (*ListUserEventsResponse) ProtoMessage()    {}
func (*ListUserEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2db17806dba696a, []int{6}
}

func (m *ListUserEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserEventsResponse.Unmarshal(m, b)
}
func (m *ListUserEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserEventsResponse.Marshal(b, m, deterministic)
}
func (m *ListUserEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserEventsResponse.Merge(m, src)
}
func (m *ListUserEventsResponse) XXX_Size() int {
	return xxx_messageInfo_ListUserEventsResponse.Size(m)
}
func (m *ListUserEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserEventsResponse proto.InternalMessageInfo

func (m *ListUserEventsResponse) GetUserEvents() []*UserEvent {
	if m != nil {
		return m.UserEvents
	}
	return nil
}

func (m *ListUserEventsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*PurgeUserEventsRequest)(nil), "google.cloud.recommendationengine.v1beta1.PurgeUserEventsRequest")
	proto.RegisterType((*PurgeUserEventsMetadata)(nil), "google.cloud.recommendationengine.v1beta1.PurgeUserEventsMetadata")
	proto.RegisterType((*PurgeUserEventsResponse)(nil), "google.cloud.recommendationengine.v1beta1.PurgeUserEventsResponse")
	proto.RegisterType((*WriteUserEventRequest)(nil), "google.cloud.recommendationengine.v1beta1.WriteUserEventRequest")
	proto.RegisterType((*CollectUserEventRequest)(nil), "google.cloud.recommendationengine.v1beta1.CollectUserEventRequest")
	proto.RegisterType((*ListUserEventsRequest)(nil), "google.cloud.recommendationengine.v1beta1.ListUserEventsRequest")
	proto.RegisterType((*ListUserEventsResponse)(nil), "google.cloud.recommendationengine.v1beta1.ListUserEventsResponse")
}

func init() {
	proto.RegisterFile("google/cloud/recommendationengine/v1beta1/user_event_service.proto", fileDescriptor_b2db17806dba696a)
}

var fileDescriptor_b2db17806dba696a = []byte{
	// 909 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0x26, 0x6d, 0x55, 0x4f, 0xd4, 0xd4, 0x0c, 0x6d, 0x62, 0x0c, 0x15, 0xd5, 0x56, 0xa0,
	0xd4, 0x52, 0x77, 0x94, 0x80, 0x38, 0xb8, 0x42, 0x22, 0x8e, 0x22, 0xa8, 0x54, 0x4a, 0xb4, 0x4e,
	0x1b, 0x84, 0x68, 0xcd, 0x78, 0xfd, 0xbc, 0x5d, 0xd8, 0x9d, 0x19, 0x66, 0x66, 0x13, 0x52, 0xd4,
	0x0b, 0xf0, 0x0d, 0x38, 0x70, 0xe7, 0x88, 0xc4, 0xb7, 0x40, 0x1c, 0x7a, 0x2c, 0xe2, 0x1b, 0x70,
	0xe3, 0xc2, 0x47, 0x40, 0xf3, 0xc7, 0xf6, 0xc6, 0x35, 0xc2, 0x4e, 0x73, 0xf3, 0xbe, 0xf7, 0xe6,
	0xf7, 0x7e, 0xbf, 0x37, 0x3f, 0xcf, 0x0c, 0xea, 0xa4, 0x9c, 0xa7, 0x39, 0x90, 0x24, 0xe7, 0xe5,
	0x80, 0x48, 0x48, 0x78, 0x51, 0x00, 0x1b, 0x50, 0x9d, 0x71, 0x06, 0x2c, 0xcd, 0x18, 0x90, 0xc3,
	0xcd, 0x3e, 0x68, 0xba, 0x49, 0x4a, 0x05, 0xb2, 0x07, 0x87, 0xc0, 0x74, 0x4f, 0x81, 0x3c, 0xcc,
	0x12, 0x88, 0x84, 0xe4, 0x9a, 0xe3, 0x9b, 0x0e, 0x23, 0xb2, 0x18, 0xd1, 0x2c, 0x8c, 0xc8, 0x63,
	0x34, 0xdf, 0xf0, 0xed, 0xa8, 0xc8, 0x08, 0x65, 0x8c, 0x6b, 0x5b, 0xa4, 0x1c, 0x50, 0x73, 0xbd,
	0x92, 0x4d, 0xf2, 0x0c, 0x98, 0xf6, 0x89, 0xd7, 0x2a, 0x89, 0xc7, 0x5a, 0x8b, 0x3e, 0x1f, 0x1c,
	0xfb, 0xd4, 0x7b, 0xf3, 0x0b, 0xc8, 0x0a, 0xc1, 0xe5, 0x08, 0xb2, 0x7d, 0x1a, 0xe1, 0x7e, 0xed,
	0x0d, 0xbf, 0x36, 0xe7, 0x2c, 0x95, 0x25, 0x63, 0x19, 0x4b, 0x09, 0x17, 0x20, 0x4f, 0x88, 0x79,
	0xd3, 0x17, 0xd9, 0xaf, 0x7e, 0x39, 0x24, 0x3a, 0x2b, 0x40, 0x69, 0x5a, 0x08, 0x5f, 0xb0, 0xe6,
	0x0b, 0xf4, 0xb1, 0x00, 0x32, 0xa0, 0xda, 0x8f, 0x33, 0x7c, 0x84, 0xd6, 0xf6, 0x4a, 0x99, 0xc2,
	0x7d, 0x05, 0x72, 0xd7, 0x74, 0x55, 0x31, 0x7c, 0x5d, 0x82, 0xd2, 0x78, 0x0d, 0x5d, 0x10, 0x54,
	0x02, 0xd3, 0x8d, 0xe0, 0x7a, 0xb0, 0x51, 0x8b, 0xfd, 0x97, 0x89, 0x0f, 0xb3, 0x5c, 0x83, 0x6c,
	0x2c, 0xb9, 0xb8, 0xfb, 0xc2, 0x57, 0xd0, 0xf9, 0x21, 0x97, 0x09, 0x34, 0x96, 0xaf, 0x07, 0x1b,
	0x17, 0x63, 0xf7, 0x11, 0x3e, 0x45, 0xeb, 0x53, 0xf8, 0x1f, 0x83, 0xa6, 0x03, 0xaa, 0x29, 0x7e,
	0x0b, 0xad, 0x8e, 0x75, 0xf4, 0x18, 0x2d, 0xc0, 0x37, 0xba, 0x34, 0x8e, 0xde, 0xa3, 0x05, 0xe0,
	0xdb, 0x68, 0x25, 0x91, 0x40, 0x35, 0xf4, 0x8c, 0x26, 0xdb, 0x74, 0x65, 0xab, 0x19, 0x79, 0x1b,
	0x8c, 0x04, 0x47, 0xfb, 0x23, 0xc1, 0x31, 0x72, 0xe5, 0x26, 0x10, 0xfe, 0x1a, 0xbc, 0xd0, 0x3f,
	0x06, 0x25, 0x38, 0x53, 0x80, 0x23, 0xf4, 0xaa, 0x30, 0xa9, 0x81, 0x1b, 0xb7, 0xea, 0x25, 0xbc,
	0xf4, 0x6a, 0x97, 0xe3, 0x57, 0x5c, 0xca, 0x2d, 0xd9, 0x31, 0x09, 0xdc, 0x47, 0x78, 0xb2, 0x39,
	0xaa, 0xa7, 0x68, 0x21, 0x72, 0xc3, 0x67, 0x79, 0x63, 0x65, 0xeb, 0xdd, 0x68, 0x6e, 0x5b, 0x46,
	0x63, 0x2a, 0x71, 0xbd, 0x1c, 0xb3, 0xea, 0x5a, 0xb4, 0xf0, 0x87, 0x00, 0x5d, 0x3d, 0x90, 0x99,
	0x9e, 0xf0, 0xfd, 0xbf, 0xed, 0xe8, 0x22, 0x34, 0x61, 0xe5, 0xa7, 0x73, 0x3a, 0x36, 0xb5, 0x31,
	0x9b, 0x50, 0xa3, 0xf5, 0x1d, 0x9e, 0xe7, 0x90, 0xe8, 0xb9, 0x79, 0x5c, 0x7b, 0x81, 0x47, 0xad,
	0x82, 0x88, 0xeb, 0x68, 0xb9, 0x94, 0x99, 0xf5, 0x46, 0x2d, 0x36, 0x3f, 0x4d, 0x04, 0xb4, 0x6a,
	0x9c, 0xb3, 0xe3, 0x36, 0x3f, 0xc3, 0xef, 0x03, 0x74, 0xf5, 0x6e, 0xa6, 0xf4, 0xfc, 0x5e, 0x7c,
	0x1d, 0xd5, 0x04, 0x4d, 0xa1, 0xa7, 0xb2, 0x27, 0xce, 0x19, 0xe7, 0xe3, 0x8b, 0x26, 0xd0, 0xcd,
	0x9e, 0x80, 0x61, 0x64, 0x93, 0x9a, 0x7f, 0x05, 0xcc, 0x77, 0xb6, 0xe5, 0xfb, 0x26, 0x50, 0xf1,
	0xf1, 0xb9, 0xaa, 0x8f, 0xc3, 0x9f, 0x02, 0xb4, 0x36, 0xcd, 0xc2, 0x3b, 0xe6, 0x3e, 0x5a, 0xa9,
	0x38, 0xa0, 0x11, 0xbc, 0xc4, 0xd6, 0xa3, 0xc9, 0xd6, 0xe3, 0xb7, 0xd1, 0x65, 0x06, 0xdf, 0xe8,
	0x5e, 0x85, 0xad, 0x9b, 0xdf, 0x25, 0x13, 0xde, 0x1b, 0x31, 0xde, 0xfa, 0xbd, 0x86, 0xea, 0x63,
	0x84, 0xae, 0x3b, 0x15, 0xf1, 0x3f, 0x01, 0x5a, 0x3d, 0xe9, 0x18, 0xfc, 0xc1, 0x02, 0x8c, 0x66,
	0x9a, 0xad, 0x79, 0x2a, 0x4d, 0xe1, 0xf0, 0xbb, 0x3f, 0xfe, 0xfa, 0x71, 0xe9, 0x8b, 0x30, 0x1e,
	0x1f, 0x66, 0xdf, 0xba, 0x7d, 0x7a, 0x5f, 0x48, 0xfe, 0x25, 0x24, 0x5a, 0x91, 0x16, 0xc9, 0x79,
	0xe2, 0x4e, 0x2e, 0xd2, 0x22, 0x09, 0xd5, 0x34, 0xe7, 0xa9, 0xf9, 0x69, 0x47, 0xda, 0xd5, 0x5c,
	0x82, 0x22, 0xad, 0xa7, 0x64, 0x32, 0x9d, 0xf6, 0x91, 0x61, 0xd8, 0xae, 0x98, 0x0b, 0x3f, 0x0b,
	0x50, 0x7d, 0xda, 0x9e, 0xb8, 0xb3, 0x00, 0xe5, 0xff, 0xf0, 0x76, 0xf3, 0xca, 0x08, 0x83, 0x8a,
	0x2c, 0xfa, 0x48, 0x6b, 0xd1, 0xe1, 0x83, 0xe3, 0xf0, 0x73, 0x2b, 0xeb, 0x01, 0xde, 0x3f, 0x53,
	0x59, 0x89, 0xe3, 0x80, 0xff, 0x0e, 0xd0, 0xea, 0x49, 0xbb, 0x2d, 0xb4, 0x7f, 0x33, 0xff, 0x2f,
	0xcd, 0xed, 0x97, 0x40, 0x70, 0x5e, 0x0f, 0xf7, 0xad, 0xea, 0x7b, 0xf8, 0xee, 0x59, 0xaa, 0xc6,
	0xcf, 0x03, 0x74, 0x79, 0xea, 0x3c, 0xc6, 0x8b, 0x90, 0x9d, 0x7d, 0x57, 0x35, 0xaf, 0x8d, 0x20,
	0x2a, 0x97, 0x64, 0xf4, 0xc9, 0xe8, 0x1a, 0x09, 0x1f, 0x5a, 0x2d, 0x07, 0x67, 0x6c, 0x4c, 0x7b,
	0x43, 0xb4, 0x83, 0x16, 0xfe, 0x33, 0x40, 0xf5, 0x3b, 0xf6, 0xba, 0xaf, 0xa8, 0x5a, 0xc4, 0x8f,
	0xd3, 0x8b, 0xe7, 0x94, 0xf5, 0xc8, 0xca, 0xfa, 0x34, 0xec, 0x9e, 0xa9, 0x2c, 0xf7, 0x70, 0x69,
	0x07, 0xad, 0xe6, 0xc1, 0xb3, 0xed, 0x1b, 0x33, 0x79, 0x3b, 0x56, 0x54, 0x64, 0x2a, 0x4a, 0x78,
	0xf1, 0x7c, 0x3b, 0x32, 0xcf, 0x24, 0xd5, 0x26, 0xe4, 0xe8, 0xe8, 0x68, 0x2a, 0x49, 0x68, 0xa9,
	0x1f, 0xbb, 0x77, 0xcf, 0x2d, 0x91, 0x53, 0x3d, 0xe4, 0xb2, 0xe8, 0xfc, 0x16, 0xa0, 0x5b, 0x09,
	0x2f, 0xe6, 0x9f, 0xd0, 0x5e, 0xf0, 0xd9, 0x43, 0x5f, 0x9c, 0xf2, 0x9c, 0xb2, 0x34, 0xe2, 0x32,
	0x25, 0x29, 0x30, 0x7b, 0xfd, 0x93, 0x49, 0xcb, 0x39, 0x5e, 0x58, 0xb7, 0x67, 0x25, 0x7f, 0x5e,
	0x3a, 0x1f, 0xef, 0xee, 0x6c, 0xdf, 0xf9, 0x65, 0xe9, 0xe6, 0x87, 0xae, 0xcf, 0x8e, 0x25, 0x15,
	0x9f, 0xa8, 0xdd, 0x75, 0xa4, 0x1e, 0x6c, 0x76, 0x0c, 0x50, 0xff, 0x82, 0xed, 0xfe, 0xce, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0xb3, 0x11, 0x07, 0xd7, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserEventServiceClient is the client API for UserEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserEventServiceClient interface {
	// Writes a single user event.
	WriteUserEvent(ctx context.Context, in *WriteUserEventRequest, opts ...grpc.CallOption) (*UserEvent, error)
	// Writes a single user event from the browser. This uses a GET request to
	// due to browser restriction of POST-ing to a 3rd party domain.
	//
	// This method is used only by the Recommendations AI JavaScript pixel.
	// Users should not call this method directly.
	CollectUserEvent(ctx context.Context, in *CollectUserEventRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// Gets a list of user events within a time range, with potential filtering.
	ListUserEvents(ctx context.Context, in *ListUserEventsRequest, opts ...grpc.CallOption) (*ListUserEventsResponse, error)
	// Deletes permanently all user events specified by the filter provided.
	// Depending on the number of events specified by the filter, this operation
	// could take hours or days to complete. To test a filter, use the list
	// command first.
	PurgeUserEvents(ctx context.Context, in *PurgeUserEventsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Bulk import of User events. Request processing might be
	// synchronous. Events that already exist are skipped.
	// Use this method for backfilling historical user events.
	//
	// Operation.response is of type ImportResponse. Note that it is
	// possible for a subset of the items to be successfully inserted.
	// Operation.metadata is of type ImportMetadata.
	ImportUserEvents(ctx context.Context, in *ImportUserEventsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type userEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserEventServiceClient(cc grpc.ClientConnInterface) UserEventServiceClient {
	return &userEventServiceClient{cc}
}

func (c *userEventServiceClient) WriteUserEvent(ctx context.Context, in *WriteUserEventRequest, opts ...grpc.CallOption) (*UserEvent, error) {
	out := new(UserEvent)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.UserEventService/WriteUserEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventServiceClient) CollectUserEvent(ctx context.Context, in *CollectUserEventRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.UserEventService/CollectUserEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventServiceClient) ListUserEvents(ctx context.Context, in *ListUserEventsRequest, opts ...grpc.CallOption) (*ListUserEventsResponse, error) {
	out := new(ListUserEventsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.UserEventService/ListUserEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventServiceClient) PurgeUserEvents(ctx context.Context, in *PurgeUserEventsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.UserEventService/PurgeUserEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventServiceClient) ImportUserEvents(ctx context.Context, in *ImportUserEventsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.UserEventService/ImportUserEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserEventServiceServer is the server API for UserEventService service.
type UserEventServiceServer interface {
	// Writes a single user event.
	WriteUserEvent(context.Context, *WriteUserEventRequest) (*UserEvent, error)
	// Writes a single user event from the browser. This uses a GET request to
	// due to browser restriction of POST-ing to a 3rd party domain.
	//
	// This method is used only by the Recommendations AI JavaScript pixel.
	// Users should not call this method directly.
	CollectUserEvent(context.Context, *CollectUserEventRequest) (*httpbody.HttpBody, error)
	// Gets a list of user events within a time range, with potential filtering.
	ListUserEvents(context.Context, *ListUserEventsRequest) (*ListUserEventsResponse, error)
	// Deletes permanently all user events specified by the filter provided.
	// Depending on the number of events specified by the filter, this operation
	// could take hours or days to complete. To test a filter, use the list
	// command first.
	PurgeUserEvents(context.Context, *PurgeUserEventsRequest) (*longrunning.Operation, error)
	// Bulk import of User events. Request processing might be
	// synchronous. Events that already exist are skipped.
	// Use this method for backfilling historical user events.
	//
	// Operation.response is of type ImportResponse. Note that it is
	// possible for a subset of the items to be successfully inserted.
	// Operation.metadata is of type ImportMetadata.
	ImportUserEvents(context.Context, *ImportUserEventsRequest) (*longrunning.Operation, error)
}

// UnimplementedUserEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserEventServiceServer struct {
}

func (*UnimplementedUserEventServiceServer) WriteUserEvent(ctx context.Context, req *WriteUserEventRequest) (*UserEvent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteUserEvent not implemented")
}
func (*UnimplementedUserEventServiceServer) CollectUserEvent(ctx context.Context, req *CollectUserEventRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectUserEvent not implemented")
}
func (*UnimplementedUserEventServiceServer) ListUserEvents(ctx context.Context, req *ListUserEventsRequest) (*ListUserEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserEvents not implemented")
}
func (*UnimplementedUserEventServiceServer) PurgeUserEvents(ctx context.Context, req *PurgeUserEventsRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeUserEvents not implemented")
}
func (*UnimplementedUserEventServiceServer) ImportUserEvents(ctx context.Context, req *ImportUserEventsRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportUserEvents not implemented")
}

func RegisterUserEventServiceServer(s *grpc.Server, srv UserEventServiceServer) {
	s.RegisterService(&_UserEventService_serviceDesc, srv)
}

func _UserEventService_WriteUserEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteUserEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServiceServer).WriteUserEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.UserEventService/WriteUserEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServiceServer).WriteUserEvent(ctx, req.(*WriteUserEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEventService_CollectUserEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectUserEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServiceServer).CollectUserEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.UserEventService/CollectUserEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServiceServer).CollectUserEvent(ctx, req.(*CollectUserEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEventService_ListUserEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServiceServer).ListUserEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.UserEventService/ListUserEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServiceServer).ListUserEvents(ctx, req.(*ListUserEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEventService_PurgeUserEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeUserEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServiceServer).PurgeUserEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.UserEventService/PurgeUserEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServiceServer).PurgeUserEvents(ctx, req.(*PurgeUserEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEventService_ImportUserEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportUserEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServiceServer).ImportUserEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.UserEventService/ImportUserEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServiceServer).ImportUserEvents(ctx, req.(*ImportUserEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserEventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.recommendationengine.v1beta1.UserEventService",
	HandlerType: (*UserEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteUserEvent",
			Handler:    _UserEventService_WriteUserEvent_Handler,
		},
		{
			MethodName: "CollectUserEvent",
			Handler:    _UserEventService_CollectUserEvent_Handler,
		},
		{
			MethodName: "ListUserEvents",
			Handler:    _UserEventService_ListUserEvents_Handler,
		},
		{
			MethodName: "PurgeUserEvents",
			Handler:    _UserEventService_PurgeUserEvents_Handler,
		},
		{
			MethodName: "ImportUserEvents",
			Handler:    _UserEventService_ImportUserEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/recommendationengine/v1beta1/user_event_service.proto",
}
