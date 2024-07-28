// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/msg_swap.proto

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

type OrderType int32

const (
	OrderType_market OrderType = 0
	OrderType_limit  OrderType = 1
)

var OrderType_name = map[int32]string{
	0: "market",
	1: "limit",
}

var OrderType_value = map[string]int32{
	"market": 0,
	"limit":  1,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}

func (OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b1915766ca9ad929, []int{0}
}

type MsgSwap struct {
	Tx                      common.Tx                                     `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx"`
	TargetAsset             common.Asset                                  `protobuf:"bytes,2,opt,name=target_asset,json=targetAsset,proto3" json:"target_asset"`
	Destination             gitlab_com_mayachain_mayanode_common.Address  `protobuf:"bytes,3,opt,name=destination,proto3,casttype=gitlab.com/mayachain/mayanode/common.Address" json:"destination,omitempty"`
	TradeTarget             github_com_cosmos_cosmos_sdk_types.Uint       `protobuf:"bytes,4,opt,name=trade_target,json=tradeTarget,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"trade_target"`
	AffiliateAddress        gitlab_com_mayachain_mayanode_common.Address  `protobuf:"bytes,5,opt,name=affiliate_address,json=affiliateAddress,proto3,casttype=gitlab.com/mayachain/mayanode/common.Address" json:"affiliate_address,omitempty"`
	AffiliateBasisPoints    github_com_cosmos_cosmos_sdk_types.Uint       `protobuf:"bytes,6,opt,name=affiliate_basis_points,json=affiliateBasisPoints,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"affiliate_basis_points"`
	Signer                  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,7,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
	Aggregator              string                                        `protobuf:"bytes,8,opt,name=aggregator,proto3" json:"aggregator,omitempty"`
	AggregatorTargetAddress string                                        `protobuf:"bytes,9,opt,name=aggregator_target_address,json=aggregatorTargetAddress,proto3" json:"aggregator_target_address,omitempty"`
	AggregatorTargetLimit   *github_com_cosmos_cosmos_sdk_types.Uint      `protobuf:"bytes,10,opt,name=aggregator_target_limit,json=aggregatorTargetLimit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"aggregator_target_limit,omitempty"`
	OrderType               OrderType                                     `protobuf:"varint,11,opt,name=order_type,json=orderType,proto3,enum=types.OrderType" json:"order_type,omitempty"`
	StreamQuantity          uint64                                        `protobuf:"varint,12,opt,name=stream_quantity,json=streamQuantity,proto3" json:"stream_quantity,omitempty"`
	StreamInterval          uint64                                        `protobuf:"varint,13,opt,name=stream_interval,json=streamInterval,proto3" json:"stream_interval,omitempty"`
}

func (m *MsgSwap) Reset()         { *m = MsgSwap{} }
func (m *MsgSwap) String() string { return proto.CompactTextString(m) }
func (*MsgSwap) ProtoMessage()    {}
func (*MsgSwap) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1915766ca9ad929, []int{0}
}
func (m *MsgSwap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwap.Merge(m, src)
}
func (m *MsgSwap) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwap) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwap.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwap proto.InternalMessageInfo

func (m *MsgSwap) GetTx() common.Tx {
	if m != nil {
		return m.Tx
	}
	return common.Tx{}
}

func (m *MsgSwap) GetTargetAsset() common.Asset {
	if m != nil {
		return m.TargetAsset
	}
	return common.Asset{}
}

func (m *MsgSwap) GetDestination() gitlab_com_mayachain_mayanode_common.Address {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *MsgSwap) GetAffiliateAddress() gitlab_com_mayachain_mayanode_common.Address {
	if m != nil {
		return m.AffiliateAddress
	}
	return ""
}

func (m *MsgSwap) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *MsgSwap) GetAggregator() string {
	if m != nil {
		return m.Aggregator
	}
	return ""
}

func (m *MsgSwap) GetAggregatorTargetAddress() string {
	if m != nil {
		return m.AggregatorTargetAddress
	}
	return ""
}

func (m *MsgSwap) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_market
}

func (m *MsgSwap) GetStreamQuantity() uint64 {
	if m != nil {
		return m.StreamQuantity
	}
	return 0
}

