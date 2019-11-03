// Code generated by protoc-gen-go. DO NOT EDIT.
// source: UserDetails.proto

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
	TicketType           UserDetails_TicketType `protobuf:"varint,1,opt,name=ticketType,proto3,enum=beilmo.spectre.dto.UserDetails_TicketType" json:"ticketType,omitempty"`
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
	proto.RegisterEnum("beilmo.spectre.dto.UserDetails_TicketType", UserDetails_TicketType_name, UserDetails_TicketType_value)
	proto.RegisterType((*UserDetails)(nil), "beilmo.spectre.dto.UserDetails")
}

func init() { proto.RegisterFile("UserDetails.proto", fileDescriptor_db864463f0c45ba5) }

var fileDescriptor_db864463f0c45ba5 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x9b, 0x5a, 0xaa, 0x4c, 0xa1, 0xac, 0x7b, 0xd2, 0x9b, 0xf4, 0x24, 0x42, 0x36, 0xa0,
	0x17, 0xaf, 0x0d, 0x59, 0x45, 0x0d, 0x11, 0x96, 0x06, 0xff, 0x5c, 0x24, 0xd9, 0x8e, 0x75, 0xb1,
	0xe9, 0x84, 0xdd, 0xf1, 0xd0, 0x27, 0xf2, 0x1d, 0x7c, 0x3a, 0xa9, 0x8a, 0x06, 0x7b, 0x1c, 0xe6,
	0xf7, 0xfb, 0xf8, 0xf8, 0x60, 0xbf, 0x0c, 0xe8, 0x33, 0xe4, 0xca, 0x2d, 0x83, 0x6a, 0x3d, 0x31,
	0x49, 0x59, 0xa3, 0x5b, 0x36, 0xa4, 0x42, 0x8b, 0x96, 0x3d, 0xaa, 0x39, 0xd3, 0xe4, 0x3d, 0x82,
	0x51, 0x87, 0x94, 0xd7, 0x00, 0xec, 0xec, 0x2b, 0xf2, 0x6c, 0xdd, 0xe2, 0x41, 0x74, 0x14, 0x1d,
	0x8f, 0x4f, 0x4f, 0xd4, 0xb6, 0xa8, 0xba, 0xf1, 0xb3, 0x5f, 0xc3, 0x74, 0xec, 0x49, 0x0e, 0xf0,
	0xf7, 0x91, 0x23, 0xd8, 0x2d, 0x8b, 0x9b, 0xe2, 0xf6, 0xae, 0x10, 0x3d, 0xb9, 0x07, 0x83, 0x0b,
	0xa3, 0xb5, 0x88, 0xe4, 0x18, 0x40, 0x4f, 0x4d, 0xfe, 0xf0, 0x94, 0x5e, 0x99, 0x4c, 0xf4, 0x37,
	0x98, 0xd1, 0x97, 0x65, 0x3e, 0x35, 0x62, 0x47, 0x02, 0x0c, 0xd3, 0xb2, 0xc8, 0x72, 0x2d, 0x06,
	0xe9, 0x3d, 0x1c, 0x5a, 0x6a, 0xfe, 0x57, 0xc1, 0x15, 0x3b, 0x5e, 0x3f, 0x9e, 0x2f, 0x1c, 0xbf,
	0xbc, 0xd5, 0xca, 0x52, 0x93, 0x7c, 0x13, 0xc9, 0x0f, 0x11, 0x2f, 0x28, 0xf6, 0x18, 0x38, 0xae,
	0x5a, 0x97, 0xb8, 0x15, 0xa3, 0x7f, 0xae, 0x2c, 0x86, 0x64, 0x73, 0xce, 0x99, 0x3e, 0xfa, 0xbd,
	0x7a, 0xf8, 0x35, 0xcf, 0xd9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0xed, 0xc0, 0xd0, 0x33,
	0x01, 0x00, 0x00,
}