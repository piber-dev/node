// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta1legacy/resource.proto

package legacy

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	legacy "github.com/ovrclk/akash/types/legacy"
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

// Resource stores unit, total count and price of resource
type Resource struct {
	Resources legacy.ResourceUnits `protobuf:"bytes,1,opt,name=resources,proto3" json:"unit" yaml:"unit"`
	Count     uint32               `protobuf:"varint,2,opt,name=count,proto3" json:"count" yaml:"count"`
	Price     types.Coin           `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9bb19359db18816, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetResources() legacy.ResourceUnits {
	if m != nil {
		return m.Resources
	}
	return legacy.ResourceUnits{}
}

func (m *Resource) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Resource) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Resource)(nil), "akash.deployment.v1beta1legacy.Resource")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta1legacy/resource.proto", fileDescriptor_d9bb19359db18816)
}

var fileDescriptor_d9bb19359db18816 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xff, 0x6f, 0x11, 0xa4, 0xb0, 0x44, 0x0c, 0x6d, 0x11, 0x4e, 0x95, 0x85, 0x0e,
	0x60, 0xab, 0xb0, 0x75, 0x0c, 0x23, 0x12, 0x43, 0x24, 0x16, 0xb6, 0xc4, 0x58, 0x69, 0xd4, 0x24,
	0x37, 0x8a, 0x9d, 0x8a, 0xbc, 0x05, 0x8f, 0xc0, 0xe3, 0x74, 0xec, 0xc8, 0x14, 0xa1, 0x76, 0x41,
	0x1d, 0x2b, 0xb1, 0xa3, 0xd8, 0xad, 0x42, 0x06, 0x36, 0xdf, 0xeb, 0xef, 0x9c, 0x7b, 0x7c, 0x6d,
	0xde, 0xf8, 0x73, 0x5f, 0xcc, 0xe8, 0x0b, 0xcf, 0x62, 0x28, 0x13, 0x9e, 0x4a, 0xba, 0x98, 0x04,
	0x5c, 0xfa, 0x93, 0x98, 0x87, 0x3e, 0x2b, 0x69, 0xce, 0x05, 0x14, 0x39, 0xe3, 0x24, 0xcb, 0x41,
	0x82, 0x85, 0x15, 0x4e, 0x1a, 0x9c, 0xb4, 0xf0, 0xe1, 0x79, 0x08, 0x21, 0x28, 0x94, 0xd6, 0x27,
	0xad, 0x1a, 0x5e, 0xeb, 0x21, 0x81, 0x2f, 0xf8, 0x1f, 0xf6, 0x45, 0x1a, 0x49, 0xb1, 0xa7, 0x31,
	0x03, 0x91, 0x80, 0x68, 0xe1, 0x94, 0x41, 0x94, 0xea, 0x7b, 0xe7, 0x1b, 0x99, 0xc7, 0xde, 0x5e,
	0x67, 0x05, 0xe6, 0xc9, 0xc1, 0x43, 0xf4, 0xd1, 0x08, 0x8d, 0x7b, 0xb7, 0x57, 0x44, 0x87, 0xac,
	0xf5, 0xed, 0x78, 0xe4, 0x20, 0x7b, 0xaa, 0xc7, 0xb9, 0x17, 0xcb, 0xca, 0x36, 0xb6, 0x95, 0xdd,
	0xa9, 0xa7, 0xef, 0x2a, 0xbb, 0x57, 0xfa, 0x49, 0x3c, 0x75, 0xea, 0xca, 0xf1, 0x1a, 0x5b, 0x8b,
	0x9a, 0x5d, 0x06, 0x45, 0x2a, 0xfb, 0xff, 0x46, 0x68, 0x7c, 0xe6, 0x0e, 0xb6, 0x95, 0xad, 0x1b,
	0xbb, 0xca, 0x3e, 0xd5, 0x1a, 0x55, 0x3a, 0x9e, 0x6e, 0x5b, 0x8f, 0x66, 0x37, 0xcb, 0x23, 0xc6,
	0xfb, 0xff, 0x55, 0xa0, 0x01, 0xd1, 0x2f, 0x6a, 0x25, 0x22, 0xf7, 0x10, 0xa5, 0xee, 0xe5, 0x3e,
	0x82, 0xe6, 0x1b, 0x3f, 0x55, 0x3a, 0x9e, 0x6e, 0x4f, 0x3b, 0x5f, 0xef, 0xb6, 0xe1, 0x3e, 0x2c,
	0xd7, 0x18, 0xad, 0xd6, 0x18, 0x7d, 0xae, 0x31, 0x7a, 0xdb, 0x60, 0x63, 0xb5, 0xc1, 0xc6, 0xc7,
	0x06, 0x1b, 0xcf, 0x93, 0x30, 0x92, 0xb3, 0x22, 0x20, 0x0c, 0x12, 0x0a, 0x8b, 0x9c, 0xc5, 0x73,
	0xaa, 0x37, 0xfe, 0xfa, 0xfb, 0x63, 0x65, 0x99, 0x71, 0x41, 0xf5, 0x26, 0x82, 0x23, 0xb5, 0xcb,
	0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x1c, 0xb8, 0x76, 0x00, 0x02, 0x00, 0x00,
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintResource(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Count != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Resources.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintResource(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	offset -= sovResource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Resources.Size()
	n += 1 + l + sovResource(uint64(l))
	if m.Count != 0 {
		n += 1 + sovResource(uint64(m.Count))
	}
	l = m.Price.Size()
	n += 1 + l + sovResource(uint64(l))
	return n
}

func sovResource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resources.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
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
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
				return 0, ErrInvalidLengthResource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResource = fmt.Errorf("proto: unexpected end of group")
)
