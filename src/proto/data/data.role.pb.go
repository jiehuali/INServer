// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data.role.proto

package data

import (
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

// 离线数据
type RoleSummaryData struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Zone                 int32    `protobuf:"varint,2,opt,name=Zone,proto3" json:"Zone,omitempty"`
	RoleUUID             string   `protobuf:"bytes,3,opt,name=RoleUUID,proto3" json:"RoleUUID,omitempty"`
	MapUUID              string   `protobuf:"bytes,4,opt,name=MapUUID,proto3" json:"MapUUID,omitempty"`
	MailUUID             string   `protobuf:"bytes,5,opt,name=MailUUID,proto3" json:"MailUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleSummaryData) Reset()         { *m = RoleSummaryData{} }
func (m *RoleSummaryData) String() string { return proto.CompactTextString(m) }
func (*RoleSummaryData) ProtoMessage()    {}
func (*RoleSummaryData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bf5fdd57cc20ee2, []int{0}
}
func (m *RoleSummaryData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoleSummaryData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoleSummaryData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoleSummaryData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleSummaryData.Merge(m, src)
}
func (m *RoleSummaryData) XXX_Size() int {
	return m.Size()
}
func (m *RoleSummaryData) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleSummaryData.DiscardUnknown(m)
}

var xxx_messageInfo_RoleSummaryData proto.InternalMessageInfo

func (m *RoleSummaryData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoleSummaryData) GetZone() int32 {
	if m != nil {
		return m.Zone
	}
	return 0
}

func (m *RoleSummaryData) GetRoleUUID() string {
	if m != nil {
		return m.RoleUUID
	}
	return ""
}

func (m *RoleSummaryData) GetMapUUID() string {
	if m != nil {
		return m.MapUUID
	}
	return ""
}

func (m *RoleSummaryData) GetMailUUID() string {
	if m != nil {
		return m.MailUUID
	}
	return ""
}

// 在线数据
type RoleOnlineData struct {
	EntityData           *EntityData `protobuf:"bytes,1,opt,name=EntityData,proto3" json:"EntityData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RoleOnlineData) Reset()         { *m = RoleOnlineData{} }
func (m *RoleOnlineData) String() string { return proto.CompactTextString(m) }
func (*RoleOnlineData) ProtoMessage()    {}
func (*RoleOnlineData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bf5fdd57cc20ee2, []int{1}
}
func (m *RoleOnlineData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoleOnlineData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoleOnlineData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoleOnlineData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleOnlineData.Merge(m, src)
}
func (m *RoleOnlineData) XXX_Size() int {
	return m.Size()
}
func (m *RoleOnlineData) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleOnlineData.DiscardUnknown(m)
}

var xxx_messageInfo_RoleOnlineData proto.InternalMessageInfo

func (m *RoleOnlineData) GetEntityData() *EntityData {
	if m != nil {
		return m.EntityData
	}
	return nil
}

