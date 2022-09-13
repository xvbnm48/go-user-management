// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screen_message.proto

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

type Screen_Panel int32

const (
	Screen_UNKNOWN Screen_Panel = 0
	Screen_IPS     Screen_Panel = 1
	Screen_OLED    Screen_Panel = 2
)

var Screen_Panel_name = map[int32]string{
	0: "UNKNOWN",
	1: "IPS",
	2: "OLED",
}

var Screen_Panel_value = map[string]int32{
	"UNKNOWN": 0,
	"IPS":     1,
	"OLED":    2,
}

func (x Screen_Panel) String() string {
	return proto.EnumName(Screen_Panel_name, int32(x))
}

func (Screen_Panel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

type Screen struct {
	SizeInc              float32            `protobuf:"fixed32,1,opt,name=size_inc,json=sizeInc,proto3" json:"size_inc,omitempty"`
	Resolution           *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Panel                Screen_Panel       `protobuf:"varint,3,opt,name=panel,proto3,enum=otakutech.pcbook.Screen_Panel" json:"panel,omitempty"`
	Multitouch           bool               `protobuf:"varint,4,opt,name=multitouch,proto3" json:"multitouch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Screen) Reset()         { *m = Screen{} }
func (m *Screen) String() string { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()    {}
func (*Screen) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0}
}

func (m *Screen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen.Unmarshal(m, b)
}
func (m *Screen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen.Marshal(b, m, deterministic)
}
func (m *Screen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen.Merge(m, src)
}
func (m *Screen) XXX_Size() int {
	return xxx_messageInfo_Screen.Size(m)
}
func (m *Screen) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen.DiscardUnknown(m)
}

var xxx_messageInfo_Screen proto.InternalMessageInfo

func (m *Screen) GetSizeInc() float32 {
	if m != nil {
		return m.SizeInc
	}
	return 0
}

func (m *Screen) GetResolution() *Screen_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Screen) GetPanel() Screen_Panel {
	if m != nil {
		return m.Panel
	}
	return Screen_UNKNOWN
}

func (m *Screen) GetMultitouch() bool {
	if m != nil {
		return m.Multitouch
	}
	return false
}

type Screen_Resolution struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Screen_Resolution) Reset()         { *m = Screen_Resolution{} }
func (m *Screen_Resolution) String() string { return proto.CompactTextString(m) }
func (*Screen_Resolution) ProtoMessage()    {}
func (*Screen_Resolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

func (m *Screen_Resolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen_Resolution.Unmarshal(m, b)
}
func (m *Screen_Resolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen_Resolution.Marshal(b, m, deterministic)
}
func (m *Screen_Resolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen_Resolution.Merge(m, src)
}
func (m *Screen_Resolution) XXX_Size() int {
	return xxx_messageInfo_Screen_Resolution.Size(m)
}
func (m *Screen_Resolution) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen_Resolution.DiscardUnknown(m)
}

var xxx_messageInfo_Screen_Resolution proto.InternalMessageInfo

func (m *Screen_Resolution) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Screen_Resolution) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("otakutech.pcbook.Screen_Panel", Screen_Panel_name, Screen_Panel_value)
	proto.RegisterType((*Screen)(nil), "otakutech.pcbook.Screen")
	proto.RegisterType((*Screen_Resolution)(nil), "otakutech.pcbook.Screen.Resolution")
}

func init() {
	proto.RegisterFile("screen_message.proto", fileDescriptor_8a2814cd02b8aba7)
}

var fileDescriptor_8a2814cd02b8aba7 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xdd, 0x34, 0xff, 0x98, 0x52, 0x09, 0x43, 0x91, 0xe8, 0xa1, 0x84, 0x7a, 0x30, 0xa7,
	0x1c, 0xaa, 0x27, 0x8f, 0xfe, 0x39, 0x14, 0x25, 0x2d, 0x5b, 0x44, 0xf0, 0x52, 0x92, 0x75, 0x69,
	0x96, 0xa6, 0xbb, 0x21, 0xbb, 0x41, 0xf0, 0xeb, 0xf8, 0x45, 0xa5, 0x1b, 0xd1, 0x22, 0xf4, 0xf8,
	0x86, 0xdf, 0x7b, 0x6f, 0x66, 0x60, 0xac, 0x59, 0xcb, 0xb9, 0x5c, 0xef, 0xb8, 0xd6, 0xc5, 0x86,
	0x67, 0x4d, 0xab, 0x8c, 0xc2, 0x48, 0x99, 0x62, 0xdb, 0x19, 0xce, 0xaa, 0xac, 0x61, 0xa5, 0x52,
	0xdb, 0xe9, 0x97, 0x03, 0xfe, 0xca, 0xa2, 0x78, 0x0e, 0xa1, 0x16, 0x9f, 0x7c, 0x2d, 0x24, 0x8b,
	0x49, 0x42, 0x52, 0x87, 0x06, 0x7b, 0x3d, 0x97, 0x0c, 0xef, 0x01, 0x5a, 0xae, 0x55, 0xdd, 0x19,
	0xa1, 0x64, 0xec, 0x24, 0x24, 0x1d, 0xce, 0x2e, 0xb3, 0xff, 0x61, 0x59, 0x1f, 0x94, 0xd1, 0x5f,
	0x94, 0x1e, 0xd8, 0xf0, 0x06, 0xbc, 0xa6, 0x90, 0xbc, 0x8e, 0x07, 0x09, 0x49, 0x4f, 0x67, 0x93,
	0xa3, 0xfe, 0xe5, 0x9e, 0xa2, 0x3d, 0x8c, 0x13, 0x80, 0x5d, 0x57, 0x1b, 0x61, 0x54, 0xc7, 0xaa,
	0xd8, 0x4d, 0x48, 0x1a, 0xd2, 0x83, 0xc9, 0xc5, 0x2d, 0xc0, 0x5f, 0x1f, 0x8e, 0xc1, 0xfb, 0x10,
	0xef, 0xa6, 0xb2, 0x07, 0x8c, 0x68, 0x2f, 0xf0, 0x0c, 0xfc, 0x8a, 0x8b, 0x4d, 0x65, 0xec, 0xea,
	0x23, 0xfa, 0xa3, 0xa6, 0x57, 0xe0, 0xd9, 0x2e, 0x1c, 0x42, 0xf0, 0x92, 0x3f, 0xe5, 0x8b, 0xd7,
	0x3c, 0x3a, 0xc1, 0x00, 0x06, 0xf3, 0xe5, 0x2a, 0x22, 0x18, 0x82, 0xbb, 0x78, 0x7e, 0x7c, 0x88,
	0x9c, 0x3b, 0xf7, 0xcd, 0x69, 0xca, 0xd2, 0xb7, 0x4f, 0xbc, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0x90, 0x7f, 0xda, 0x5c, 0x01, 0x00, 0x00,
}