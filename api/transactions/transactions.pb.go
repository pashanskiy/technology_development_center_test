// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: transactions/transactions.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	core "technology_development_center_test/api/core"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionsInfo []*TransactionInfo `protobuf:"bytes,1,rep,name=TransactionsInfo,proto3" json:"TransactionsInfo,omitempty"`
}

func (x *TransactionsInfo) Reset() {
	*x = TransactionsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_transactions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsInfo) ProtoMessage() {}

func (x *TransactionsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_transactions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsInfo.ProtoReflect.Descriptor instead.
func (*TransactionsInfo) Descriptor() ([]byte, []int) {
	return file_transactions_transactions_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionsInfo) GetTransactionsInfo() []*TransactionInfo {
	if x != nil {
		return x.TransactionsInfo
	}
	return nil
}

type TransactionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         *string     `protobuf:"bytes,1,opt,name=uid,proto3,oneof" json:"uid,omitempty"`
	FromUserUid *string     `protobuf:"bytes,2,opt,name=from_user_uid,json=fromUserUid,proto3,oneof" json:"from_user_uid,omitempty"`
	ToUserUid   *string     `protobuf:"bytes,3,opt,name=to_user_uid,json=toUserUid,proto3,oneof" json:"to_user_uid,omitempty"`
	Money       *core.Money `protobuf:"bytes,4,opt,name=money,proto3,oneof" json:"money,omitempty"`
	Status      *string     `protobuf:"bytes,5,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Created     *int64      `protobuf:"varint,6,opt,name=created,proto3,oneof" json:"created,omitempty"`
}

func (x *TransactionInfo) Reset() {
	*x = TransactionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_transactions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionInfo) ProtoMessage() {}

func (x *TransactionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_transactions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionInfo.ProtoReflect.Descriptor instead.
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return file_transactions_transactions_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionInfo) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *TransactionInfo) GetFromUserUid() string {
	if x != nil && x.FromUserUid != nil {
		return *x.FromUserUid
	}
	return ""
}

func (x *TransactionInfo) GetToUserUid() string {
	if x != nil && x.ToUserUid != nil {
		return *x.ToUserUid
	}
	return ""
}

func (x *TransactionInfo) GetMoney() *core.Money {
	if x != nil {
		return x.Money
	}
	return nil
}

func (x *TransactionInfo) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *TransactionInfo) GetCreated() int64 {
	if x != nil && x.Created != nil {
		return *x.Created
	}
	return 0
}

type CreateDepositRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUid *core.Uid   `protobuf:"bytes,1,opt,name=user_uid,json=userUid,proto3" json:"user_uid,omitempty"`
	Money   *core.Money `protobuf:"bytes,2,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *CreateDepositRequest) Reset() {
	*x = CreateDepositRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_transactions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDepositRequest) ProtoMessage() {}

