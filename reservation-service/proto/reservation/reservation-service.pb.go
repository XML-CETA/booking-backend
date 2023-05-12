// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: reservation/reservation-service.proto

package reservation

import (
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
		mi := &file_reservation_reservation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationCreateRequest) ProtoMessage() {}

func (x *ReservationCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_reservation_service_proto_msgTypes[0]
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
	return file_reservation_reservation_service_proto_rawDescGZIP(), []int{0}
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

	Greeting string `protobuf:"bytes,2,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *ReservationCreateResponse) Reset() {
	*x = ReservationCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_reservation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationCreateResponse) ProtoMessage() {}

func (x *ReservationCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_reservation_service_proto_msgTypes[1]
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
	return file_reservation_reservation_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReservationCreateResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

var File_reservation_reservation_service_proto protoreflect.FileDescriptor

var file_reservation_reservation_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x37, 0x0a, 0x19,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x32, 0x57, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13,
	0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reservation_reservation_service_proto_rawDescOnce sync.Once
	file_reservation_reservation_service_proto_rawDescData = file_reservation_reservation_service_proto_rawDesc
)

func file_reservation_reservation_service_proto_rawDescGZIP() []byte {
	file_reservation_reservation_service_proto_rawDescOnce.Do(func() {
		file_reservation_reservation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservation_reservation_service_proto_rawDescData)
	})
	return file_reservation_reservation_service_proto_rawDescData
}

var file_reservation_reservation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_reservation_reservation_service_proto_goTypes = []interface{}{
	(*ReservationCreateRequest)(nil),  // 0: ReservationCreateRequest
	(*ReservationCreateResponse)(nil), // 1: ReservationCreateResponse
}
var file_reservation_reservation_service_proto_depIdxs = []int32{
	0, // 0: ReservationService.Create:input_type -> ReservationCreateRequest
	1, // 1: ReservationService.Create:output_type -> ReservationCreateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reservation_reservation_service_proto_init() }
func file_reservation_reservation_service_proto_init() {
	if File_reservation_reservation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservation_reservation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_reservation_reservation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reservation_reservation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservation_reservation_service_proto_goTypes,
		DependencyIndexes: file_reservation_reservation_service_proto_depIdxs,
		MessageInfos:      file_reservation_reservation_service_proto_msgTypes,
	}.Build()
	File_reservation_reservation_service_proto = out.File
	file_reservation_reservation_service_proto_rawDesc = nil
	file_reservation_reservation_service_proto_goTypes = nil
	file_reservation_reservation_service_proto_depIdxs = nil
}
