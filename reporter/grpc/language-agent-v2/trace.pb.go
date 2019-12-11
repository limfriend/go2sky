// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent-v2/trace.proto

package language_agent_v2

import (
	context "context"
	fmt "fmt"
	common "github.com/SkyAPM/go2sky/reporter/grpc/common"
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

type SegmentObject struct {
	TraceSegmentId       *common.UniqueId `protobuf:"bytes,1,opt,name=traceSegmentId,proto3" json:"traceSegmentId,omitempty"`
	Spans                []*SpanObjectV2  `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	ServiceId            int32            `protobuf:"varint,3,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	ServiceInstanceId    int32            `protobuf:"varint,4,opt,name=serviceInstanceId,proto3" json:"serviceInstanceId,omitempty"`
	IsSizeLimited        bool             `protobuf:"varint,5,opt,name=isSizeLimited,proto3" json:"isSizeLimited,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SegmentObject) Reset()         { *m = SegmentObject{} }
func (m *SegmentObject) String() string { return proto.CompactTextString(m) }
func (*SegmentObject) ProtoMessage()    {}
func (*SegmentObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{0}
}

func (m *SegmentObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentObject.Unmarshal(m, b)
}
func (m *SegmentObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentObject.Marshal(b, m, deterministic)
}
func (m *SegmentObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentObject.Merge(m, src)
}
func (m *SegmentObject) XXX_Size() int {
	return xxx_messageInfo_SegmentObject.Size(m)
}
func (m *SegmentObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentObject.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentObject proto.InternalMessageInfo

func (m *SegmentObject) GetTraceSegmentId() *common.UniqueId {
	if m != nil {
		return m.TraceSegmentId
	}
	return nil
}

func (m *SegmentObject) GetSpans() []*SpanObjectV2 {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *SegmentObject) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *SegmentObject) GetServiceInstanceId() int32 {
	if m != nil {
		return m.ServiceInstanceId
	}
	return 0
}

func (m *SegmentObject) GetIsSizeLimited() bool {
	if m != nil {
		return m.IsSizeLimited
	}
	return false
}

type SegmentReference struct {
	RefType                 common.RefType   `protobuf:"varint,1,opt,name=refType,proto3,enum=RefType" json:"refType,omitempty"`
	ParentTraceSegmentId    *common.UniqueId `protobuf:"bytes,2,opt,name=parentTraceSegmentId,proto3" json:"parentTraceSegmentId,omitempty"`
	ParentSpanId            int32            `protobuf:"varint,3,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	ParentServiceInstanceId int32            `protobuf:"varint,4,opt,name=parentServiceInstanceId,proto3" json:"parentServiceInstanceId,omitempty"`
	NetworkAddress          string           `protobuf:"bytes,5,opt,name=networkAddress,proto3" json:"networkAddress,omitempty"`
	NetworkAddressId        int32            `protobuf:"varint,6,opt,name=networkAddressId,proto3" json:"networkAddressId,omitempty"`
	EntryServiceInstanceId  int32            `protobuf:"varint,7,opt,name=entryServiceInstanceId,proto3" json:"entryServiceInstanceId,omitempty"`
	EntryEndpoint           string           `protobuf:"bytes,8,opt,name=entryEndpoint,proto3" json:"entryEndpoint,omitempty"`
	EntryEndpointId         int32            `protobuf:"varint,9,opt,name=entryEndpointId,proto3" json:"entryEndpointId,omitempty"`
	ParentEndpoint          string           `protobuf:"bytes,10,opt,name=parentEndpoint,proto3" json:"parentEndpoint,omitempty"`
	ParentEndpointId        int32            `protobuf:"varint,11,opt,name=parentEndpointId,proto3" json:"parentEndpointId,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}         `json:"-"`
	XXX_unrecognized        []byte           `json:"-"`
	XXX_sizecache           int32            `json:"-"`
}

func (m *SegmentReference) Reset()         { *m = SegmentReference{} }
func (m *SegmentReference) String() string { return proto.CompactTextString(m) }
func (*SegmentReference) ProtoMessage()    {}
func (*SegmentReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{1}
}

func (m *SegmentReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentReference.Unmarshal(m, b)
}
func (m *SegmentReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentReference.Marshal(b, m, deterministic)
}
func (m *SegmentReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentReference.Merge(m, src)
}
func (m *SegmentReference) XXX_Size() int {
	return xxx_messageInfo_SegmentReference.Size(m)
}
func (m *SegmentReference) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentReference.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentReference proto.InternalMessageInfo

func (m *SegmentReference) GetRefType() common.RefType {
	if m != nil {
		return m.RefType
	}
	return common.RefType_CrossProcess
}

func (m *SegmentReference) GetParentTraceSegmentId() *common.UniqueId {
	if m != nil {
		return m.ParentTraceSegmentId
	}
	return nil
}

func (m *SegmentReference) GetParentSpanId() int32 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *SegmentReference) GetParentServiceInstanceId() int32 {
	if m != nil {
		return m.ParentServiceInstanceId
	}
	return 0
}

func (m *SegmentReference) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

func (m *SegmentReference) GetNetworkAddressId() int32 {
	if m != nil {
		return m.NetworkAddressId
	}
	return 0
}

func (m *SegmentReference) GetEntryServiceInstanceId() int32 {
	if m != nil {
		return m.EntryServiceInstanceId
	}
	return 0
}

func (m *SegmentReference) GetEntryEndpoint() string {
	if m != nil {
		return m.EntryEndpoint
	}
	return ""
}

func (m *SegmentReference) GetEntryEndpointId() int32 {
	if m != nil {
		return m.EntryEndpointId
	}
	return 0
}

func (m *SegmentReference) GetParentEndpoint() string {
	if m != nil {
		return m.ParentEndpoint
	}
	return ""
}

func (m *SegmentReference) GetParentEndpointId() int32 {
	if m != nil {
		return m.ParentEndpointId
	}
	return 0
}

type SpanObjectV2 struct {
	SpanId               int32                        `protobuf:"varint,1,opt,name=spanId,proto3" json:"spanId,omitempty"`
	ParentSpanId         int32                        `protobuf:"varint,2,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	StartTime            int64                        `protobuf:"varint,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64                        `protobuf:"varint,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Refs                 []*SegmentReference          `protobuf:"bytes,5,rep,name=refs,proto3" json:"refs,omitempty"`
	OperationNameId      int32                        `protobuf:"varint,6,opt,name=operationNameId,proto3" json:"operationNameId,omitempty"`
	OperationName        string                       `protobuf:"bytes,7,opt,name=operationName,proto3" json:"operationName,omitempty"`
	PeerId               int32                        `protobuf:"varint,8,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Peer                 string                       `protobuf:"bytes,9,opt,name=peer,proto3" json:"peer,omitempty"`
	SpanType             common.SpanType              `protobuf:"varint,10,opt,name=spanType,proto3,enum=SpanType" json:"spanType,omitempty"`
	SpanLayer            common.SpanLayer             `protobuf:"varint,11,opt,name=spanLayer,proto3,enum=SpanLayer" json:"spanLayer,omitempty"`
	ComponentId          int32                        `protobuf:"varint,12,opt,name=componentId,proto3" json:"componentId,omitempty"`
	Component            string                       `protobuf:"bytes,13,opt,name=component,proto3" json:"component,omitempty"`
	IsError              bool                         `protobuf:"varint,14,opt,name=isError,proto3" json:"isError,omitempty"`
	Tags                 []*common.KeyStringValuePair `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty"`
	Logs                 []*Log                       `protobuf:"bytes,16,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SpanObjectV2) Reset()         { *m = SpanObjectV2{} }
func (m *SpanObjectV2) String() string { return proto.CompactTextString(m) }
func (*SpanObjectV2) ProtoMessage()    {}
func (*SpanObjectV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{2}
}

func (m *SpanObjectV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanObjectV2.Unmarshal(m, b)
}
func (m *SpanObjectV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanObjectV2.Marshal(b, m, deterministic)
}
func (m *SpanObjectV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanObjectV2.Merge(m, src)
}
func (m *SpanObjectV2) XXX_Size() int {
	return xxx_messageInfo_SpanObjectV2.Size(m)
}
func (m *SpanObjectV2) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanObjectV2.DiscardUnknown(m)
}

var xxx_messageInfo_SpanObjectV2 proto.InternalMessageInfo

func (m *SpanObjectV2) GetSpanId() int32 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *SpanObjectV2) GetParentSpanId() int32 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *SpanObjectV2) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *SpanObjectV2) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *SpanObjectV2) GetRefs() []*SegmentReference {
	if m != nil {
		return m.Refs
	}
	return nil
}

func (m *SpanObjectV2) GetOperationNameId() int32 {
	if m != nil {
		return m.OperationNameId
	}
	return 0
}

func (m *SpanObjectV2) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *SpanObjectV2) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *SpanObjectV2) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *SpanObjectV2) GetSpanType() common.SpanType {
	if m != nil {
		return m.SpanType
	}
	return common.SpanType_Entry
}

func (m *SpanObjectV2) GetSpanLayer() common.SpanLayer {
	if m != nil {
		return m.SpanLayer
	}
	return common.SpanLayer_Unknown
}

func (m *SpanObjectV2) GetComponentId() int32 {
	if m != nil {
		return m.ComponentId
	}
	return 0
}

func (m *SpanObjectV2) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *SpanObjectV2) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *SpanObjectV2) GetTags() []*common.KeyStringValuePair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SpanObjectV2) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

type Log struct {
	Time                 int64                        `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Data                 []*common.KeyStringValuePair `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{3}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Log) GetData() []*common.KeyStringValuePair {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SegmentObject)(nil), "SegmentObject")
	proto.RegisterType((*SegmentReference)(nil), "SegmentReference")
	proto.RegisterType((*SpanObjectV2)(nil), "SpanObjectV2")
	proto.RegisterType((*Log)(nil), "Log")
}

func init() { proto.RegisterFile("language-agent-v2/trace.proto", fileDescriptor_8124ab206744863a) }

var fileDescriptor_8124ab206744863a = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0xa7, 0x44, 0xf9, 0xd0, 0x49, 0xe2, 0xa6, 0xec, 0xd0, 0xa9, 0xc1, 0x06, 0x18, 0xde,
	0xba, 0x19, 0xc5, 0x42, 0x63, 0x2a, 0x30, 0xec, 0xa6, 0x17, 0xeb, 0x50, 0x0c, 0xc6, 0x82, 0x2e,
	0xa0, 0xd3, 0x0e, 0xd8, 0x1d, 0x23, 0x9d, 0xa8, 0x9a, 0x25, 0x52, 0x23, 0xe9, 0x14, 0xde, 0xed,
	0x1e, 0x21, 0x6f, 0xb1, 0x37, 0xda, 0xdb, 0x14, 0x3c, 0x92, 0x3f, 0x64, 0x27, 0x57, 0xe6, 0xf9,
	0xfd, 0x8f, 0x29, 0x9e, 0xf3, 0xd7, 0xa1, 0xe0, 0xab, 0x52, 0xaa, 0x7c, 0x26, 0x73, 0x3c, 0x97,
	0x39, 0x2a, 0x77, 0x7e, 0x9b, 0x8c, 0x9c, 0x91, 0x29, 0xf2, 0xda, 0x68, 0xa7, 0xcf, 0x9e, 0xa4,
	0xba, 0xaa, 0xb4, 0x1a, 0x35, 0x3f, 0x2d, 0x7c, 0xd6, 0x42, 0x4a, 0x3c, 0x5f, 0x97, 0x06, 0xff,
	0x07, 0x70, 0x32, 0xc1, 0xbc, 0x42, 0xe5, 0x7e, 0xbf, 0xfe, 0x0b, 0x53, 0xc7, 0x7e, 0x80, 0x1e,
	0xe5, 0xb5, 0x74, 0x9c, 0xc5, 0x41, 0x3f, 0x18, 0x1e, 0x25, 0x11, 0x7f, 0xa7, 0x8a, 0xbf, 0x67,
	0x38, 0xce, 0xc4, 0x46, 0x02, 0xfb, 0x1a, 0xf6, 0x6c, 0x2d, 0x95, 0x8d, 0x77, 0xfa, 0xbb, 0xc3,
	0xa3, 0xe4, 0x84, 0x4f, 0x6a, 0xa9, 0x9a, 0xed, 0xde, 0x27, 0xa2, 0xd1, 0xd8, 0x97, 0x10, 0x59,
	0x34, 0xb7, 0x45, 0x8a, 0xe3, 0x2c, 0xde, 0xed, 0x07, 0xc3, 0x3d, 0xb1, 0x02, 0xec, 0x7b, 0x78,
	0xbc, 0x08, 0x94, 0x75, 0x52, 0x51, 0x56, 0x48, 0x59, 0xdb, 0x02, 0xfb, 0x06, 0x4e, 0x0a, 0x3b,
	0x29, 0xfe, 0xc1, 0x8b, 0xa2, 0x2a, 0x1c, 0x66, 0xf1, 0x5e, 0x3f, 0x18, 0x1e, 0x8a, 0x2e, 0x1c,
	0xfc, 0x1b, 0xc2, 0x69, 0x7b, 0x48, 0x81, 0x37, 0x68, 0x50, 0xa5, 0xc8, 0x06, 0x70, 0x60, 0xf0,
	0xe6, 0x6a, 0x5e, 0x23, 0xd5, 0xd5, 0x4b, 0x0e, 0xb9, 0x68, 0x62, 0xb1, 0x10, 0xd8, 0x2b, 0xf8,
	0xbc, 0x96, 0x06, 0x95, 0xbb, 0xea, 0x36, 0x62, 0x67, 0xb3, 0x11, 0xf7, 0xa6, 0xb1, 0x01, 0x1c,
	0x37, 0xdc, 0xb7, 0x61, 0x59, 0x6c, 0x87, 0xb1, 0x9f, 0xe0, 0x8b, 0x36, 0x7e, 0xa0, 0xea, 0x87,
	0x64, 0xf6, 0x2d, 0xf4, 0x14, 0xba, 0x8f, 0xda, 0x4c, 0x7f, 0xce, 0x32, 0x83, 0xd6, 0x52, 0xf1,
	0x91, 0xd8, 0xa0, 0xec, 0x05, 0x9c, 0x76, 0xc9, 0x38, 0x8b, 0xf7, 0x69, 0xeb, 0x2d, 0xce, 0x7e,
	0x84, 0xa7, 0xa8, 0x9c, 0x99, 0x6f, 0x1f, 0xe6, 0x80, 0xfe, 0xf1, 0x80, 0xea, 0x7d, 0x20, 0xe5,
	0x8d, 0xca, 0x6a, 0x5d, 0x28, 0x17, 0x1f, 0xd2, 0x51, 0xba, 0x90, 0x0d, 0xe1, 0x51, 0x07, 0x8c,
	0xb3, 0x38, 0xa2, 0x6d, 0x37, 0xb1, 0xaf, 0xad, 0x29, 0x7b, 0xb9, 0x21, 0x34, 0xb5, 0x75, 0xa9,
	0xaf, 0xad, 0x4b, 0xc6, 0x59, 0x7c, 0xd4, 0xd4, 0xb6, 0xc9, 0x07, 0x77, 0x21, 0x1c, 0xaf, 0xbf,
	0x8f, 0xec, 0x29, 0xec, 0xdb, 0xc6, 0x98, 0x80, 0xfe, 0xd2, 0x46, 0x5b, 0xb6, 0xed, 0xdc, 0x63,
	0x9b, 0x7f, 0x89, 0x9d, 0x34, 0xee, 0xaa, 0xa8, 0x90, 0x7c, 0xdd, 0x15, 0x2b, 0xc0, 0x62, 0x38,
	0x40, 0x95, 0x91, 0x16, 0x92, 0xb6, 0x08, 0xd9, 0x73, 0x08, 0x0d, 0xde, 0x78, 0xab, 0xfc, 0x80,
	0x3c, 0xe6, 0x9b, 0xaf, 0xa5, 0x20, 0xd9, 0x77, 0x4a, 0xd7, 0x68, 0xa4, 0x2b, 0xb4, 0x7a, 0x2b,
	0x2b, 0x5c, 0x5a, 0xb6, 0x89, 0x7d, 0xe7, 0x3b, 0x88, 0x8c, 0x8a, 0x44, 0x17, 0xfa, 0x52, 0x6b,
	0x44, 0x33, 0xce, 0xc8, 0x98, 0x3d, 0xd1, 0x46, 0x8c, 0x41, 0xe8, 0x57, 0x64, 0x43, 0x24, 0x68,
	0xcd, 0x9e, 0xc3, 0xa1, 0x6f, 0x04, 0x4d, 0x06, 0xd0, 0x64, 0x44, 0x34, 0xc7, 0x34, 0x1a, 0x4b,
	0x89, 0x0d, 0x21, 0xf2, 0xeb, 0x0b, 0x39, 0x47, 0x43, 0x3d, 0xef, 0x25, 0x40, 0x79, 0x44, 0xc4,
	0x4a, 0x64, 0x7d, 0x38, 0x4a, 0x75, 0x55, 0x6b, 0xd5, 0x0c, 0xcf, 0x31, 0x9d, 0x60, 0x1d, 0xf9,
	0x6e, 0x2e, 0xc3, 0xf8, 0x84, 0xce, 0xb2, 0x02, 0xbe, 0x9b, 0x85, 0x7d, 0x63, 0x8c, 0x36, 0x71,
	0x8f, 0xc6, 0x7b, 0x11, 0xb2, 0xef, 0x20, 0x74, 0x32, 0xb7, 0xf1, 0x23, 0xea, 0xe6, 0x13, 0xfe,
	0x1b, 0xce, 0x27, 0xce, 0x14, 0x2a, 0x7f, 0x2f, 0xcb, 0x19, 0x5e, 0xca, 0xc2, 0x08, 0x4a, 0x60,
	0x31, 0x84, 0xa5, 0xce, 0x6d, 0x7c, 0x4a, 0x89, 0x21, 0xbf, 0xd0, 0xb9, 0x20, 0x32, 0x78, 0x0d,
	0xbb, 0x17, 0x3a, 0xf7, 0x8d, 0x70, 0xde, 0xae, 0x80, 0xec, 0xa2, 0xb5, 0xdf, 0x3d, 0x93, 0x4e,
	0xb6, 0x97, 0xd9, 0xfd, 0xbb, 0xfb, 0x84, 0xe4, 0x57, 0x78, 0xb6, 0x3e, 0xf9, 0x02, 0x6b, 0x6d,
	0x16, 0x03, 0xcb, 0x5e, 0xc0, 0x41, 0xaa, 0xcb, 0xd2, 0xdf, 0xa8, 0xa7, 0xfc, 0x5d, 0x6d, 0x9d,
	0x41, 0x59, 0xb5, 0x99, 0x67, 0x11, 0xff, 0x45, 0x57, 0x95, 0x54, 0x99, 0x1d, 0x7c, 0x36, 0x0c,
	0x5e, 0xdf, 0x05, 0xf0, 0x52, 0x9b, 0x9c, 0xcb, 0x5a, 0xa6, 0x1f, 0x90, 0xdb, 0xe9, 0xfc, 0xa3,
	0x2c, 0xa7, 0x85, 0xf2, 0xa4, 0xe2, 0xed, 0xb4, 0xf2, 0xc5, 0xc5, 0xcf, 0xe9, 0xe2, 0xe7, 0xb7,
	0xc9, 0x65, 0xf0, 0xe7, 0xab, 0xbc, 0x70, 0x1f, 0x66, 0xd7, 0x3c, 0xd5, 0xd5, 0xc8, 0xa1, 0x33,
	0xd2, 0x61, 0x29, 0xaf, 0xed, 0x28, 0xd7, 0x89, 0x9d, 0xce, 0x47, 0x86, 0x8e, 0x84, 0x66, 0x94,
	0x9b, 0x3a, 0x1d, 0x6d, 0x7d, 0x39, 0xfe, 0xdb, 0x39, 0x9b, 0x4c, 0xe7, 0x7f, 0xb4, 0xcf, 0x7a,
	0xdb, 0x3c, 0xe7, 0xd2, 0x7f, 0x17, 0x52, 0x5d, 0x5e, 0xef, 0xd3, 0x17, 0xe2, 0xe5, 0xa7, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x09, 0x88, 0x7b, 0xc8, 0x72, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceSegmentReportServiceClient is the client API for TraceSegmentReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceSegmentReportServiceClient interface {
	Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentReportService_CollectClient, error)
}

type traceSegmentReportServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceSegmentReportServiceClient(cc *grpc.ClientConn) TraceSegmentReportServiceClient {
	return &traceSegmentReportServiceClient{cc}
}

func (c *traceSegmentReportServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentReportService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceSegmentReportService_serviceDesc.Streams[0], "/TraceSegmentReportService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceSegmentReportServiceCollectClient{stream}
	return x, nil
}

type TraceSegmentReportService_CollectClient interface {
	Send(*common.UpstreamSegment) error
	CloseAndRecv() (*common.Commands, error)
	grpc.ClientStream
}

type traceSegmentReportServiceCollectClient struct {
	grpc.ClientStream
}

func (x *traceSegmentReportServiceCollectClient) Send(m *common.UpstreamSegment) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceSegmentReportServiceCollectClient) CloseAndRecv() (*common.Commands, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(common.Commands)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceSegmentReportServiceServer is the server API for TraceSegmentReportService service.
type TraceSegmentReportServiceServer interface {
	Collect(TraceSegmentReportService_CollectServer) error
}

// UnimplementedTraceSegmentReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceSegmentReportServiceServer struct {
}

func (*UnimplementedTraceSegmentReportServiceServer) Collect(srv TraceSegmentReportService_CollectServer) error {
	return status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterTraceSegmentReportServiceServer(s *grpc.Server, srv TraceSegmentReportServiceServer) {
	s.RegisterService(&_TraceSegmentReportService_serviceDesc, srv)
}

func _TraceSegmentReportService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceSegmentReportServiceServer).Collect(&traceSegmentReportServiceCollectServer{stream})
}

type TraceSegmentReportService_CollectServer interface {
	SendAndClose(*common.Commands) error
	Recv() (*common.UpstreamSegment, error)
	grpc.ServerStream
}

type traceSegmentReportServiceCollectServer struct {
	grpc.ServerStream
}

func (x *traceSegmentReportServiceCollectServer) SendAndClose(m *common.Commands) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceSegmentReportServiceCollectServer) Recv() (*common.UpstreamSegment, error) {
	m := new(common.UpstreamSegment)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceSegmentReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TraceSegmentReportService",
	HandlerType: (*TraceSegmentReportServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _TraceSegmentReportService_Collect_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "language-agent-v2/trace.proto",
}
