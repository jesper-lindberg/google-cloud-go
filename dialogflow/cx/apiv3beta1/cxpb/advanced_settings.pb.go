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
// source: google/cloud/dialogflow/cx/v3beta1/advanced_settings.proto

package cxpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Hierarchical advanced settings for agent/flow/page/fulfillment/parameter.
// Settings exposed at lower level overrides the settings exposed at higher
// level. Overriding occurs at the sub-setting level. For example, the
// playback_interruption_settings at fulfillment level only overrides the
// playback_interruption_settings at the agent level, leaving other settings
// at the agent level unchanged.
//
// DTMF settings does not override each other. DTMF settings set at different
// levels define DTMF detections running in parallel.
//
// Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
type AdvancedSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If present, incoming audio is exported by Dialogflow to the configured
	// Google Cloud Storage destination.
	// Exposed at the following levels:
	// - Agent level
	// - Flow level
	AudioExportGcsDestination *GcsDestination `protobuf:"bytes,2,opt,name=audio_export_gcs_destination,json=audioExportGcsDestination,proto3" json:"audio_export_gcs_destination,omitempty"`
	// Settings for speech to text detection.
	// Exposed at the following levels:
	// - Agent level
	// - Flow level
	// - Page level
	// - Parameter level
	SpeechSettings *AdvancedSettings_SpeechSettings `protobuf:"bytes,3,opt,name=speech_settings,json=speechSettings,proto3" json:"speech_settings,omitempty"`
	// Settings for DTMF.
	// Exposed at the following levels:
	// - Agent level
	// - Flow level
	// - Page level
	// - Parameter level.
	DtmfSettings *AdvancedSettings_DtmfSettings `protobuf:"bytes,5,opt,name=dtmf_settings,json=dtmfSettings,proto3" json:"dtmf_settings,omitempty"`
	// Settings for logging.
	// Settings for Dialogflow History, Contact Center messages, StackDriver logs,
	// and speech logging.
	// Exposed at the following levels:
	// - Agent level.
	LoggingSettings *AdvancedSettings_LoggingSettings `protobuf:"bytes,6,opt,name=logging_settings,json=loggingSettings,proto3" json:"logging_settings,omitempty"`
}

func (x *AdvancedSettings) Reset() {
	*x = AdvancedSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvancedSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedSettings) ProtoMessage() {}

