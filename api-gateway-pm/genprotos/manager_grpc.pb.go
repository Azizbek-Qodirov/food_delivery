// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.1
// source: food-delivery-protos/manager.proto

package genprotos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ProductManagerService_CreateProduct_FullMethodName  = "/delivery.ProductManagerService/CreateProduct"
	ProductManagerService_UpdateProduct_FullMethodName  = "/delivery.ProductManagerService/UpdateProduct"
	ProductManagerService_DeleteProduct_FullMethodName  = "/delivery.ProductManagerService/DeleteProduct"
	ProductManagerService_GetProduct_FullMethodName     = "/delivery.ProductManagerService/GetProduct"
	ProductManagerService_GetAllProducts_FullMethodName = "/delivery.ProductManagerService/GetAllProducts"
)

// ProductManagerServiceClient is the client API for ProductManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductManagerServiceClient interface {
	CreateProduct(ctx context.Context, in *ProductCReq, opts ...grpc.CallOption) (*Void, error)
	UpdateProduct(ctx context.Context, in *ProductUReq, opts ...grpc.CallOption) (*ProductGRes, error)
	DeleteProduct(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	GetProduct(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*ProductGRes, error)
	GetAllProducts(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ProductGARes, error)
}

type productManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductManagerServiceClient(cc grpc.ClientConnInterface) ProductManagerServiceClient {
	return &productManagerServiceClient{cc}
}

func (c *productManagerServiceClient) CreateProduct(ctx context.Context, in *ProductCReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ProductManagerService_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagerServiceClient) UpdateProduct(ctx context.Context, in *ProductUReq, opts ...grpc.CallOption) (*ProductGRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductGRes)
	err := c.cc.Invoke(ctx, ProductManagerService_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagerServiceClient) DeleteProduct(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ProductManagerService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagerServiceClient) GetProduct(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*ProductGRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductGRes)
	err := c.cc.Invoke(ctx, ProductManagerService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagerServiceClient) GetAllProducts(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ProductGARes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductGARes)
	err := c.cc.Invoke(ctx, ProductManagerService_GetAllProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductManagerServiceServer is the server API for ProductManagerService service.
// All implementations must embed UnimplementedProductManagerServiceServer
// for forward compatibility
type ProductManagerServiceServer interface {
	CreateProduct(context.Context, *ProductCReq) (*Void, error)
	UpdateProduct(context.Context, *ProductUReq) (*ProductGRes, error)
	DeleteProduct(context.Context, *ByID) (*Void, error)
	GetProduct(context.Context, *ByID) (*ProductGRes, error)
	GetAllProducts(context.Context, *Void) (*ProductGARes, error)
	mustEmbedUnimplementedProductManagerServiceServer()
}

// UnimplementedProductManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductManagerServiceServer struct {
}

func (UnimplementedProductManagerServiceServer) CreateProduct(context.Context, *ProductCReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductManagerServiceServer) UpdateProduct(context.Context, *ProductUReq) (*ProductGRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductManagerServiceServer) DeleteProduct(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductManagerServiceServer) GetProduct(context.Context, *ByID) (*ProductGRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductManagerServiceServer) GetAllProducts(context.Context, *Void) (*ProductGARes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedProductManagerServiceServer) mustEmbedUnimplementedProductManagerServiceServer() {}

// UnsafeProductManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductManagerServiceServer will
// result in compilation errors.
type UnsafeProductManagerServiceServer interface {
	mustEmbedUnimplementedProductManagerServiceServer()
}

func RegisterProductManagerServiceServer(s grpc.ServiceRegistrar, srv ProductManagerServiceServer) {
	s.RegisterService(&ProductManagerService_ServiceDesc, srv)
}

func _ProductManagerService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagerServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductManagerService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagerServiceServer).CreateProduct(ctx, req.(*ProductCReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagerService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductUReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagerServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductManagerService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagerServiceServer).UpdateProduct(ctx, req.(*ProductUReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagerService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagerServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductManagerService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagerServiceServer).DeleteProduct(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagerService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagerServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductManagerService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagerServiceServer).GetProduct(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagerService_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagerServiceServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductManagerService_GetAllProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagerServiceServer).GetAllProducts(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductManagerService_ServiceDesc is the grpc.ServiceDesc for ProductManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "delivery.ProductManagerService",
	HandlerType: (*ProductManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductManagerService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductManagerService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductManagerService_DeleteProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductManagerService_GetProduct_Handler,
		},
		{
			MethodName: "GetAllProducts",
			Handler:    _ProductManagerService_GetAllProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "food-delivery-protos/manager.proto",
}