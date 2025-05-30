// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: imgservice.proto

package proto

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

type UploadImageRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*UploadImageRequest_Meta
	//	*UploadImageRequest_Chunk
	Data          isUploadImageRequest_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	mi := &file_imgservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imgservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_imgservice_proto_rawDescGZIP(), []int{0}
}

func (x *UploadImageRequest) GetData() isUploadImageRequest_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadImageRequest) GetMeta() *UploadMeta {
	if x != nil {
		if x, ok := x.Data.(*UploadImageRequest_Meta); ok {
			return x.Meta
		}
	}
	return nil
}

func (x *UploadImageRequest) GetChunk() []byte {
	if x != nil {
		if x, ok := x.Data.(*UploadImageRequest_Chunk); ok {
			return x.Chunk
		}
	}
	return nil
}

type isUploadImageRequest_Data interface {
	isUploadImageRequest_Data()
}

type UploadImageRequest_Meta struct {
	Meta *UploadMeta `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type UploadImageRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*UploadImageRequest_Meta) isUploadImageRequest_Data() {}

func (*UploadImageRequest_Chunk) isUploadImageRequest_Data() {}

type UploadMeta struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadMeta) Reset() {
	*x = UploadMeta{}
	mi := &file_imgservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMeta) ProtoMessage() {}

func (x *UploadMeta) ProtoReflect() protoreflect.Message {
	mi := &file_imgservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadMeta.ProtoReflect.Descriptor instead.
func (*UploadMeta) Descriptor() ([]byte, []int) {
	return file_imgservice_proto_rawDescGZIP(), []int{1}
}

func (x *UploadMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UploadImageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Info          *ImageInfo             `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	mi := &file_imgservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imgservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_imgservice_proto_rawDescGZIP(), []int{2}
}

func (x *UploadImageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UploadImageResponse) GetInfo() *ImageInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type ImageInfo struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Name           string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt      string                 `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastModifiedAt string                 `protobuf:"bytes,3,opt,name=last_modified_at,json=lastModifiedAt,proto3" json:"last_modified_at,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ImageInfo) Reset() {
	*x = ImageInfo{}
	mi := &file_imgservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfo) ProtoMessage() {}

func (x *ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_imgservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return file_imgservice_proto_rawDescGZIP(), []int{3}
}

func (x *ImageInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImageInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ImageInfo) GetLastModifiedAt() string {
	if x != nil {
		return x.LastModifiedAt
	}
	return ""
}

type DownloadImageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadImageRequest) Reset() {
	*x = DownloadImageRequest{}
	mi := &file_imgservice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadImageRequest) ProtoMessage() {}

