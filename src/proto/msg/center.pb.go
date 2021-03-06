// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: center.proto

package msg

import (
	etc "INServer/src/proto/etc"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type NodeState int32

const (
	NodeState_Unset   NodeState = 0
	NodeState_Ready   NodeState = 1
	NodeState_Running NodeState = 2
	NodeState_Offline NodeState = 3
)

var NodeState_name = map[int32]string{
	0: "Unset",
	1: "Ready",
	2: "Running",
	3: "Offline",
}

var NodeState_value = map[string]int32{
	"Unset":   0,
	"Ready":   1,
	"Running": 2,
	"Offline": 3,
}

func (x NodeState) String() string {
	return proto.EnumName(NodeState_name, int32(x))
}

func (NodeState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{0}
}

type ETCSyncNTF struct {
	BasicConfig          *etc.BasicConfig `protobuf:"bytes,1,opt,name=BasicConfig,proto3" json:"BasicConfig,omitempty"`
	Database             *etc.Database    `protobuf:"bytes,2,opt,name=Database,proto3" json:"Database,omitempty"`
	ServerList           *etc.ServerList  `protobuf:"bytes,3,opt,name=ServerList,proto3" json:"ServerList,omitempty"`
	ZoneList             *etc.ZoneList    `protobuf:"bytes,4,opt,name=ZoneList,proto3" json:"ZoneList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ETCSyncNTF) Reset()         { *m = ETCSyncNTF{} }
func (m *ETCSyncNTF) String() string { return proto.CompactTextString(m) }
func (*ETCSyncNTF) ProtoMessage()    {}
func (*ETCSyncNTF) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{0}
}
func (m *ETCSyncNTF) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ETCSyncNTF) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ETCSyncNTF.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ETCSyncNTF) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ETCSyncNTF.Merge(m, src)
}
func (m *ETCSyncNTF) XXX_Size() int {
	return m.Size()
}
func (m *ETCSyncNTF) XXX_DiscardUnknown() {
	xxx_messageInfo_ETCSyncNTF.DiscardUnknown(m)
}

var xxx_messageInfo_ETCSyncNTF proto.InternalMessageInfo

func (m *ETCSyncNTF) GetBasicConfig() *etc.BasicConfig {
	if m != nil {
		return m.BasicConfig
	}
	return nil
}

func (m *ETCSyncNTF) GetDatabase() *etc.Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *ETCSyncNTF) GetServerList() *etc.ServerList {
	if m != nil {
		return m.ServerList
	}
	return nil
}

func (m *ETCSyncNTF) GetZoneList() *etc.ZoneList {
	if m != nil {
		return m.ZoneList
	}
	return nil
}

type NodeStartNTF struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStartNTF) Reset()         { *m = NodeStartNTF{} }
func (m *NodeStartNTF) String() string { return proto.CompactTextString(m) }
func (*NodeStartNTF) ProtoMessage()    {}
func (*NodeStartNTF) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{1}
}
func (m *NodeStartNTF) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeStartNTF) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeStartNTF.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeStartNTF) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStartNTF.Merge(m, src)
}
func (m *NodeStartNTF) XXX_Size() int {
	return m.Size()
}
func (m *NodeStartNTF) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStartNTF.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStartNTF proto.InternalMessageInfo

func (m *NodeStartNTF) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type Node struct {
	NodeState            NodeState `protobuf:"varint,1,opt,name=NodeState,proto3,enum=NodeState" json:"NodeState,omitempty"`
	NodeAddress          []byte    `protobuf:"bytes,2,opt,name=NodeAddress,proto3" json:"NodeAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{2}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNodeState() NodeState {
	if m != nil {
		return m.NodeState
	}
	return NodeState_Unset
}

func (m *Node) GetNodeAddress() []byte {
	if m != nil {
		return m.NodeAddress
	}
	return nil
}

type NodesInfoNTF struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesInfoNTF) Reset()         { *m = NodesInfoNTF{} }
func (m *NodesInfoNTF) String() string { return proto.CompactTextString(m) }
func (*NodesInfoNTF) ProtoMessage()    {}
func (*NodesInfoNTF) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{3}
}
func (m *NodesInfoNTF) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodesInfoNTF) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodesInfoNTF.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodesInfoNTF) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesInfoNTF.Merge(m, src)
}
func (m *NodesInfoNTF) XXX_Size() int {
	return m.Size()
}
func (m *NodesInfoNTF) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesInfoNTF.DiscardUnknown(m)
}

var xxx_messageInfo_NodesInfoNTF proto.InternalMessageInfo

