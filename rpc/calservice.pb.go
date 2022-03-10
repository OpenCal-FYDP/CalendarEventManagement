// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: rpc/calservice.proto

package rpc

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

// implemented to google’s spec https://developers.google.com/calendar/api/guides/create-events
type CalEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary    string   `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	Location   string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Start      int64    `protobuf:"varint,3,opt,name=Start,proto3" json:"Start,omitempty"` //iso timestamp
	End        int64    `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`     //iso timestamp
	Recurrence []string `protobuf:"bytes,5,rep,name=recurrence,proto3" json:"recurrence,omitempty"`
	Attendees  []string `protobuf:"bytes,6,rep,name=attendees,proto3" json:"attendees,omitempty"`
}

func (x *CalEvent) Reset() {
	*x = CalEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_calservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalEvent) ProtoMessage() {}

func (x *CalEvent) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_calservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalEvent.ProtoReflect.Descriptor instead.
func (*CalEvent) Descriptor() ([]byte, []int) {
	return file_rpc_calservice_proto_rawDescGZIP(), []int{0}
}

func (x *CalEvent) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *CalEvent) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CalEvent) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *CalEvent) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *CalEvent) GetRecurrence() []string {
	if x != nil {
		return x.Recurrence
	}
	return nil
}

func (x *CalEvent) GetAttendees() []string {
	if x != nil {
		return x.Attendees
	}
	return nil
}

type CreateEventReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CalendarId string `protobuf:"bytes,1,opt,name=calendarId,proto3" json:"calendarId,omitempty"`
	EventId    string `protobuf:"bytes,2,opt,name=eventId,proto3" json:"eventId,omitempty"`
	//  string oauthToken = 3;
	Event        *CalEvent `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	OwnerOfEvent string    `protobuf:"bytes,5,opt,name=ownerOfEvent,proto3" json:"ownerOfEvent,omitempty"` // email used to fetch oauth token from identity service
}

func (x *CreateEventReq) Reset() {
	*x = CreateEventReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_calservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventReq) ProtoMessage() {}

func (x *CreateEventReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_calservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventReq.ProtoReflect.Descriptor instead.
func (*CreateEventReq) Descriptor() ([]byte, []int) {
	return file_rpc_calservice_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventReq) GetCalendarId() string {
	if x != nil {
		return x.CalendarId
	}
	return ""
}

func (x *CreateEventReq) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *CreateEventReq) GetEvent() *CalEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *CreateEventReq) GetOwnerOfEvent() string {
	if x != nil {
		return x.OwnerOfEvent
	}
	return ""
}

type CreateEventRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CalendarId string    `protobuf:"bytes,1,opt,name=calendarId,proto3" json:"calendarId,omitempty"`
	EventId    string    `protobuf:"bytes,2,opt,name=eventId,proto3" json:"eventId,omitempty"`
	Event      *CalEvent `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *CreateEventRes) Reset() {
	*x = CreateEventRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_calservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRes) ProtoMessage() {}

func (x *CreateEventRes) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_calservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRes.ProtoReflect.Descriptor instead.
func (*CreateEventRes) Descriptor() ([]byte, []int) {
	return file_rpc_calservice_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEventRes) GetCalendarId() string {
	if x != nil {
		return x.CalendarId
	}
	return ""
}

func (x *CreateEventRes) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *CreateEventRes) GetEvent() *CalEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type UpdateEventReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CalendarId string `protobuf:"bytes,1,opt,name=calendarId,proto3" json:"calendarId,omitempty"`
	EventId    string `protobuf:"bytes,2,opt,name=eventId,proto3" json:"eventId,omitempty"`
	//  string oauthToken = 3;
	Event        *CalEvent `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	OwnerOfEvent string    `protobuf:"bytes,5,opt,name=ownerOfEvent,proto3" json:"ownerOfEvent,omitempty"` // email used to fetch oauth token from identity service
}

func (x *UpdateEventReq) Reset() {
	*x = UpdateEventReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_calservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventReq) ProtoMessage() {}

func (x *UpdateEventReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_calservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventReq.ProtoReflect.Descriptor instead.
func (*UpdateEventReq) Descriptor() ([]byte, []int) {
	return file_rpc_calservice_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEventReq) GetCalendarId() string {
	if x != nil {
		return x.CalendarId
	}
	return ""
}

func (x *UpdateEventReq) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *UpdateEventReq) GetEvent() *CalEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *UpdateEventReq) GetOwnerOfEvent() string {
	if x != nil {
		return x.OwnerOfEvent
	}
	return ""
}

type UpdateEventRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CalendarId string    `protobuf:"bytes,1,opt,name=calendarId,proto3" json:"calendarId,omitempty"`
	EventId    string    `protobuf:"bytes,2,opt,name=eventId,proto3" json:"eventId,omitempty"`
	Event      *CalEvent `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *UpdateEventRes) Reset() {
	*x = UpdateEventRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_calservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventRes) ProtoMessage() {}

func (x *UpdateEventRes) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_calservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventRes.ProtoReflect.Descriptor instead.
func (*UpdateEventRes) Descriptor() ([]byte, []int) {
	return file_rpc_calservice_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEventRes) GetCalendarId() string {
	if x != nil {
		return x.CalendarId
	}
	return ""
}

func (x *UpdateEventRes) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *UpdateEventRes) GetEvent() *CalEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type DeleteEventReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CalendarId string `protobuf:"bytes,1,opt,name=calendarId,proto3" json:"calendarId,omitempty"`
	EventId    string `protobuf:"bytes,2,opt,name=eventId,proto3" json:"eventId,omitempty"`
	//  string oauthToken = 3;
	OwnerOfEvent string `protobuf:"bytes,5,opt,name=ownerOfEvent,proto3" json:"ownerOfEvent,omitempty"` // email used to fetch oauth token from identity service
}

func (x *DeleteEventReq) Reset() {
	*x = DeleteEventReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_calservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventReq) ProtoMessage() {}

func (x *DeleteEventReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_calservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventReq.ProtoReflect.Descriptor instead.
func (*DeleteEventReq) Descriptor() ([]byte, []int) {
	return file_rpc_calservice_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteEventReq) GetCalendarId() string {
	if x != nil {
		return x.CalendarId
	}
	return ""
}

func (x *DeleteEventReq) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *DeleteEventReq) GetOwnerOfEvent() string {
	if x != nil {
		return x.OwnerOfEvent
	}
	return ""
}

type DeleteEventRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEventRes) Reset() {
	*x = DeleteEventRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_calservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventRes) ProtoMessage() {}

func (x *DeleteEventRes) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_calservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventRes.ProtoReflect.Descriptor instead.
func (*DeleteEventRes) Descriptor() ([]byte, []int) {
	return file_rpc_calservice_proto_rawDescGZIP(), []int{6}
}

type GetEventReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=eventId,proto3" json:"eventId,omitempty"`
}

