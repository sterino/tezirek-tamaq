// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: order.proto

package orderpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Элемент заказа
type Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Item) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

// Запрос на создание
type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RestaurantId  string                 `protobuf:"bytes,2,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Items         []*Item                `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	TotalPrice    float64                `protobuf:"fixed64,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *CreateOrderRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateOrderRequest) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *CreateOrderRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Запрос на обновление
type UpdateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items         []*Item                `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	TotalPrice    float64                `protobuf:"fixed64,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderRequest) Reset() {
	*x = UpdateOrderRequest{}
	mi := &file_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderRequest) ProtoMessage() {}

func (x *UpdateOrderRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateOrderRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *UpdateOrderRequest) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *UpdateOrderRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Запросы по ID
type GetOrderByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderByIDRequest) Reset() {
	*x = GetOrderByIDRequest{}
	mi := &file_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByIDRequest) ProtoMessage() {}

func (x *GetOrderByIDRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetOrderByIDRequest.ProtoReflect.Descriptor instead.
func (*GetOrderByIDRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOrderRequest) Reset() {
	*x = DeleteOrderRequest{}
	mi := &file_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderRequest) ProtoMessage() {}

func (x *DeleteOrderRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteOrderRequest.ProtoReflect.Descriptor instead.
func (*DeleteOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListByUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListByUserRequest) Reset() {
	*x = ListByUserRequest{}
	mi := &file_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByUserRequest) ProtoMessage() {}

func (x *ListByUserRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListByUserRequest.ProtoReflect.Descriptor instead.
func (*ListByUserRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *ListByUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListByRestaurantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RestaurantId  string                 `protobuf:"bytes,1,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListByRestaurantRequest) Reset() {
	*x = ListByRestaurantRequest{}
	mi := &file_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListByRestaurantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByRestaurantRequest) ProtoMessage() {}

func (x *ListByRestaurantRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListByRestaurantRequest.ProtoReflect.Descriptor instead.
func (*ListByRestaurantRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *ListByRestaurantRequest) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

// Ответ с заказом
type OrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RestaurantId  string                 `protobuf:"bytes,3,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Items         []*Item                `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	TotalPrice    float64                `protobuf:"fixed64,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	mi := &file_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *OrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderResponse) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *OrderResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OrderResponse) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *OrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OrderResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Ответ со списком заказов
type ListOrdersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*OrderResponse       `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersResponse) Reset() {
	*x = ListOrdersResponse{}
	mi := &file_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResponse) ProtoMessage() {}

func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{8}
}

