// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: accommodation_service/accommodation_service.proto

package accommodation_service

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
	AccommodationService_Create_FullMethodName                   = "/AccommodationService/Create"
	AccommodationService_Update_FullMethodName                   = "/AccommodationService/Update"
	AccommodationService_Delete_FullMethodName                   = "/AccommodationService/Delete"
	AccommodationService_GetAll_FullMethodName                   = "/AccommodationService/GetAll"
	AccommodationService_SearchAccommodations_FullMethodName     = "/AccommodationService/SearchAccommodations"
	AccommodationService_GetById_FullMethodName                  = "/AccommodationService/GetById"
	AccommodationService_CreateAppointment_FullMethodName        = "/AccommodationService/CreateAppointment"
	AccommodationService_UpdateAppointment_FullMethodName        = "/AccommodationService/UpdateAppointment"
	AccommodationService_ValidateReservation_FullMethodName      = "/AccommodationService/ValidateReservation"
	AccommodationService_IsAutomaticConfirmation_FullMethodName  = "/AccommodationService/IsAutomaticConfirmation"
	AccommodationService_DeleteHostAccommodations_FullMethodName = "/AccommodationService/DeleteHostAccommodations"
)

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	Create(ctx context.Context, in *AccommodationCreateRequest, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *SingleAccommodation, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *AccommodationIdRequest, opts ...grpc.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *GetAllAccommodationRequest, opts ...grpc.CallOption) (*GetAllAccommodationResponse, error)
	SearchAccommodations(ctx context.Context, in *SearchAccommodationsRequest, opts ...grpc.CallOption) (*SearchAccommodationsResponse, error)
	GetById(ctx context.Context, in *AccommodationIdRequest, opts ...grpc.CallOption) (*SingleAccommodation, error)
	CreateAppointment(ctx context.Context, in *SingleAppointment, opts ...grpc.CallOption) (*Response, error)
	UpdateAppointment(ctx context.Context, in *UpdateAppointmentRequest, opts ...grpc.CallOption) (*Response, error)
	ValidateReservation(ctx context.Context, in *ValidateReservationRequest, opts ...grpc.CallOption) (*ValidateReservationResponse, error)
	IsAutomaticConfirmation(ctx context.Context, in *AccommodationIdRequest, opts ...grpc.CallOption) (*IsAutomaticConfirmationResponse, error)
	DeleteHostAccommodations(ctx context.Context, in *DeleteHostAccommodationsRequest, opts ...grpc.CallOption) (*DeleteHostAccommodationsResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) Create(ctx context.Context, in *AccommodationCreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, AccommodationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Update(ctx context.Context, in *SingleAccommodation, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, AccommodationService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Delete(ctx context.Context, in *AccommodationIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, AccommodationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAll(ctx context.Context, in *GetAllAccommodationRequest, opts ...grpc.CallOption) (*GetAllAccommodationResponse, error) {
	out := new(GetAllAccommodationResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) SearchAccommodations(ctx context.Context, in *SearchAccommodationsRequest, opts ...grpc.CallOption) (*SearchAccommodationsResponse, error) {
	out := new(SearchAccommodationsResponse)
	err := c.cc.Invoke(ctx, AccommodationService_SearchAccommodations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetById(ctx context.Context, in *AccommodationIdRequest, opts ...grpc.CallOption) (*SingleAccommodation, error) {
	out := new(SingleAccommodation)
	err := c.cc.Invoke(ctx, AccommodationService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateAppointment(ctx context.Context, in *SingleAppointment, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, AccommodationService_CreateAppointment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) UpdateAppointment(ctx context.Context, in *UpdateAppointmentRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, AccommodationService_UpdateAppointment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) ValidateReservation(ctx context.Context, in *ValidateReservationRequest, opts ...grpc.CallOption) (*ValidateReservationResponse, error) {
	out := new(ValidateReservationResponse)
	err := c.cc.Invoke(ctx, AccommodationService_ValidateReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) IsAutomaticConfirmation(ctx context.Context, in *AccommodationIdRequest, opts ...grpc.CallOption) (*IsAutomaticConfirmationResponse, error) {
	out := new(IsAutomaticConfirmationResponse)
	err := c.cc.Invoke(ctx, AccommodationService_IsAutomaticConfirmation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) DeleteHostAccommodations(ctx context.Context, in *DeleteHostAccommodationsRequest, opts ...grpc.CallOption) (*DeleteHostAccommodationsResponse, error) {
	out := new(DeleteHostAccommodationsResponse)
	err := c.cc.Invoke(ctx, AccommodationService_DeleteHostAccommodations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	Create(context.Context, *AccommodationCreateRequest) (*Response, error)
	Update(context.Context, *SingleAccommodation) (*Response, error)
	Delete(context.Context, *AccommodationIdRequest) (*Response, error)
	GetAll(context.Context, *GetAllAccommodationRequest) (*GetAllAccommodationResponse, error)
	SearchAccommodations(context.Context, *SearchAccommodationsRequest) (*SearchAccommodationsResponse, error)
	GetById(context.Context, *AccommodationIdRequest) (*SingleAccommodation, error)
	CreateAppointment(context.Context, *SingleAppointment) (*Response, error)
	UpdateAppointment(context.Context, *UpdateAppointmentRequest) (*Response, error)
	ValidateReservation(context.Context, *ValidateReservationRequest) (*ValidateReservationResponse, error)
	IsAutomaticConfirmation(context.Context, *AccommodationIdRequest) (*IsAutomaticConfirmationResponse, error)
	DeleteHostAccommodations(context.Context, *DeleteHostAccommodationsRequest) (*DeleteHostAccommodationsResponse, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) Create(context.Context, *AccommodationCreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccommodationServiceServer) Update(context.Context, *SingleAccommodation) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAccommodationServiceServer) Delete(context.Context, *AccommodationIdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAll(context.Context, *GetAllAccommodationRequest) (*GetAllAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAccommodationServiceServer) SearchAccommodations(context.Context, *SearchAccommodationsRequest) (*SearchAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccommodations not implemented")
}
func (UnimplementedAccommodationServiceServer) GetById(context.Context, *AccommodationIdRequest) (*SingleAccommodation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateAppointment(context.Context, *SingleAppointment) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppointment not implemented")
}
func (UnimplementedAccommodationServiceServer) UpdateAppointment(context.Context, *UpdateAppointmentRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppointment not implemented")
}
func (UnimplementedAccommodationServiceServer) ValidateReservation(context.Context, *ValidateReservationRequest) (*ValidateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateReservation not implemented")
}
func (UnimplementedAccommodationServiceServer) IsAutomaticConfirmation(context.Context, *AccommodationIdRequest) (*IsAutomaticConfirmationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAutomaticConfirmation not implemented")
}
func (UnimplementedAccommodationServiceServer) DeleteHostAccommodations(context.Context, *DeleteHostAccommodationsRequest) (*DeleteHostAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHostAccommodations not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Create(ctx, req.(*AccommodationCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleAccommodation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Update(ctx, req.(*SingleAccommodation))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Delete(ctx, req.(*AccommodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAll(ctx, req.(*GetAllAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_SearchAccommodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAccommodationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).SearchAccommodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_SearchAccommodations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).SearchAccommodations(ctx, req.(*SearchAccommodationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetById(ctx, req.(*AccommodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleAppointment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_CreateAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAppointment(ctx, req.(*SingleAppointment))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_UpdateAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).UpdateAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_UpdateAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).UpdateAppointment(ctx, req.(*UpdateAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_ValidateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).ValidateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_ValidateReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).ValidateReservation(ctx, req.(*ValidateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_IsAutomaticConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).IsAutomaticConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_IsAutomaticConfirmation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).IsAutomaticConfirmation(ctx, req.(*AccommodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_DeleteHostAccommodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHostAccommodationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).DeleteHostAccommodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_DeleteHostAccommodations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).DeleteHostAccommodations(ctx, req.(*DeleteHostAccommodationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AccommodationService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AccommodationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccommodationService_Delete_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AccommodationService_GetAll_Handler,
		},
		{
			MethodName: "SearchAccommodations",
			Handler:    _AccommodationService_SearchAccommodations_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _AccommodationService_GetById_Handler,
		},
		{
			MethodName: "CreateAppointment",
			Handler:    _AccommodationService_CreateAppointment_Handler,
		},
		{
			MethodName: "UpdateAppointment",
			Handler:    _AccommodationService_UpdateAppointment_Handler,
		},
		{
			MethodName: "ValidateReservation",
			Handler:    _AccommodationService_ValidateReservation_Handler,
		},
		{
			MethodName: "IsAutomaticConfirmation",
			Handler:    _AccommodationService_IsAutomaticConfirmation_Handler,
		},
		{
			MethodName: "DeleteHostAccommodations",
			Handler:    _AccommodationService_DeleteHostAccommodations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation_service/accommodation_service.proto",
}
