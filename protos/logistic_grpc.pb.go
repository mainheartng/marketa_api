// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: protos/logistic.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LogisticsService_GetShippingRates_FullMethodName           = "/logistic.LogisticsService/GetShippingRates"
	LogisticsService_CreateShipment_FullMethodName             = "/logistic.LogisticsService/CreateShipment"
	LogisticsService_TrackShipment_FullMethodName              = "/logistic.LogisticsService/TrackShipment"
	LogisticsService_UpdateShipmentStatus_FullMethodName       = "/logistic.LogisticsService/UpdateShipmentStatus"
	LogisticsService_ListShipments_FullMethodName              = "/logistic.LogisticsService/ListShipments"
	LogisticsService_ConfigureLogisticsProvider_FullMethodName = "/logistic.LogisticsService/ConfigureLogisticsProvider"
	LogisticsService_CreateReturnLabel_FullMethodName          = "/logistic.LogisticsService/CreateReturnLabel"
)

// LogisticsServiceClient is the client API for LogisticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogisticsServiceClient interface {
	GetShippingRates(ctx context.Context, in *GetShippingRatesRequest, opts ...grpc.CallOption) (*GetShippingRatesResponse, error)
	CreateShipment(ctx context.Context, in *CreateShipmentRequest, opts ...grpc.CallOption) (*CreateShipmentResponse, error)
	TrackShipment(ctx context.Context, in *TrackShipmentRequest, opts ...grpc.CallOption) (*TrackShipmentResponse, error)
	UpdateShipmentStatus(ctx context.Context, in *UpdateShipmentStatusRequest, opts ...grpc.CallOption) (*UpdateShipmentStatusResponse, error)
	ListShipments(ctx context.Context, in *ListShipmentsRequest, opts ...grpc.CallOption) (*ListShipmentsResponse, error)
	ConfigureLogisticsProvider(ctx context.Context, in *ConfigureLogisticsProviderRequest, opts ...grpc.CallOption) (*ConfigureLogisticsProviderResponse, error)
	CreateReturnLabel(ctx context.Context, in *CreateReturnLabelRequest, opts ...grpc.CallOption) (*CreateReturnLabelResponse, error)
}

type logisticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogisticsServiceClient(cc grpc.ClientConnInterface) LogisticsServiceClient {
	return &logisticsServiceClient{cc}
}

