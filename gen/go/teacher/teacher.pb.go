// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: teacher/teacher.proto

package teacher17

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

type TeacherPhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *TeacherPhoneNumber) Reset() {
	*x = TeacherPhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_teacher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherPhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherPhoneNumber) ProtoMessage() {}

func (x *TeacherPhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_teacher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherPhoneNumber.ProtoReflect.Descriptor instead.
func (*TeacherPhoneNumber) Descriptor() ([]byte, []int) {
	return file_teacher_teacher_proto_rawDescGZIP(), []int{0}
}

func (x *TeacherPhoneNumber) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type CreateTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherFname    string                `protobuf:"bytes,1,opt,name=teacher_fname,json=teacherFname,proto3" json:"teacher_fname,omitempty"`
	TeacherLname    string                `protobuf:"bytes,2,opt,name=teacher_lname,json=teacherLname,proto3" json:"teacher_lname,omitempty"`
	TeacherBirthday string                `protobuf:"bytes,3,opt,name=teacher_birthday,json=teacherBirthday,proto3" json:"teacher_birthday,omitempty"`
	Phones          []*TeacherPhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
}

func (x *CreateTeacherRequest) Reset() {
	*x = CreateTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_teacher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeacherRequest) ProtoMessage() {}

func (x *CreateTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_teacher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeacherRequest.ProtoReflect.Descriptor instead.
func (*CreateTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_teacher_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTeacherRequest) GetTeacherFname() string {
	if x != nil {
		return x.TeacherFname
	}
	return ""
}

func (x *CreateTeacherRequest) GetTeacherLname() string {
	if x != nil {
		return x.TeacherLname
	}
	return ""
}

func (x *CreateTeacherRequest) GetTeacherBirthday() string {
	if x != nil {
		return x.TeacherBirthday
	}
	return ""
}

func (x *CreateTeacherRequest) GetPhones() []*TeacherPhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

type CreateTeacherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId string `protobuf:"bytes,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *CreateTeacherResponse) Reset() {
	*x = CreateTeacherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_teacher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeacherResponse) ProtoMessage() {}

func (x *CreateTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_teacher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeacherResponse.ProtoReflect.Descriptor instead.
func (*CreateTeacherResponse) Descriptor() ([]byte, []int) {
	return file_teacher_teacher_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTeacherResponse) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

type Teacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherFname    string                `protobuf:"bytes,1,opt,name=teacher_fname,json=teacherFname,proto3" json:"teacher_fname,omitempty"`
	TeacherLname    string                `protobuf:"bytes,2,opt,name=teacher_lname,json=teacherLname,proto3" json:"teacher_lname,omitempty"`
	TeacherBirthday string                `protobuf:"bytes,3,opt,name=teacher_birthday,json=teacherBirthday,proto3" json:"teacher_birthday,omitempty"`
	Phones          []*TeacherPhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	Salary          float32               `protobuf:"fixed32,5,opt,name=salary,proto3" json:"salary,omitempty"`
	TeacherPoint    int64                 `protobuf:"varint,6,opt,name=teacher_point,json=teacherPoint,proto3" json:"teacher_point,omitempty"`
	Rating          int32                 `protobuf:"varint,7,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt       string                `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string                `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       string                `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	TeacherId       string                `protobuf:"bytes,11,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_teacher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_teacher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_teacher_teacher_proto_rawDescGZIP(), []int{3}
}

func (x *Teacher) GetTeacherFname() string {
	if x != nil {
		return x.TeacherFname
	}
	return ""
}

func (x *Teacher) GetTeacherLname() string {
	if x != nil {
		return x.TeacherLname
	}
	return ""
}

func (x *Teacher) GetTeacherBirthday() string {
	if x != nil {
		return x.TeacherBirthday
	}
	return ""
}

func (x *Teacher) GetPhones() []*TeacherPhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *Teacher) GetSalary() float32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *Teacher) GetTeacherPoint() int64 {
	if x != nil {
		return x.TeacherPoint
	}
	return 0
}

func (x *Teacher) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Teacher) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Teacher) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Teacher) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Teacher) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

type GetAllTeachersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Field  string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	Values string `protobuf:"bytes,4,opt,name=values,proto3" json:"values,omitempty"`
}

func (x *GetAllTeachersRequest) Reset() {
	*x = GetAllTeachersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_teacher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTeachersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTeachersRequest) ProtoMessage() {}

func (x *GetAllTeachersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_teacher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTeachersRequest.ProtoReflect.Descriptor instead.
func (*GetAllTeachersRequest) Descriptor() ([]byte, []int) {
	return file_teacher_teacher_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllTeachersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllTeachersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllTeachersRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *GetAllTeachersRequest) GetValues() string {
	if x != nil {
		return x.Values
	}
	return ""
}

type GetAllTeachersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teachers []*Teacher `protobuf:"bytes,1,rep,name=teachers,proto3" json:"teachers,omitempty"`
}

func (x *GetAllTeachersResponse) Reset() {
	*x = GetAllTeachersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_teacher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTeachersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTeachersResponse) ProtoMessage() {}

