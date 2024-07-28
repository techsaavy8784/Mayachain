// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/msg_set_node_keys.proto

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

type MsgSetNodeKeys struct {
	PubKeySetSet        common.PubKeySet                              `protobuf:"bytes,1,opt,name=pub_key_set_set,json=pubKeySetSet,proto3" json:"pub_key_set_set"`
	ValidatorConsPubKey string                                        `protobuf:"bytes,2,opt,name=validator_cons_pub_key,json=validatorConsPubKey,proto3" json:"validator_cons_pub_key,omitempty"`
	Signer              github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgSetNodeKeys) Reset()         { *m = MsgSetNodeKeys{} }
func (m *MsgSetNodeKeys) String() string { return proto.CompactTextString(m) }
func (*MsgSetNodeKeys) ProtoMessage()    {}
func (*MsgSetNodeKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_555203e8a5c8c348, []int{0}
}
func (m *MsgSetNodeKeys) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetNodeKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetNodeKeys.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetNodeKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetNodeKeys.Merge(m, src)
}
func (m *MsgSetNodeKeys) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetNodeKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetNodeKeys.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetNodeKeys proto.InternalMessageInfo

func (m *MsgSetNodeKeys) GetPubKeySetSet() common.PubKeySet {
	if m != nil {
		return m.PubKeySetSet
	}
	return common.PubKeySet{}
}

func (m *MsgSetNodeKeys) GetValidatorConsPubKey() string {
	if m != nil {
		return m.ValidatorConsPubKey
	}
	return ""
}

func (m *MsgSetNodeKeys) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSetNodeKeys)(nil), "types.MsgSetNodeKeys")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/msg_set_node_keys.proto", fileDescriptor_555203e8a5c8c348)
}

var fileDescriptor_555203e8a5c8c348 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xb3, 0xef, 0xab, 0x05, 0x63, 0x51, 0x8c, 0x22, 0xa5, 0x87, 0x34, 0x78, 0xea, 0xa5,
	0x5d, 0x6a, 0xc1, 0xa3, 0xd0, 0x7a, 0x2a, 0x45, 0x91, 0xf6, 0xe6, 0x25, 0x24, 0xbb, 0xc3, 0x36,
	0xb4, 0xd9, 0x09, 0x99, 0x6d, 0x31, 0xdf, 0xc2, 0x8f, 0xd5, 0x63, 0x8f, 0x9e, 0x8a, 0xb4, 0xdf,
	0xc2, 0x93, 0xe4, 0x8f, 0x16, 0xf1, 0x34, 0x4f, 0x66, 0xb2, 0xbf, 0x07, 0x7e, 0xf6, 0x5d, 0x1c,
	0x64, 0x81, 0x98, 0x05, 0x91, 0xe6, 0xab, 0x1e, 0x7f, 0xe5, 0x87, 0x4f, 0x93, 0x25, 0x40, 0x3c,
	0x26, 0xe5, 0x13, 0x18, 0x5f, 0xa3, 0x04, 0x7f, 0x0e, 0x19, 0x75, 0x93, 0x14, 0x0d, 0x3a, 0xc7,
	0xc5, 0xb9, 0xe9, 0xfd, 0x7a, 0x2e, 0x30, 0x8e, 0x51, 0x57, 0xa3, 0xfc, 0xb1, 0x79, 0xa5, 0x50,
	0x61, 0x11, 0x79, 0x9e, 0xca, 0xed, 0xcd, 0x86, 0xd9, 0x67, 0x8f, 0xa4, 0xa6, 0x60, 0x9e, 0x50,
	0xc2, 0x18, 0x32, 0x72, 0xee, 0xed, 0xf3, 0x64, 0x19, 0xe6, 0x1d, 0x45, 0x21, 0x81, 0x69, 0x30,
	0x8f, 0xb5, 0x4f, 0x6f, 0x2f, 0xba, 0x15, 0xf0, 0x79, 0x19, 0x8e, 0x21, 0x9b, 0x82, 0x19, 0x1e,
	0xad, 0xb7, 0x2d, 0x6b, 0x52, 0x4f, 0xbe, 0x17, 0x53, 0x30, 0x4e, 0xdf, 0xbe, 0x5e, 0x05, 0x8b,
	0x48, 0x06, 0x06, 0x53, 0x5f, 0xa0, 0x26, 0xbf, 0xc2, 0x35, 0xfe, 0x79, 0xac, 0x7d, 0x32, 0xb9,
	0xfc, 0xb9, 0x3e, 0xa0, 0xa6, 0x92, 0xe5, 0x8c, 0xec, 0x1a, 0x45, 0x4a, 0x43, 0xda, 0xf8, 0xef,
	0xb1, 0x76, 0x7d, 0xd8, 0xfb, 0xdc, 0xb6, 0x3a, 0x2a, 0x32, 0xb3, 0x65, 0x98, 0xb7, 0x72, 0x81,
	0x14, 0x23, 0x55, 0xa3, 0x43, 0x72, 0x5e, 0x6a, 0xe9, 0x0e, 0x84, 0x18, 0x48, 0x99, 0x02, 0xd1,
	0xa4, 0x02, 0x0c, 0x47, 0xeb, 0x9d, 0xcb, 0x36, 0x3b, 0x97, 0x7d, 0xec, 0x5c, 0xf6, 0xb6, 0x77,
	0xad, 0xcd, 0xde, 0xb5, 0xde, 0xf7, 0xae, 0xf5, 0xc2, 0x55, 0x64, 0x16, 0x41, 0x09, 0x3c, 0xf8,
	0xca, 0x53, 0x6e, 0xf5, 0xaf, 0xf4, 0xb0, 0x56, 0x48, 0xea, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xcb, 0x05, 0xe9, 0x5b, 0x9d, 0x01, 0x00, 0x00,
}

