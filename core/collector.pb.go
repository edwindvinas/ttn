// Code generated by protoc-gen-gogo.
// source: collector.proto
// DO NOT EDIT!

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetApplicationsCollectorReq struct {
}

func (m *GetApplicationsCollectorReq) Reset()         { *m = GetApplicationsCollectorReq{} }
func (m *GetApplicationsCollectorReq) String() string { return proto.CompactTextString(m) }
func (*GetApplicationsCollectorReq) ProtoMessage()    {}
func (*GetApplicationsCollectorReq) Descriptor() ([]byte, []int) {
	return fileDescriptorCollector, []int{0}
}

type GetApplicationsCollectorRes struct {
	Applications []*CollectorApplication `protobuf:"bytes,1,rep,name=Applications" json:"Applications,omitempty"`
}

func (m *GetApplicationsCollectorRes) Reset()         { *m = GetApplicationsCollectorRes{} }
func (m *GetApplicationsCollectorRes) String() string { return proto.CompactTextString(m) }
func (*GetApplicationsCollectorRes) ProtoMessage()    {}
func (*GetApplicationsCollectorRes) Descriptor() ([]byte, []int) {
	return fileDescriptorCollector, []int{1}
}

func (m *GetApplicationsCollectorRes) GetApplications() []*CollectorApplication {
	if m != nil {
		return m.Applications
	}
	return nil
}

type CollectorApplication struct {
	AppEUI []byte `protobuf:"bytes,1,opt,name=AppEUI,proto3" json:"AppEUI,omitempty"`
}

func (m *CollectorApplication) Reset()                    { *m = CollectorApplication{} }
func (m *CollectorApplication) String() string            { return proto.CompactTextString(m) }
func (*CollectorApplication) ProtoMessage()               {}
func (*CollectorApplication) Descriptor() ([]byte, []int) { return fileDescriptorCollector, []int{2} }

type AddApplicationCollectorReq struct {
	AppEUI       []byte `protobuf:"bytes,1,opt,name=AppEUI,proto3" json:"AppEUI,omitempty"`
	AppAccessKey string `protobuf:"bytes,2,opt,name=AppAccessKey,proto3" json:"AppAccessKey,omitempty"`
}

func (m *AddApplicationCollectorReq) Reset()         { *m = AddApplicationCollectorReq{} }
func (m *AddApplicationCollectorReq) String() string { return proto.CompactTextString(m) }
func (*AddApplicationCollectorReq) ProtoMessage()    {}
func (*AddApplicationCollectorReq) Descriptor() ([]byte, []int) {
	return fileDescriptorCollector, []int{3}
}

type AddApplicationCollectorRes struct {
}

func (m *AddApplicationCollectorRes) Reset()         { *m = AddApplicationCollectorRes{} }
func (m *AddApplicationCollectorRes) String() string { return proto.CompactTextString(m) }
func (*AddApplicationCollectorRes) ProtoMessage()    {}
func (*AddApplicationCollectorRes) Descriptor() ([]byte, []int) {
	return fileDescriptorCollector, []int{4}
}

type RemoveApplicationCollectorReq struct {
	AppEUI []byte `protobuf:"bytes,1,opt,name=AppEUI,proto3" json:"AppEUI,omitempty"`
}

func (m *RemoveApplicationCollectorReq) Reset()         { *m = RemoveApplicationCollectorReq{} }
func (m *RemoveApplicationCollectorReq) String() string { return proto.CompactTextString(m) }
func (*RemoveApplicationCollectorReq) ProtoMessage()    {}
func (*RemoveApplicationCollectorReq) Descriptor() ([]byte, []int) {
	return fileDescriptorCollector, []int{5}
}

type RemoveApplicationCollectorRes struct {
}

func (m *RemoveApplicationCollectorRes) Reset()         { *m = RemoveApplicationCollectorRes{} }
func (m *RemoveApplicationCollectorRes) String() string { return proto.CompactTextString(m) }
func (*RemoveApplicationCollectorRes) ProtoMessage()    {}
func (*RemoveApplicationCollectorRes) Descriptor() ([]byte, []int) {
	return fileDescriptorCollector, []int{6}
}