// 实时数据 与场景相关的数据
type RoleRealtimeData struct {
	LastStaticMapUUID    string   `protobuf:"bytes,1,opt,name=LastStaticMapUUID,proto3" json:"LastStaticMapUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleRealtimeData) Reset()         { *m = RoleRealtimeData{} }
func (m *RoleRealtimeData) String() string { return proto.CompactTextString(m) }
func (*RoleRealtimeData) ProtoMessage()    {}
func (*RoleRealtimeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bf5fdd57cc20ee2, []int{2}
}
func (m *RoleRealtimeData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoleRealtimeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoleRealtimeData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoleRealtimeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleRealtimeData.Merge(m, src)
}
func (m *RoleRealtimeData) XXX_Size() int {
	return m.Size()
}
func (m *RoleRealtimeData) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleRealtimeData.DiscardUnknown(m)
}

var xxx_messageInfo_RoleRealtimeData proto.InternalMessageInfo

func (m *RoleRealtimeData) GetLastStaticMapUUID() string {
	if m != nil {
		return m.LastStaticMapUUID
	}
	return ""
}

type Role struct {
	SummaryData          *RoleSummaryData `protobuf:"bytes,1,opt,name=SummaryData,proto3" json:"SummaryData,omitempty"`
	OnlineData           *RoleOnlineData  `protobuf:"bytes,2,opt,name=OnlineData,proto3" json:"OnlineData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bf5fdd57cc20ee2, []int{3}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Role.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return m.Size()
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetSummaryData() *RoleSummaryData {
	if m != nil {
		return m.SummaryData
	}
	return nil
}

func (m *Role) GetOnlineData() *RoleOnlineData {
	if m != nil {
		return m.OnlineData
	}
	return nil
}

func init() {
	proto.RegisterType((*RoleSummaryData)(nil), "RoleSummaryData")
	proto.RegisterType((*RoleOnlineData)(nil), "RoleOnlineData")
	proto.RegisterType((*RoleRealtimeData)(nil), "RoleRealtimeData")
	proto.RegisterType((*Role)(nil), "Role")
}

func init() { proto.RegisterFile("data.role.proto", fileDescriptor_4bf5fdd57cc20ee2) }

var fileDescriptor_4bf5fdd57cc20ee2 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x76, 0xb5, 0xf5, 0x67, 0x02, 0xa6, 0xdd, 0x8b, 0xa1, 0x87, 0x50, 0x72, 0xaa, 0x28, 0x09,
	0xc4, 0xb3, 0x20, 0x52, 0x0f, 0x05, 0x5b, 0x61, 0x43, 0x2f, 0xbd, 0x8d, 0x75, 0x0f, 0xc1, 0x4d,
	0xb6, 0x6c, 0x57, 0xa1, 0xcf, 0xe0, 0x0b, 0xf8, 0x48, 0x1e, 0x7d, 0x04, 0x89, 0x2f, 0x22, 0x3b,
	0x31, 0x36, 0xea, 0x6d, 0xbe, 0xef, 0x9b, 0x6f, 0xf6, 0xdb, 0x19, 0xf0, 0x1f, 0xd0, 0x62, 0x6c,
	0xb4, 0x92, 0xf1, 0xca, 0x68, 0xab, 0x07, 0x7d, 0x22, 0x64, 0x69, 0x73, 0xbb, 0xa9, 0xa9, 0xe8,
	0x85, 0x81, 0x2f, 0xb4, 0x92, 0xd9, 0x53, 0x51, 0xa0, 0xd9, 0x8c, 0xd1, 0x22, 0xe7, 0xd0, 0x99,
	0x61, 0x21, 0x03, 0x36, 0x64, 0xa3, 0x23, 0x41, 0xb5, 0xe3, 0x16, 0xba, 0x94, 0xc1, 0xee, 0x90,
	0x8d, 0xba, 0x82, 0x6a, 0x3e, 0x80, 0x43, 0x67, 0x9d, 0xcf, 0x27, 0xe3, 0x60, 0x8f, 0x7a, 0x7f,
	0x30, 0x0f, 0xe0, 0x60, 0x8a, 0x2b, 0x92, 0x3a, 0x24, 0x35, 0xd0, 0xb9, 0xa6, 0x98, 0x2b, 0x92,
	0xba, 0xb5, 0xab, 0xc1, 0xd1, 0x25, 0x1c, 0xbb, 0x09, 0x77, 0xa5, 0xca, 0x4b, 0x49, 0x59, 0xce,
	0x00, 0x6e, 0x28, 0xaf, 0x43, 0x94, 0xc8, 0x4b, 0xbd, 0x78, 0x4b, 0x89, 0x96, 0x1c, 0x5d, 0x41,
	0xcf, 0xd9, 0x85, 0x44, 0x65, 0xf3, 0xa2, 0x1e, 0x70, 0x0e, 0xfd, 0x5b, 0x5c, 0xdb, 0xcc, 0xa2,
	0xcd, 0x97, 0x4d, 0xa4, 0xfa, 0x67, 0xff, 0x85, 0xe8, 0x11, 0x3a, 0x6e, 0x02, 0x4f, 0xc1, 0x6b,
	0x6d, 0xe4, 0xfb, 0xdd, 0x5e, 0xfc, 0x67, 0x53, 0xa2, 0xdd, 0xc4, 0x13, 0x80, 0x6d, 0x70, 0x5a,
	0x94, 0x97, 0xfa, 0xf1, 0xef, 0xff, 0x88, 0x56, 0xcb, 0xf5, 0xe9, 0x5b, 0x15, 0xb2, 0xf7, 0x2a,
	0x64, 0x1f, 0x55, 0xc8, 0x5e, 0x3f, 0xc3, 0x9d, 0xc5, 0xc9, 0x64, 0x96, 0x49, 0xf3, 0x2c, 0x4d,
	0xb2, 0x36, 0xcb, 0x84, 0x2e, 0x94, 0xb8, 0x9b, 0xdd, 0xef, 0x53, 0x7d, 0xf1, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xb5, 0x1c, 0xe9, 0xf8, 0xd3, 0x01, 0x00, 0x00,
}