func (m *MsgSetNodeKeys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetNodeKeys) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetNodeKeys) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgSetNodeKeys(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValidatorConsPubKey) > 0 {
		i -= len(m.ValidatorConsPubKey)
		copy(dAtA[i:], m.ValidatorConsPubKey)
		i = encodeVarintMsgSetNodeKeys(dAtA, i, uint64(len(m.ValidatorConsPubKey)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.PubKeySetSet.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgSetNodeKeys(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMsgSetNodeKeys(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgSetNodeKeys(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetNodeKeys) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PubKeySetSet.Size()
	n += 1 + l + sovMsgSetNodeKeys(uint64(l))
	l = len(m.ValidatorConsPubKey)
	if l > 0 {
		n += 1 + l + sovMsgSetNodeKeys(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgSetNodeKeys(uint64(l))
	}
	return n
}

func sovMsgSetNodeKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgSetNodeKeys(x uint64) (n int) {
	return sovMsgSetNodeKeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetNodeKeys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgSetNodeKeys
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
			return fmt.Errorf("proto: MsgSetNodeKeys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetNodeKeys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeySetSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSetNodeKeys
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
				return ErrInvalidLengthMsgSetNodeKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSetNodeKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PubKeySetSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorConsPubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSetNodeKeys
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
				return ErrInvalidLengthMsgSetNodeKeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSetNodeKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorConsPubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSetNodeKeys
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
				return ErrInvalidLengthMsgSetNodeKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSetNodeKeys
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
			skippy, err := skipMsgSetNodeKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgSetNodeKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgSetNodeKeys
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
func skipMsgSetNodeKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgSetNodeKeys
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
					return 0, ErrIntOverflowMsgSetNodeKeys
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
					return 0, ErrIntOverflowMsgSetNodeKeys
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
				return 0, ErrInvalidLengthMsgSetNodeKeys
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgSetNodeKeys
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgSetNodeKeys
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgSetNodeKeys        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgSetNodeKeys          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgSetNodeKeys = fmt.Errorf("proto: unexpected end of group")
)