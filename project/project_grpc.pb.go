// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: project/project.proto

package project

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ProjectManagement_GetProject_FullMethodName     = "/project.ProjectManagement/GetProject"
	ProjectManagement_GetProjects_FullMethodName    = "/project.ProjectManagement/GetProjects"
	ProjectManagement_FetchProjects_FullMethodName  = "/project.ProjectManagement/FetchProjects"
	ProjectManagement_StreamProjects_FullMethodName = "/project.ProjectManagement/StreamProjects"
)

// ProjectManagementClient is the client API for ProjectManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectManagementClient interface {
	GetProject(ctx context.Context, in *ProjectId, opts ...grpc.CallOption) (*Project, error)
	GetProjects(ctx context.Context, in *ProjectId, opts ...grpc.CallOption) (ProjectManagement_GetProjectsClient, error)
	FetchProjects(ctx context.Context, opts ...grpc.CallOption) (ProjectManagement_FetchProjectsClient, error)
	StreamProjects(ctx context.Context, opts ...grpc.CallOption) (ProjectManagement_StreamProjectsClient, error)
}

type projectManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectManagementClient(cc grpc.ClientConnInterface) ProjectManagementClient {
	return &projectManagementClient{cc}
}

func (c *projectManagementClient) GetProject(ctx context.Context, in *ProjectId, opts ...grpc.CallOption) (*Project, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Project)
	err := c.cc.Invoke(ctx, ProjectManagement_GetProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectManagementClient) GetProjects(ctx context.Context, in *ProjectId, opts ...grpc.CallOption) (ProjectManagement_GetProjectsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProjectManagement_ServiceDesc.Streams[0], ProjectManagement_GetProjects_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &projectManagementGetProjectsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectManagement_GetProjectsClient interface {
	Recv() (*Project, error)
	grpc.ClientStream
}

type projectManagementGetProjectsClient struct {
	grpc.ClientStream
}

func (x *projectManagementGetProjectsClient) Recv() (*Project, error) {
	m := new(Project)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectManagementClient) FetchProjects(ctx context.Context, opts ...grpc.CallOption) (ProjectManagement_FetchProjectsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProjectManagement_ServiceDesc.Streams[1], ProjectManagement_FetchProjects_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &projectManagementFetchProjectsClient{ClientStream: stream}
	return x, nil
}

type ProjectManagement_FetchProjectsClient interface {
	Send(*ProjectId) error
	CloseAndRecv() (*Projects, error)
	grpc.ClientStream
}

type projectManagementFetchProjectsClient struct {
	grpc.ClientStream
}

func (x *projectManagementFetchProjectsClient) Send(m *ProjectId) error {
	return x.ClientStream.SendMsg(m)
}

func (x *projectManagementFetchProjectsClient) CloseAndRecv() (*Projects, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Projects)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectManagementClient) StreamProjects(ctx context.Context, opts ...grpc.CallOption) (ProjectManagement_StreamProjectsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProjectManagement_ServiceDesc.Streams[2], ProjectManagement_StreamProjects_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &projectManagementStreamProjectsClient{ClientStream: stream}
	return x, nil
}

type ProjectManagement_StreamProjectsClient interface {
	Send(*ProjectId) error
	Recv() (*Project, error)
	grpc.ClientStream
}

type projectManagementStreamProjectsClient struct {
	grpc.ClientStream
}

func (x *projectManagementStreamProjectsClient) Send(m *ProjectId) error {
	return x.ClientStream.SendMsg(m)
}

func (x *projectManagementStreamProjectsClient) Recv() (*Project, error) {
	m := new(Project)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProjectManagementServer is the server API for ProjectManagement service.
// All implementations must embed UnimplementedProjectManagementServer
// for forward compatibility
type ProjectManagementServer interface {
	GetProject(context.Context, *ProjectId) (*Project, error)
	GetProjects(*ProjectId, ProjectManagement_GetProjectsServer) error
	FetchProjects(ProjectManagement_FetchProjectsServer) error
	StreamProjects(ProjectManagement_StreamProjectsServer) error
	mustEmbedUnimplementedProjectManagementServer()
}

// UnimplementedProjectManagementServer must be embedded to have forward compatible implementations.
type UnimplementedProjectManagementServer struct {
}

func (UnimplementedProjectManagementServer) GetProject(context.Context, *ProjectId) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectManagementServer) GetProjects(*ProjectId, ProjectManagement_GetProjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProjects not implemented")
}
func (UnimplementedProjectManagementServer) FetchProjects(ProjectManagement_FetchProjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchProjects not implemented")
}
func (UnimplementedProjectManagementServer) StreamProjects(ProjectManagement_StreamProjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamProjects not implemented")
}
func (UnimplementedProjectManagementServer) mustEmbedUnimplementedProjectManagementServer() {}

// UnsafeProjectManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectManagementServer will
// result in compilation errors.
type UnsafeProjectManagementServer interface {
	mustEmbedUnimplementedProjectManagementServer()
}

func RegisterProjectManagementServer(s grpc.ServiceRegistrar, srv ProjectManagementServer) {
	s.RegisterService(&ProjectManagement_ServiceDesc, srv)
}

func _ProjectManagement_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectManagementServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectManagement_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectManagementServer).GetProject(ctx, req.(*ProjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectManagement_GetProjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProjectId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectManagementServer).GetProjects(m, &projectManagementGetProjectsServer{ServerStream: stream})
}

type ProjectManagement_GetProjectsServer interface {
	Send(*Project) error
	grpc.ServerStream
}

type projectManagementGetProjectsServer struct {
	grpc.ServerStream
}

func (x *projectManagementGetProjectsServer) Send(m *Project) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectManagement_FetchProjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProjectManagementServer).FetchProjects(&projectManagementFetchProjectsServer{ServerStream: stream})
}

type ProjectManagement_FetchProjectsServer interface {
	SendAndClose(*Projects) error
	Recv() (*ProjectId, error)
	grpc.ServerStream
}

type projectManagementFetchProjectsServer struct {
	grpc.ServerStream
}

func (x *projectManagementFetchProjectsServer) SendAndClose(m *Projects) error {
	return x.ServerStream.SendMsg(m)
}

func (x *projectManagementFetchProjectsServer) Recv() (*ProjectId, error) {
	m := new(ProjectId)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProjectManagement_StreamProjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProjectManagementServer).StreamProjects(&projectManagementStreamProjectsServer{ServerStream: stream})
}

type ProjectManagement_StreamProjectsServer interface {
	Send(*Project) error
	Recv() (*ProjectId, error)
	grpc.ServerStream
}

type projectManagementStreamProjectsServer struct {
	grpc.ServerStream
}

func (x *projectManagementStreamProjectsServer) Send(m *Project) error {
	return x.ServerStream.SendMsg(m)
}

func (x *projectManagementStreamProjectsServer) Recv() (*ProjectId, error) {
	m := new(ProjectId)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProjectManagement_ServiceDesc is the grpc.ServiceDesc for ProjectManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "project.ProjectManagement",
	HandlerType: (*ProjectManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProject",
			Handler:    _ProjectManagement_GetProject_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetProjects",
			Handler:       _ProjectManagement_GetProjects_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FetchProjects",
			Handler:       _ProjectManagement_FetchProjects_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamProjects",
			Handler:       _ProjectManagement_StreamProjects_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "project/project.proto",
}
