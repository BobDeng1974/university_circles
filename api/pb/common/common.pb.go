// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package university_circles_srv_common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ReportReq struct {
	Type                 int64    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Pic                  string   `protobuf:"bytes,3,opt,name=pic,proto3" json:"pic,omitempty"`
	MsgId                string   `protobuf:"bytes,4,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	GoodsId              string   `protobuf:"bytes,5,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	Uid                  string   `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportReq) Reset()         { *m = ReportReq{} }
func (m *ReportReq) String() string { return proto.CompactTextString(m) }
func (*ReportReq) ProtoMessage()    {}
func (*ReportReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *ReportReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportReq.Unmarshal(m, b)
}
func (m *ReportReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportReq.Marshal(b, m, deterministic)
}
func (m *ReportReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportReq.Merge(m, src)
}
func (m *ReportReq) XXX_Size() int {
	return xxx_messageInfo_ReportReq.Size(m)
}
func (m *ReportReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReportReq proto.InternalMessageInfo

func (m *ReportReq) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReportReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ReportReq) GetPic() string {
	if m != nil {
		return m.Pic
	}
	return ""
}

func (m *ReportReq) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *ReportReq) GetGoodsId() string {
	if m != nil {
		return m.GoodsId
	}
	return ""
}

func (m *ReportReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type Response struct {
	// 响应是否成功
	Success int64 `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// 响应信息
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() int64 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ReportReq)(nil), "university_circles.srv.common.ReportReq")
	proto.RegisterType((*Response)(nil), "university_circles.srv.common.Response")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0x4d, 0x4b, 0x03, 0x31,
	0x14, 0x74, 0xdd, 0x76, 0x6d, 0x1f, 0x08, 0x12, 0x11, 0xa2, 0x20, 0x94, 0xbd, 0xb8, 0xa7, 0x08,
	0x0a, 0xfe, 0x01, 0x4f, 0xbd, 0xc6, 0xb3, 0x14, 0x4c, 0x1e, 0x4b, 0xc0, 0x7c, 0x98, 0x97, 0x5d,
	0xe8, 0xaf, 0xf0, 0x2f, 0x4b, 0x92, 0xae, 0x47, 0x3d, 0x65, 0xe6, 0xe5, 0x65, 0x32, 0x33, 0x70,
	0xad, 0xbc, 0xb5, 0xde, 0x3d, 0xd6, 0x43, 0x84, 0xe8, 0x93, 0x67, 0xf7, 0x93, 0x33, 0x33, 0x46,
	0x32, 0xe9, 0x78, 0x50, 0x26, 0xaa, 0x4f, 0x24, 0x41, 0x71, 0x16, 0x75, 0xa9, 0xff, 0x6e, 0x60,
	0x2b, 0x31, 0xf8, 0x98, 0x24, 0x7e, 0x31, 0x06, 0xab, 0x74, 0x0c, 0xc8, 0x9b, 0x5d, 0x33, 0xb4,
	0xb2, 0x60, 0xc6, 0xe1, 0x42, 0x79, 0x97, 0xd0, 0x25, 0x7e, 0xbe, 0x6b, 0x86, 0xad, 0x5c, 0x28,
	0xbb, 0x82, 0x36, 0x18, 0xc5, 0xdb, 0x32, 0xcd, 0x90, 0xdd, 0x40, 0x67, 0x69, 0x3c, 0x18, 0xcd,
	0x57, 0x65, 0xb8, 0xb6, 0x34, 0xee, 0x35, 0xbb, 0x85, 0xcd, 0xe8, 0xbd, 0xa6, 0x7c, 0xb1, 0xae,
	0x1a, 0x85, 0xef, 0x75, 0xd6, 0x98, 0x8c, 0xe6, 0x5d, 0xd5, 0x98, 0x8c, 0xee, 0x5f, 0x60, 0x23,
	0x91, 0x82, 0x77, 0x54, 0xfe, 0xa6, 0x49, 0x29, 0x24, 0x3a, 0x59, 0x5a, 0x68, 0x7e, 0x67, 0x69,
	0x3c, 0x39, 0xca, 0xf0, 0xc9, 0xc1, 0xe5, 0x6b, 0xc9, 0xf4, 0x86, 0x71, 0x36, 0x0a, 0xd9, 0x3b,
	0x74, 0x35, 0x19, 0x1b, 0xc4, 0x9f, 0x25, 0x88, 0xdf, 0x02, 0xee, 0x1e, 0xfe, 0xdd, 0xac, 0xce,
	0xfa, 0xb3, 0x8f, 0xae, 0xf4, 0xfb, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x67, 0x49, 0x91, 0x3e,
	0x76, 0x01, 0x00, 0x00,
}
