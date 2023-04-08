// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: project_service.proto

package project

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

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	Index(ctx context.Context, in *IndexMessage, opts ...grpc.CallOption) (*IndexResponse, error)
	FindProjectByMemId(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*MyProjectResponse, error)
	FindProjectTemplate(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*ProjectTemplateResponse, error)
	SaveProject(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*SaveProjectMessage, error)
	FindProjectDetail(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*ProjectDetailMessage, error)
	UpdateDeletedProject(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*DeletedProjectResponse, error)
	UpdateCollectProject(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*CollectProjectResponse, error)
	UpdateProject(ctx context.Context, in *UpdateProjectMessage, opts ...grpc.CallOption) (*UpdateProjectResponse, error)
	GetLogBySelfProject(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*ProjectLogResponse, error)
	NodeList(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*ProjectNodeResponseMessage, error)
	FindProjectByMemberId(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*FindProjectByMemberIdResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Index(ctx context.Context, in *IndexMessage, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) FindProjectByMemId(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*MyProjectResponse, error) {
	out := new(MyProjectResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/FindProjectByMemId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) FindProjectTemplate(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*ProjectTemplateResponse, error) {
	out := new(ProjectTemplateResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/FindProjectTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) SaveProject(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*SaveProjectMessage, error) {
	out := new(SaveProjectMessage)
	err := c.cc.Invoke(ctx, "/project.ProjectService/SaveProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) FindProjectDetail(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*ProjectDetailMessage, error) {
	out := new(ProjectDetailMessage)
	err := c.cc.Invoke(ctx, "/project.ProjectService/FindProjectDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateDeletedProject(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*DeletedProjectResponse, error) {
	out := new(DeletedProjectResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/UpdateDeletedProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateCollectProject(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*CollectProjectResponse, error) {
	out := new(CollectProjectResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/UpdateCollectProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProject(ctx context.Context, in *UpdateProjectMessage, opts ...grpc.CallOption) (*UpdateProjectResponse, error) {
	out := new(UpdateProjectResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetLogBySelfProject(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*ProjectLogResponse, error) {
	out := new(ProjectLogResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/GetLogBySelfProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) NodeList(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*ProjectNodeResponseMessage, error) {
	out := new(ProjectNodeResponseMessage)
	err := c.cc.Invoke(ctx, "/project.ProjectService/NodeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) FindProjectByMemberId(ctx context.Context, in *ProjectRpcMessage, opts ...grpc.CallOption) (*FindProjectByMemberIdResponse, error) {
	out := new(FindProjectByMemberIdResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/FindProjectByMemberId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	Index(context.Context, *IndexMessage) (*IndexResponse, error)
	FindProjectByMemId(context.Context, *ProjectRpcMessage) (*MyProjectResponse, error)
	FindProjectTemplate(context.Context, *ProjectRpcMessage) (*ProjectTemplateResponse, error)
	SaveProject(context.Context, *ProjectRpcMessage) (*SaveProjectMessage, error)
	FindProjectDetail(context.Context, *ProjectRpcMessage) (*ProjectDetailMessage, error)
	UpdateDeletedProject(context.Context, *ProjectRpcMessage) (*DeletedProjectResponse, error)
	UpdateCollectProject(context.Context, *ProjectRpcMessage) (*CollectProjectResponse, error)
	UpdateProject(context.Context, *UpdateProjectMessage) (*UpdateProjectResponse, error)
	GetLogBySelfProject(context.Context, *ProjectRpcMessage) (*ProjectLogResponse, error)
	NodeList(context.Context, *ProjectRpcMessage) (*ProjectNodeResponseMessage, error)
	FindProjectByMemberId(context.Context, *ProjectRpcMessage) (*FindProjectByMemberIdResponse, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) Index(context.Context, *IndexMessage) (*IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedProjectServiceServer) FindProjectByMemId(context.Context, *ProjectRpcMessage) (*MyProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProjectByMemId not implemented")
}
func (UnimplementedProjectServiceServer) FindProjectTemplate(context.Context, *ProjectRpcMessage) (*ProjectTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProjectTemplate not implemented")
}
func (UnimplementedProjectServiceServer) SaveProject(context.Context, *ProjectRpcMessage) (*SaveProjectMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProject not implemented")
}
func (UnimplementedProjectServiceServer) FindProjectDetail(context.Context, *ProjectRpcMessage) (*ProjectDetailMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProjectDetail not implemented")
}
func (UnimplementedProjectServiceServer) UpdateDeletedProject(context.Context, *ProjectRpcMessage) (*DeletedProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeletedProject not implemented")
}
func (UnimplementedProjectServiceServer) UpdateCollectProject(context.Context, *ProjectRpcMessage) (*CollectProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollectProject not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProject(context.Context, *UpdateProjectMessage) (*UpdateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedProjectServiceServer) GetLogBySelfProject(context.Context, *ProjectRpcMessage) (*ProjectLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogBySelfProject not implemented")
}
func (UnimplementedProjectServiceServer) NodeList(context.Context, *ProjectRpcMessage) (*ProjectNodeResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeList not implemented")
}
func (UnimplementedProjectServiceServer) FindProjectByMemberId(context.Context, *ProjectRpcMessage) (*FindProjectByMemberIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProjectByMemberId not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Index(ctx, req.(*IndexMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_FindProjectByMemId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRpcMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).FindProjectByMemId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/FindProjectByMemId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).FindProjectByMemId(ctx, req.(*ProjectRpcMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_FindProjectTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRpcMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).FindProjectTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/FindProjectTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).FindProjectTemplate(ctx, req.(*ProjectRpcMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_SaveProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRpcMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).SaveProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/SaveProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).SaveProject(ctx, req.(*ProjectRpcMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_FindProjectDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRpcMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).FindProjectDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/FindProjectDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).FindProjectDetail(ctx, req.(*ProjectRpcMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateDeletedProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRpcMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateDeletedProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/UpdateDeletedProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateDeletedProject(ctx, req.(*ProjectRpcMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateCollectProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRpcMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateCollectProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/UpdateCollectProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateCollectProject(ctx, req.(*ProjectRpcMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/UpdateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProject(ctx, req.(*UpdateProjectMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetLogBySelfProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRpcMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetLogBySelfProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/GetLogBySelfProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetLogBySelfProject(ctx, req.(*ProjectRpcMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_NodeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRpcMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).NodeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/NodeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).NodeList(ctx, req.(*ProjectRpcMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_FindProjectByMemberId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRpcMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).FindProjectByMemberId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/FindProjectByMemberId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).FindProjectByMemberId(ctx, req.(*ProjectRpcMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "project.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _ProjectService_Index_Handler,
		},
		{
			MethodName: "FindProjectByMemId",
			Handler:    _ProjectService_FindProjectByMemId_Handler,
		},
		{
			MethodName: "FindProjectTemplate",
			Handler:    _ProjectService_FindProjectTemplate_Handler,
		},
		{
			MethodName: "SaveProject",
			Handler:    _ProjectService_SaveProject_Handler,
		},
		{
			MethodName: "FindProjectDetail",
			Handler:    _ProjectService_FindProjectDetail_Handler,
		},
		{
			MethodName: "UpdateDeletedProject",
			Handler:    _ProjectService_UpdateDeletedProject_Handler,
		},
		{
			MethodName: "UpdateCollectProject",
			Handler:    _ProjectService_UpdateCollectProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _ProjectService_UpdateProject_Handler,
		},
		{
			MethodName: "GetLogBySelfProject",
			Handler:    _ProjectService_GetLogBySelfProject_Handler,
		},
		{
			MethodName: "NodeList",
			Handler:    _ProjectService_NodeList_Handler,
		},
		{
			MethodName: "FindProjectByMemberId",
			Handler:    _ProjectService_FindProjectByMemberId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project_service.proto",
}
