// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/type_solvency_voter.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "gitlab.com/mayachain/mayanode/common"
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

type SolvencyVoter struct {
	Id                   gitlab_com_mayachain_mayanode_common.TxID   `protobuf:"bytes,1,opt,name=id,proto3,casttype=gitlab.com/mayachain/mayanode/common.TxID" json:"id,omitempty"`
	Chain                gitlab_com_mayachain_mayanode_common.Chain  `protobuf:"bytes,2,opt,name=chain,proto3,casttype=gitlab.com/mayachain/mayanode/common.Chain" json:"chain,omitempty"`
	PubKey               gitlab_com_mayachain_mayanode_common.PubKey `protobuf:"bytes,3,opt,name=pub_key,json=pubKey,proto3,casttype=gitlab.com/mayachain/mayanode/common.PubKey" json:"pub_key,omitempty"`
	Coins                gitlab_com_mayachain_mayanode_common.Coins  `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=gitlab.com/mayachain/mayanode/common.Coins" json:"coins"`
	Height               int64                                       `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	ConsensusBlockHeight int64                                       `protobuf:"varint,6,opt,name=consensus_block_height,json=consensusBlockHeight,proto3" json:"consensus_block_height,omitempty"`
	Signers              []string                                    `protobuf:"bytes,7,rep,name=signers,proto3" json:"signers,omitempty"`
}

func (m *SolvencyVoter) Reset()      { *m = SolvencyVoter{} }
func (*SolvencyVoter) ProtoMessage() {}
func (*SolvencyVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_419ca5f2fbdd59a9, []int{0}
}
func (m *SolvencyVoter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SolvencyVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SolvencyVoter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SolvencyVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SolvencyVoter.Merge(m, src)
}
func (m *SolvencyVoter) XXX_Size() int {
	return m.Size()
}
func (m *SolvencyVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_SolvencyVoter.DiscardUnknown(m)
}

var xxx_messageInfo_SolvencyVoter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SolvencyVoter)(nil), "types.SolvencyVoter")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/type_solvency_voter.proto", fileDescriptor_419ca5f2fbdd59a9)
}

var fileDescriptor_419ca5f2fbdd59a9 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xc8, 0x4d, 0xac, 0x4c,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xaf, 0xd0, 0x47, 0x70, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xc1, 0x64, 0x7c, 0x71, 0x7e, 0x4e, 0x59, 0x6a, 0x5e, 0x72, 0x65, 0x7c, 0x59, 0x7e,
	0x49, 0x6a, 0x91, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x10, 0x2b, 0x58, 0x81, 0x94, 0x48, 0x7a,
	0x7e, 0x7a, 0x3e, 0x58, 0x44, 0x1f, 0xc4, 0x82, 0x48, 0x4a, 0x29, 0xa0, 0x18, 0x9b, 0x9c, 0x9f,
	0x9b, 0x9b, 0x9f, 0x07, 0xa5, 0x20, 0x2a, 0x94, 0xe6, 0x31, 0x73, 0xf1, 0x06, 0x43, 0xcd, 0x0d,
	0x03, 0x19, 0x2b, 0x64, 0xcb, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xa4,
	0xfb, 0xeb, 0x9e, 0xbc, 0x66, 0x7a, 0x66, 0x49, 0x4e, 0x62, 0x92, 0x5e, 0x72, 0x7e, 0x2e, 0x92,
	0xb3, 0x40, 0xac, 0xbc, 0xfc, 0x94, 0x54, 0x98, 0x69, 0x21, 0x15, 0x9e, 0x2e, 0x41, 0x4c, 0x99,
	0x29, 0x42, 0x2e, 0x5c, 0xac, 0x60, 0x15, 0x12, 0x4c, 0x60, 0x13, 0xf4, 0x7e, 0xdd, 0x93, 0xd7,
	0x22, 0xca, 0x04, 0x67, 0x90, 0x68, 0x10, 0x44, 0xb3, 0x90, 0x07, 0x17, 0x7b, 0x41, 0x69, 0x52,
	0x7c, 0x76, 0x6a, 0xa5, 0x04, 0x33, 0xd8, 0x1c, 0xfd, 0x5f, 0xf7, 0xe4, 0xb5, 0x89, 0x32, 0x27,
	0xa0, 0x34, 0xc9, 0x3b, 0xb5, 0x32, 0x88, 0xad, 0x00, 0x4c, 0x0b, 0x85, 0x71, 0xb1, 0x26, 0xe7,
	0x67, 0xe6, 0x15, 0x4b, 0xb0, 0x28, 0x30, 0x6b, 0x70, 0x1b, 0xf1, 0xe8, 0xc1, 0xac, 0xcb, 0xcf,
	0xcc, 0x73, 0x32, 0x3a, 0x71, 0x4f, 0x9e, 0x61, 0xd5, 0x7d, 0x62, 0x5d, 0x08, 0x32, 0x27, 0x08,
	0x62, 0x9c, 0x90, 0x18, 0x17, 0x5b, 0x46, 0x6a, 0x66, 0x7a, 0x46, 0x89, 0x04, 0xab, 0x02, 0xa3,
	0x06, 0x73, 0x10, 0x94, 0x27, 0x64, 0xc2, 0x25, 0x96, 0x9c, 0x9f, 0x57, 0x9c, 0x9a, 0x57, 0x5c,
	0x5a, 0x1c, 0x9f, 0x94, 0x93, 0x9f, 0x9c, 0x1d, 0x0f, 0x55, 0xc7, 0x06, 0x56, 0x27, 0x02, 0x97,
	0x75, 0x02, 0x49, 0x7a, 0x40, 0x74, 0x49, 0x70, 0xb1, 0x17, 0x67, 0xa6, 0xe7, 0xa5, 0x16, 0x15,
	0x4b, 0xb0, 0x2b, 0x30, 0x6b, 0x70, 0x06, 0xc1, 0xb8, 0x4e, 0xa1, 0x27, 0x1e, 0xca, 0x31, 0xdc,
	0x78, 0x28, 0xc7, 0xd0, 0xf0, 0x48, 0x8e, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5,
	0x18, 0xa2, 0xf4, 0xf1, 0x7b, 0x00, 0x23, 0x41, 0x25, 0xb1, 0x81, 0xa3, 0xdf, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x5a, 0x81, 0x9f, 0x3d, 0x79, 0x02, 0x00, 0x00,
}

func (m *SolvencyVoter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SolvencyVoter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SolvencyVoter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTypeSolvencyVoter(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.ConsensusBlockHeight != 0 {
		i = encodeVarintTypeSolvencyVoter(dAtA, i, uint64(m.ConsensusBlockHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.Height != 0 {
		i = encodeVarintTypeSolvencyVoter(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypeSolvencyVoter(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintTypeSolvencyVoter(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTypeSolvencyVoter(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTypeSolvencyVoter(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeSolvencyVoter(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeSolvencyVoter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SolvencyVoter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTypeSolvencyVoter(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypeSolvencyVoter(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovTypeSolvencyVoter(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTypeSolvencyVoter(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovTypeSolvencyVoter(uint64(m.Height))
	}
	if m.ConsensusBlockHeight != 0 {
		n += 1 + sovTypeSolvencyVoter(uint64(m.ConsensusBlockHeight))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTypeSolvencyVoter(uint64(l))
		}
	}
	return n
}

func sovTypeSolvencyVoter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeSolvencyVoter(x uint64) (n int) {
	return sovTypeSolvencyVoter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SolvencyVoter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeSolvencyVoter
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
			return fmt.Errorf("proto: SolvencyVoter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SolvencyVoter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeSolvencyVoter
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
				return ErrInvalidLengthTypeSolvencyVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeSolvencyVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = gitlab_com_mayachain_mayanode_common.TxID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeSolvencyVoter
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
				return ErrInvalidLengthTypeSolvencyVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeSolvencyVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = gitlab_com_mayachain_mayanode_common.Chain(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeSolvencyVoter
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
				return ErrInvalidLengthTypeSolvencyVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeSolvencyVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = gitlab_com_mayachain_mayanode_common.PubKey(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeSolvencyVoter
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
				return ErrInvalidLengthTypeSolvencyVoter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeSolvencyVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, common.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeSolvencyVoter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusBlockHeight", wireType)
			}
			m.ConsensusBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeSolvencyVoter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConsensusBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeSolvencyVoter
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
				return ErrInvalidLengthTypeSolvencyVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeSolvencyVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeSolvencyVoter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeSolvencyVoter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeSolvencyVoter
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
func skipTypeSolvencyVoter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeSolvencyVoter
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
					return 0, ErrIntOverflowTypeSolvencyVoter
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
					return 0, ErrIntOverflowTypeSolvencyVoter
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
				return 0, ErrInvalidLengthTypeSolvencyVoter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeSolvencyVoter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeSolvencyVoter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeSolvencyVoter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeSolvencyVoter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeSolvencyVoter = fmt.Errorf("proto: unexpected end of group")
)