func (m *RoleSummaryData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoleSummaryData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoleSummaryData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.MailUUID) > 0 {
		i -= len(m.MailUUID)
		copy(dAtA[i:], m.MailUUID)
		i = encodeVarintDataRole(dAtA, i, uint64(len(m.MailUUID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MapUUID) > 0 {
		i -= len(m.MapUUID)
		copy(dAtA[i:], m.MapUUID)
		i = encodeVarintDataRole(dAtA, i, uint64(len(m.MapUUID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RoleUUID) > 0 {
		i -= len(m.RoleUUID)
		copy(dAtA[i:], m.RoleUUID)
		i = encodeVarintDataRole(dAtA, i, uint64(len(m.RoleUUID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Zone != 0 {
		i = encodeVarintDataRole(dAtA, i, uint64(m.Zone))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDataRole(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RoleOnlineData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoleOnlineData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoleOnlineData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EntityData != nil {
		{
			size, err := m.EntityData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDataRole(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RoleRealtimeData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoleRealtimeData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoleRealtimeData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.LastStaticMapUUID) > 0 {
		i -= len(m.LastStaticMapUUID)
		copy(dAtA[i:], m.LastStaticMapUUID)
		i = encodeVarintDataRole(dAtA, i, uint64(len(m.LastStaticMapUUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Role) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Role) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Role) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OnlineData != nil {
		{
			size, err := m.OnlineData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDataRole(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.SummaryData != nil {
		{
			size, err := m.SummaryData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDataRole(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDataRole(dAtA []byte, offset int, v uint64) int {
	offset -= sovDataRole(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RoleSummaryData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDataRole(uint64(l))
	}
	if m.Zone != 0 {
		n += 1 + sovDataRole(uint64(m.Zone))
	}
	l = len(m.RoleUUID)
	if l > 0 {
		n += 1 + l + sovDataRole(uint64(l))
	}
	l = len(m.MapUUID)
	if l > 0 {
		n += 1 + l + sovDataRole(uint64(l))
	}
	l = len(m.MailUUID)
	if l > 0 {
		n += 1 + l + sovDataRole(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RoleOnlineData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EntityData != nil {
		l = m.EntityData.Size()
		n += 1 + l + sovDataRole(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RoleRealtimeData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LastStaticMapUUID)
	if l > 0 {
		n += 1 + l + sovDataRole(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Role) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SummaryData != nil {
		l = m.SummaryData.Size()
		n += 1 + l + sovDataRole(uint64(l))
	}
	if m.OnlineData != nil {
		l = m.OnlineData.Size()
		n += 1 + l + sovDataRole(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDataRole(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDataRole(x uint64) (n int) {
	return sovDataRole(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RoleSummaryData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataRole
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
			return fmt.Errorf("proto: RoleSummaryData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoleSummaryData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zone", wireType)
			}
			m.Zone = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Zone |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MapUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MapUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MailUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MailUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDataRole
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDataRole
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
func (m *RoleOnlineData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataRole
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
			return fmt.Errorf("proto: RoleOnlineData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoleOnlineData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRole
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
				return ErrInvalidLengthDataRole
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EntityData == nil {
				m.EntityData = &EntityData{}
			}
			if err := m.EntityData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDataRole
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDataRole
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
func (m *RoleRealtimeData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataRole
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
			return fmt.Errorf("proto: RoleRealtimeData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoleRealtimeData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastStaticMapUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastStaticMapUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDataRole
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDataRole
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
func (m *Role) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataRole
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
			return fmt.Errorf("proto: Role: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Role: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SummaryData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRole
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
				return ErrInvalidLengthDataRole
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SummaryData == nil {
				m.SummaryData = &RoleSummaryData{}
			}
			if err := m.SummaryData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnlineData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRole
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
				return ErrInvalidLengthDataRole
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OnlineData == nil {
				m.OnlineData = &RoleOnlineData{}
			}
			if err := m.OnlineData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDataRole
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDataRole
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
func skipDataRole(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataRole
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
					return 0, ErrIntOverflowDataRole
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
					return 0, ErrIntOverflowDataRole
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
				return 0, ErrInvalidLengthDataRole
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDataRole
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDataRole
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDataRole        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataRole          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDataRole = fmt.Errorf("proto: unexpected end of group")
)