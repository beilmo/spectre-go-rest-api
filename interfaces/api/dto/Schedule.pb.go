// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Schedule.proto

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

type Schedule struct {
	Items                []*Session `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	InterestFilter       string     `protobuf:"bytes,2,opt,name=interest_filter,json=interestFilter,proto3" json:"interest_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bc31960f3ce8c15, []int{0}
}

func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schedule.Unmarshal(m, b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
}
func (m *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(m, src)
}
func (m *Schedule) XXX_Size() int {
	return xxx_messageInfo_Schedule.Size(m)
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

func (m *Schedule) GetItems() []*Session {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Schedule) GetInterestFilter() string {
	if m != nil {
		return m.InterestFilter
	}
	return ""
}

func init() {
	proto.RegisterType((*Schedule)(nil), "beilmo.spectre.dto.Schedule")
}

func init() { proto.RegisterFile("Schedule.proto", fileDescriptor_9bc31960f3ce8c15) }

var fileDescriptor_9bc31960f3ce8c15 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xbd, 0x4e, 0x85, 0x40,
	0x10, 0x46, 0x05, 0xa3, 0xd1, 0x35, 0x62, 0xb2, 0x15, 0x6a, 0x43, 0x6c, 0xa4, 0x61, 0x37, 0x6a,
	0x63, 0x6d, 0xe1, 0x03, 0x40, 0x63, 0x6c, 0x0c, 0x2c, 0x03, 0x4c, 0xc2, 0x32, 0x64, 0x77, 0x28,
	0x7c, 0xb5, 0xfb, 0x74, 0x37, 0xfc, 0x35, 0xf7, 0x96, 0xf3, 0xe5, 0x9c, 0x93, 0x8c, 0x88, 0x0a,
	0xd3, 0x41, 0x3d, 0xf5, 0xa0, 0x46, 0x47, 0x4c, 0x52, 0x56, 0x80, 0xbd, 0x25, 0xe5, 0x47, 0x30,
	0xec, 0x40, 0xd5, 0x4c, 0x4f, 0xf7, 0x05, 0x78, 0x8f, 0x34, 0xac, 0xc8, 0x4b, 0x23, 0x6e, 0x76,
	0x49, 0xbe, 0x89, 0x2b, 0x64, 0xb0, 0x3e, 0x0e, 0x92, 0xcb, 0xf4, 0xee, 0xfd, 0x59, 0x9d, 0xeb,
	0x6a, 0xb3, 0xf3, 0x95, 0x94, 0xaf, 0xe2, 0x01, 0x07, 0x06, 0x07, 0x9e, 0xff, 0x1a, 0xec, 0x19,
	0x5c, 0x1c, 0x26, 0x41, 0x7a, 0x9b, 0x47, 0xfb, 0xfc, 0xbd, 0xac, 0x5f, 0x3f, 0xe2, 0xd1, 0x90,
	0x3d, 0x2d, 0xc2, 0xc0, 0xc8, 0xff, 0xbf, 0x9f, 0x2d, 0x72, 0x37, 0x55, 0xca, 0x90, 0xd5, 0x2b,
	0xa1, 0x37, 0x22, 0x6b, 0x29, 0x9b, 0x2b, 0x59, 0x39, 0xa2, 0x5e, 0x92, 0x4d, 0x69, 0xc0, 0xeb,
	0xf9, 0xac, 0x99, 0x0e, 0xe1, 0x45, 0x75, 0xbd, 0x3c, 0xf2, 0x71, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0x8e, 0xbf, 0xbd, 0xfd, 0x00, 0x00, 0x00,
}