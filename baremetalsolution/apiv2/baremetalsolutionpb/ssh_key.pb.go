// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/baremetalsolution/v2/ssh_key.proto

package baremetalsolutionpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An SSH key, used for authorizing with the interactive serial console feature.
type SSHKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The name of this SSH key.
	// Currently, the only valid value for the location is "global".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The public SSH key. This must be in OpenSSH .authorized_keys format.
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *SSHKey) Reset() {
	*x = SSHKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHKey) ProtoMessage() {}

func (x *SSHKey) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHKey.ProtoReflect.Descriptor instead.
func (*SSHKey) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescGZIP(), []int{0}
}

func (x *SSHKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SSHKey) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// Message for listing the public SSH keys in a project.
type ListSSHKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent containing the SSH keys.
	// Currently, the only valid value for the location is "global".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListSSHKeysRequest) Reset() {
	*x = ListSSHKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSSHKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSSHKeysRequest) ProtoMessage() {}

func (x *ListSSHKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSSHKeysRequest.ProtoReflect.Descriptor instead.
func (*ListSSHKeysRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescGZIP(), []int{1}
}

func (x *ListSSHKeysRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListSSHKeysRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSSHKeysRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Message for response of ListSSHKeys.
type ListSSHKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The SSH keys registered in the project.
	SshKeys []*SSHKey `protobuf:"bytes,1,rep,name=ssh_keys,json=sshKeys,proto3" json:"ssh_keys,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the list.
	NextPageToken string `protobuf:"bytes,90,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListSSHKeysResponse) Reset() {
	*x = ListSSHKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSSHKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSSHKeysResponse) ProtoMessage() {}

func (x *ListSSHKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSSHKeysResponse.ProtoReflect.Descriptor instead.
func (*ListSSHKeysResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescGZIP(), []int{2}
}

func (x *ListSSHKeysResponse) GetSshKeys() []*SSHKey {
	if x != nil {
		return x.SshKeys
	}
	return nil
}

func (x *ListSSHKeysResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Message for registering a public SSH key in a project.
type CreateSSHKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent containing the SSH keys.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The SSH key to register.
	SshKey *SSHKey `protobuf:"bytes,2,opt,name=ssh_key,json=sshKey,proto3" json:"ssh_key,omitempty"`
	// Required. The ID to use for the key, which will become the final component
	// of the key's resource name.
	//
	// This value must match the regex:
	//
	//	[a-zA-Z0-9@.\-_]{1,64}
	SshKeyId string `protobuf:"bytes,3,opt,name=ssh_key_id,json=sshKeyId,proto3" json:"ssh_key_id,omitempty"`
}

func (x *CreateSSHKeyRequest) Reset() {
	*x = CreateSSHKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSSHKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSSHKeyRequest) ProtoMessage() {}

func (x *CreateSSHKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSSHKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateSSHKeyRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSSHKeyRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateSSHKeyRequest) GetSshKey() *SSHKey {
	if x != nil {
		return x.SshKey
	}
	return nil
}

func (x *CreateSSHKeyRequest) GetSshKeyId() string {
	if x != nil {
		return x.SshKeyId
	}
	return ""
}

// Message for deleting an SSH key from a project.
type DeleteSSHKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the SSH key to delete.
	// Currently, the only valid value for the location is "global".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteSSHKeyRequest) Reset() {
	*x = DeleteSSHKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSSHKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSSHKeyRequest) ProtoMessage() {}

func (x *DeleteSSHKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSSHKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteSSHKeyRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSSHKeyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_baremetalsolution_v2_ssh_key_proto protoreflect.FileDescriptor

var file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x73, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa9, 0x01, 0x0a, 0x06, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x3a, 0x67, 0xea, 0x41, 0x64, 0x0a, 0x27, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x73, 0x68, 0x4b, 0x65,
	0x79, 0x12, 0x39, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x73, 0x68, 0x4b, 0x65,
	0x79, 0x73, 0x2f, 0x7b, 0x73, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0x22, 0x93, 0x01, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x83, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x53, 0x48, 0x4b, 0x65,
	0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x73, 0x73,
	0x68, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x73, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc4, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x07, 0x73, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x73, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0a,
	0x73, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x73, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x22,
	0x5a, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x62, 0x61,
	0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x73, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xfa, 0x01, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x0b, 0x53, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42, 0x61, 0x72,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56,
	0x32, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x42, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescOnce sync.Once
	file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescData = file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDesc
)

func file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescGZIP() []byte {
	file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescOnce.Do(func() {
		file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescData)
	})
	return file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDescData
}

var file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_baremetalsolution_v2_ssh_key_proto_goTypes = []interface{}{
	(*SSHKey)(nil),              // 0: google.cloud.baremetalsolution.v2.SSHKey
	(*ListSSHKeysRequest)(nil),  // 1: google.cloud.baremetalsolution.v2.ListSSHKeysRequest
	(*ListSSHKeysResponse)(nil), // 2: google.cloud.baremetalsolution.v2.ListSSHKeysResponse
	(*CreateSSHKeyRequest)(nil), // 3: google.cloud.baremetalsolution.v2.CreateSSHKeyRequest
	(*DeleteSSHKeyRequest)(nil), // 4: google.cloud.baremetalsolution.v2.DeleteSSHKeyRequest
}
var file_google_cloud_baremetalsolution_v2_ssh_key_proto_depIdxs = []int32{
	0, // 0: google.cloud.baremetalsolution.v2.ListSSHKeysResponse.ssh_keys:type_name -> google.cloud.baremetalsolution.v2.SSHKey
	0, // 1: google.cloud.baremetalsolution.v2.CreateSSHKeyRequest.ssh_key:type_name -> google.cloud.baremetalsolution.v2.SSHKey
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_baremetalsolution_v2_ssh_key_proto_init() }
func file_google_cloud_baremetalsolution_v2_ssh_key_proto_init() {
	if File_google_cloud_baremetalsolution_v2_ssh_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHKey); i {
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
		file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSSHKeysRequest); i {
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
		file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSSHKeysResponse); i {
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
		file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSSHKeyRequest); i {
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
		file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSSHKeyRequest); i {
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
			RawDescriptor: file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_baremetalsolution_v2_ssh_key_proto_goTypes,
		DependencyIndexes: file_google_cloud_baremetalsolution_v2_ssh_key_proto_depIdxs,
		MessageInfos:      file_google_cloud_baremetalsolution_v2_ssh_key_proto_msgTypes,
	}.Build()
	File_google_cloud_baremetalsolution_v2_ssh_key_proto = out.File
	file_google_cloud_baremetalsolution_v2_ssh_key_proto_rawDesc = nil
	file_google_cloud_baremetalsolution_v2_ssh_key_proto_goTypes = nil
	file_google_cloud_baremetalsolution_v2_ssh_key_proto_depIdxs = nil
}
