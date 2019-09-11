// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package models

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

type UserGender int32

const (
	UserGender_MALE   UserGender = 0
	UserGender_FEMALE UserGender = 1
)

var UserGender_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}

var UserGender_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x UserGender) String() string {
	return proto.EnumName(UserGender_name, int32(x))
}

func (UserGender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

type User struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Gender               UserGender `protobuf:"varint,4,opt,name=gender,proto3,enum=models.UserGender" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetGender() UserGender {
	if m != nil {
		return m.Gender
	}
	return UserGender_MALE
}

type UserList struct {
	List                 []*User  `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterEnum("models.UserGender", UserGender_name, UserGender_value)
	proto.RegisterType((*User)(nil), "models.User")
	proto.RegisterType((*UserList)(nil), "models.UserList")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0xca, 0xc2, 0x30,
	0x10, 0x84, 0xff, 0xb4, 0x21, 0xf4, 0x5f, 0xa5, 0xc8, 0x9e, 0x82, 0xa7, 0xd0, 0x53, 0x29, 0xd2,
	0x43, 0x7d, 0x02, 0x0f, 0xd5, 0x4b, 0xbd, 0x14, 0x7c, 0x80, 0x4a, 0x82, 0x04, 0xda, 0xa6, 0x24,
	0x15, 0x5f, 0x5f, 0xba, 0x8a, 0x7a, 0x9b, 0xfd, 0x66, 0x77, 0x96, 0x01, 0xb8, 0x07, 0xe3, 0xcb,
	0xc9, 0xbb, 0xd9, 0xa1, 0x18, 0x9c, 0x36, 0x7d, 0xc8, 0x3c, 0xf0, 0x4b, 0x30, 0x1e, 0x53, 0x88,
	0xac, 0x96, 0x4c, 0xb1, 0xfc, 0xbf, 0x8d, 0xac, 0x46, 0x04, 0x3e, 0x76, 0x83, 0x91, 0x11, 0x11,
	0xd2, 0xb8, 0x85, 0x64, 0xea, 0x42, 0x78, 0x38, 0xaf, 0x65, 0x4c, 0xfc, 0x33, 0x63, 0x01, 0xe2,
	0x66, 0x46, 0x6d, 0xbc, 0xe4, 0x8a, 0xe5, 0x69, 0x85, 0xe5, 0xeb, 0x41, 0xb9, 0xa4, 0x9f, 0xc8,
	0x69, 0xdf, 0x1b, 0xd9, 0x0e, 0x92, 0x85, 0x36, 0x36, 0xcc, 0xa8, 0x80, 0xf7, 0x36, 0xcc, 0x92,
	0xa9, 0x38, 0x5f, 0x55, 0xeb, 0xdf, 0xab, 0x96, 0x9c, 0x22, 0x03, 0xf8, 0x66, 0x60, 0x02, 0xfc,
	0x7c, 0x68, 0xea, 0xcd, 0x1f, 0x02, 0x88, 0x63, 0x4d, 0x9a, 0x5d, 0x05, 0x95, 0xda, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x79, 0x08, 0x9c, 0xe8, 0xe2, 0x00, 0x00, 0x00,
}
