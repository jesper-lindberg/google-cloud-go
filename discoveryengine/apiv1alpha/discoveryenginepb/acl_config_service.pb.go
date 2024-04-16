// Copyright 2022 Google LLC
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
// source: google/cloud/discoveryengine/v1alpha/acl_config_service.proto

package discoveryenginepb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for GetAclConfigRequest method.
type GetAclConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of
	// [AclConfig][google.cloud.discoveryengine.v1alpha.AclConfig], such as
	// `projects/*/locations/*/aclConfig`.
	//
	// If the caller does not have permission to access the
	// [AclConfig][google.cloud.discoveryengine.v1alpha.AclConfig], regardless of
	// whether or not it exists, a PERMISSION_DENIED error is returned.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAclConfigRequest) Reset() {
	*x = GetAclConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAclConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAclConfigRequest) ProtoMessage() {}

func (x *GetAclConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAclConfigRequest.ProtoReflect.Descriptor instead.
func (*GetAclConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAclConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for UpdateAclConfig method.
type UpdateAclConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AclConfig *AclConfig `protobuf:"bytes,1,opt,name=acl_config,json=aclConfig,proto3" json:"acl_config,omitempty"`
}

func (x *UpdateAclConfigRequest) Reset() {
	*x = UpdateAclConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAclConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAclConfigRequest) ProtoMessage() {}

func (x *UpdateAclConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAclConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateAclConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateAclConfigRequest) GetAclConfig() *AclConfig {
	if x != nil {
		return x.AclConfig
	}
	return nil
}

var File_google_cloud_discoveryengine_v1alpha_acl_config_service_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x63, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61,
	0x63, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x63, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6d, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x0a, 0x61, 0x63, 0x6c, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x41, 0x63, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x09, 0x61, 0x63, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0xf8, 0x03, 0x0a, 0x10,
	0x41, 0x63, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xd1, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x63, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x4f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x49, 0x3a, 0x0a, 0x61, 0x63, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x3b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x7b, 0x61, 0x63, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x63, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x7d, 0x12, 0xbb, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x63, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x63, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x3f, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32,
	0x12, 0x30, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x63, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x7d, 0x1a, 0x52, 0xca, 0x41, 0x1e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xa1, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x15, 0x41, 0x63, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62,
	0xa2, 0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49,
	0x4e, 0x45, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0xea, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDescData = file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_goTypes = []interface{}{
	(*GetAclConfigRequest)(nil),    // 0: google.cloud.discoveryengine.v1alpha.GetAclConfigRequest
	(*UpdateAclConfigRequest)(nil), // 1: google.cloud.discoveryengine.v1alpha.UpdateAclConfigRequest
	(*AclConfig)(nil),              // 2: google.cloud.discoveryengine.v1alpha.AclConfig
}
var file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_depIdxs = []int32{
	2, // 0: google.cloud.discoveryengine.v1alpha.UpdateAclConfigRequest.acl_config:type_name -> google.cloud.discoveryengine.v1alpha.AclConfig
	1, // 1: google.cloud.discoveryengine.v1alpha.AclConfigService.UpdateAclConfig:input_type -> google.cloud.discoveryengine.v1alpha.UpdateAclConfigRequest
	0, // 2: google.cloud.discoveryengine.v1alpha.AclConfigService.GetAclConfig:input_type -> google.cloud.discoveryengine.v1alpha.GetAclConfigRequest
	2, // 3: google.cloud.discoveryengine.v1alpha.AclConfigService.UpdateAclConfig:output_type -> google.cloud.discoveryengine.v1alpha.AclConfig
	2, // 4: google.cloud.discoveryengine.v1alpha.AclConfigService.GetAclConfig:output_type -> google.cloud.discoveryengine.v1alpha.AclConfig
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_init() }
func file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_init() {
	if File_google_cloud_discoveryengine_v1alpha_acl_config_service_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1alpha_acl_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAclConfigRequest); i {
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
		file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAclConfigRequest); i {
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
			RawDescriptor: file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1alpha_acl_config_service_proto = out.File
	file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1alpha_acl_config_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AclConfigServiceClient is the client API for AclConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AclConfigServiceClient interface {
	// Default Acl Configuration for use in a location of a customer's project.
	// Updates will only reflect to new data stores. Existing data stores will
	// still use the old value.
	UpdateAclConfig(ctx context.Context, in *UpdateAclConfigRequest, opts ...grpc.CallOption) (*AclConfig, error)
	// Gets the [AclConfig][google.cloud.discoveryengine.v1alpha.AclConfig].
	GetAclConfig(ctx context.Context, in *GetAclConfigRequest, opts ...grpc.CallOption) (*AclConfig, error)
}

type aclConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAclConfigServiceClient(cc grpc.ClientConnInterface) AclConfigServiceClient {
	return &aclConfigServiceClient{cc}
}

func (c *aclConfigServiceClient) UpdateAclConfig(ctx context.Context, in *UpdateAclConfigRequest, opts ...grpc.CallOption) (*AclConfig, error) {
	out := new(AclConfig)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.AclConfigService/UpdateAclConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aclConfigServiceClient) GetAclConfig(ctx context.Context, in *GetAclConfigRequest, opts ...grpc.CallOption) (*AclConfig, error) {
	out := new(AclConfig)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.AclConfigService/GetAclConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AclConfigServiceServer is the server API for AclConfigService service.
type AclConfigServiceServer interface {
	// Default Acl Configuration for use in a location of a customer's project.
	// Updates will only reflect to new data stores. Existing data stores will
	// still use the old value.
	UpdateAclConfig(context.Context, *UpdateAclConfigRequest) (*AclConfig, error)
	// Gets the [AclConfig][google.cloud.discoveryengine.v1alpha.AclConfig].
	GetAclConfig(context.Context, *GetAclConfigRequest) (*AclConfig, error)
}

// UnimplementedAclConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAclConfigServiceServer struct {
}

func (*UnimplementedAclConfigServiceServer) UpdateAclConfig(context.Context, *UpdateAclConfigRequest) (*AclConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAclConfig not implemented")
}
func (*UnimplementedAclConfigServiceServer) GetAclConfig(context.Context, *GetAclConfigRequest) (*AclConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAclConfig not implemented")
}

func RegisterAclConfigServiceServer(s *grpc.Server, srv AclConfigServiceServer) {
	s.RegisterService(&_AclConfigService_serviceDesc, srv)
}

func _AclConfigService_UpdateAclConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAclConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclConfigServiceServer).UpdateAclConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.AclConfigService/UpdateAclConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclConfigServiceServer).UpdateAclConfig(ctx, req.(*UpdateAclConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AclConfigService_GetAclConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAclConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclConfigServiceServer).GetAclConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.AclConfigService/GetAclConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclConfigServiceServer).GetAclConfig(ctx, req.(*GetAclConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AclConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1alpha.AclConfigService",
	HandlerType: (*AclConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateAclConfig",
			Handler:    _AclConfigService_UpdateAclConfig_Handler,
		},
		{
			MethodName: "GetAclConfig",
			Handler:    _AclConfigService_GetAclConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1alpha/acl_config_service.proto",
}
