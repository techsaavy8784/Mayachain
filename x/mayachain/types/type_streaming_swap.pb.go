// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/type_streaming_swap.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type StreamingSwap struct {
	TxID              gitlab_com_mayachain_mayanode_common.TxID `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,casttype=gitlab.com/mayachain/mayanode/common.TxID" json:"tx_id,omitempty"`
	Interval          uint64                                    `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Quantity          uint64                                    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Count             uint64                                    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	LastHeight        int64                                     `protobuf:"varint,5,opt,name=last_height,json=lastHeight,proto3" json:"last_height,omitempty"`
	TradeTarget       github_com_cosmos_cosmos_sdk_types.Uint   `protobuf:"bytes,6,opt,name=trade_target,json=tradeTarget,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"trade_target"`
	Deposit           github_com_cosmos_cosmos_sdk_types.Uint   `protobuf:"bytes,7,opt,name=deposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"deposit"`
	In                github_com_cosmos_cosmos_sdk_types.Uint   `protobuf:"bytes,8,opt,name=in,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"in"`
	Out               github_com_cosmos_cosmos_sdk_types.Uint   `protobuf:"bytes,9,opt,name=out,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"out"`
	FailedSwaps       []uint64                                  `protobuf:"varint,10,rep,packed,name=failed_swaps,json=failedSwaps,proto3" json:"failed_swaps,omitempty"`
	FailedSwapReasons []string                                  `protobuf:"bytes,11,rep,name=failed_swap_reasons,json=failedSwapReasons,proto3" json:"failed_swap_reasons,omitempty"`
}

func (m *StreamingSwap) Reset()         { *m = StreamingSwap{} }
func (m *StreamingSwap) String() string { return proto.CompactTextString(m) }
func (*StreamingSwap) ProtoMessage()    {}
func (*StreamingSwap) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfaea44c1b899b2e, []int{0}
}
func (m *StreamingSwap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamingSwap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamingSwap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamingSwap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingSwap.Merge(m, src)
}
func (m *StreamingSwap) XXX_Size() int {
	return m.Size()
}
func (m *StreamingSwap) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingSwap.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingSwap proto.InternalMessageInfo

func (m *StreamingSwap) GetTxID() gitlab_com_mayachain_mayanode_common.TxID {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *StreamingSwap) GetInterval() uint64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *StreamingSwap) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *StreamingSwap) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *StreamingSwap) GetLastHeight() int64 {
	if m != nil {
		return m.LastHeight
	}
	return 0
}

func (m *StreamingSwap) GetFailedSwaps() []uint64 {
	if m != nil {
		return m.FailedSwaps
	}
	return nil
}

func (m *StreamingSwap) GetFailedSwapReasons() []string {
	if m != nil {
		return m.FailedSwapReasons
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamingSwap)(nil), "types.StreamingSwap")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/type_streaming_swap.proto", fileDescriptor_cfaea44c1b899b2e)
}