func (x *CreateDepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_transactions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDepositRequest.ProtoReflect.Descriptor instead.
func (*CreateDepositRequest) Descriptor() ([]byte, []int) {
	return file_transactions_transactions_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDepositRequest) GetUserUid() *core.Uid {
	if x != nil {
		return x.UserUid
	}
	return nil
}

func (x *CreateDepositRequest) GetMoney() *core.Money {
	if x != nil {
		return x.Money
	}
	return nil
}

type CreateWithdrawalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUserUid *core.Uid   `protobuf:"bytes,1,opt,name=from_user_uid,json=fromUserUid,proto3" json:"from_user_uid,omitempty"`
	ToUserUid   *core.Uid   `protobuf:"bytes,2,opt,name=to_user_uid,json=toUserUid,proto3" json:"to_user_uid,omitempty"`
	Money       *core.Money `protobuf:"bytes,3,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *CreateWithdrawalRequest) Reset() {
	*x = CreateWithdrawalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_transactions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWithdrawalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWithdrawalRequest) ProtoMessage() {}

func (x *CreateWithdrawalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_transactions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWithdrawalRequest.ProtoReflect.Descriptor instead.
func (*CreateWithdrawalRequest) Descriptor() ([]byte, []int) {
	return file_transactions_transactions_proto_rawDescGZIP(), []int{3}
}

func (x *CreateWithdrawalRequest) GetFromUserUid() *core.Uid {
	if x != nil {
		return x.FromUserUid
	}
	return nil
}

func (x *CreateWithdrawalRequest) GetToUserUid() *core.Uid {
	if x != nil {
		return x.ToUserUid
	}
	return nil
}

func (x *CreateWithdrawalRequest) GetMoney() *core.Money {
	if x != nil {
		return x.Money
	}
	return nil
}

var File_transactions_transactions_proto protoreflect.FileDescriptor

var file_transactions_transactions_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x51, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0xa5, 0x02, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x03, 0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x55, 0x69, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x74, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x48, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04,
	0x5f, 0x75, 0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x5f, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x69, 0x64, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0x96, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x55, 0x69, 0x64, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x55, 0x69, 0x64, 0x52, 0x09, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x32, 0xe8, 0x02, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x26,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x12, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x26,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x2a, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x69, 0x64, 0x12, 0x4c,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x61, 0x6c, 0x12, 0x2d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x09, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x69, 0x64, 0x42, 0x2e, 0x5a, 0x2c,
	0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transactions_transactions_proto_rawDescOnce sync.Once
	file_transactions_transactions_proto_rawDescData = file_transactions_transactions_proto_rawDesc
)

func file_transactions_transactions_proto_rawDescGZIP() []byte {
	file_transactions_transactions_proto_rawDescOnce.Do(func() {
		file_transactions_transactions_proto_rawDescData = protoimpl.X.CompressGZIP(file_transactions_transactions_proto_rawDescData)
	})
	return file_transactions_transactions_proto_rawDescData
}

var file_transactions_transactions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transactions_transactions_proto_goTypes = []interface{}{
	(*TransactionsInfo)(nil),        // 0: transactions_service.TransactionsInfo
	(*TransactionInfo)(nil),         // 1: transactions_service.TransactionInfo
	(*CreateDepositRequest)(nil),    // 2: transactions_service.CreateDepositRequest
	(*CreateWithdrawalRequest)(nil), // 3: transactions_service.CreateWithdrawalRequest
	(*core.Money)(nil),              // 4: core.Money
	(*core.Uid)(nil),                // 5: core.Uid
}
var file_transactions_transactions_proto_depIdxs = []int32{
	1,  // 0: transactions_service.TransactionsInfo.TransactionsInfo:type_name -> transactions_service.TransactionInfo
	4,  // 1: transactions_service.TransactionInfo.money:type_name -> core.Money
	5,  // 2: transactions_service.CreateDepositRequest.user_uid:type_name -> core.Uid
	4,  // 3: transactions_service.CreateDepositRequest.money:type_name -> core.Money
	5,  // 4: transactions_service.CreateWithdrawalRequest.from_user_uid:type_name -> core.Uid
	5,  // 5: transactions_service.CreateWithdrawalRequest.to_user_uid:type_name -> core.Uid
	4,  // 6: transactions_service.CreateWithdrawalRequest.money:type_name -> core.Money
	1,  // 7: transactions_service.TransactionsService.GetDeposit:input_type -> transactions_service.TransactionInfo
	1,  // 8: transactions_service.TransactionsService.GetWithdrawal:input_type -> transactions_service.TransactionInfo
	2,  // 9: transactions_service.TransactionsService.CreateDeposit:input_type -> transactions_service.CreateDepositRequest
	3,  // 10: transactions_service.TransactionsService.CreateWithdrawal:input_type -> transactions_service.CreateWithdrawalRequest
	0,  // 11: transactions_service.TransactionsService.GetDeposit:output_type -> transactions_service.TransactionsInfo
	0,  // 12: transactions_service.TransactionsService.GetWithdrawal:output_type -> transactions_service.TransactionsInfo
	5,  // 13: transactions_service.TransactionsService.CreateDeposit:output_type -> core.Uid
	5,  // 14: transactions_service.TransactionsService.CreateWithdrawal:output_type -> core.Uid
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_transactions_transactions_proto_init() }
func file_transactions_transactions_proto_init() {
	if File_transactions_transactions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transactions_transactions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsInfo); i {
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
		file_transactions_transactions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionInfo); i {
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
		file_transactions_transactions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDepositRequest); i {
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
		file_transactions_transactions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWithdrawalRequest); i {
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
	file_transactions_transactions_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transactions_transactions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transactions_transactions_proto_goTypes,
		DependencyIndexes: file_transactions_transactions_proto_depIdxs,
		MessageInfos:      file_transactions_transactions_proto_msgTypes,
	}.Build()
	File_transactions_transactions_proto = out.File
	file_transactions_transactions_proto_rawDesc = nil
	file_transactions_transactions_proto_goTypes = nil
	file_transactions_transactions_proto_depIdxs = nil
}
