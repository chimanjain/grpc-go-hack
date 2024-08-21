// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: project/project.proto

package project

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProjectId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProjectId) Reset() {
	*x = ProjectId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectId) ProtoMessage() {}

func (x *ProjectId) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectId.ProtoReflect.Descriptor instead.
func (*ProjectId) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DeployId             string `protobuf:"bytes,3,opt,name=deployId,proto3" json:"deployId,omitempty"`
	DeployAt             string `protobuf:"bytes,4,opt,name=deployAt,proto3" json:"deployAt,omitempty"`
	NoticeTotalCount     uint32 `protobuf:"varint,5,opt,name=noticeTotalCount,proto3" json:"noticeTotalCount,omitempty"`
	RejectionCount       uint32 `protobuf:"varint,6,opt,name=rejectionCount,proto3" json:"rejectionCount,omitempty"`
	FileCount            uint32 `protobuf:"varint,7,opt,name=fileCount,proto3" json:"fileCount,omitempty"`
	DeployCount          uint32 `protobuf:"varint,8,opt,name=deployCount,proto3" json:"deployCount,omitempty"`
	GroupResolvedCount   uint32 `protobuf:"varint,9,opt,name=groupResolvedCount,proto3" json:"groupResolvedCount,omitempty"`
	GroupUnresolvedCount uint32 `protobuf:"varint,10,opt,name=groupUnresolvedCount,proto3" json:"groupUnresolvedCount,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{1}
}

func (x *Project) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetDeployId() string {
	if x != nil {
		return x.DeployId
	}
	return ""
}

func (x *Project) GetDeployAt() string {
	if x != nil {
		return x.DeployAt
	}
	return ""
}

func (x *Project) GetNoticeTotalCount() uint32 {
	if x != nil {
		return x.NoticeTotalCount
	}
	return 0
}

func (x *Project) GetRejectionCount() uint32 {
	if x != nil {
		return x.RejectionCount
	}
	return 0
}

func (x *Project) GetFileCount() uint32 {
	if x != nil {
		return x.FileCount
	}
	return 0
}

func (x *Project) GetDeployCount() uint32 {
	if x != nil {
		return x.DeployCount
	}
	return 0
}

func (x *Project) GetGroupResolvedCount() uint32 {
	if x != nil {
		return x.GroupResolvedCount
	}
	return 0
}

func (x *Project) GetGroupUnresolvedCount() uint32 {
	if x != nil {
		return x.GroupUnresolvedCount
	}
	return 0
}

type Projects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *Projects) Reset() {
	*x = Projects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Projects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Projects) ProtoMessage() {}

func (x *Projects) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Projects.ProtoReflect.Descriptor instead.
func (*Projects) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{2}
}

func (x *Projects) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

var File_project_project_proto protoreflect.FileDescriptor

var file_project_project_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x1b, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdd, 0x02,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x41, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x10, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x6e,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x38, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x32, 0xfc, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x0d,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3c, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x69, 0x6d, 0x61, 0x6e, 0x6a, 0x61, 0x69, 0x6e, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2d, 0x68, 0x61, 0x63, 0x6b, 0x3b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_project_proto_rawDescOnce sync.Once
	file_project_project_proto_rawDescData = file_project_project_proto_rawDesc
)

func file_project_project_proto_rawDescGZIP() []byte {
	file_project_project_proto_rawDescOnce.Do(func() {
		file_project_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_project_proto_rawDescData)
	})
	return file_project_project_proto_rawDescData
}

var file_project_project_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_project_project_proto_goTypes = []any{
	(*ProjectId)(nil), // 0: project.ProjectId
	(*Project)(nil),   // 1: project.Project
	(*Projects)(nil),  // 2: project.Projects
}
var file_project_project_proto_depIdxs = []int32{
	1, // 0: project.Projects.projects:type_name -> project.Project
	0, // 1: project.ProjectManagement.GetProject:input_type -> project.ProjectId
	0, // 2: project.ProjectManagement.GetProjects:input_type -> project.ProjectId
	0, // 3: project.ProjectManagement.FetchProjects:input_type -> project.ProjectId
	0, // 4: project.ProjectManagement.StreamProjects:input_type -> project.ProjectId
	1, // 5: project.ProjectManagement.GetProject:output_type -> project.Project
	1, // 6: project.ProjectManagement.GetProjects:output_type -> project.Project
	2, // 7: project.ProjectManagement.FetchProjects:output_type -> project.Projects
	1, // 8: project.ProjectManagement.StreamProjects:output_type -> project.Project
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_project_project_proto_init() }
func file_project_project_proto_init() {
	if File_project_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_project_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProjectId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_project_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Project); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_project_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Projects); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_project_proto_goTypes,
		DependencyIndexes: file_project_project_proto_depIdxs,
		MessageInfos:      file_project_project_proto_msgTypes,
	}.Build()
	File_project_project_proto = out.File
	file_project_project_proto_rawDesc = nil
	file_project_project_proto_goTypes = nil
	file_project_project_proto_depIdxs = nil
}
