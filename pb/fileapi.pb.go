// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: fileapi.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_fileapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_fileapi_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FileInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type SaveFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *SaveFileRequest) Reset() {
	*x = SaveFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveFileRequest) ProtoMessage() {}

func (x *SaveFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveFileRequest.ProtoReflect.Descriptor instead.
func (*SaveFileRequest) Descriptor() ([]byte, []int) {
	return file_fileapi_proto_rawDescGZIP(), []int{1}
}

func (x *SaveFileRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type SaveFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SaveFileResponse) Reset() {
	*x = SaveFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveFileResponse) ProtoMessage() {}

func (x *SaveFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveFileResponse.ProtoReflect.Descriptor instead.
func (*SaveFileResponse) Descriptor() ([]byte, []int) {
	return file_fileapi_proto_rawDescGZIP(), []int{2}
}

func (x *SaveFileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListFilesRequest) Reset() {
	*x = ListFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesRequest) ProtoMessage() {}

func (x *ListFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesRequest.ProtoReflect.Descriptor instead.
func (*ListFilesRequest) Descriptor() ([]byte, []int) {
	return file_fileapi_proto_rawDescGZIP(), []int{3}
}

func (x *ListFilesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListFilesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*FileInfo `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ListFilesResponse) Reset() {
	*x = ListFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesResponse) ProtoMessage() {}

func (x *ListFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesResponse.ProtoReflect.Descriptor instead.
func (*ListFilesResponse) Descriptor() ([]byte, []int) {
	return file_fileapi_proto_rawDescGZIP(), []int{4}
}

func (x *ListFilesResponse) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_fileapi_proto protoreflect.FileDescriptor

var file_fileapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x53, 0x61, 0x76,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x22, 0x26, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x37, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x32, 0x7e, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_fileapi_proto_rawDescOnce sync.Once
	file_fileapi_proto_rawDescData = file_fileapi_proto_rawDesc
)

func file_fileapi_proto_rawDescGZIP() []byte {
	file_fileapi_proto_rawDescOnce.Do(func() {
		file_fileapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_fileapi_proto_rawDescData)
	})
	return file_fileapi_proto_rawDescData
}

var file_fileapi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_fileapi_proto_goTypes = []interface{}{
	(*FileInfo)(nil),              // 0: pb.FileInfo
	(*SaveFileRequest)(nil),       // 1: pb.SaveFileRequest
	(*SaveFileResponse)(nil),      // 2: pb.SaveFileResponse
	(*ListFilesRequest)(nil),      // 3: pb.ListFilesRequest
	(*ListFilesResponse)(nil),     // 4: pb.ListFilesResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_fileapi_proto_depIdxs = []int32{
	5, // 0: pb.FileInfo.createdAt:type_name -> google.protobuf.Timestamp
	5, // 1: pb.FileInfo.updatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: pb.ListFilesResponse.files:type_name -> pb.FileInfo
	1, // 3: pb.FileService.SaveFile:input_type -> pb.SaveFileRequest
	3, // 4: pb.FileService.ListFiles:input_type -> pb.ListFilesRequest
	2, // 5: pb.FileService.SaveFile:output_type -> pb.SaveFileResponse
	4, // 6: pb.FileService.ListFiles:output_type -> pb.ListFilesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_fileapi_proto_init() }
func file_fileapi_proto_init() {
	if File_fileapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fileapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_fileapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveFileRequest); i {
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
		file_fileapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveFileResponse); i {
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
		file_fileapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesRequest); i {
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
		file_fileapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesResponse); i {
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
			RawDescriptor: file_fileapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fileapi_proto_goTypes,
		DependencyIndexes: file_fileapi_proto_depIdxs,
		MessageInfos:      file_fileapi_proto_msgTypes,
	}.Build()
	File_fileapi_proto = out.File
	file_fileapi_proto_rawDesc = nil
	file_fileapi_proto_goTypes = nil
	file_fileapi_proto_depIdxs = nil
}
