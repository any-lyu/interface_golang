// Code generated by protoc-gen-go-api. DO NOT EDIT.

package proto

import (
	context "context"
	messagebus_request_to_load_test_interface "github.com/any-lyu/interface_golang/golang/messagebus_request_to_load_test_interface"
)

type requestToLoadTestContextKey struct{}

var _requestToLoadTestContextKey requestToLoadTestContextKey

func ContextWithRequestToLoadTest(ctx context.Context, v *messagebus_request_to_load_test_interface.Message) context.Context {
	if v == nil {
		return ctx
	}
	return context.WithValue(ctx, _requestToLoadTestContextKey, v)
}

func RequestToLoadTestFromContext(ctx context.Context) (v *messagebus_request_to_load_test_interface.Message, ok bool) {
	v, ok = ctx.Value(_requestToLoadTestContextKey).(*messagebus_request_to_load_test_interface.Message)
	return v, ok
}
