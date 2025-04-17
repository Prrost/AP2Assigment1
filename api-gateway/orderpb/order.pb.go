// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: order.proto

package orderpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     int32                  `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Amount        int64                  `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateOrderRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateOrderRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Order         *Order                 `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int32                  `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type GetOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Order         *Order                 `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	mi := &file_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type UpdateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int32                  `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	ProductId     int32                  `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Amount        int64                  `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderRequest) Reset() {
	*x = UpdateOrderRequest{}
	mi := &file_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderRequest) ProtoMessage() {}

func (x *UpdateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateOrderRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *UpdateOrderRequest) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *UpdateOrderRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UpdateOrderRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Order         *Order                 `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderResponse) Reset() {
	*x = UpdateOrderResponse{}
	mi := &file_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderResponse) ProtoMessage() {}

func (x *UpdateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type ListAllOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllOrdersRequest) Reset() {
	*x = ListAllOrdersRequest{}
	mi := &file_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllOrdersRequest) ProtoMessage() {}

func (x *ListAllOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListAllOrdersRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

type ListAllOrdersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllOrdersResponse) Reset() {
	*x = ListAllOrdersResponse{}
	mi := &file_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllOrdersResponse) ProtoMessage() {}

func (x *ListAllOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListAllOrdersResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *ListAllOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int32                  `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	ProductId     int32                  `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	UserId        int32                  `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Amount        int64                  `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{8}
}

func (x *Order) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Order) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Order) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

const file_order_proto_rawDesc = "" +
	"\n" +
	"\vorder.proto\x12\x05order\"b\n" +
	"\x12CreateOrderRequest\x12\x1c\n" +
	"\tProductId\x18\x01 \x01(\x05R\tProductId\x12\x16\n" +
	"\x06UserId\x18\x02 \x01(\x05R\x06UserId\x12\x16\n" +
	"\x06Amount\x18\x03 \x01(\x03R\x06Amount\"S\n" +
	"\x13CreateOrderResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\"\n" +
	"\x05order\x18\x02 \x01(\v2\f.order.OrderR\x05order\"+\n" +
	"\x0fGetOrderRequest\x12\x18\n" +
	"\aOrderId\x18\x01 \x01(\x05R\aOrderId\"P\n" +
	"\x10GetOrderResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\"\n" +
	"\x05order\x18\x02 \x01(\v2\f.order.OrderR\x05order\"|\n" +
	"\x12UpdateOrderRequest\x12\x18\n" +
	"\aOrderId\x18\x01 \x01(\x05R\aOrderId\x12\x1c\n" +
	"\tProductId\x18\x02 \x01(\x05R\tProductId\x12\x16\n" +
	"\x06Amount\x18\x03 \x01(\x03R\x06Amount\x12\x16\n" +
	"\x06Status\x18\x04 \x01(\tR\x06Status\"S\n" +
	"\x13UpdateOrderResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\"\n" +
	"\x05order\x18\x02 \x01(\v2\f.order.OrderR\x05order\"\x16\n" +
	"\x14ListAllOrdersRequest\"=\n" +
	"\x15ListAllOrdersResponse\x12$\n" +
	"\x06orders\x18\x01 \x03(\v2\f.order.OrderR\x06orders\"\x87\x01\n" +
	"\x05Order\x12\x18\n" +
	"\aOrderId\x18\x01 \x01(\x05R\aOrderId\x12\x1c\n" +
	"\tProductId\x18\x02 \x01(\x05R\tProductId\x12\x16\n" +
	"\x06UserId\x18\x03 \x01(\x05R\x06UserId\x12\x16\n" +
	"\x06Amount\x18\x04 \x01(\x03R\x06Amount\x12\x16\n" +
	"\x06Status\x18\x05 \x01(\tR\x06Status2\xa3\x02\n" +
	"\fOrderService\x12D\n" +
	"\vCreateOrder\x12\x19.order.CreateOrderRequest\x1a\x1a.order.CreateOrderResponse\x12;\n" +
	"\bGetOrder\x12\x16.order.GetOrderRequest\x1a\x17.order.GetOrderResponse\x12D\n" +
	"\vUpdateOrder\x12\x19.order.UpdateOrderRequest\x1a\x1a.order.UpdateOrderResponse\x12J\n" +
	"\rListAllOrders\x12\x1b.order.ListAllOrdersRequest\x1a\x1c.order.ListAllOrdersResponseB\x14Z\x12../orderpb;orderpbb\x06proto3"

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData []byte
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)))
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_order_proto_goTypes = []any{
	(*CreateOrderRequest)(nil),    // 0: order.CreateOrderRequest
	(*CreateOrderResponse)(nil),   // 1: order.CreateOrderResponse
	(*GetOrderRequest)(nil),       // 2: order.GetOrderRequest
	(*GetOrderResponse)(nil),      // 3: order.GetOrderResponse
	(*UpdateOrderRequest)(nil),    // 4: order.UpdateOrderRequest
	(*UpdateOrderResponse)(nil),   // 5: order.UpdateOrderResponse
	(*ListAllOrdersRequest)(nil),  // 6: order.ListAllOrdersRequest
	(*ListAllOrdersResponse)(nil), // 7: order.ListAllOrdersResponse
	(*Order)(nil),                 // 8: order.Order
}
var file_order_proto_depIdxs = []int32{
	8, // 0: order.CreateOrderResponse.order:type_name -> order.Order
	8, // 1: order.GetOrderResponse.order:type_name -> order.Order
	8, // 2: order.UpdateOrderResponse.order:type_name -> order.Order
	8, // 3: order.ListAllOrdersResponse.orders:type_name -> order.Order
	0, // 4: order.OrderService.CreateOrder:input_type -> order.CreateOrderRequest
	2, // 5: order.OrderService.GetOrder:input_type -> order.GetOrderRequest
	4, // 6: order.OrderService.UpdateOrder:input_type -> order.UpdateOrderRequest
	6, // 7: order.OrderService.ListAllOrders:input_type -> order.ListAllOrdersRequest
	1, // 8: order.OrderService.CreateOrder:output_type -> order.CreateOrderResponse
	3, // 9: order.OrderService.GetOrder:output_type -> order.GetOrderResponse
	5, // 10: order.OrderService.UpdateOrder:output_type -> order.UpdateOrderResponse
	7, // 11: order.OrderService.ListAllOrders:output_type -> order.ListAllOrdersResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
