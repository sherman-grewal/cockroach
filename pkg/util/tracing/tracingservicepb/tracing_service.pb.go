// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/tracing/tracingservicepb/tracing_service.proto

package tracingservicepb

import (
	context "context"
	fmt "fmt"
	tracingpb "github.com/cockroachdb/cockroach/pkg/util/tracing/tracingpb"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetSpanRecordingsRequest struct {
	TraceID uint64 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (m *GetSpanRecordingsRequest) Reset()         { *m = GetSpanRecordingsRequest{} }
func (m *GetSpanRecordingsRequest) String() string { return proto.CompactTextString(m) }
func (*GetSpanRecordingsRequest) ProtoMessage()    {}
func (*GetSpanRecordingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_29b78bec82996ca3, []int{0}
}
func (m *GetSpanRecordingsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSpanRecordingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GetSpanRecordingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpanRecordingsRequest.Merge(m, src)
}
func (m *GetSpanRecordingsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetSpanRecordingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpanRecordingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpanRecordingsRequest proto.InternalMessageInfo

type GetSpanRecordingsResponse struct {
	Recordings []GetSpanRecordingsResponse_Recording `protobuf:"bytes,1,rep,name=recordings,proto3" json:"recordings"`
}

func (m *GetSpanRecordingsResponse) Reset()         { *m = GetSpanRecordingsResponse{} }
func (m *GetSpanRecordingsResponse) String() string { return proto.CompactTextString(m) }
func (*GetSpanRecordingsResponse) ProtoMessage()    {}
func (*GetSpanRecordingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_29b78bec82996ca3, []int{1}
}
func (m *GetSpanRecordingsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSpanRecordingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GetSpanRecordingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpanRecordingsResponse.Merge(m, src)
}
func (m *GetSpanRecordingsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetSpanRecordingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpanRecordingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpanRecordingsResponse proto.InternalMessageInfo

type GetSpanRecordingsResponse_Recording struct {
	RecordedSpans []tracingpb.RecordedSpan `protobuf:"bytes,1,rep,name=recorded_spans,json=recordedSpans,proto3" json:"recorded_spans"`
}

func (m *GetSpanRecordingsResponse_Recording) Reset()         { *m = GetSpanRecordingsResponse_Recording{} }
func (m *GetSpanRecordingsResponse_Recording) String() string { return proto.CompactTextString(m) }
func (*GetSpanRecordingsResponse_Recording) ProtoMessage()    {}
func (*GetSpanRecordingsResponse_Recording) Descriptor() ([]byte, []int) {
	return fileDescriptor_29b78bec82996ca3, []int{1, 0}
}
func (m *GetSpanRecordingsResponse_Recording) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSpanRecordingsResponse_Recording) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GetSpanRecordingsResponse_Recording) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpanRecordingsResponse_Recording.Merge(m, src)
}
func (m *GetSpanRecordingsResponse_Recording) XXX_Size() int {
	return m.Size()
}
func (m *GetSpanRecordingsResponse_Recording) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpanRecordingsResponse_Recording.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpanRecordingsResponse_Recording proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetSpanRecordingsRequest)(nil), "cockroach.util.tracing.GetSpanRecordingsRequest")
	proto.RegisterType((*GetSpanRecordingsResponse)(nil), "cockroach.util.tracing.GetSpanRecordingsResponse")
	proto.RegisterType((*GetSpanRecordingsResponse_Recording)(nil), "cockroach.util.tracing.GetSpanRecordingsResponse.Recording")
}

func init() {
	proto.RegisterFile("util/tracing/tracingservicepb/tracing_service.proto", fileDescriptor_29b78bec82996ca3)
}

