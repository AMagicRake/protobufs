// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: complex.proto

package proto

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

type Dummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Dummy) Reset() {
	*x = Dummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dummy) ProtoMessage() {}

func (x *Dummy) ProtoReflect() protoreflect.Message {
	mi := &file_complex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dummy.ProtoReflect.Descriptor instead.
func (*Dummy) Descriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{0}
}

func (x *Dummy) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dummy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Complex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OneDummy   *Dummy   `protobuf:"bytes,1,opt,name=one_dummy,json=oneDummy,proto3" json:"one_dummy,omitempty"`
	MultiDummy []*Dummy `protobuf:"bytes,2,rep,name=multi_dummy,json=multiDummy,proto3" json:"multi_dummy,omitempty"`
}

func (x *Complex) Reset() {
	*x = Complex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Complex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Complex) ProtoMessage() {}

func (x *Complex) ProtoReflect() protoreflect.Message {
	mi := &file_complex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Complex.ProtoReflect.Descriptor instead.
func (*Complex) Descriptor() ([]byte, []int) {
	return file_complex_proto_rawDescGZIP(), []int{1}
}

func (x *Complex) GetOneDummy() *Dummy {
	if x != nil {
		return x.OneDummy
	}
	return nil
}

func (x *Complex) GetMultiDummy() []*Dummy {
	if x != nil {
		return x.MultiDummy
	}
	return nil
}

var File_complex_proto protoreflect.FileDescriptor

var file_complex_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x22, 0x2b, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x77, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x12, 0x33, 0x0a, 0x09, 0x6f, 0x6e, 0x65, 0x5f,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2e, 0x44, 0x75,
	0x6d, 0x6d, 0x79, 0x52, 0x08, 0x6f, 0x6e, 0x65, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x12, 0x37, 0x0a,
	0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x0a, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_complex_proto_rawDescOnce sync.Once
	file_complex_proto_rawDescData = file_complex_proto_rawDesc
)

func file_complex_proto_rawDescGZIP() []byte {
	file_complex_proto_rawDescOnce.Do(func() {
		file_complex_proto_rawDescData = protoimpl.X.CompressGZIP(file_complex_proto_rawDescData)
	})
	return file_complex_proto_rawDescData
}

var file_complex_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_complex_proto_goTypes = []any{
	(*Dummy)(nil),   // 0: example.complex.Dummy
	(*Complex)(nil), // 1: example.complex.Complex
}
var file_complex_proto_depIdxs = []int32{
	0, // 0: example.complex.Complex.one_dummy:type_name -> example.complex.Dummy
	0, // 1: example.complex.Complex.multi_dummy:type_name -> example.complex.Dummy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_complex_proto_init() }
func file_complex_proto_init() {
	if File_complex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_complex_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Dummy); i {
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
		file_complex_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Complex); i {
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
			RawDescriptor: file_complex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_complex_proto_goTypes,
		DependencyIndexes: file_complex_proto_depIdxs,
		MessageInfos:      file_complex_proto_msgTypes,
	}.Build()
	File_complex_proto = out.File
	file_complex_proto_rawDesc = nil
	file_complex_proto_goTypes = nil
	file_complex_proto_depIdxs = nil
}
