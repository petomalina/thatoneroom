// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: thatoneroom/server/v1/service.proto

package serverv1

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

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thatoneroom_server_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thatoneroom_server_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_thatoneroom_server_v1_service_proto_rawDescGZIP(), []int{0}
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thatoneroom_server_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thatoneroom_server_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_thatoneroom_server_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConnectResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*Request_Move
	Action isRequest_Action `protobuf_oneof:"action"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thatoneroom_server_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_thatoneroom_server_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_thatoneroom_server_v1_service_proto_rawDescGZIP(), []int{2}
}

func (m *Request) GetAction() isRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *Request) GetMove() *Move {
	if x, ok := x.GetAction().(*Request_Move); ok {
		return x.Move
	}
	return nil
}

type isRequest_Action interface {
	isRequest_Action()
}

type Request_Move struct {
	Move *Move `protobuf:"bytes,1,opt,name=move,proto3,oneof"`
}

func (*Request_Move) isRequest_Action() {}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*Response_UpdateEntity
	Action isResponse_Action `protobuf_oneof:"action"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thatoneroom_server_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_thatoneroom_server_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_thatoneroom_server_v1_service_proto_rawDescGZIP(), []int{3}
}

func (m *Response) GetAction() isResponse_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *Response) GetUpdateEntity() *UpdateEntity {
	if x, ok := x.GetAction().(*Response_UpdateEntity); ok {
		return x.UpdateEntity
	}
	return nil
}

type isResponse_Action interface {
	isResponse_Action()
}

type Response_UpdateEntity struct {
	UpdateEntity *UpdateEntity `protobuf:"bytes,2,opt,name=updateEntity,proto3,oneof"`
}

func (*Response_UpdateEntity) isResponse_Action() {}

type Move struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To *Coordinate `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *Move) Reset() {
	*x = Move{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thatoneroom_server_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Move) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Move) ProtoMessage() {}

func (x *Move) ProtoReflect() protoreflect.Message {
	mi := &file_thatoneroom_server_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Move.ProtoReflect.Descriptor instead.
func (*Move) Descriptor() ([]byte, []int) {
	return file_thatoneroom_server_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *Move) GetTo() *Coordinate {
	if x != nil {
		return x.To
	}
	return nil
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Entity:
	//
	//	*Entity_Player
	Entity isEntity_Entity `protobuf_oneof:"entity"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thatoneroom_server_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_thatoneroom_server_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_thatoneroom_server_v1_service_proto_rawDescGZIP(), []int{5}
}

func (m *Entity) GetEntity() isEntity_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (x *Entity) GetPlayer() *Player {
	if x, ok := x.GetEntity().(*Entity_Player); ok {
		return x.Player
	}
	return nil
}

type isEntity_Entity interface {
	isEntity_Entity()
}

type Entity_Player struct {
	Player *Player `protobuf:"bytes,1,opt,name=player,proto3,oneof"`
}

