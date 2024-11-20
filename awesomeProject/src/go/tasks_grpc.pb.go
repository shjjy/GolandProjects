// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: tasks.proto

package tasks

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TaskService_CreateTask_FullMethodName  = "/tasks.TaskService/CreateTask"
	TaskService_GetTask_FullMethodName     = "/tasks.TaskService/GetTask"
	TaskService_ListTasks_FullMethodName   = "/tasks.TaskService/ListTasks"
	TaskService_RecordTasks_FullMethodName = "/tasks.TaskService/RecordTasks"
	TaskService_TaskChat_FullMethodName    = "/tasks.TaskService/TaskChat"
)

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error)
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Task], error)
	RecordTasks(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CreateTaskRequest, TaskSummary], error)
	TaskChat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[TaskComment, TaskComment], error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Task)
	err := c.cc.Invoke(ctx, TaskService_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Task)
	err := c.cc.Invoke(ctx, TaskService_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Task], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TaskService_ServiceDesc.Streams[0], TaskService_ListTasks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListTasksRequest, Task]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskService_ListTasksClient = grpc.ServerStreamingClient[Task]

func (c *taskServiceClient) RecordTasks(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CreateTaskRequest, TaskSummary], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TaskService_ServiceDesc.Streams[1], TaskService_RecordTasks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CreateTaskRequest, TaskSummary]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskService_RecordTasksClient = grpc.ClientStreamingClient[CreateTaskRequest, TaskSummary]

func (c *taskServiceClient) TaskChat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[TaskComment, TaskComment], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TaskService_ServiceDesc.Streams[2], TaskService_TaskChat_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TaskComment, TaskComment]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskService_TaskChatClient = grpc.BidiStreamingClient[TaskComment, TaskComment]

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility.
type TaskServiceServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*Task, error)
	GetTask(context.Context, *GetTaskRequest) (*Task, error)
	ListTasks(*ListTasksRequest, grpc.ServerStreamingServer[Task]) error
	RecordTasks(grpc.ClientStreamingServer[CreateTaskRequest, TaskSummary]) error
	TaskChat(grpc.BidiStreamingServer[TaskComment, TaskComment]) error
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskServiceServer struct{}

func (UnimplementedTaskServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskServiceServer) GetTask(context.Context, *GetTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskServiceServer) ListTasks(*ListTasksRequest, grpc.ServerStreamingServer[Task]) error {
	return status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedTaskServiceServer) RecordTasks(grpc.ClientStreamingServer[CreateTaskRequest, TaskSummary]) error {
	return status.Errorf(codes.Unimplemented, "method RecordTasks not implemented")
}
func (UnimplementedTaskServiceServer) TaskChat(grpc.BidiStreamingServer[TaskComment, TaskComment]) error {
	return status.Errorf(codes.Unimplemented, "method TaskChat not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}
func (UnimplementedTaskServiceServer) testEmbeddedByValue()                     {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	// If the following call pancis, it indicates UnimplementedTaskServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTasksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskServiceServer).ListTasks(m, &grpc.GenericServerStream[ListTasksRequest, Task]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskService_ListTasksServer = grpc.ServerStreamingServer[Task]

func _TaskService_RecordTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServiceServer).RecordTasks(&grpc.GenericServerStream[CreateTaskRequest, TaskSummary]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskService_RecordTasksServer = grpc.ClientStreamingServer[CreateTaskRequest, TaskSummary]

func _TaskService_TaskChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServiceServer).TaskChat(&grpc.GenericServerStream[TaskComment, TaskComment]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskService_TaskChatServer = grpc.BidiStreamingServer[TaskComment, TaskComment]

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tasks.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskService_GetTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTasks",
			Handler:       _TaskService_ListTasks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordTasks",
			Handler:       _TaskService_RecordTasks_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TaskChat",
			Handler:       _TaskService_TaskChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "tasks.proto",
}