func (c *logisticsServiceClient) GetShippingRates(ctx context.Context, in *GetShippingRatesRequest, opts ...grpc.CallOption) (*GetShippingRatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetShippingRatesResponse)
	err := c.cc.Invoke(ctx, LogisticsService_GetShippingRates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsServiceClient) CreateShipment(ctx context.Context, in *CreateShipmentRequest, opts ...grpc.CallOption) (*CreateShipmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShipmentResponse)
	err := c.cc.Invoke(ctx, LogisticsService_CreateShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsServiceClient) TrackShipment(ctx context.Context, in *TrackShipmentRequest, opts ...grpc.CallOption) (*TrackShipmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TrackShipmentResponse)
	err := c.cc.Invoke(ctx, LogisticsService_TrackShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsServiceClient) UpdateShipmentStatus(ctx context.Context, in *UpdateShipmentStatusRequest, opts ...grpc.CallOption) (*UpdateShipmentStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateShipmentStatusResponse)
	err := c.cc.Invoke(ctx, LogisticsService_UpdateShipmentStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsServiceClient) ListShipments(ctx context.Context, in *ListShipmentsRequest, opts ...grpc.CallOption) (*ListShipmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListShipmentsResponse)
	err := c.cc.Invoke(ctx, LogisticsService_ListShipments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsServiceClient) ConfigureLogisticsProvider(ctx context.Context, in *ConfigureLogisticsProviderRequest, opts ...grpc.CallOption) (*ConfigureLogisticsProviderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfigureLogisticsProviderResponse)
	err := c.cc.Invoke(ctx, LogisticsService_ConfigureLogisticsProvider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsServiceClient) CreateReturnLabel(ctx context.Context, in *CreateReturnLabelRequest, opts ...grpc.CallOption) (*CreateReturnLabelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateReturnLabelResponse)
	err := c.cc.Invoke(ctx, LogisticsService_CreateReturnLabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogisticsServiceServer is the server API for LogisticsService service.
// All implementations must embed UnimplementedLogisticsServiceServer
// for forward compatibility.
type LogisticsServiceServer interface {
	GetShippingRates(context.Context, *GetShippingRatesRequest) (*GetShippingRatesResponse, error)
	CreateShipment(context.Context, *CreateShipmentRequest) (*CreateShipmentResponse, error)
	TrackShipment(context.Context, *TrackShipmentRequest) (*TrackShipmentResponse, error)
	UpdateShipmentStatus(context.Context, *UpdateShipmentStatusRequest) (*UpdateShipmentStatusResponse, error)
	ListShipments(context.Context, *ListShipmentsRequest) (*ListShipmentsResponse, error)
	ConfigureLogisticsProvider(context.Context, *ConfigureLogisticsProviderRequest) (*ConfigureLogisticsProviderResponse, error)
	CreateReturnLabel(context.Context, *CreateReturnLabelRequest) (*CreateReturnLabelResponse, error)
	mustEmbedUnimplementedLogisticsServiceServer()
}

// UnimplementedLogisticsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogisticsServiceServer struct{}

func (UnimplementedLogisticsServiceServer) GetShippingRates(context.Context, *GetShippingRatesRequest) (*GetShippingRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShippingRates not implemented")
}
func (UnimplementedLogisticsServiceServer) CreateShipment(context.Context, *CreateShipmentRequest) (*CreateShipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShipment not implemented")
}
func (UnimplementedLogisticsServiceServer) TrackShipment(context.Context, *TrackShipmentRequest) (*TrackShipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackShipment not implemented")
}
func (UnimplementedLogisticsServiceServer) UpdateShipmentStatus(context.Context, *UpdateShipmentStatusRequest) (*UpdateShipmentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShipmentStatus not implemented")
}
func (UnimplementedLogisticsServiceServer) ListShipments(context.Context, *ListShipmentsRequest) (*ListShipmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShipments not implemented")
}
func (UnimplementedLogisticsServiceServer) ConfigureLogisticsProvider(context.Context, *ConfigureLogisticsProviderRequest) (*ConfigureLogisticsProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureLogisticsProvider not implemented")
}
func (UnimplementedLogisticsServiceServer) CreateReturnLabel(context.Context, *CreateReturnLabelRequest) (*CreateReturnLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReturnLabel not implemented")
}
func (UnimplementedLogisticsServiceServer) mustEmbedUnimplementedLogisticsServiceServer() {}
func (UnimplementedLogisticsServiceServer) testEmbeddedByValue()                          {}

// UnsafeLogisticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogisticsServiceServer will
// result in compilation errors.
type UnsafeLogisticsServiceServer interface {
	mustEmbedUnimplementedLogisticsServiceServer()
}

func RegisterLogisticsServiceServer(s grpc.ServiceRegistrar, srv LogisticsServiceServer) {
	// If the following call pancis, it indicates UnimplementedLogisticsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LogisticsService_ServiceDesc, srv)
}

func _LogisticsService_GetShippingRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShippingRatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticsServiceServer).GetShippingRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogisticsService_GetShippingRates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticsServiceServer).GetShippingRates(ctx, req.(*GetShippingRatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticsService_CreateShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticsServiceServer).CreateShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogisticsService_CreateShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticsServiceServer).CreateShipment(ctx, req.(*CreateShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticsService_TrackShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticsServiceServer).TrackShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogisticsService_TrackShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticsServiceServer).TrackShipment(ctx, req.(*TrackShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticsService_UpdateShipmentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShipmentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticsServiceServer).UpdateShipmentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogisticsService_UpdateShipmentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticsServiceServer).UpdateShipmentStatus(ctx, req.(*UpdateShipmentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticsService_ListShipments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShipmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticsServiceServer).ListShipments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogisticsService_ListShipments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticsServiceServer).ListShipments(ctx, req.(*ListShipmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticsService_ConfigureLogisticsProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureLogisticsProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticsServiceServer).ConfigureLogisticsProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogisticsService_ConfigureLogisticsProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticsServiceServer).ConfigureLogisticsProvider(ctx, req.(*ConfigureLogisticsProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticsService_CreateReturnLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReturnLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticsServiceServer).CreateReturnLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogisticsService_CreateReturnLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticsServiceServer).CreateReturnLabel(ctx, req.(*CreateReturnLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogisticsService_ServiceDesc is the grpc.ServiceDesc for LogisticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogisticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logistic.LogisticsService",
	HandlerType: (*LogisticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShippingRates",
			Handler:    _LogisticsService_GetShippingRates_Handler,
		},
		{
			MethodName: "CreateShipment",
			Handler:    _LogisticsService_CreateShipment_Handler,
		},
		{
			MethodName: "TrackShipment",
			Handler:    _LogisticsService_TrackShipment_Handler,
		},
		{
			MethodName: "UpdateShipmentStatus",
			Handler:    _LogisticsService_UpdateShipmentStatus_Handler,
		},
		{
			MethodName: "ListShipments",
			Handler:    _LogisticsService_ListShipments_Handler,
		},
		{
			MethodName: "ConfigureLogisticsProvider",
			Handler:    _LogisticsService_ConfigureLogisticsProvider_Handler,
		},
		{
			MethodName: "CreateReturnLabel",
			Handler:    _LogisticsService_CreateReturnLabel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/logistic.proto",
}
