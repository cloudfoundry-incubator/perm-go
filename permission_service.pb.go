// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: permission_service.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HasPermissionRequest struct {
	Actor    *Actor `protobuf:"bytes,1,opt,name=actor" json:"actor,omitempty"`
	Action   string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Resource string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (m *HasPermissionRequest) Reset()         { *m = HasPermissionRequest{} }
func (m *HasPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*HasPermissionRequest) ProtoMessage()    {}
func (*HasPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPermissionService, []int{0}
}

func (m *HasPermissionRequest) GetActor() *Actor {
	if m != nil {
		return m.Actor
	}
	return nil
}

func (m *HasPermissionRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *HasPermissionRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type HasPermissionResponse struct {
	HasPermission bool `protobuf:"varint,1,opt,name=has_permission,json=hasPermission,proto3" json:"has_permission,omitempty"`
}

func (m *HasPermissionResponse) Reset()         { *m = HasPermissionResponse{} }
func (m *HasPermissionResponse) String() string { return proto.CompactTextString(m) }
func (*HasPermissionResponse) ProtoMessage()    {}
func (*HasPermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPermissionService, []int{1}
}

func (m *HasPermissionResponse) GetHasPermission() bool {
	if m != nil {
		return m.HasPermission
	}
	return false
}

type ListResourcePatternsRequest struct {
	Actor  *Actor `protobuf:"bytes,2,opt,name=actor" json:"actor,omitempty"`
	Action string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (m *ListResourcePatternsRequest) Reset()         { *m = ListResourcePatternsRequest{} }
func (m *ListResourcePatternsRequest) String() string { return proto.CompactTextString(m) }
func (*ListResourcePatternsRequest) ProtoMessage()    {}
func (*ListResourcePatternsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPermissionService, []int{2}
}

func (m *ListResourcePatternsRequest) GetActor() *Actor {
	if m != nil {
		return m.Actor
	}
	return nil
}

func (m *ListResourcePatternsRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type ListResourcePatternsResponse struct {
	ResourcePatterns []string `protobuf:"bytes,2,rep,name=resource_patterns,json=resourcePatterns" json:"resource_patterns,omitempty"`
}

func (m *ListResourcePatternsResponse) Reset()         { *m = ListResourcePatternsResponse{} }
func (m *ListResourcePatternsResponse) String() string { return proto.CompactTextString(m) }
func (*ListResourcePatternsResponse) ProtoMessage()    {}
func (*ListResourcePatternsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPermissionService, []int{3}
}

func (m *ListResourcePatternsResponse) GetResourcePatterns() []string {
	if m != nil {
		return m.ResourcePatterns
	}
	return nil
}

func init() {
	proto.RegisterType((*HasPermissionRequest)(nil), "cloud_foundry.perm.protos.HasPermissionRequest")
	proto.RegisterType((*HasPermissionResponse)(nil), "cloud_foundry.perm.protos.HasPermissionResponse")
	proto.RegisterType((*ListResourcePatternsRequest)(nil), "cloud_foundry.perm.protos.ListResourcePatternsRequest")
	proto.RegisterType((*ListResourcePatternsResponse)(nil), "cloud_foundry.perm.protos.ListResourcePatternsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PermissionService service

type PermissionServiceClient interface {
	HasPermission(ctx context.Context, in *HasPermissionRequest, opts ...grpc.CallOption) (*HasPermissionResponse, error)
	ListResourcePatterns(ctx context.Context, in *ListResourcePatternsRequest, opts ...grpc.CallOption) (*ListResourcePatternsResponse, error)
}

type permissionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPermissionServiceClient(cc *grpc.ClientConn) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) HasPermission(ctx context.Context, in *HasPermissionRequest, opts ...grpc.CallOption) (*HasPermissionResponse, error) {
	out := new(HasPermissionResponse)
	err := grpc.Invoke(ctx, "/cloud_foundry.perm.protos.PermissionService/HasPermission", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) ListResourcePatterns(ctx context.Context, in *ListResourcePatternsRequest, opts ...grpc.CallOption) (*ListResourcePatternsResponse, error) {
	out := new(ListResourcePatternsResponse)
	err := grpc.Invoke(ctx, "/cloud_foundry.perm.protos.PermissionService/ListResourcePatterns", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PermissionService service

type PermissionServiceServer interface {
	HasPermission(context.Context, *HasPermissionRequest) (*HasPermissionResponse, error)
	ListResourcePatterns(context.Context, *ListResourcePatternsRequest) (*ListResourcePatternsResponse, error)
}

func RegisterPermissionServiceServer(s *grpc.Server, srv PermissionServiceServer) {
	s.RegisterService(&_PermissionService_serviceDesc, srv)
}

func _PermissionService_HasPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).HasPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud_foundry.perm.protos.PermissionService/HasPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).HasPermission(ctx, req.(*HasPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_ListResourcePatterns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcePatternsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).ListResourcePatterns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud_foundry.perm.protos.PermissionService/ListResourcePatterns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).ListResourcePatterns(ctx, req.(*ListResourcePatternsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PermissionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud_foundry.perm.protos.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HasPermission",
			Handler:    _PermissionService_HasPermission_Handler,
		},
		{
			MethodName: "ListResourcePatterns",
			Handler:    _PermissionService_ListResourcePatterns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permission_service.proto",
}

func (m *HasPermissionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HasPermissionRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Actor != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPermissionService(dAtA, i, uint64(m.Actor.Size()))
		n1, err := m.Actor.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Action) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPermissionService(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	if len(m.Resource) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPermissionService(dAtA, i, uint64(len(m.Resource)))
		i += copy(dAtA[i:], m.Resource)
	}
	return i, nil
}

func (m *HasPermissionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HasPermissionResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HasPermission {
		dAtA[i] = 0x8
		i++
		if m.HasPermission {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *ListResourcePatternsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListResourcePatternsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Actor != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPermissionService(dAtA, i, uint64(m.Actor.Size()))
		n2, err := m.Actor.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Action) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPermissionService(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	return i, nil
}

func (m *ListResourcePatternsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListResourcePatternsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ResourcePatterns) > 0 {
		for _, s := range m.ResourcePatterns {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintPermissionService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HasPermissionRequest) Size() (n int) {
	var l int
	_ = l
	if m.Actor != nil {
		l = m.Actor.Size()
		n += 1 + l + sovPermissionService(uint64(l))
	}
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovPermissionService(uint64(l))
	}
	l = len(m.Resource)
	if l > 0 {
		n += 1 + l + sovPermissionService(uint64(l))
	}
	return n
}

func (m *HasPermissionResponse) Size() (n int) {
	var l int
	_ = l
	if m.HasPermission {
		n += 2
	}
	return n
}

func (m *ListResourcePatternsRequest) Size() (n int) {
	var l int
	_ = l
	if m.Actor != nil {
		l = m.Actor.Size()
		n += 1 + l + sovPermissionService(uint64(l))
	}
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovPermissionService(uint64(l))
	}
	return n
}

func (m *ListResourcePatternsResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.ResourcePatterns) > 0 {
		for _, s := range m.ResourcePatterns {
			l = len(s)
			n += 1 + l + sovPermissionService(uint64(l))
		}
	}
	return n
}

func sovPermissionService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPermissionService(x uint64) (n int) {
	return sovPermissionService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HasPermissionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissionService
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
			return fmt.Errorf("proto: HasPermissionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HasPermissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissionService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPermissionService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Actor == nil {
				m.Actor = &Actor{}
			}
			if err := m.Actor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissionService
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
				return ErrInvalidLengthPermissionService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissionService
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
				return ErrInvalidLengthPermissionService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermissionService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermissionService
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
func (m *HasPermissionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissionService
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
			return fmt.Errorf("proto: HasPermissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HasPermissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasPermission", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissionService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasPermission = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPermissionService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermissionService
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
func (m *ListResourcePatternsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissionService
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
			return fmt.Errorf("proto: ListResourcePatternsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListResourcePatternsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissionService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPermissionService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Actor == nil {
				m.Actor = &Actor{}
			}
			if err := m.Actor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissionService
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
				return ErrInvalidLengthPermissionService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermissionService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermissionService
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
func (m *ListResourcePatternsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissionService
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
			return fmt.Errorf("proto: ListResourcePatternsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListResourcePatternsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourcePatterns", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissionService
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
				return ErrInvalidLengthPermissionService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourcePatterns = append(m.ResourcePatterns, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermissionService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPermissionService
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
func skipPermissionService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPermissionService
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
					return 0, ErrIntOverflowPermissionService
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
					return 0, ErrIntOverflowPermissionService
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
				return 0, ErrInvalidLengthPermissionService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPermissionService
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
				next, err := skipPermissionService(dAtA[start:])
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
	ErrInvalidLengthPermissionService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPermissionService   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("permission_service.proto", fileDescriptorPermissionService) }

var fileDescriptorPermissionService = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x75, 0x53, 0x2d, 0xed, 0x94, 0x4a, 0xbb, 0x54, 0x89, 0x51, 0x42, 0x08, 0x08, 0x05, 0x31,
	0x95, 0x0a, 0xf5, 0x26, 0xe8, 0x49, 0xc4, 0x43, 0x8d, 0x37, 0x2f, 0x21, 0x4d, 0xb7, 0x6d, 0xc0,
	0x66, 0xe3, 0xce, 0x46, 0xf0, 0xea, 0xc5, 0xb3, 0x37, 0x3f, 0xc9, 0xa3, 0x9f, 0x20, 0xf5, 0x47,
	0xa4, 0xd9, 0xd8, 0xd6, 0x5a, 0x8b, 0x7a, 0x4a, 0xde, 0xce, 0x7b, 0x33, 0xef, 0xcd, 0x80, 0x1e,
	0x33, 0x31, 0x0c, 0x11, 0x43, 0x1e, 0x79, 0xc8, 0xc4, 0x5d, 0x18, 0x30, 0x27, 0x16, 0x5c, 0x72,
	0xba, 0x15, 0xdc, 0xf0, 0xa4, 0xeb, 0xf5, 0x78, 0x12, 0x75, 0xc5, 0xbd, 0x33, 0xe6, 0xa9, 0x0a,
	0x1a, 0xfb, 0xfd, 0x50, 0x0e, 0x92, 0x8e, 0x13, 0xf0, 0x61, 0xa3, 0xcf, 0xfb, 0xbc, 0x91, 0xbe,
	0x77, 0x92, 0x5e, 0x8a, 0x52, 0x90, 0xfe, 0x29, 0xbe, 0x51, 0xf2, 0x03, 0xc9, 0x85, 0x02, 0xf6,
	0x03, 0x81, 0xda, 0x99, 0x8f, 0xed, 0xc9, 0x58, 0x97, 0xdd, 0x26, 0x0c, 0x25, 0x6d, 0xc1, 0x5a,
	0xca, 0xd3, 0x89, 0x45, 0xea, 0xa5, 0xa6, 0xe5, 0xfc, 0x38, 0xdf, 0x39, 0x19, 0xf3, 0x5c, 0x45,
	0xa7, 0x9b, 0x90, 0xf7, 0x03, 0x19, 0xf2, 0x48, 0xd7, 0x2c, 0x52, 0x2f, 0xba, 0x19, 0xa2, 0x06,
	0x14, 0x04, 0x43, 0x9e, 0x88, 0x80, 0xe9, 0xb9, 0xb4, 0x32, 0xc1, 0xf6, 0x31, 0x6c, 0xcc, 0x79,
	0xc0, 0x98, 0x47, 0xc8, 0xe8, 0x2e, 0xac, 0x0f, 0x7c, 0xf4, 0xa6, 0x4b, 0x49, 0xdd, 0x14, 0xdc,
	0xf2, 0x60, 0x96, 0x6e, 0x23, 0x6c, 0x5f, 0x84, 0x28, 0xdd, 0xac, 0x5f, 0xdb, 0x97, 0x92, 0x89,
	0x08, 0xbf, 0x45, 0xd1, 0xfe, 0x1b, 0x25, 0x37, 0x1b, 0xe5, 0x7c, 0xb5, 0x40, 0x2a, 0x9a, 0x7d,
	0x09, 0x3b, 0x8b, 0x87, 0x66, 0xde, 0xf7, 0xa0, 0xfa, 0x19, 0xd0, 0x8b, 0xb3, 0xa2, 0xae, 0x59,
	0xb9, 0x7a, 0xd1, 0xad, 0x88, 0x39, 0x91, 0x6a, 0xd9, 0x7c, 0xd2, 0xa0, 0x3a, 0x8d, 0x75, 0xa5,
	0xee, 0x4f, 0x05, 0x94, 0xbf, 0x6c, 0x87, 0x36, 0x96, 0x04, 0x58, 0x74, 0x4b, 0xe3, 0xe0, 0xf7,
	0x82, 0xcc, 0xfc, 0x23, 0x81, 0xda, 0xa2, 0x74, 0xb4, 0xb5, 0xa4, 0xd5, 0x92, 0x1b, 0x18, 0x47,
	0x7f, 0xd6, 0x29, 0x27, 0xa7, 0xfa, 0xcb, 0xc8, 0x24, 0xaf, 0x23, 0x93, 0xbc, 0x8d, 0x4c, 0xf2,
	0xfc, 0x6e, 0xae, 0x5c, 0xe7, 0x95, 0xac, 0xa3, 0xbe, 0x87, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xec, 0x67, 0x8c, 0x74, 0x34, 0x03, 0x00, 0x00,
}
