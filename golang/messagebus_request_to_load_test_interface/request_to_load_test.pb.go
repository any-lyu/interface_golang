// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: messagebus/request_to_load_test/request_to_load_test.proto

package messagebus_request_to_load_test_interface

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

// Header http请求的header
type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key header的key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// values header的values
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagebus_request_to_load_test_request_to_load_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_messagebus_request_to_load_test_request_to_load_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_messagebus_request_to_load_test_request_to_load_test_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// Message 压测的一个请求体
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// headers http headers
	Headers []*Header `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	// method http请求的path
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// body http请求体
	Body []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagebus_request_to_load_test_request_to_load_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_messagebus_request_to_load_test_request_to_load_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_messagebus_request_to_load_test_request_to_load_test_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Message) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Message) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_messagebus_request_to_load_test_request_to_load_test_proto protoreflect.FileDescriptor

var file_messagebus_request_to_load_test_request_to_load_test_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x75, 0x73, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x75, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x22, 0x78, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x75, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x56, 0x5a, 0x54, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x79, 0x2d, 0x6c, 0x79,
	0x75, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x62, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x5f,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messagebus_request_to_load_test_request_to_load_test_proto_rawDescOnce sync.Once
	file_messagebus_request_to_load_test_request_to_load_test_proto_rawDescData = file_messagebus_request_to_load_test_request_to_load_test_proto_rawDesc
)

func file_messagebus_request_to_load_test_request_to_load_test_proto_rawDescGZIP() []byte {
	file_messagebus_request_to_load_test_request_to_load_test_proto_rawDescOnce.Do(func() {
		file_messagebus_request_to_load_test_request_to_load_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_messagebus_request_to_load_test_request_to_load_test_proto_rawDescData)
	})
	return file_messagebus_request_to_load_test_request_to_load_test_proto_rawDescData
}

var file_messagebus_request_to_load_test_request_to_load_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_messagebus_request_to_load_test_request_to_load_test_proto_goTypes = []interface{}{
	(*Header)(nil),  // 0: messagebus.request_to_load_test.Header
	(*Message)(nil), // 1: messagebus.request_to_load_test.Message
}
var file_messagebus_request_to_load_test_request_to_load_test_proto_depIdxs = []int32{
	0, // 0: messagebus.request_to_load_test.Message.headers:type_name -> messagebus.request_to_load_test.Header
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messagebus_request_to_load_test_request_to_load_test_proto_init() }
func file_messagebus_request_to_load_test_request_to_load_test_proto_init() {
	if File_messagebus_request_to_load_test_request_to_load_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messagebus_request_to_load_test_request_to_load_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_messagebus_request_to_load_test_request_to_load_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_messagebus_request_to_load_test_request_to_load_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messagebus_request_to_load_test_request_to_load_test_proto_goTypes,
		DependencyIndexes: file_messagebus_request_to_load_test_request_to_load_test_proto_depIdxs,
		MessageInfos:      file_messagebus_request_to_load_test_request_to_load_test_proto_msgTypes,
	}.Build()
	File_messagebus_request_to_load_test_request_to_load_test_proto = out.File
	file_messagebus_request_to_load_test_request_to_load_test_proto_rawDesc = nil
	file_messagebus_request_to_load_test_request_to_load_test_proto_goTypes = nil
	file_messagebus_request_to_load_test_request_to_load_test_proto_depIdxs = nil
}