func (x *AdvancedSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedSettings.ProtoReflect.Descriptor instead.
func (*AdvancedSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescGZIP(), []int{0}
}

func (x *AdvancedSettings) GetAudioExportGcsDestination() *GcsDestination {
	if x != nil {
		return x.AudioExportGcsDestination
	}
	return nil
}

func (x *AdvancedSettings) GetSpeechSettings() *AdvancedSettings_SpeechSettings {
	if x != nil {
		return x.SpeechSettings
	}
	return nil
}

func (x *AdvancedSettings) GetDtmfSettings() *AdvancedSettings_DtmfSettings {
	if x != nil {
		return x.DtmfSettings
	}
	return nil
}

func (x *AdvancedSettings) GetLoggingSettings() *AdvancedSettings_LoggingSettings {
	if x != nil {
		return x.LoggingSettings
	}
	return nil
}

// Define behaviors of speech to text detection.
type AdvancedSettings_SpeechSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sensitivity of the speech model that detects the end of speech.
	// Scale from 0 to 100.
	EndpointerSensitivity int32 `protobuf:"varint,1,opt,name=endpointer_sensitivity,json=endpointerSensitivity,proto3" json:"endpointer_sensitivity,omitempty"`
	// Timeout before detecting no speech.
	NoSpeechTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=no_speech_timeout,json=noSpeechTimeout,proto3" json:"no_speech_timeout,omitempty"`
	// Use timeout based endpointing, interpreting endpointer sensitivy as
	// seconds of timeout value.
	UseTimeoutBasedEndpointing bool `protobuf:"varint,3,opt,name=use_timeout_based_endpointing,json=useTimeoutBasedEndpointing,proto3" json:"use_timeout_based_endpointing,omitempty"`
	// Mapping from language to Speech-to-Text model. The mapped Speech-to-Text
	// model will be selected for requests from its corresponding language.
	// For more information, see
	// [Speech
	// models](https://cloud.google.com/dialogflow/cx/docs/concept/speech-models).
	Models map[string]string `protobuf:"bytes,5,rep,name=models,proto3" json:"models,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AdvancedSettings_SpeechSettings) Reset() {
	*x = AdvancedSettings_SpeechSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvancedSettings_SpeechSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedSettings_SpeechSettings) ProtoMessage() {}

func (x *AdvancedSettings_SpeechSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedSettings_SpeechSettings.ProtoReflect.Descriptor instead.
func (*AdvancedSettings_SpeechSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AdvancedSettings_SpeechSettings) GetEndpointerSensitivity() int32 {
	if x != nil {
		return x.EndpointerSensitivity
	}
	return 0
}

func (x *AdvancedSettings_SpeechSettings) GetNoSpeechTimeout() *durationpb.Duration {
	if x != nil {
		return x.NoSpeechTimeout
	}
	return nil
}

func (x *AdvancedSettings_SpeechSettings) GetUseTimeoutBasedEndpointing() bool {
	if x != nil {
		return x.UseTimeoutBasedEndpointing
	}
	return false
}

func (x *AdvancedSettings_SpeechSettings) GetModels() map[string]string {
	if x != nil {
		return x.Models
	}
	return nil
}

// Define behaviors for DTMF (dual tone multi frequency).
type AdvancedSettings_DtmfSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, incoming audio is processed for DTMF (dual tone multi frequency)
	// events. For example, if the caller presses a button on their telephone
	// keypad and DTMF processing is enabled, Dialogflow will detect the
	// event (e.g. a "3" was pressed) in the incoming audio and pass the event
	// to the bot to drive business logic (e.g. when 3 is pressed, return the
	// account balance).
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Max length of DTMF digits.
	MaxDigits int32 `protobuf:"varint,2,opt,name=max_digits,json=maxDigits,proto3" json:"max_digits,omitempty"`
	// The digit that terminates a DTMF digit sequence.
	FinishDigit string `protobuf:"bytes,3,opt,name=finish_digit,json=finishDigit,proto3" json:"finish_digit,omitempty"`
}

func (x *AdvancedSettings_DtmfSettings) Reset() {
	*x = AdvancedSettings_DtmfSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvancedSettings_DtmfSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedSettings_DtmfSettings) ProtoMessage() {}

func (x *AdvancedSettings_DtmfSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedSettings_DtmfSettings.ProtoReflect.Descriptor instead.
func (*AdvancedSettings_DtmfSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AdvancedSettings_DtmfSettings) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *AdvancedSettings_DtmfSettings) GetMaxDigits() int32 {
	if x != nil {
		return x.MaxDigits
	}
	return 0
}

func (x *AdvancedSettings_DtmfSettings) GetFinishDigit() string {
	if x != nil {
		return x.FinishDigit
	}
	return ""
}

// Define behaviors on logging.
type AdvancedSettings_LoggingSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, StackDriver logging is currently enabled.
	EnableStackdriverLogging bool `protobuf:"varint,2,opt,name=enable_stackdriver_logging,json=enableStackdriverLogging,proto3" json:"enable_stackdriver_logging,omitempty"`
	// If true, DF Interaction logging is currently enabled.
	EnableInteractionLogging bool `protobuf:"varint,3,opt,name=enable_interaction_logging,json=enableInteractionLogging,proto3" json:"enable_interaction_logging,omitempty"`
}

func (x *AdvancedSettings_LoggingSettings) Reset() {
	*x = AdvancedSettings_LoggingSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvancedSettings_LoggingSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedSettings_LoggingSettings) ProtoMessage() {}

func (x *AdvancedSettings_LoggingSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedSettings_LoggingSettings.ProtoReflect.Descriptor instead.
func (*AdvancedSettings_LoggingSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescGZIP(), []int{0, 2}
}

func (x *AdvancedSettings_LoggingSettings) GetEnableStackdriverLogging() bool {
	if x != nil {
		return x.EnableStackdriverLogging
	}
	return false
}

func (x *AdvancedSettings_LoggingSettings) GetEnableInteractionLogging() bool {
	if x != nil {
		return x.EnableInteractionLogging
	}
	return false
}

var File_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc2, 0x08, 0x0a, 0x10, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x73, 0x0a, 0x1c, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x67, 0x63, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x47, 0x63, 0x73, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x19,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x63, 0x73, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6c, 0x0a, 0x0f, 0x73, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e,
	0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x66, 0x0a, 0x0d, 0x64, 0x74, 0x6d, 0x66, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x44, 0x74, 0x6d, 0x66, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x0c, 0x64, 0x74, 0x6d, 0x66, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x6f, 0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x0f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x1a, 0xf5, 0x02, 0x0a, 0x0e, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x35, 0x0a, 0x16, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x15, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x11, 0x6e, 0x6f,
	0x5f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x6e, 0x6f, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x41, 0x0a, 0x1d, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x75, 0x73, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x42, 0x61, 0x73, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x67, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63,
	0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63,
	0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x6a, 0x0a, 0x0c, 0x44, 0x74, 0x6d, 0x66,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x44, 0x69, 0x67, 0x69, 0x74,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x44,
	0x69, 0x67, 0x69, 0x74, 0x1a, 0x8d, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x0a, 0x1a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x42, 0xcf, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42,
	0x15, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x78, 0x70, 0x62, 0x3b, 0x63, 0x78, 0x70, 0x62,
	0xf8, 0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x43, 0x78, 0x2e, 0x56, 0x33, 0x42, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x26,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x43, 0x58, 0x3a, 0x3a, 0x56,
	0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescData = file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDesc
)

func file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescData
}

var file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_goTypes = []interface{}{
	(*AdvancedSettings)(nil),                 // 0: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings
	(*AdvancedSettings_SpeechSettings)(nil),  // 1: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings
	(*AdvancedSettings_DtmfSettings)(nil),    // 2: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.DtmfSettings
	(*AdvancedSettings_LoggingSettings)(nil), // 3: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.LoggingSettings
	nil,                                      // 4: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings.ModelsEntry
	(*GcsDestination)(nil),                   // 5: google.cloud.dialogflow.cx.v3beta1.GcsDestination
	(*durationpb.Duration)(nil),              // 6: google.protobuf.Duration
}
var file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_depIdxs = []int32{
	5, // 0: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.audio_export_gcs_destination:type_name -> google.cloud.dialogflow.cx.v3beta1.GcsDestination
	1, // 1: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.speech_settings:type_name -> google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings
	2, // 2: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.dtmf_settings:type_name -> google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.DtmfSettings
	3, // 3: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.logging_settings:type_name -> google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.LoggingSettings
	6, // 4: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings.no_speech_timeout:type_name -> google.protobuf.Duration
	4, // 5: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings.models:type_name -> google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.SpeechSettings.ModelsEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_init() }
func file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_init() {
	if File_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto != nil {
		return
	}
	file_google_cloud_dialogflow_cx_v3beta1_gcs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvancedSettings); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvancedSettings_SpeechSettings); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvancedSettings_DtmfSettings); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvancedSettings_LoggingSettings); i {
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
			RawDescriptor: file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_depIdxs,
		MessageInfos:      file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto = out.File
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDesc = nil
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_goTypes = nil
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_depIdxs = nil
}
