// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: common/method_options.proto

package common_interface

import (
	errorcode "github.com/any-lyu/interface_golang/golang/errorcode"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// CacheMode 缓存方式
type CacheMode int32

const (
	// CACHE_MODE_UNDEFINED 未定义
	CacheMode_CACHE_MODE_UNDEFINED CacheMode = 0
	// CACHE_MODE_DEFAULT 默认缓存方式 - 优先使用未过期的缓存, 同时尝试请求网络刷新过了保鲜期的数据
	CacheMode_CACHE_MODE_DEFAULT CacheMode = 1
	// CACHE_MODE_NO_STORE 完全不使用缓存
	CacheMode_CACHE_MODE_NO_STORE CacheMode = 2
	// CACHE_MODE_NO_CACHE 不使用缓存, 但是会保存收到的数据到缓存
	CacheMode_CACHE_MODE_NO_CACHE CacheMode = 3
	// CACHE_MODE_FORCE_CACHE 强制使用缓存, 即使缓存已经过期, 有缓存的话不再刷新, 否则请求网络并保存结果到缓存
	CacheMode_CACHE_MODE_FORCE_CACHE CacheMode = 4
	// CACHE_MODE_ONLY_IF_CACHED 只使用缓存的内容, 不尝试从网络加载
	CacheMode_CACHE_MODE_ONLY_IF_CACHED CacheMode = 5
)

// Enum value maps for CacheMode.
var (
	CacheMode_name = map[int32]string{
		0: "CACHE_MODE_UNDEFINED",
		1: "CACHE_MODE_DEFAULT",
		2: "CACHE_MODE_NO_STORE",
		3: "CACHE_MODE_NO_CACHE",
		4: "CACHE_MODE_FORCE_CACHE",
		5: "CACHE_MODE_ONLY_IF_CACHED",
	}
	CacheMode_value = map[string]int32{
		"CACHE_MODE_UNDEFINED":      0,
		"CACHE_MODE_DEFAULT":        1,
		"CACHE_MODE_NO_STORE":       2,
		"CACHE_MODE_NO_CACHE":       3,
		"CACHE_MODE_FORCE_CACHE":    4,
		"CACHE_MODE_ONLY_IF_CACHED": 5,
	}
)

func (x CacheMode) Enum() *CacheMode {
	p := new(CacheMode)
	*p = x
	return p
}

func (x CacheMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CacheMode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_method_options_proto_enumTypes[0].Descriptor()
}

func (CacheMode) Type() protoreflect.EnumType {
	return &file_common_method_options_proto_enumTypes[0]
}

func (x CacheMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CacheMode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CacheMode(num)
	return nil
}

// Deprecated: Use CacheMode.Descriptor instead.
func (CacheMode) EnumDescriptor() ([]byte, []int) {
	return file_common_method_options_proto_rawDescGZIP(), []int{0}
}

// MethodOptions 是 google.protobuf.MethodOptions 的扩展
type MethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// requiredPermission 是请求方法所需要的权限.
	RequiredPermission *Permission `protobuf:"varint,1,opt,name=requiredPermission,enum=common.Permission" json:"requiredPermission,omitempty"`
	// authFree 是否不验证登录, 默认为 false, 即需要验证 token
	AuthFree *bool `protobuf:"varint,2,opt,name=authFree" json:"authFree,omitempty"`
	// officeNetworkRequired 是否要求客户端来自办公网络, 默认 false, 即不做要求
	OfficeNetworkRequired *bool `protobuf:"varint,3,opt,name=officeNetworkRequired" json:"officeNetworkRequired,omitempty"`
	// timeoutMS 超时时间, 单位为毫秒, 如果不指定则使用默认超时时间
	TimeoutMS *int32 `protobuf:"varint,4,opt,name=timeoutMS" json:"timeoutMS,omitempty"`
	// cacheControl 是方法生成的前端ts/js sdk对缓存使用的默认策略, 默认不做缓存, 设置了这个选项以后在ts/js sdk会生成对应的 XxxFetcher 对象
	CacheControl *CacheControl `protobuf:"bytes,5,opt,name=cacheControl" json:"cacheControl,omitempty"`
	// requestEncrypted request经过加密的接口, 在面向客户端的时候内容以加密的文本传输, 默认为 false, 即不加密
	RequestEncrypted *bool `protobuf:"varint,6,opt,name=requestEncrypted" json:"requestEncrypted,omitempty"`
	// replyEncrypted reply 经过加密的接口, 在面向客户端的时候内容以加密的文本传输, 默认为 false, 即不加密
	ReplyEncrypted *bool `protobuf:"varint,7,opt,name=replyEncrypted" json:"replyEncrypted,omitempty"`
	// errors 是可能返回的错误集合
	Errors []errorcode.Code `protobuf:"varint,8,rep,name=errors,enum=errorcode.Code" json:"errors,omitempty"`
	// serviceHost 是请求这个 method 的 service host, 比如 localhost, localhost:8888
	ServiceHost *string `protobuf:"bytes,9,opt,name=serviceHost" json:"serviceHost,omitempty"`
	// qaServiceHost 是测试环境请求这个 method 的 service host, 比如 localhost, localhost:8888
	QaServiceHost *string `protobuf:"bytes,100,opt,name=qaServiceHost" json:"qaServiceHost,omitempty"` // TODO: 当前没有 naming 系统, 有了 naming 系统后废除这个 option
	// preServiceHost 是预生产环境请求这个 method 的 service host, 比如 localhost, localhost:8888
	PreServiceHost *string `protobuf:"bytes,101,opt,name=preServiceHost" json:"preServiceHost,omitempty"` // TODO: 当前没有 naming 系统, 有了 naming 系统后废除这个 option
}

