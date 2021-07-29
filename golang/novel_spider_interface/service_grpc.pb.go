// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package novel_spider_interface

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// DefaultClient is the client API for Default service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DefaultClient interface {
	// SearchNovel 搜索小说
	SearchNovel(ctx context.Context, in *SearchNovelRequest) (*SearchNovelReply, error)
	// SearchMenu 搜索目录
	SearchMenu(ctx context.Context, in *SearchMenuRequest) (*SearchMenuReply, error)
	// SearchChapter 搜索章节
	SearchChapter(ctx context.Context, in *SearchChapterRequest) (*SearchChapterReply, error)
}

type defaultClient struct {
	cc interface {
		Invoke(ctx context.Context, method string, args, reply proto.Message) error
	}
}

func NewDefaultClient(cc interface {
	Invoke(ctx context.Context, method string, args, reply proto.Message) error
}) DefaultClient {
	return &defaultClient{cc}
}

func (c *defaultClient) SearchNovel(ctx context.Context, in *SearchNovelRequest) (*SearchNovelReply, error) {
	out := new(SearchNovelReply)
	err := c.cc.Invoke(ctx, "/novel_spider.Default/SearchNovel", in, out)
	return out, err
}

func (c *defaultClient) SearchMenu(ctx context.Context, in *SearchMenuRequest) (*SearchMenuReply, error) {
	out := new(SearchMenuReply)
	err := c.cc.Invoke(ctx, "/novel_spider.Default/SearchMenu", in, out)
	return out, err
}

func (c *defaultClient) SearchChapter(ctx context.Context, in *SearchChapterRequest) (*SearchChapterReply, error) {
	out := new(SearchChapterReply)
	err := c.cc.Invoke(ctx, "/novel_spider.Default/SearchChapter", in, out)
	return out, err
}

// DefaultServer is the server API for Default service.
type DefaultServer interface {
	// SearchNovel 搜索小说
	SearchNovel(context.Context, *SearchNovelRequest) (*SearchNovelReply, error)
	// SearchMenu 搜索目录
	SearchMenu(context.Context, *SearchMenuRequest) (*SearchMenuReply, error)
	// SearchChapter 搜索章节
	SearchChapter(context.Context, *SearchChapterRequest) (*SearchChapterReply, error)
}

// UnimplementedDefaultServer can be embedded to have forward compatible implementations.
type UnimplementedDefaultServer struct {
}

func (*UnimplementedDefaultServer) SearchNovel(context.Context, *SearchNovelRequest) (*SearchNovelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchNovel not implemented")
}
func (*UnimplementedDefaultServer) SearchMenu(context.Context, *SearchMenuRequest) (*SearchMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMenu not implemented")
}
func (*UnimplementedDefaultServer) SearchChapter(context.Context, *SearchChapterRequest) (*SearchChapterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchChapter not implemented")
}

func RegisterDefaultServer(s interface {
	RegisterService(sd *grpc.ServiceDesc, ss interface{})
}, srv DefaultServer) {
	s.RegisterService(&_Default_serviceDesc, srv)
}

func _Default_SearchNovel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchNovelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServer).SearchNovel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novel_spider.Default/SearchNovel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServer).SearchNovel(ctx, req.(*SearchNovelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Default_SearchMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServer).SearchMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novel_spider.Default/SearchMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServer).SearchMenu(ctx, req.(*SearchMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Default_SearchChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServer).SearchChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novel_spider.Default/SearchChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServer).SearchChapter(ctx, req.(*SearchChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Default_serviceDesc = grpc.ServiceDesc{
	ServiceName: "novel_spider.Default",
	HandlerType: (*DefaultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchNovel",
			Handler:    _Default_SearchNovel_Handler,
		},
		{
			MethodName: "SearchMenu",
			Handler:    _Default_SearchMenu_Handler,
		},
		{
			MethodName: "SearchChapter",
			Handler:    _Default_SearchChapter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "novel_spider/service.proto",
}