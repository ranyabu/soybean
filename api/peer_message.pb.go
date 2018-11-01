// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer_message.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PeerMessage struct {
	Type                 int32             `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Meta                 map[string][]byte `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PeerMessage) Reset()         { *m = PeerMessage{} }
func (m *PeerMessage) String() string { return proto.CompactTextString(m) }
func (*PeerMessage) ProtoMessage()    {}
func (*PeerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c9f06aaec5d2f7d, []int{0}
}

func (m *PeerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerMessage.Unmarshal(m, b)
}
func (m *PeerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerMessage.Marshal(b, m, deterministic)
}
func (m *PeerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerMessage.Merge(m, src)
}
func (m *PeerMessage) XXX_Size() int {
	return xxx_messageInfo_PeerMessage.Size(m)
}
func (m *PeerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PeerMessage proto.InternalMessageInfo

func (m *PeerMessage) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *PeerMessage) GetMeta() map[string][]byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PeerMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerMessage)(nil), "types.PeerMessage")
	proto.RegisterMapType((map[string][]byte)(nil), "types.PeerMessage.MetaEntry")
}

func init() { proto.RegisterFile("peer_message.proto", fileDescriptor_7c9f06aaec5d2f7d) }

var fileDescriptor_7c9f06aaec5d2f7d = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x48, 0x4d, 0x2d,
	0x8a, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x5a, 0xc0, 0xc8, 0xc5, 0x1d, 0x90, 0x9a, 0x5a, 0xe4, 0x0b,
	0x91, 0x14, 0x12, 0xe2, 0x62, 0x01, 0x49, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0xd9,
	0x42, 0x06, 0x5c, 0x2c, 0xb9, 0xa9, 0x25, 0x89, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x32,
	0x7a, 0x60, 0x9d, 0x7a, 0x48, 0xba, 0xf4, 0x7c, 0x53, 0x4b, 0x12, 0x5d, 0xf3, 0x4a, 0x8a, 0x2a,
	0x83, 0xc0, 0x2a, 0x41, 0xa6, 0x24, 0xe5, 0xa7, 0x54, 0x4a, 0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x04,
	0x81, 0xd9, 0x52, 0xe6, 0x5c, 0x9c, 0x70, 0x65, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x60,
	0x5b, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26,
	0xb0, 0x1e, 0x08, 0xc7, 0x8a, 0xc9, 0x82, 0xd1, 0x89, 0x3d, 0x0a, 0xe2, 0xd6, 0x24, 0x36, 0xb0,
	0xcb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xe0, 0x52, 0x30, 0xcf, 0x00, 0x00, 0x00,
}
