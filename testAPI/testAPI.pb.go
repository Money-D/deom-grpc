// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: testAPI.proto

package testAPI

import (
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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age   int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testAPI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_testAPI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_testAPI_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person   *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	SomeData string  `protobuf:"bytes,2,opt,name=someData,proto3" json:"someData,omitempty"`
}

func (x *GetDataRequest) Reset() {
	*x = GetDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testAPI_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataRequest) ProtoMessage() {}

func (x *GetDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testAPI_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataRequest.ProtoReflect.Descriptor instead.
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return file_testAPI_proto_rawDescGZIP(), []int{1}
}

func (x *GetDataRequest) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

func (x *GetDataRequest) GetSomeData() string {
	if x != nil {
		return x.SomeData
	}
	return ""
}

type GetDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pdf []byte `protobuf:"bytes,1,opt,name=pdf,proto3" json:"pdf,omitempty"`
	Png []byte `protobuf:"bytes,2,opt,name=png,proto3" json:"png,omitempty"`
}

func (x *GetDataResponse) Reset() {
	*x = GetDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testAPI_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataResponse) ProtoMessage() {}

func (x *GetDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testAPI_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataResponse.ProtoReflect.Descriptor instead.
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return file_testAPI_proto_rawDescGZIP(), []int{2}
}

func (x *GetDataResponse) GetPdf() []byte {
	if x != nil {
		return x.Pdf
	}
	return nil
}

func (x *GetDataResponse) GetPng() []byte {
	if x != nil {
		return x.Png
	}
	return nil
}

var File_testAPI_proto protoreflect.FileDescriptor

var file_testAPI_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x41, 0x50, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x44, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6d, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6d, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x64, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x64, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x6e, 0x67, 0x32, 0x37, 0x0a, 0x07, 0x74,
	0x65, 0x73, 0x74, 0x41, 0x50, 0x49, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x44, 0x5a, 0x42, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x6f, 0x6e,
	0x4d, 0x61, 0x72, 0x63, 0x6f, 0x74, 0x74, 0x65, 0x2f, 0x4c, 0x69, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_testAPI_proto_rawDescOnce sync.Once
	file_testAPI_proto_rawDescData = file_testAPI_proto_rawDesc
)

func file_testAPI_proto_rawDescGZIP() []byte {
	file_testAPI_proto_rawDescOnce.Do(func() {
		file_testAPI_proto_rawDescData = protoimpl.X.CompressGZIP(file_testAPI_proto_rawDescData)
	})
	return file_testAPI_proto_rawDescData
}

var file_testAPI_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_testAPI_proto_goTypes = []interface{}{
	(*Person)(nil),          // 0: Person
	(*GetDataRequest)(nil),  // 1: GetDataRequest
	(*GetDataResponse)(nil), // 2: GetDataResponse
}
var file_testAPI_proto_depIdxs = []int32{
	0, // 0: GetDataRequest.person:type_name -> Person
	1, // 1: testAPI.GetData:input_type -> GetDataRequest
	2, // 2: testAPI.GetData:output_type -> GetDataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_testAPI_proto_init() }
func file_testAPI_proto_init() {
	if File_testAPI_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testAPI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_testAPI_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataRequest); i {
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
		file_testAPI_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataResponse); i {
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
			RawDescriptor: file_testAPI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testAPI_proto_goTypes,
		DependencyIndexes: file_testAPI_proto_depIdxs,
		MessageInfos:      file_testAPI_proto_msgTypes,
	}.Build()
	File_testAPI_proto = out.File
	file_testAPI_proto_rawDesc = nil
	file_testAPI_proto_goTypes = nil
	file_testAPI_proto_depIdxs = nil
}
