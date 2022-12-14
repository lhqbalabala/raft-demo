// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: Server.proto

package models

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

type RequestVoteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int32  `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	CandidateId  string `protobuf:"bytes,2,opt,name=CandidateId,proto3" json:"CandidateId,omitempty"`
	LastLogIndex int32  `protobuf:"varint,3,opt,name=LastLogIndex,proto3" json:"LastLogIndex,omitempty"`
	LastLogTerm  int32  `protobuf:"varint,4,opt,name=LastLogTerm,proto3" json:"LastLogTerm,omitempty"`
	Type         string `protobuf:"bytes,5,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *RequestVoteArgs) Reset() {
	*x = RequestVoteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteArgs) ProtoMessage() {}

func (x *RequestVoteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteArgs.ProtoReflect.Descriptor instead.
func (*RequestVoteArgs) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{0}
}

func (x *RequestVoteArgs) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteArgs) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

func (x *RequestVoteArgs) GetLastLogIndex() int32 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *RequestVoteArgs) GetLastLogTerm() int32 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

func (x *RequestVoteArgs) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type RequestVoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term      int32 `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	VoteReply bool  `protobuf:"varint,2,opt,name=VoteReply,proto3" json:"VoteReply,omitempty"`
}

func (x *RequestVoteReply) Reset() {
	*x = RequestVoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteReply) ProtoMessage() {}

func (x *RequestVoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteReply.ProtoReflect.Descriptor instead.
func (*RequestVoteReply) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{1}
}

func (x *RequestVoteReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteReply) GetVoteReply() bool {
	if x != nil {
		return x.VoteReply
	}
	return false
}

type AppendEntriesArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int32       `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`                 //??????????????????
	LeaderId     string      `protobuf:"bytes,2,opt,name=LeaderId,proto3" json:"LeaderId,omitempty"`          //????????? ID ??????????????????????????????????????????????????????????????????????????????????????? ID ??????????????????????????????????????????????????????????????????????????????????????????????????????????????????
	PrevLogIndex int32       `protobuf:"varint,3,opt,name=PrevLogIndex,proto3" json:"PrevLogIndex,omitempty"` //?????????????????????????????????????????????????????????
	PrevLogTerm  int32       `protobuf:"varint,4,opt,name=PrevLogTerm,proto3" json:"PrevLogTerm,omitempty"`   //?????????????????????????????????????????????????????????
	Entries      []*LogEntry `protobuf:"bytes,5,rep,name=Entries,proto3" json:"Entries,omitempty"`            //??????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????
	LeaderCommit int32       `protobuf:"varint,6,opt,name=LeaderCommit,proto3" json:"LeaderCommit,omitempty"` //????????????????????????????????????????????????????????????
}

func (x *AppendEntriesArgs) Reset() {
	*x = AppendEntriesArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesArgs) ProtoMessage() {}

func (x *AppendEntriesArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesArgs.ProtoReflect.Descriptor instead.
func (*AppendEntriesArgs) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{2}
}

func (x *AppendEntriesArgs) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesArgs) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *AppendEntriesArgs) GetPrevLogIndex() int32 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntriesArgs) GetPrevLogTerm() int32 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntriesArgs) GetEntries() []*LogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *AppendEntriesArgs) GetLeaderCommit() int32 {
	if x != nil {
		return x.LeaderCommit
	}
	return 0
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command *Command `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	Term    int32    `protobuf:"varint,2,opt,name=Term,proto3" json:"Term,omitempty"`
	Index   int32    `protobuf:"varint,3,opt,name=Index,proto3" json:"Index,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{3}
}

func (x *LogEntry) GetCommand() *Command {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *LogEntry) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *LogEntry) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Key          string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	Value        string `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
	Id           int32  `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`
	TimeDuration int32  `protobuf:"varint,5,opt,name=TimeDuration,proto3" json:"TimeDuration,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{4}
}

func (x *Command) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Command) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Command) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Command) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Command) GetTimeDuration() int32 {
	if x != nil {
		return x.TimeDuration
	}
	return 0
}

type AppendEntriesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    int32 `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`       //???????????????????????????????????? ???????????????????????????
	Success bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"` //???????????????????????????????????? prevLogIndex ?????? prevLogTerm ????????????????????? true
}

