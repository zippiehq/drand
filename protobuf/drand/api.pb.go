//
// This protobuf file contains the definition of the public API endpoints as
// well as messages. All client implementations should use this reference
// protobuf to implement a compatible drand client.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: drand/api.proto

package drand

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// PublicRandRequest requests a public random value that has been generated in a
// unbiasable way and verifiable.
type PublicRandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// round uniquely identifies a beacon. If round == 0 (or unspecified), then
	// the response will contain the last.
	Round uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *PublicRandRequest) Reset() {
	*x = PublicRandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicRandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicRandRequest) ProtoMessage() {}

func (x *PublicRandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drand_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicRandRequest.ProtoReflect.Descriptor instead.
func (*PublicRandRequest) Descriptor() ([]byte, []int) {
	return file_drand_api_proto_rawDescGZIP(), []int{0}
}

func (x *PublicRandRequest) GetRound() uint64 {
	if x != nil {
		return x.Round
	}
	return 0
}

// PublicRandResponse holds a signature which is the random value. It can be
// verified thanks to the distributed public key of the nodes that have ran the
// DKG protocol and is unbiasable. The randomness can be verified using the BLS
// verification routine with the message "round || previous_rand".
type PublicRandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round             uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	Signature         []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	PreviousSignature []byte `protobuf:"bytes,3,opt,name=previous_signature,json=previousSignature,proto3" json:"previous_signature,omitempty"`
	// randomness is simply there to demonstrate - it is the hash of the
	// signature. It should be computed locally.
	Randomness []byte `protobuf:"bytes,4,opt,name=randomness,proto3" json:"randomness,omitempty"`
}

func (x *PublicRandResponse) Reset() {
	*x = PublicRandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicRandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicRandResponse) ProtoMessage() {}

func (x *PublicRandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drand_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicRandResponse.ProtoReflect.Descriptor instead.
func (*PublicRandResponse) Descriptor() ([]byte, []int) {
	return file_drand_api_proto_rawDescGZIP(), []int{1}
}

func (x *PublicRandResponse) GetRound() uint64 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *PublicRandResponse) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *PublicRandResponse) GetPreviousSignature() []byte {
	if x != nil {
		return x.PreviousSignature
	}
	return nil
}

func (x *PublicRandResponse) GetRandomness() []byte {
	if x != nil {
		return x.Randomness
	}
	return nil
}

// PrivateRandRequest is the message to send when requesting a private random
// value.
type PrivateRandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Request is the ECIES encryption of an ephemereal public key towards which
	// to encrypt the private randomness. The format of the bytes is denoted by
	// the ECIES encryption used by drand.
	Request []byte `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *PrivateRandRequest) Reset() {
	*x = PrivateRandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateRandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateRandRequest) ProtoMessage() {}

func (x *PrivateRandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drand_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateRandRequest.ProtoReflect.Descriptor instead.
func (*PrivateRandRequest) Descriptor() ([]byte, []int) {
	return file_drand_api_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateRandRequest) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

type PrivateRandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Responses is the ECIES encryption of the private randomness using the
	// ephemereal public key sent in the request.  The format of the bytes is
	// denoted by the ECIES  encryption used by drand.
	Response []byte `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PrivateRandResponse) Reset() {
	*x = PrivateRandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateRandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateRandResponse) ProtoMessage() {}

func (x *PrivateRandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drand_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateRandResponse.ProtoReflect.Descriptor instead.
func (*PrivateRandResponse) Descriptor() ([]byte, []int) {
	return file_drand_api_proto_rawDescGZIP(), []int{3}
}

func (x *PrivateRandResponse) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

type HomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HomeRequest) Reset() {
	*x = HomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeRequest) ProtoMessage() {}

func (x *HomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drand_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeRequest.ProtoReflect.Descriptor instead.
func (*HomeRequest) Descriptor() ([]byte, []int) {
	return file_drand_api_proto_rawDescGZIP(), []int{4}
}

type HomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HomeResponse) Reset() {
	*x = HomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeResponse) ProtoMessage() {}

func (x *HomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drand_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeResponse.ProtoReflect.Descriptor instead.
func (*HomeResponse) Descriptor() ([]byte, []int) {
	return file_drand_api_proto_rawDescGZIP(), []int{5}
}

func (x *HomeResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_drand_api_proto protoreflect.FileDescriptor

var file_drand_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x1a, 0x12, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x11,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x52, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65, 0x73,
	0x73, 0x22, 0x2e, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x31, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x0c, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xcb, 0x02, 0x0a, 0x06,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x41, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x52, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x52, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x52, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x2e,
	0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65,
	0x12, 0x12, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x48, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x64, 0x72,
	0x61, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x72, 0x61,
	0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_drand_api_proto_rawDescOnce sync.Once
	file_drand_api_proto_rawDescData = file_drand_api_proto_rawDesc
)

func file_drand_api_proto_rawDescGZIP() []byte {
	file_drand_api_proto_rawDescOnce.Do(func() {
		file_drand_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_drand_api_proto_rawDescData)
	})
	return file_drand_api_proto_rawDescData
}

var file_drand_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_drand_api_proto_goTypes = []interface{}{
	(*PublicRandRequest)(nil),   // 0: drand.PublicRandRequest
	(*PublicRandResponse)(nil),  // 1: drand.PublicRandResponse
	(*PrivateRandRequest)(nil),  // 2: drand.PrivateRandRequest
	(*PrivateRandResponse)(nil), // 3: drand.PrivateRandResponse
	(*HomeRequest)(nil),         // 4: drand.HomeRequest
	(*HomeResponse)(nil),        // 5: drand.HomeResponse
	(*ChainInfoRequest)(nil),    // 6: drand.ChainInfoRequest
	(*ChainInfoPacket)(nil),     // 7: drand.ChainInfoPacket
}
var file_drand_api_proto_depIdxs = []int32{
	0, // 0: drand.Public.PublicRand:input_type -> drand.PublicRandRequest
	0, // 1: drand.Public.PublicRandStream:input_type -> drand.PublicRandRequest
	2, // 2: drand.Public.PrivateRand:input_type -> drand.PrivateRandRequest
	6, // 3: drand.Public.ChainInfo:input_type -> drand.ChainInfoRequest
	4, // 4: drand.Public.Home:input_type -> drand.HomeRequest
	1, // 5: drand.Public.PublicRand:output_type -> drand.PublicRandResponse
	1, // 6: drand.Public.PublicRandStream:output_type -> drand.PublicRandResponse
	3, // 7: drand.Public.PrivateRand:output_type -> drand.PrivateRandResponse
	7, // 8: drand.Public.ChainInfo:output_type -> drand.ChainInfoPacket
	5, // 9: drand.Public.Home:output_type -> drand.HomeResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_drand_api_proto_init() }
func file_drand_api_proto_init() {
	if File_drand_api_proto != nil {
		return
	}
	file_drand_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_drand_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicRandRequest); i {
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
		file_drand_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicRandResponse); i {
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
		file_drand_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateRandRequest); i {
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
		file_drand_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateRandResponse); i {
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
		file_drand_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeRequest); i {
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
		file_drand_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeResponse); i {
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
			RawDescriptor: file_drand_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_drand_api_proto_goTypes,
		DependencyIndexes: file_drand_api_proto_depIdxs,
		MessageInfos:      file_drand_api_proto_msgTypes,
	}.Build()
	File_drand_api_proto = out.File
	file_drand_api_proto_rawDesc = nil
	file_drand_api_proto_goTypes = nil
	file_drand_api_proto_depIdxs = nil
}