func (x *MethodOptions) Reset() {
	*x = MethodOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_method_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodOptions) ProtoMessage() {}

func (x *MethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_common_method_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodOptions.ProtoReflect.Descriptor instead.
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return file_common_method_options_proto_rawDescGZIP(), []int{0}
}

func (x *MethodOptions) GetRequiredPermission() Permission {
	if x != nil && x.RequiredPermission != nil {
		return *x.RequiredPermission
	}
	return Permission_PERMISSION_UNDEFINED
}

func (x *MethodOptions) GetAuthFree() bool {
	if x != nil && x.AuthFree != nil {
		return *x.AuthFree
	}
	return false
}

func (x *MethodOptions) GetOfficeNetworkRequired() bool {
	if x != nil && x.OfficeNetworkRequired != nil {
		return *x.OfficeNetworkRequired
	}
	return false
}

func (x *MethodOptions) GetTimeoutMS() int32 {
	if x != nil && x.TimeoutMS != nil {
		return *x.TimeoutMS
	}
	return 0
}

func (x *MethodOptions) GetCacheControl() *CacheControl {
	if x != nil {
		return x.CacheControl
	}
	return nil
}

func (x *MethodOptions) GetRequestEncrypted() bool {
	if x != nil && x.RequestEncrypted != nil {
		return *x.RequestEncrypted
	}
	return false
}

func (x *MethodOptions) GetReplyEncrypted() bool {
	if x != nil && x.ReplyEncrypted != nil {
		return *x.ReplyEncrypted
	}
	return false
}

func (x *MethodOptions) GetErrors() []errorcode.Code {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *MethodOptions) GetServiceHost() string {
	if x != nil && x.ServiceHost != nil {
		return *x.ServiceHost
	}
	return ""
}

func (x *MethodOptions) GetQaServiceHost() string {
	if x != nil && x.QaServiceHost != nil {
		return *x.QaServiceHost
	}
	return ""
}

func (x *MethodOptions) GetPreServiceHost() string {
	if x != nil && x.PreServiceHost != nil {
		return *x.PreServiceHost
	}
	return ""
}

// CacheControl 缓存选项
type CacheControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// mode 缓存方式
	Mode *CacheMode `protobuf:"varint,1,opt,name=mode,enum=common.CacheMode" json:"mode,omitempty"`
	// minFresh 默认缓存方式下缓存的保鲜时间, 以秒为单位 - 在这个时间间隔内客户端不会尝试刷新数据
	MinFresh *int32 `protobuf:"varint,2,opt,name=minFresh" json:"minFresh,omitempty"`
	// maxAge 默认缓存方式下缓存的过期时间, 以秒为单位 - 在这个时间间隔内的缓存才会展示给客户端
	MaxAge *int32 `protobuf:"varint,3,opt,name=maxAge" json:"maxAge,omitempty"`
}

func (x *CacheControl) Reset() {
	*x = CacheControl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_method_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheControl) ProtoMessage() {}