func (m *NodesInfoNTF) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type ResetConnectionNTF struct {
	ServerID             int32    `protobuf:"varint,1,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetConnectionNTF) Reset()         { *m = ResetConnectionNTF{} }
func (m *ResetConnectionNTF) String() string { return proto.CompactTextString(m) }
func (*ResetConnectionNTF) ProtoMessage()    {}
func (*ResetConnectionNTF) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de517c49d537f4b, []int{4}
}
func (m *ResetConnectionNTF) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResetConnectionNTF) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResetConnectionNTF.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResetConnectionNTF) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetConnectionNTF.Merge(m, src)
}
func (m *ResetConnectionNTF) XXX_Size() int {
	return m.Size()
}
func (m *ResetConnectionNTF) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetConnectionNTF.DiscardUnknown(m)
}

var xxx_messageInfo_ResetConnectionNTF proto.InternalMessageInfo

func (m *ResetConnectionNTF) GetServerID() int32 {
	if m != nil {
		return m.ServerID
	}
	return 0
}

func init() {
	proto.RegisterEnum("NodeState", NodeState_name, NodeState_value)
	proto.RegisterType((*ETCSyncNTF)(nil), "ETCSyncNTF")
	proto.RegisterType((*NodeStartNTF)(nil), "NodeStartNTF")
	proto.RegisterType((*Node)(nil), "Node")
	proto.RegisterType((*NodesInfoNTF)(nil), "NodesInfoNTF")
	proto.RegisterType((*ResetConnectionNTF)(nil), "ResetConnectionNTF")
}

func init() { proto.RegisterFile("center.proto", fileDescriptor_1de517c49d537f4b) }

var fileDescriptor_1de517c49d537f4b = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcd, 0x6a, 0xdb, 0x40,
	0x14, 0x85, 0x3d, 0xb6, 0x55, 0xdb, 0x57, 0xa2, 0x55, 0x67, 0x51, 0x84, 0x0b, 0xc2, 0x08, 0x0a,
	0xa2, 0x06, 0xb9, 0xb8, 0xcb, 0xae, 0x6a, 0xbb, 0x05, 0x43, 0x71, 0x61, 0xec, 0x6e, 0xbc, 0x93,
	0xa5, 0x2b, 0x23, 0x68, 0x66, 0x82, 0x66, 0x12, 0x70, 0x9e, 0x24, 0x6f, 0x92, 0x57, 0xc8, 0x32,
	0x8f, 0x10, 0x9c, 0x17, 0x09, 0x33, 0xfa, 0xb1, 0xb2, 0x9b, 0xf3, 0xdd, 0x73, 0xcf, 0x3d, 0x42,
	0xe0, 0x24, 0xc8, 0x15, 0x16, 0xd1, 0x75, 0x21, 0x94, 0x18, 0x7f, 0x44, 0x95, 0x44, 0x12, 0x8b,
	0x5b, 0x2c, 0x64, 0x85, 0x3e, 0x68, 0x74, 0x27, 0x38, 0xbe, 0x01, 0x87, 0x58, 0xe6, 0x49, 0x05,
	0xa8, 0x06, 0x69, 0xac, 0xe2, 0x43, 0x2c, 0xb1, 0x64, 0xc1, 0x03, 0x01, 0xf8, 0xb5, 0x5b, 0x6e,
	0x4f, 0x3c, 0xd9, 0xec, 0x7e, 0xd3, 0x08, 0xec, 0x85, 0xde, 0x58, 0x0a, 0x9e, 0xe5, 0x47, 0x8f,
	0x4c, 0x48, 0x68, 0xcf, 0x9d, 0xa8, 0xc5, 0x58, 0xdb, 0x40, 0xbf, 0xc0, 0x70, 0x55, 0x05, 0x7a,
	0x5d, 0x63, 0x1e, 0x45, 0x35, 0x60, 0xcd, 0x88, 0x4e, 0x01, 0xb6, 0xa6, 0xec, 0x9f, 0x5c, 0x2a,
	0xaf, 0x67, 0x8c, 0x76, 0x74, 0x41, 0xac, 0x35, 0xd6, 0x99, 0x7b, 0xc1, 0xd1, 0x58, 0xfb, 0x55,
	0x66, 0x0d, 0x58, 0x33, 0x0a, 0x42, 0x70, 0x36, 0x22, 0xc5, 0xad, 0x8a, 0x0b, 0xa5, 0xab, 0x7b,
	0x30, 0xf8, 0x99, 0xa6, 0x05, 0x4a, 0x69, 0x6a, 0x3b, 0xac, 0x96, 0x01, 0x83, 0xbe, 0x76, 0xd2,
	0x10, 0x46, 0xd5, 0x86, 0x42, 0xe3, 0x79, 0x3f, 0x87, 0xa8, 0x21, 0xec, 0x32, 0xa4, 0x13, 0xb0,
	0xb5, 0xa8, 0xf3, 0xba, 0x26, 0xaf, 0x8d, 0x82, 0x69, 0x79, 0x5d, 0xae, 0x79, 0x26, 0xf4, 0xf5,
	0xcf, 0x60, 0x19, 0xed, 0x91, 0x49, 0x2f, 0xb4, 0xe7, 0x96, 0xc9, 0x65, 0x25, 0x0b, 0xbe, 0x01,
	0x65, 0x28, 0x51, 0x2d, 0x05, 0xe7, 0x98, 0xa8, 0x5c, 0x70, 0xbd, 0x32, 0x86, 0x61, 0xf9, 0xd5,
	0xeb, 0x95, 0x69, 0x63, 0xb1, 0x46, 0x7f, 0xfd, 0xd1, 0xaa, 0x4a, 0x47, 0x60, 0xfd, 0xe3, 0x12,
	0x95, 0xdb, 0xd1, 0x4f, 0x86, 0x71, 0x7a, 0x72, 0x09, 0xb5, 0x61, 0xc0, 0x6e, 0x38, 0xcf, 0xf9,
	0xd1, 0xed, 0x6a, 0xf1, 0x37, 0xcb, 0xfe, 0xe7, 0x1c, 0xdd, 0xde, 0x22, 0x7c, 0x3c, 0xfb, 0xe4,
	0xe9, 0xec, 0x93, 0xe7, 0xb3, 0x4f, 0xee, 0x5f, 0xfc, 0xce, 0xfe, 0xd3, 0x7a, 0x53, 0x46, 0xcf,
	0x64, 0x91, 0xcc, 0xcc, 0x9f, 0x9f, 0x5d, 0xc9, 0xe3, 0xe1, 0x9d, 0x79, 0x7e, 0x7f, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x2f, 0x2e, 0x76, 0x9a, 0x5d, 0x02, 0x00, 0x00,
}

func (m *ETCSyncNTF) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ETCSyncNTF) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ETCSyncNTF) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ZoneList != nil {
		{
			size, err := m.ZoneList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCenter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ServerList != nil {
		{
			size, err := m.ServerList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCenter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Database != nil {
		{
			size, err := m.Database.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCenter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BasicConfig != nil {
		{
			size, err := m.BasicConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCenter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NodeStartNTF) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeStartNTF) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeStartNTF) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCenter(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintCenter(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.NodeState != 0 {
		i = encodeVarintCenter(dAtA, i, uint64(m.NodeState))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NodesInfoNTF) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodesInfoNTF) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodesInfoNTF) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCenter(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ResetConnectionNTF) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResetConnectionNTF) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResetConnectionNTF) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ServerID != 0 {
		i = encodeVarintCenter(dAtA, i, uint64(m.ServerID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCenter(dAtA []byte, offset int, v uint64) int {
	offset -= sovCenter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ETCSyncNTF) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BasicConfig != nil {
		l = m.BasicConfig.Size()
		n += 1 + l + sovCenter(uint64(l))
	}
	if m.Database != nil {
		l = m.Database.Size()
		n += 1 + l + sovCenter(uint64(l))
	}
	if m.ServerList != nil {
		l = m.ServerList.Size()
		n += 1 + l + sovCenter(uint64(l))
	}
	if m.ZoneList != nil {
		l = m.ZoneList.Size()
		n += 1 + l + sovCenter(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NodeStartNTF) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCenter(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NodeState != 0 {
		n += 1 + sovCenter(uint64(m.NodeState))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovCenter(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NodesInfoNTF) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovCenter(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResetConnectionNTF) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ServerID != 0 {
		n += 1 + sovCenter(uint64(m.ServerID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCenter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCenter(x uint64) (n int) {
	return sovCenter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ETCSyncNTF) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCenter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ETCSyncNTF: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ETCSyncNTF: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasicConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCenter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCenter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BasicConfig == nil {
				m.BasicConfig = &etc.BasicConfig{}
			}
			if err := m.BasicConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Database", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCenter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCenter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Database == nil {
				m.Database = &etc.Database{}
			}
			if err := m.Database.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCenter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCenter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ServerList == nil {
				m.ServerList = &etc.ServerList{}
			}
			if err := m.ServerList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZoneList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCenter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCenter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ZoneList == nil {
				m.ZoneList = &etc.ZoneList{}
			}
			if err := m.ZoneList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCenter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCenter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCenter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeStartNTF) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCenter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeStartNTF: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeStartNTF: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCenter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCenter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCenter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCenter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCenter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCenter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeState", wireType)
			}
			m.NodeState = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeState |= NodeState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCenter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCenter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = append(m.NodeAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.NodeAddress == nil {
				m.NodeAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCenter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCenter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCenter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodesInfoNTF) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCenter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodesInfoNTF: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodesInfoNTF: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCenter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCenter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, &Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCenter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCenter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCenter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResetConnectionNTF) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCenter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResetConnectionNTF: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResetConnectionNTF: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerID", wireType)
			}
			m.ServerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCenter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCenter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCenter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCenter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCenter
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCenter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCenter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCenter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCenter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCenter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCenter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCenter = fmt.Errorf("proto: unexpected end of group")
)
