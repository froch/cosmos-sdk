// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: apis/orm/v1alpha1/orm.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StateObjectDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypePrefix        []byte                              `protobuf:"bytes,1,opt,name=type_prefix,json=typePrefix,proto3" json:"type_prefix,omitempty"`
	TableDescriptor   *TableDescriptor                    `protobuf:"bytes,2,opt,name=table_descriptor,json=tableDescriptor,proto3" json:"table_descriptor,omitempty"`
	FileDescriptor    *descriptorpb.FileDescriptorProto   `protobuf:"bytes,3,opt,name=file_descriptor,json=fileDescriptor,proto3" json:"file_descriptor,omitempty"`
	ProtoDependencies []*descriptorpb.FileDescriptorProto `protobuf:"bytes,4,rep,name=proto_dependencies,json=protoDependencies,proto3" json:"proto_dependencies,omitempty"`
}

func (x *StateObjectDescriptor) Reset() {
	*x = StateObjectDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_orm_v1alpha1_orm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateObjectDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateObjectDescriptor) ProtoMessage() {}

func (x *StateObjectDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_apis_orm_v1alpha1_orm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateObjectDescriptor.ProtoReflect.Descriptor instead.
func (*StateObjectDescriptor) Descriptor() ([]byte, []int) {
	return file_apis_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{0}
}

func (x *StateObjectDescriptor) GetTypePrefix() []byte {
	if x != nil {
		return x.TypePrefix
	}
	return nil
}

func (x *StateObjectDescriptor) GetTableDescriptor() *TableDescriptor {
	if x != nil {
		return x.TableDescriptor
	}
	return nil
}

func (x *StateObjectDescriptor) GetFileDescriptor() *descriptorpb.FileDescriptorProto {
	if x != nil {
		return x.FileDescriptor
	}
	return nil
}

func (x *StateObjectDescriptor) GetProtoDependencies() []*descriptorpb.FileDescriptorProto {
	if x != nil {
		return x.ProtoDependencies
	}
	return nil
}

type TableDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// primary_key_descriptor describes how the state object primary key
	// is formed.
	PrimaryKeyDescriptor *PrimaryKeyDescriptor `protobuf:"bytes,1,opt,name=primary_key_descriptor,json=primaryKeyDescriptor,proto3" json:"primary_key_descriptor,omitempty"`
	// secondary_key_descriptors describes how the state object is indexed.
	SecondaryKeyDescriptors []*SecondaryKeyDescriptor `protobuf:"bytes,2,rep,name=secondary_key_descriptors,json=secondaryKeyDescriptors,proto3" json:"secondary_key_descriptors,omitempty"`
	// singleton signals that can exist only one instance of this object
	// if singleton is true, primary_key_descriptor and secondary_key_descriptors
	// must not be set.
	Singleton bool `protobuf:"varint,3,opt,name=singleton,proto3" json:"singleton,omitempty"`
}

func (x *TableDescriptor) Reset() {
	*x = TableDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_orm_v1alpha1_orm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableDescriptor) ProtoMessage() {}

func (x *TableDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_apis_orm_v1alpha1_orm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableDescriptor.ProtoReflect.Descriptor instead.
func (*TableDescriptor) Descriptor() ([]byte, []int) {
	return file_apis_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{1}
}

func (x *TableDescriptor) GetPrimaryKeyDescriptor() *PrimaryKeyDescriptor {
	if x != nil {
		return x.PrimaryKeyDescriptor
	}
	return nil
}

func (x *TableDescriptor) GetSecondaryKeyDescriptors() []*SecondaryKeyDescriptor {
	if x != nil {
		return x.SecondaryKeyDescriptors
	}
	return nil
}

func (x *TableDescriptor) GetSingleton() bool {
	if x != nil {
		return x.Singleton
	}
	return false
}

type PrimaryKeyDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtobufFieldNames []string `protobuf:"bytes,1,rep,name=protobuf_field_names,json=protobufFieldNames,proto3" json:"protobuf_field_names,omitempty"`
}

func (x *PrimaryKeyDescriptor) Reset() {
	*x = PrimaryKeyDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_orm_v1alpha1_orm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryKeyDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryKeyDescriptor) ProtoMessage() {}

func (x *PrimaryKeyDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_apis_orm_v1alpha1_orm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryKeyDescriptor.ProtoReflect.Descriptor instead.
func (*PrimaryKeyDescriptor) Descriptor() ([]byte, []int) {
	return file_apis_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{2}
}

func (x *PrimaryKeyDescriptor) GetProtobufFieldNames() []string {
	if x != nil {
		return x.ProtobufFieldNames
	}
	return nil
}

type SecondaryKeyDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtobufFieldName string `protobuf:"bytes,2,opt,name=protobuf_field_name,json=protobufFieldName,proto3" json:"protobuf_field_name,omitempty"`
}

