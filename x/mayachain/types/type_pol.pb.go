// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/type_pol.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ProtocolOwnedLiquidity struct {
	CacaoDeposited github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,1,opt,name=cacao_deposited,json=cacaoDeposited,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"cacao_deposited"`
	CacaoWithdrawn github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=cacao_withdrawn,json=cacaoWithdrawn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"cacao_withdrawn"`
}

func (m *ProtocolOwnedLiquidity) Reset()         { *m = ProtocolOwnedLiquidity{} }
func (m *ProtocolOwnedLiquidity) String() string { return proto.CompactTextString(m) }
func (*ProtocolOwnedLiquidity) ProtoMessage()    {}
func (*ProtocolOwnedLiquidity) Descriptor() ([]byte, []int) {
	return fileDescriptor_56049012f7da4012, []int{0}
}
func (m *ProtocolOwnedLiquidity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolOwnedLiquidity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolOwnedLiquidity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolOwnedLiquidity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolOwnedLiquidity.Merge(m, src)
}
func (m *ProtocolOwnedLiquidity) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolOwnedLiquidity) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolOwnedLiquidity.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolOwnedLiquidity proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProtocolOwnedLiquidity)(nil), "types.ProtocolOwnedLiquidity")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/type_pol.proto", fileDescriptor_56049012f7da4012)
}

var fileDescriptor_56049012f7da4012 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcd, 0x4d, 0xac, 0x4c,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xaf, 0xd0, 0x47, 0x70, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xc1, 0x64, 0x7c, 0x41, 0x7e, 0x8e, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x10, 0x2b,
	0x58, 0x54, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0x2c, 0xa2, 0x0f, 0x62, 0x41, 0x24, 0x95, 0xce,
	0x30, 0x72, 0x89, 0x05, 0x80, 0x58, 0xc9, 0xf9, 0x39, 0xfe, 0xe5, 0x79, 0xa9, 0x29, 0x3e, 0x99,
	0x85, 0xa5, 0x99, 0x29, 0x99, 0x25, 0x95, 0x42, 0x11, 0x5c, 0xfc, 0xc9, 0x89, 0xc9, 0x89, 0xf9,
	0xf1, 0x29, 0xa9, 0x05, 0xf9, 0xc5, 0x99, 0x25, 0xa9, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x4e, 0xfa, 0x27, 0xee, 0xc9, 0x33, 0xdc, 0xba, 0x27, 0xaf, 0x9e, 0x9e, 0x59, 0x92, 0x51, 0x9a,
	0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0x0c, 0xa5, 0x74, 0x8b, 0x53,
	0xb2, 0x21, 0x6e, 0xd1, 0x0b, 0xcd, 0xcc, 0x2b, 0x09, 0xe2, 0x03, 0x9b, 0xe3, 0x02, 0x33, 0x06,
	0x61, 0x72, 0x79, 0x66, 0x49, 0x46, 0x4a, 0x51, 0x62, 0x79, 0x9e, 0x04, 0x13, 0x25, 0x26, 0x87,
	0xc3, 0x8c, 0x71, 0xf2, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xfd,
	0xf4, 0xcc, 0x92, 0x9c, 0x44, 0x88, 0x91, 0x88, 0x00, 0x03, 0xb1, 0xf2, 0xf2, 0x53, 0x52, 0x31,
	0x43, 0x31, 0x89, 0x0d, 0x1c, 0x40, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x43, 0xcd,
	0xd7, 0x6e, 0x01, 0x00, 0x00,
}

func (m *ProtocolOwnedLiquidity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolOwnedLiquidity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolOwnedLiquidity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CacaoWithdrawn.Size()
		i -= size
		if _, err := m.CacaoWithdrawn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypePol(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.CacaoDeposited.Size()
		i -= size
		if _, err := m.CacaoDeposited.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypePol(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypePol(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypePol(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProtocolOwnedLiquidity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CacaoDeposited.Size()
	n += 1 + l + sovTypePol(uint64(l))
	l = m.CacaoWithdrawn.Size()
	n += 1 + l + sovTypePol(uint64(l))
	return n
}

func sovTypePol(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypePol(x uint64) (n int) {
	return sovTypePol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProtocolOwnedLiquidity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypePol
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
			return fmt.Errorf("proto: ProtocolOwnedLiquidity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolOwnedLiquidity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacaoDeposited", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePol
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
				return ErrInvalidLengthTypePol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypePol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CacaoDeposited.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacaoWithdrawn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePol
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
				return ErrInvalidLengthTypePol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypePol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CacaoWithdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypePol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypePol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypePol
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
func skipTypePol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypePol
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
					return 0, ErrIntOverflowTypePol
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
					return 0, ErrIntOverflowTypePol
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
				return 0, ErrInvalidLengthTypePol
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypePol
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypePol
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypePol        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypePol          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypePol = fmt.Errorf("proto: unexpected end of group")
)
