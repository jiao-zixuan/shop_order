// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package order

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	// 获取用户的购物车信息
	CartItemList(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*CartItemListResponse, error)
	// 添加商品到购物车
	CreateCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*ShopCartInfoResponse, error)
	// 修改购物车记录 主要修改他是否勾选和购买数量
	UpdateCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*OrderEmpty, error)
	// 删除购物车 删除多个商品的话用 切片 repeated
	DeleteOrders(ctx context.Context, in *DeleteOrdersRequest, opts ...grpc.CallOption) (*OrderEmpty, error)
	// 创建订单
	Create(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfoResponse, error)
	// 订单列表
	OrderList(ctx context.Context, in *OrderFilterRequest, opts ...grpc.CallOption) (*OrderListResponse, error)
	// 订单详情
	OrderDetail(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfoDetailResponse, error)
	// 修改订单状态
	UpdateOrderStatus(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*OrderEmpty, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) CartItemList(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*CartItemListResponse, error) {
	out := new(CartItemListResponse)
	err := c.cc.Invoke(ctx, "/Order/CartItemList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*ShopCartInfoResponse, error) {
	out := new(ShopCartInfoResponse)
	err := c.cc.Invoke(ctx, "/Order/CreateCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*OrderEmpty, error) {
	out := new(OrderEmpty)
	err := c.cc.Invoke(ctx, "/Order/UpdateCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) DeleteOrders(ctx context.Context, in *DeleteOrdersRequest, opts ...grpc.CallOption) (*OrderEmpty, error) {
	out := new(OrderEmpty)
	err := c.cc.Invoke(ctx, "/Order/DeleteOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) Create(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfoResponse, error) {
	out := new(OrderInfoResponse)
	err := c.cc.Invoke(ctx, "/Order/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderList(ctx context.Context, in *OrderFilterRequest, opts ...grpc.CallOption) (*OrderListResponse, error) {
	out := new(OrderListResponse)
	err := c.cc.Invoke(ctx, "/Order/OrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderDetail(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfoDetailResponse, error) {
	out := new(OrderInfoDetailResponse)
	err := c.cc.Invoke(ctx, "/Order/OrderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateOrderStatus(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*OrderEmpty, error) {
	out := new(OrderEmpty)
	err := c.cc.Invoke(ctx, "/Order/UpdateOrderStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	// 获取用户的购物车信息
	CartItemList(context.Context, *UserInfo) (*CartItemListResponse, error)
	// 添加商品到购物车
	CreateCartItem(context.Context, *CartItemRequest) (*ShopCartInfoResponse, error)
	// 修改购物车记录 主要修改他是否勾选和购买数量
	UpdateCartItem(context.Context, *CartItemRequest) (*OrderEmpty, error)
	// 删除购物车 删除多个商品的话用 切片 repeated
	DeleteOrders(context.Context, *DeleteOrdersRequest) (*OrderEmpty, error)
	// 创建订单
	Create(context.Context, *OrderRequest) (*OrderInfoResponse, error)
	// 订单列表
	OrderList(context.Context, *OrderFilterRequest) (*OrderListResponse, error)
	// 订单详情
	OrderDetail(context.Context, *OrderRequest) (*OrderInfoDetailResponse, error)
	// 修改订单状态
	UpdateOrderStatus(context.Context, *UpdateOrderRequest) (*OrderEmpty, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) CartItemList(context.Context, *UserInfo) (*CartItemListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemList not implemented")
}
func (UnimplementedOrderServer) CreateCartItem(context.Context, *CartItemRequest) (*ShopCartInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCartItem not implemented")
}
func (UnimplementedOrderServer) UpdateCartItem(context.Context, *CartItemRequest) (*OrderEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCartItem not implemented")
}
func (UnimplementedOrderServer) DeleteOrders(context.Context, *DeleteOrdersRequest) (*OrderEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrders not implemented")
}
func (UnimplementedOrderServer) Create(context.Context, *OrderRequest) (*OrderInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderServer) OrderList(context.Context, *OrderFilterRequest) (*OrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedOrderServer) OrderDetail(context.Context, *OrderRequest) (*OrderInfoDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDetail not implemented")
}
func (UnimplementedOrderServer) UpdateOrderStatus(context.Context, *UpdateOrderRequest) (*OrderEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_CartItemList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CartItemList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/CartItemList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CartItemList(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/CreateCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateCartItem(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/UpdateCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateCartItem(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_DeleteOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).DeleteOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/DeleteOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).DeleteOrders(ctx, req.(*DeleteOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Create(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/OrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderList(ctx, req.(*OrderFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/OrderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderDetail(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/UpdateOrderStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateOrderStatus(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CartItemList",
			Handler:    _Order_CartItemList_Handler,
		},
		{
			MethodName: "CreateCartItem",
			Handler:    _Order_CreateCartItem_Handler,
		},
		{
			MethodName: "UpdateCartItem",
			Handler:    _Order_UpdateCartItem_Handler,
		},
		{
			MethodName: "DeleteOrders",
			Handler:    _Order_DeleteOrders_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Order_Create_Handler,
		},
		{
			MethodName: "OrderList",
			Handler:    _Order_OrderList_Handler,
		},
		{
			MethodName: "OrderDetail",
			Handler:    _Order_OrderDetail_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _Order_UpdateOrderStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
