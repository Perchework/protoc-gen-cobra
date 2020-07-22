// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crud.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CRUDObject struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CRUDObject) Reset()         { *m = CRUDObject{} }
func (m *CRUDObject) String() string { return proto.CompactTextString(m) }
func (*CRUDObject) ProtoMessage()    {}
func (*CRUDObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_478bbe1b22b2e995, []int{0}
}
func (m *CRUDObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CRUDObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CRUDObject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CRUDObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CRUDObject.Merge(m, src)
}
func (m *CRUDObject) XXX_Size() int {
	return m.Size()
}
func (m *CRUDObject) XXX_DiscardUnknown() {
	xxx_messageInfo_CRUDObject.DiscardUnknown(m)
}

var xxx_messageInfo_CRUDObject proto.InternalMessageInfo

func (m *CRUDObject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CRUDObject) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CreateCRUD struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCRUD) Reset()         { *m = CreateCRUD{} }
func (m *CreateCRUD) String() string { return proto.CompactTextString(m) }
func (*CreateCRUD) ProtoMessage()    {}
func (*CreateCRUD) Descriptor() ([]byte, []int) {
	return fileDescriptor_478bbe1b22b2e995, []int{1}
}
func (m *CreateCRUD) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateCRUD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateCRUD.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateCRUD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCRUD.Merge(m, src)
}
func (m *CreateCRUD) XXX_Size() int {
	return m.Size()
}
func (m *CreateCRUD) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCRUD.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCRUD proto.InternalMessageInfo

func (m *CreateCRUD) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCRUD) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetCRUD struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCRUD) Reset()         { *m = GetCRUD{} }
func (m *GetCRUD) String() string { return proto.CompactTextString(m) }
func (*GetCRUD) ProtoMessage()    {}
func (*GetCRUD) Descriptor() ([]byte, []int) {
	return fileDescriptor_478bbe1b22b2e995, []int{2}
}
func (m *GetCRUD) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCRUD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCRUD.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetCRUD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCRUD.Merge(m, src)
}
func (m *GetCRUD) XXX_Size() int {
	return m.Size()
}
func (m *GetCRUD) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCRUD.DiscardUnknown(m)
}

var xxx_messageInfo_GetCRUD proto.InternalMessageInfo

func (m *GetCRUD) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_478bbe1b22b2e995, []int{3}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CRUDObject)(nil), "pb.CRUDObject")
	proto.RegisterType((*CreateCRUD)(nil), "pb.CreateCRUD")
	proto.RegisterType((*GetCRUD)(nil), "pb.GetCRUD")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

func init() { proto.RegisterFile("crud.proto", fileDescriptor_478bbe1b22b2e995) }

var fileDescriptor_478bbe1b22b2e995 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2e, 0x2a, 0x4d,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe3, 0xe2, 0x72, 0x0e,
	0x0a, 0x75, 0xf1, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73,
	0x4a, 0x53, 0x25, 0x98, 0xc0, 0x82, 0x10, 0x0e, 0x58, 0x5f, 0x51, 0x6a, 0x62, 0x49, 0x2a, 0x48,
	0x37, 0x09, 0xfa, 0x64, 0xb9, 0xd8, 0xdd, 0x53, 0x4b, 0x70, 0x69, 0x52, 0x62, 0xe7, 0x62, 0x75,
	0xcd, 0x2d, 0x28, 0xa9, 0x34, 0x5a, 0xcc, 0xc8, 0xc5, 0x02, 0x56, 0xa5, 0xc1, 0xc5, 0x06, 0xb1,
	0x48, 0x88, 0x4f, 0xaf, 0x20, 0x49, 0x0f, 0x61, 0xa9, 0x14, 0x84, 0x8f, 0x70, 0xbc, 0x12, 0x17,
	0xb3, 0x7b, 0x6a, 0x89, 0x10, 0x37, 0x48, 0x18, 0x6a, 0x07, 0x86, 0x1a, 0x0d, 0x2e, 0xb6, 0xd0,
	0x82, 0x14, 0x84, 0x69, 0x70, 0x19, 0x0c, 0x95, 0xca, 0x5c, 0x6c, 0x2e, 0xa9, 0x39, 0xa9, 0x58,
	0x54, 0x72, 0x82, 0xf8, 0x60, 0x57, 0x3a, 0x09, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0x83, 0xd6, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xcf, 0xb9, 0x4b, 0xa9, 0x68, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CRUDClient is the client API for CRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CRUDClient interface {
	Create(ctx context.Context, in *CreateCRUD, opts ...grpc.CallOption) (*CRUDObject, error)
	Get(ctx context.Context, in *GetCRUD, opts ...grpc.CallOption) (*CRUDObject, error)
	Update(ctx context.Context, in *CRUDObject, opts ...grpc.CallOption) (*CRUDObject, error)
	Delete(ctx context.Context, in *CRUDObject, opts ...grpc.CallOption) (*Empty, error)
}

