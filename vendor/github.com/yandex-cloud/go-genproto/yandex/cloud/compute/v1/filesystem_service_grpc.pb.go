// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/compute/v1/filesystem_service.proto

package compute

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FilesystemServiceClient is the client API for FilesystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilesystemServiceClient interface {
	// Returns the specified filesystem.
	//
	// To get the list of available filesystems, make a [List] request.
	Get(ctx context.Context, in *GetFilesystemRequest, opts ...grpc.CallOption) (*Filesystem, error)
	// Lists filesystems in the specified folder.
	List(ctx context.Context, in *ListFilesystemsRequest, opts ...grpc.CallOption) (*ListFilesystemsResponse, error)
	// Creates a filesystem in the specified folder.
	Create(ctx context.Context, in *CreateFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified filesystem.
	Update(ctx context.Context, in *UpdateFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified filesystem.
	//
	// Deleting a filesystem removes its data permanently and is irreversible.
	//
	// It is not possible to delete a filesystem that is attached to an instance.
	Delete(ctx context.Context, in *DeleteFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified filesystem.
	ListOperations(ctx context.Context, in *ListFilesystemOperationsRequest, opts ...grpc.CallOption) (*ListFilesystemOperationsResponse, error)
}

type filesystemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesystemServiceClient(cc grpc.ClientConnInterface) FilesystemServiceClient {
	return &filesystemServiceClient{cc}
}

func (c *filesystemServiceClient) Get(ctx context.Context, in *GetFilesystemRequest, opts ...grpc.CallOption) (*Filesystem, error) {
	out := new(Filesystem)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.FilesystemService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) List(ctx context.Context, in *ListFilesystemsRequest, opts ...grpc.CallOption) (*ListFilesystemsResponse, error) {
	out := new(ListFilesystemsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.FilesystemService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) Create(ctx context.Context, in *CreateFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.FilesystemService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) Update(ctx context.Context, in *UpdateFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.FilesystemService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) Delete(ctx context.Context, in *DeleteFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.FilesystemService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) ListOperations(ctx context.Context, in *ListFilesystemOperationsRequest, opts ...grpc.CallOption) (*ListFilesystemOperationsResponse, error) {
	out := new(ListFilesystemOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.FilesystemService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilesystemServiceServer is the server API for FilesystemService service.
// All implementations should embed UnimplementedFilesystemServiceServer
// for forward compatibility
type FilesystemServiceServer interface {
	// Returns the specified filesystem.
	//
	// To get the list of available filesystems, make a [List] request.
	Get(context.Context, *GetFilesystemRequest) (*Filesystem, error)
	// Lists filesystems in the specified folder.
	List(context.Context, *ListFilesystemsRequest) (*ListFilesystemsResponse, error)
	// Creates a filesystem in the specified folder.
	Create(context.Context, *CreateFilesystemRequest) (*operation.Operation, error)
	// Updates the specified filesystem.
	Update(context.Context, *UpdateFilesystemRequest) (*operation.Operation, error)
	// Deletes the specified filesystem.
	//
	// Deleting a filesystem removes its data permanently and is irreversible.
	//
	// It is not possible to delete a filesystem that is attached to an instance.
	Delete(context.Context, *DeleteFilesystemRequest) (*operation.Operation, error)
	// Lists operations for the specified filesystem.
	ListOperations(context.Context, *ListFilesystemOperationsRequest) (*ListFilesystemOperationsResponse, error)
}

// UnimplementedFilesystemServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFilesystemServiceServer struct {
}

func (UnimplementedFilesystemServiceServer) Get(context.Context, *GetFilesystemRequest) (*Filesystem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFilesystemServiceServer) List(context.Context, *ListFilesystemsRequest) (*ListFilesystemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFilesystemServiceServer) Create(context.Context, *CreateFilesystemRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFilesystemServiceServer) Update(context.Context, *UpdateFilesystemRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFilesystemServiceServer) Delete(context.Context, *DeleteFilesystemRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFilesystemServiceServer) ListOperations(context.Context, *ListFilesystemOperationsRequest) (*ListFilesystemOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}

// UnsafeFilesystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilesystemServiceServer will
// result in compilation errors.
type UnsafeFilesystemServiceServer interface {
	mustEmbedUnimplementedFilesystemServiceServer()
}

func RegisterFilesystemServiceServer(s grpc.ServiceRegistrar, srv FilesystemServiceServer) {
	s.RegisterService(&FilesystemService_ServiceDesc, srv)
}

func _FilesystemService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.FilesystemService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).Get(ctx, req.(*GetFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesystemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.FilesystemService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).List(ctx, req.(*ListFilesystemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.FilesystemService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).Create(ctx, req.(*CreateFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.FilesystemService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).Update(ctx, req.(*UpdateFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.FilesystemService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).Delete(ctx, req.(*DeleteFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesystemOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.FilesystemService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).ListOperations(ctx, req.(*ListFilesystemOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FilesystemService_ServiceDesc is the grpc.ServiceDesc for FilesystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FilesystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.FilesystemService",
	HandlerType: (*FilesystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FilesystemService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FilesystemService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FilesystemService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FilesystemService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FilesystemService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _FilesystemService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/filesystem_service.proto",
}
