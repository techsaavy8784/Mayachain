// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/type_ban_voter.proto

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

type BanVoter struct {
	NodeAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=node_address,json=nodeAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"node_address,omitempty"`
	BlockHeight int64                                         `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Signers     []string                                      `protobuf:"bytes,3,rep,name=signers,proto3" json:"signers,omitempty"`
}

func (m *BanVoter) Reset()      { *m = BanVoter{} }
func (*BanVoter) ProtoMessage() {}
func (*BanVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f1e07404e27649, []int{0}
}
func (m *BanVoter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BanVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BanVoter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BanVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BanVoter.Merge(m, src)
}
func (m *BanVoter) XXX_Size() int {
	return m.Size()
}
func (m *BanVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_BanVoter.DiscardUnknown(m)
}

var xxx_messageInfo_BanVoter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BanVoter)(nil), "types.BanVoter")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/type_ban_voter.proto", fileDescriptor_76f1e07404e27649)
}

var fileDescriptor_76f1e07404e27649 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xce, 0x4d, 0xac, 0x4c,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xaf, 0xd0, 0x47, 0x70, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xc1, 0x64, 0x7c, 0x52, 0x62, 0x5e, 0x7c, 0x59, 0x7e, 0x49, 0x6a, 0x91, 0x5e, 0x41,
	0x51, 0x7e, 0x49, 0xbe, 0x10, 0x2b, 0x58, 0x4e, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0x2c, 0xa2,
	0x0f, 0x62, 0x41, 0x24, 0x95, 0xe6, 0x32, 0x72, 0x71, 0x38, 0x25, 0xe6, 0x85, 0x81, 0xd4, 0x0b,
	0x85, 0x70, 0xf1, 0xe4, 0xe5, 0xa7, 0xa4, 0xc6, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b,
	0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x38, 0x19, 0xfe, 0xba, 0x27, 0xaf, 0x9b, 0x9e, 0x59, 0x92, 0x51,
	0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0x0c, 0xa5, 0x74, 0x8b,
	0x53, 0xb2, 0x21, 0x96, 0xeb, 0x39, 0x26, 0x27, 0x3b, 0x42, 0x34, 0x06, 0x71, 0x83, 0x8c, 0x81,
	0x72, 0x84, 0x14, 0xb9, 0x78, 0x92, 0x72, 0xf2, 0x93, 0xb3, 0xe3, 0x33, 0x52, 0x33, 0xd3, 0x33,
	0x4a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0xb8, 0xc1, 0x62, 0x1e, 0x60, 0x21, 0x21, 0x09,
	0x2e, 0xf6, 0xe2, 0xcc, 0xf4, 0xbc, 0xd4, 0xa2, 0x62, 0x09, 0x66, 0x05, 0x66, 0x0d, 0xce, 0x20,
	0x18, 0xd7, 0x29, 0xf4, 0xc4, 0x43, 0x39, 0x86, 0x1b, 0x0f, 0xe5, 0x18, 0x1a, 0x1e, 0xc9, 0x31,
	0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb,
	0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x7e, 0x7a, 0x66, 0x49, 0x4e,
	0x22, 0xc4, 0x79, 0x88, 0x30, 0x01, 0xb1, 0x40, 0x0e, 0xc1, 0x0c, 0xa8, 0x24, 0x36, 0xb0, 0xef,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x1a, 0xdb, 0x93, 0x51, 0x01, 0x00, 0x00,
}

func (m *BanVoter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BanVoter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BanVoter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTypeBanVoter(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTypeBanVoter(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintTypeBanVoter(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeBanVoter(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeBanVoter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BanVoter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovTypeBanVoter(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTypeBanVoter(uint64(m.BlockHeight))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTypeBanVoter(uint64(l))
		}
	}
	return n
}

func sovTypeBanVoter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeBanVoter(x uint64) (n int) {
	return sovTypeBanVoter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BanVoter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeBanVoter
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
			return fmt.Errorf("proto: BanVoter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BanVoter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeBanVoter
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
				return ErrInvalidLengthTypeBanVoter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeBanVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = append(m.NodeAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.NodeAddress == nil {
				m.NodeAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeBanVoter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeBanVoter
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
				return ErrInvalidLengthTypeBanVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeBanVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeBanVoter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeBanVoter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeBanVoter
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
func skipTypeBanVoter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeBanVoter
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
					return 0, ErrIntOverflowTypeBanVoter
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
					return 0, ErrIntOverflowTypeBanVoter
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
				return 0, ErrInvalidLengthTypeBanVoter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeBanVoter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeBanVoter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeBanVoter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeBanVoter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeBanVoter = fmt.Errorf("proto: unexpected end of group")
)