type cRUDClient struct {
	cc *grpc.ClientConn
}

func NewCRUDClient(cc *grpc.ClientConn) CRUDClient {
	return &cRUDClient{cc}
}

func (c *cRUDClient) Create(ctx context.Context, in *CreateCRUD, opts ...grpc.CallOption) (*CRUDObject, error) {
	out := new(CRUDObject)
	err := c.cc.Invoke(ctx, "/pb.CRUD/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Get(ctx context.Context, in *GetCRUD, opts ...grpc.CallOption) (*CRUDObject, error) {
	out := new(CRUDObject)
	err := c.cc.Invoke(ctx, "/pb.CRUD/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Update(ctx context.Context, in *CRUDObject, opts ...grpc.CallOption) (*CRUDObject, error) {
	out := new(CRUDObject)
	err := c.cc.Invoke(ctx, "/pb.CRUD/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Delete(ctx context.Context, in *CRUDObject, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.CRUD/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServer is the server API for CRUD service.
type CRUDServer interface {
	Create(context.Context, *CreateCRUD) (*CRUDObject, error)
	Get(context.Context, *GetCRUD) (*CRUDObject, error)
	Update(context.Context, *CRUDObject) (*CRUDObject, error)
	Delete(context.Context, *CRUDObject) (*Empty, error)
}

// UnimplementedCRUDServer can be embedded to have forward compatible implementations.
type UnimplementedCRUDServer struct {
}

func (*UnimplementedCRUDServer) Create(ctx context.Context, req *CreateCRUD) (*CRUDObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCRUDServer) Get(ctx context.Context, req *GetCRUD) (*CRUDObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCRUDServer) Update(ctx context.Context, req *CRUDObject) (*CRUDObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCRUDServer) Delete(ctx context.Context, req *CRUDObject) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCRUDServer(s *grpc.Server, srv CRUDServer) {
	s.RegisterService(&_CRUD_serviceDesc, srv)
}

func _CRUD_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCRUD)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CRUD/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Create(ctx, req.(*CreateCRUD))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCRUD)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CRUD/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Get(ctx, req.(*GetCRUD))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CRUDObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CRUD/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Update(ctx, req.(*CRUDObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CRUDObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CRUD/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Delete(ctx, req.(*CRUDObject))
	}
	return interceptor(ctx, in, info, handler)
}

var _CRUD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CRUD",
	HandlerType: (*CRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CRUD_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CRUD_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CRUD_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CRUD_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud.proto",
}

func (m *CRUDObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CRUDObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCrud(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCrud(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CreateCRUD) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateCRUD) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCrud(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCrud(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GetCRUD) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCRUD) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCrud(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintCrud(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CRUDObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCrud(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCrud(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CreateCRUD) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCrud(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCrud(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetCRUD) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCrud(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCrud(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCrud(x uint64) (n int) {
	return sovCrud(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CRUDObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrud
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CRUDObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CRUDObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrud
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrud
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrud
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrud
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrud(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCrud
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCrud
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateCRUD) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrud
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateCRUD: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateCRUD: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrud
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrud
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrud
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrud
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrud(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCrud
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCrud
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetCRUD) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrud
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetCRUD: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCRUD: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrud
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrud
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrud(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCrud
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCrud
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrud
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCrud(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCrud
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCrud
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCrud(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCrud
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCrud
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCrud
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCrud
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCrud
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCrud
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCrud(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCrud
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCrud = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCrud   = fmt.Errorf("proto: integer overflow")
)
