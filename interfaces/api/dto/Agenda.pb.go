// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Agenda.proto

package dto

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

type Agenda struct {
	SessionItems         []string `protobuf:"bytes,1,rep,name=session_items,json=sessionItems,proto3" json:"session_items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Agenda) Reset()         { *m = Agenda{} }
func (m *Agenda) String() string { return proto.CompactTextString(m) }
func (*Agenda) ProtoMessage()    {}
func (*Agenda) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e07cbe4508d3447, []int{0}
}

func (m *Agenda) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Agenda.Unmarshal(m, b)
}
func (m *Agenda) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Agenda.Marshal(b, m, deterministic)
}
func (m *Agenda) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Agenda.Merge(m, src)
}
func (m *Agenda) XXX_Size() int {
	return xxx_messageInfo_Agenda.Size(m)
}
func (m *Agenda) XXX_DiscardUnknown() {
	xxx_messageInfo_Agenda.DiscardUnknown(m)
}

var xxx_messageInfo_Agenda proto.InternalMessageInfo

func (m *Agenda) GetSessionItems() []string {
	if m != nil {
		return m.SessionItems
	}
	return nil
}

func init() {
	proto.RegisterType((*Agenda)(nil), "beilmo.spectre.dto.Agenda")
}

func init() { proto.RegisterFile("Agenda.proto", fileDescriptor_3e07cbe4508d3447) }

var fileDescriptor_3e07cbe4508d3447 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x71, 0x4c, 0x4f, 0xcd,
	0x4b, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0x4a, 0xcd, 0xcc, 0xc9, 0xcd,
	0xd7, 0x2b, 0x2e, 0x48, 0x4d, 0x2e, 0x29, 0x4a, 0xd5, 0x4b, 0x29, 0xc9, 0x57, 0xd2, 0xe5, 0x62,
	0x83, 0xa8, 0x11, 0x52, 0xe6, 0xe2, 0x2d, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x8b, 0xcf, 0x2c,
	0x49, 0xcd, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0xe2, 0x81, 0x0a, 0x7a, 0x82, 0xc4,
	0x9c, 0x22, 0xb8, 0x24, 0x93, 0xf3, 0x73, 0xf5, 0xd0, 0x0c, 0x4a, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9,
	0x8c, 0xb2, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0xa8, 0xd0,
	0x87, 0xaa, 0xd0, 0x4d, 0xcf, 0xd7, 0x2d, 0x4a, 0x2d, 0x2e, 0xd1, 0x4d, 0x2c, 0xc8, 0xd4, 0xcf,
	0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x4e, 0x2d, 0xd6, 0x07, 0x71, 0x53, 0x4a, 0xf2, 0x77,
	0x31, 0x31, 0x24, 0xb1, 0x81, 0xdd, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x21, 0xd4, 0x57,
	0x83, 0xb3, 0x00, 0x00, 0x00,
}