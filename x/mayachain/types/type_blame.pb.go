// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mayachain/v1/x/mayachain/types/type_blame.proto

package types

import (
	fmt "fmt"
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

type Node struct {
	Pubkey         string `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	BlameData      []byte `protobuf:"bytes,2,opt,name=blame_data,json=blameData,proto3" json:"blame_data,omitempty"`
	BlameSignature []byte `protobuf:"bytes,3,opt,name=blame_signature,json=blameSignature,proto3" json:"blame_signature,omitempty"`
}

func (m *Node) Reset()      { *m = Node{} }
func (*Node) ProtoMessage() {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bd05d18f755d777, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

type Blame struct {
	FailReason string `protobuf:"bytes,1,opt,name=fail_reason,json=failReason,proto3" json:"fail_reason,omitempty"`
	IsUnicast  bool   `protobuf:"varint,2,opt,name=is_unicast,json=isUnicast,proto3" json:"is_unicast,omitempty"`
	BlameNodes []Node `protobuf:"bytes,3,rep,name=blame_nodes,json=blameNodes,proto3" json:"blame_nodes"`
	Round      string `protobuf:"bytes,4,opt,name=round,proto3" json:"round,omitempty"`
}

func (m *Blame) Reset()      { *m = Blame{} }
func (*Blame) ProtoMessage() {}
func (*Blame) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bd05d18f755d777, []int{1}
}
func (m *Blame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Blame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Blame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Blame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blame.Merge(m, src)
}
func (m *Blame) XXX_Size() int {
	return m.Size()
}
func (m *Blame) XXX_DiscardUnknown() {
	xxx_messageInfo_Blame.DiscardUnknown(m)
}

var xxx_messageInfo_Blame proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Node)(nil), "types.Node")
	proto.RegisterType((*Blame)(nil), "types.Blame")
}

func init() {
	proto.RegisterFile("mayachain/v1/x/mayachain/types/type_blame.proto", fileDescriptor_0bd05d18f755d777)
}

var fileDescriptor_0bd05d18f755d777 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xbf, 0x6a, 0xf3, 0x30,
	0x10, 0xc0, 0xa5, 0x2f, 0x7f, 0xf8, 0x22, 0x97, 0x16, 0x44, 0x28, 0xa6, 0x50, 0x25, 0x64, 0x69,
	0x26, 0x9b, 0xa6, 0x6f, 0x10, 0x3a, 0x77, 0x70, 0xc9, 0xd2, 0xc5, 0x9c, 0x63, 0xc5, 0x15, 0x4d,
	0xac, 0x60, 0xc9, 0xa5, 0xd9, 0xfa, 0x08, 0xed, 0x5b, 0x65, 0xcc, 0x98, 0xa9, 0x34, 0xf6, 0x8b,
	0x14, 0x9d, 0x53, 0x08, 0x74, 0x11, 0x77, 0xbf, 0x13, 0xfa, 0xdd, 0x9d, 0x58, 0xb8, 0x82, 0x0d,
	0xcc, 0x9f, 0x41, 0xe5, 0xe1, 0xeb, 0x6d, 0xf8, 0x76, 0x92, 0xda, 0xcd, 0x5a, 0x1a, 0x3c, 0xe3,
	0x64, 0x09, 0x2b, 0x19, 0xac, 0x0b, 0x6d, 0x35, 0xef, 0x20, 0xbf, 0xea, 0x67, 0x3a, 0xd3, 0x48,
	0x42, 0x17, 0x35, 0xc5, 0xd1, 0x82, 0xb5, 0x1f, 0x74, 0x2a, 0xf9, 0x25, 0xeb, 0xae, 0xcb, 0xe4,
	0x45, 0x6e, 0x7c, 0x3a, 0xa4, 0xe3, 0x5e, 0x74, 0xcc, 0xf8, 0x35, 0x63, 0xf8, 0x56, 0x9c, 0x82,
	0x05, 0xff, 0xdf, 0x90, 0x8e, 0xcf, 0xa2, 0x1e, 0x92, 0x7b, 0xb0, 0xc0, 0x6f, 0xd8, 0x45, 0x53,
	0x36, 0x2a, 0xcb, 0xc1, 0x96, 0x85, 0xf4, 0x5b, 0x78, 0xe7, 0x1c, 0xf1, 0xe3, 0x2f, 0x1d, 0x7d,
	0x52, 0xd6, 0x99, 0x3a, 0xc4, 0x07, 0xcc, 0x5b, 0x80, 0x5a, 0xc6, 0x85, 0x04, 0xa3, 0xf3, 0xa3,
	0x8e, 0x39, 0x14, 0x21, 0x71, 0x4a, 0x65, 0xe2, 0x32, 0x57, 0x73, 0x30, 0x16, 0x95, 0xff, 0xa3,
	0x9e, 0x32, 0xb3, 0x06, 0xf0, 0x09, 0xf3, 0x1a, 0x65, 0xae, 0x53, 0x69, 0xfc, 0xd6, 0xb0, 0x35,
	0xf6, 0x26, 0x5e, 0x80, 0x43, 0x06, 0x6e, 0x96, 0x69, 0x7b, 0xfb, 0x35, 0x20, 0x51, 0xd3, 0xb7,
	0x03, 0x86, 0xf7, 0x59, 0xa7, 0xd0, 0x65, 0x9e, 0xfa, 0x6d, 0xb4, 0x35, 0xc9, 0x74, 0xb6, 0x3d,
	0x08, 0xb2, 0x3f, 0x08, 0xf2, 0x5e, 0x09, 0xb2, 0xad, 0x04, 0xdd, 0x55, 0x82, 0x7e, 0x57, 0x82,
	0x7e, 0xd4, 0x82, 0xec, 0x6a, 0x41, 0xf6, 0xb5, 0x20, 0x4f, 0x61, 0xa6, 0xec, 0x12, 0x92, 0x60,
	0xae, 0x57, 0x27, 0xbb, 0x76, 0x91, 0x6b, 0xe0, 0xef, 0x07, 0x24, 0x5d, 0xdc, 0xec, 0xdd, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xc9, 0x3e, 0x06, 0xa9, 0x01, 0x00, 0x00,
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlameSignature) > 0 {
		i -= len(m.BlameSignature)
		copy(dAtA[i:], m.BlameSignature)
		i = encodeVarintTypeBlame(dAtA, i, uint64(len(m.BlameSignature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BlameData) > 0 {
		i -= len(m.BlameData)
		copy(dAtA[i:], m.BlameData)
		i = encodeVarintTypeBlame(dAtA, i, uint64(len(m.BlameData)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintTypeBlame(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Blame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Blame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Blame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Round) > 0 {
		i -= len(m.Round)
		copy(dAtA[i:], m.Round)
		i = encodeVarintTypeBlame(dAtA, i, uint64(len(m.Round)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BlameNodes) > 0 {
		for iNdEx := len(m.BlameNodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlameNodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypeBlame(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.IsUnicast {
		i--
		if m.IsUnicast {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.FailReason) > 0 {
		i -= len(m.FailReason)
		copy(dAtA[i:], m.FailReason)
		i = encodeVarintTypeBlame(dAtA, i, uint64(len(m.FailReason)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeBlame(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeBlame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovTypeBlame(uint64(l))
	}
	l = len(m.BlameData)
	if l > 0 {
		n += 1 + l + sovTypeBlame(uint64(l))
	}
	l = len(m.BlameSignature)
	if l > 0 {
		n += 1 + l + sovTypeBlame(uint64(l))
	}
	return n
}

func (m *Blame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FailReason)
	if l > 0 {
		n += 1 + l + sovTypeBlame(uint64(l))
	}
	if m.IsUnicast {
		n += 2
	}
	if len(m.BlameNodes) > 0 {
		for _, e := range m.BlameNodes {
			l = e.Size()
			n += 1 + l + sovTypeBlame(uint64(l))
		}
	}
	l = len(m.Round)
	if l > 0 {
		n += 1 + l + sovTypeBlame(uint64(l))
	}
	return n
}

func sovTypeBlame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeBlame(x uint64) (n int) {
	return sovTypeBlame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeBlame
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeBlame
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
				return ErrInvalidLengthTypeBlame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeBlame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlameData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeBlame
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
				return ErrInvalidLengthTypeBlame
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeBlame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlameData = append(m.BlameData[:0], dAtA[iNdEx:postIndex]...)
			if m.BlameData == nil {
				m.BlameData = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlameSignature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeBlame
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
				return ErrInvalidLengthTypeBlame
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeBlame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlameSignature = append(m.BlameSignature[:0], dAtA[iNdEx:postIndex]...)
			if m.BlameSignature == nil {
				m.BlameSignature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeBlame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeBlame
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeBlame
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
func (m *Blame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeBlame
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
			return fmt.Errorf("proto: Blame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Blame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeBlame
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
				return ErrInvalidLengthTypeBlame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeBlame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FailReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsUnicast", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeBlame
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
			m.IsUnicast = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlameNodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeBlame
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
				return ErrInvalidLengthTypeBlame
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeBlame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlameNodes = append(m.BlameNodes, Node{})
			if err := m.BlameNodes[len(m.BlameNodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeBlame
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
				return ErrInvalidLengthTypeBlame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeBlame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Round = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeBlame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeBlame
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeBlame
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
func skipTypeBlame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeBlame
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
					return 0, ErrIntOverflowTypeBlame
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
					return 0, ErrIntOverflowTypeBlame
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
				return 0, ErrInvalidLengthTypeBlame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeBlame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeBlame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeBlame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeBlame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeBlame = fmt.Errorf("proto: unexpected end of group")
)
