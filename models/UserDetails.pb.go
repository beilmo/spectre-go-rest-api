// Code generated by protoc-gen-go. DO NOT EDIT.
// source: UserDetails.proto

package beilmo_spectre_entity

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

type UserDetails_TicketType int32

const (
	UserDetails_UNKNOWN    UserDetails_TicketType = 0
	UserDetails_FREE       UserDetails_TicketType = 1
	UserDetails_EARLY_BIRD UserDetails_TicketType = 2
	UserDetails_REGULAR    UserDetails_TicketType = 3
	UserDetails_BUNDLE     UserDetails_TicketType = 4
)

var UserDetails_TicketType_name = map[int32]string{
	0: "UNKNOWN",
	1: "FREE",
	2: "EARLY_BIRD",
	3: "REGULAR",
	4: "BUNDLE",
}

var UserDetails_TicketType_value = map[string]int32{
	"UNKNOWN":    0,
	"FREE":       1,
	"EARLY_BIRD": 2,
	"REGULAR":    3,
	"BUNDLE":     4,
}

func (x UserDetails_TicketType) String() string {
	return proto.EnumName(UserDetails_TicketType_name, int32(x))
}

func (UserDetails_TicketType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db864463f0c45ba5, []int{0, 0}
}

type UserDetails struct {
	TicketType           UserDetails_TicketType `protobuf:"varint,1,opt,name=ticketType,proto3,enum=beilmo.spectre.entity.UserDetails_TicketType" json:"ticketType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UserDetails) Reset()         { *m = UserDetails{} }
func (m *UserDetails) String() string { return proto.CompactTextString(m) }
func (*UserDetails) ProtoMessage()    {}
func (*UserDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_db864463f0c45ba5, []int{0}
}

func (m *UserDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetails.Unmarshal(m, b)
}
func (m *UserDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetails.Marshal(b, m, deterministic)
}
func (m *UserDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetails.Merge(m, src)
}
func (m *UserDetails) XXX_Size() int {
	return xxx_messageInfo_UserDetails.Size(m)
}
func (m *UserDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetails.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetails proto.InternalMessageInfo

func (m *UserDetails) GetTicketType() UserDetails_TicketType {
	if m != nil {
		return m.TicketType
	}
	return UserDetails_UNKNOWN
}

func init() {
	proto.RegisterEnum("beilmo.spectre.entity.UserDetails_TicketType", UserDetails_TicketType_name, UserDetails_TicketType_value)
	proto.RegisterType((*UserDetails)(nil), "beilmo.spectre.entity.UserDetails")
}

func init() { proto.RegisterFile("UserDetails.proto", fileDescriptor_db864463f0c45ba5) }

var fileDescriptor_db864463f0c45ba5 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0c, 0x2d, 0x4e, 0x2d,
	0x72, 0x49, 0x2d, 0x49, 0xcc, 0xcc, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4d,
	0x4a, 0xcd, 0xcc, 0xc9, 0xcd, 0xd7, 0x2b, 0x2e, 0x48, 0x4d, 0x2e, 0x29, 0x4a, 0xd5, 0x4b, 0xcd,
	0x2b, 0xc9, 0x2c, 0xa9, 0x54, 0x5a, 0xc5, 0xc8, 0xc5, 0x8d, 0xa4, 0x58, 0xc8, 0x97, 0x8b, 0xab,
	0x24, 0x33, 0x39, 0x3b, 0xb5, 0x24, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xcf,
	0x48, 0x57, 0x0f, 0xab, 0x5e, 0x3d, 0x64, 0x4b, 0x42, 0xe0, 0x9a, 0x82, 0x90, 0x0c, 0x50, 0xf2,
	0xe1, 0xe2, 0x42, 0xc8, 0x08, 0x71, 0x73, 0xb1, 0x87, 0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0x09,
	0x30, 0x08, 0x71, 0x70, 0xb1, 0xb8, 0x05, 0xb9, 0xba, 0x0a, 0x30, 0x0a, 0xf1, 0x71, 0x71, 0xb9,
	0x3a, 0x06, 0xf9, 0x44, 0xc6, 0x3b, 0x79, 0x06, 0xb9, 0x08, 0x30, 0x81, 0x94, 0x05, 0xb9, 0xba,
	0x87, 0xfa, 0x38, 0x06, 0x09, 0x30, 0x0b, 0x71, 0x71, 0xb1, 0x39, 0x85, 0xfa, 0xb9, 0xf8, 0xb8,
	0x0a, 0xb0, 0x38, 0xc9, 0x71, 0x49, 0x26, 0xe7, 0xe7, 0x62, 0x77, 0xcd, 0x2e, 0x26, 0x86, 0x24,
	0x36, 0xb0, 0x57, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x62, 0xd7, 0x0d, 0xf8, 0xff, 0x00,
	0x00, 0x00,
}
