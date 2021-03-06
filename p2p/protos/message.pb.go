// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p/protos/message.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Node struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	IP                   string   `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_a4392b4d6f123090, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Node) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Node) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Ping struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_a4392b4d6f123090, []int{1}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (dst *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(dst, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

type Pong struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Ping                 []byte   `protobuf:"bytes,2,opt,name=Ping,proto3" json:"Ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_a4392b4d6f123090, []int{2}
}
func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (dst *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(dst, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Pong) GetPing() []byte {
	if m != nil {
		return m.Ping
	}
	return nil
}

type FindNode struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Target               []byte   `protobuf:"bytes,2,opt,name=Target,proto3" json:"Target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNode) Reset()         { *m = FindNode{} }
func (m *FindNode) String() string { return proto.CompactTextString(m) }
func (*FindNode) ProtoMessage()    {}
func (*FindNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_a4392b4d6f123090, []int{3}
}
func (m *FindNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNode.Unmarshal(m, b)
}
func (m *FindNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNode.Marshal(b, m, deterministic)
}
func (dst *FindNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNode.Merge(dst, src)
}
func (m *FindNode) XXX_Size() int {
	return xxx_messageInfo_FindNode.Size(m)
}
func (m *FindNode) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNode.DiscardUnknown(m)
}

var xxx_messageInfo_FindNode proto.InternalMessageInfo

func (m *FindNode) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *FindNode) GetTarget() []byte {
	if m != nil {
		return m.Target
	}
	return nil
}

type Neighbors struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Nodes                []*Node  `protobuf:"bytes,2,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Neighbors) Reset()         { *m = Neighbors{} }
func (m *Neighbors) String() string { return proto.CompactTextString(m) }
func (*Neighbors) ProtoMessage()    {}
func (*Neighbors) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_a4392b4d6f123090, []int{4}
}
func (m *Neighbors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Neighbors.Unmarshal(m, b)
}
func (m *Neighbors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Neighbors.Marshal(b, m, deterministic)
}
func (dst *Neighbors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Neighbors.Merge(dst, src)
}
func (m *Neighbors) XXX_Size() int {
	return xxx_messageInfo_Neighbors.Size(m)
}
func (m *Neighbors) XXX_DiscardUnknown() {
	xxx_messageInfo_Neighbors.DiscardUnknown(m)
}

var xxx_messageInfo_Neighbors proto.InternalMessageInfo

func (m *Neighbors) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Neighbors) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// tcp message
type Disc struct {
	Reason               uint32   `protobuf:"varint,1,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Disc) Reset()         { *m = Disc{} }
func (m *Disc) String() string { return proto.CompactTextString(m) }
func (*Disc) ProtoMessage()    {}
func (*Disc) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_a4392b4d6f123090, []int{5}
}
func (m *Disc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Disc.Unmarshal(m, b)
}
func (m *Disc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Disc.Marshal(b, m, deterministic)
}
func (dst *Disc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Disc.Merge(dst, src)
}
func (m *Disc) XXX_Size() int {
	return xxx_messageInfo_Disc.Size(m)
}
func (m *Disc) XXX_DiscardUnknown() {
	xxx_messageInfo_Disc.DiscardUnknown(m)
}

var xxx_messageInfo_Disc proto.InternalMessageInfo

func (m *Disc) GetReason() uint32 {
	if m != nil {
		return m.Reason
	}
	return 0
}

type Handshake struct {
	NetID                uint32   `protobuf:"varint,1,opt,name=NetID,proto3" json:"NetID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ID                   []byte   `protobuf:"bytes,3,opt,name=ID,proto3" json:"ID,omitempty"`
	Version              uint32   `protobuf:"varint,4,opt,name=Version,proto3" json:"Version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Handshake) Reset()         { *m = Handshake{} }
func (m *Handshake) String() string { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()    {}
func (*Handshake) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_a4392b4d6f123090, []int{6}
}
func (m *Handshake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Handshake.Unmarshal(m, b)
}
func (m *Handshake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Handshake.Marshal(b, m, deterministic)
}
func (dst *Handshake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handshake.Merge(dst, src)
}
func (m *Handshake) XXX_Size() int {
	return xxx_messageInfo_Handshake.Size(m)
}
func (m *Handshake) XXX_DiscardUnknown() {
	xxx_messageInfo_Handshake.DiscardUnknown(m)
}

var xxx_messageInfo_Handshake proto.InternalMessageInfo

func (m *Handshake) GetNetID() uint32 {
	if m != nil {
		return m.NetID
	}
	return 0
}

func (m *Handshake) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Handshake) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Handshake) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*Node)(nil), "protos.Node")
	proto.RegisterType((*Ping)(nil), "protos.Ping")
	proto.RegisterType((*Pong)(nil), "protos.Pong")
	proto.RegisterType((*FindNode)(nil), "protos.FindNode")
	proto.RegisterType((*Neighbors)(nil), "protos.Neighbors")
	proto.RegisterType((*Disc)(nil), "protos.Disc")
	proto.RegisterType((*Handshake)(nil), "protos.Handshake")
}

