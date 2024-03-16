// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: logservice.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the log line
type StoreLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LogLine   string                 `protobuf:"bytes,2,opt,name=log_line,json=logLine,proto3" json:"log_line,omitempty"`
	Bucket    string                 `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Token     string                 `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *StoreLogRequest) Reset() {
	*x = StoreLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreLogRequest) ProtoMessage() {}

func (x *StoreLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreLogRequest.ProtoReflect.Descriptor instead.
func (*StoreLogRequest) Descriptor() ([]byte, []int) {
	return file_logservice_proto_rawDescGZIP(), []int{0}
}

func (x *StoreLogRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoreLogRequest) GetLogLine() string {
	if x != nil {
		return x.LogLine
	}
	return ""
}

func (x *StoreLogRequest) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *StoreLogRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StoreLogRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// The request message containing the batch of log lines
type StoreLogBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogLines []*StoreLogBatchRequest_LogLine `protobuf:"bytes,1,rep,name=log_lines,json=logLines,proto3" json:"log_lines,omitempty"`
	Bucket   string                          `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Token    string                          `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *StoreLogBatchRequest) Reset() {
	*x = StoreLogBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreLogBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreLogBatchRequest) ProtoMessage() {}

func (x *StoreLogBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreLogBatchRequest.ProtoReflect.Descriptor instead.
func (*StoreLogBatchRequest) Descriptor() ([]byte, []int) {
	return file_logservice_proto_rawDescGZIP(), []int{1}
}

func (x *StoreLogBatchRequest) GetLogLines() []*StoreLogBatchRequest_LogLine {
	if x != nil {
		return x.LogLines
	}
	return nil
}

func (x *StoreLogBatchRequest) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *StoreLogBatchRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// The request message for retrieving logs
type GetLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastX  int32  `protobuf:"varint,1,opt,name=last_x,json=lastX,proto3" json:"last_x,omitempty"`
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Token  string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetLogsRequest) Reset() {
	*x = GetLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsRequest) ProtoMessage() {}

func (x *GetLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsRequest.ProtoReflect.Descriptor instead.
func (*GetLogsRequest) Descriptor() ([]byte, []int) {
	return file_logservice_proto_rawDescGZIP(), []int{2}
}

func (x *GetLogsRequest) GetLastX() int32 {
	if x != nil {
		return x.LastX
	}
	return 0
}

func (x *GetLogsRequest) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *GetLogsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// The response message containing the logs
type GetLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*StoreLogRequest `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *GetLogsResponse) Reset() {
	*x = GetLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsResponse) ProtoMessage() {}

func (x *GetLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsResponse.ProtoReflect.Descriptor instead.
func (*GetLogsResponse) Descriptor() ([]byte, []int) {
	return file_logservice_proto_rawDescGZIP(), []int{3}
}

func (x *GetLogsResponse) GetLogs() []*StoreLogRequest {
	if x != nil {
		return x.Logs
	}
	return nil
}

// The request message for retrieving the number of logs
type GetLogCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetLogCountRequest) Reset() {
	*x = GetLogCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogCountRequest) ProtoMessage() {}

func (x *GetLogCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogCountRequest.ProtoReflect.Descriptor instead.
func (*GetLogCountRequest) Descriptor() ([]byte, []int) {
	return file_logservice_proto_rawDescGZIP(), []int{4}
}

func (x *GetLogCountRequest) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *GetLogCountRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// The response message containing the number of logs
type GetLogCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetLogCountResponse) Reset() {
	*x = GetLogCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogCountResponse) ProtoMessage() {}

func (x *GetLogCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogCountResponse.ProtoReflect.Descriptor instead.
func (*GetLogCountResponse) Descriptor() ([]byte, []int) {
	return file_logservice_proto_rawDescGZIP(), []int{5}
}

func (x *GetLogCountResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// The login request message
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_logservice_proto_rawDescGZIP(), []int{6}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// The login response message
type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_logservice_proto_rawDescGZIP(), []int{7}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type StoreLogBatchRequest_LogLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LogLine   string                 `protobuf:"bytes,2,opt,name=log_line,json=logLine,proto3" json:"log_line,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *StoreLogBatchRequest_LogLine) Reset() {
	*x = StoreLogBatchRequest_LogLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logservice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreLogBatchRequest_LogLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreLogBatchRequest_LogLine) ProtoMessage() {}

