// Code generated by protoc-gen-go-api. DO NOT EDIT.

package mall_interface

import (
	context "context"
	fmt "fmt"
	mall_interface "github.com/any-lyu/interface_golang/golang/mall_interface"
	internal "github.com/any-lyu/interface_golang/golang_api/internal"
	proto "github.com/any-lyu/interface_golang/golang_api/proto"
)

func init() {
	internal.RegisterServiceFuncArray = append(internal.RegisterServiceFuncArray, registerDefaultServer)
}

func registerDefaultServer(env string, server proto.Server, invokerFactory proto.InvokerFactory, methodOptionFuncs proto.MethodOptionFuncs) {
	mall_interface.RegisterDefaultServer(server, newDefaultServer(env, invokerFactory, methodOptionFuncs))
}

func newDefaultServer(env string, invokerFactory proto.InvokerFactory, methodOptionFuncs proto.MethodOptionFuncs) mall_interface.DefaultServer {
	var invoker proto.Invoker
	switch env {
	case "prd":
		invoker = invokerFactory.MakeInvoker("mall-inner.any-lyu.net:10924")
	case "pre":
		invoker = invokerFactory.MakeInvoker("127.0.0.1:10924")
	case "qa":
		invoker = invokerFactory.MakeInvoker("127.0.0.1:10924")
	default:
		panic(fmt.Sprintf("invalid env: %q", env))
	}

	return &defaultServer{
		client:            mall_interface.NewDefaultClient(invoker),
		methodOptionFuncs: methodOptionFuncs,
	}
}

// defaultServer implements mall_interface.DefaultServer interface
type defaultServer struct {
	client            mall_interface.DefaultClient
	methodOptionFuncs proto.MethodOptionFuncs
}

// Login 登录
func (srv *defaultServer) Login(ctx context.Context, in *mall_interface.LoginRequest) (*mall_interface.LoginReply, error) {
	return srv.client.Login(ctx, in)
}