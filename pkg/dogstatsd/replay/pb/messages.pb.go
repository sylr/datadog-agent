// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package pb

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The message contains the payload and the ancillary info
type UnixDogstatsdMsg struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PayloadSize          int32    `protobuf:"varint,2,opt,name=payloadSize,proto3" json:"payloadSize,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	AncillarySize        int32    `protobuf:"varint,4,opt,name=ancillarySize,proto3" json:"ancillarySize,omitempty"`
	Ancillary            []byte   `protobuf:"bytes,5,opt,name=ancillary,proto3" json:"ancillary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnixDogstatsdMsg) Reset()         { *m = UnixDogstatsdMsg{} }
func (m *UnixDogstatsdMsg) String() string { return proto.CompactTextString(m) }
func (*UnixDogstatsdMsg) ProtoMessage()    {}
func (*UnixDogstatsdMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *UnixDogstatsdMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnixDogstatsdMsg.Unmarshal(m, b)
}
func (m *UnixDogstatsdMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnixDogstatsdMsg.Marshal(b, m, deterministic)
}
func (m *UnixDogstatsdMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnixDogstatsdMsg.Merge(m, src)
}
func (m *UnixDogstatsdMsg) XXX_Size() int {
	return xxx_messageInfo_UnixDogstatsdMsg.Size(m)
}
func (m *UnixDogstatsdMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UnixDogstatsdMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UnixDogstatsdMsg proto.InternalMessageInfo

func (m *UnixDogstatsdMsg) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *UnixDogstatsdMsg) GetPayloadSize() int32 {
	if m != nil {
		return m.PayloadSize
	}
	return 0
}

func (m *UnixDogstatsdMsg) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *UnixDogstatsdMsg) GetAncillarySize() int32 {
	if m != nil {
		return m.AncillarySize
	}
	return 0
}

func (m *UnixDogstatsdMsg) GetAncillary() []byte {
	if m != nil {
		return m.Ancillary
	}
	return nil
}

func init() {
	proto.RegisterType((*UnixDogstatsdMsg)(nil), "pb.UnixDogstatsdMsg")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xda,
	0xc0, 0xc8, 0x25, 0x10, 0x9a, 0x97, 0x59, 0xe1, 0x92, 0x9f, 0x5e, 0x5c, 0x92, 0x58, 0x52, 0x9c,
	0xe2, 0x5b, 0x9c, 0x2e, 0x24, 0xc3, 0xc5, 0x59, 0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b,
	0x20, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1c, 0x84, 0x10, 0x10, 0x52, 0xe0, 0xe2, 0x2e, 0x48, 0xac,
	0xcc, 0xc9, 0x4f, 0x4c, 0x09, 0xce, 0xac, 0x4a, 0x95, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0d, 0x42,
	0x16, 0x12, 0x92, 0xe0, 0x62, 0x87, 0x72, 0x25, 0x98, 0x15, 0x18, 0x35, 0x78, 0x82, 0x60, 0x5c,
	0x21, 0x15, 0x2e, 0xde, 0xc4, 0xbc, 0xe4, 0xcc, 0x9c, 0x9c, 0xc4, 0xa2, 0x4a, 0xb0, 0x6e, 0x16,
	0xb0, 0x6e, 0x54, 0x41, 0x90, 0xfd, 0x70, 0x01, 0x09, 0x56, 0xb0, 0x09, 0x08, 0x81, 0x24, 0x36,
	0xb0, 0xeb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xe3, 0x4a, 0x9a, 0xcf, 0x00, 0x00,
	0x00,
}