func (m *MsgSwap) GetStreamInterval() uint64 {
	if m != nil {
		return m.StreamInterval
	}
	return 0
}

func init() {
	proto.RegisterEnum("types.OrderType", OrderType_name, OrderType_value)
	proto.RegisterType((*MsgSwap)(nil), "types.MsgSwap")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/msg_swap.proto", fileDescriptor_b1915766ca9ad929)
}

var fileDescriptor_b1915766ca9ad929 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6f, 0x12, 0x41,
	0x14, 0xc7, 0xd9, 0x5a, 0xa8, 0x3c, 0x68, 0xc5, 0x49, 0xb5, 0x63, 0x0f, 0xcb, 0xa6, 0x97, 0x12,
	0x23, 0xac, 0xad, 0x89, 0x07, 0x6f, 0x70, 0x23, 0xd1, 0xa8, 0x2b, 0x5e, 0x4c, 0xcc, 0x66, 0x60,
	0xa7, 0xd3, 0x49, 0xd9, 0x9d, 0x75, 0xe6, 0xb5, 0x85, 0xff, 0xc2, 0x3f, 0xab, 0xc7, 0x1e, 0x8d,
	0x07, 0x62, 0xc0, 0xbf, 0xa2, 0x27, 0xb3, 0xbf, 0x00, 0x6d, 0x62, 0xb4, 0xa7, 0x79, 0xf3, 0xe5,
	0xfb, 0x3e, 0xef, 0x47, 0x86, 0x85, 0x76, 0xc8, 0xa6, 0x6c, 0x74, 0xca, 0x64, 0xe4, 0x5e, 0x1c,
	0xb9, 0x13, 0x77, 0x75, 0xc5, 0x69, 0xcc, 0x8d, 0x1b, 0x1a, 0xe1, 0x9b, 0x4b, 0x16, 0x77, 0x62,
	0xad, 0x50, 0x91, 0x72, 0xaa, 0xee, 0x3b, 0xbf, 0x65, 0x8d, 0x54, 0x18, 0xaa, 0x28, 0x3f, 0x32,
	0xe3, 0xfe, 0xae, 0x50, 0x42, 0xa5, 0xa1, 0x9b, 0x44, 0x99, 0x7a, 0xf0, 0xb3, 0x02, 0x5b, 0x6f,
	0x8c, 0xf8, 0x70, 0xc9, 0x62, 0xe2, 0xc0, 0x06, 0x4e, 0xa8, 0xe5, 0x58, 0xad, 0xda, 0x31, 0x74,
	0xf2, 0xe4, 0xc1, 0xa4, 0xb7, 0x79, 0x35, 0x6b, 0x96, 0xbc, 0x0d, 0x9c, 0x90, 0x97, 0x50, 0x47,
	0xa6, 0x05, 0x47, 0x9f, 0x19, 0xc3, 0x91, 0x6e, 0xa4, 0xde, 0xed, 0xc2, 0xdb, 0x4d, 0xc4, 0xdc,
	0x5e, 0xcb, 0x8c, 0xa9, 0x44, 0x3c, 0xa8, 0x05, 0xdc, 0xa0, 0x8c, 0x18, 0x4a, 0x15, 0xd1, 0x7b,
	0x8e, 0xd5, 0xaa, 0xf6, 0x9e, 0xdf, 0xcc, 0x9a, 0xcf, 0x84, 0xc4, 0x31, 0x1b, 0x26, 0x80, 0xb5,
	0x41, 0x93, 0x28, 0x52, 0x01, 0x2f, 0x06, 0xe8, 0x06, 0x81, 0xe6, 0xc6, 0x78, 0xeb, 0x10, 0xe2,
	0x41, 0x1d, 0x35, 0x0b, 0xb8, 0x9f, 0x15, 0xa2, 0x9b, 0x29, 0xd4, 0x4d, 0x8a, 0x7f, 0x9f, 0x35,
	0x0f, 0x85, 0xc4, 0xd3, 0xf3, 0x0c, 0x3c, 0x52, 0x26, 0x54, 0x26, 0x3f, 0xda, 0x26, 0x38, 0xcb,
	0x36, 0xd9, 0xf9, 0x28, 0x23, 0xf4, 0x6a, 0x29, 0x64, 0x90, 0x32, 0xc8, 0x67, 0x78, 0xc8, 0x4e,
	0x4e, 0xe4, 0x58, 0x32, 0xe4, 0x3e, 0xcb, 0xaa, 0xd2, 0xf2, 0x1d, 0xbb, 0x6d, 0x2c, 0x51, 0xb9,
	0x42, 0x38, 0x3c, 0x5e, 0xe1, 0x87, 0xcc, 0x48, 0xe3, 0xc7, 0x4a, 0x46, 0x68, 0x68, 0xe5, 0x6e,
	0xcd, 0xef, 0x2e, 0x71, 0xbd, 0x84, 0xf6, 0x2e, 0x85, 0x91, 0x3e, 0x54, 0x8c, 0x14, 0x11, 0xd7,
	0x74, 0xcb, 0xb1, 0x5a, 0xf5, 0xde, 0xd1, 0xcd, 0xac, 0xd9, 0xfe, 0x07, 0x64, 0x77, 0x34, 0x2a,
	0x7a, 0xcf, 0x01, 0xc4, 0x06, 0x60, 0x42, 0x68, 0x2e, 0x18, 0x2a, 0x4d, 0xef, 0x27, 0x5d, 0x7a,
	0x6b, 0x0a, 0x79, 0x05, 0x4f, 0x56, 0x37, 0xbf, 0x78, 0x1b, 0xf9, 0xe2, 0xaa, 0xa9, 0x7d, 0x6f,
	0x65, 0xc8, 0xb6, 0x5c, 0x6c, 0x43, 0xc0, 0xde, 0xed, 0xdc, 0xb1, 0x0c, 0x25, 0x52, 0x58, 0xae,
	0xc3, 0xfa, 0x9f, 0x75, 0x3c, 0xfa, 0xb3, 0xd4, 0xeb, 0x84, 0x46, 0x5c, 0x00, 0xa5, 0x03, 0xae,
	0xfd, 0xc4, 0x4a, 0x6b, 0x8e, 0xd5, 0xda, 0x39, 0x6e, 0x74, 0xb2, 0xbc, 0xb7, 0xc9, 0x0f, 0x83,
	0x69, 0xcc, 0xbd, 0xaa, 0x2a, 0x42, 0x72, 0x08, 0x0f, 0x0c, 0x6a, 0xce, 0x42, 0xff, 0xcb, 0x39,
	0x8b, 0x50, 0xe2, 0x94, 0xd6, 0x1d, 0xab, 0xb5, 0xe9, 0xed, 0x64, 0xf2, 0xfb, 0x5c, 0x5d, 0x33,
	0xca, 0x08, 0xb9, 0xbe, 0x60, 0x63, 0xba, 0xbd, 0x6e, 0xec, 0xe7, 0xea, 0xd3, 0x03, 0xa8, 0x2e,
	0x2b, 0x11, 0x80, 0x4a, 0xc8, 0xf4, 0x19, 0xc7, 0x46, 0x89, 0x54, 0xa1, 0x9c, 0x8e, 0xdc, 0xb0,
	0x7a, 0xfd, 0xab, 0xb9, 0x6d, 0x5d, 0xcf, 0x6d, 0xeb, 0xc7, 0xdc, 0xb6, 0xbe, 0x2e, 0xec, 0xd2,
	0xf5, 0xc2, 0x2e, 0x7d, 0x5b, 0xd8, 0xa5, 0x4f, 0xee, 0xdf, 0xdf, 0xdd, 0xad, 0x6f, 0xc4, 0xb0,
	0x92, 0xfe, 0xb9, 0x5f, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x66, 0x70, 0x63, 0xa9, 0x4c, 0x04,
	0x00, 0x00,
}

