// Code generated by protoc-gen-go. DO NOT EDIT.
// source: DbLedgerStorageDataFormats.proto

package bk_proto

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

//*
// Ledger metadata stored in the bookie
type LedgerData struct {
	Exists               *bool    `protobuf:"varint,1,req,name=exists" json:"exists,omitempty"`
	Fenced               *bool    `protobuf:"varint,2,req,name=fenced" json:"fenced,omitempty"`
	MasterKey            []byte   `protobuf:"bytes,3,req,name=masterKey" json:"masterKey,omitempty"`
	ExplicitLac          []byte   `protobuf:"bytes,4,opt,name=explicitLac" json:"explicitLac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LedgerData) Reset()         { *m = LedgerData{} }
func (m *LedgerData) String() string { return proto.CompactTextString(m) }
func (*LedgerData) ProtoMessage()    {}
func (*LedgerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac68653cb2d5bd9, []int{0}
}

func (m *LedgerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerData.Unmarshal(m, b)
}
func (m *LedgerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerData.Marshal(b, m, deterministic)
}
func (m *LedgerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerData.Merge(m, src)
}
func (m *LedgerData) XXX_Size() int {
	return xxx_messageInfo_LedgerData.Size(m)
}
func (m *LedgerData) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerData.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerData proto.InternalMessageInfo

func (m *LedgerData) GetExists() bool {
	if m != nil && m.Exists != nil {
		return *m.Exists
	}
	return false
}

func (m *LedgerData) GetFenced() bool {
	if m != nil && m.Fenced != nil {
		return *m.Fenced
	}
	return false
}

func (m *LedgerData) GetMasterKey() []byte {
	if m != nil {
		return m.MasterKey
	}
	return nil
}

func (m *LedgerData) GetExplicitLac() []byte {
	if m != nil {
		return m.ExplicitLac
	}
	return nil
}

func init() {
	proto.RegisterType((*LedgerData)(nil), "bk.proto.LedgerData")
}

func init() { proto.RegisterFile("DbLedgerStorageDataFormats.proto", fileDescriptor_cac68653cb2d5bd9) }

var fileDescriptor_cac68653cb2d5bd9 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0xb1, 0x0a, 0xc2, 0x30,
	0x18, 0x04, 0x60, 0x5a, 0x45, 0x6a, 0xec, 0x94, 0x41, 0x32, 0x38, 0x04, 0xa7, 0x0c, 0x92, 0x87,
	0x90, 0x22, 0x82, 0x9d, 0xea, 0x13, 0xfc, 0x49, 0x7f, 0x6b, 0x68, 0x6b, 0x42, 0x92, 0xa1, 0x82,
	0x0f, 0x2f, 0x35, 0x82, 0x6e, 0x77, 0x1f, 0xb7, 0x1c, 0xe1, 0x95, 0xaa, 0xb1, 0xed, 0xd0, 0x5f,
	0xa3, 0xf5, 0xd0, 0x61, 0x05, 0x11, 0x4e, 0xd6, 0x8f, 0x10, 0x83, 0x74, 0xde, 0x46, 0x4b, 0x0b,
	0xd5, 0xa7, 0xb4, 0x7f, 0x11, 0x92, 0xb6, 0xf3, 0x88, 0x6e, 0xc9, 0x0a, 0x27, 0x13, 0x62, 0x60,
	0x19, 0xcf, 0x45, 0xd1, 0x7c, 0xdb, 0xec, 0x37, 0x7c, 0x68, 0x6c, 0x59, 0x9e, 0x3c, 0x35, 0xba,
	0x23, 0xeb, 0x11, 0x42, 0x44, 0x7f, 0xc1, 0x27, 0x5b, 0xf0, 0x5c, 0x94, 0xcd, 0x0f, 0x28, 0x27,
	0x1b, 0x9c, 0xdc, 0x60, 0xb4, 0x89, 0x35, 0x68, 0xb6, 0xe4, 0x99, 0x28, 0x9b, 0x7f, 0x3a, 0x1e,
	0x88, 0xb0, 0xbe, 0x93, 0xe0, 0x40, 0xdf, 0x51, 0x2a, 0x6b, 0xfb, 0x1e, 0xd1, 0xa1, 0xff, 0x44,
	0x83, 0x32, 0xa4, 0x03, 0x72, 0x68, 0xd5, 0x39, 0x7b, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xf8,
	0x8c, 0x84, 0xd8, 0x00, 0x00, 0x00,
}