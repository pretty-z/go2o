// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shop_service.proto

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

type TurnShopRequest struct {
	ShopId               int64    `protobuf:"zigzag64,1,opt,name=shopId,proto3" json:"shopId"`
	On                   bool     `protobuf:"varint,2,opt,name=on,proto3" json:"on"`
	Reason               string   `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TurnShopRequest) Reset()         { *m = TurnShopRequest{} }
func (m *TurnShopRequest) String() string { return proto.CompactTextString(m) }
func (*TurnShopRequest) ProtoMessage()    {}
func (*TurnShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_e0b4ce123fd750f5, []int{0}
}
func (m *TurnShopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TurnShopRequest.Unmarshal(m, b)
}
func (m *TurnShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TurnShopRequest.Marshal(b, m, deterministic)
}
func (dst *TurnShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TurnShopRequest.Merge(dst, src)
}
func (m *TurnShopRequest) XXX_Size() int {
	return xxx_messageInfo_TurnShopRequest.Size(m)
}
func (m *TurnShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TurnShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TurnShopRequest proto.InternalMessageInfo

func (m *TurnShopRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *TurnShopRequest) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *TurnShopRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*TurnShopRequest)(nil), "TurnShopRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShopServiceClient interface {
	// * 获取店铺,shopId
	GetShop(ctx context.Context, in *ShopId, opts ...grpc.CallOption) (*SShop, error)
	// rpc GetVendorShop_ (Int64) returns (SShop) {}
	// 检查商户是否开通店铺
	CheckMerchantShopState(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*CheckShopResponse, error)
	// * 获取门店,storeId
	GetStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*SStore, error)
	// * 根据主机头获取店铺编号,host
	QueryShopByHost(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int64, error)
	// 获取门店
	// rpc GetOfflineShop(1:sint32 shopId)returns(Shop)
	// 打开或关闭商店
	TurnShop(ctx context.Context, in *TurnShopRequest, opts ...grpc.CallOption) (*Result, error)
	// 保存门店
	SaveShop(ctx context.Context, in *SShop, opts ...grpc.CallOption) (*Result, error)
	// 保存门店
	SaveOfflineShop(ctx context.Context, in *SStore, opts ...grpc.CallOption) (*Result, error)
	// 删除商店
	DeleteStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*Result, error)
}

type shopServiceClient struct {
	cc *grpc.ClientConn
}

func NewShopServiceClient(cc *grpc.ClientConn) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) GetShop(ctx context.Context, in *ShopId, opts ...grpc.CallOption) (*SShop, error) {
	out := new(SShop)
	err := c.cc.Invoke(ctx, "/ShopService/GetShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) CheckMerchantShopState(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*CheckShopResponse, error) {
	out := new(CheckShopResponse)
	err := c.cc.Invoke(ctx, "/ShopService/CheckMerchantShopState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*SStore, error) {
	out := new(SStore)
	err := c.cc.Invoke(ctx, "/ShopService/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) QueryShopByHost(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int64, error) {
	out := new(Int64)
	err := c.cc.Invoke(ctx, "/ShopService/QueryShopByHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) TurnShop(ctx context.Context, in *TurnShopRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/TurnShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SaveShop(ctx context.Context, in *SShop, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/SaveShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SaveOfflineShop(ctx context.Context, in *SStore, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/SaveOfflineShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) DeleteStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/DeleteStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
type ShopServiceServer interface {
	// * 获取店铺,shopId
	GetShop(context.Context, *ShopId) (*SShop, error)
	// rpc GetVendorShop_ (Int64) returns (SShop) {}
	// 检查商户是否开通店铺
	CheckMerchantShopState(context.Context, *MerchantId) (*CheckShopResponse, error)
	// * 获取门店,storeId
	GetStore(context.Context, *StoreId) (*SStore, error)
	// * 根据主机头获取店铺编号,host
	QueryShopByHost(context.Context, *String) (*Int64, error)
	// 获取门店
	// rpc GetOfflineShop(1:sint32 shopId)returns(Shop)
	// 打开或关闭商店
	TurnShop(context.Context, *TurnShopRequest) (*Result, error)
	// 保存门店
	SaveShop(context.Context, *SShop) (*Result, error)
	// 保存门店
	SaveOfflineShop(context.Context, *SStore) (*Result, error)
	// 删除商店
	DeleteStore(context.Context, *StoreId) (*Result, error)
}

func RegisterShopServiceServer(s *grpc.Server, srv ShopServiceServer) {
	s.RegisterService(&_ShopService_serviceDesc, srv)
}

func _ShopService_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShop(ctx, req.(*ShopId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_CheckMerchantShopState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CheckMerchantShopState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/CheckMerchantShopState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CheckMerchantShopState(ctx, req.(*MerchantId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetStore(ctx, req.(*StoreId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_QueryShopByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).QueryShopByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/QueryShopByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).QueryShopByHost(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_TurnShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TurnShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).TurnShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/TurnShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).TurnShop(ctx, req.(*TurnShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SaveShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SShop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SaveShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/SaveShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SaveShop(ctx, req.(*SShop))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SaveOfflineShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SaveOfflineShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/SaveOfflineShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SaveOfflineShop(ctx, req.(*SStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_DeleteStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).DeleteStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/DeleteStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).DeleteStore(ctx, req.(*StoreId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShop",
			Handler:    _ShopService_GetShop_Handler,
		},
		{
			MethodName: "CheckMerchantShopState",
			Handler:    _ShopService_CheckMerchantShopState_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _ShopService_GetStore_Handler,
		},
		{
			MethodName: "QueryShopByHost",
			Handler:    _ShopService_QueryShopByHost_Handler,
		},
		{
			MethodName: "TurnShop",
			Handler:    _ShopService_TurnShop_Handler,
		},
		{
			MethodName: "SaveShop",
			Handler:    _ShopService_SaveShop_Handler,
		},
		{
			MethodName: "SaveOfflineShop",
			Handler:    _ShopService_SaveOfflineShop_Handler,
		},
		{
			MethodName: "DeleteStore",
			Handler:    _ShopService_DeleteStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop_service.proto",
}

func init() { proto.RegisterFile("shop_service.proto", fileDescriptor_shop_service_e0b4ce123fd750f5) }

var fileDescriptor_shop_service_e0b4ce123fd750f5 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x86, 0x93, 0x5c, 0x48, 0x72, 0xa7, 0x97, 0x5b, 0x19, 0xa1, 0x94, 0xe0, 0x47, 0x09, 0x8a,
	0x5d, 0x8d, 0x50, 0xc5, 0x4d, 0x77, 0x55, 0xd0, 0x2c, 0x44, 0x9a, 0xb8, 0x72, 0x23, 0x69, 0x72,
	0x9a, 0x04, 0xd3, 0x39, 0x71, 0x66, 0x52, 0xe8, 0xdf, 0xf3, 0x97, 0xc9, 0x4c, 0x13, 0x28, 0xba,
	0x3a, 0x5f, 0xcf, 0xbc, 0xbc, 0xe7, 0x0c, 0xa1, 0xb2, 0xc4, 0xe6, 0x5d, 0x82, 0xd8, 0x56, 0x19,
	0xb0, 0x46, 0xa0, 0xc2, 0xe0, 0x5f, 0x51, 0xe3, 0x2a, 0xad, 0xbb, 0x6a, 0xb4, 0x01, 0x29, 0xd3,
	0x02, 0xae, 0x0d, 0x99, 0x2b, 0xdc, 0xf7, 0xc3, 0x25, 0x19, 0xbe, 0xb6, 0x82, 0x27, 0x25, 0x36,
	0x31, 0x7c, 0xb6, 0x20, 0x15, 0x1d, 0x11, 0x57, 0x43, 0x51, 0x3e, 0xb6, 0x27, 0xf6, 0x94, 0xc6,
	0x5d, 0x45, 0xff, 0x13, 0x07, 0xf9, 0xd8, 0x99, 0xd8, 0x53, 0x3f, 0x76, 0x90, 0x6b, 0x4e, 0x40,
	0x2a, 0x91, 0x8f, 0xff, 0x4c, 0xec, 0xe9, 0xdf, 0xb8, 0xab, 0x66, 0x5f, 0x0e, 0x19, 0x68, 0xbd,
	0x64, 0x6f, 0x87, 0x9e, 0x10, 0xef, 0x11, 0x94, 0xee, 0x50, 0x8f, 0x25, 0x46, 0x2b, 0x70, 0x59,
	0xa2, 0xb3, 0xd0, 0xa2, 0x73, 0x32, 0xba, 0x2f, 0x21, 0xfb, 0x78, 0x06, 0x91, 0x95, 0x29, 0x37,
	0x5c, 0xa2, 0x52, 0x05, 0x74, 0xc0, 0xfa, 0x5e, 0x94, 0x07, 0x94, 0x19, 0x6a, 0xef, 0x53, 0x36,
	0xc8, 0x25, 0x84, 0x16, 0x3d, 0x27, 0xbe, 0x96, 0x56, 0x28, 0x80, 0xfa, 0xcc, 0xc4, 0x28, 0x0f,
	0x3c, 0x96, 0x98, 0x34, 0xb4, 0xe8, 0x05, 0x19, 0x2e, 0x5b, 0x10, 0x3b, 0xfd, 0x6e, 0xb1, 0x7b,
	0x42, 0xa9, 0xb4, 0x07, 0x25, 0x2a, 0x5e, 0x04, 0x2e, 0x8b, 0xb8, 0xba, 0xbb, 0x0d, 0x2d, 0x7a,
	0x45, 0xfc, 0xfe, 0x08, 0xf4, 0x88, 0xfd, 0xb8, 0x47, 0xe0, 0xb1, 0x18, 0x64, 0x5b, 0xab, 0xd0,
	0xa2, 0xa7, 0xc4, 0x4f, 0xd2, 0x2d, 0x18, 0xb0, 0x5b, 0xe1, 0x70, 0x7c, 0x49, 0x86, 0x7a, 0xfc,
	0xb2, 0x5e, 0xd7, 0x15, 0x87, 0x7e, 0x63, 0xe3, 0xe5, 0x10, 0x0b, 0xc9, 0xe0, 0x01, 0x6a, 0x50,
	0xf0, 0xdb, 0x78, 0xcf, 0x2c, 0xce, 0xc8, 0x71, 0x86, 0x1b, 0x56, 0x54, 0xaa, 0x6c, 0x57, 0xac,
	0xc0, 0x19, 0x32, 0xd1, 0x64, 0x6f, 0x1e, 0x9b, 0x9b, 0x7f, 0x5b, 0xb9, 0x26, 0xdc, 0x7c, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x5c, 0x8c, 0x65, 0xa7, 0xfa, 0x01, 0x00, 0x00,
}
