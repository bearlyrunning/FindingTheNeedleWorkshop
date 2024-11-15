// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: signal.proto

package signalpb

import (
	normalizedlogpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*Signal_BadDomainFiltered
	//	*Signal_BadDomain
	//	*Signal_DnsTunnel
	//	*Signal_BrowserSubProc
	Event isSignal_Event `protobuf_oneof:"event"`
}

func (x *Signal) Reset() {
	*x = Signal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{0}
}

func (m *Signal) GetEvent() isSignal_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *Signal) GetBadDomainFiltered() *normalizedlogpb.DNS {
	if x, ok := x.GetEvent().(*Signal_BadDomainFiltered); ok {
		return x.BadDomainFiltered
	}
	return nil
}

func (x *Signal) GetBadDomain() *BadDomain {
	if x, ok := x.GetEvent().(*Signal_BadDomain); ok {
		return x.BadDomain
	}
	return nil
}

func (x *Signal) GetDnsTunnel() *DNSTunnel {
	if x, ok := x.GetEvent().(*Signal_DnsTunnel); ok {
		return x.DnsTunnel
	}
	return nil
}

func (x *Signal) GetBrowserSubProc() *BrowserSubProc {
	if x, ok := x.GetEvent().(*Signal_BrowserSubProc); ok {
		return x.BrowserSubProc
	}
	return nil
}

type isSignal_Event interface {
	isSignal_Event()
}

type Signal_BadDomainFiltered struct {
	BadDomainFiltered *normalizedlogpb.DNS `protobuf:"bytes,1,opt,name=bad_domain_filtered,json=badDomainFiltered,proto3,oneof"`
}

type Signal_BadDomain struct {
	BadDomain *BadDomain `protobuf:"bytes,2,opt,name=bad_domain,json=badDomain,proto3,oneof"`
}

type Signal_DnsTunnel struct {
	DnsTunnel *DNSTunnel `protobuf:"bytes,3,opt,name=dns_tunnel,json=dnsTunnel,proto3,oneof"`
}

type Signal_BrowserSubProc struct {
	BrowserSubProc *BrowserSubProc `protobuf:"bytes,4,opt,name=browser_sub_proc,json=browserSubProc,proto3,oneof"`
}

func (*Signal_BadDomainFiltered) isSignal_Event() {}

func (*Signal_BadDomain) isSignal_Event() {}

func (*Signal_DnsTunnel) isSignal_Event() {}

func (*Signal_BrowserSubProc) isSignal_Event() {}

// BadDomain detection event message.
type BadDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp of the first observed DNS query.
	TimestampStart *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp_start,json=timestampStart,proto3" json:"timestamp_start,omitempty"`
	// Timestamp of the last observed DNS query.
	TimestampEnd *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp_end,json=timestampEnd,proto3" json:"timestamp_end,omitempty"`
	// Bad domain.
	BadDomain string `protobuf:"bytes,3,opt,name=bad_domain,json=badDomain,proto3" json:"bad_domain,omitempty"`
	// Source client IP.
	SourceIp string `protobuf:"bytes,4,opt,name=source_ip,json=sourceIp,proto3" json:"source_ip,omitempty"`
	// Hostname.
	Hostname string                 `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`
	DnsLog   []*normalizedlogpb.DNS `protobuf:"bytes,6,rep,name=dns_log,json=dnsLog,proto3" json:"dns_log,omitempty"`
}

func (x *BadDomain) Reset() {
	*x = BadDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadDomain) ProtoMessage() {}

func (x *BadDomain) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadDomain.ProtoReflect.Descriptor instead.
func (*BadDomain) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{1}
}

func (x *BadDomain) GetTimestampStart() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampStart
	}
	return nil
}

func (x *BadDomain) GetTimestampEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampEnd
	}
	return nil
}

func (x *BadDomain) GetBadDomain() string {
	if x != nil {
		return x.BadDomain
	}
	return ""
}

func (x *BadDomain) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

func (x *BadDomain) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *BadDomain) GetDnsLog() []*normalizedlogpb.DNS {
	if x != nil {
		return x.DnsLog
	}
	return nil
}