func (x *GetAllTeachersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_teacher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTeachersResponse.ProtoReflect.Descriptor instead.
func (*GetAllTeachersResponse) Descriptor() ([]byte, []int) {
	return file_teacher_teacher_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllTeachersResponse) GetTeachers() []*Teacher {
	if x != nil {
		return x.Teachers
	}
	return nil
}

type GetTeacherByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId string `protobuf:"bytes,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
}

func (x *GetTeacherByIdRequest) Reset() {
	*x = GetTeacherByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_teacher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeacherByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeacherByIdRequest) ProtoMessage() {}

func (x *GetTeacherByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_teacher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeacherByIdRequest.ProtoReflect.Descriptor instead.
func (*GetTeacherByIdRequest) Descriptor() ([]byte, []int) {
	return file_teacher_teacher_proto_rawDescGZIP(), []int{6}
}

func (x *GetTeacherByIdRequest) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

type UpdateTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherFname string                `protobuf:"bytes,1,opt,name=teacher_fname,json=teacherFname,proto3" json:"teacher_fname,omitempty"`
	TeacherLname string                `protobuf:"bytes,2,opt,name=teacher_lname,json=teacherLname,proto3" json:"teacher_lname,omitempty"`
	Phones       []*TeacherPhoneNumber `protobuf:"bytes,3,rep,name=phones,proto3" json:"phones,omitempty"`
}

func (x *UpdateTeacherRequest) Reset() {
	*x = UpdateTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teacher_teacher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeacherRequest) ProtoMessage() {}

func (x *UpdateTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_teacher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeacherRequest.ProtoReflect.Descriptor instead.
func (*UpdateTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_teacher_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTeacherRequest) GetTeacherFname() string {
	if x != nil {
		return x.TeacherFname
	}
	return ""
}

func (x *UpdateTeacherRequest) GetTeacherLname() string {
	if x != nil {
		return x.TeacherLname
	}
	return ""
}

func (x *UpdateTeacherRequest) GetPhones() []*TeacherPhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

var File_teacher_teacher_proto protoreflect.FileDescriptor

var file_teacher_teacher_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x12, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0xb8, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x46, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x2b,
	0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xfc, 0x02, 0x0a, 0x07, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x46,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x4c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x6f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x08, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x73, 0x22, 0x36, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x46, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x32, 0xf8, 0x01, 0x0a, 0x0e,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x16, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x42, 0x1c, 0x5a, 0x1a, 0x74, 0x75, 0x7a, 0x6f, 0x76, 0x2e,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x31, 0x37, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teacher_teacher_proto_rawDescOnce sync.Once
	file_teacher_teacher_proto_rawDescData = file_teacher_teacher_proto_rawDesc
)

func file_teacher_teacher_proto_rawDescGZIP() []byte {
	file_teacher_teacher_proto_rawDescOnce.Do(func() {
		file_teacher_teacher_proto_rawDescData = protoimpl.X.CompressGZIP(file_teacher_teacher_proto_rawDescData)
	})
	return file_teacher_teacher_proto_rawDescData
}

var file_teacher_teacher_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_teacher_teacher_proto_goTypes = []any{
	(*TeacherPhoneNumber)(nil),     // 0: TeacherPhoneNumber
	(*CreateTeacherRequest)(nil),   // 1: CreateTeacherRequest
	(*CreateTeacherResponse)(nil),  // 2: CreateTeacherResponse
	(*Teacher)(nil),                // 3: Teacher
	(*GetAllTeachersRequest)(nil),  // 4: GetAllTeachersRequest
	(*GetAllTeachersResponse)(nil), // 5: GetAllTeachersResponse
	(*GetTeacherByIdRequest)(nil),  // 6: GetTeacherByIdRequest
	(*UpdateTeacherRequest)(nil),   // 7: UpdateTeacherRequest
}
var file_teacher_teacher_proto_depIdxs = []int32{
	0, // 0: CreateTeacherRequest.phones:type_name -> TeacherPhoneNumber
	0, // 1: Teacher.phones:type_name -> TeacherPhoneNumber
	3, // 2: GetAllTeachersResponse.teachers:type_name -> Teacher
	0, // 3: UpdateTeacherRequest.phones:type_name -> TeacherPhoneNumber
	1, // 4: TeacherService.CreateTeacher:input_type -> CreateTeacherRequest
	4, // 5: TeacherService.GetAllTeacher:input_type -> GetAllTeachersRequest
	6, // 6: TeacherService.GetTeacherById:input_type -> GetTeacherByIdRequest
	7, // 7: TeacherService.UpdateTeacher:input_type -> UpdateTeacherRequest
	2, // 8: TeacherService.CreateTeacher:output_type -> CreateTeacherResponse
	5, // 9: TeacherService.GetAllTeacher:output_type -> GetAllTeachersResponse
	3, // 10: TeacherService.GetTeacherById:output_type -> Teacher
	3, // 11: TeacherService.UpdateTeacher:output_type -> Teacher
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_teacher_teacher_proto_init() }
func file_teacher_teacher_proto_init() {
	if File_teacher_teacher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teacher_teacher_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TeacherPhoneNumber); i {
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
		file_teacher_teacher_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTeacherRequest); i {
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
		file_teacher_teacher_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTeacherResponse); i {
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
		file_teacher_teacher_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Teacher); i {
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
		file_teacher_teacher_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTeachersRequest); i {
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
		file_teacher_teacher_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTeachersResponse); i {
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
		file_teacher_teacher_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetTeacherByIdRequest); i {
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
		file_teacher_teacher_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTeacherRequest); i {
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
			RawDescriptor: file_teacher_teacher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teacher_teacher_proto_goTypes,
		DependencyIndexes: file_teacher_teacher_proto_depIdxs,
		MessageInfos:      file_teacher_teacher_proto_msgTypes,
	}.Build()
	File_teacher_teacher_proto = out.File
	file_teacher_teacher_proto_rawDesc = nil
	file_teacher_teacher_proto_goTypes = nil
	file_teacher_teacher_proto_depIdxs = nil
}
