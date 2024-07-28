// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/msg_yggdrasil.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type MsgYggdrasil struct {
	Tx          common.Tx                                     `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx"`
	PubKey      gitlab_com_mayachain_mayanode_common.PubKey   `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3,casttype=gitlab.com/mayachain/mayanode/common.PubKey" json:"pub_key,omitempty"`
	AddFunds    bool                                          `protobuf:"varint,3,opt,name=add_funds,json=addFunds,proto3" json:"add_funds,omitempty"`
	Coins       gitlab_com_mayachain_mayanode_common.Coins    `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=gitlab.com/mayachain/mayanode/common.Coins" json:"coins"`
	BlockHeight int64                                         `protobuf:"varint,5,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Signer      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgYggdrasil) Reset()         { *m = MsgYggdrasil{} }
func (m *MsgYggdrasil) String() string { return proto.CompactTextString(m) }
func (*MsgYggdrasil) ProtoMessage()    {}
func (*MsgYggdrasil) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbf701e894cf7e06, []int{0}
}
func (m *MsgYggdrasil) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgYggdrasil) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgYggdrasil.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgYggdrasil) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgYggdrasil.Merge(m, src)
}
func (m *MsgYggdrasil) XXX_Size() int {
	return m.Size()
}
func (m *MsgYggdrasil) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgYggdrasil.DiscardUnknown(m)
}

var xxx_messageInfo_MsgYggdrasil proto.InternalMessageInfo

func (m *MsgYggdrasil) GetTx() common.Tx {
	if m != nil {
		return m.Tx
	}
	return common.Tx{}
}

func (m *MsgYggdrasil) GetPubKey() gitlab_com_mayachain_mayanode_common.PubKey {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *MsgYggdrasil) GetAddFunds() bool {
	if m != nil {
		return m.AddFunds
	}
	return false
}

func (m *MsgYggdrasil) GetCoins() gitlab_com_mayachain_mayanode_common.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *MsgYggdrasil) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MsgYggdrasil) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgYggdrasil)(nil), "types.MsgYggdrasil")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/msg_yggdrasil.proto", fileDescriptor_dbf701e894cf7e06)
}

var fileDescriptor_dbf701e894cf7e06 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x6a, 0xdb, 0x40,
	0x18, 0xd4, 0xca, 0xb6, 0x6a, 0xaf, 0x75, 0x12, 0x3d, 0x08, 0x17, 0xa4, 0x6d, 0x4f, 0xa2, 0xc5,
	0x16, 0x56, 0x9f, 0xc0, 0x2a, 0x14, 0x9b, 0x52, 0x28, 0xa2, 0x14, 0x9a, 0x8b, 0x90, 0xb4, 0xca,
	0x6a, 0xb1, 0xa5, 0x15, 0x5a, 0x29, 0x48, 0x6f, 0x91, 0xe7, 0xc8, 0x93, 0xf8, 0xe8, 0x63, 0x2e,
	0x71, 0x82, 0xfd, 0x16, 0x3e, 0x05, 0xfd, 0x98, 0x24, 0x04, 0x42, 0x4e, 0xdf, 0x7c, 0xb3, 0x3b,
	0xf3, 0x0d, 0x0c, 0xb4, 0x62, 0xaf, 0xf2, 0x82, 0xc8, 0xa3, 0x89, 0x79, 0x35, 0x37, 0x4b, 0xf3,
	0x69, 0xcd, 0xab, 0x34, 0xe4, 0x66, 0xcc, 0x89, 0x5b, 0x11, 0x82, 0x33, 0x8f, 0xd3, 0xcd, 0x2c,
	0xcd, 0x58, 0xce, 0x94, 0x41, 0xf3, 0x34, 0x41, 0x2f, 0xa4, 0x01, 0x8b, 0x63, 0x96, 0x74, 0xa3,
	0xfd, 0x38, 0xf9, 0x48, 0x18, 0x61, 0x0d, 0x34, 0x6b, 0xd4, 0xb2, 0x5f, 0xee, 0x44, 0x28, 0xff,
	0xe6, 0xe4, 0xff, 0xd9, 0x55, 0x41, 0x50, 0xcc, 0x4b, 0x15, 0x20, 0x60, 0x8c, 0x2d, 0x38, 0xeb,
	0x1c, 0xfe, 0x96, 0x76, 0x7f, 0xbb, 0xd7, 0x05, 0x47, 0xcc, 0x4b, 0x65, 0x09, 0x3f, 0xa4, 0x85,
	0xef, 0xae, 0xc3, 0x4a, 0x15, 0x11, 0x30, 0x46, 0xb6, 0x79, 0xda, 0xeb, 0xdf, 0x08, 0xcd, 0x37,
	0x9e, 0x5f, 0x0b, 0x9e, 0xc5, 0xae, 0x51, 0xc2, 0x70, 0x78, 0x4e, 0xf2, 0xa7, 0xf0, 0x7f, 0x85,
	0x95, 0x23, 0xa5, 0xcd, 0x54, 0x3e, 0xc1, 0x91, 0x87, 0xb1, 0x7b, 0x59, 0x24, 0x98, 0xab, 0x3d,
	0x04, 0x8c, 0xa1, 0x33, 0xf4, 0x30, 0xfe, 0x59, 0xef, 0xca, 0x3f, 0x38, 0x08, 0x18, 0x4d, 0xb8,
	0xda, 0x47, 0x3d, 0x63, 0x6c, 0xc9, 0xe7, 0x2c, 0x3f, 0x18, 0x4d, 0x6c, 0xab, 0x4e, 0x73, 0x73,
	0xaf, 0x7f, 0x7d, 0xd7, 0xd9, 0x5a, 0xc2, 0x9d, 0xd6, 0x4e, 0xf9, 0x0c, 0x65, 0x7f, 0xc3, 0x82,
	0xb5, 0x1b, 0x85, 0x94, 0x44, 0xb9, 0x3a, 0x40, 0xc0, 0xe8, 0x39, 0xe3, 0x86, 0x5b, 0x36, 0x94,
	0xb2, 0x82, 0x12, 0xa7, 0x24, 0x09, 0x33, 0x55, 0x42, 0xc0, 0x90, 0xed, 0xf9, 0x69, 0xaf, 0x4f,
	0x09, 0xcd, 0xa3, 0xa2, 0xbd, 0x14, 0x30, 0x1e, 0x33, 0xde, 0x8d, 0x29, 0xc7, 0xeb, 0xb6, 0x9f,
	0xd9, 0x22, 0x08, 0x16, 0x18, 0x67, 0x21, 0xe7, 0x4e, 0x67, 0x60, 0xaf, 0xb6, 0x07, 0x0d, 0xec,
	0x0e, 0x1a, 0x78, 0x38, 0x68, 0xe0, 0xfa, 0xa8, 0x09, 0xbb, 0xa3, 0x26, 0xdc, 0x1e, 0x35, 0xe1,
	0xc2, 0x7c, 0x3b, 0xfa, 0xab, 0xf6, 0x7d, 0xa9, 0x69, 0xec, 0xfb, 0x63, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcf, 0x4f, 0x0a, 0xd4, 0x26, 0x02, 0x00, 0x00,
}

func (m *MsgYggdrasil) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgYggdrasil) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgYggdrasil) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgYggdrasil(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x32
	}
	if m.BlockHeight != 0 {
		i = encodeVarintMsgYggdrasil(dAtA, i, uint64(m.BlockHeight))
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
				i = encodeVarintMsgYggdrasil(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.AddFunds {
		i--
		if m.AddFunds {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintMsgYggdrasil(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgYggdrasil(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMsgYggdrasil(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgYggdrasil(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgYggdrasil) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tx.Size()
	n += 1 + l + sovMsgYggdrasil(uint64(l))
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovMsgYggdrasil(uint64(l))
	}
	if m.AddFunds {
		n += 2
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovMsgYggdrasil(uint64(l))
		}
	}
	if m.BlockHeight != 0 {
		n += 1 + sovMsgYggdrasil(uint64(m.BlockHeight))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgYggdrasil(uint64(l))
	}
	return n
}

func sovMsgYggdrasil(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgYggdrasil(x uint64) (n int) {
	return sovMsgYggdrasil(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgYggdrasil) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgYggdrasil
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
			return fmt.Errorf("proto: MsgYggdrasil: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgYggdrasil: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgYggdrasil
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
				return ErrInvalidLengthMsgYggdrasil
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgYggdrasil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgYggdrasil
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
				return ErrInvalidLengthMsgYggdrasil
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgYggdrasil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = gitlab_com_mayachain_mayanode_common.PubKey(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddFunds", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgYggdrasil
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
			m.AddFunds = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgYggdrasil
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
				return ErrInvalidLengthMsgYggdrasil
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgYggdrasil
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
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgYggdrasil
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgYggdrasil
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
				return ErrInvalidLengthMsgYggdrasil
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgYggdrasil
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = append(m.Signer[:0], dAtA[iNdEx:postIndex]...)
			if m.Signer == nil {
				m.Signer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgYggdrasil(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgYggdrasil
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgYggdrasil
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
func skipMsgYggdrasil(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgYggdrasil
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
					return 0, ErrIntOverflowMsgYggdrasil
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
					return 0, ErrIntOverflowMsgYggdrasil
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
				return 0, ErrInvalidLengthMsgYggdrasil
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgYggdrasil
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgYggdrasil
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgYggdrasil        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgYggdrasil          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgYggdrasil = fmt.Errorf("proto: unexpected end of group")
)