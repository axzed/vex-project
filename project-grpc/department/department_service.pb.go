// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: department_service.proto

package department

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

type DepartmentReqMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId             int64  `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	OrganizationCode     string `protobuf:"bytes,2,opt,name=organizationCode,proto3" json:"organizationCode,omitempty"`
	Page                 int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int64  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	DepartmentCode       string `protobuf:"bytes,5,opt,name=departmentCode,proto3" json:"departmentCode,omitempty"`
	ParentDepartmentCode string `protobuf:"bytes,6,opt,name=parentDepartmentCode,proto3" json:"parentDepartmentCode,omitempty"`
	Name                 string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DepartmentReqMessage) Reset() {
	*x = DepartmentReqMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_department_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartmentReqMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartmentReqMessage) ProtoMessage() {}

func (x *DepartmentReqMessage) ProtoReflect() protoreflect.Message {
	mi := &file_department_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartmentReqMessage.ProtoReflect.Descriptor instead.
func (*DepartmentReqMessage) Descriptor() ([]byte, []int) {
	return file_department_service_proto_rawDescGZIP(), []int{0}
}

func (x *DepartmentReqMessage) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *DepartmentReqMessage) GetOrganizationCode() string {
	if x != nil {
		return x.OrganizationCode
	}
	return ""
}

func (x *DepartmentReqMessage) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *DepartmentReqMessage) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DepartmentReqMessage) GetDepartmentCode() string {
	if x != nil {
		return x.DepartmentCode
	}
	return ""
}

func (x *DepartmentReqMessage) GetParentDepartmentCode() string {
	if x != nil {
		return x.ParentDepartmentCode
	}
	return ""
}

func (x *DepartmentReqMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DepartmentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	OrganizationCode string `protobuf:"bytes,3,opt,name=OrganizationCode,proto3" json:"OrganizationCode,omitempty"`
	Name             string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	CreateTime       string `protobuf:"bytes,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Pcode            string `protobuf:"bytes,6,opt,name=pcode,proto3" json:"pcode,omitempty"`
	Path             string `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DepartmentMessage) Reset() {
	*x = DepartmentMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_department_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartmentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartmentMessage) ProtoMessage() {}

func (x *DepartmentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_department_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartmentMessage.ProtoReflect.Descriptor instead.
func (*DepartmentMessage) Descriptor() ([]byte, []int) {
	return file_department_service_proto_rawDescGZIP(), []int{1}
}

func (x *DepartmentMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DepartmentMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DepartmentMessage) GetOrganizationCode() string {
	if x != nil {
		return x.OrganizationCode
	}
	return ""
}

func (x *DepartmentMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DepartmentMessage) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *DepartmentMessage) GetPcode() string {
	if x != nil {
		return x.Pcode
	}
	return ""
}

func (x *DepartmentMessage) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ListDepartmentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*DepartmentMessage `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64                `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListDepartmentMessage) Reset() {
	*x = ListDepartmentMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_department_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDepartmentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDepartmentMessage) ProtoMessage() {}

func (x *ListDepartmentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_department_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDepartmentMessage.ProtoReflect.Descriptor instead.
func (*ListDepartmentMessage) Descriptor() ([]byte, []int) {
	return file_department_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListDepartmentMessage) GetList() []*DepartmentMessage {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListDepartmentMessage) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_department_service_proto protoreflect.FileDescriptor

var file_department_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xfe, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x32, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc1, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x2a, 0x0a, 0x10, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x60, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xf8, 0x01,
	0x0a, 0x11, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x20, 0x2e, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1d, 0x2e,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x20, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1d, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x20, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_department_service_proto_rawDescOnce sync.Once
	file_department_service_proto_rawDescData = file_department_service_proto_rawDesc
)

func file_department_service_proto_rawDescGZIP() []byte {
	file_department_service_proto_rawDescOnce.Do(func() {
		file_department_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_department_service_proto_rawDescData)
	})
	return file_department_service_proto_rawDescData
}

var file_department_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_department_service_proto_goTypes = []interface{}{
	(*DepartmentReqMessage)(nil),  // 0: department.DepartmentReqMessage
	(*DepartmentMessage)(nil),     // 1: department.DepartmentMessage
	(*ListDepartmentMessage)(nil), // 2: department.ListDepartmentMessage
}
var file_department_service_proto_depIdxs = []int32{
	1, // 0: department.ListDepartmentMessage.list:type_name -> department.DepartmentMessage
	0, // 1: department.DepartmentService.Save:input_type -> department.DepartmentReqMessage
	0, // 2: department.DepartmentService.Read:input_type -> department.DepartmentReqMessage
	0, // 3: department.DepartmentService.List:input_type -> department.DepartmentReqMessage
	1, // 4: department.DepartmentService.Save:output_type -> department.DepartmentMessage
	1, // 5: department.DepartmentService.Read:output_type -> department.DepartmentMessage
	2, // 6: department.DepartmentService.List:output_type -> department.ListDepartmentMessage
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_department_service_proto_init() }
func file_department_service_proto_init() {
	if File_department_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_department_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartmentReqMessage); i {
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
		file_department_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartmentMessage); i {
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
		file_department_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDepartmentMessage); i {
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
			RawDescriptor: file_department_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_department_service_proto_goTypes,
		DependencyIndexes: file_department_service_proto_depIdxs,
		MessageInfos:      file_department_service_proto_msgTypes,
	}.Build()
	File_department_service_proto = out.File
	file_department_service_proto_rawDesc = nil
	file_department_service_proto_goTypes = nil
	file_department_service_proto_depIdxs = nil
}
