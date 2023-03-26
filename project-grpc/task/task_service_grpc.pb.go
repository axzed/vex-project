// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: task_service.proto

package task

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	TaskStages(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskStagesResponse, error)
	MemberProjectList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*MemberProjectResponse, error)
	TaskList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskListResponse, error)
	SaveTask(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskMessage, error)
	TaskSort(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskSortResponse, error)
	MyTaskList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*MyTaskListResponse, error)
	ReadTask(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskMessage, error)
	ListTaskMember(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskMemberList, error)
	TaskLog(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskLogList, error)
	TaskWorkTimeList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskWorkTimeResponse, error)
	SaveTaskWorkTime(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*SaveTaskWorkTimeResponse, error)
	SaveTaskFile(ctx context.Context, in *TaskFileReqMessage, opts ...grpc.CallOption) (*TaskFileResponse, error)
	TaskSources(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskSourceResponse, error)
	CreateComment(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*CreateCommentResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) TaskStages(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskStagesResponse, error) {
	out := new(TaskStagesResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/TaskStages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) MemberProjectList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*MemberProjectResponse, error) {
	out := new(MemberProjectResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/MemberProjectList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskListResponse, error) {
	out := new(TaskListResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/TaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) SaveTask(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskMessage, error) {
	out := new(TaskMessage)
	err := c.cc.Invoke(ctx, "/task.TaskService/SaveTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskSort(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskSortResponse, error) {
	out := new(TaskSortResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/TaskSort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) MyTaskList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*MyTaskListResponse, error) {
	out := new(MyTaskListResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/MyTaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ReadTask(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskMessage, error) {
	out := new(TaskMessage)
	err := c.cc.Invoke(ctx, "/task.TaskService/ReadTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListTaskMember(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskMemberList, error) {
	out := new(TaskMemberList)
	err := c.cc.Invoke(ctx, "/task.TaskService/ListTaskMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskLog(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskLogList, error) {
	out := new(TaskLogList)
	err := c.cc.Invoke(ctx, "/task.TaskService/TaskLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskWorkTimeList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskWorkTimeResponse, error) {
	out := new(TaskWorkTimeResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/TaskWorkTimeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) SaveTaskWorkTime(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*SaveTaskWorkTimeResponse, error) {
	out := new(SaveTaskWorkTimeResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/SaveTaskWorkTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) SaveTaskFile(ctx context.Context, in *TaskFileReqMessage, opts ...grpc.CallOption) (*TaskFileResponse, error) {
	out := new(TaskFileResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/SaveTaskFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskSources(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskSourceResponse, error) {
	out := new(TaskSourceResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/TaskSources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) CreateComment(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	TaskStages(context.Context, *TaskReqMessage) (*TaskStagesResponse, error)
	MemberProjectList(context.Context, *TaskReqMessage) (*MemberProjectResponse, error)
	TaskList(context.Context, *TaskReqMessage) (*TaskListResponse, error)
	SaveTask(context.Context, *TaskReqMessage) (*TaskMessage, error)
	TaskSort(context.Context, *TaskReqMessage) (*TaskSortResponse, error)
	MyTaskList(context.Context, *TaskReqMessage) (*MyTaskListResponse, error)
	ReadTask(context.Context, *TaskReqMessage) (*TaskMessage, error)
	ListTaskMember(context.Context, *TaskReqMessage) (*TaskMemberList, error)
	TaskLog(context.Context, *TaskReqMessage) (*TaskLogList, error)
	TaskWorkTimeList(context.Context, *TaskReqMessage) (*TaskWorkTimeResponse, error)
	SaveTaskWorkTime(context.Context, *TaskReqMessage) (*SaveTaskWorkTimeResponse, error)
	SaveTaskFile(context.Context, *TaskFileReqMessage) (*TaskFileResponse, error)
	TaskSources(context.Context, *TaskReqMessage) (*TaskSourceResponse, error)
	CreateComment(context.Context, *TaskReqMessage) (*CreateCommentResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) TaskStages(context.Context, *TaskReqMessage) (*TaskStagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskStages not implemented")
}
func (UnimplementedTaskServiceServer) MemberProjectList(context.Context, *TaskReqMessage) (*MemberProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberProjectList not implemented")
}
func (UnimplementedTaskServiceServer) TaskList(context.Context, *TaskReqMessage) (*TaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskList not implemented")
}
func (UnimplementedTaskServiceServer) SaveTask(context.Context, *TaskReqMessage) (*TaskMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTask not implemented")
}
func (UnimplementedTaskServiceServer) TaskSort(context.Context, *TaskReqMessage) (*TaskSortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskSort not implemented")
}
func (UnimplementedTaskServiceServer) MyTaskList(context.Context, *TaskReqMessage) (*MyTaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MyTaskList not implemented")
}
func (UnimplementedTaskServiceServer) ReadTask(context.Context, *TaskReqMessage) (*TaskMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTask not implemented")
}
func (UnimplementedTaskServiceServer) ListTaskMember(context.Context, *TaskReqMessage) (*TaskMemberList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTaskMember not implemented")
}
func (UnimplementedTaskServiceServer) TaskLog(context.Context, *TaskReqMessage) (*TaskLogList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskLog not implemented")
}
func (UnimplementedTaskServiceServer) TaskWorkTimeList(context.Context, *TaskReqMessage) (*TaskWorkTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskWorkTimeList not implemented")
}
func (UnimplementedTaskServiceServer) SaveTaskWorkTime(context.Context, *TaskReqMessage) (*SaveTaskWorkTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTaskWorkTime not implemented")
}
func (UnimplementedTaskServiceServer) SaveTaskFile(context.Context, *TaskFileReqMessage) (*TaskFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTaskFile not implemented")
}
func (UnimplementedTaskServiceServer) TaskSources(context.Context, *TaskReqMessage) (*TaskSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskSources not implemented")
}
func (UnimplementedTaskServiceServer) CreateComment(context.Context, *TaskReqMessage) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_TaskStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/TaskStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskStages(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_MemberProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).MemberProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/MemberProjectList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).MemberProjectList(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/TaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskList(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_SaveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).SaveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/SaveTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).SaveTask(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskSort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskSort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/TaskSort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskSort(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_MyTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).MyTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/MyTaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).MyTaskList(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ReadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ReadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/ReadTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ReadTask(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListTaskMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ListTaskMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/ListTaskMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ListTaskMember(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/TaskLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskLog(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskWorkTimeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskWorkTimeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/TaskWorkTimeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskWorkTimeList(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_SaveTaskWorkTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).SaveTaskWorkTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/SaveTaskWorkTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).SaveTaskWorkTime(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_SaveTaskFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskFileReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).SaveTaskFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/SaveTaskFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).SaveTaskFile(ctx, req.(*TaskFileReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/TaskSources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskSources(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateComment(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskStages",
			Handler:    _TaskService_TaskStages_Handler,
		},
		{
			MethodName: "MemberProjectList",
			Handler:    _TaskService_MemberProjectList_Handler,
		},
		{
			MethodName: "TaskList",
			Handler:    _TaskService_TaskList_Handler,
		},
		{
			MethodName: "SaveTask",
			Handler:    _TaskService_SaveTask_Handler,
		},
		{
			MethodName: "TaskSort",
			Handler:    _TaskService_TaskSort_Handler,
		},
		{
			MethodName: "MyTaskList",
			Handler:    _TaskService_MyTaskList_Handler,
		},
		{
			MethodName: "ReadTask",
			Handler:    _TaskService_ReadTask_Handler,
		},
		{
			MethodName: "ListTaskMember",
			Handler:    _TaskService_ListTaskMember_Handler,
		},
		{
			MethodName: "TaskLog",
			Handler:    _TaskService_TaskLog_Handler,
		},
		{
			MethodName: "TaskWorkTimeList",
			Handler:    _TaskService_TaskWorkTimeList_Handler,
		},
		{
			MethodName: "SaveTaskWorkTime",
			Handler:    _TaskService_SaveTaskWorkTime_Handler,
		},
		{
			MethodName: "SaveTaskFile",
			Handler:    _TaskService_SaveTaskFile_Handler,
		},
		{
			MethodName: "TaskSources",
			Handler:    _TaskService_TaskSources_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _TaskService_CreateComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task_service.proto",
}
