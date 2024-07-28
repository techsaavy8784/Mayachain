// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/type_tss_metric.proto

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

type NodeTssTime struct {
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	TssTime int64                                         `protobuf:"varint,2,opt,name=tss_time,json=tssTime,proto3" json:"tss_time,omitempty"`
}

func (m *NodeTssTime) Reset()         { *m = NodeTssTime{} }
func (m *NodeTssTime) String() string { return proto.CompactTextString(m) }
func (*NodeTssTime) ProtoMessage()    {}
func (*NodeTssTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_7819b12199d37add, []int{0}
}
func (m *NodeTssTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeTssTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeTssTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeTssTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeTssTime.Merge(m, src)
}
func (m *NodeTssTime) XXX_Size() int {
	return m.Size()
}
func (m *NodeTssTime) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeTssTime.DiscardUnknown(m)
}

var xxx_messageInfo_NodeTssTime proto.InternalMessageInfo

func (m *NodeTssTime) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *NodeTssTime) GetTssTime() int64 {
	if m != nil {
		return m.TssTime
	}
	return 0
}

type TssKeygenMetric struct {
	PubKey       gitlab_com_mayachain_mayanode_common.PubKey `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3,casttype=gitlab.com/mayachain/mayanode/common.PubKey" json:"pub_key,omitempty"`
	NodeTssTimes []NodeTssTime                               `protobuf:"bytes,2,rep,name=node_tss_times,json=nodeTssTimes,proto3" json:"node_tss_times"`
}

func (m *TssKeygenMetric) Reset()         { *m = TssKeygenMetric{} }
func (m *TssKeygenMetric) String() string { return proto.CompactTextString(m) }
func (*TssKeygenMetric) ProtoMessage()    {}
func (*TssKeygenMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_7819b12199d37add, []int{1}
}
func (m *TssKeygenMetric) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TssKeygenMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TssKeygenMetric.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TssKeygenMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TssKeygenMetric.Merge(m, src)
}
func (m *TssKeygenMetric) XXX_Size() int {
	return m.Size()
}
func (m *TssKeygenMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_TssKeygenMetric.DiscardUnknown(m)
}

var xxx_messageInfo_TssKeygenMetric proto.InternalMessageInfo

func (m *TssKeygenMetric) GetPubKey() gitlab_com_mayachain_mayanode_common.PubKey {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *TssKeygenMetric) GetNodeTssTimes() []NodeTssTime {
	if m != nil {
		return m.NodeTssTimes
	}
	return nil
}

type TssKeysignMetric struct {
	TxID         gitlab_com_mayachain_mayanode_common.TxID `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,casttype=gitlab.com/mayachain/mayanode/common.TxID" json:"tx_id,omitempty"`
	NodeTssTimes []NodeTssTime                             `protobuf:"bytes,2,rep,name=node_tss_times,json=nodeTssTimes,proto3" json:"node_tss_times"`
}

func (m *TssKeysignMetric) Reset()         { *m = TssKeysignMetric{} }
func (m *TssKeysignMetric) String() string { return proto.CompactTextString(m) }
func (*TssKeysignMetric) ProtoMessage()    {}
func (*TssKeysignMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_7819b12199d37add, []int{2}
}
func (m *TssKeysignMetric) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TssKeysignMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TssKeysignMetric.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TssKeysignMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TssKeysignMetric.Merge(m, src)
}
func (m *TssKeysignMetric) XXX_Size() int {
	return m.Size()
}
func (m *TssKeysignMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_TssKeysignMetric.DiscardUnknown(m)
}

var xxx_messageInfo_TssKeysignMetric proto.InternalMessageInfo

func (m *TssKeysignMetric) GetTxID() gitlab_com_mayachain_mayanode_common.TxID {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *TssKeysignMetric) GetNodeTssTimes() []NodeTssTime {
	if m != nil {
		return m.NodeTssTimes
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeTssTime)(nil), "types.NodeTssTime")
	proto.RegisterType((*TssKeygenMetric)(nil), "types.TssKeygenMetric")
	proto.RegisterType((*TssKeysignMetric)(nil), "types.TssKeysignMetric")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/type_tss_metric.proto", fileDescriptor_7819b12199d37add)
}