var fileDescriptor_29b78bec82996ca3 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2e, 0x2d, 0xc9, 0xcc,
	0xd1, 0x2f, 0x29, 0x4a, 0x4c, 0xce, 0xcc, 0x4b, 0x87, 0xd1, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9,
	0xa9, 0x05, 0x49, 0x30, 0x81, 0x78, 0xa8, 0x88, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x58,
	0x72, 0x7e, 0x72, 0x76, 0x51, 0x7e, 0x62, 0x72, 0x86, 0x1e, 0x48, 0xbb, 0x1e, 0x54, 0x95, 0x94,
	0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x89, 0x3e, 0x88, 0x05, 0x51, 0x2d, 0xa5, 0x85, 0xcd, 0x8a,
	0x82, 0x24, 0xfd, 0xa2, 0xd4, 0xe4, 0xfc, 0xa2, 0x94, 0xd4, 0x94, 0xf8, 0xe2, 0x82, 0xc4, 0x3c,
	0x88, 0x5a, 0x25, 0x27, 0x2e, 0x09, 0xf7, 0xd4, 0x92, 0xe0, 0x82, 0xc4, 0xbc, 0x20, 0xb0, 0x2c,
	0xc8, 0x31, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x6a, 0x5c, 0x1c, 0x20, 0xcd, 0xa9,
	0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x4e, 0xdc, 0x8f, 0xee, 0xc9, 0xb3, 0x87,
	0x80, 0xc4, 0x3c, 0x5d, 0x82, 0xd8, 0xc1, 0x92, 0x9e, 0x29, 0x4a, 0x2f, 0x18, 0xb9, 0x24, 0xb1,
	0x18, 0x52, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x94, 0xc8, 0xc5, 0x55, 0x04, 0x17, 0x95, 0x60,
	0x54, 0x60, 0xd6, 0xe0, 0x36, 0xb2, 0xd6, 0xc3, 0xee, 0x21, 0x3d, 0x9c, 0xc6, 0xe8, 0xc1, 0x85,
	0x9c, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x42, 0x32, 0x54, 0x2a, 0x83, 0x8b, 0x13, 0x2e, 0x2d,
	0x14, 0xcd, 0xc5, 0x87, 0xe2, 0x51, 0x98, 0x9d, 0x7a, 0xb8, 0xec, 0x84, 0x07, 0x10, 0xd4, 0x8e,
	0xd4, 0x14, 0x90, 0x13, 0xa0, 0xd6, 0xf0, 0x16, 0x21, 0x89, 0x15, 0x1b, 0xb5, 0x32, 0x72, 0x81,
	0xfd, 0x0f, 0xb2, 0xa8, 0x8a, 0x4b, 0x10, 0xc3, 0xb9, 0x42, 0x06, 0x24, 0xf8, 0x0c, 0x1c, 0xca,
	0x52, 0x86, 0x24, 0x87, 0x85, 0x12, 0x83, 0x93, 0xd1, 0x89, 0x87, 0x72, 0x0c, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0x78, 0xe3, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x00, 0x7a, 0xd2, 0x4a,
	0x62, 0x03, 0xc7, 0xb8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x43, 0x42, 0x98, 0x9f, 0x82, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TracingClient is the client API for Tracing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TracingClient interface {
	// GetSpanRecordings contacts a nodes' inflight span registry to return the
	// tracing.Recordings for each root span with a matching TraceID.
	GetSpanRecordings(ctx context.Context, in *GetSpanRecordingsRequest, opts ...grpc.CallOption) (*GetSpanRecordingsResponse, error)
}

type tracingClient struct {
	cc *grpc.ClientConn
}

func NewTracingClient(cc *grpc.ClientConn) TracingClient {
	return &tracingClient{cc}
}