func init() { proto.RegisterFile("p2p/protos/message.proto", fileDescriptor_message_a4392b4d6f123090) }

var fileDescriptor_message_a4392b4d6f123090 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xe9, 0x9f, 0x5d, 0xed, 0xd8, 0x7a, 0x28, 0xb2, 0xe4, 0x24, 0x25, 0xa7, 0xe2, 0x61,
	0x17, 0xea, 0xcd, 0x8b, 0x97, 0x22, 0xf6, 0x52, 0x4a, 0x10, 0xaf, 0x92, 0xb5, 0x43, 0x37, 0xc8,
	0x26, 0x25, 0xd3, 0xef, 0x8f, 0x34, 0x89, 0x1e, 0x5c, 0x6f, 0xf3, 0x9b, 0xe4, 0xbd, 0x79, 0x33,
	0xc0, 0xe6, 0x66, 0x3e, 0xcc, 0xd6, 0x2c, 0x86, 0x0e, 0x67, 0x24, 0x92, 0x13, 0xee, 0x1d, 0x96,
	0x5b, 0xdf, 0xe5, 0x4f, 0x90, 0xf6, 0x66, 0xc4, 0xf2, 0x16, 0xe2, 0xae, 0x65, 0x51, 0x15, 0xd5,
	0xb9, 0x88, 0xbb, 0xd6, 0xf1, 0xc0, 0xe2, 0x2a, 0xaa, 0x33, 0x11, 0x77, 0x43, 0x59, 0x42, 0x3a,
	0x18, 0xbb, 0xb0, 0xa4, 0x8a, 0xea, 0x42, 0xb8, 0x9a, 0xef, 0x20, 0x1d, 0x94, 0x9e, 0xfe, 0x6a,
	0xf9, 0xc3, 0xfa, 0xf7, 0xb2, 0xef, 0x3c, 0x94, 0x9e, 0x9c, 0x6b, 0x2e, 0x5c, 0xcd, 0x1b, 0xb8,
	0x7e, 0x51, 0x7a, 0xfc, 0x37, 0xc3, 0x0e, 0xb6, 0x6f, 0xd2, 0x4e, 0xb8, 0x04, 0x45, 0x20, 0xfe,
	0x0c, 0x59, 0x8f, 0x6a, 0x3a, 0x1d, 0x8d, 0xa5, 0x0b, 0x11, 0x87, 0xcd, 0x6a, 0x46, 0x2c, 0xae,
	0x92, 0xfa, 0xa6, 0xc9, 0xfd, 0xbe, 0xb4, 0x5f, 0x9b, 0xc2, 0x3f, 0xf1, 0x7b, 0x48, 0x5b, 0x45,
	0x9f, 0xeb, 0x00, 0x8b, 0x92, 0x8c, 0x76, 0xfa, 0x42, 0x04, 0xe2, 0x1f, 0x90, 0xbd, 0x4a, 0x3d,
	0xd2, 0x49, 0x7e, 0x61, 0x79, 0x07, 0x9b, 0x1e, 0x97, 0x30, 0xa3, 0x10, 0x1e, 0xd6, 0x5d, 0x7a,
	0x79, 0xc6, 0x70, 0x21, 0x57, 0x87, 0x28, 0xc9, 0x6f, 0x14, 0x06, 0x57, 0xef, 0x68, 0x49, 0x19,
	0xcd, 0x52, 0xa7, 0xfd, 0xc1, 0xa3, 0xbf, 0xfe, 0xe3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4,
	0x5e, 0xad, 0x5c, 0xa0, 0x01, 0x00, 0x00,
}
