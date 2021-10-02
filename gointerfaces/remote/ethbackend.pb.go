// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: remote/ethbackend.proto

package remote

import (
	types "github.com/ledgerwatch/erigon-lib/gointerfaces/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event int32

const (
	Event_HEADER        Event = 0
	Event_PENDING_LOGS  Event = 1
	Event_PENDING_BLOCK Event = 2
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "HEADER",
		1: "PENDING_LOGS",
		2: "PENDING_BLOCK",
	}
	Event_value = map[string]int32{
		"HEADER":        0,
		"PENDING_LOGS":  1,
		"PENDING_BLOCK": 2,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_remote_ethbackend_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_remote_ethbackend_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{0}
}

type EtherbaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EtherbaseRequest) Reset() {
	*x = EtherbaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtherbaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtherbaseRequest) ProtoMessage() {}

func (x *EtherbaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtherbaseRequest.ProtoReflect.Descriptor instead.
func (*EtherbaseRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{0}
}

type EtherbaseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *types.H160 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *EtherbaseReply) Reset() {
	*x = EtherbaseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtherbaseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtherbaseReply) ProtoMessage() {}

func (x *EtherbaseReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtherbaseReply.ProtoReflect.Descriptor instead.
func (*EtherbaseReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{1}
}

func (x *EtherbaseReply) GetAddress() *types.H160 {
	if x != nil {
		return x.Address
	}
	return nil
}

type NetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetVersionRequest) Reset() {
	*x = NetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetVersionRequest) ProtoMessage() {}

func (x *NetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetVersionRequest.ProtoReflect.Descriptor instead.
func (*NetVersionRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{2}
}

type NetVersionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NetVersionReply) Reset() {
	*x = NetVersionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetVersionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetVersionReply) ProtoMessage() {}

