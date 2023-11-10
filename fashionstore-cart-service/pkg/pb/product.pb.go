// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: pkg/pb/proto/product.proto

package pb

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

type AddProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Quantity    uint32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Discount    float64 `protobuf:"fixed64,5,opt,name=discount,proto3" json:"discount,omitempty"`
	Category    uint32  `protobuf:"varint,6,opt,name=category,proto3" json:"category,omitempty"`
	Brand       uint32  `protobuf:"varint,7,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *AddProductRequest) Reset() {
	*x = AddProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_proto_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductRequest) ProtoMessage() {}

func (x *AddProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_proto_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductRequest.ProtoReflect.Descriptor instead.
func (*AddProductRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_proto_product_proto_rawDescGZIP(), []int{0}
}

func (x *AddProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddProductRequest) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *AddProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AddProductRequest) GetDiscount() float64 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *AddProductRequest) GetCategory() uint32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *AddProductRequest) GetBrand() uint32 {
	if x != nil {
		return x.Brand
	}
	return 0
}

type AddProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddProductResponse) Reset() {
	*x = AddProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_proto_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductResponse) ProtoMessage() {}

func (x *AddProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_proto_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductResponse.ProtoReflect.Descriptor instead.
func (*AddProductResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_proto_product_proto_rawDescGZIP(), []int{1}
}

func (x *AddProductResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ListProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListProductRequest) Reset() {
	*x = ListProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_proto_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductRequest) ProtoMessage() {}

func (x *ListProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_proto_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductRequest.ProtoReflect.Descriptor instead.
func (*ListProductRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_proto_product_proto_rawDescGZIP(), []int{2}
}

func (x *ListProductRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProductRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ProductDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price        float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Discount     float64 `protobuf:"fixed64,5,opt,name=discount,proto3" json:"discount,omitempty"`
	Category     string  `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	Brand        string  `protobuf:"bytes,7,opt,name=brand,proto3" json:"brand,omitempty"`
	SellingPrice float64 `protobuf:"fixed64,8,opt,name=sellingPrice,proto3" json:"sellingPrice,omitempty"`
	Status       string  `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ProductDetail) Reset() {
	*x = ProductDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_proto_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetail) ProtoMessage() {}

func (x *ProductDetail) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_proto_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDetail.ProtoReflect.Descriptor instead.
func (*ProductDetail) Descriptor() ([]byte, []int) {
	return file_pkg_pb_proto_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductDetail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductDetail) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductDetail) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductDetail) GetDiscount() float64 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *ProductDetail) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProductDetail) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *ProductDetail) GetSellingPrice() float64 {
	if x != nil {
		return x.SellingPrice
	}
	return 0
}

func (x *ProductDetail) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ListProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ProductDetail `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ListProductResponse) Reset() {
	*x = ListProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_proto_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductResponse) ProtoMessage() {}

func (x *ListProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_proto_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductResponse.ProtoReflect.Descriptor instead.
func (*ListProductResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_proto_product_proto_rawDescGZIP(), []int{4}
}

func (x *ListProductResponse) GetProducts() []*ProductDetail {
	if x != nil {
		return x.Products
	}
	return nil
}

type ProductDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []uint64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductDetailsRequest) Reset() {
	*x = ProductDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_proto_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetailsRequest) ProtoMessage() {}

