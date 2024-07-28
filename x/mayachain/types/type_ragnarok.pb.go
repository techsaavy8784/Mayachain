// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/type_ragnarok.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "gitlab.com/mayachain/mayanode/common"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RagnarokWithdrawPosition struct {
	Number int64        `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Pool   common.Asset `protobuf:"bytes,2,opt,name=pool,proto3" json:"pool"`
}

func (m *RagnarokWithdrawPosition) Reset()         { *m = RagnarokWithdrawPosition{} }
func (m *RagnarokWithdrawPosition) String() string { return proto.CompactTextString(m) }
func (*RagnarokWithdrawPosition) ProtoMessage()    {}
func (*RagnarokWithdrawPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_f729973ce4b551a2, []int{0}
}
func (m *RagnarokWithdrawPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RagnarokWithdrawPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RagnarokWithdrawPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RagnarokWithdrawPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RagnarokWithdrawPosition.Merge(m, src)
}
func (m *RagnarokWithdrawPosition) XXX_Size() int {
	return m.Size()
}
func (m *RagnarokWithdrawPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_RagnarokWithdrawPosition.DiscardUnknown(m)
}

var xxx_messageInfo_RagnarokWithdrawPosition proto.InternalMessageInfo

func (m *RagnarokWithdrawPosition) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RagnarokWithdrawPosition) GetPool() common.Asset {
	if m != nil {
		return m.Pool
	}
	return common.Asset{}
}

func init() {
	proto.RegisterType((*RagnarokWithdrawPosition)(nil), "types.RagnarokWithdrawPosition")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/type_ragnarok.proto", fileDescriptor_f729973ce4b551a2)
}

var fileDescriptor_f729973ce4b551a2 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xca, 0x4d, 0xac, 0x4c,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xaf, 0xd0, 0x47, 0x70, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xc1, 0x64, 0x7c, 0x51, 0x62, 0x7a, 0x5e, 0x62, 0x51, 0x7e, 0xb6, 0x5e, 0x41, 0x51,
	0x7e, 0x49, 0xbe, 0x10, 0x2b, 0x58, 0x4a, 0x4a, 0x01, 0x45, 0x6b, 0x72, 0x7e, 0x6e, 0x6e, 0x7e,
	0x1e, 0x94, 0x82, 0x28, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88,
	0xa8, 0x52, 0x34, 0x97, 0x44, 0x10, 0xd4, 0xc0, 0xf0, 0xcc, 0x92, 0x8c, 0x94, 0xa2, 0xc4, 0xf2,
	0x80, 0xfc, 0xe2, 0xcc, 0x92, 0xcc, 0xfc, 0x3c, 0x21, 0x31, 0x2e, 0xb6, 0xbc, 0xd2, 0xdc, 0xa4,
	0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x28, 0x4f, 0x48, 0x9d, 0x8b, 0xa5, 0x20,
	0x3f, 0x3f, 0x47, 0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x57, 0x0f, 0x6a, 0x8d, 0x63, 0x71,
	0x71, 0x6a, 0x89, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x60, 0x05, 0x4e, 0x9e, 0x27, 0x1e,
	0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17,
	0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9f, 0x9e, 0x59, 0x92, 0x93, 0x98, 0x04,
	0xd2, 0x8a, 0xe4, 0x4b, 0x10, 0x2b, 0x2f, 0x3f, 0x25, 0x15, 0xd3, 0xeb, 0x49, 0x6c, 0x60, 0xe7,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x32, 0x09, 0xe4, 0x23, 0x01, 0x00, 0x00,
}

func (m *RagnarokWithdrawPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RagnarokWithdrawPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RagnarokWithdrawPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Pool.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypeRagnarok(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Number != 0 {
		i = encodeVarintTypeRagnarok(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeRagnarok(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeRagnarok(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RagnarokWithdrawPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovTypeRagnarok(uint64(m.Number))
	}
	l = m.Pool.Size()
	n += 1 + l + sovTypeRagnarok(uint64(l))
	return n
}

func sovTypeRagnarok(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeRagnarok(x uint64) (n int) {
	return sovTypeRagnarok(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RagnarokWithdrawPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeRagnarok
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
			return fmt.Errorf("proto: RagnarokWithdrawPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RagnarokWithdrawPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRagnarok
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRagnarok
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
				return ErrInvalidLengthTypeRagnarok
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeRagnarok
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeRagnarok(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeRagnarok
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeRagnarok
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypeRagnarok(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeRagnarok
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
					return 0, ErrIntOverflowTypeRagnarok
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
					return 0, ErrIntOverflowTypeRagnarok
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
				return 0, ErrInvalidLengthTypeRagnarok
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeRagnarok
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeRagnarok
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeRagnarok        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeRagnarok          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeRagnarok = fmt.Errorf("proto: unexpected end of group")
)
