// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: reservation/reservation-service.proto

package reservation

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
	ReservationService_Create_FullMethodName = "/ReservationService/Create"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	Create(ctx context.Context, in *ReservationCreateRequest, opts ...grpc.CallOption) (*ReservationCreateResponse, error)
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

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	Create(context.Context, *ReservationCreateRequest) (*ReservationCreateResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) Create(context.Context, *ReservationCreateRequest) (*ReservationCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation/reservation-service.proto",
}