func (x *SecondaryKeyDescriptor) Reset() {
	*x = SecondaryKeyDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_orm_v1alpha1_orm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecondaryKeyDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecondaryKeyDescriptor) ProtoMessage() {}

func (x *SecondaryKeyDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_apis_orm_v1alpha1_orm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecondaryKeyDescriptor.ProtoReflect.Descriptor instead.
func (*SecondaryKeyDescriptor) Descriptor() ([]byte, []int) {
	return file_apis_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{3}
}

func (x *SecondaryKeyDescriptor) GetProtobufFieldName() string {
	if x != nil {
		return x.ProtobufFieldName
	}
	return ""
}

var File_apis_orm_v1alpha1_orm_proto protoreflect.FileDescriptor

var file_apis_orm_v1alpha1_orm_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x02, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x4f, 0x0a, 0x10, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52,
	0x0f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x12, 0x4d, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
	0x0e, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12,
	0x53, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x22, 0xf9, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x5f, 0x0a, 0x16, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x52, 0x14, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x67, 0x0a, 0x19, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x17, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e,
	0x22, 0x48, 0x0a, 0x14, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x48, 0x0a, 0x16, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6f, 0x72,
	0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_apis_orm_v1alpha1_orm_proto_rawDescOnce sync.Once
	file_apis_orm_v1alpha1_orm_proto_rawDescData = file_apis_orm_v1alpha1_orm_proto_rawDesc
)

func file_apis_orm_v1alpha1_orm_proto_rawDescGZIP() []byte {
	file_apis_orm_v1alpha1_orm_proto_rawDescOnce.Do(func() {
		file_apis_orm_v1alpha1_orm_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_orm_v1alpha1_orm_proto_rawDescData)
	})
	return file_apis_orm_v1alpha1_orm_proto_rawDescData
}

var file_apis_orm_v1alpha1_orm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apis_orm_v1alpha1_orm_proto_goTypes = []interface{}{
	(*StateObjectDescriptor)(nil),            // 0: cosmos.orm.v1alpha1.StateObjectDescriptor
	(*TableDescriptor)(nil),                  // 1: cosmos.orm.v1alpha1.TableDescriptor
	(*PrimaryKeyDescriptor)(nil),             // 2: cosmos.orm.v1alpha1.PrimaryKeyDescriptor
	(*SecondaryKeyDescriptor)(nil),           // 3: cosmos.orm.v1alpha1.SecondaryKeyDescriptor
	(*descriptorpb.FileDescriptorProto)(nil), // 4: google.protobuf.FileDescriptorProto
}
var file_apis_orm_v1alpha1_orm_proto_depIdxs = []int32{
	1, // 0: cosmos.orm.v1alpha1.StateObjectDescriptor.table_descriptor:type_name -> cosmos.orm.v1alpha1.TableDescriptor
	4, // 1: cosmos.orm.v1alpha1.StateObjectDescriptor.file_descriptor:type_name -> google.protobuf.FileDescriptorProto
	4, // 2: cosmos.orm.v1alpha1.StateObjectDescriptor.proto_dependencies:type_name -> google.protobuf.FileDescriptorProto
	2, // 3: cosmos.orm.v1alpha1.TableDescriptor.primary_key_descriptor:type_name -> cosmos.orm.v1alpha1.PrimaryKeyDescriptor
	3, // 4: cosmos.orm.v1alpha1.TableDescriptor.secondary_key_descriptors:type_name -> cosmos.orm.v1alpha1.SecondaryKeyDescriptor
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_apis_orm_v1alpha1_orm_proto_init() }
func file_apis_orm_v1alpha1_orm_proto_init() {
	if File_apis_orm_v1alpha1_orm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_orm_v1alpha1_orm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateObjectDescriptor); i {
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
		file_apis_orm_v1alpha1_orm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableDescriptor); i {
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
		file_apis_orm_v1alpha1_orm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryKeyDescriptor); i {
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
		file_apis_orm_v1alpha1_orm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecondaryKeyDescriptor); i {
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
			RawDescriptor: file_apis_orm_v1alpha1_orm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_orm_v1alpha1_orm_proto_goTypes,
		DependencyIndexes: file_apis_orm_v1alpha1_orm_proto_depIdxs,
		MessageInfos:      file_apis_orm_v1alpha1_orm_proto_msgTypes,
	}.Build()
	File_apis_orm_v1alpha1_orm_proto = out.File
	file_apis_orm_v1alpha1_orm_proto_rawDesc = nil
	file_apis_orm_v1alpha1_orm_proto_goTypes = nil
	file_apis_orm_v1alpha1_orm_proto_depIdxs = nil
}