var fileDescriptor_7819b12199d37add = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0xb3, 0x7d, 0x7d, 0x9e, 0x6d, 0x51, 0x09, 0x1e, 0xaa, 0x87, 0xa4, 0xf4, 0x54, 0x91,
	0x66, 0xa9, 0x2f, 0x57, 0xa1, 0xc1, 0x83, 0xb5, 0x28, 0x12, 0x72, 0xf2, 0x12, 0xf2, 0xb2, 0xa4,
	0x4b, 0xdd, 0x6c, 0xe8, 0x6e, 0x24, 0xf9, 0x16, 0xde, 0x05, 0x3f, 0x4f, 0x8f, 0x3d, 0x7a, 0x0a,
	0x92, 0x7e, 0x8b, 0x9e, 0x24, 0x49, 0x4b, 0x05, 0x41, 0x04, 0x2f, 0x3b, 0xb3, 0xf3, 0xdf, 0x9d,
	0xff, 0x6f, 0x60, 0xe0, 0x05, 0xb5, 0x13, 0xdb, 0x9d, 0xda, 0x24, 0x40, 0xcf, 0x43, 0x14, 0xa3,
	0xdd, 0x55, 0x24, 0x21, 0xe6, 0xc5, 0x69, 0x09, 0xce, 0x2d, 0x8a, 0xc5, 0x9c, 0xb8, 0x5a, 0x38,
	0x67, 0x82, 0xc9, 0xf5, 0x42, 0x3c, 0x3e, 0xf4, 0x99, 0xcf, 0x8a, 0x0a, 0xca, 0xb3, 0x52, 0xec,
	0x45, 0xb0, 0x75, 0xcf, 0x3c, 0x6c, 0x72, 0x6e, 0x12, 0x8a, 0xe5, 0x09, 0x6c, 0xda, 0x9e, 0x37,
	0xc7, 0x9c, 0x77, 0x40, 0x17, 0xf4, 0xdb, 0xfa, 0x70, 0x9d, 0xaa, 0x03, 0x9f, 0x88, 0x69, 0xe4,
	0x68, 0x2e, 0xa3, 0xc8, 0x65, 0x9c, 0x32, 0xbe, 0x09, 0x03, 0xee, 0xcd, 0x4a, 0x6b, 0x6d, 0xe4,
	0xba, 0xa3, 0xf2, 0xa3, 0xb1, 0xed, 0x20, 0x1f, 0xc1, 0x7f, 0x39, 0x8c, 0x20, 0x14, 0x77, 0x2a,
	0x5d, 0xd0, 0xaf, 0x1a, 0x4d, 0x51, 0xfa, 0xf4, 0x5e, 0x01, 0xdc, 0x37, 0x39, 0x9f, 0xe0, 0xc4,
	0xc7, 0xc1, 0x5d, 0x41, 0x2b, 0xdf, 0xc0, 0x66, 0x18, 0x39, 0xd6, 0x0c, 0x27, 0x85, 0xf7, 0x7f,
	0x1d, 0xad, 0x53, 0xf5, 0xd4, 0x27, 0xe2, 0xc9, 0x2e, 0xbd, 0x77, 0xe3, 0xe6, 0x59, 0xc0, 0x3c,
	0x8c, 0x5c, 0x46, 0x29, 0x0b, 0xb4, 0x87, 0xc8, 0x99, 0xe0, 0xc4, 0x68, 0x84, 0x45, 0x94, 0xaf,
	0xe0, 0x5e, 0xae, 0x5a, 0x5b, 0x77, 0xde, 0xa9, 0x74, 0xab, 0xfd, 0xd6, 0x99, 0xac, 0x95, 0xb0,
	0x5f, 0x26, 0xd6, 0x6b, 0x8b, 0x54, 0x95, 0x8c, 0x76, 0xb0, 0x2b, 0xf1, 0xde, 0x1b, 0x80, 0x07,
	0x25, 0x1d, 0x27, 0xfe, 0x16, 0xef, 0x16, 0xd6, 0x45, 0x6c, 0x11, 0x6f, 0x03, 0x77, 0x99, 0xa5,
	0x6a, 0xcd, 0x8c, 0xc7, 0xd7, 0xeb, 0x54, 0x3d, 0xf9, 0x15, 0x64, 0xfe, 0xd8, 0xa8, 0x89, 0x78,
	0xec, 0xfd, 0x15, 0x50, 0x1f, 0x2f, 0x32, 0x05, 0x2c, 0x33, 0x05, 0x7c, 0x64, 0x0a, 0x78, 0x59,
	0x29, 0xd2, 0x72, 0xa5, 0x48, 0xef, 0x2b, 0x45, 0x7a, 0x44, 0x3f, 0xa3, 0x7c, 0xdb, 0x19, 0xa7,
	0x51, 0xec, 0xc1, 0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x5b, 0x3f, 0x6b, 0x5c, 0x02,
	0x00, 0x00,
}