func (x *ListOrdersResponse) GetOrders() []*OrderResponse {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

const file_order_proto_rawDesc = "" +
	"\n" +
	"\vorder.proto\x12\aorderpb\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"L\n" +
	"\x04Item\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price\"\xb0\x01\n" +
	"\x12CreateOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12#\n" +
	"\rrestaurant_id\x18\x02 \x01(\tR\frestaurantId\x12#\n" +
	"\x05items\x18\x03 \x03(\v2\r.orderpb.ItemR\x05items\x12\x1f\n" +
	"\vtotal_price\x18\x04 \x01(\x01R\n" +
	"totalPrice\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\"\x82\x01\n" +
	"\x12UpdateOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12#\n" +
	"\x05items\x18\x02 \x03(\v2\r.orderpb.ItemR\x05items\x12\x1f\n" +
	"\vtotal_price\x18\x03 \x01(\x01R\n" +
	"totalPrice\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\"%\n" +
	"\x13GetOrderByIDRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"$\n" +
	"\x12DeleteOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\",\n" +
	"\x11ListByUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\">\n" +
	"\x17ListByRestaurantRequest\x12#\n" +
	"\rrestaurant_id\x18\x01 \x01(\tR\frestaurantId\"\xb1\x02\n" +
	"\rOrderResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12#\n" +
	"\rrestaurant_id\x18\x03 \x01(\tR\frestaurantId\x12#\n" +
	"\x05items\x18\x04 \x03(\v2\r.orderpb.ItemR\x05items\x12\x1f\n" +
	"\vtotal_price\x18\x05 \x01(\x01R\n" +
	"totalPrice\x12\x16\n" +
	"\x06status\x18\x06 \x01(\tR\x06status\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"D\n" +
	"\x12ListOrdersResponse\x12.\n" +
	"\x06orders\x18\x01 \x03(\v2\x16.orderpb.OrderResponseR\x06orders2\xe3\x03\n" +
	"\fOrderService\x12=\n" +
	"\x06Create\x12\x1b.orderpb.CreateOrderRequest\x1a\x16.orderpb.OrderResponse\x12?\n" +
	"\aGetByID\x12\x1c.orderpb.GetOrderByIDRequest\x1a\x16.orderpb.OrderResponse\x12=\n" +
	"\x06Update\x12\x1b.orderpb.UpdateOrderRequest\x1a\x16.google.protobuf.Empty\x12=\n" +
	"\x06Delete\x12\x1b.orderpb.DeleteOrderRequest\x1a\x16.google.protobuf.Empty\x12;\n" +
	"\x04List\x12\x16.google.protobuf.Empty\x1a\x1b.orderpb.ListOrdersResponse\x12E\n" +
	"\n" +
	"ListByUser\x12\x1a.orderpb.ListByUserRequest\x1a\x1b.orderpb.ListOrdersResponse\x12Q\n" +
	"\x10ListByRestaurant\x12 .orderpb.ListByRestaurantRequest\x1a\x1b.orderpb.ListOrdersResponseB\x1fZ\x1dapi-service/proto/gen/orderpbb\x06proto3"

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
	(*Item)(nil),                    // 0: orderpb.Item
	(*CreateOrderRequest)(nil),      // 1: orderpb.CreateOrderRequest
	(*UpdateOrderRequest)(nil),      // 2: orderpb.UpdateOrderRequest
	(*GetOrderByIDRequest)(nil),     // 3: orderpb.GetOrderByIDRequest
	(*DeleteOrderRequest)(nil),      // 4: orderpb.DeleteOrderRequest
	(*ListByUserRequest)(nil),       // 5: orderpb.ListByUserRequest
	(*ListByRestaurantRequest)(nil), // 6: orderpb.ListByRestaurantRequest
	(*OrderResponse)(nil),           // 7: orderpb.OrderResponse
	(*ListOrdersResponse)(nil),      // 8: orderpb.ListOrdersResponse
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),           // 10: google.protobuf.Empty
}
var file_order_proto_depIdxs = []int32{
	0,  // 0: orderpb.CreateOrderRequest.items:type_name -> orderpb.Item
	0,  // 1: orderpb.UpdateOrderRequest.items:type_name -> orderpb.Item
	0,  // 2: orderpb.OrderResponse.items:type_name -> orderpb.Item
	9,  // 3: orderpb.OrderResponse.created_at:type_name -> google.protobuf.Timestamp
	9,  // 4: orderpb.OrderResponse.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 5: orderpb.ListOrdersResponse.orders:type_name -> orderpb.OrderResponse
	1,  // 6: orderpb.OrderService.Create:input_type -> orderpb.CreateOrderRequest
	3,  // 7: orderpb.OrderService.GetByID:input_type -> orderpb.GetOrderByIDRequest
	2,  // 8: orderpb.OrderService.Update:input_type -> orderpb.UpdateOrderRequest
	4,  // 9: orderpb.OrderService.Delete:input_type -> orderpb.DeleteOrderRequest
	10, // 10: orderpb.OrderService.List:input_type -> google.protobuf.Empty
	5,  // 11: orderpb.OrderService.ListByUser:input_type -> orderpb.ListByUserRequest
	6,  // 12: orderpb.OrderService.ListByRestaurant:input_type -> orderpb.ListByRestaurantRequest
	7,  // 13: orderpb.OrderService.Create:output_type -> orderpb.OrderResponse
	7,  // 14: orderpb.OrderService.GetByID:output_type -> orderpb.OrderResponse
	10, // 15: orderpb.OrderService.Update:output_type -> google.protobuf.Empty
	10, // 16: orderpb.OrderService.Delete:output_type -> google.protobuf.Empty
	8,  // 17: orderpb.OrderService.List:output_type -> orderpb.ListOrdersResponse
	8,  // 18: orderpb.OrderService.ListByUser:output_type -> orderpb.ListOrdersResponse
	8,  // 19: orderpb.OrderService.ListByRestaurant:output_type -> orderpb.ListOrdersResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
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
