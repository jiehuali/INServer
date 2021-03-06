// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: db.role.proto

package db

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type DBRole struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty" db:"UUID"`
	SummaryData          []byte   `protobuf:"bytes,2,opt,name=SummaryData,proto3" json:"SummaryData,omitempty" db:"SummaryData"`
	OnlineData           []byte   `protobuf:"bytes,3,opt,name=OnlineData,proto3" json:"OnlineData,omitempty" db:"OnlineData"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DBRole) Reset()         { *m = DBRole{} }
func (m *DBRole) String() string { return proto.CompactTextString(m) }
func (*DBRole) ProtoMessage()    {}
func (*DBRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e0ef268370bd29, []int{0}
}
func (m *DBRole) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DBRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DBRole.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DBRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBRole.Merge(m, src)
}
func (m *DBRole) XXX_Size() int {
	return m.Size()
}
func (m *DBRole) XXX_DiscardUnknown() {
	xxx_messageInfo_DBRole.DiscardUnknown(m)
}

var xxx_messageInfo_DBRole proto.InternalMessageInfo

func (m *DBRole) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *DBRole) GetSummaryData() []byte {
	if m != nil {
		return m.SummaryData
	}
	return nil
}

func (m *DBRole) GetOnlineData() []byte {
	if m != nil {
		return m.OnlineData
	}
	return nil
}

func init() {
	proto.RegisterType((*DBRole)(nil), "DBRole")
}

func init() { proto.RegisterFile("db.role.proto", fileDescriptor_d3e0ef268370bd29) }

var fileDescriptor_d3e0ef268370bd29 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x49, 0xd2, 0x2b,
	0xca, 0xcf, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x4a, 0xcf, 0x4f, 0xcf, 0x87,
	0xb0, 0x95, 0x66, 0x30, 0x72, 0xb1, 0xb9, 0x38, 0x05, 0xe5, 0xe7, 0xa4, 0x0a, 0x29, 0x72, 0xb1,
	0x84, 0x86, 0x7a, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0xf1, 0x7e, 0xba, 0x27, 0xcf,
	0x99, 0x92, 0x64, 0xa5, 0x04, 0x12, 0x53, 0x0a, 0x02, 0x4b, 0x09, 0x99, 0x71, 0x71, 0x07, 0x97,
	0xe6, 0xe6, 0x26, 0x16, 0x55, 0xba, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x38,
	0x89, 0x7c, 0xba, 0x27, 0x2f, 0x00, 0x52, 0x89, 0x24, 0xa5, 0x14, 0x84, 0xac, 0x50, 0xc8, 0x98,
	0x8b, 0xcb, 0x3f, 0x2f, 0x27, 0x33, 0x2f, 0x15, 0xac, 0x8d, 0x19, 0xac, 0x4d, 0xf8, 0xd3, 0x3d,
	0x79, 0x7e, 0x90, 0x36, 0x84, 0x8c, 0x52, 0x10, 0x92, 0x32, 0x27, 0xf5, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x51, 0x4f,
	0xbf, 0xe0, 0xd4, 0xa2, 0xb2, 0xd4, 0x22, 0xfd, 0xe2, 0xa2, 0x64, 0x7d, 0xb0, 0x07, 0xf4, 0x53,
	0x92, 0x92, 0xd8, 0xc0, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0xab, 0xe9, 0x6a,
	0xe7, 0x00, 0x00, 0x00,
}

func (m *DBRole) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DBRole) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DBRole) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.OnlineData) > 0 {
		i -= len(m.OnlineData)
		copy(dAtA[i:], m.OnlineData)
		i = encodeVarintDbRole(dAtA, i, uint64(len(m.OnlineData)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SummaryData) > 0 {
		i -= len(m.SummaryData)
		copy(dAtA[i:], m.SummaryData)
		i = encodeVarintDbRole(dAtA, i, uint64(len(m.SummaryData)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UUID) > 0 {
		i -= len(m.UUID)
		copy(dAtA[i:], m.UUID)
		i = encodeVarintDbRole(dAtA, i, uint64(len(m.UUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDbRole(dAtA []byte, offset int, v uint64) int {
	offset -= sovDbRole(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DBRole) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovDbRole(uint64(l))
	}
	l = len(m.SummaryData)
	if l > 0 {
		n += 1 + l + sovDbRole(uint64(l))
	}
	l = len(m.OnlineData)
	if l > 0 {
		n += 1 + l + sovDbRole(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDbRole(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDbRole(x uint64) (n int) {
	return sovDbRole(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DBRole) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDbRole
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
			return fmt.Errorf("proto: DBRole: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DBRole: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDbRole
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
				return ErrInvalidLengthDbRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDbRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SummaryData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDbRole
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
				return ErrInvalidLengthDbRole
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDbRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SummaryData = append(m.SummaryData[:0], dAtA[iNdEx:postIndex]...)
			if m.SummaryData == nil {
				m.SummaryData = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnlineData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDbRole
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
				return ErrInvalidLengthDbRole
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDbRole
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OnlineData = append(m.OnlineData[:0], dAtA[iNdEx:postIndex]...)
			if m.OnlineData == nil {
				m.OnlineData = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDbRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDbRole
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDbRole
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
func skipDbRole(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDbRole
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
					return 0, ErrIntOverflowDbRole
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
					return 0, ErrIntOverflowDbRole
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
				return 0, ErrInvalidLengthDbRole
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDbRole
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDbRole
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDbRole        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDbRole          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDbRole = fmt.Errorf("proto: unexpected end of group")
)