var fileDescriptor_cfaea44c1b899b2e = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x3d, 0x8f, 0x94, 0x40,
	0x18, 0x5e, 0x16, 0xf6, 0xee, 0x76, 0xf6, 0x2c, 0x1c, 0xaf, 0x98, 0x5c, 0x01, 0x68, 0x23, 0x16,
	0x42, 0x8c, 0x31, 0xb1, 0x33, 0x5e, 0x2c, 0xc4, 0x72, 0xee, 0x6c, 0x6c, 0xc8, 0x1c, 0x8c, 0x30,
	0x11, 0x66, 0x90, 0x79, 0xb9, 0x63, 0xff, 0x85, 0x3f, 0xc9, 0xf2, 0xca, 0x2b, 0x8d, 0x05, 0x31,
	0xec, 0xbf, 0xb0, 0x32, 0x0c, 0xae, 0xb7, 0x89, 0xc9, 0x15, 0xdb, 0xc0, 0xfb, 0x7c, 0xbc, 0x0f,
	0x1f, 0xf3, 0xa0, 0xd7, 0x15, 0x5b, 0xb3, 0xb4, 0x60, 0x42, 0x46, 0x57, 0x2f, 0xa2, 0x2e, 0xba,
	0x83, 0xb0, 0xae, 0xb9, 0x36, 0xd7, 0x44, 0x43, 0xc3, 0x59, 0x25, 0x64, 0x9e, 0xe8, 0x6b, 0x56,
	0x87, 0x75, 0xa3, 0x40, 0xe1, 0x85, 0x31, 0x9c, 0x9e, 0xe4, 0x2a, 0x57, 0x86, 0x89, 0xc6, 0x69,
	0x12, 0x9f, 0x7c, 0x77, 0xd0, 0x83, 0xf3, 0xed, 0xd6, 0xf9, 0x35, 0xab, 0xf1, 0x07, 0xb4, 0x80,
	0x2e, 0x11, 0x19, 0xb1, 0x7c, 0x2b, 0x58, 0x9e, 0xbd, 0x1a, 0x7a, 0xcf, 0xb9, 0xe8, 0xe2, 0x77,
	0xbf, 0x7b, 0xef, 0x59, 0x2e, 0xa0, 0x64, 0x97, 0x61, 0xaa, 0xaa, 0x9d, 0xe7, 0x8f, 0x93, 0x54,
	0x19, 0x8f, 0x52, 0x55, 0x55, 0x4a, 0x86, 0xa3, 0x99, 0x3a, 0xd0, 0xc5, 0x19, 0x3e, 0x45, 0x47,
	0x42, 0x02, 0x6f, 0xae, 0x58, 0x49, 0xe6, 0xbe, 0x15, 0x38, 0xf4, 0x1f, 0x1e, 0xb5, 0xaf, 0x2d,
	0x93, 0x20, 0x60, 0x4d, 0xec, 0x49, 0xdb, 0x62, 0x7c, 0x82, 0x16, 0xa9, 0x6a, 0x25, 0x10, 0xc7,
	0x08, 0x13, 0xc0, 0x1e, 0x5a, 0x95, 0x4c, 0x43, 0x52, 0x70, 0x91, 0x17, 0x40, 0x16, 0xbe, 0x15,
	0xd8, 0x14, 0x8d, 0xd4, 0x7b, 0xc3, 0x60, 0x8a, 0x8e, 0xa1, 0x61, 0x19, 0x4f, 0x80, 0x35, 0x39,
	0x07, 0x72, 0x60, 0xbe, 0x20, 0xba, 0xe9, 0xbd, 0xd9, 0xcf, 0xde, 0x7b, 0x9a, 0x0b, 0x28, 0xda,
	0xe9, 0xed, 0x53, 0xa5, 0x2b, 0xa5, 0xff, 0xde, 0x9e, 0xeb, 0xec, 0xcb, 0xf4, 0x17, 0xc3, 0x8f,
	0x42, 0x02, 0x5d, 0x99, 0x90, 0x0b, 0x93, 0x81, 0x63, 0x74, 0x98, 0xf1, 0x5a, 0x69, 0x01, 0xe4,
	0x70, 0xbf, 0xb8, 0xed, 0x3e, 0x7e, 0x83, 0xe6, 0x42, 0x92, 0xa3, 0xfd, 0x52, 0xe6, 0x42, 0xe2,
	0xb7, 0xc8, 0x56, 0x2d, 0x90, 0xe5, 0x7e, 0x09, 0xe3, 0x2e, 0x7e, 0x8c, 0x8e, 0x3f, 0x33, 0x51,
	0xf2, 0xcc, 0x34, 0x44, 0x13, 0xe4, 0xdb, 0x81, 0x43, 0x57, 0x13, 0x37, 0x9e, 0xbf, 0xc6, 0x21,
	0x7a, 0xb4, 0x63, 0x49, 0x1a, 0xce, 0xb4, 0x92, 0x9a, 0xac, 0x7c, 0x3b, 0x58, 0xd2, 0x87, 0x77,
	0x4e, 0x3a, 0x09, 0x67, 0xf1, 0xcd, 0xe0, 0x5a, 0xb7, 0x83, 0x6b, 0xfd, 0x1a, 0x5c, 0xeb, 0xdb,
	0xc6, 0x9d, 0xdd, 0x6e, 0xdc, 0xd9, 0x8f, 0x8d, 0x3b, 0xfb, 0x14, 0xdd, 0xdf, 0x97, 0xff, 0x4a,
	0x7c, 0x79, 0x60, 0x4a, 0xf9, 0xf2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0xa6, 0x5d, 0x2d,
	0xed, 0x02, 0x00, 0x00,
}

