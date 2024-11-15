// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/type_errata_tx_voter.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	gitlab_com_mayachain_mayanode_common "gitlab.com/mayachain/mayanode/common"
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

type ErrataTxVoter struct {
	TxID        gitlab_com_mayachain_mayanode_common.TxID  `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,casttype=gitlab.com/mayachain/mayanode/common.TxID" json:"tx_id,omitempty"`
	Chain       gitlab_com_mayachain_mayanode_common.Chain `protobuf:"bytes,2,opt,name=chain,proto3,casttype=gitlab.com/mayachain/mayanode/common.Chain" json:"chain,omitempty"`
	BlockHeight int64                                      `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Signers     []string                                   `protobuf:"bytes,4,rep,name=signers,proto3" json:"signers,omitempty"`
}

func (m *ErrataTxVoter) Reset()      { *m = ErrataTxVoter{} }
func (*ErrataTxVoter) ProtoMessage() {}
func (*ErrataTxVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8f81aeea701bb5, []int{0}
}
func (m *ErrataTxVoter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrataTxVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrataTxVoter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrataTxVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrataTxVoter.Merge(m, src)
}
func (m *ErrataTxVoter) XXX_Size() int {
	return m.Size()
}
func (m *ErrataTxVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrataTxVoter.DiscardUnknown(m)
}

var xxx_messageInfo_ErrataTxVoter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ErrataTxVoter)(nil), "types.ErrataTxVoter")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/type_errata_tx_voter.proto", fileDescriptor_8b8f81aeea701bb5)
}

var fileDescriptor_8b8f81aeea701bb5 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xcc, 0x4d, 0xac, 0x4c,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xaf, 0xd0, 0x47, 0x70, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xc1, 0x64, 0x7c, 0x6a, 0x51, 0x51, 0x62, 0x49, 0x62, 0x7c, 0x49, 0x45, 0x7c, 0x59,
	0x7e, 0x49, 0x6a, 0x91, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x10, 0x2b, 0x58, 0x85, 0x94, 0x48,
	0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x44, 0x1f, 0xc4, 0x82, 0x48, 0x2a, 0xdd, 0x63, 0xe4, 0xe2, 0x75,
	0x05, 0x6b, 0x0b, 0xa9, 0x08, 0x03, 0x69, 0x12, 0xf2, 0xe2, 0x62, 0x2d, 0xa9, 0x88, 0xcf, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x32, 0x7d, 0x74, 0x4f, 0x9e, 0x25, 0xa4, 0xc2, 0xd3,
	0xe5, 0xd7, 0x3d, 0x79, 0xcd, 0xf4, 0xcc, 0x92, 0x9c, 0xc4, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x24,
	0x07, 0x80, 0x58, 0x79, 0xf9, 0x29, 0xa9, 0xfa, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x7a, 0x20,
	0xc5, 0x41, 0x2c, 0x25, 0x15, 0x9e, 0x29, 0x42, 0x2e, 0x5c, 0xac, 0x60, 0x35, 0x12, 0x4c, 0x60,
	0xb3, 0xf4, 0x7e, 0xdd, 0x93, 0xd7, 0x22, 0xca, 0x0c, 0x67, 0x90, 0x68, 0x10, 0x44, 0xb3, 0x90,
	0x22, 0x17, 0x4f, 0x52, 0x4e, 0x7e, 0x72, 0x76, 0x7c, 0x46, 0x6a, 0x66, 0x7a, 0x46, 0x89, 0x04,
	0xb3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x37, 0x58, 0xcc, 0x03, 0x2c, 0x24, 0x24, 0xc1, 0xc5, 0x5e,
	0x9c, 0x99, 0x9e, 0x97, 0x5a, 0x54, 0x2c, 0xc1, 0xa2, 0xc0, 0xac, 0xc1, 0x19, 0x04, 0xe3, 0x3a,
	0x85, 0x9e, 0x78, 0x28, 0xc7, 0x70, 0xe3, 0xa1, 0x1c, 0x43, 0xc3, 0x23, 0x39, 0x86, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0xc7, 0xef, 0x2a, 0x8c, 0xf0, 0x4e, 0x62,
	0x03, 0x07, 0x9f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x5d, 0xcb, 0x1d, 0x98, 0x01, 0x00,
	0x00,
}

func (m *ErrataTxVoter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrataTxVoter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrataTxVoter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTypeErrataTxVoter(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTypeErrataTxVoter(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTypeErrataTxVoter(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxID) > 0 {
		i -= len(m.TxID)
		copy(dAtA[i:], m.TxID)
		i = encodeVarintTypeErrataTxVoter(dAtA, i, uint64(len(m.TxID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeErrataTxVoter(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeErrataTxVoter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ErrataTxVoter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxID)
	if l > 0 {
		n += 1 + l + sovTypeErrataTxVoter(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypeErrataTxVoter(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTypeErrataTxVoter(uint64(m.BlockHeight))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTypeErrataTxVoter(uint64(l))
		}
	}
	return n
}

func sovTypeErrataTxVoter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeErrataTxVoter(x uint64) (n int) {
	return sovTypeErrataTxVoter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ErrataTxVoter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeErrataTxVoter
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
			return fmt.Errorf("proto: ErrataTxVoter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrataTxVoter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeErrataTxVoter
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
				return ErrInvalidLengthTypeErrataTxVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeErrataTxVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxID = gitlab_com_mayachain_mayanode_common.TxID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeErrataTxVoter
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
				return ErrInvalidLengthTypeErrataTxVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeErrataTxVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = gitlab_com_mayachain_mayanode_common.Chain(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeErrataTxVoter
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeErrataTxVoter
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
				return ErrInvalidLengthTypeErrataTxVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeErrataTxVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeErrataTxVoter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeErrataTxVoter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeErrataTxVoter
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
func skipTypeErrataTxVoter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeErrataTxVoter
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
					return 0, ErrIntOverflowTypeErrataTxVoter
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
					return 0, ErrIntOverflowTypeErrataTxVoter
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
				return 0, ErrInvalidLengthTypeErrataTxVoter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeErrataTxVoter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeErrataTxVoter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeErrataTxVoter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeErrataTxVoter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeErrataTxVoter = fmt.Errorf("proto: unexpected end of group")
)
