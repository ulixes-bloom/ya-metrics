// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: metrics.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Metric struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mtype         string                 `protobuf:"bytes,2,opt,name=mtype,proto3" json:"mtype,omitempty"`
	Delta         *int64                 `protobuf:"varint,3,opt,name=delta,proto3,oneof" json:"delta,omitempty"`
	Value         *float64               `protobuf:"fixed64,4,opt,name=value,proto3,oneof" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Metric) Reset() {
	*x = Metric{}
	mi := &file_metrics_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Metric) GetMtype() string {
	if x != nil {
		return x.Mtype
	}
	return ""
}

func (x *Metric) GetDelta() int64 {
	if x != nil && x.Delta != nil {
		return *x.Delta
	}
	return 0
}

func (x *Metric) GetValue() float64 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

type UpdateMetricRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metric        *Metric                `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	Hash          *string                `protobuf:"bytes,2,opt,name=hash,proto3,oneof" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateMetricRequest) Reset() {
	*x = UpdateMetricRequest{}
	mi := &file_metrics_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMetricRequest) ProtoMessage() {}

func (x *UpdateMetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMetricRequest.ProtoReflect.Descriptor instead.
func (*UpdateMetricRequest) Descriptor() ([]byte, []int) {
	return file_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateMetricRequest) GetMetric() *Metric {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *UpdateMetricRequest) GetHash() string {
	if x != nil && x.Hash != nil {
		return *x.Hash
	}
	return ""
}

var File_metrics_proto protoreflect.FileDescriptor

var file_metrics_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x63, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x17, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x32, 0x55, 0x0a, 0x0a, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x47, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x1f, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0f,
	0x5a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_metrics_proto_rawDescOnce sync.Once
	file_metrics_proto_rawDescData []byte
)

func file_metrics_proto_rawDescGZIP() []byte {
	file_metrics_proto_rawDescOnce.Do(func() {
		file_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metrics_proto_rawDesc), len(file_metrics_proto_rawDesc)))
	})
	return file_metrics_proto_rawDescData
}

var file_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_metrics_proto_goTypes = []any{
	(*Metric)(nil),              // 0: monitoring.Metric
	(*UpdateMetricRequest)(nil), // 1: monitoring.UpdateMetricRequest
	(*emptypb.Empty)(nil),       // 2: google.protobuf.Empty
}
var file_metrics_proto_depIdxs = []int32{
	0, // 0: monitoring.UpdateMetricRequest.metric:type_name -> monitoring.Metric
	1, // 1: monitoring.Monitoring.UpdateMetric:input_type -> monitoring.UpdateMetricRequest
	2, // 2: monitoring.Monitoring.UpdateMetric:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_metrics_proto_init() }
func file_metrics_proto_init() {
	if File_metrics_proto != nil {
		return
	}
	file_metrics_proto_msgTypes[0].OneofWrappers = []any{}
	file_metrics_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metrics_proto_rawDesc), len(file_metrics_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metrics_proto_goTypes,
		DependencyIndexes: file_metrics_proto_depIdxs,
		MessageInfos:      file_metrics_proto_msgTypes,
	}.Build()
	File_metrics_proto = out.File
	file_metrics_proto_goTypes = nil
	file_metrics_proto_depIdxs = nil
}
