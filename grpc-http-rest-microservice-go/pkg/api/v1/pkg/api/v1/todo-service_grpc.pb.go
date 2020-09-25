// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ToDoServiceClient is the client API for ToDoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToDoServiceClient interface {
	// Create new todo task
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read todo task
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update todo task
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete todo task
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type toDoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToDoServiceClient(cc grpc.ClientConnInterface) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

var toDoServiceCreateStreamDesc = &grpc.StreamDesc{
	StreamName: "Create",
}

func (c *toDoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var toDoServiceReadStreamDesc = &grpc.StreamDesc{
	StreamName: "Read",
}

func (c *toDoServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var toDoServiceUpdateStreamDesc = &grpc.StreamDesc{
	StreamName: "Update",
}

func (c *toDoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var toDoServiceDeleteStreamDesc = &grpc.StreamDesc{
	StreamName: "Delete",
}

func (c *toDoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoServiceService is the service API for ToDoService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterToDoServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type ToDoServiceService struct {
	// Create new todo task
	Create func(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read todo task
	Read func(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update todo task
	Update func(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete todo task
	Delete func(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func (s *ToDoServiceService) create(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v1.ToDoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *ToDoServiceService) read(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v1.ToDoService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *ToDoServiceService) update(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v1.ToDoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *ToDoServiceService) delete(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v1.ToDoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterToDoServiceService registers a service implementation with a gRPC server.
func RegisterToDoServiceService(s grpc.ServiceRegistrar, srv *ToDoServiceService) {
	srvCopy := *srv
	if srvCopy.Create == nil {
		srvCopy.Create = func(context.Context, *CreateRequest) (*CreateResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
		}
	}
	if srvCopy.Read == nil {
		srvCopy.Read = func(context.Context, *ReadRequest) (*ReadResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
		}
	}
	if srvCopy.Update == nil {
		srvCopy.Update = func(context.Context, *UpdateRequest) (*UpdateResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
		}
	}
	if srvCopy.Delete == nil {
		srvCopy.Delete = func(context.Context, *DeleteRequest) (*DeleteResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "v1.ToDoService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Create",
				Handler:    srvCopy.create,
			},
			{
				MethodName: "Read",
				Handler:    srvCopy.read,
			},
			{
				MethodName: "Update",
				Handler:    srvCopy.update,
			},
			{
				MethodName: "Delete",
				Handler:    srvCopy.delete,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "todo-service.proto",
	}

	s.RegisterService(&sd, nil)
}
