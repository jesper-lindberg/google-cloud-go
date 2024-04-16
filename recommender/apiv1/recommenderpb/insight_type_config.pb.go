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
// source: google/cloud/recommender/v1/insight_type_config.proto

package recommenderpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration for an InsightType.
type InsightTypeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of insight type config.
	// Eg,
	// projects/[PROJECT_NUMBER]/locations/[LOCATION]/insightTypes/[INSIGHT_TYPE_ID]/config
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// InsightTypeGenerationConfig which configures the generation of
	// insights for this insight type.
	InsightTypeGenerationConfig *InsightTypeGenerationConfig `protobuf:"bytes,2,opt,name=insight_type_generation_config,json=insightTypeGenerationConfig,proto3" json:"insight_type_generation_config,omitempty"`
	// Fingerprint of the InsightTypeConfig. Provides optimistic locking when
	// updating.
	Etag string `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	// Last time when the config was updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Immutable. The revision ID of the config.
	// A new revision is committed whenever the config is changed in any way.
	// The format is an 8-character hexadecimal string.
	RevisionId string `protobuf:"bytes,5,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Allows clients to store small amounts of arbitrary data. Annotations must
	// follow the Kubernetes syntax.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A user-settable field to provide a human-readable name to be used in user
	// interfaces.
	DisplayName string `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *InsightTypeConfig) Reset() {
	*x = InsightTypeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommender_v1_insight_type_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightTypeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightTypeConfig) ProtoMessage() {}

func (x *InsightTypeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommender_v1_insight_type_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightTypeConfig.ProtoReflect.Descriptor instead.
func (*InsightTypeConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommender_v1_insight_type_config_proto_rawDescGZIP(), []int{0}
}

func (x *InsightTypeConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InsightTypeConfig) GetInsightTypeGenerationConfig() *InsightTypeGenerationConfig {
	if x != nil {
		return x.InsightTypeGenerationConfig
	}
	return nil
}

func (x *InsightTypeConfig) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *InsightTypeConfig) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *InsightTypeConfig) GetRevisionId() string {
	if x != nil {
		return x.RevisionId
	}
	return ""
}

func (x *InsightTypeConfig) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *InsightTypeConfig) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// A configuration to customize the generation of insights.
// Eg, customizing the lookback period considered when generating a
// insight.
type InsightTypeGenerationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Parameters for this InsightTypeGenerationConfig. These configs can be used
	// by or are applied to all subtypes.
	Params *structpb.Struct `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *InsightTypeGenerationConfig) Reset() {
	*x = InsightTypeGenerationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommender_v1_insight_type_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightTypeGenerationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightTypeGenerationConfig) ProtoMessage() {}

func (x *InsightTypeGenerationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommender_v1_insight_type_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightTypeGenerationConfig.ProtoReflect.Descriptor instead.
func (*InsightTypeGenerationConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommender_v1_insight_type_config_proto_rawDescGZIP(), []int{1}
}

func (x *InsightTypeGenerationConfig) GetParams() *structpb.Struct {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_google_cloud_recommender_v1_insight_type_config_proto protoreflect.FileDescriptor

var file_google_cloud_recommender_v1_insight_type_config_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x98, 0x06, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x7d, 0x0a, 0x1e, 0x69, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x1b, 0x69, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0b, 0x72, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe0, 0x41, 0x05, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x61, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0xaf, 0x02, 0xea, 0x41, 0xab, 0x02, 0x0a,
	0x2c, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x54, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x59, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x2f, 0x7b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4e, 0x0a, 0x1b, 0x49, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0xa3, 0x01, 0x0a, 0x1f, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x16,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x70, 0x62, 0xa2, 0x02, 0x04, 0x43, 0x52,
	0x45, 0x43, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_recommender_v1_insight_type_config_proto_rawDescOnce sync.Once
	file_google_cloud_recommender_v1_insight_type_config_proto_rawDescData = file_google_cloud_recommender_v1_insight_type_config_proto_rawDesc
)

func file_google_cloud_recommender_v1_insight_type_config_proto_rawDescGZIP() []byte {
	file_google_cloud_recommender_v1_insight_type_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_recommender_v1_insight_type_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_recommender_v1_insight_type_config_proto_rawDescData)
	})
	return file_google_cloud_recommender_v1_insight_type_config_proto_rawDescData
}

var file_google_cloud_recommender_v1_insight_type_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_recommender_v1_insight_type_config_proto_goTypes = []interface{}{
	(*InsightTypeConfig)(nil),           // 0: google.cloud.recommender.v1.InsightTypeConfig
	(*InsightTypeGenerationConfig)(nil), // 1: google.cloud.recommender.v1.InsightTypeGenerationConfig
	nil,                                 // 2: google.cloud.recommender.v1.InsightTypeConfig.AnnotationsEntry
	(*timestamppb.Timestamp)(nil),       // 3: google.protobuf.Timestamp
	(*structpb.Struct)(nil),             // 4: google.protobuf.Struct
}
var file_google_cloud_recommender_v1_insight_type_config_proto_depIdxs = []int32{
	1, // 0: google.cloud.recommender.v1.InsightTypeConfig.insight_type_generation_config:type_name -> google.cloud.recommender.v1.InsightTypeGenerationConfig
	3, // 1: google.cloud.recommender.v1.InsightTypeConfig.update_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.cloud.recommender.v1.InsightTypeConfig.annotations:type_name -> google.cloud.recommender.v1.InsightTypeConfig.AnnotationsEntry
	4, // 3: google.cloud.recommender.v1.InsightTypeGenerationConfig.params:type_name -> google.protobuf.Struct
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_recommender_v1_insight_type_config_proto_init() }
func file_google_cloud_recommender_v1_insight_type_config_proto_init() {
	if File_google_cloud_recommender_v1_insight_type_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_recommender_v1_insight_type_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightTypeConfig); i {
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
		file_google_cloud_recommender_v1_insight_type_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightTypeGenerationConfig); i {
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
			RawDescriptor: file_google_cloud_recommender_v1_insight_type_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_recommender_v1_insight_type_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_recommender_v1_insight_type_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_recommender_v1_insight_type_config_proto_msgTypes,
	}.Build()
	File_google_cloud_recommender_v1_insight_type_config_proto = out.File
	file_google_cloud_recommender_v1_insight_type_config_proto_rawDesc = nil
	file_google_cloud_recommender_v1_insight_type_config_proto_goTypes = nil
	file_google_cloud_recommender_v1_insight_type_config_proto_depIdxs = nil
}