func (x *CacheControl) ProtoReflect() protoreflect.Message {
	mi := &file_common_method_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheControl.ProtoReflect.Descriptor instead.
func (*CacheControl) Descriptor() ([]byte, []int) {
	return file_common_method_options_proto_rawDescGZIP(), []int{1}
}

func (x *CacheControl) GetMode() CacheMode {
	if x != nil && x.Mode != nil {
		return *x.Mode
	}
	return CacheMode_CACHE_MODE_UNDEFINED
}

func (x *CacheControl) GetMinFresh() int32 {
	if x != nil && x.MinFresh != nil {
		return *x.MinFresh
	}
	return 0
}

func (x *CacheControl) GetMaxAge() int32 {
	if x != nil && x.MaxAge != nil {
		return *x.MaxAge
	}
	return 0
}

var file_common_method_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodOptions)(nil),
		Field:         50000,
		Name:          "common.methodOptions",
		Tag:           "bytes,50000,opt,name=methodOptions",
		Filename:      "common/method_options.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// methodOptions 是扩展的 MethodOptions
	//
	// optional common.MethodOptions methodOptions = 50000;
	E_MethodOptions = &file_common_method_options_proto_extTypes[0]
)

var File_common_method_options_proto protoreflect.FileDescriptor

var file_common_method_options_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xea, 0x03, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x46,
	0x72, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x46,
	0x72, 0x65, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x15, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x4d, 0x53, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x53, 0x12, 0x38, 0x0a, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x52, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x71, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x71, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x70, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x22, 0x69,
	0x0a, 0x0c, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x25,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x46, 0x72, 0x65, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x46, 0x72, 0x65, 0x73,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x2a, 0xaa, 0x01, 0x0a, 0x09, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x43, 0x48, 0x45,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x43,
	0x48, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45,
	0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x4e, 0x4f, 0x5f, 0x43, 0x41, 0x43, 0x48, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x43,
	0x41, 0x43, 0x48, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f,
	0x43, 0x41, 0x43, 0x48, 0x45, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x41, 0x43, 0x48, 0x45,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x49, 0x46, 0x5f, 0x43, 0x41,
	0x43, 0x48, 0x45, 0x44, 0x10, 0x05, 0x3a, 0x5d, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x79, 0x2d, 0x6c, 0x79, 0x75, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65,
}

var (
	file_common_method_options_proto_rawDescOnce sync.Once
	file_common_method_options_proto_rawDescData = file_common_method_options_proto_rawDesc
)

func file_common_method_options_proto_rawDescGZIP() []byte {
	file_common_method_options_proto_rawDescOnce.Do(func() {
		file_common_method_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_method_options_proto_rawDescData)
	})
	return file_common_method_options_proto_rawDescData
}

var file_common_method_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_method_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_method_options_proto_goTypes = []interface{}{
	(CacheMode)(0),                     // 0: common.CacheMode
	(*MethodOptions)(nil),              // 1: common.MethodOptions
	(*CacheControl)(nil),               // 2: common.CacheControl
	(Permission)(0),                    // 3: common.Permission
	(errorcode.Code)(0),                // 4: errorcode.Code
	(*descriptorpb.MethodOptions)(nil), // 5: google.protobuf.MethodOptions
}
var file_common_method_options_proto_depIdxs = []int32{
	3, // 0: common.MethodOptions.requiredPermission:type_name -> common.Permission
	2, // 1: common.MethodOptions.cacheControl:type_name -> common.CacheControl
	4, // 2: common.MethodOptions.errors:type_name -> errorcode.Code
	0, // 3: common.CacheControl.mode:type_name -> common.CacheMode
	5, // 4: common.methodOptions:extendee -> google.protobuf.MethodOptions
	1, // 5: common.methodOptions:type_name -> common.MethodOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	5, // [5:6] is the sub-list for extension type_name
	4, // [4:5] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_method_options_proto_init() }
func file_common_method_options_proto_init() {
	if File_common_method_options_proto != nil {
		return
	}
	file_common_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_method_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodOptions); i {
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
		file_common_method_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheControl); i {
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
			RawDescriptor: file_common_method_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_common_method_options_proto_goTypes,
		DependencyIndexes: file_common_method_options_proto_depIdxs,
		EnumInfos:         file_common_method_options_proto_enumTypes,
		MessageInfos:      file_common_method_options_proto_msgTypes,
		ExtensionInfos:    file_common_method_options_proto_extTypes,
	}.Build()
	File_common_method_options_proto = out.File
	file_common_method_options_proto_rawDesc = nil
	file_common_method_options_proto_goTypes = nil
	file_common_method_options_proto_depIdxs = nil
}
