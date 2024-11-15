// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/msg_withdraw_liquidity.proto

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

type MsgWithdrawLiquidity struct {
	Tx              common.Tx                                     `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx"`
	WithdrawAddress gitlab_com_mayachain_mayanode_common.Address  `protobuf:"bytes,2,opt,name=withdraw_address,json=withdrawAddress,proto3,casttype=gitlab.com/mayachain/mayanode/common.Address" json:"withdraw_address,omitempty"`
	BasisPoints     github_com_cosmos_cosmos_sdk_types.Uint       `protobuf:"bytes,3,opt,name=basis_points,json=basisPoints,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"basis_points"`
	Asset           common.Asset                                  `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset"`
	WithdrawalAsset common.Asset                                  `protobuf:"bytes,5,opt,name=withdrawal_asset,json=withdrawalAsset,proto3" json:"withdrawal_asset"`
	Signer          github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgWithdrawLiquidity) Reset()         { *m = MsgWithdrawLiquidity{} }
func (m *MsgWithdrawLiquidity) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawLiquidity) ProtoMessage()    {}
func (*MsgWithdrawLiquidity) Descriptor() ([]byte, []int) {
	return fileDescriptor_60fc185a6989792c, []int{0}
}
func (m *MsgWithdrawLiquidity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawLiquidity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawLiquidity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawLiquidity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawLiquidity.Merge(m, src)
}
func (m *MsgWithdrawLiquidity) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawLiquidity) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawLiquidity.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawLiquidity proto.InternalMessageInfo

func (m *MsgWithdrawLiquidity) GetTx() common.Tx {
	if m != nil {
		return m.Tx
	}
	return common.Tx{}
}

func (m *MsgWithdrawLiquidity) GetWithdrawAddress() gitlab_com_mayachain_mayanode_common.Address {
	if m != nil {
		return m.WithdrawAddress
	}
	return ""
}

func (m *MsgWithdrawLiquidity) GetAsset() common.Asset {
	if m != nil {
		return m.Asset
	}
	return common.Asset{}
}

func (m *MsgWithdrawLiquidity) GetWithdrawalAsset() common.Asset {
	if m != nil {
		return m.WithdrawalAsset
	}
	return common.Asset{}
}

func (m *MsgWithdrawLiquidity) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgWithdrawLiquidity)(nil), "types.MsgWithdrawLiquidity")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/msg_withdraw_liquidity.proto", fileDescriptor_60fc185a6989792c)
}

var fileDescriptor_60fc185a6989792c = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x93, 0xfe, 0x83, 0x3b, 0xed, 0xa5, 0x97, 0xd0, 0x45, 0xe8, 0x22, 0x09, 0x77, 0x63,
	0x05, 0xdb, 0x58, 0x5d, 0x0a, 0x42, 0xb3, 0x2b, 0x28, 0x48, 0x50, 0x04, 0x5d, 0x84, 0x69, 0x12,
	0xd2, 0xc1, 0x26, 0x53, 0x73, 0xa6, 0x36, 0x7d, 0x0b, 0xdf, 0xc0, 0xd7, 0xe9, 0xb2, 0x4b, 0x71,
	0x11, 0xa4, 0x7d, 0x8b, 0xae, 0x24, 0x93, 0x89, 0x55, 0x44, 0x71, 0x75, 0x4e, 0x0e, 0xdf, 0xf7,
	0xcb, 0xf9, 0x4e, 0x82, 0x4e, 0x42, 0xbc, 0xc0, 0xee, 0x18, 0x93, 0xc8, 0x7c, 0xe8, 0x9b, 0x89,
	0xb9, 0x7b, 0x64, 0x8b, 0xa9, 0x0f, 0x66, 0x08, 0x81, 0x33, 0x27, 0x6c, 0xec, 0xc5, 0x78, 0xee,
	0x4c, 0xc8, 0xfd, 0x8c, 0x78, 0x84, 0x2d, 0x7a, 0xd3, 0x98, 0x32, 0xaa, 0x54, 0xb9, 0xa6, 0x6d,
	0x7c, 0x62, 0xb8, 0x34, 0x0c, 0x69, 0x24, 0x4a, 0x2e, 0x6c, 0xb7, 0x02, 0x1a, 0x50, 0xde, 0x9a,
	0x59, 0x97, 0x4f, 0xff, 0x3f, 0x95, 0x51, 0xeb, 0x1c, 0x82, 0x6b, 0x81, 0x3f, 0x2b, 0xe8, 0x8a,
	0x81, 0x4a, 0x2c, 0x51, 0x65, 0x43, 0xee, 0xd4, 0x8f, 0x50, 0x4f, 0x90, 0x2e, 0x13, 0xab, 0xb2,
	0x4c, 0x75, 0xc9, 0x2e, 0xb1, 0x44, 0xb9, 0x45, 0xff, 0xde, 0xb7, 0xc2, 0x9e, 0x17, 0xfb, 0x00,
	0x6a, 0xc9, 0x90, 0x3b, 0x7f, 0xac, 0xc3, 0x6d, 0xaa, 0x1f, 0x04, 0x84, 0x4d, 0xf0, 0x28, 0x73,
	0x7e, 0x08, 0x94, 0x75, 0x11, 0xf5, 0xfc, 0x62, 0xb5, 0x41, 0xee, 0xb3, 0x9b, 0x05, 0x49, 0x0c,
	0x14, 0x1b, 0x35, 0x46, 0x18, 0x08, 0x38, 0x53, 0x4a, 0x22, 0x06, 0x6a, 0x99, 0x83, 0xcd, 0xec,
	0xe5, 0x2f, 0xa9, 0xbe, 0x17, 0x10, 0x36, 0x9e, 0xe5, 0x70, 0x97, 0x42, 0x48, 0x41, 0x94, 0x2e,
	0x78, 0x77, 0xf9, 0xd5, 0x7a, 0x57, 0x24, 0x62, 0x76, 0x9d, 0x43, 0x2e, 0x38, 0x43, 0xd9, 0x47,
	0x55, 0x0c, 0xe0, 0x33, 0xb5, 0xc2, 0x53, 0xfd, 0x2d, 0x52, 0x0d, 0xb2, 0xa1, 0x08, 0x96, 0x2b,
	0x94, 0xd3, 0x5d, 0x36, 0x3c, 0x71, 0x72, 0x57, 0xf5, 0x7b, 0x57, 0x73, 0x27, 0xe6, 0x63, 0x65,
	0x88, 0x6a, 0x40, 0x82, 0xc8, 0x8f, 0xd5, 0x9a, 0x21, 0x77, 0x1a, 0x56, 0x7f, 0x9b, 0xea, 0xdd,
	0x5f, 0x2c, 0x3d, 0x70, 0xdd, 0xe2, 0x24, 0x02, 0x60, 0x0d, 0x97, 0x6b, 0x4d, 0x5e, 0xad, 0x35,
	0xf9, 0x75, 0xad, 0xc9, 0x8f, 0x1b, 0x4d, 0x5a, 0x6d, 0x34, 0xe9, 0x79, 0xa3, 0x49, 0x37, 0xe6,
	0xcf, 0x27, 0xfe, 0xf2, 0x23, 0x8d, 0x6a, 0xfc, 0x9b, 0x1f, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xdb, 0x27, 0x58, 0xad, 0x71, 0x02, 0x00, 0x00,
}

