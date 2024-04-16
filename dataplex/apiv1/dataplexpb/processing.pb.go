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
// source: google/cloud/dataplex/v1/processing.proto

package dataplexpb

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

// DataScan scheduling and trigger settings.
type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DataScan scheduling and trigger settings.
	//
	// If not specified, the default is `onDemand`.
	//
	// Types that are assignable to Mode:
	//
	//	*Trigger_OnDemand_
	//	*Trigger_Schedule_
	Mode isTrigger_Mode `protobuf_oneof:"mode"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger.ProtoReflect.Descriptor instead.
func (*Trigger) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_processing_proto_rawDescGZIP(), []int{0}
}

func (m *Trigger) GetMode() isTrigger_Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (x *Trigger) GetOnDemand() *Trigger_OnDemand {
	if x, ok := x.GetMode().(*Trigger_OnDemand_); ok {
		return x.OnDemand
	}
	return nil
}

func (x *Trigger) GetSchedule() *Trigger_Schedule {
	if x, ok := x.GetMode().(*Trigger_Schedule_); ok {
		return x.Schedule
	}
	return nil
}

type isTrigger_Mode interface {
	isTrigger_Mode()
}

type Trigger_OnDemand_ struct {
	// The scan runs once via `RunDataScan` API.
	OnDemand *Trigger_OnDemand `protobuf:"bytes,100,opt,name=on_demand,json=onDemand,proto3,oneof"`
}

type Trigger_Schedule_ struct {
	// The scan is scheduled to run periodically.
	Schedule *Trigger_Schedule `protobuf:"bytes,101,opt,name=schedule,proto3,oneof"`
}

func (*Trigger_OnDemand_) isTrigger_Mode() {}

func (*Trigger_Schedule_) isTrigger_Mode() {}

// The data source for DataScan.
type DataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source is required and immutable. Once it is set, it cannot be change
	// to others.
	//
	// Types that are assignable to Source:
	//
	//	*DataSource_Entity
	//	*DataSource_Resource
	Source isDataSource_Source `protobuf_oneof:"source"`
}

func (x *DataSource) Reset() {
	*x = DataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource) ProtoMessage() {}

func (x *DataSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource.ProtoReflect.Descriptor instead.
func (*DataSource) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_processing_proto_rawDescGZIP(), []int{1}
}

func (m *DataSource) GetSource() isDataSource_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *DataSource) GetEntity() string {
	if x, ok := x.GetSource().(*DataSource_Entity); ok {
		return x.Entity
	}
	return ""
}

func (x *DataSource) GetResource() string {
	if x, ok := x.GetSource().(*DataSource_Resource); ok {
		return x.Resource
	}
	return ""
}

type isDataSource_Source interface {
	isDataSource_Source()
}

type DataSource_Entity struct {
	// Immutable. The Dataplex entity that represents the data source (e.g.
	// BigQuery table) for DataScan, of the form:
	// `projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/entities/{entity_id}`.
	Entity string `protobuf:"bytes,100,opt,name=entity,proto3,oneof"`
}

type DataSource_Resource struct {
	// Immutable. The service-qualified full resource name of the cloud resource
	// for a DataScan job to scan against. The field could be: BigQuery table of
	// type "TABLE" for DataProfileScan/DataQualityScan Format:
	// //bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID
	Resource string `protobuf:"bytes,101,opt,name=resource,proto3,oneof"`
}

func (*DataSource_Entity) isDataSource_Source() {}

func (*DataSource_Resource) isDataSource_Source() {}

// The data scanned during processing (e.g. in incremental DataScan)
type ScannedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The range of scanned data
	//
	// Types that are assignable to DataRange:
	//
	//	*ScannedData_IncrementalField_
	DataRange isScannedData_DataRange `protobuf_oneof:"data_range"`
}

func (x *ScannedData) Reset() {
	*x = ScannedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScannedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScannedData) ProtoMessage() {}

func (x *ScannedData) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScannedData.ProtoReflect.Descriptor instead.
func (*ScannedData) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_processing_proto_rawDescGZIP(), []int{2}
}

func (m *ScannedData) GetDataRange() isScannedData_DataRange {
	if m != nil {
		return m.DataRange
	}
	return nil
}

func (x *ScannedData) GetIncrementalField() *ScannedData_IncrementalField {
	if x, ok := x.GetDataRange().(*ScannedData_IncrementalField_); ok {
		return x.IncrementalField
	}
	return nil
}

type isScannedData_DataRange interface {
	isScannedData_DataRange()
}

type ScannedData_IncrementalField_ struct {
	// The range denoted by values of an incremental field
	IncrementalField *ScannedData_IncrementalField `protobuf:"bytes,1,opt,name=incremental_field,json=incrementalField,proto3,oneof"`
}

func (*ScannedData_IncrementalField_) isScannedData_DataRange() {}

// The scan runs once via `RunDataScan` API.
type Trigger_OnDemand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Trigger_OnDemand) Reset() {
	*x = Trigger_OnDemand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger_OnDemand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger_OnDemand) ProtoMessage() {}

func (x *Trigger_OnDemand) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger_OnDemand.ProtoReflect.Descriptor instead.
func (*Trigger_OnDemand) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_processing_proto_rawDescGZIP(), []int{0, 0}
}

// The scan is scheduled to run periodically.
type Trigger_Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. [Cron](https://en.wikipedia.org/wiki/Cron) schedule for running
	// scans periodically.
	//
	// To explicitly set a timezone in the cron tab, apply a prefix in the
	// cron tab: **"CRON_TZ=${IANA_TIME_ZONE}"** or **"TZ=${IANA_TIME_ZONE}"**.
	// The **${IANA_TIME_ZONE}** may only be a valid string from IANA time zone
	// database
	// ([wikipedia](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)).
	// For example, `CRON_TZ=America/New_York 1 * * * *`, or
	// `TZ=America/New_York 1 * * * *`.
	//
	// This field is required for Schedule scans.
	Cron string `protobuf:"bytes,1,opt,name=cron,proto3" json:"cron,omitempty"`
}

func (x *Trigger_Schedule) Reset() {
	*x = Trigger_Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger_Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger_Schedule) ProtoMessage() {}

func (x *Trigger_Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger_Schedule.ProtoReflect.Descriptor instead.
func (*Trigger_Schedule) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_processing_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Trigger_Schedule) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

// A data range denoted by a pair of start/end values of a field.
type ScannedData_IncrementalField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field that contains values which monotonically increases over time
	// (e.g. a timestamp column).
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// Value that marks the start of the range.
	Start string `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	// Value that marks the end of the range.
	End string `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ScannedData_IncrementalField) Reset() {
	*x = ScannedData_IncrementalField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScannedData_IncrementalField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScannedData_IncrementalField) ProtoMessage() {}

func (x *ScannedData_IncrementalField) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_processing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScannedData_IncrementalField.ProtoReflect.Descriptor instead.
func (*ScannedData_IncrementalField) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_processing_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ScannedData_IncrementalField) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *ScannedData_IncrementalField) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *ScannedData_IncrementalField) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

var File_google_cloud_dataplex_v1_processing_proto protoreflect.FileDescriptor

var file_google_cloud_dataplex_v1_processing_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x49, 0x0a,
	0x09, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x08,
	0x6f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x48, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x1a, 0x0a, 0x0a, 0x08, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x23,
	0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x72,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x63,
	0x72, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x7b, 0x0a, 0x0a, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xe0, 0x41, 0x05, 0xfa, 0x41,
	0x20, 0x0a, 0x1e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x48, 0x00, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x05, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x0b, 0x53, 0x63, 0x61,
	0x6e, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x65, 0x0a, 0x11, 0x69, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x00, 0x52, 0x10, 0x69,
	0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a,
	0x50, 0x0a, 0x10, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x42,
	0x6b, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x42,
	0x0f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x70,
	0x62, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dataplex_v1_processing_proto_rawDescOnce sync.Once
	file_google_cloud_dataplex_v1_processing_proto_rawDescData = file_google_cloud_dataplex_v1_processing_proto_rawDesc
)

func file_google_cloud_dataplex_v1_processing_proto_rawDescGZIP() []byte {
	file_google_cloud_dataplex_v1_processing_proto_rawDescOnce.Do(func() {
		file_google_cloud_dataplex_v1_processing_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dataplex_v1_processing_proto_rawDescData)
	})
	return file_google_cloud_dataplex_v1_processing_proto_rawDescData
}

var file_google_cloud_dataplex_v1_processing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_dataplex_v1_processing_proto_goTypes = []interface{}{
	(*Trigger)(nil),                      // 0: google.cloud.dataplex.v1.Trigger
	(*DataSource)(nil),                   // 1: google.cloud.dataplex.v1.DataSource
	(*ScannedData)(nil),                  // 2: google.cloud.dataplex.v1.ScannedData
	(*Trigger_OnDemand)(nil),             // 3: google.cloud.dataplex.v1.Trigger.OnDemand
	(*Trigger_Schedule)(nil),             // 4: google.cloud.dataplex.v1.Trigger.Schedule
	(*ScannedData_IncrementalField)(nil), // 5: google.cloud.dataplex.v1.ScannedData.IncrementalField
}
var file_google_cloud_dataplex_v1_processing_proto_depIdxs = []int32{
	3, // 0: google.cloud.dataplex.v1.Trigger.on_demand:type_name -> google.cloud.dataplex.v1.Trigger.OnDemand
	4, // 1: google.cloud.dataplex.v1.Trigger.schedule:type_name -> google.cloud.dataplex.v1.Trigger.Schedule
	5, // 2: google.cloud.dataplex.v1.ScannedData.incremental_field:type_name -> google.cloud.dataplex.v1.ScannedData.IncrementalField
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_dataplex_v1_processing_proto_init() }
func file_google_cloud_dataplex_v1_processing_proto_init() {
	if File_google_cloud_dataplex_v1_processing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dataplex_v1_processing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger); i {
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
		file_google_cloud_dataplex_v1_processing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSource); i {
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
		file_google_cloud_dataplex_v1_processing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScannedData); i {
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
		file_google_cloud_dataplex_v1_processing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger_OnDemand); i {
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
		file_google_cloud_dataplex_v1_processing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger_Schedule); i {
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
		file_google_cloud_dataplex_v1_processing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScannedData_IncrementalField); i {
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
	file_google_cloud_dataplex_v1_processing_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Trigger_OnDemand_)(nil),
		(*Trigger_Schedule_)(nil),
	}
	file_google_cloud_dataplex_v1_processing_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DataSource_Entity)(nil),
		(*DataSource_Resource)(nil),
	}
	file_google_cloud_dataplex_v1_processing_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ScannedData_IncrementalField_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dataplex_v1_processing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dataplex_v1_processing_proto_goTypes,
		DependencyIndexes: file_google_cloud_dataplex_v1_processing_proto_depIdxs,
		MessageInfos:      file_google_cloud_dataplex_v1_processing_proto_msgTypes,
	}.Build()
	File_google_cloud_dataplex_v1_processing_proto = out.File
	file_google_cloud_dataplex_v1_processing_proto_rawDesc = nil
	file_google_cloud_dataplex_v1_processing_proto_goTypes = nil
	file_google_cloud_dataplex_v1_processing_proto_depIdxs = nil
}