func (x *ProductDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_proto_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDetailsRequest.ProtoReflect.Descriptor instead.
func (*ProductDetailsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_proto_product_proto_rawDescGZIP(), []int{5}
}

func (x *ProductDetailsRequest) GetId() []uint64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type ProductDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*FetchProductResponse `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ProductDetailsResponse) Reset() {
	*x = ProductDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_proto_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetailsResponse) ProtoMessage() {}

func (x *ProductDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_proto_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDetailsResponse.ProtoReflect.Descriptor instead.
func (*ProductDetailsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_proto_product_proto_rawDescGZIP(), []int{6}
}

func (x *ProductDetailsResponse) GetProducts() []*FetchProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

type FetchProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FetchProductRequest) Reset() {
	*x = FetchProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_proto_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchProductRequest) ProtoMessage() {}

func (x *FetchProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_proto_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchProductRequest.ProtoReflect.Descriptor instead.
func (*FetchProductRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_proto_product_proto_rawDescGZIP(), []int{7}
}

func (x *FetchProductRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FetchProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Quantity     uint32  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	SellingPrice float64 `protobuf:"fixed64,5,opt,name=sellingPrice,proto3" json:"sellingPrice,omitempty"`
}

func (x *FetchProductResponse) Reset() {
	*x = FetchProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_proto_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchProductResponse) ProtoMessage() {}

func (x *FetchProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_proto_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchProductResponse.ProtoReflect.Descriptor instead.
func (*FetchProductResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_proto_product_proto_rawDescGZIP(), []int{8}
}

func (x *FetchProductResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FetchProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FetchProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FetchProductResponse) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *FetchProductResponse) GetSellingPrice() float64 {
	if x != nil {
		return x.SellingPrice
	}
	return 0
}

var File_pkg_pb_proto_product_proto protoreflect.FileDescriptor

var file_pkg_pb_proto_product_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0xc9, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x2a, 0x0a, 0x12,
	0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf5, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x44, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x4e, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22,
	0x25, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0xa2, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_proto_product_proto_rawDescOnce sync.Once
	file_pkg_pb_proto_product_proto_rawDescData = file_pkg_pb_proto_product_proto_rawDesc
)

func file_pkg_pb_proto_product_proto_rawDescGZIP() []byte {
	file_pkg_pb_proto_product_proto_rawDescOnce.Do(func() {
		file_pkg_pb_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_proto_product_proto_rawDescData)
	})
	return file_pkg_pb_proto_product_proto_rawDescData
}

var file_pkg_pb_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_pb_proto_product_proto_goTypes = []interface{}{
	(*AddProductRequest)(nil),      // 0: pb.AddProductRequest
	(*AddProductResponse)(nil),     // 1: pb.AddProductResponse
	(*ListProductRequest)(nil),     // 2: pb.ListProductRequest
	(*ProductDetail)(nil),          // 3: pb.ProductDetail
	(*ListProductResponse)(nil),    // 4: pb.ListProductResponse
	(*ProductDetailsRequest)(nil),  // 5: pb.ProductDetailsRequest
	(*ProductDetailsResponse)(nil), // 6: pb.ProductDetailsResponse
	(*FetchProductRequest)(nil),    // 7: pb.FetchProductRequest
	(*FetchProductResponse)(nil),   // 8: pb.FetchProductResponse
}
var file_pkg_pb_proto_product_proto_depIdxs = []int32{
	3, // 0: pb.ListProductResponse.products:type_name -> pb.ProductDetail
	8, // 1: pb.ProductDetailsResponse.products:type_name -> pb.FetchProductResponse
	0, // 2: pb.ProductService.AddProduct:input_type -> pb.AddProductRequest
	2, // 3: pb.ProductService.ListProducts:input_type -> pb.ListProductRequest
	5, // 4: pb.ProductService.ProductDetails:input_type -> pb.ProductDetailsRequest
	7, // 5: pb.ProductService.FetchProduct:input_type -> pb.FetchProductRequest
	1, // 6: pb.ProductService.AddProduct:output_type -> pb.AddProductResponse
	4, // 7: pb.ProductService.ListProducts:output_type -> pb.ListProductResponse
	6, // 8: pb.ProductService.ProductDetails:output_type -> pb.ProductDetailsResponse
	8, // 9: pb.ProductService.FetchProduct:output_type -> pb.FetchProductResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_pb_proto_product_proto_init() }
func file_pkg_pb_proto_product_proto_init() {
	if File_pkg_pb_proto_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_proto_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductRequest); i {
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
		file_pkg_pb_proto_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductResponse); i {
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
		file_pkg_pb_proto_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductRequest); i {
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
		file_pkg_pb_proto_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDetail); i {
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
		file_pkg_pb_proto_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductResponse); i {
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
		file_pkg_pb_proto_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDetailsRequest); i {
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
		file_pkg_pb_proto_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDetailsResponse); i {
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
		file_pkg_pb_proto_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchProductRequest); i {
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
		file_pkg_pb_proto_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchProductResponse); i {
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
			RawDescriptor: file_pkg_pb_proto_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_proto_product_proto_goTypes,
		DependencyIndexes: file_pkg_pb_proto_product_proto_depIdxs,
		MessageInfos:      file_pkg_pb_proto_product_proto_msgTypes,
	}.Build()
	File_pkg_pb_proto_product_proto = out.File
	file_pkg_pb_proto_product_proto_rawDesc = nil
	file_pkg_pb_proto_product_proto_goTypes = nil
	file_pkg_pb_proto_product_proto_depIdxs = nil
}