func (x *DownloadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imgservice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadImageRequest.ProtoReflect.Descriptor instead.
func (*DownloadImageRequest) Descriptor() ([]byte, []int) {
	return file_imgservice_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadImageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DownloadImageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chunk         []byte                 `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadImageResponse) Reset() {
	*x = DownloadImageResponse{}
	mi := &file_imgservice_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadImageResponse) ProtoMessage() {}

func (x *DownloadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imgservice_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadImageResponse.ProtoReflect.Descriptor instead.
func (*DownloadImageResponse) Descriptor() ([]byte, []int) {
	return file_imgservice_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadImageResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type GetImagesListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetImagesListRequest) Reset() {
	*x = GetImagesListRequest{}
	mi := &file_imgservice_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetImagesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesListRequest) ProtoMessage() {}

func (x *GetImagesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imgservice_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesListRequest.ProtoReflect.Descriptor instead.
func (*GetImagesListRequest) Descriptor() ([]byte, []int) {
	return file_imgservice_proto_rawDescGZIP(), []int{6}
}

type GetImagesListStrResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageInfo     []string               `protobuf:"bytes,1,rep,name=image_info,json=imageInfo,proto3" json:"image_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetImagesListStrResponse) Reset() {
	*x = GetImagesListStrResponse{}
	mi := &file_imgservice_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetImagesListStrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesListStrResponse) ProtoMessage() {}

func (x *GetImagesListStrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imgservice_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesListStrResponse.ProtoReflect.Descriptor instead.
func (*GetImagesListStrResponse) Descriptor() ([]byte, []int) {
	return file_imgservice_proto_rawDescGZIP(), []int{7}
}

func (x *GetImagesListStrResponse) GetImageInfo() []string {
	if x != nil {
		return x.ImageInfo
	}
	return nil
}

var File_imgservice_proto protoreflect.FileDescriptor

const file_imgservice_proto_rawDesc = "" +
	"\n" +
	"\x10imgservice.proto\x12\x05proto\"]\n" +
	"\x12UploadImageRequest\x12'\n" +
	"\x04meta\x18\x01 \x01(\v2\x11.proto.UploadMetaH\x00R\x04meta\x12\x16\n" +
	"\x05chunk\x18\x02 \x01(\fH\x00R\x05chunkB\x06\n" +
	"\x04data\" \n" +
	"\n" +
	"UploadMeta\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"U\n" +
	"\x13UploadImageResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12$\n" +
	"\x04info\x18\x02 \x01(\v2\x10.proto.ImageInfoR\x04info\"h\n" +
	"\tImageInfo\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1d\n" +
	"\n" +
	"created_at\x18\x02 \x01(\tR\tcreatedAt\x12(\n" +
	"\x10last_modified_at\x18\x03 \x01(\tR\x0elastModifiedAt\"*\n" +
	"\x14DownloadImageRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"-\n" +
	"\x15DownloadImageResponse\x12\x14\n" +
	"\x05chunk\x18\x01 \x01(\fR\x05chunk\"\x16\n" +
	"\x14GetImagesListRequest\"9\n" +
	"\x18GetImagesListStrResponse\x12\x1d\n" +
	"\n" +
	"image_info\x18\x01 \x03(\tR\timageInfo2\xf9\x01\n" +
	"\fImageService\x12H\n" +
	"\vUploadImage\x12\x19.proto.UploadImageRequest\x1a\x1a.proto.UploadImageResponse\"\x00(\x01\x12N\n" +
	"\rDownloadImage\x12\x1b.proto.DownloadImageRequest\x1a\x1c.proto.DownloadImageResponse\"\x000\x01\x12O\n" +
	"\rGetImagesList\x12\x1b.proto.GetImagesListRequest\x1a\x1f.proto.GetImagesListStrResponse\"\x00B$Z\"github.com/emrzvv/tages-test/protob\x06proto3"

var (
	file_imgservice_proto_rawDescOnce sync.Once
	file_imgservice_proto_rawDescData []byte
)

func file_imgservice_proto_rawDescGZIP() []byte {
	file_imgservice_proto_rawDescOnce.Do(func() {
		file_imgservice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_imgservice_proto_rawDesc), len(file_imgservice_proto_rawDesc)))
	})
	return file_imgservice_proto_rawDescData
}

var file_imgservice_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_imgservice_proto_goTypes = []any{
	(*UploadImageRequest)(nil),       // 0: proto.UploadImageRequest
	(*UploadMeta)(nil),               // 1: proto.UploadMeta
	(*UploadImageResponse)(nil),      // 2: proto.UploadImageResponse
	(*ImageInfo)(nil),                // 3: proto.ImageInfo
	(*DownloadImageRequest)(nil),     // 4: proto.DownloadImageRequest
	(*DownloadImageResponse)(nil),    // 5: proto.DownloadImageResponse
	(*GetImagesListRequest)(nil),     // 6: proto.GetImagesListRequest
	(*GetImagesListStrResponse)(nil), // 7: proto.GetImagesListStrResponse
}
var file_imgservice_proto_depIdxs = []int32{
	1, // 0: proto.UploadImageRequest.meta:type_name -> proto.UploadMeta
	3, // 1: proto.UploadImageResponse.info:type_name -> proto.ImageInfo
	0, // 2: proto.ImageService.UploadImage:input_type -> proto.UploadImageRequest
	4, // 3: proto.ImageService.DownloadImage:input_type -> proto.DownloadImageRequest
	6, // 4: proto.ImageService.GetImagesList:input_type -> proto.GetImagesListRequest
	2, // 5: proto.ImageService.UploadImage:output_type -> proto.UploadImageResponse
	5, // 6: proto.ImageService.DownloadImage:output_type -> proto.DownloadImageResponse
	7, // 7: proto.ImageService.GetImagesList:output_type -> proto.GetImagesListStrResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_imgservice_proto_init() }
func file_imgservice_proto_init() {
	if File_imgservice_proto != nil {
		return
	}
	file_imgservice_proto_msgTypes[0].OneofWrappers = []any{
		(*UploadImageRequest_Meta)(nil),
		(*UploadImageRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_imgservice_proto_rawDesc), len(file_imgservice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imgservice_proto_goTypes,
		DependencyIndexes: file_imgservice_proto_depIdxs,
		MessageInfos:      file_imgservice_proto_msgTypes,
	}.Build()
	File_imgservice_proto = out.File
	file_imgservice_proto_goTypes = nil
	file_imgservice_proto_depIdxs = nil
}
