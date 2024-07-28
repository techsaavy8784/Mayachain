// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/msg_reserve_contributor.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type MsgReserveContributor struct {
	Tx          common.Tx                                     `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx"`
	Contributor ReserveContributor                            `protobuf:"bytes,2,opt,name=contributor,proto3" json:"contributor"`
	Signer      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgReserveContributor) Reset()         { *m = MsgReserveContributor{} }
func (m *MsgReserveContributor) String() string { return proto.CompactTextString(m) }
func (*MsgReserveContributor) ProtoMessage()    {}
func (*MsgReserveContributor) Descriptor() ([]byte, []int) {
	return fileDescriptor_44afd1d80ec7ebe3, []int{0}
}
func (m *MsgReserveContributor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReserveContributor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReserveContributor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReserveContributor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReserveContributor.Merge(m, src)
}
func (m *MsgReserveContributor) XXX_Size() int {
	return m.Size()
}
func (m *MsgReserveContributor) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReserveContributor.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReserveContributor proto.InternalMessageInfo

func (m *MsgReserveContributor) GetTx() common.Tx {
	if m != nil {
		return m.Tx
	}
	return common.Tx{}
}

func (m *MsgReserveContributor) GetContributor() ReserveContributor {
	if m != nil {
		return m.Contributor
	}
	return ReserveContributor{}
}

func (m *MsgReserveContributor) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgReserveContributor)(nil), "types.MsgReserveContributor")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/msg_reserve_contributor.proto", fileDescriptor_44afd1d80ec7ebe3)
}

var fileDescriptor_44afd1d80ec7ebe3 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xc9, 0x4d, 0xac, 0x4c,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xaf, 0xd0, 0x47, 0x70, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xf5, 0x73, 0x8b, 0xd3, 0xe3, 0x8b, 0x52, 0x8b, 0x53, 0x8b, 0xca, 0x52, 0xe3, 0x93,
	0xf3, 0xf3, 0x4a, 0x8a, 0x32, 0x93, 0x4a, 0x4b, 0xf2, 0x8b, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x58, 0xc1, 0x8a, 0xa4, 0x6c, 0x09, 0x18, 0x02, 0x22, 0x71, 0x9b, 0x22, 0xa5, 0x80, 0xa2,
	0x3d, 0x39, 0x3f, 0x37, 0x37, 0x3f, 0x0f, 0x4a, 0x41, 0x55, 0x88, 0xa4, 0xe7, 0xa7, 0xe7, 0x83,
	0x99, 0xfa, 0x20, 0x16, 0x44, 0x54, 0xe9, 0x38, 0x23, 0x97, 0xa8, 0x6f, 0x71, 0x7a, 0x10, 0xc4,
	0x60, 0x67, 0x84, 0xb9, 0x42, 0x0a, 0x5c, 0x4c, 0x25, 0x15, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0x5c, 0x7a, 0x50, 0xa3, 0x42, 0x2a, 0x9c, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x62, 0x2a,
	0xa9, 0x10, 0x72, 0xe4, 0xe2, 0x46, 0x72, 0x88, 0x04, 0x13, 0x58, 0xa9, 0xa4, 0x1e, 0xd8, 0xbd,
	0x7a, 0x98, 0x26, 0x42, 0x75, 0x22, 0xeb, 0x11, 0xf2, 0xe4, 0x62, 0x2b, 0xce, 0x4c, 0xcf, 0x4b,
	0x2d, 0x92, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x71, 0x32, 0xfc, 0x75, 0x4f, 0x5e, 0x37, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x09, 0x64, 0xa5, 0x7e, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0x31, 0x94, 0xd2, 0x2d,
	0x4e, 0xc9, 0x86, 0x84, 0x86, 0x9e, 0x63, 0x72, 0xb2, 0x63, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0x71,
	0x10, 0xd4, 0x00, 0x27, 0xcf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2,
	0x4f, 0xcf, 0x2c, 0xc9, 0x49, 0x84, 0x18, 0x88, 0x08, 0x26, 0x10, 0x2b, 0x2f, 0x3f, 0x25, 0x15,
	0x33, 0xac, 0x93, 0xd8, 0xc0, 0x61, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xca, 0x72, 0x13,
	0x5c, 0xd9, 0x01, 0x00, 0x00,
}

func (m *MsgReserveContributor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReserveContributor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReserveContributor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgReserveContributor(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Contributor.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgReserveContributor(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgReserveContributor(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMsgReserveContributor(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgReserveContributor(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgReserveContributor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tx.Size()
	n += 1 + l + sovMsgReserveContributor(uint64(l))
	l = m.Contributor.Size()
	n += 1 + l + sovMsgReserveContributor(uint64(l))
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgReserveContributor(uint64(l))
	}
	return n
}

func sovMsgReserveContributor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgReserveContributor(x uint64) (n int) {
	return sovMsgReserveContributor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgReserveContributor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgReserveContributor
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
			return fmt.Errorf("proto: MsgReserveContributor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReserveContributor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgReserveContributor
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
				return ErrInvalidLengthMsgReserveContributor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgReserveContributor
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
				return fmt.Errorf("proto: wrong wireType = %d for field Contributor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgReserveContributor
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
				return ErrInvalidLengthMsgReserveContributor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgReserveContributor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Contributor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgReserveContributor
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
				return ErrInvalidLengthMsgReserveContributor
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgReserveContributor
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
			skippy, err := skipMsgReserveContributor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgReserveContributor
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgReserveContributor
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
func skipMsgReserveContributor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgReserveContributor
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
					return 0, ErrIntOverflowMsgReserveContributor
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
					return 0, ErrIntOverflowMsgReserveContributor
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
				return 0, ErrInvalidLengthMsgReserveContributor
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgReserveContributor
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgReserveContributor
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgReserveContributor        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgReserveContributor          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgReserveContributor = fmt.Errorf("proto: unexpected end of group")
)