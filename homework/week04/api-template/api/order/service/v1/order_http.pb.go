// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type OrderHTTPServer interface {
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderReply, error)
	DeleteOrder(context.Context, *DeleteOrderReq) (*DeleteOrderReply, error)
	GetOrder(context.Context, *GetOrderReq) (*GetOrderReply, error)
	ListOrder(context.Context, *ListOrderReq) (*ListOrderReply, error)
	SearchOrder(context.Context, *SearchOrderReq) (*SearchOrderReply, error)
	UpdateOrder(context.Context, *UpdateOrderReq) (*UpdateOrderReply, error)
}

func RegisterOrderHTTPServer(s *http.Server, srv OrderHTTPServer) {
	r := s.Route("/")
	r.GET("/orders/v1/{id}", _Order_GetOrder0_HTTP_Handler(srv))
	r.DELETE("/orders/v1", _Order_DeleteOrder0_HTTP_Handler(srv))
	r.POST("/orders/v1", _Order_CreateOrder0_HTTP_Handler(srv))
	r.PUT("/orders/v1", _Order_UpdateOrder0_HTTP_Handler(srv))
	r.GET("/orders/v1", _Order_ListOrder0_HTTP_Handler(srv))
	r.GET("/orders/v1/search", _Order_SearchOrder0_HTTP_Handler(srv))
}

func _Order_GetOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOrderReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/order.service.v1.Order/GetOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOrder(ctx, req.(*GetOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_DeleteOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteOrderReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/order.service.v1.Order/DeleteOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteOrder(ctx, req.(*DeleteOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_CreateOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOrderReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/order.service.v1.Order/CreateOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOrder(ctx, req.(*CreateOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_UpdateOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateOrderReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/order.service.v1.Order/UpdateOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateOrder(ctx, req.(*UpdateOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_ListOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOrderReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/order.service.v1.Order/ListOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListOrder(ctx, req.(*ListOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_SearchOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SearchOrderReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/order.service.v1.Order/SearchOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SearchOrder(ctx, req.(*SearchOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SearchOrderReply)
		return ctx.Result(200, reply)
	}
}

type OrderHTTPClient interface {
	CreateOrder(ctx context.Context, req *CreateOrderReq, opts ...http.CallOption) (rsp *CreateOrderReply, err error)
	DeleteOrder(ctx context.Context, req *DeleteOrderReq, opts ...http.CallOption) (rsp *DeleteOrderReply, err error)
	GetOrder(ctx context.Context, req *GetOrderReq, opts ...http.CallOption) (rsp *GetOrderReply, err error)
	ListOrder(ctx context.Context, req *ListOrderReq, opts ...http.CallOption) (rsp *ListOrderReply, err error)
	SearchOrder(ctx context.Context, req *SearchOrderReq, opts ...http.CallOption) (rsp *SearchOrderReply, err error)
	UpdateOrder(ctx context.Context, req *UpdateOrderReq, opts ...http.CallOption) (rsp *UpdateOrderReply, err error)
}

type OrderHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderHTTPClient(client *http.Client) OrderHTTPClient {
	return &OrderHTTPClientImpl{client}
}

func (c *OrderHTTPClientImpl) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...http.CallOption) (*CreateOrderReply, error) {
	var out CreateOrderReply
	pattern := "/orders/v1"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/order.service.v1.Order/CreateOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...http.CallOption) (*DeleteOrderReply, error) {
	var out DeleteOrderReply
	pattern := "/orders/v1"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/order.service.v1.Order/DeleteOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) GetOrder(ctx context.Context, in *GetOrderReq, opts ...http.CallOption) (*GetOrderReply, error) {
	var out GetOrderReply
	pattern := "/orders/v1/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/order.service.v1.Order/GetOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) ListOrder(ctx context.Context, in *ListOrderReq, opts ...http.CallOption) (*ListOrderReply, error) {
	var out ListOrderReply
	pattern := "/orders/v1"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/order.service.v1.Order/ListOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) SearchOrder(ctx context.Context, in *SearchOrderReq, opts ...http.CallOption) (*SearchOrderReply, error) {
	var out SearchOrderReply
	pattern := "/orders/v1/search"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/order.service.v1.Order/SearchOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...http.CallOption) (*UpdateOrderReply, error) {
	var out UpdateOrderReply
	pattern := "/orders/v1"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/order.service.v1.Order/UpdateOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}