func (x *GetEventReq) Reset() {
	*x = GetEventReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_calservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventReq) ProtoMessage() {}

func (x *GetEventReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_calservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventReq.ProtoReflect.Descriptor instead.
func (*GetEventReq) Descriptor() ([]byte, []int) {
	return file_rpc_calservice_proto_rawDescGZIP(), []int{7}
}

func (x *GetEventReq) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type GetEventRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *CalEvent `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *GetEventRes) Reset() {
	*x = GetEventRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_calservice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRes) ProtoMessage() {}

func (x *GetEventRes) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_calservice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRes.ProtoReflect.Descriptor instead.
func (*GetEventRes) Descriptor() ([]byte, []int) {
	return file_rpc_calservice_proto_rawDescGZIP(), []int{8}
}

func (x *GetEventRes) GetEvent() *CalEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_rpc_calservice_proto protoreflect.FileDescriptor

var file_rpc_calservice_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x65, 0x73, 0x22,
	0x8f, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x61,
	0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x6b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43,
	0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x8f,
	0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x6c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x6b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x61,
	0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x6e, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x10, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22,
	0x27, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0xdb, 0x01, 0x0a, 0x1e, 0x43, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2f, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0c, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_calservice_proto_rawDescOnce sync.Once
	file_rpc_calservice_proto_rawDescData = file_rpc_calservice_proto_rawDesc
)

func file_rpc_calservice_proto_rawDescGZIP() []byte {
	file_rpc_calservice_proto_rawDescOnce.Do(func() {
		file_rpc_calservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_calservice_proto_rawDescData)
	})
	return file_rpc_calservice_proto_rawDescData
}

var file_rpc_calservice_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rpc_calservice_proto_goTypes = []interface{}{
	(*CalEvent)(nil),       // 0: CalEvent
	(*CreateEventReq)(nil), // 1: CreateEventReq
	(*CreateEventRes)(nil), // 2: CreateEventRes
	(*UpdateEventReq)(nil), // 3: UpdateEventReq
	(*UpdateEventRes)(nil), // 4: UpdateEventRes
	(*DeleteEventReq)(nil), // 5: DeleteEventReq
	(*DeleteEventRes)(nil), // 6: DeleteEventRes
	(*GetEventReq)(nil),    // 7: GetEventReq
	(*GetEventRes)(nil),    // 8: GetEventRes
}
var file_rpc_calservice_proto_depIdxs = []int32{
	0, // 0: CreateEventReq.event:type_name -> CalEvent
	0, // 1: CreateEventRes.event:type_name -> CalEvent
	0, // 2: UpdateEventReq.event:type_name -> CalEvent
	0, // 3: UpdateEventRes.event:type_name -> CalEvent
	0, // 4: GetEventRes.event:type_name -> CalEvent
	1, // 5: CalendarEventManagementService.CreateEvent:input_type -> CreateEventReq
	3, // 6: CalendarEventManagementService.UpdateEvent:input_type -> UpdateEventReq
	5, // 7: CalendarEventManagementService.DeleteEvent:input_type -> DeleteEventReq
	7, // 8: CalendarEventManagementService.GetEvent:input_type -> GetEventReq
	2, // 9: CalendarEventManagementService.CreateEvent:output_type -> CreateEventRes
	4, // 10: CalendarEventManagementService.UpdateEvent:output_type -> UpdateEventRes
	6, // 11: CalendarEventManagementService.DeleteEvent:output_type -> DeleteEventRes
	8, // 12: CalendarEventManagementService.GetEvent:output_type -> GetEventRes
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_rpc_calservice_proto_init() }
func file_rpc_calservice_proto_init() {
	if File_rpc_calservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_calservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalEvent); i {
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
		file_rpc_calservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventReq); i {
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
		file_rpc_calservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRes); i {
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
		file_rpc_calservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventReq); i {
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
		file_rpc_calservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventRes); i {
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
		file_rpc_calservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventReq); i {
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
		file_rpc_calservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventRes); i {
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
		file_rpc_calservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventReq); i {
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
		file_rpc_calservice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRes); i {
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
			RawDescriptor: file_rpc_calservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_calservice_proto_goTypes,
		DependencyIndexes: file_rpc_calservice_proto_depIdxs,
		MessageInfos:      file_rpc_calservice_proto_msgTypes,
	}.Build()
	File_rpc_calservice_proto = out.File
	file_rpc_calservice_proto_rawDesc = nil
	file_rpc_calservice_proto_goTypes = nil
	file_rpc_calservice_proto_depIdxs = nil
}
