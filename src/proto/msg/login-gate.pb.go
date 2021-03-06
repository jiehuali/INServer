// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: login-gate.proto

package msg

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

type LoginToGate struct {
	Cert                 *SessionCert `protobuf:"bytes,1,opt,name=Cert,proto3" json:"Cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LoginToGate) Reset()         { *m = LoginToGate{} }
func (m *LoginToGate) String() string { return proto.CompactTextString(m) }
func (*LoginToGate) ProtoMessage()    {}
func (*LoginToGate) Descriptor() ([]byte, []int) {
	return fileDescriptor_68343c62cb2204e9, []int{0}
}
func (m *LoginToGate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginToGate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginToGate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginToGate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginToGate.Merge(m, src)
}
func (m *LoginToGate) XXX_Size() int {
	return m.Size()
}
func (m *LoginToGate) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginToGate.DiscardUnknown(m)
}

var xxx_messageInfo_LoginToGate proto.InternalMessageInfo

func (m *LoginToGate) GetCert() *SessionCert {
	if m != nil {
		return m.Cert
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginToGate)(nil), "LoginToGate")
}

func init() { proto.RegisterFile("login-gate.proto", fileDescriptor_68343c62cb2204e9) }

var fileDescriptor_68343c62cb2204e9 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xc9, 0x4f, 0xcf,
	0xcc, 0xd3, 0x4d, 0x4f, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x12, 0x4c, 0xce,
	0xc9, 0x4c, 0xcd, 0x2b, 0x41, 0x12, 0x52, 0xd2, 0xe7, 0xe2, 0xf6, 0x01, 0x29, 0x0b, 0xc9, 0x77,
	0x4f, 0x2c, 0x49, 0x15, 0x52, 0xe0, 0x62, 0x71, 0x4e, 0x2d, 0x2a, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x36, 0xe2, 0xd1, 0x0b, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x03, 0x89, 0x05, 0x81, 0x65,
	0x9c, 0x34, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19,
	0x8f, 0xe5, 0x18, 0xa2, 0xc4, 0x3c, 0xfd, 0x82, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0xf4, 0x8b, 0x8b,
	0x92, 0xf5, 0xc1, 0xc6, 0xea, 0xe7, 0x16, 0xa7, 0x27, 0xb1, 0x81, 0x99, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0xd6, 0x72, 0x2b, 0x88, 0x00, 0x00, 0x00,
}

func (m *LoginToGate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginToGate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginToGate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Cert != nil {
		{
			size, err := m.Cert.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLoginGate(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLoginGate(dAtA []byte, offset int, v uint64) int {
	offset -= sovLoginGate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LoginToGate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cert != nil {
		l = m.Cert.Size()
		n += 1 + l + sovLoginGate(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLoginGate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLoginGate(x uint64) (n int) {
	return sovLoginGate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoginToGate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoginGate
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
			return fmt.Errorf("proto: LoginToGate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginToGate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cert", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginGate
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
				return ErrInvalidLengthLoginGate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoginGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cert == nil {
				m.Cert = &SessionCert{}
			}
			if err := m.Cert.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoginGate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLoginGate
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLoginGate
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
func skipLoginGate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoginGate
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
					return 0, ErrIntOverflowLoginGate
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
					return 0, ErrIntOverflowLoginGate
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
				return 0, ErrInvalidLengthLoginGate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLoginGate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLoginGate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLoginGate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoginGate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLoginGate = fmt.Errorf("proto: unexpected end of group")
)
