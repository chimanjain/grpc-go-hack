// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: project/project.proto

package project

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProjectID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectID) Reset() {
	*x = ProjectID{}
	mi := &file_project_project_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectID) ProtoMessage() {}

func (x *ProjectID) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectID.ProtoReflect.Descriptor instead.
func (*ProjectID) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectID) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Project struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Id                   uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DeployId             string                 `protobuf:"bytes,3,opt,name=deployId,proto3" json:"deployId,omitempty"`
	DeployAt             string                 `protobuf:"bytes,4,opt,name=deployAt,proto3" json:"deployAt,omitempty"`
	NoticeTotalCount     uint32                 `protobuf:"varint,5,opt,name=noticeTotalCount,proto3" json:"noticeTotalCount,omitempty"`
	RejectionCount       uint32                 `protobuf:"varint,6,opt,name=rejectionCount,proto3" json:"rejectionCount,omitempty"`
	FileCount            uint32                 `protobuf:"varint,7,opt,name=fileCount,proto3" json:"fileCount,omitempty"`
	DeployCount          uint32                 `protobuf:"varint,8,opt,name=deployCount,proto3" json:"deployCount,omitempty"`
	GroupResolvedCount   uint32                 `protobuf:"varint,9,opt,name=groupResolvedCount,proto3" json:"groupResolvedCount,omitempty"`
	GroupUnresolvedCount uint32                 `protobuf:"varint,10,opt,name=groupUnresolvedCount,proto3" json:"groupUnresolvedCount,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Project) Reset() {
	*x = Project{}
	mi := &file_project_project_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[1]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Projects      []*Project             `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Projects) Reset() {
	*x = Projects{}
	mi := &file_project_project_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Projects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Projects) ProtoMessage() {}

func (x *Projects) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[2]
	if x != nil {
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

const file_project_project_proto_rawDesc = "" +
	"\n" +
	"\x15project/project.proto\x12\aproject\"\x1b\n" +
	"\tProjectID\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"\xdd\x02\n" +
	"\aProject\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bdeployId\x18\x03 \x01(\tR\bdeployId\x12\x1a\n" +
	"\bdeployAt\x18\x04 \x01(\tR\bdeployAt\x12*\n" +
	"\x10noticeTotalCount\x18\x05 \x01(\rR\x10noticeTotalCount\x12&\n" +
	"\x0erejectionCount\x18\x06 \x01(\rR\x0erejectionCount\x12\x1c\n" +
	"\tfileCount\x18\a \x01(\rR\tfileCount\x12 \n" +
	"\vdeployCount\x18\b \x01(\rR\vdeployCount\x12.\n" +
	"\x12groupResolvedCount\x18\t \x01(\rR\x12groupResolvedCount\x122\n" +
	"\x14groupUnresolvedCount\x18\n" +
	" \x01(\rR\x14groupUnresolvedCount\"8\n" +
	"\bProjects\x12,\n" +
	"\bprojects\x18\x01 \x03(\v2\x10.project.ProjectR\bprojects2\xfc\x01\n" +
	"\x11ProjectManagement\x124\n" +
	"\n" +
	"GetProject\x12\x12.project.ProjectID\x1a\x10.project.Project\"\x00\x127\n" +
	"\vGetProjects\x12\x12.project.ProjectID\x1a\x10.project.Project\"\x000\x01\x12:\n" +
	"\rFetchProjects\x12\x12.project.ProjectID\x1a\x11.project.Projects\"\x00(\x01\x12<\n" +
	"\x0eStreamProjects\x12\x12.project.ProjectID\x1a\x10.project.Project\"\x00(\x010\x01B,Z*github.com/chimanjain/grpc-go-hack;projectb\x06proto3"

var (
	file_project_project_proto_rawDescOnce sync.Once
	file_project_project_proto_rawDescData []byte
)

func file_project_project_proto_rawDescGZIP() []byte {
	file_project_project_proto_rawDescOnce.Do(func() {
		file_project_project_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_project_proto_rawDesc), len(file_project_project_proto_rawDesc)))
	})
	return file_project_project_proto_rawDescData
}

var file_project_project_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_project_project_proto_goTypes = []any{
	(*ProjectID)(nil), // 0: project.ProjectID
	(*Project)(nil),   // 1: project.Project
	(*Projects)(nil),  // 2: project.Projects
}
var file_project_project_proto_depIdxs = []int32{
	1, // 0: project.Projects.projects:type_name -> project.Project
	0, // 1: project.ProjectManagement.GetProject:input_type -> project.ProjectID
	0, // 2: project.ProjectManagement.GetProjects:input_type -> project.ProjectID
	0, // 3: project.ProjectManagement.FetchProjects:input_type -> project.ProjectID
	0, // 4: project.ProjectManagement.StreamProjects:input_type -> project.ProjectID
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_project_proto_rawDesc), len(file_project_project_proto_rawDesc)),
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
	file_project_project_proto_goTypes = nil
	file_project_project_proto_depIdxs = nil
}