func (x *NetVersionReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetVersionReply.ProtoReflect.Descriptor instead.
func (*NetVersionReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{3}
}

func (x *NetVersionReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NetPeerCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetPeerCountRequest) Reset() {
	*x = NetPeerCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetPeerCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetPeerCountRequest) ProtoMessage() {}

func (x *NetPeerCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetPeerCountRequest.ProtoReflect.Descriptor instead.
func (*NetPeerCountRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{4}
}

type NetPeerCountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *NetPeerCountReply) Reset() {
	*x = NetPeerCountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetPeerCountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetPeerCountReply) ProtoMessage() {}

func (x *NetPeerCountReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetPeerCountReply.ProtoReflect.Descriptor instead.
func (*NetPeerCountReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{5}
}

func (x *NetPeerCountReply) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ProtocolVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProtocolVersionRequest) Reset() {
	*x = ProtocolVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolVersionRequest) ProtoMessage() {}

func (x *ProtocolVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolVersionRequest.ProtoReflect.Descriptor instead.
func (*ProtocolVersionRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{6}
}

type ProtocolVersionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProtocolVersionReply) Reset() {
	*x = ProtocolVersionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolVersionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolVersionReply) ProtoMessage() {}

func (x *ProtocolVersionReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolVersionReply.ProtoReflect.Descriptor instead.
func (*ProtocolVersionReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{7}
}

func (x *ProtocolVersionReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ClientVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientVersionRequest) Reset() {
	*x = ClientVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientVersionRequest) ProtoMessage() {}

func (x *ClientVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientVersionRequest.ProtoReflect.Descriptor instead.
func (*ClientVersionRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{8}
}

type ClientVersionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
}

func (x *ClientVersionReply) Reset() {
	*x = ClientVersionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientVersionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientVersionReply) ProtoMessage() {}

func (x *ClientVersionReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientVersionReply.ProtoReflect.Descriptor instead.
func (*ClientVersionReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{9}
}

func (x *ClientVersionReply) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Event `protobuf:"varint,1,opt,name=type,proto3,enum=remote.Event" json:"type,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{10}
}

func (x *SubscribeRequest) GetType() Event {
	if x != nil {
		return x.Type
	}
	return Event_HEADER
}

type SubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Event  `protobuf:"varint,1,opt,name=type,proto3,enum=remote.Event" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` //  serialized data
}

func (x *SubscribeReply) Reset() {
	*x = SubscribeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReply) ProtoMessage() {}

func (x *SubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReply.ProtoReflect.Descriptor instead.
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{11}
}

func (x *SubscribeReply) GetType() Event {
	if x != nil {
		return x.Type
	}
	return Event_HEADER
}

func (x *SubscribeReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_remote_ethbackend_proto protoreflect.FileDescriptor

var file_remote_ethbackend_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x65, 0x74, 0x68, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x12, 0x0a, 0x10, 0x45, 0x74, 0x68, 0x65, 0x72, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x0e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x48, 0x31, 0x36, 0x30, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x13,
	0x0a, 0x11, 0x4e, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x4e, 0x65, 0x74, 0x50, 0x65, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a,
	0x11, 0x4e, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x30, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x0e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x2a, 0x38, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0a, 0x0a,
	0x06, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x02, 0x32, 0xea,
	0x03, 0x0a, 0x0a, 0x45, 0x54, 0x48, 0x42, 0x41, 0x43, 0x4b, 0x45, 0x4e, 0x44, 0x12, 0x3d, 0x0a,
	0x09, 0x45, 0x74, 0x68, 0x65, 0x72, 0x62, 0x61, 0x73, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x45, 0x74,
	0x68, 0x65, 0x72, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x0a,
	0x4e, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x4e,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46,
	0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f,
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x49, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3f, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x3b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_ethbackend_proto_rawDescOnce sync.Once
	file_remote_ethbackend_proto_rawDescData = file_remote_ethbackend_proto_rawDesc
)

func file_remote_ethbackend_proto_rawDescGZIP() []byte {
	file_remote_ethbackend_proto_rawDescOnce.Do(func() {
		file_remote_ethbackend_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_ethbackend_proto_rawDescData)
	})
	return file_remote_ethbackend_proto_rawDescData
}

var file_remote_ethbackend_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_remote_ethbackend_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_remote_ethbackend_proto_goTypes = []interface{}{
	(Event)(0),                     // 0: remote.Event
	(*EtherbaseRequest)(nil),       // 1: remote.EtherbaseRequest
	(*EtherbaseReply)(nil),         // 2: remote.EtherbaseReply
	(*NetVersionRequest)(nil),      // 3: remote.NetVersionRequest
	(*NetVersionReply)(nil),        // 4: remote.NetVersionReply
	(*NetPeerCountRequest)(nil),    // 5: remote.NetPeerCountRequest
	(*NetPeerCountReply)(nil),      // 6: remote.NetPeerCountReply
	(*ProtocolVersionRequest)(nil), // 7: remote.ProtocolVersionRequest
	(*ProtocolVersionReply)(nil),   // 8: remote.ProtocolVersionReply
	(*ClientVersionRequest)(nil),   // 9: remote.ClientVersionRequest
	(*ClientVersionReply)(nil),     // 10: remote.ClientVersionReply
	(*SubscribeRequest)(nil),       // 11: remote.SubscribeRequest
	(*SubscribeReply)(nil),         // 12: remote.SubscribeReply
	(*types.H160)(nil),             // 13: types.H160
	(*emptypb.Empty)(nil),          // 14: google.protobuf.Empty
	(*types.VersionReply)(nil),     // 15: types.VersionReply
}
var file_remote_ethbackend_proto_depIdxs = []int32{
	13, // 0: remote.EtherbaseReply.address:type_name -> types.H160
	0,  // 1: remote.SubscribeRequest.type:type_name -> remote.Event
	0,  // 2: remote.SubscribeReply.type:type_name -> remote.Event
	1,  // 3: remote.ETHBACKEND.Etherbase:input_type -> remote.EtherbaseRequest
	3,  // 4: remote.ETHBACKEND.NetVersion:input_type -> remote.NetVersionRequest
	5,  // 5: remote.ETHBACKEND.NetPeerCount:input_type -> remote.NetPeerCountRequest
	14, // 6: remote.ETHBACKEND.Version:input_type -> google.protobuf.Empty
	7,  // 7: remote.ETHBACKEND.ProtocolVersion:input_type -> remote.ProtocolVersionRequest
	9,  // 8: remote.ETHBACKEND.ClientVersion:input_type -> remote.ClientVersionRequest
	11, // 9: remote.ETHBACKEND.Subscribe:input_type -> remote.SubscribeRequest
	2,  // 10: remote.ETHBACKEND.Etherbase:output_type -> remote.EtherbaseReply
	4,  // 11: remote.ETHBACKEND.NetVersion:output_type -> remote.NetVersionReply
	6,  // 12: remote.ETHBACKEND.NetPeerCount:output_type -> remote.NetPeerCountReply
	15, // 13: remote.ETHBACKEND.Version:output_type -> types.VersionReply
	8,  // 14: remote.ETHBACKEND.ProtocolVersion:output_type -> remote.ProtocolVersionReply
	10, // 15: remote.ETHBACKEND.ClientVersion:output_type -> remote.ClientVersionReply
	12, // 16: remote.ETHBACKEND.Subscribe:output_type -> remote.SubscribeReply
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_remote_ethbackend_proto_init() }
func file_remote_ethbackend_proto_init() {
	if File_remote_ethbackend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_ethbackend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EtherbaseRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EtherbaseReply); i {
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
		file_remote_ethbackend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetVersionRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetVersionReply); i {
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
		file_remote_ethbackend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetPeerCountRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetPeerCountReply); i {
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
		file_remote_ethbackend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolVersionRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolVersionReply); i {
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
		file_remote_ethbackend_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientVersionRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientVersionReply); i {
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
		file_remote_ethbackend_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeReply); i {
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
			RawDescriptor: file_remote_ethbackend_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remote_ethbackend_proto_goTypes,
		DependencyIndexes: file_remote_ethbackend_proto_depIdxs,
		EnumInfos:         file_remote_ethbackend_proto_enumTypes,
		MessageInfos:      file_remote_ethbackend_proto_msgTypes,
	}.Build()
	File_remote_ethbackend_proto = out.File
	file_remote_ethbackend_proto_rawDesc = nil
	file_remote_ethbackend_proto_goTypes = nil
	file_remote_ethbackend_proto_depIdxs = nil
}
