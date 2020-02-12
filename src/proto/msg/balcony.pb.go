// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: balcony.proto

package msg

import (
	data "INServer/src/proto/data"
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

type RoleEnterReq struct {
	RoleUUID             string   `protobuf:"bytes,1,opt,name=RoleUUID,proto3" json:"RoleUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleEnterReq) Reset()         { *m = RoleEnterReq{} }
func (m *RoleEnterReq) String() string { return proto.CompactTextString(m) }
func (*RoleEnterReq) ProtoMessage()    {}
func (*RoleEnterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac93eeb62f414a57, []int{0}
}
func (m *RoleEnterReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoleEnterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoleEnterReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoleEnterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleEnterReq.Merge(m, src)
}
func (m *RoleEnterReq) XXX_Size() int {
	return m.Size()
}
func (m *RoleEnterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleEnterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoleEnterReq proto.InternalMessageInfo

func (m *RoleEnterReq) GetRoleUUID() string {
	if m != nil {
		return m.RoleUUID
	}
	return ""
}

type RoleEnterResp struct {
	Success              bool       `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	MapID                int32      `protobuf:"varint,2,opt,name=MapID,proto3" json:"MapID,omitempty"`
	Role                 *data.Role `protobuf:"bytes,3,opt,name=Role,proto3" json:"Role,omitempty"`
	WorldID              int32      `protobuf:"varint,4,opt,name=WorldID,proto3" json:"WorldID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RoleEnterResp) Reset()         { *m = RoleEnterResp{} }
func (m *RoleEnterResp) String() string { return proto.CompactTextString(m) }
func (*RoleEnterResp) ProtoMessage()    {}
func (*RoleEnterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac93eeb62f414a57, []int{1}
}
func (m *RoleEnterResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoleEnterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoleEnterResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoleEnterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleEnterResp.Merge(m, src)
}
func (m *RoleEnterResp) XXX_Size() int {
	return m.Size()
}
func (m *RoleEnterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleEnterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RoleEnterResp proto.InternalMessageInfo

func (m *RoleEnterResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RoleEnterResp) GetMapID() int32 {
	if m != nil {
		return m.MapID
	}
	return 0
}

func (m *RoleEnterResp) GetRole() *data.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *RoleEnterResp) GetWorldID() int32 {
	if m != nil {
		return m.WorldID
	}
	return 0
}

func init() {
	proto.RegisterType((*RoleEnterReq)(nil), "RoleEnterReq")
	proto.RegisterType((*RoleEnterResp)(nil), "RoleEnterResp")
}

func init() { proto.RegisterFile("balcony.proto", fileDescriptor_ac93eeb62f414a57) }

var fileDescriptor_ac93eeb62f414a57 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4a, 0xcc, 0x49,
	0xce, 0xcf, 0xab, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x4f, 0x49, 0x2c, 0x49, 0xd4,
	0x2b, 0xca, 0xcf, 0x49, 0x85, 0x08, 0x28, 0x69, 0x71, 0xf1, 0x04, 0xe5, 0xe7, 0xa4, 0xba, 0xe6,
	0x95, 0xa4, 0x16, 0x05, 0xa5, 0x16, 0x0a, 0x49, 0x71, 0x71, 0x80, 0xf8, 0xa1, 0xa1, 0x9e, 0x2e,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x52, 0x09, 0x17, 0x2f, 0x92, 0xda, 0xe2,
	0x02, 0x21, 0x09, 0x2e, 0xf6, 0xe0, 0xd2, 0xe4, 0xe4, 0xd4, 0xe2, 0x62, 0xb0, 0x5a, 0x8e, 0x20,
	0x18, 0x57, 0x48, 0x84, 0x8b, 0xd5, 0x37, 0xb1, 0xc0, 0xd3, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83,
	0x35, 0x08, 0xc2, 0x11, 0x92, 0xe4, 0x62, 0x01, 0x19, 0x20, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x6d,
	0xc4, 0xaa, 0x07, 0xe2, 0x04, 0x81, 0x85, 0x40, 0x46, 0x85, 0xe7, 0x17, 0xe5, 0xa4, 0x78, 0xba,
	0x48, 0xb0, 0x80, 0xb5, 0xc0, 0xb8, 0x4e, 0x1a, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x51, 0x62, 0x9e, 0x7e, 0xc1, 0xa9, 0x45,
	0x65, 0xa9, 0x45, 0xfa, 0xc5, 0x45, 0xc9, 0xfa, 0x60, 0x7f, 0xe8, 0xe7, 0x16, 0xa7, 0x27, 0xb1,
	0x81, 0x99, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0xfc, 0x73, 0xb7, 0xf4, 0x00, 0x00,
	0x00,
}

func (m *RoleEnterReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoleEnterReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoleEnterReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RoleUUID) > 0 {
		i -= len(m.RoleUUID)
		copy(dAtA[i:], m.RoleUUID)
		i = encodeVarintBalcony(dAtA, i, uint64(len(m.RoleUUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RoleEnterResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoleEnterResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoleEnterResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.WorldID != 0 {
		i = encodeVarintBalcony(dAtA, i, uint64(m.WorldID))
		i--
		dAtA[i] = 0x20
	}
	if m.Role != nil {
		{
			size, err := m.Role.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBalcony(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MapID != 0 {
		i = encodeVarintBalcony(dAtA, i, uint64(m.MapID))
		i--
		dAtA[i] = 0x10
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBalcony(dAtA []byte, offset int, v uint64) int {
	offset -= sovBalcony(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RoleEnterReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RoleUUID)
	if l > 0 {
		n += 1 + l + sovBalcony(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RoleEnterResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	if m.MapID != 0 {
		n += 1 + sovBalcony(uint64(m.MapID))
	}
	if m.Role != nil {
		l = m.Role.Size()
		n += 1 + l + sovBalcony(uint64(l))
	}
	if m.WorldID != 0 {
		n += 1 + sovBalcony(uint64(m.WorldID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBalcony(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBalcony(x uint64) (n int) {
	return sovBalcony(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RoleEnterReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalcony
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
			return fmt.Errorf("proto: RoleEnterReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoleEnterReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalcony
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
				return ErrInvalidLengthBalcony
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalcony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBalcony(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBalcony
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBalcony
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
func (m *RoleEnterResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalcony
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
			return fmt.Errorf("proto: RoleEnterResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoleEnterResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalcony
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MapID", wireType)
			}
			m.MapID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalcony
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MapID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalcony
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
				return ErrInvalidLengthBalcony
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalcony
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Role == nil {
				m.Role = &data.Role{}
			}
			if err := m.Role.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorldID", wireType)
			}
			m.WorldID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalcony
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WorldID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBalcony(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBalcony
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBalcony
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
func skipBalcony(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBalcony
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
					return 0, ErrIntOverflowBalcony
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
					return 0, ErrIntOverflowBalcony
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
				return 0, ErrInvalidLengthBalcony
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBalcony
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBalcony
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBalcony        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBalcony          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBalcony = fmt.Errorf("proto: unexpected end of group")
)