func (x *AppendEntriesReply) Reset() {
	*x = AppendEntriesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesReply) ProtoMessage() {}

func (x *AppendEntriesReply) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesReply.ProtoReflect.Descriptor instead.
func (*AppendEntriesReply) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{5}
}

func (x *AppendEntriesReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type AppendCommandArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command *Command `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
}

func (x *AppendCommandArgs) Reset() {
	*x = AppendCommandArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendCommandArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendCommandArgs) ProtoMessage() {}

func (x *AppendCommandArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendCommandArgs.ProtoReflect.Descriptor instead.
func (*AppendCommandArgs) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{6}
}

func (x *AppendCommandArgs) GetCommand() *Command {
	if x != nil {
		return x.Command
	}
	return nil
}

type AppendCommandReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsReceive bool `protobuf:"varint,1,opt,name=IsReceive,proto3" json:"IsReceive,omitempty"`
}

func (x *AppendCommandReply) Reset() {
	*x = AppendCommandReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendCommandReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendCommandReply) ProtoMessage() {}

func (x *AppendCommandReply) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendCommandReply.ProtoReflect.Descriptor instead.
func (*AppendCommandReply) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{7}
}

func (x *AppendCommandReply) GetIsReceive() bool {
	if x != nil {
		return x.IsReceive
	}
	return false
}

type GetKeyArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *GetKeyArgs) Reset() {
	*x = GetKeyArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyArgs) ProtoMessage() {}

func (x *GetKeyArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyArgs.ProtoReflect.Descriptor instead.
func (*GetKeyArgs) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{8}
}

func (x *GetKeyArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetKeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	IsOk  bool   `protobuf:"varint,2,opt,name=IsOk,proto3" json:"IsOk,omitempty"`
}

func (x *GetKeyReply) Reset() {
	*x = GetKeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyReply) ProtoMessage() {}

func (x *GetKeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyReply.ProtoReflect.Descriptor instead.
func (*GetKeyReply) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{9}
}

func (x *GetKeyReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetKeyReply) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

type SnapArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SnapArgs) Reset() {
	*x = SnapArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapArgs) ProtoMessage() {}

func (x *SnapArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapArgs.ProtoReflect.Descriptor instead.
func (*SnapArgs) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{10}
}

func (x *SnapArgs) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SnapReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAccept bool `protobuf:"varint,1,opt,name=IsAccept,proto3" json:"IsAccept,omitempty"`
}

func (x *SnapReply) Reset() {
	*x = SnapReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Server_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapReply) ProtoMessage() {}

func (x *SnapReply) ProtoReflect() protoreflect.Message {
	mi := &file_Server_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapReply.ProtoReflect.Descriptor instead.
func (*SnapReply) Descriptor() ([]byte, []int) {
	return file_Server_proto_rawDescGZIP(), []int{11}
}

func (x *SnapReply) GetIsAccept() bool {
	if x != nil {
		return x.IsAccept
	}
	return false
}

var File_Server_proto protoreflect.FileDescriptor

var file_Server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65,
	0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54,
	0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x44, 0x0a, 0x10, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x65,
	0x72, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0xd9, 0x01, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x50, 0x72,
	0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72,
	0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x2a, 0x0a, 0x07,
	0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x5f, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x29, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x79, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x54, 0x69, 0x6d, 0x65,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x65,
	0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3e, 0x0a, 0x11,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x72, 0x67,
	0x73, 0x12, 0x29, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x32, 0x0a, 0x12,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79,
	0x22, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x73, 0x4f, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x49, 0x73, 0x4f, 0x6b, 0x22, 0x1e, 0x0a, 0x08, 0x53, 0x6e, 0x61,
	0x70, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x09, 0x53, 0x6e, 0x61,
	0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x32, 0xbf, 0x02, 0x0a, 0x07, 0x52, 0x61, 0x66, 0x74, 0x52, 0x50, 0x43, 0x12, 0x40,
	0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x46, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x1a, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x41, 0x72, 0x67, 0x73, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x31, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x13,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x12,
	0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x41, 0x72, 0x67,
	0x73, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Server_proto_rawDescOnce sync.Once
	file_Server_proto_rawDescData = file_Server_proto_rawDesc
)

func file_Server_proto_rawDescGZIP() []byte {
	file_Server_proto_rawDescOnce.Do(func() {
		file_Server_proto_rawDescData = protoimpl.X.CompressGZIP(file_Server_proto_rawDescData)
	})
	return file_Server_proto_rawDescData
}

var file_Server_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_Server_proto_goTypes = []interface{}{
	(*RequestVoteArgs)(nil),    // 0: models.RequestVoteArgs
	(*RequestVoteReply)(nil),   // 1: models.RequestVoteReply
	(*AppendEntriesArgs)(nil),  // 2: models.AppendEntriesArgs
	(*LogEntry)(nil),           // 3: models.LogEntry
	(*Command)(nil),            // 4: models.Command
	(*AppendEntriesReply)(nil), // 5: models.AppendEntriesReply
	(*AppendCommandArgs)(nil),  // 6: models.AppendCommandArgs
	(*AppendCommandReply)(nil), // 7: models.AppendCommandReply
	(*GetKeyArgs)(nil),         // 8: models.GetKeyArgs
	(*GetKeyReply)(nil),        // 9: models.GetKeyReply
	(*SnapArgs)(nil),           // 10: models.SnapArgs
	(*SnapReply)(nil),          // 11: models.SnapReply
}
var file_Server_proto_depIdxs = []int32{
	3,  // 0: models.AppendEntriesArgs.Entries:type_name -> models.LogEntry
	4,  // 1: models.LogEntry.Command:type_name -> models.Command
	4,  // 2: models.AppendCommandArgs.Command:type_name -> models.Command
	0,  // 3: models.RaftRPC.RequestVote:input_type -> models.RequestVoteArgs
	2,  // 4: models.RaftRPC.AppendEntries:input_type -> models.AppendEntriesArgs
	6,  // 5: models.RaftRPC.AppendCommand:input_type -> models.AppendCommandArgs
	8,  // 6: models.RaftRPC.GetKey:input_type -> models.GetKeyArgs
	10, // 7: models.RaftRPC.SendSnap:input_type -> models.SnapArgs
	1,  // 8: models.RaftRPC.RequestVote:output_type -> models.RequestVoteReply
	5,  // 9: models.RaftRPC.AppendEntries:output_type -> models.AppendEntriesReply
	7,  // 10: models.RaftRPC.AppendCommand:output_type -> models.AppendCommandReply
	9,  // 11: models.RaftRPC.GetKey:output_type -> models.GetKeyReply
	11, // 12: models.RaftRPC.SendSnap:output_type -> models.SnapReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_Server_proto_init() }
func file_Server_proto_init() {
	if File_Server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteArgs); i {
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
		file_Server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteReply); i {
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
		file_Server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesArgs); i {
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
		file_Server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
		file_Server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_Server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesReply); i {
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
		file_Server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendCommandArgs); i {
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
		file_Server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendCommandReply); i {
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
		file_Server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyArgs); i {
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
		file_Server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyReply); i {
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
		file_Server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapArgs); i {
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
		file_Server_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapReply); i {
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
			RawDescriptor: file_Server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Server_proto_goTypes,
		DependencyIndexes: file_Server_proto_depIdxs,
		MessageInfos:      file_Server_proto_msgTypes,
	}.Build()
	File_Server_proto = out.File
	file_Server_proto_rawDesc = nil
	file_Server_proto_goTypes = nil
	file_Server_proto_depIdxs = nil
}