func (m *MsgWithdrawLiquidity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawLiquidity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawLiquidity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgWithdrawLiquidity(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.WithdrawalAsset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgWithdrawLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgWithdrawLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.BasisPoints.Size()
		i -= size
		if _, err := m.BasisPoints.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsgWithdrawLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.WithdrawAddress) > 0 {
		i -= len(m.WithdrawAddress)
		copy(dAtA[i:], m.WithdrawAddress)
		i = encodeVarintMsgWithdrawLiquidity(dAtA, i, uint64(len(m.WithdrawAddress)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgWithdrawLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMsgWithdrawLiquidity(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgWithdrawLiquidity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgWithdrawLiquidity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tx.Size()
	n += 1 + l + sovMsgWithdrawLiquidity(uint64(l))
	l = len(m.WithdrawAddress)
	if l > 0 {
		n += 1 + l + sovMsgWithdrawLiquidity(uint64(l))
	}
	l = m.BasisPoints.Size()
	n += 1 + l + sovMsgWithdrawLiquidity(uint64(l))
	l = m.Asset.Size()
	n += 1 + l + sovMsgWithdrawLiquidity(uint64(l))
	l = m.WithdrawalAsset.Size()
	n += 1 + l + sovMsgWithdrawLiquidity(uint64(l))
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgWithdrawLiquidity(uint64(l))
	}
	return n
}

func sovMsgWithdrawLiquidity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgWithdrawLiquidity(x uint64) (n int) {
	return sovMsgWithdrawLiquidity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgWithdrawLiquidity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgWithdrawLiquidity
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
			return fmt.Errorf("proto: MsgWithdrawLiquidity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawLiquidity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgWithdrawLiquidity
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
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgWithdrawLiquidity
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
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgWithdrawLiquidity
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
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawAddress = gitlab_com_mayachain_mayanode_common.Address(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasisPoints", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgWithdrawLiquidity
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
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BasisPoints.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgWithdrawLiquidity
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
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalAsset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgWithdrawLiquidity
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
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WithdrawalAsset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgWithdrawLiquidity
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
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgWithdrawLiquidity
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
			skippy, err := skipMsgWithdrawLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgWithdrawLiquidity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgWithdrawLiquidity
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
func skipMsgWithdrawLiquidity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgWithdrawLiquidity
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
					return 0, ErrIntOverflowMsgWithdrawLiquidity
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
					return 0, ErrIntOverflowMsgWithdrawLiquidity
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
				return 0, ErrInvalidLengthMsgWithdrawLiquidity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgWithdrawLiquidity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgWithdrawLiquidity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgWithdrawLiquidity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgWithdrawLiquidity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgWithdrawLiquidity = fmt.Errorf("proto: unexpected end of group")
)
