// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: reservation_service/reservation_service.proto

package reservation_service

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

const (
	ReservationService_Create_FullMethodName                  = "/ReservationService/Create"
	ReservationService_GetAll_FullMethodName                  = "/ReservationService/GetAll"
	ReservationService_GetWaitingReservations_FullMethodName  = "/ReservationService/GetWaitingReservations"
	ReservationService_ConfirmReservation_FullMethodName      = "/ReservationService/ConfirmReservation"
	ReservationService_Delete_FullMethodName                  = "/ReservationService/Delete"
	ReservationService_GetHostAnalytics_FullMethodName        = "/ReservationService/GetHostAnalytics"
	ReservationService_HasLeftoverReservations_FullMethodName = "/ReservationService/HasLeftoverReservations"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	Create(ctx context.Context, in *ReservationCreateRequest, opts ...grpc.CallOption) (*ReservationCreateResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetWaitingReservations(ctx context.Context, in *WaitingReservationsForHostRequest, opts ...grpc.CallOption) (*WaitingReservationsForHostResponse, error)
	ConfirmReservation(ctx context.Context, in *ConfirmReservationRequest, opts ...grpc.CallOption) (*ConfirmReservationResponse, error)
	Delete(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error)
	GetHostAnalytics(ctx context.Context, in *HostAnalyticsRequest, opts ...grpc.CallOption) (*HostAnalyticsResponse, error)
	HasLeftoverReservations(ctx context.Context, in *LeftoverReservationsRequest, opts ...grpc.CallOption) (*LeftoverReservationsResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) Create(ctx context.Context, in *ReservationCreateRequest, opts ...grpc.CallOption) (*ReservationCreateResponse, error) {
	out := new(ReservationCreateResponse)
	err := c.cc.Invoke(ctx, ReservationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetWaitingReservations(ctx context.Context, in *WaitingReservationsForHostRequest, opts ...grpc.CallOption) (*WaitingReservationsForHostResponse, error) {
	out := new(WaitingReservationsForHostResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetWaitingReservations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) ConfirmReservation(ctx context.Context, in *ConfirmReservationRequest, opts ...grpc.CallOption) (*ConfirmReservationResponse, error) {
	out := new(ConfirmReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_ConfirmReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Delete(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error) {
	out := new(DeleteReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetHostAnalytics(ctx context.Context, in *HostAnalyticsRequest, opts ...grpc.CallOption) (*HostAnalyticsResponse, error) {
	out := new(HostAnalyticsResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetHostAnalytics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) HasLeftoverReservations(ctx context.Context, in *LeftoverReservationsRequest, opts ...grpc.CallOption) (*LeftoverReservationsResponse, error) {
	out := new(LeftoverReservationsResponse)
	err := c.cc.Invoke(ctx, ReservationService_HasLeftoverReservations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	Create(context.Context, *ReservationCreateRequest) (*ReservationCreateResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetWaitingReservations(context.Context, *WaitingReservationsForHostRequest) (*WaitingReservationsForHostResponse, error)
	ConfirmReservation(context.Context, *ConfirmReservationRequest) (*ConfirmReservationResponse, error)
	Delete(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error)
	GetHostAnalytics(context.Context, *HostAnalyticsRequest) (*HostAnalyticsResponse, error)
	HasLeftoverReservations(context.Context, *LeftoverReservationsRequest) (*LeftoverReservationsResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) Create(context.Context, *ReservationCreateRequest) (*ReservationCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedReservationServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedReservationServiceServer) GetWaitingReservations(context.Context, *WaitingReservationsForHostRequest) (*WaitingReservationsForHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaitingReservations not implemented")
}
func (UnimplementedReservationServiceServer) ConfirmReservation(context.Context, *ConfirmReservationRequest) (*ConfirmReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmReservation not implemented")
}
func (UnimplementedReservationServiceServer) Delete(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedReservationServiceServer) GetHostAnalytics(context.Context, *HostAnalyticsRequest) (*HostAnalyticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostAnalytics not implemented")
}
func (UnimplementedReservationServiceServer) HasLeftoverReservations(context.Context, *LeftoverReservationsRequest) (*LeftoverReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasLeftoverReservations not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Create(ctx, req.(*ReservationCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetWaitingReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitingReservationsForHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetWaitingReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetWaitingReservations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetWaitingReservations(ctx, req.(*WaitingReservationsForHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_ConfirmReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ConfirmReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_ConfirmReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ConfirmReservation(ctx, req.(*ConfirmReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Delete(ctx, req.(*DeleteReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetHostAnalytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostAnalyticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetHostAnalytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetHostAnalytics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetHostAnalytics(ctx, req.(*HostAnalyticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_HasLeftoverReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeftoverReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).HasLeftoverReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_HasLeftoverReservations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).HasLeftoverReservations(ctx, req.(*LeftoverReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReservationService_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ReservationService_GetAll_Handler,
		},
		{
			MethodName: "GetWaitingReservations",
			Handler:    _ReservationService_GetWaitingReservations_Handler,
		},
		{
			MethodName: "ConfirmReservation",
			Handler:    _ReservationService_ConfirmReservation_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReservationService_Delete_Handler,
		},
		{
			MethodName: "GetHostAnalytics",
			Handler:    _ReservationService_GetHostAnalytics_Handler,
		},
		{
			MethodName: "HasLeftoverReservations",
			Handler:    _ReservationService_HasLeftoverReservations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation_service/reservation_service.proto",
}
