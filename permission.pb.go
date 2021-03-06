// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: permission.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Permission struct {
	Action          string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	ResourcePattern string `protobuf:"bytes,3,opt,name=resource_pattern,json=resourcePattern,proto3" json:"resource_pattern,omitempty"`
}

func (m *Permission) Reset()                    { *m = Permission{} }
func (m *Permission) String() string            { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()               {}
func (*Permission) Descriptor() ([]byte, []int) { return fileDescriptorPermission, []int{0} }

func (m *Permission) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Permission) GetResourcePattern() string {
	if m != nil {
		return m.ResourcePattern
	}
	return ""
}

func init() {
	proto.RegisterType((*Permission)(nil), "cloud_foundry.perm.protos.Permission")
}
func (m *Permission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Permission) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Action) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPermission(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	if len(m.ResourcePattern) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPermission(dAtA, i, uint64(len(m.ResourcePattern)))
		i += copy(dAtA[i:], m.ResourcePattern)
	}
	return i, nil
}

func encodeVarintPermission(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Permission) Size() (n int) {
	var l int
	_ = l
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	l = len(m.ResourcePattern)
	if l > 0 {
		n += 1 + l + sovPermission(uint64(l))
	}
	return n
}

func sovPermission(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPermission(x uint64) (n int) {
	return sovPermission(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Permission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermission
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Permission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Permission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourcePattern", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPermission
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourcePattern = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermission
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
func skipPermission(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPermission
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
					return 0, ErrIntOverflowPermission
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPermission
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPermission
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPermission
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPermission(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPermission = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPermission   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("permission.proto", fileDescriptorPermission) }

var fileDescriptorPermission = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0x2d, 0xca,
	0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4c, 0xce,
	0xc9, 0x2f, 0x4d, 0x89, 0x4f, 0xcb, 0x2f, 0xcd, 0x4b, 0x29, 0xaa, 0xd4, 0x03, 0xc9, 0x43, 0x64,
	0x8a, 0xa5, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3,
	0xd3, 0xf3, 0xf5, 0xc1, 0xe2, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xd4, 0x2b,
	0x85, 0x72, 0x71, 0x05, 0xc0, 0x4d, 0x17, 0x12, 0xe3, 0x62, 0x4b, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf,
	0x93, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x34, 0xb9, 0x04, 0x8a, 0x52, 0x8b,
	0xf3, 0x4b, 0x8b, 0x92, 0x53, 0xe3, 0x0b, 0x12, 0x4b, 0x4a, 0x52, 0x8b, 0xf2, 0x24, 0x98, 0xc1,
	0x2a, 0xf8, 0x61, 0xe2, 0x01, 0x10, 0x61, 0x2f, 0x16, 0x0e, 0x46, 0x01, 0x26, 0x27, 0x89, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0x36, 0x88, 0xfb, 0x92, 0x20, 0xb4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xb5, 0x75,
	0x97, 0xd5, 0x00, 0x00, 0x00,
}
