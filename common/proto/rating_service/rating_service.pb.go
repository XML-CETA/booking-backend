// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: rating_service/rating_service.proto

package rating_service

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

type RatingAccommodationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accommodation string `protobuf:"bytes,1,opt,name=accommodation,proto3" json:"accommodation,omitempty"`
	Rate          int32  `protobuf:"varint,2,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *RatingAccommodationRequest) Reset() {
	*x = RatingAccommodationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_rating_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingAccommodationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingAccommodationRequest) ProtoMessage() {}

func (x *RatingAccommodationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_rating_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingAccommodationRequest.ProtoReflect.Descriptor instead.
func (*RatingAccommodationRequest) Descriptor() ([]byte, []int) {
	return file_rating_service_rating_service_proto_rawDescGZIP(), []int{0}
}

func (x *RatingAccommodationRequest) GetAccommodation() string {
	if x != nil {
		return x.Accommodation
	}
	return ""
}

func (x *RatingAccommodationRequest) GetRate() int32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type RateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_rating_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_rating_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_rating_service_rating_service_proto_rawDescGZIP(), []int{1}
}

func (x *RateResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type RateAccommodationIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RateAccommodationIdRequest) Reset() {
	*x = RateAccommodationIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_rating_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateAccommodationIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateAccommodationIdRequest) ProtoMessage() {}

func (x *RateAccommodationIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_rating_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateAccommodationIdRequest.ProtoReflect.Descriptor instead.
func (*RateAccommodationIdRequest) Descriptor() ([]byte, []int) {
	return file_rating_service_rating_service_proto_rawDescGZIP(), []int{2}
}

func (x *RateAccommodationIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AverageRateAccommodationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float32 `protobuf:"fixed32,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *AverageRateAccommodationResponse) Reset() {
	*x = AverageRateAccommodationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_rating_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRateAccommodationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRateAccommodationResponse) ProtoMessage() {}

func (x *AverageRateAccommodationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_rating_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRateAccommodationResponse.ProtoReflect.Descriptor instead.
func (*AverageRateAccommodationResponse) Descriptor() ([]byte, []int) {
	return file_rating_service_rating_service_proto_rawDescGZIP(), []int{3}
}

func (x *AverageRateAccommodationResponse) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

type RateAccommodationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Date string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Rate int32  `protobuf:"varint,3,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *RateAccommodationResponse) Reset() {
	*x = RateAccommodationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_rating_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateAccommodationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateAccommodationResponse) ProtoMessage() {}

func (x *RateAccommodationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_rating_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateAccommodationResponse.ProtoReflect.Descriptor instead.
func (*RateAccommodationResponse) Descriptor() ([]byte, []int) {
	return file_rating_service_rating_service_proto_rawDescGZIP(), []int{4}
}

func (x *RateAccommodationResponse) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RateAccommodationResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *RateAccommodationResponse) GetRate() int32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type AllAccommodationRatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rates []*RateAccommodationResponse `protobuf:"bytes,1,rep,name=rates,proto3" json:"rates,omitempty"`
}

func (x *AllAccommodationRatesResponse) Reset() {
	*x = AllAccommodationRatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_service_rating_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllAccommodationRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllAccommodationRatesResponse) ProtoMessage() {}

func (x *AllAccommodationRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_service_rating_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllAccommodationRatesResponse.ProtoReflect.Descriptor instead.
func (*AllAccommodationRatesResponse) Descriptor() ([]byte, []int) {
	return file_rating_service_rating_service_proto_rawDescGZIP(), []int{5}
}

func (x *AllAccommodationRatesResponse) GetRates() []*RateAccommodationResponse {
	if x != nil {
		return x.Rates
	}
	return nil
}

var File_rating_service_rating_service_proto protoreflect.FileDescriptor

var file_rating_service_rating_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x1a, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x2c, 0x0a, 0x1a, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a,
	0x20, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x19, 0x52,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x72, 0x61, 0x74, 0x65, 0x22, 0x51, 0x0a, 0x1d, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x73, 0x32, 0xda, 0x04, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x68, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1b,
	0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x1a, 0x16, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6a, 0x0a,
	0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7c, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x8a, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12,
	0x23, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2d,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rating_service_rating_service_proto_rawDescOnce sync.Once
	file_rating_service_rating_service_proto_rawDescData = file_rating_service_rating_service_proto_rawDesc
)

func file_rating_service_rating_service_proto_rawDescGZIP() []byte {
	file_rating_service_rating_service_proto_rawDescOnce.Do(func() {
		file_rating_service_rating_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rating_service_rating_service_proto_rawDescData)
	})
	return file_rating_service_rating_service_proto_rawDescData
}

var file_rating_service_rating_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rating_service_rating_service_proto_goTypes = []interface{}{
	(*RatingAccommodationRequest)(nil),       // 0: RatingAccommodationRequest
	(*RateResponse)(nil),                     // 1: RateResponse
	(*RateAccommodationIdRequest)(nil),       // 2: RateAccommodationIdRequest
	(*AverageRateAccommodationResponse)(nil), // 3: AverageRateAccommodationResponse
	(*RateAccommodationResponse)(nil),        // 4: RateAccommodationResponse
	(*AllAccommodationRatesResponse)(nil),    // 5: AllAccommodationRatesResponse
}
var file_rating_service_rating_service_proto_depIdxs = []int32{
	4, // 0: AllAccommodationRatesResponse.rates:type_name -> RateAccommodationResponse
	0, // 1: RatingService.CreateAccommodationRate:input_type -> RatingAccommodationRequest
	0, // 2: RatingService.UpdateAccommodationRate:input_type -> RatingAccommodationRequest
	2, // 3: RatingService.DeleteAccommodationRate:input_type -> RateAccommodationIdRequest
	2, // 4: RatingService.GetAllAccommodationRates:input_type -> RateAccommodationIdRequest
	2, // 5: RatingService.GetAverageAccommodationRate:input_type -> RateAccommodationIdRequest
	1, // 6: RatingService.CreateAccommodationRate:output_type -> RateResponse
	1, // 7: RatingService.UpdateAccommodationRate:output_type -> RateResponse
	1, // 8: RatingService.DeleteAccommodationRate:output_type -> RateResponse
	5, // 9: RatingService.GetAllAccommodationRates:output_type -> AllAccommodationRatesResponse
	3, // 10: RatingService.GetAverageAccommodationRate:output_type -> AverageRateAccommodationResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rating_service_rating_service_proto_init() }
func file_rating_service_rating_service_proto_init() {
	if File_rating_service_rating_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rating_service_rating_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingAccommodationRequest); i {
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
		file_rating_service_rating_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateResponse); i {
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
		file_rating_service_rating_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateAccommodationIdRequest); i {
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
		file_rating_service_rating_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRateAccommodationResponse); i {
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
		file_rating_service_rating_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateAccommodationResponse); i {
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
		file_rating_service_rating_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllAccommodationRatesResponse); i {
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
			RawDescriptor: file_rating_service_rating_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rating_service_rating_service_proto_goTypes,
		DependencyIndexes: file_rating_service_rating_service_proto_depIdxs,
		MessageInfos:      file_rating_service_rating_service_proto_msgTypes,
	}.Build()
	File_rating_service_rating_service_proto = out.File
	file_rating_service_rating_service_proto_rawDesc = nil
	file_rating_service_rating_service_proto_goTypes = nil
	file_rating_service_rating_service_proto_depIdxs = nil
}