func (m *NodeTssTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeTssTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeTssTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TssTime != 0 {
		i = encodeVarintTypeTssMetric(dAtA, i, uint64(m.TssTime))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypeTssMetric(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TssKeygenMetric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TssKeygenMetric) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TssKeygenMetric) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NodeTssTimes) > 0 {
		for iNdEx := len(m.NodeTssTimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodeTssTimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypeTssMetric(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintTypeTssMetric(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TssKeysignMetric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TssKeysignMetric) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TssKeysignMetric) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NodeTssTimes) > 0 {
		for iNdEx := len(m.NodeTssTimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodeTssTimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypeTssMetric(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.TxID) > 0 {
		i -= len(m.TxID)
		copy(dAtA[i:], m.TxID)
		i = encodeVarintTypeTssMetric(dAtA, i, uint64(len(m.TxID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeTssMetric(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeTssMetric(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NodeTssTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypeTssMetric(uint64(l))
	}
	if m.TssTime != 0 {
		n += 1 + sovTypeTssMetric(uint64(m.TssTime))
	}
	return n
}

func (m *TssKeygenMetric) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovTypeTssMetric(uint64(l))
	}
	if len(m.NodeTssTimes) > 0 {
		for _, e := range m.NodeTssTimes {
			l = e.Size()
			n += 1 + l + sovTypeTssMetric(uint64(l))
		}
	}
	return n
}

func (m *TssKeysignMetric) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxID)
	if l > 0 {
		n += 1 + l + sovTypeTssMetric(uint64(l))
	}
	if len(m.NodeTssTimes) > 0 {
		for _, e := range m.NodeTssTimes {
			l = e.Size()
			n += 1 + l + sovTypeTssMetric(uint64(l))
		}
	}
	return n
}

func sovTypeTssMetric(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeTssMetric(x uint64) (n int) {
	return sovTypeTssMetric(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NodeTssTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeTssMetric
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
			return fmt.Errorf("proto: NodeTssTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeTssTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTssMetric
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
				return ErrInvalidLengthTypeTssMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTssMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TssTime", wireType)
			}
			m.TssTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTssMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TssTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypeTssMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeTssMetric
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeTssMetric
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
func (m *TssKeygenMetric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeTssMetric
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
			return fmt.Errorf("proto: TssKeygenMetric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TssKeygenMetric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTssMetric
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
				return ErrInvalidLengthTypeTssMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTssMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = gitlab_com_mayachain_mayanode_common.PubKey(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeTssTimes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTssMetric
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
				return ErrInvalidLengthTypeTssMetric
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTssMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeTssTimes = append(m.NodeTssTimes, NodeTssTime{})
			if err := m.NodeTssTimes[len(m.NodeTssTimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeTssMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeTssMetric
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeTssMetric
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
func (m *TssKeysignMetric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeTssMetric
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
			return fmt.Errorf("proto: TssKeysignMetric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TssKeysignMetric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTssMetric
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
				return ErrInvalidLengthTypeTssMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTssMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxID = gitlab_com_mayachain_mayanode_common.TxID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeTssTimes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTssMetric
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
				return ErrInvalidLengthTypeTssMetric
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTssMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeTssTimes = append(m.NodeTssTimes, NodeTssTime{})
			if err := m.NodeTssTimes[len(m.NodeTssTimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeTssMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeTssMetric
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeTssMetric
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
func skipTypeTssMetric(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeTssMetric
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
					return 0, ErrIntOverflowTypeTssMetric
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
					return 0, ErrIntOverflowTypeTssMetric
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
				return 0, ErrInvalidLengthTypeTssMetric
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeTssMetric
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeTssMetric
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeTssMetric        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeTssMetric          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeTssMetric = fmt.Errorf("proto: unexpected end of group")
)
