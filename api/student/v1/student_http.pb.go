// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v3.21.9
// source: student/v1/student.proto

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

const OperationStudentCallStudent = "/api.student.v1.Student/CallStudent"

type StudentHTTPServer interface {
	CallStudent(context.Context, *StudentRequest) (*StudentReply, error)
}

func RegisterStudentHTTPServer(s *http.Server, srv StudentHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/student", _Student_CallStudent0_HTTP_Handler(srv))
	r.GET("/v1/student/{name}/{age}", _Student_CallStudent1_HTTP_Handler(srv))
}

func _Student_CallStudent0_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StudentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentCallStudent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CallStudent(ctx, req.(*StudentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StudentReply)
		return ctx.Result(200, reply)
	}
}

func _Student_CallStudent1_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StudentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentCallStudent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CallStudent(ctx, req.(*StudentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StudentReply)
		return ctx.Result(200, reply)
	}
}

type StudentHTTPClient interface {
	CallStudent(ctx context.Context, req *StudentRequest, opts ...http.CallOption) (rsp *StudentReply, err error)
}

type StudentHTTPClientImpl struct {
	cc *http.Client
}

func NewStudentHTTPClient(client *http.Client) StudentHTTPClient {
	return &StudentHTTPClientImpl{client}
}

func (c *StudentHTTPClientImpl) CallStudent(ctx context.Context, in *StudentRequest, opts ...http.CallOption) (*StudentReply, error) {
	var out StudentReply
	pattern := "/v1/student/{name}/{age}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStudentCallStudent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