func init() {
	proto.RegisterType((*GetApplicationsCollectorReq)(nil), "core.GetApplicationsCollectorReq")
	proto.RegisterType((*GetApplicationsCollectorRes)(nil), "core.GetApplicationsCollectorRes")
	proto.RegisterType((*CollectorApplication)(nil), "core.CollectorApplication")
	proto.RegisterType((*AddApplicationCollectorReq)(nil), "core.AddApplicationCollectorReq")
	proto.RegisterType((*AddApplicationCollectorRes)(nil), "core.AddApplicationCollectorRes")
	proto.RegisterType((*RemoveApplicationCollectorReq)(nil), "core.RemoveApplicationCollectorReq")
	proto.RegisterType((*RemoveApplicationCollectorRes)(nil), "core.RemoveApplicationCollectorRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for CollectorManager service

type CollectorManagerClient interface {
	GetApplications(ctx context.Context, in *GetApplicationsCollectorReq, opts ...grpc.CallOption) (*GetApplicationsCollectorRes, error)
	AddApplication(ctx context.Context, in *AddApplicationCollectorReq, opts ...grpc.CallOption) (*AddApplicationCollectorRes, error)
	RemoveApplication(ctx context.Context, in *RemoveApplicationCollectorReq, opts ...grpc.CallOption) (*RemoveApplicationCollectorRes, error)
}

type collectorManagerClient struct {
	cc *grpc.ClientConn
}

func NewCollectorManagerClient(cc *grpc.ClientConn) CollectorManagerClient {
	return &collectorManagerClient{cc}
}

func (c *collectorManagerClient) GetApplications(ctx context.Context, in *GetApplicationsCollectorReq, opts ...grpc.CallOption) (*GetApplicationsCollectorRes, error) {
	out := new(GetApplicationsCollectorRes)
	err := grpc.Invoke(ctx, "/core.CollectorManager/GetApplications", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorManagerClient) AddApplication(ctx context.Context, in *AddApplicationCollectorReq, opts ...grpc.CallOption) (*AddApplicationCollectorRes, error) {
	out := new(AddApplicationCollectorRes)
	err := grpc.Invoke(ctx, "/core.CollectorManager/AddApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorManagerClient) RemoveApplication(ctx context.Context, in *RemoveApplicationCollectorReq, opts ...grpc.CallOption) (*RemoveApplicationCollectorRes, error) {
	out := new(RemoveApplicationCollectorRes)
	err := grpc.Invoke(ctx, "/core.CollectorManager/RemoveApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CollectorManager service

type CollectorManagerServer interface {
	GetApplications(context.Context, *GetApplicationsCollectorReq) (*GetApplicationsCollectorRes, error)
	AddApplication(context.Context, *AddApplicationCollectorReq) (*AddApplicationCollectorRes, error)
	RemoveApplication(context.Context, *RemoveApplicationCollectorReq) (*RemoveApplicationCollectorRes, error)
}

func RegisterCollectorManagerServer(s *grpc.Server, srv CollectorManagerServer) {
	s.RegisterService(&_CollectorManager_serviceDesc, srv)
}

func _CollectorManager_GetApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationsCollectorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorManagerServer).GetApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CollectorManager/GetApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorManagerServer).GetApplications(ctx, req.(*GetApplicationsCollectorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectorManager_AddApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddApplicationCollectorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorManagerServer).AddApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CollectorManager/AddApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorManagerServer).AddApplication(ctx, req.(*AddApplicationCollectorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectorManager_RemoveApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveApplicationCollectorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorManagerServer).RemoveApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CollectorManager/RemoveApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorManagerServer).RemoveApplication(ctx, req.(*RemoveApplicationCollectorReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollectorManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.CollectorManager",
	HandlerType: (*CollectorManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApplications",
			Handler:    _CollectorManager_GetApplications_Handler,
		},
		{
			MethodName: "AddApplication",
			Handler:    _CollectorManager_AddApplication_Handler,
		},
		{
			MethodName: "RemoveApplication",
			Handler:    _CollectorManager_RemoveApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorCollector,
}

func (m *GetApplicationsCollectorReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetApplicationsCollectorReq) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetApplicationsCollectorRes) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetApplicationsCollectorRes) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Applications) > 0 {
		for _, msg := range m.Applications {
			data[i] = 0xa
			i++
			i = encodeVarintCollector(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CollectorApplication) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CollectorApplication) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AppEUI) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCollector(data, i, uint64(len(m.AppEUI)))
		i += copy(data[i:], m.AppEUI)
	}
	return i, nil
}

func (m *AddApplicationCollectorReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *AddApplicationCollectorReq) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AppEUI) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCollector(data, i, uint64(len(m.AppEUI)))
		i += copy(data[i:], m.AppEUI)
	}
	if len(m.AppAccessKey) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintCollector(data, i, uint64(len(m.AppAccessKey)))
		i += copy(data[i:], m.AppAccessKey)
	}
	return i, nil
}

func (m *AddApplicationCollectorRes) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *AddApplicationCollectorRes) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *RemoveApplicationCollectorReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RemoveApplicationCollectorReq) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AppEUI) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCollector(data, i, uint64(len(m.AppEUI)))
		i += copy(data[i:], m.AppEUI)
	}
	return i, nil
}

func (m *RemoveApplicationCollectorRes) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RemoveApplicationCollectorRes) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64Collector(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Collector(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCollector(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *GetApplicationsCollectorReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetApplicationsCollectorRes) Size() (n int) {
	var l int
	_ = l
	if len(m.Applications) > 0 {
		for _, e := range m.Applications {
			l = e.Size()
			n += 1 + l + sovCollector(uint64(l))
		}
	}
	return n
}

func (m *CollectorApplication) Size() (n int) {
	var l int
	_ = l
	l = len(m.AppEUI)
	if l > 0 {
		n += 1 + l + sovCollector(uint64(l))
	}
	return n
}

func (m *AddApplicationCollectorReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.AppEUI)
	if l > 0 {
		n += 1 + l + sovCollector(uint64(l))
	}
	l = len(m.AppAccessKey)
	if l > 0 {
		n += 1 + l + sovCollector(uint64(l))
	}
	return n
}

func (m *AddApplicationCollectorRes) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *RemoveApplicationCollectorReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.AppEUI)
	if l > 0 {
		n += 1 + l + sovCollector(uint64(l))
	}
	return n
}

