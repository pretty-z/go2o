// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shipment_service.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShipmentServiceClient is the client API for ShipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShipmentServiceClient interface {
	// 创建一个配送覆盖的区域
	CreateCoverageArea_(ctx context.Context, in *SCoverageValue, opts ...grpc.CallOption) (*Result, error)
	// 获取订单的发货单信息
	GetShipOrderOfOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*ShipmentOrderListResponse, error)
	// * 物流追踪
	GetLogisticFlowTrack(ctx context.Context, in *LogisticFlowTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error)
	// * 获取发货单的物流追踪信息,$shipOrderId:发货单编号 $invert:是否倒序排列
	ShipOrderLogisticTrack(ctx context.Context, in *OrderLogisticTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error)
}

type shipmentServiceClient struct {
	cc *grpc.ClientConn
}

func NewShipmentServiceClient(cc *grpc.ClientConn) ShipmentServiceClient {
	return &shipmentServiceClient{cc}
}

func (c *shipmentServiceClient) CreateCoverageArea_(ctx context.Context, in *SCoverageValue, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShipmentService/CreateCoverageArea_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) GetShipOrderOfOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*ShipmentOrderListResponse, error) {
	out := new(ShipmentOrderListResponse)
	err := c.cc.Invoke(ctx, "/ShipmentService/GetShipOrderOfOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) GetLogisticFlowTrack(ctx context.Context, in *LogisticFlowTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error) {
	out := new(SShipOrderTrack)
	err := c.cc.Invoke(ctx, "/ShipmentService/GetLogisticFlowTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) ShipOrderLogisticTrack(ctx context.Context, in *OrderLogisticTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error) {
	out := new(SShipOrderTrack)
	err := c.cc.Invoke(ctx, "/ShipmentService/ShipOrderLogisticTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShipmentServiceServer is the server API for ShipmentService service.
type ShipmentServiceServer interface {
	// 创建一个配送覆盖的区域
	CreateCoverageArea_(context.Context, *SCoverageValue) (*Result, error)
	// 获取订单的发货单信息
	GetShipOrderOfOrder(context.Context, *OrderId) (*ShipmentOrderListResponse, error)
	// * 物流追踪
	GetLogisticFlowTrack(context.Context, *LogisticFlowTrackRequest) (*SShipOrderTrack, error)
	// * 获取发货单的物流追踪信息,$shipOrderId:发货单编号 $invert:是否倒序排列
	ShipOrderLogisticTrack(context.Context, *OrderLogisticTrackRequest) (*SShipOrderTrack, error)
}

func RegisterShipmentServiceServer(s *grpc.Server, srv ShipmentServiceServer) {
	s.RegisterService(&_ShipmentService_serviceDesc, srv)
}

func _ShipmentService_CreateCoverageArea__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SCoverageValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).CreateCoverageArea_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShipmentService/CreateCoverageArea_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).CreateCoverageArea_(ctx, req.(*SCoverageValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_GetShipOrderOfOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetShipOrderOfOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShipmentService/GetShipOrderOfOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetShipOrderOfOrder(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_GetLogisticFlowTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogisticFlowTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetLogisticFlowTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShipmentService/GetLogisticFlowTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetLogisticFlowTrack(ctx, req.(*LogisticFlowTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_ShipOrderLogisticTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderLogisticTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).ShipOrderLogisticTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShipmentService/ShipOrderLogisticTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).ShipOrderLogisticTrack(ctx, req.(*OrderLogisticTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShipmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShipmentService",
	HandlerType: (*ShipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoverageArea_",
			Handler:    _ShipmentService_CreateCoverageArea__Handler,
		},
		{
			MethodName: "GetShipOrderOfOrder",
			Handler:    _ShipmentService_GetShipOrderOfOrder_Handler,
		},
		{
			MethodName: "GetLogisticFlowTrack",
			Handler:    _ShipmentService_GetLogisticFlowTrack_Handler,
		},
		{
			MethodName: "ShipOrderLogisticTrack",
			Handler:    _ShipmentService_ShipOrderLogisticTrack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipment_service.proto",
}

func init() {
	proto.RegisterFile("shipment_service.proto", fileDescriptor_shipment_service_c3b63cac0f16d282)
}

var fileDescriptor_shipment_service_c3b63cac0f16d282 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xc3, 0xbf, 0xf8, 0x2b, 0x83, 0x50, 0x99, 0x48, 0xc5, 0x59, 0xb8, 0xc8, 0x03, 0x5c,
	0xb1, 0x2e, 0xc5, 0x85, 0x16, 0xad, 0x42, 0xa1, 0x90, 0x88, 0x0b, 0x37, 0x65, 0x92, 0x5c, 0xa7,
	0x83, 0x49, 0x6f, 0x9c, 0xb9, 0xa9, 0x2f, 0xe2, 0x03, 0x4b, 0xa7, 0x49, 0x5c, 0x14, 0x5c, 0x1d,
	0xe6, 0x3b, 0x67, 0xce, 0x81, 0x2b, 0x26, 0x7e, 0x6d, 0x9b, 0x1a, 0x37, 0xbc, 0xf2, 0xe8, 0xb6,
	0xb6, 0x40, 0x68, 0x1c, 0x31, 0xa9, 0x63, 0x53, 0x51, 0xae, 0xab, 0xee, 0xa5, 0x6a, 0xf4, 0x5e,
	0x1b, 0xbc, 0x1c, 0xd2, 0x25, 0x53, 0xe7, 0x9d, 0xf5, 0x1e, 0xb9, 0x12, 0xdd, 0xaf, 0x31, 0xfd,
	0xfe, 0x27, 0xc6, 0x59, 0x97, 0xcf, 0xf6, 0xe5, 0xf2, 0x4a, 0xc4, 0x33, 0x87, 0x9a, 0x71, 0x46,
	0x5b, 0x74, 0xda, 0xe0, 0x9d, 0x43, 0xbd, 0x92, 0x63, 0xc8, 0x7a, 0xf0, 0xaa, 0xab, 0x16, 0xd5,
	0x08, 0x52, 0xf4, 0x6d, 0xc5, 0x49, 0x24, 0x6f, 0x45, 0x3c, 0x47, 0xde, 0x15, 0x2d, 0x77, 0x03,
	0xcb, 0xf7, 0x20, 0xf2, 0x08, 0x82, 0x3e, 0x97, 0x4a, 0x41, 0xbf, 0x12, 0xc8, 0xc2, 0x7a, 0x4e,
	0xd1, 0x37, 0xb4, 0xf1, 0x98, 0x44, 0xf2, 0x41, 0x9c, 0xce, 0x91, 0x17, 0x64, 0xac, 0x67, 0x5b,
	0x3c, 0x56, 0xf4, 0xf5, 0xe2, 0x74, 0xf1, 0x21, 0xcf, 0xe1, 0x80, 0xa5, 0xf8, 0xd9, 0xa2, 0x67,
	0x75, 0x02, 0xd9, 0x30, 0x17, 0x8c, 0x24, 0x92, 0x4f, 0x62, 0x32, 0xb0, 0xfe, 0xe3, 0xbe, 0x48,
	0xc1, 0x21, 0xfc, 0xa3, 0xe9, 0xfe, 0x42, 0xc4, 0x05, 0xd5, 0x60, 0x2c, 0xaf, 0xdb, 0x1c, 0x0c,
	0x4d, 0x09, 0x5c, 0x53, 0xbc, 0x8d, 0xe0, 0x26, 0x9c, 0x2d, 0xff, 0x1f, 0xe4, 0xfa, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x04, 0xd2, 0xae, 0x12, 0x9a, 0x01, 0x00, 0x00,
}
