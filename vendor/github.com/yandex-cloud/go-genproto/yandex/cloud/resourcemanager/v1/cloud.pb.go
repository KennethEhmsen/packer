// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/resourcemanager/v1/cloud.proto

package resourcemanager

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A Cloud resource. For more information, see [Cloud](/docs/resource-manager/concepts/resources-hierarchy#cloud).
type Cloud struct {
	// ID of the cloud.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the cloud. 3-63 characters long.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the cloud. 0-256 characters long.
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cloud) Reset()         { *m = Cloud{} }
func (m *Cloud) String() string { return proto.CompactTextString(m) }
func (*Cloud) ProtoMessage()    {}
func (*Cloud) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0119578dc2f92ce, []int{0}
}

func (m *Cloud) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cloud.Unmarshal(m, b)
}
func (m *Cloud) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cloud.Marshal(b, m, deterministic)
}
func (m *Cloud) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cloud.Merge(m, src)
}
func (m *Cloud) XXX_Size() int {
	return xxx_messageInfo_Cloud.Size(m)
}
func (m *Cloud) XXX_DiscardUnknown() {
	xxx_messageInfo_Cloud.DiscardUnknown(m)
}

var xxx_messageInfo_Cloud proto.InternalMessageInfo

func (m *Cloud) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cloud) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Cloud) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cloud) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Cloud)(nil), "yandex.cloud.resourcemanager.v1.Cloud")
}

func init() {
	proto.RegisterFile("yandex/cloud/resourcemanager/v1/cloud.proto", fileDescriptor_d0119578dc2f92ce)
}

var fileDescriptor_d0119578dc2f92ce = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0xb5, 0x0a, 0x4d, 0xc1, 0x43, 0x4e, 0x4b, 0x2f, 0x5d, 0xf4, 0x52, 0x90, 0x26,
	0x54, 0x4f, 0xe2, 0x49, 0x7d, 0x83, 0xea, 0xc9, 0x8b, 0x64, 0x93, 0x31, 0x06, 0x9a, 0xcc, 0x92,
	0x9d, 0x14, 0xf5, 0x09, 0x7c, 0x6c, 0x31, 0xb1, 0xa0, 0x8b, 0xd0, 0x5b, 0xf8, 0xe7, 0x9b, 0x8f,
	0xcc, 0xcf, 0x2e, 0xde, 0x55, 0x30, 0xf0, 0x26, 0xf5, 0x16, 0x93, 0x91, 0x11, 0x06, 0x4c, 0x51,
	0x83, 0x57, 0x41, 0x59, 0x88, 0x72, 0xb7, 0x2e, 0x03, 0xd1, 0x47, 0x24, 0xe4, 0x8b, 0x02, 0x8b,
	0x92, 0x8d, 0x60, 0xb1, 0x5b, 0xcf, 0x17, 0x16, 0xd1, 0x6e, 0x41, 0x66, 0xbc, 0x4b, 0x2f, 0x92,
	0x9c, 0x87, 0x81, 0x94, 0xef, 0x8b, 0xe1, 0xec, 0xb3, 0x62, 0xc7, 0xf7, 0xdf, 0xdb, 0xfc, 0x94,
	0xd5, 0xce, 0x34, 0x55, 0x5b, 0x2d, 0xa7, 0x9b, 0xda, 0x19, 0x7e, 0xcd, 0x98, 0x8e, 0xa0, 0x08,
	0xcc, 0xb3, 0xa2, 0xa6, 0x6e, 0xab, 0xe5, 0xec, 0x72, 0x2e, 0x8a, 0x4f, 0xec, 0x7d, 0xe2, 0x71,
	0xef, 0xdb, 0x4c, 0x7f, 0xe8, 0x5b, 0xe2, 0x9c, 0x4d, 0x82, 0xf2, 0xd0, 0x1c, 0x65, 0x59, 0x7e,
	0xf3, 0x96, 0xcd, 0x0c, 0x0c, 0x3a, 0xba, 0x9e, 0x1c, 0x86, 0x66, 0x92, 0x47, 0xbf, 0xa3, 0xbb,
	0x0f, 0x76, 0xfe, 0xe7, 0x1c, 0xd5, 0xbb, 0x7f, 0x4e, 0x7a, 0x7a, 0xb0, 0x8e, 0x5e, 0x53, 0x27,
	0x34, 0x7a, 0x59, 0xf8, 0x55, 0xe9, 0xca, 0xe2, 0xca, 0x42, 0xc8, 0x3f, 0x93, 0x07, 0x4a, 0xbc,
	0x19, 0x45, 0xdd, 0x49, 0x5e, 0xbb, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x84, 0xe5, 0x1e,
	0x7e, 0x01, 0x00, 0x00,
}
