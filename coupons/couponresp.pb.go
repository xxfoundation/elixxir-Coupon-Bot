// Code generated by protoc-gen-go. DO NOT EDIT.
// source: couponresp.proto

package coupons

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

type CMIXText struct {
	Version              uint32     `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Text                 string     `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Reply                *TextReply `protobuf:"bytes,3,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CMIXText) Reset()         { *m = CMIXText{} }
func (m *CMIXText) String() string { return proto.CompactTextString(m) }
func (*CMIXText) ProtoMessage()    {}
func (*CMIXText) Descriptor() ([]byte, []int) {
	return fileDescriptor_0df6afd1dcee6c11, []int{0}
}

func (m *CMIXText) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMIXText.Unmarshal(m, b)
}
func (m *CMIXText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMIXText.Marshal(b, m, deterministic)
}
func (m *CMIXText) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMIXText.Merge(m, src)
}
func (m *CMIXText) XXX_Size() int {
	return xxx_messageInfo_CMIXText.Size(m)
}
func (m *CMIXText) XXX_DiscardUnknown() {
	xxx_messageInfo_CMIXText.DiscardUnknown(m)
}

var xxx_messageInfo_CMIXText proto.InternalMessageInfo

func (m *CMIXText) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *CMIXText) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *CMIXText) GetReply() *TextReply {
	if m != nil {
		return m.Reply
	}
	return nil
}

type TextReply struct {
	MessageId            []byte   `protobuf:"bytes,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	SenderId             []byte   `protobuf:"bytes,2,opt,name=senderId,proto3" json:"senderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextReply) Reset()         { *m = TextReply{} }
func (m *TextReply) String() string { return proto.CompactTextString(m) }
func (*TextReply) ProtoMessage()    {}
func (*TextReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0df6afd1dcee6c11, []int{1}
}

func (m *TextReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextReply.Unmarshal(m, b)
}
func (m *TextReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextReply.Marshal(b, m, deterministic)
}
func (m *TextReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextReply.Merge(m, src)
}
func (m *TextReply) XXX_Size() int {
	return xxx_messageInfo_TextReply.Size(m)
}
func (m *TextReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TextReply.DiscardUnknown(m)
}

var xxx_messageInfo_TextReply proto.InternalMessageInfo

func (m *TextReply) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

func (m *TextReply) GetSenderId() []byte {
	if m != nil {
		return m.SenderId
	}
	return nil
}

func init() {
	proto.RegisterType((*CMIXText)(nil), "coupons.CMIXText")
	proto.RegisterType((*TextReply)(nil), "coupons.TextReply")
}

func init() {
	proto.RegisterFile("couponresp.proto", fileDescriptor_0df6afd1dcee6c11)
}

var fileDescriptor_0df6afd1dcee6c11 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xce, 0x2f, 0x2d,
	0xc8, 0xcf, 0x2b, 0x4a, 0x2d, 0x2e, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x88,
	0x14, 0x2b, 0x25, 0x71, 0x71, 0x38, 0xfb, 0x7a, 0x46, 0x84, 0xa4, 0x56, 0x94, 0x08, 0x49, 0x70,
	0xb1, 0x97, 0xa5, 0x16, 0x15, 0x67, 0xe6, 0xe7, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0xc1,
	0xb8, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0xb6, 0x90, 0x06, 0x17, 0x6b, 0x51, 0x6a, 0x41, 0x4e, 0xa5, 0x04, 0xb3, 0x02, 0xa3, 0x06,
	0xb7, 0x91, 0x90, 0x1e, 0xd4, 0x48, 0x3d, 0x90, 0x59, 0x41, 0x20, 0x99, 0x20, 0x88, 0x02, 0x25,
	0x57, 0x2e, 0x4e, 0xb8, 0x98, 0x90, 0x0c, 0x17, 0x67, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa,
	0x67, 0x0a, 0xd8, 0x1a, 0x9e, 0x20, 0x84, 0x80, 0x90, 0x14, 0x17, 0x47, 0x71, 0x6a, 0x5e, 0x4a,
	0x6a, 0x91, 0x67, 0x0a, 0xd8, 0x32, 0x9e, 0x20, 0x38, 0x3f, 0x89, 0x0d, 0xec, 0x74, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x78, 0xf4, 0x8f, 0xce, 0x00, 0x00, 0x00,
}