func (m *MsgSwap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StreamInterval != 0 {
		i = encodeVarintMsgSwap(dAtA, i, uint64(m.StreamInterval))
		i--
		dAtA[i] = 0x68
	}
	if m.StreamQuantity != 0 {
		i = encodeVarintMsgSwap(dAtA, i, uint64(m.StreamQuantity))
		i--
		dAtA[i] = 0x60
	}
	if m.OrderType != 0 {
		i = encodeVarintMsgSwap(dAtA, i, uint64(m.OrderType))
		i--
		dAtA[i] = 0x58
	}
	if m.AggregatorTargetLimit != nil {
		{
			size := m.AggregatorTargetLimit.Size()
			i -= size
			if _, err := m.AggregatorTargetLimit.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintMsgSwap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.AggregatorTargetAddress) > 0 {
		i -= len(m.AggregatorTargetAddress)
		copy(dAtA[i:], m.AggregatorTargetAddress)
		i = encodeVarintMsgSwap(dAtA, i, uint64(len(m.AggregatorTargetAddress)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Aggregator) > 0 {
		i -= len(m.Aggregator)
		copy(dAtA[i:], m.Aggregator)
		i = encodeVarintMsgSwap(dAtA, i, uint64(len(m.Aggregator)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgSwap(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size := m.AffiliateBasisPoints.Size()
		i -= size
		if _, err := m.AffiliateBasisPoints.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsgSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.AffiliateAddress) > 0 {
		i -= len(m.AffiliateAddress)
		copy(dAtA[i:], m.AffiliateAddress)
		i = encodeVarintMsgSwap(dAtA, i, uint64(len(m.AffiliateAddress)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.TradeTarget.Size()
		i -= size
		if _, err := m.TradeTarget.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsgSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Destination) > 0 {
		i -= len(m.Destination)
		copy(dAtA[i:], m.Destination)
		i = encodeVarintMsgSwap(dAtA, i, uint64(len(m.Destination)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.TargetAsset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMsgSwap(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgSwap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSwap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tx.Size()
	n += 1 + l + sovMsgSwap(uint64(l))
	l = m.TargetAsset.Size()
	n += 1 + l + sovMsgSwap(uint64(l))
	l = len(m.Destination)
	if l > 0 {
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	l = m.TradeTarget.Size()
	n += 1 + l + sovMsgSwap(uint64(l))
	l = len(m.AffiliateAddress)
	if l > 0 {
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	l = m.AffiliateBasisPoints.Size()
	n += 1 + l + sovMsgSwap(uint64(l))
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	l = len(m.Aggregator)
	if l > 0 {
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	l = len(m.AggregatorTargetAddress)
	if l > 0 {
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	if m.AggregatorTargetLimit != nil {
		l = m.AggregatorTargetLimit.Size()
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	if m.OrderType != 0 {
		n += 1 + sovMsgSwap(uint64(m.OrderType))
	}
	if m.StreamQuantity != 0 {
		n += 1 + sovMsgSwap(uint64(m.StreamQuantity))
	}
	if m.StreamInterval != 0 {
		n += 1 + sovMsgSwap(uint64(m.StreamInterval))
	}
	return n
}

func sovMsgSwap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgSwap(x uint64) (n int) {
	return sovMsgSwap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSwap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgSwap
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
			return fmt.Errorf("proto: MsgSwap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
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
				return fmt.Errorf("proto: wrong wireType = %d for field TargetAsset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TargetAsset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destination = gitlab_com_mayachain_mayanode_common.Address(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TradeTarget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AffiliateAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AffiliateAddress = gitlab_com_mayachain_mayanode_common.Address(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AffiliateBasisPoints", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AffiliateBasisPoints.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = append(m.Signer[:0], dAtA[iNdEx:postIndex]...)
			if m.Signer == nil {
				m.Signer = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aggregator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Aggregator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregatorTargetAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregatorTargetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregatorTargetLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Uint
			m.AggregatorTargetLimit = &v
			if err := m.AggregatorTargetLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderType", wireType)
			}
			m.OrderType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderType |= OrderType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamQuantity", wireType)
			}
			m.StreamQuantity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamQuantity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamInterval", wireType)
			}
			m.StreamInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsgSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgSwap
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
func skipMsgSwap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgSwap
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
					return 0, ErrIntOverflowMsgSwap
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
					return 0, ErrIntOverflowMsgSwap
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
				return 0, ErrInvalidLengthMsgSwap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgSwap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgSwap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgSwap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgSwap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgSwap = fmt.Errorf("proto: unexpected end of group")
)