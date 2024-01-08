// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: capture.proto

package openbytes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Capture struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iface  string `protobuf:"bytes,1,opt,name=iface,proto3" json:"iface,omitempty"`
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *Capture) Reset() {
	*x = Capture{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capture_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capture) ProtoMessage() {}

func (x *Capture) ProtoReflect() protoreflect.Message {
	mi := &file_capture_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capture.ProtoReflect.Descriptor instead.
func (*Capture) Descriptor() ([]byte, []int) {
	return file_capture_proto_rawDescGZIP(), []int{0}
}

func (x *Capture) GetIface() string {
	if x != nil {
		return x.Iface
	}
	return ""
}

func (x *Capture) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string type = 1;
	Protocol string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Src      *Address `protobuf:"bytes,3,opt,name=src,proto3" json:"src,omitempty"`
	Dst      *Address `protobuf:"bytes,4,opt,name=dst,proto3" json:"dst,omitempty"`
	Layers   string   `protobuf:"bytes,5,opt,name=layers,proto3" json:"layers,omitempty"`
	Payload  []byte   `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	Data     string   `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capture_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_capture_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_capture_proto_rawDescGZIP(), []int{1}
}

func (x *Packet) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Packet) GetSrc() *Address {
	if x != nil {
		return x.Src
	}
	return nil
}

func (x *Packet) GetDst() *Address {
	if x != nil {
		return x.Dst
	}
	return nil
}

func (x *Packet) GetLayers() string {
	if x != nil {
		return x.Layers
	}
	return ""
}

func (x *Packet) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Packet) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Mac  string `protobuf:"bytes,3,opt,name=mac,proto3" json:"mac,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capture_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_capture_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_capture_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Address) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Address) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label      string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	TcpCount   int32  `protobuf:"varint,2,opt,name=tcp_count,json=tcpCount,proto3" json:"tcp_count,omitempty"`
	UdpCount   int32  `protobuf:"varint,3,opt,name=udp_count,json=udpCount,proto3" json:"udp_count,omitempty"`
	OtherCount int32  `protobuf:"varint,4,opt,name=other_count,json=otherCount,proto3" json:"other_count,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capture_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_capture_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_capture_proto_rawDescGZIP(), []int{3}
}

func (x *Point) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Point) GetTcpCount() int32 {
	if x != nil {
		return x.TcpCount
	}
	return 0
}

func (x *Point) GetUdpCount() int32 {
	if x != nil {
		return x.UdpCount
	}
	return 0
}

func (x *Point) GetOtherCount() int32 {
	if x != nil {
		return x.OtherCount
	}
	return 0
}

type CopyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iface        string                     `protobuf:"bytes,1,opt,name=iface,proto3" json:"iface,omitempty"`
	Destinations []*CopyRequest_Destination `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
}

func (x *CopyRequest) Reset() {
	*x = CopyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capture_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyRequest) ProtoMessage() {}

func (x *CopyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_capture_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyRequest.ProtoReflect.Descriptor instead.
func (*CopyRequest) Descriptor() ([]byte, []int) {
	return file_capture_proto_rawDescGZIP(), []int{4}
}

func (x *CopyRequest) GetIface() string {
	if x != nil {
		return x.Iface
	}
	return ""
}

func (x *CopyRequest) GetDestinations() []*CopyRequest_Destination {
	if x != nil {
		return x.Destinations
	}
	return nil
}

type CopyRequest_Destination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *CopyRequest_Destination) Reset() {
	*x = CopyRequest_Destination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capture_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyRequest_Destination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyRequest_Destination) ProtoMessage() {}

func (x *CopyRequest_Destination) ProtoReflect() protoreflect.Message {
	mi := &file_capture_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyRequest_Destination.ProtoReflect.Descriptor instead.
func (*CopyRequest_Destination) Descriptor() ([]byte, []int) {
	return file_capture_proto_rawDescGZIP(), []int{4, 0}
}

func (x *CopyRequest_Destination) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CopyRequest_Destination) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

var File_capture_proto protoreflect.FileDescriptor

var file_capture_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xb6,
	0x01, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x24, 0x0a, 0x03, 0x64,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x03, 0x64, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3f, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x22, 0x78, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x63, 0x70, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x63, 0x70, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x64, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x64, 0x70, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x31, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x32, 0xf2, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x11, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x12,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x1a, 0x10, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x04, 0x43, 0x6f, 0x70, 0x79,
	0x12, 0x16, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x70,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x3b, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_capture_proto_rawDescOnce sync.Once
	file_capture_proto_rawDescData = file_capture_proto_rawDesc
)

func file_capture_proto_rawDescGZIP() []byte {
	file_capture_proto_rawDescOnce.Do(func() {
		file_capture_proto_rawDescData = protoimpl.X.CompressGZIP(file_capture_proto_rawDescData)
	})
	return file_capture_proto_rawDescData
}

var file_capture_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_capture_proto_goTypes = []interface{}{
	(*Capture)(nil),                 // 0: openbytes.Capture
	(*Packet)(nil),                  // 1: openbytes.Packet
	(*Address)(nil),                 // 2: openbytes.Address
	(*Point)(nil),                   // 3: openbytes.Point
	(*CopyRequest)(nil),             // 4: openbytes.CopyRequest
	(*CopyRequest_Destination)(nil), // 5: openbytes.CopyRequest.Destination
	(*emptypb.Empty)(nil),           // 6: google.protobuf.Empty
	(*structpb.ListValue)(nil),      // 7: google.protobuf.ListValue
	(*wrapperspb.StringValue)(nil),  // 8: google.protobuf.StringValue
}
var file_capture_proto_depIdxs = []int32{
	2, // 0: openbytes.Packet.src:type_name -> openbytes.Address
	2, // 1: openbytes.Packet.dst:type_name -> openbytes.Address
	5, // 2: openbytes.CopyRequest.destinations:type_name -> openbytes.CopyRequest.Destination
	6, // 3: openbytes.Captures.Device:input_type -> google.protobuf.Empty
	0, // 4: openbytes.Captures.List:input_type -> openbytes.Capture
	0, // 5: openbytes.Captures.Traffic:input_type -> openbytes.Capture
	4, // 6: openbytes.Captures.Copy:input_type -> openbytes.CopyRequest
	7, // 7: openbytes.Captures.Device:output_type -> google.protobuf.ListValue
	1, // 8: openbytes.Captures.List:output_type -> openbytes.Packet
	3, // 9: openbytes.Captures.Traffic:output_type -> openbytes.Point
	8, // 10: openbytes.Captures.Copy:output_type -> google.protobuf.StringValue
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_capture_proto_init() }
func file_capture_proto_init() {
	if File_capture_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_capture_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capture); i {
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
		file_capture_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
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
		file_capture_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_capture_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_capture_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyRequest); i {
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
		file_capture_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyRequest_Destination); i {
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
			RawDescriptor: file_capture_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_capture_proto_goTypes,
		DependencyIndexes: file_capture_proto_depIdxs,
		MessageInfos:      file_capture_proto_msgTypes,
	}.Build()
	File_capture_proto = out.File
	file_capture_proto_rawDesc = nil
	file_capture_proto_goTypes = nil
	file_capture_proto_depIdxs = nil
}