func (m *RemoveApplicationCollectorRes) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovCollector(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCollector(x uint64) (n int) {
	return sovCollector(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetApplicationsCollectorReq) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollector
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetApplicationsCollectorReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetApplicationsCollectorReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCollector(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollector
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
func (m *GetApplicationsCollectorRes) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollector
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetApplicationsCollectorRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetApplicationsCollectorRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Applications", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollector
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCollector
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Applications = append(m.Applications, &CollectorApplication{})
			if err := m.Applications[len(m.Applications)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollector(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollector
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
func (m *CollectorApplication) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollector
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CollectorApplication: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectorApplication: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollector
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCollector
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppEUI = append(m.AppEUI[:0], data[iNdEx:postIndex]...)
			if m.AppEUI == nil {
				m.AppEUI = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollector(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollector
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
func (m *AddApplicationCollectorReq) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollector
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddApplicationCollectorReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddApplicationCollectorReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollector
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCollector
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppEUI = append(m.AppEUI[:0], data[iNdEx:postIndex]...)
			if m.AppEUI == nil {
				m.AppEUI = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppAccessKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollector
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCollector
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppAccessKey = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollector(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollector
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
func (m *AddApplicationCollectorRes) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollector
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddApplicationCollectorRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddApplicationCollectorRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCollector(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollector
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
func (m *RemoveApplicationCollectorReq) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollector
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RemoveApplicationCollectorReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveApplicationCollectorReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollector
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCollector
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppEUI = append(m.AppEUI[:0], data[iNdEx:postIndex]...)
			if m.AppEUI == nil {
				m.AppEUI = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollector(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollector
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
func (m *RemoveApplicationCollectorRes) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollector
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RemoveApplicationCollectorRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveApplicationCollectorRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCollector(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollector
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
func skipCollector(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollector
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowCollector
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowCollector
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCollector
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCollector
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipCollector(data[start:])
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
	ErrInvalidLengthCollector = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollector   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("collector.proto", fileDescriptorCollector) }

var fileDescriptorCollector = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0xcf, 0xc9,
	0x49, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0xce, 0x2f,
	0x4a, 0x55, 0x92, 0xe5, 0x92, 0x76, 0x4f, 0x2d, 0x71, 0x2c, 0x28, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c,
	0xc9, 0xcc, 0xcf, 0x2b, 0x76, 0x86, 0xa9, 0x0b, 0x4a, 0x2d, 0x54, 0x8a, 0xc5, 0x27, 0x5d, 0x2c,
	0x64, 0xc7, 0xc5, 0x83, 0x2c, 0x27, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xa5, 0x07, 0x32,
	0x5a, 0x0f, 0xae, 0x12, 0x49, 0x49, 0x10, 0x8a, 0x7a, 0x25, 0x3d, 0x2e, 0x11, 0x6c, 0xaa, 0x84,
	0xc4, 0xb8, 0xd8, 0x1c, 0x0b, 0x0a, 0x5c, 0x43, 0x3d, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82,
	0xa0, 0x3c, 0xa5, 0x08, 0x2e, 0x29, 0xc7, 0x94, 0x14, 0x24, 0x95, 0xc8, 0x8e, 0xc5, 0xa5, 0x4b,
	0x48, 0x09, 0xec, 0x4a, 0xc7, 0xe4, 0xe4, 0xd4, 0xe2, 0x62, 0xef, 0xd4, 0x4a, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x14, 0x31, 0x25, 0x19, 0x3c, 0x26, 0x17, 0x2b, 0x99, 0x73, 0xc9, 0x06,
	0xa5, 0xe6, 0xe6, 0x97, 0xa5, 0x92, 0x68, 0xb5, 0x92, 0x3c, 0x7e, 0x8d, 0xc5, 0x46, 0x0b, 0x98,
	0xb8, 0x04, 0xe0, 0x02, 0xbe, 0x89, 0x79, 0x89, 0xe9, 0xa9, 0x45, 0x42, 0xe1, 0x5c, 0xfc, 0x68,
	0xa1, 0x2e, 0xa4, 0x08, 0x09, 0x53, 0x3c, 0x71, 0x25, 0x45, 0x50, 0x49, 0xb1, 0x50, 0x08, 0x17,
	0x1f, 0xaa, 0x2f, 0x85, 0x14, 0x20, 0x9a, 0x70, 0x87, 0xaa, 0x14, 0x21, 0x15, 0xc5, 0x42, 0xb1,
	0x5c, 0x82, 0x18, 0x9e, 0x14, 0x52, 0x86, 0x68, 0xc3, 0x1b, 0x6c, 0x52, 0x44, 0x28, 0x2a, 0x76,
	0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c,
	0x96, 0x63, 0x48, 0x62, 0x03, 0xa7, 0x60, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0x44,
	0x60, 0xbb, 0xd4, 0x02, 0x00, 0x00,
}
