// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: reservation_service/reservation_service.proto

package reservation_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReservationCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accommodation int32  `protobuf:"varint,1,opt,name=accommodation,proto3" json:"accommodation,omitempty"`
	Offer         int32  `protobuf:"varint,2,opt,name=offer,proto3" json:"offer,omitempty"`
	DateFrom      string `protobuf:"bytes,3,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo        string `protobuf:"bytes,4,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
	Guests        int32  `protobuf:"varint,5,opt,name=guests,proto3" json:"guests,omitempty"`
}

func (x *ReservationCreateRequest) Reset() {
	*x = ReservationCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_reservation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationCreateRequest) ProtoMessage() {}

func (x *ReservationCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_reservation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationCreateRequest.ProtoReflect.Descriptor instead.
func (*ReservationCreateRequest) Descriptor() ([]byte, []int) {
	return file_reservation_service_reservation_service_proto_rawDescGZIP(), []int{0}
}

func (x *ReservationCreateRequest) GetAccommodation() int32 {
	if x != nil {
		return x.Accommodation
	}
	return 0
}

func (x *ReservationCreateRequest) GetOffer() int32 {
	if x != nil {
		return x.Offer
	}
	return 0
}

func (x *ReservationCreateRequest) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *ReservationCreateRequest) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *ReservationCreateRequest) GetGuests() int32 {
	if x != nil {
		return x.Guests
	}
	return 0
}

type ReservationCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReservationCreateResponse) Reset() {
	*x = ReservationCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_reservation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationCreateResponse) ProtoMessage() {}

func (x *ReservationCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_reservation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationCreateResponse.ProtoReflect.Descriptor instead.
func (*ReservationCreateResponse) Descriptor() ([]byte, []int) {
	return file_reservation_service_reservation_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReservationCreateResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Accommodation int32  `protobuf:"varint,2,opt,name=accommodation,proto3" json:"accommodation,omitempty"`
	Offer         int32  `protobuf:"varint,3,opt,name=offer,proto3" json:"offer,omitempty"`
	DateFrom      string `protobuf:"bytes,4,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo        string `protobuf:"bytes,5,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
	Guests        int32  `protobuf:"varint,6,opt,name=guests,proto3" json:"guests,omitempty"`
	Status        int32  `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_reservation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_reservation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_reservation_service_reservation_service_proto_rawDescGZIP(), []int{2}
}

func (x *Reservation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reservation) GetAccommodation() int32 {
	if x != nil {
		return x.Accommodation
	}
	return 0
}

func (x *Reservation) GetOffer() int32 {
	if x != nil {
		return x.Offer
	}
	return 0
}

func (x *Reservation) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *Reservation) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *Reservation) GetGuests() int32 {
	if x != nil {
		return x.Guests
	}
	return 0
}

func (x *Reservation) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_reservation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_reservation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_reservation_service_reservation_service_proto_rawDescGZIP(), []int{3}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*Reservation `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_service_reservation_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_service_reservation_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_reservation_service_reservation_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllResponse) GetReservations() []*Reservation {
	if x != nil {
		return x.Reservations
	}
	return nil
}

var File_reservation_service_reservation_service_proto protoreflect.FileDescriptor

var file_reservation_service_reservation_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01,
	0x0a, 0x18, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x22, 0x2f, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xbd, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xb1, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x59, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x32, 0x5a, 0x30,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reservation_service_reservation_service_proto_rawDescOnce sync.Once
	file_reservation_service_reservation_service_proto_rawDescData = file_reservation_service_reservation_service_proto_rawDesc
)

func file_reservation_service_reservation_service_proto_rawDescGZIP() []byte {
	file_reservation_service_reservation_service_proto_rawDescOnce.Do(func() {
		file_reservation_service_reservation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservation_service_reservation_service_proto_rawDescData)
	})
	return file_reservation_service_reservation_service_proto_rawDescData
}

var file_reservation_service_reservation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_reservation_service_reservation_service_proto_goTypes = []interface{}{
	(*ReservationCreateRequest)(nil),  // 0: ReservationCreateRequest
	(*ReservationCreateResponse)(nil), // 1: ReservationCreateResponse
	(*Reservation)(nil),               // 2: Reservation
	(*GetAllRequest)(nil),             // 3: GetAllRequest
	(*GetAllResponse)(nil),            // 4: GetAllResponse
}
var file_reservation_service_reservation_service_proto_depIdxs = []int32{
	2, // 0: GetAllResponse.reservations:type_name -> Reservation
	0, // 1: ReservationService.Create:input_type -> ReservationCreateRequest
	3, // 2: ReservationService.GetAll:input_type -> GetAllRequest
	1, // 3: ReservationService.Create:output_type -> ReservationCreateResponse
	4, // 4: ReservationService.GetAll:output_type -> GetAllResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reservation_service_reservation_service_proto_init() }
func file_reservation_service_reservation_service_proto_init() {
	if File_reservation_service_reservation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservation_service_reservation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reservation_service_reservation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationCreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reservation_service_reservation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reservation_service_reservation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reservation_service_reservation_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reservation_service_reservation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservation_service_reservation_service_proto_goTypes,
		DependencyIndexes: file_reservation_service_reservation_service_proto_depIdxs,
		MessageInfos:      file_reservation_service_reservation_service_proto_msgTypes,
	}.Build()
	File_reservation_service_reservation_service_proto = out.File
	file_reservation_service_reservation_service_proto_rawDesc = nil
	file_reservation_service_reservation_service_proto_goTypes = nil
	file_reservation_service_reservation_service_proto_depIdxs = nil
}
