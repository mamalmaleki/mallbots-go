// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: depotpb/messages.proto

package depotpb

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

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	StoreId   string `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Quantity  int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_depotpb_messages_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderItem) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ShoppingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       string           `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Stops         map[string]*Stop `protobuf:"bytes,3,rep,name=stops,proto3" json:"stops,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AssignedBotId string           `protobuf:"bytes,4,opt,name=assigned_bot_id,json=assignedBotId,proto3" json:"assigned_bot_id,omitempty"`
	Status        string           `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ShoppingList) Reset() {
	*x = ShoppingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShoppingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShoppingList) ProtoMessage() {}

func (x *ShoppingList) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShoppingList.ProtoReflect.Descriptor instead.
func (*ShoppingList) Descriptor() ([]byte, []int) {
	return file_depotpb_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ShoppingList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShoppingList) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ShoppingList) GetStops() map[string]*Stop {
	if x != nil {
		return x.Stops
	}
	return nil
}

func (x *ShoppingList) GetAssignedBotId() string {
	if x != nil {
		return x.AssignedBotId
	}
	return ""
}

func (x *ShoppingList) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Stop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreName     string           `protobuf:"bytes,1,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	StoreLocation string           `protobuf:"bytes,2,opt,name=store_location,json=storeLocation,proto3" json:"store_location,omitempty"`
	Items         map[string]*Item `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Stop) Reset() {
	*x = Stop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop) ProtoMessage() {}

func (x *Stop) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop.ProtoReflect.Descriptor instead.
func (*Stop) Descriptor() ([]byte, []int) {
	return file_depotpb_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Stop) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *Stop) GetStoreLocation() string {
	if x != nil {
		return x.StoreLocation
	}
	return ""
}

func (x *Stop) GetItems() map[string]*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_depotpb_messages_proto_rawDescGZIP(), []int{3}
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

var File_depotpb_messages_proto protoreflect.FileDescriptor

var file_depotpb_messages_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70,
	0x62, 0x22, 0x61, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0xfa, 0x01, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x70, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x6f, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x47, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x70,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xc5, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x1a, 0x47, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x36, 0x0a, 0x04, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x42, 0x96, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70,
	0x62, 0x42, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x75, 0x73, 0x2f, 0x65, 0x64, 0x61, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x2d,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x68, 0x34, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x74,
	0x2f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62,
	0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62,
	0xca, 0x02, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0xe2, 0x02, 0x13, 0x44, 0x65, 0x70,
	0x6f, 0x74, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_depotpb_messages_proto_rawDescOnce sync.Once
	file_depotpb_messages_proto_rawDescData = file_depotpb_messages_proto_rawDesc
)

func file_depotpb_messages_proto_rawDescGZIP() []byte {
	file_depotpb_messages_proto_rawDescOnce.Do(func() {
		file_depotpb_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_depotpb_messages_proto_rawDescData)
	})
	return file_depotpb_messages_proto_rawDescData
}

var file_depotpb_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_depotpb_messages_proto_goTypes = []interface{}{
	(*OrderItem)(nil),    // 0: depotpb.OrderItem
	(*ShoppingList)(nil), // 1: depotpb.ShoppingList
	(*Stop)(nil),         // 2: depotpb.Stop
	(*Item)(nil),         // 3: depotpb.Item
	nil,                  // 4: depotpb.ShoppingList.StopsEntry
	nil,                  // 5: depotpb.Stop.ItemsEntry
}
var file_depotpb_messages_proto_depIdxs = []int32{
	4, // 0: depotpb.ShoppingList.stops:type_name -> depotpb.ShoppingList.StopsEntry
	5, // 1: depotpb.Stop.items:type_name -> depotpb.Stop.ItemsEntry
	2, // 2: depotpb.ShoppingList.StopsEntry.value:type_name -> depotpb.Stop
	3, // 3: depotpb.Stop.ItemsEntry.value:type_name -> depotpb.Item
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_depotpb_messages_proto_init() }
func file_depotpb_messages_proto_init() {
	if File_depotpb_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_depotpb_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_depotpb_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShoppingList); i {
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
		file_depotpb_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop); i {
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
		file_depotpb_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_depotpb_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_depotpb_messages_proto_goTypes,
		DependencyIndexes: file_depotpb_messages_proto_depIdxs,
		MessageInfos:      file_depotpb_messages_proto_msgTypes,
	}.Build()
	File_depotpb_messages_proto = out.File
	file_depotpb_messages_proto_rawDesc = nil
	file_depotpb_messages_proto_goTypes = nil
	file_depotpb_messages_proto_depIdxs = nil
}
