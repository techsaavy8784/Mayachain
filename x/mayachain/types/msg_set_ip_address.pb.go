// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/msg_set_ip_address.proto

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

type MsgSetIPAddress struct {
	IPAddress string                                        `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Signer    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgSetIPAddress) Reset()         { *m = MsgSetIPAddress{} }
func (m *MsgSetIPAddress) String() string { return proto.CompactTextString(m) }
func (*MsgSetIPAddress) ProtoMessage()    {}
func (*MsgSetIPAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_87406621ca648130, []int{0}
}
func (m *MsgSetIPAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetIPAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetIPAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetIPAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetIPAddress.Merge(m, src)
}
func (m *MsgSetIPAddress) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetIPAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetIPAddress.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetIPAddress proto.InternalMessageInfo

func (m *MsgSetIPAddress) GetIPAddress() string {
	if m != nil {
		return m.IPAddress
	}
	return ""
}

func (m *MsgSetIPAddress) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSetIPAddress)(nil), "types.MsgSetIPAddress")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/msg_set_ip_address.proto", fileDescriptor_87406621ca648130)
}

var fileDescriptor_87406621ca648130 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xcf, 0x4d, 0xac, 0x4c,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xaf, 0xd0, 0x47, 0x70, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xf5, 0x73, 0x8b, 0xd3, 0xe3, 0x8b, 0x53, 0x4b, 0xe2, 0x33, 0x0b, 0xe2, 0x13, 0x53,
	0x52, 0x8a, 0x52, 0x8b, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0xc1, 0xf2, 0x52,
	0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x11, 0x7d, 0x10, 0x0b, 0x22, 0xa9, 0xd4, 0xc5, 0xc8, 0xc5,
	0xef, 0x5b, 0x9c, 0x1e, 0x9c, 0x5a, 0xe2, 0x19, 0xe0, 0x08, 0xd1, 0x26, 0xa4, 0xc3, 0xc5, 0x85,
	0x30, 0x44, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x89, 0xf7, 0xd1, 0x3d, 0x79, 0x4e, 0xb8, 0x92,
	0x20, 0xce, 0xcc, 0x02, 0x98, 0x6a, 0x4f, 0x2e, 0xb6, 0xe2, 0xcc, 0xf4, 0xbc, 0xd4, 0x22, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0x1e, 0x27, 0xc3, 0x5f, 0xf7, 0xe4, 0x75, 0xd3, 0x33, 0x4b, 0x32, 0x4a,
	0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xa1, 0x94, 0x6e, 0x71,
	0x4a, 0x36, 0xc4, 0xbd, 0x7a, 0x8e, 0xc9, 0xc9, 0x30, 0xd3, 0xa0, 0x06, 0x38, 0x79, 0x9e, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c,
	0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x7e, 0x7a, 0x66, 0x49, 0x4e, 0x22, 0xc4,
	0x40, 0x84, 0xc7, 0x41, 0xac, 0xbc, 0xfc, 0x94, 0x54, 0xcc, 0xd0, 0x48, 0x62, 0x03, 0x7b, 0xcf,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x94, 0x82, 0x81, 0xb5, 0x36, 0x01, 0x00, 0x00,
}

func (m *MsgSetIPAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetIPAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetIPAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgSetIpAddress(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IPAddress) > 0 {
		i -= len(m.IPAddress)
		copy(dAtA[i:], m.IPAddress)
		i = encodeVarintMsgSetIpAddress(dAtA, i, uint64(len(m.IPAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgSetIpAddress(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgSetIpAddress(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetIPAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IPAddress)
	if l > 0 {
		n += 1 + l + sovMsgSetIpAddress(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgSetIpAddress(uint64(l))
	}
	return n
}

func sovMsgSetIpAddress(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgSetIpAddress(x uint64) (n int) {
	return sovMsgSetIpAddress(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetIPAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgSetIpAddress
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
			return fmt.Errorf("proto: MsgSetIPAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetIPAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSetIpAddress
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
				return ErrInvalidLengthMsgSetIpAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSetIpAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSetIpAddress
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
				return ErrInvalidLengthMsgSetIpAddress
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSetIpAddress
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
			skippy, err := skipMsgSetIpAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgSetIpAddress
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgSetIpAddress
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
func skipMsgSetIpAddress(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgSetIpAddress
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
					return 0, ErrIntOverflowMsgSetIpAddress
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
					return 0, ErrIntOverflowMsgSetIpAddress
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
				return 0, ErrInvalidLengthMsgSetIpAddress
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgSetIpAddress
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgSetIpAddress
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgSetIpAddress        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgSetIpAddress          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgSetIpAddress = fmt.Errorf("proto: unexpected end of group")
)