func (m *StreamingSwap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamingSwap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamingSwap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FailedSwapReasons) > 0 {
		for iNdEx := len(m.FailedSwapReasons) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FailedSwapReasons[iNdEx])
			copy(dAtA[i:], m.FailedSwapReasons[iNdEx])
			i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(len(m.FailedSwapReasons[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.FailedSwaps) > 0 {
		dAtA2 := make([]byte, len(m.FailedSwaps)*10)
		var j1 int
		for _, num := range m.FailedSwaps {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x52
	}
	{
		size := m.Out.Size()
		i -= size
		if _, err := m.Out.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.In.Size()
		i -= size
		if _, err := m.In.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.Deposit.Size()
		i -= size
		if _, err := m.Deposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.TradeTarget.Size()
		i -= size
		if _, err := m.TradeTarget.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.LastHeight != 0 {
		i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(m.LastHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.Count != 0 {
		i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x20
	}
	if m.Quantity != 0 {
		i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(m.Quantity))
		i--
		dAtA[i] = 0x18
	}
	if m.Interval != 0 {
		i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(m.Interval))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TxID) > 0 {
		i -= len(m.TxID)
		copy(dAtA[i:], m.TxID)
		i = encodeVarintTypeStreamingSwap(dAtA, i, uint64(len(m.TxID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeStreamingSwap(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeStreamingSwap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StreamingSwap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxID)
	if l > 0 {
		n += 1 + l + sovTypeStreamingSwap(uint64(l))
	}
	if m.Interval != 0 {
		n += 1 + sovTypeStreamingSwap(uint64(m.Interval))
	}
	if m.Quantity != 0 {
		n += 1 + sovTypeStreamingSwap(uint64(m.Quantity))
	}
	if m.Count != 0 {
		n += 1 + sovTypeStreamingSwap(uint64(m.Count))
	}
	if m.LastHeight != 0 {
		n += 1 + sovTypeStreamingSwap(uint64(m.LastHeight))
	}
	l = m.TradeTarget.Size()
	n += 1 + l + sovTypeStreamingSwap(uint64(l))
	l = m.Deposit.Size()
	n += 1 + l + sovTypeStreamingSwap(uint64(l))
	l = m.In.Size()
	n += 1 + l + sovTypeStreamingSwap(uint64(l))
	l = m.Out.Size()
	n += 1 + l + sovTypeStreamingSwap(uint64(l))
	if len(m.FailedSwaps) > 0 {
		l = 0
		for _, e := range m.FailedSwaps {
			l += sovTypeStreamingSwap(uint64(e))
		}
		n += 1 + sovTypeStreamingSwap(uint64(l)) + l
	}
	if len(m.FailedSwapReasons) > 0 {
		for _, s := range m.FailedSwapReasons {
			l = len(s)
			n += 1 + l + sovTypeStreamingSwap(uint64(l))
		}
	}
	return n
}

func sovTypeStreamingSwap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeStreamingSwap(x uint64) (n int) {
	return sovTypeStreamingSwap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamingSwap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeStreamingSwap
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
			return fmt.Errorf("proto: StreamingSwap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamingSwap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeStreamingSwap
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
				return ErrInvalidLengthTypeStreamingSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeStreamingSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxID = gitlab_com_mayachain_mayanode_common.TxID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			m.Interval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeStreamingSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Interval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			m.Quantity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeStreamingSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quantity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeStreamingSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeight", wireType)
			}
			m.LastHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeStreamingSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeStreamingSwap
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
				return ErrInvalidLengthTypeStreamingSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeStreamingSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TradeTarget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeStreamingSwap
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
				return ErrInvalidLengthTypeStreamingSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeStreamingSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field In", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeStreamingSwap
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
				return ErrInvalidLengthTypeStreamingSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeStreamingSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.In.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Out", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeStreamingSwap
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
				return ErrInvalidLengthTypeStreamingSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeStreamingSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Out.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypeStreamingSwap
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.FailedSwaps = append(m.FailedSwaps, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypeStreamingSwap
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTypeStreamingSwap
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTypeStreamingSwap
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.FailedSwaps) == 0 {
					m.FailedSwaps = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypeStreamingSwap
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.FailedSwaps = append(m.FailedSwaps, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field FailedSwaps", wireType)
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailedSwapReasons", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeStreamingSwap
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
				return ErrInvalidLengthTypeStreamingSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeStreamingSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FailedSwapReasons = append(m.FailedSwapReasons, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeStreamingSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeStreamingSwap
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeStreamingSwap
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
func skipTypeStreamingSwap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeStreamingSwap
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
					return 0, ErrIntOverflowTypeStreamingSwap
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
					return 0, ErrIntOverflowTypeStreamingSwap
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
				return 0, ErrInvalidLengthTypeStreamingSwap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeStreamingSwap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeStreamingSwap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeStreamingSwap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeStreamingSwap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeStreamingSwap = fmt.Errorf("proto: unexpected end of group")
)
