// Code generated by protoc-gen-go-api. DO NOT EDIT.

package registry

import (
	internal "github.com/any-lyu/interface_golang/golang_api/internal"
	_ "github.com/any-lyu/interface_golang/golang_api/mall_interface"
	proto "github.com/any-lyu/interface_golang/golang_api/proto"
)

// RegisterServices register services to rpc server
func RegisterServices(env string, server proto.Server, invokerFactory proto.InvokerFactory, methodOptionFuncs proto.MethodOptionFuncs) {
	for _, fn := range internal.RegisterServiceFuncArray {
		fn(env, server, invokerFactory, methodOptionFuncs)
	}
}