func (x *StoreLogBatchRequest_LogLine) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreLogBatchRequest_LogLine.ProtoReflect.Descriptor instead.
func (*StoreLogBatchRequest_LogLine) Descriptor() ([]byte, []int) {
	return file_logservice_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StoreLogBatchRequest_LogLine) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoreLogBatchRequest_LogLine) GetLogLine() string {
	if x != nil {
		return x.LogLine
	}
	return ""
}

func (x *StoreLogBatchRequest_LogLine) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_logservice_proto protoreflect.FileDescriptor

var file_logservice_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a,
	0x0f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0xfb, 0x01, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x67,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x09,
	0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x69,
	0x6e, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x1a, 0x6e, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x61, 0x73, 0x74, 0x58, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x42, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x2b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xea, 0x02, 0x0a,
	0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x0d, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x20, 0x2e, 0x6c, 0x6f,
	0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x6f,
	0x67, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73,
	0x12, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c,
	0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6c, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x61,
	0x73, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logservice_proto_rawDescOnce sync.Once
	file_logservice_proto_rawDescData = file_logservice_proto_rawDesc
)

func file_logservice_proto_rawDescGZIP() []byte {
	file_logservice_proto_rawDescOnce.Do(func() {
		file_logservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_logservice_proto_rawDescData)
	})
	return file_logservice_proto_rawDescData
}

var file_logservice_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_logservice_proto_goTypes = []interface{}{
	(*StoreLogRequest)(nil),              // 0: logservice.StoreLogRequest
	(*StoreLogBatchRequest)(nil),         // 1: logservice.StoreLogBatchRequest
	(*GetLogsRequest)(nil),               // 2: logservice.GetLogsRequest
	(*GetLogsResponse)(nil),              // 3: logservice.GetLogsResponse
	(*GetLogCountRequest)(nil),           // 4: logservice.GetLogCountRequest
	(*GetLogCountResponse)(nil),          // 5: logservice.GetLogCountResponse
	(*LoginRequest)(nil),                 // 6: logservice.LoginRequest
	(*LoginResponse)(nil),                // 7: logservice.LoginResponse
	(*StoreLogBatchRequest_LogLine)(nil), // 8: logservice.StoreLogBatchRequest.LogLine
	(*timestamppb.Timestamp)(nil),        // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                // 10: google.protobuf.Empty
}
var file_logservice_proto_depIdxs = []int32{
	9,  // 0: logservice.StoreLogRequest.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 1: logservice.StoreLogBatchRequest.log_lines:type_name -> logservice.StoreLogBatchRequest.LogLine
	0,  // 2: logservice.GetLogsResponse.logs:type_name -> logservice.StoreLogRequest
	9,  // 3: logservice.StoreLogBatchRequest.LogLine.timestamp:type_name -> google.protobuf.Timestamp
	6,  // 4: logservice.LogService.Login:input_type -> logservice.LoginRequest
	0,  // 5: logservice.LogService.StoreLog:input_type -> logservice.StoreLogRequest
	1,  // 6: logservice.LogService.StoreLogBatch:input_type -> logservice.StoreLogBatchRequest
	2,  // 7: logservice.LogService.GetLogs:input_type -> logservice.GetLogsRequest
	4,  // 8: logservice.LogService.GetLogCount:input_type -> logservice.GetLogCountRequest
	7,  // 9: logservice.LogService.Login:output_type -> logservice.LoginResponse
	10, // 10: logservice.LogService.StoreLog:output_type -> google.protobuf.Empty
	10, // 11: logservice.LogService.StoreLogBatch:output_type -> google.protobuf.Empty
	3,  // 12: logservice.LogService.GetLogs:output_type -> logservice.GetLogsResponse
	5,  // 13: logservice.LogService.GetLogCount:output_type -> logservice.GetLogCountResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_logservice_proto_init() }
func file_logservice_proto_init() {
	if File_logservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreLogRequest); i {
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
		file_logservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreLogBatchRequest); i {
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
		file_logservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsRequest); i {
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
		file_logservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsResponse); i {
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
		file_logservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogCountRequest); i {
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
		file_logservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogCountResponse); i {
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
		file_logservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_logservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_logservice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreLogBatchRequest_LogLine); i {
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
			RawDescriptor: file_logservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logservice_proto_goTypes,
		DependencyIndexes: file_logservice_proto_depIdxs,
		MessageInfos:      file_logservice_proto_msgTypes,
	}.Build()
	File_logservice_proto = out.File
	file_logservice_proto_rawDesc = nil
	file_logservice_proto_goTypes = nil
	file_logservice_proto_depIdxs = nil
}
