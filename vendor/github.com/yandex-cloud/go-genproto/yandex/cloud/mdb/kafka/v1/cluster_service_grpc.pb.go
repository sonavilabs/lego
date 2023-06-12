// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/mdb/kafka/v1/cluster_service.proto

package kafka

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

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	// Returns the specified Apache Kafka® cluster.
	//
	// To get the list of available Apache Kafka® clusters, make a [List] request.
	Get(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	// Retrieves the list of Apache Kafka® clusters that belong to the specified folder.
	List(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Creates a new Apache Kafka® cluster in the specified folder.
	Create(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified Apache Kafka® cluster.
	Update(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified Apache Kafka® cluster.
	Delete(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Moves the specified Apache Kafka® cluster to the specified folder.
	Move(ctx context.Context, in *MoveClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Starts the specified Apache Kafka® cluster.
	Start(ctx context.Context, in *StartClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Stops the specified Apache Kafka® cluster.
	Stop(ctx context.Context, in *StopClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Reschedule planned maintenance operation.
	RescheduleMaintenance(ctx context.Context, in *RescheduleMaintenanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Retrieves logs for the specified Apache Kafka® cluster.
	//
	// For more information about logs, see the [Logs](/docs/managed-kafka/operations/cluster-logs) section in the documentation.
	ListLogs(ctx context.Context, in *ListClusterLogsRequest, opts ...grpc.CallOption) (*ListClusterLogsResponse, error)
	// Same as [ListLogs] but using server-side streaming. Also allows for `tail -f` semantics.
	StreamLogs(ctx context.Context, in *StreamClusterLogsRequest, opts ...grpc.CallOption) (ClusterService_StreamLogsClient, error)
	// Retrieves the list of operations for the specified Apache Kafka® cluster.
	ListOperations(ctx context.Context, in *ListClusterOperationsRequest, opts ...grpc.CallOption) (*ListClusterOperationsResponse, error)
	// Retrieves a list of hosts for the specified Apache Kafka® cluster.
	ListHosts(ctx context.Context, in *ListClusterHostsRequest, opts ...grpc.CallOption) (*ListClusterHostsResponse, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) Get(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) List(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Create(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Update(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Delete(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Move(ctx context.Context, in *MoveClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/Move", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Start(ctx context.Context, in *StartClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Stop(ctx context.Context, in *StopClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) RescheduleMaintenance(ctx context.Context, in *RescheduleMaintenanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/RescheduleMaintenance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListLogs(ctx context.Context, in *ListClusterLogsRequest, opts ...grpc.CallOption) (*ListClusterLogsResponse, error) {
	out := new(ListClusterLogsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/ListLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) StreamLogs(ctx context.Context, in *StreamClusterLogsRequest, opts ...grpc.CallOption) (ClusterService_StreamLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClusterService_ServiceDesc.Streams[0], "/yandex.cloud.mdb.kafka.v1.ClusterService/StreamLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterServiceStreamLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClusterService_StreamLogsClient interface {
	Recv() (*StreamLogRecord, error)
	grpc.ClientStream
}

type clusterServiceStreamLogsClient struct {
	grpc.ClientStream
}

func (x *clusterServiceStreamLogsClient) Recv() (*StreamLogRecord, error) {
	m := new(StreamLogRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterServiceClient) ListOperations(ctx context.Context, in *ListClusterOperationsRequest, opts ...grpc.CallOption) (*ListClusterOperationsResponse, error) {
	out := new(ListClusterOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListHosts(ctx context.Context, in *ListClusterHostsRequest, opts ...grpc.CallOption) (*ListClusterHostsResponse, error) {
	out := new(ListClusterHostsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ClusterService/ListHosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations should embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	// Returns the specified Apache Kafka® cluster.
	//
	// To get the list of available Apache Kafka® clusters, make a [List] request.
	Get(context.Context, *GetClusterRequest) (*Cluster, error)
	// Retrieves the list of Apache Kafka® clusters that belong to the specified folder.
	List(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Creates a new Apache Kafka® cluster in the specified folder.
	Create(context.Context, *CreateClusterRequest) (*operation.Operation, error)
	// Updates the specified Apache Kafka® cluster.
	Update(context.Context, *UpdateClusterRequest) (*operation.Operation, error)
	// Deletes the specified Apache Kafka® cluster.
	Delete(context.Context, *DeleteClusterRequest) (*operation.Operation, error)
	// Moves the specified Apache Kafka® cluster to the specified folder.
	Move(context.Context, *MoveClusterRequest) (*operation.Operation, error)
	// Starts the specified Apache Kafka® cluster.
	Start(context.Context, *StartClusterRequest) (*operation.Operation, error)
	// Stops the specified Apache Kafka® cluster.
	Stop(context.Context, *StopClusterRequest) (*operation.Operation, error)
	// Reschedule planned maintenance operation.
	RescheduleMaintenance(context.Context, *RescheduleMaintenanceRequest) (*operation.Operation, error)
	// Retrieves logs for the specified Apache Kafka® cluster.
	//
	// For more information about logs, see the [Logs](/docs/managed-kafka/operations/cluster-logs) section in the documentation.
	ListLogs(context.Context, *ListClusterLogsRequest) (*ListClusterLogsResponse, error)
	// Same as [ListLogs] but using server-side streaming. Also allows for `tail -f` semantics.
	StreamLogs(*StreamClusterLogsRequest, ClusterService_StreamLogsServer) error
	// Retrieves the list of operations for the specified Apache Kafka® cluster.
	ListOperations(context.Context, *ListClusterOperationsRequest) (*ListClusterOperationsResponse, error)
	// Retrieves a list of hosts for the specified Apache Kafka® cluster.
	ListHosts(context.Context, *ListClusterHostsRequest) (*ListClusterHostsResponse, error)
}

// UnimplementedClusterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (UnimplementedClusterServiceServer) Get(context.Context, *GetClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedClusterServiceServer) List(context.Context, *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedClusterServiceServer) Create(context.Context, *CreateClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedClusterServiceServer) Update(context.Context, *UpdateClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedClusterServiceServer) Delete(context.Context, *DeleteClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClusterServiceServer) Move(context.Context, *MoveClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedClusterServiceServer) Start(context.Context, *StartClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedClusterServiceServer) Stop(context.Context, *StopClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedClusterServiceServer) RescheduleMaintenance(context.Context, *RescheduleMaintenanceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RescheduleMaintenance not implemented")
}
func (UnimplementedClusterServiceServer) ListLogs(context.Context, *ListClusterLogsRequest) (*ListClusterLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLogs not implemented")
}
func (UnimplementedClusterServiceServer) StreamLogs(*StreamClusterLogsRequest, ClusterService_StreamLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLogs not implemented")
}
func (UnimplementedClusterServiceServer) ListOperations(context.Context, *ListClusterOperationsRequest) (*ListClusterOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedClusterServiceServer) ListHosts(context.Context, *ListClusterHostsRequest) (*ListClusterHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHosts not implemented")
}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
	s.RegisterService(&ClusterService_ServiceDesc, srv)
}

func _ClusterService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Get(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).List(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Create(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Update(ctx, req.(*UpdateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Delete(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Move(ctx, req.(*MoveClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Start(ctx, req.(*StartClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Stop(ctx, req.(*StopClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_RescheduleMaintenance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RescheduleMaintenanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).RescheduleMaintenance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/RescheduleMaintenance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).RescheduleMaintenance(ctx, req.(*RescheduleMaintenanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/ListLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListLogs(ctx, req.(*ListClusterLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_StreamLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamClusterLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClusterServiceServer).StreamLogs(m, &clusterServiceStreamLogsServer{stream})
}

type ClusterService_StreamLogsServer interface {
	Send(*StreamLogRecord) error
	grpc.ServerStream
}

type clusterServiceStreamLogsServer struct {
	grpc.ServerStream
}

func (x *clusterServiceStreamLogsServer) Send(m *StreamLogRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _ClusterService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListOperations(ctx, req.(*ListClusterOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ClusterService/ListHosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListHosts(ctx, req.(*ListClusterHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.kafka.v1.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ClusterService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ClusterService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ClusterService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClusterService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClusterService_Delete_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _ClusterService_Move_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _ClusterService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ClusterService_Stop_Handler,
		},
		{
			MethodName: "RescheduleMaintenance",
			Handler:    _ClusterService_RescheduleMaintenance_Handler,
		},
		{
			MethodName: "ListLogs",
			Handler:    _ClusterService_ListLogs_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _ClusterService_ListOperations_Handler,
		},
		{
			MethodName: "ListHosts",
			Handler:    _ClusterService_ListHosts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLogs",
			Handler:       _ClusterService_StreamLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "yandex/cloud/mdb/kafka/v1/cluster_service.proto",
}
