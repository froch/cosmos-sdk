// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testpb/test.proto

package testpb

import (
	_ "cosmossdk.io/depinject/appconfig/v1alpha1"
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type TestModuleGogo struct {
}

func (m *TestModuleGogo) Reset()         { *m = TestModuleGogo{} }
func (m *TestModuleGogo) String() string { return proto.CompactTextString(m) }
func (*TestModuleGogo) ProtoMessage()    {}
func (*TestModuleGogo) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{0}
}
func (m *TestModuleGogo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestModuleGogo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestModuleGogo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestModuleGogo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestModuleGogo.Merge(m, src)
}
func (m *TestModuleGogo) XXX_Size() int {
	return m.Size()
}
func (m *TestModuleGogo) XXX_DiscardUnknown() {
	xxx_messageInfo_TestModuleGogo.DiscardUnknown(m)
}

var xxx_messageInfo_TestModuleGogo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TestModuleGogo)(nil), "testpb.TestModuleGogo")
}

func init() { proto.RegisterFile("testpb/test.proto", fileDescriptor_41c67e33ca9d1f26) }

var fileDescriptor_41c67e33ca9d1f26 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x49, 0x2d, 0x2e,
	0x29, 0x48, 0xd2, 0x07, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x21, 0x29,
	0x85, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0x62, 0xfd, 0xc4, 0x82, 0x02, 0xfd, 0x32, 0xc3, 0xc4, 0x9c,
	0x82, 0x8c, 0x44, 0x43, 0xfd, 0xdc, 0xfc, 0x94, 0xd2, 0x9c, 0x54, 0x88, 0x4a, 0x25, 0x4f, 0x2e,
	0xbe, 0x90, 0xd4, 0xe2, 0x12, 0x5f, 0xb0, 0x98, 0x7b, 0x7e, 0x7a, 0xbe, 0x95, 0xf9, 0xae, 0x03,
	0xd3, 0x6e, 0x31, 0x1a, 0x72, 0xe9, 0x43, 0xf4, 0x16, 0xa7, 0x64, 0xeb, 0x65, 0xe6, 0xeb, 0x27,
	0xe7, 0x17, 0xa5, 0xea, 0x67, 0xe6, 0x95, 0xa4, 0x16, 0xe5, 0x25, 0xe6, 0x80, 0xcc, 0x4b, 0xce,
	0xcf, 0x4b, 0xcb, 0x4c, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x87, 0x58, 0xe6, 0xe4, 0x77, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1,
	0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x26, 0x28, 0x46, 0xa5, 0xa4, 0x16, 0x64, 0xe6,
	0x65, 0xa5, 0x26, 0x97, 0xe0, 0x37, 0x2f, 0x89, 0x0d, 0xec, 0x42, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0xdb, 0x22, 0xe2, 0xe0, 0x00, 0x00, 0x00,
}

func (m *TestModuleGogo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestModuleGogo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestModuleGogo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TestModuleGogo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTest(x uint64) (n int) {
	return sovTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestModuleGogo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: TestModuleGogo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestModuleGogo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
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
func skipTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
				return 0, ErrInvalidLengthTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTest = fmt.Errorf("proto: unexpected end of group")
)