func (c *tracingClient) GetSpanRecordings(ctx context.Context, in *GetSpanRecordingsRequest, opts ...grpc.CallOption) (*GetSpanRecordingsResponse, error) {
	out := new(GetSpanRecordingsResponse)
	err := c.cc.Invoke(ctx, "/cockroach.util.tracing.Tracing/GetSpanRecordings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TracingServer is the server API for Tracing service.
type TracingServer interface {
	// GetSpanRecordings contacts a nodes' inflight span registry to return the
	// tracing.Recordings for each root span with a matching TraceID.
	GetSpanRecordings(context.Context, *GetSpanRecordingsRequest) (*GetSpanRecordingsResponse, error)
}

// UnimplementedTracingServer can be embedded to have forward compatible implementations.
type UnimplementedTracingServer struct {
}

func (*UnimplementedTracingServer) GetSpanRecordings(ctx context.Context, req *GetSpanRecordingsRequest) (*GetSpanRecordingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpanRecordings not implemented")
}

func RegisterTracingServer(s *grpc.Server, srv TracingServer) {
	s.RegisterService(&_Tracing_serviceDesc, srv)
}

func _Tracing_GetSpanRecordings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpanRecordingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracingServer).GetSpanRecordings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.util.tracing.Tracing/GetSpanRecordings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracingServer).GetSpanRecordings(ctx, req.(*GetSpanRecordingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.util.tracing.Tracing",
	HandlerType: (*TracingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSpanRecordings",
			Handler:    _Tracing_GetSpanRecordings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "util/tracing/tracingservicepb/tracing_service.proto",
}

func (m *GetSpanRecordingsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSpanRecordingsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSpanRecordingsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TraceID != 0 {
		i = encodeVarintTracingService(dAtA, i, uint64(m.TraceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetSpanRecordingsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSpanRecordingsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSpanRecordingsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Recordings) > 0 {
		for iNdEx := len(m.Recordings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Recordings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTracingService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetSpanRecordingsResponse_Recording) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSpanRecordingsResponse_Recording) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSpanRecordingsResponse_Recording) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RecordedSpans) > 0 {
		for iNdEx := len(m.RecordedSpans) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecordedSpans[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTracingService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTracingService(dAtA []byte, offset int, v uint64) int {
	offset -= sovTracingService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetSpanRecordingsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TraceID != 0 {
		n += 1 + sovTracingService(uint64(m.TraceID))
	}
	return n
}

func (m *GetSpanRecordingsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Recordings) > 0 {
		for _, e := range m.Recordings {
			l = e.Size()
			n += 1 + l + sovTracingService(uint64(l))
		}
	}
	return n
}

func (m *GetSpanRecordingsResponse_Recording) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RecordedSpans) > 0 {
		for _, e := range m.RecordedSpans {
			l = e.Size()
			n += 1 + l + sovTracingService(uint64(l))
		}
	}
	return n
}

func sovTracingService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTracingService(x uint64) (n int) {
	return sovTracingService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetSpanRecordingsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracingService
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
			return fmt.Errorf("proto: GetSpanRecordingsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSpanRecordingsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceID", wireType)
			}
			m.TraceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracingService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TraceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTracingService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTracingService
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
func (m *GetSpanRecordingsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracingService
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
			return fmt.Errorf("proto: GetSpanRecordingsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSpanRecordingsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recordings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracingService
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
				return ErrInvalidLengthTracingService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracingService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recordings = append(m.Recordings, GetSpanRecordingsResponse_Recording{})
			if err := m.Recordings[len(m.Recordings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracingService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTracingService
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
func (m *GetSpanRecordingsResponse_Recording) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracingService
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
			return fmt.Errorf("proto: Recording: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Recording: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordedSpans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracingService
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
				return ErrInvalidLengthTracingService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracingService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecordedSpans = append(m.RecordedSpans, tracingpb.RecordedSpan{})
			if err := m.RecordedSpans[len(m.RecordedSpans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracingService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTracingService
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
func skipTracingService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTracingService
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
					return 0, ErrIntOverflowTracingService
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
					return 0, ErrIntOverflowTracingService
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
				return 0, ErrInvalidLengthTracingService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTracingService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTracingService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTracingService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTracingService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTracingService = fmt.Errorf("proto: unexpected end of group")
)