type DNSTunnel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp of the first observed netflow.
	TimestampStart *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp_start,json=timestampStart,proto3" json:"timestamp_start,omitempty"`
	// Timestamp of the last observed netflow.
	TimestampEnd *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp_end,json=timestampEnd,proto3" json:"timestamp_end,omitempty"`
	// Tunnel IP.
	TunnelIp string `protobuf:"bytes,3,opt,name=tunnel_ip,json=tunnelIp,proto3" json:"tunnel_ip,omitempty"`
	// Source client IP.
	SourceIp string `protobuf:"bytes,4,opt,name=source_ip,json=sourceIp,proto3" json:"source_ip,omitempty"`
	// Hostname.
	Hostname string `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Total bytes in.
	BytesInTotal int64 `protobuf:"varint,6,opt,name=bytes_in_total,json=bytesInTotal,proto3" json:"bytes_in_total,omitempty"`
	// Total bytes out.
	BytesOutTotal int64                      `protobuf:"varint,7,opt,name=bytes_out_total,json=bytesOutTotal,proto3" json:"bytes_out_total,omitempty"`
	NetflowLog    []*normalizedlogpb.Netflow `protobuf:"bytes,8,rep,name=netflow_log,json=netflowLog,proto3" json:"netflow_log,omitempty"`
}

func (x *DNSTunnel) Reset() {
	*x = DNSTunnel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSTunnel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSTunnel) ProtoMessage() {}

func (x *DNSTunnel) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSTunnel.ProtoReflect.Descriptor instead.
func (*DNSTunnel) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{2}
}

func (x *DNSTunnel) GetTimestampStart() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampStart
	}
	return nil
}

func (x *DNSTunnel) GetTimestampEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampEnd
	}
	return nil
}

func (x *DNSTunnel) GetTunnelIp() string {
	if x != nil {
		return x.TunnelIp
	}
	return ""
}

func (x *DNSTunnel) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

func (x *DNSTunnel) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *DNSTunnel) GetBytesInTotal() int64 {
	if x != nil {
		return x.BytesInTotal
	}
	return 0
}

func (x *DNSTunnel) GetBytesOutTotal() int64 {
	if x != nil {
		return x.BytesOutTotal
	}
	return 0
}

func (x *DNSTunnel) GetNetflowLog() []*normalizedlogpb.Netflow {
	if x != nil {
		return x.NetflowLog
	}
	return nil
}

type BrowserSubProc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Matched anomalous log.
	Execution *normalizedlogpb.Execution `protobuf:"bytes,1,opt,name=execution,proto3" json:"execution,omitempty"`
	// IP address of the source host.
	SourceIp string `protobuf:"bytes,2,opt,name=source_ip,json=sourceIp,proto3" json:"source_ip,omitempty"`
}

func (x *BrowserSubProc) Reset() {
	*x = BrowserSubProc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrowserSubProc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowserSubProc) ProtoMessage() {}

func (x *BrowserSubProc) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowserSubProc.ProtoReflect.Descriptor instead.
func (*BrowserSubProc) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{3}
}

func (x *BrowserSubProc) GetExecution() *normalizedlogpb.Execution {
	if x != nil {
		return x.Execution
	}
	return nil
}

func (x *BrowserSubProc) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

var File_signal_proto protoreflect.FileDescriptor

var file_signal_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02,
	0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x44, 0x0a, 0x13, 0x62, 0x61, 0x64, 0x5f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x4e, 0x53, 0x48, 0x00, 0x52, 0x11, 0x62, 0x61, 0x64,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x32,
	0x0a, 0x0a, 0x62, 0x61, 0x64, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x42, 0x61, 0x64, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x62, 0x61, 0x64, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x32, 0x0a, 0x0a, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e,
	0x44, 0x4e, 0x53, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x09, 0x64, 0x6e, 0x73,
	0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x42, 0x0a, 0x10, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65,
	0x72, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x63, 0x48, 0x00, 0x52, 0x0e, 0x62, 0x72, 0x6f, 0x77,
	0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x63, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x96, 0x02, 0x0a, 0x09, 0x42, 0x61, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x43, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x45, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x64, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x64,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2b, 0x0a, 0x07, 0x64, 0x6e, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x6c, 0x6f, 0x67,
	0x2e, 0x44, 0x4e, 0x53, 0x52, 0x06, 0x64, 0x6e, 0x73, 0x4c, 0x6f, 0x67, 0x22, 0xee, 0x02, 0x0a,
	0x09, 0x44, 0x4e, 0x53, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x43, 0x0a, 0x0f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x3f, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x65, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x45, 0x6e, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x70, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x0f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x6c, 0x6f, 0x67, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x6c, 0x6f, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x66, 0x6c, 0x6f,
	0x77, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x22, 0x65, 0x0a,
	0x0e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x63, 0x12,
	0x36, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x6c,
	0x6f, 0x67, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x70, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x2f, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x68, 0x65, 0x4e, 0x65, 0x65, 0x64, 0x6c,
	0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signal_proto_rawDescOnce sync.Once
	file_signal_proto_rawDescData = file_signal_proto_rawDesc
)

func file_signal_proto_rawDescGZIP() []byte {
	file_signal_proto_rawDescOnce.Do(func() {
		file_signal_proto_rawDescData = protoimpl.X.CompressGZIP(file_signal_proto_rawDescData)
	})
	return file_signal_proto_rawDescData
}

var file_signal_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_signal_proto_goTypes = []interface{}{
	(*Signal)(nil),                    // 0: signal.Signal
	(*BadDomain)(nil),                 // 1: signal.BadDomain
	(*DNSTunnel)(nil),                 // 2: signal.DNSTunnel
	(*BrowserSubProc)(nil),            // 3: signal.BrowserSubProc
	(*normalizedlogpb.DNS)(nil),       // 4: normalizedlog.DNS
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
	(*normalizedlogpb.Netflow)(nil),   // 6: normalizedlog.Netflow
	(*normalizedlogpb.Execution)(nil), // 7: normalizedlog.Execution
}
var file_signal_proto_depIdxs = []int32{
	4,  // 0: signal.Signal.bad_domain_filtered:type_name -> normalizedlog.DNS
	1,  // 1: signal.Signal.bad_domain:type_name -> signal.BadDomain
	2,  // 2: signal.Signal.dns_tunnel:type_name -> signal.DNSTunnel
	3,  // 3: signal.Signal.browser_sub_proc:type_name -> signal.BrowserSubProc
	5,  // 4: signal.BadDomain.timestamp_start:type_name -> google.protobuf.Timestamp
	5,  // 5: signal.BadDomain.timestamp_end:type_name -> google.protobuf.Timestamp
	4,  // 6: signal.BadDomain.dns_log:type_name -> normalizedlog.DNS
	5,  // 7: signal.DNSTunnel.timestamp_start:type_name -> google.protobuf.Timestamp
	5,  // 8: signal.DNSTunnel.timestamp_end:type_name -> google.protobuf.Timestamp
	6,  // 9: signal.DNSTunnel.netflow_log:type_name -> normalizedlog.Netflow
	7,  // 10: signal.BrowserSubProc.execution:type_name -> normalizedlog.Execution
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_signal_proto_init() }
func file_signal_proto_init() {
	if File_signal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal); i {
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
		file_signal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadDomain); i {
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
		file_signal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSTunnel); i {
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
		file_signal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrowserSubProc); i {
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
	file_signal_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Signal_BadDomainFiltered)(nil),
		(*Signal_BadDomain)(nil),
		(*Signal_DnsTunnel)(nil),
		(*Signal_BrowserSubProc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_signal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_signal_proto_goTypes,
		DependencyIndexes: file_signal_proto_depIdxs,
		MessageInfos:      file_signal_proto_msgTypes,
	}.Build()
	File_signal_proto = out.File
	file_signal_proto_rawDesc = nil
	file_signal_proto_goTypes = nil
	file_signal_proto_depIdxs = nil
}