func (*Entity_Player) isEntity_Entity() {}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Position *Coordinate `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thatoneroom_server_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_thatoneroom_server_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_thatoneroom_server_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetPosition() *Coordinate {
	if x != nil {
		return x.Position
	}
	return nil
}

type Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Coordinate) Reset() {
	*x = Coordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thatoneroom_server_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinate) ProtoMessage() {}

func (x *Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_thatoneroom_server_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinate.ProtoReflect.Descriptor instead.
func (*Coordinate) Descriptor() ([]byte, []int) {
	return file_thatoneroom_server_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *Coordinate) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coordinate) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type UpdateEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *UpdateEntity) Reset() {
	*x = UpdateEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thatoneroom_server_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntity) ProtoMessage() {}

func (x *UpdateEntity) ProtoReflect() protoreflect.Message {
	mi := &file_thatoneroom_server_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntity.ProtoReflect.Descriptor instead.
func (*UpdateEntity) Descriptor() ([]byte, []int) {
	return file_thatoneroom_server_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateEntity) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

var File_thatoneroom_server_v1_service_proto protoreflect.FileDescriptor

var file_thatoneroom_server_v1_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x10, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37,
	0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x48, 0x00, 0x52,
	0x04, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x5f, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x39, 0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x31, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x06, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72,
	0x6f, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x57, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3d, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x28, 0x0a, 0x0a, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x45, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x68,
	0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x32, 0xbc, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12,
	0x25, 0x2e, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65,
	0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1e, 0x2e, 0x74, 0x68, 0x61,
	0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x68, 0x61,
	0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0xf0, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e,
	0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x74, 0x6f,
	0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x61, 0x2f, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f,
	0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x15, 0x54, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65,
	0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x15, 0x54, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x54, 0x68, 0x61, 0x74, 0x6f, 0x6e, 0x65,
	0x72, 0x6f, 0x6f, 0x6d, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x54, 0x68, 0x61,
	0x74, 0x6f, 0x6e, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thatoneroom_server_v1_service_proto_rawDescOnce sync.Once
	file_thatoneroom_server_v1_service_proto_rawDescData = file_thatoneroom_server_v1_service_proto_rawDesc
)

func file_thatoneroom_server_v1_service_proto_rawDescGZIP() []byte {
	file_thatoneroom_server_v1_service_proto_rawDescOnce.Do(func() {
		file_thatoneroom_server_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_thatoneroom_server_v1_service_proto_rawDescData)
	})
	return file_thatoneroom_server_v1_service_proto_rawDescData
}

var file_thatoneroom_server_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_thatoneroom_server_v1_service_proto_goTypes = []interface{}{
	(*ConnectRequest)(nil),  // 0: thatoneroom.server.v1.ConnectRequest
	(*ConnectResponse)(nil), // 1: thatoneroom.server.v1.ConnectResponse
	(*Request)(nil),         // 2: thatoneroom.server.v1.Request
	(*Response)(nil),        // 3: thatoneroom.server.v1.Response
	(*Move)(nil),            // 4: thatoneroom.server.v1.Move
	(*Entity)(nil),          // 5: thatoneroom.server.v1.Entity
	(*Player)(nil),          // 6: thatoneroom.server.v1.Player
	(*Coordinate)(nil),      // 7: thatoneroom.server.v1.Coordinate
	(*UpdateEntity)(nil),    // 8: thatoneroom.server.v1.UpdateEntity
}
var file_thatoneroom_server_v1_service_proto_depIdxs = []int32{
	4, // 0: thatoneroom.server.v1.Request.move:type_name -> thatoneroom.server.v1.Move
	8, // 1: thatoneroom.server.v1.Response.updateEntity:type_name -> thatoneroom.server.v1.UpdateEntity
	7, // 2: thatoneroom.server.v1.Move.to:type_name -> thatoneroom.server.v1.Coordinate
	6, // 3: thatoneroom.server.v1.Entity.player:type_name -> thatoneroom.server.v1.Player
	7, // 4: thatoneroom.server.v1.Player.position:type_name -> thatoneroom.server.v1.Coordinate
	5, // 5: thatoneroom.server.v1.UpdateEntity.entity:type_name -> thatoneroom.server.v1.Entity
	0, // 6: thatoneroom.server.v1.ServerService.Connect:input_type -> thatoneroom.server.v1.ConnectRequest
	2, // 7: thatoneroom.server.v1.ServerService.Stream:input_type -> thatoneroom.server.v1.Request
	1, // 8: thatoneroom.server.v1.ServerService.Connect:output_type -> thatoneroom.server.v1.ConnectResponse
	3, // 9: thatoneroom.server.v1.ServerService.Stream:output_type -> thatoneroom.server.v1.Response
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_thatoneroom_server_v1_service_proto_init() }
func file_thatoneroom_server_v1_service_proto_init() {
	if File_thatoneroom_server_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thatoneroom_server_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_thatoneroom_server_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_thatoneroom_server_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_thatoneroom_server_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_thatoneroom_server_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Move); i {
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
		file_thatoneroom_server_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_thatoneroom_server_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_thatoneroom_server_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinate); i {
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
		file_thatoneroom_server_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntity); i {
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
	file_thatoneroom_server_v1_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Request_Move)(nil),
	}
	file_thatoneroom_server_v1_service_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Response_UpdateEntity)(nil),
	}
	file_thatoneroom_server_v1_service_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Entity_Player)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_thatoneroom_server_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thatoneroom_server_v1_service_proto_goTypes,
		DependencyIndexes: file_thatoneroom_server_v1_service_proto_depIdxs,
		MessageInfos:      file_thatoneroom_server_v1_service_proto_msgTypes,
	}.Build()
	File_thatoneroom_server_v1_service_proto = out.File
	file_thatoneroom_server_v1_service_proto_rawDesc = nil
	file_thatoneroom_server_v1_service_proto_goTypes = nil
	file_thatoneroom_server_v1_service_proto_depIdxs = nil
}
