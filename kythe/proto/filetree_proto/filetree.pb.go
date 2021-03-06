// Code generated by protoc-gen-gogo.
// source: kythe/proto/filetree.proto
// DO NOT EDIT!

/*
	Package filetree_proto is a generated protocol buffer package.

	It is generated from these files:
		kythe/proto/filetree.proto

	It has these top-level messages:
		CorpusRootsRequest
		CorpusRootsReply
		DirectoryRequest
		DirectoryReply
*/
package filetree_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CorpusRootsRequest struct {
}

func (m *CorpusRootsRequest) Reset()                    { *m = CorpusRootsRequest{} }
func (m *CorpusRootsRequest) String() string            { return proto.CompactTextString(m) }
func (*CorpusRootsRequest) ProtoMessage()               {}
func (*CorpusRootsRequest) Descriptor() ([]byte, []int) { return fileDescriptorFiletree, []int{0} }

type CorpusRootsReply struct {
	Corpus []*CorpusRootsReply_Corpus `protobuf:"bytes,1,rep,name=corpus" json:"corpus,omitempty"`
}

func (m *CorpusRootsReply) Reset()                    { *m = CorpusRootsReply{} }
func (m *CorpusRootsReply) String() string            { return proto.CompactTextString(m) }
func (*CorpusRootsReply) ProtoMessage()               {}
func (*CorpusRootsReply) Descriptor() ([]byte, []int) { return fileDescriptorFiletree, []int{1} }

func (m *CorpusRootsReply) GetCorpus() []*CorpusRootsReply_Corpus {
	if m != nil {
		return m.Corpus
	}
	return nil
}

type CorpusRootsReply_Corpus struct {
	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Root []string `protobuf:"bytes,2,rep,name=root" json:"root,omitempty"`
}

func (m *CorpusRootsReply_Corpus) Reset()         { *m = CorpusRootsReply_Corpus{} }
func (m *CorpusRootsReply_Corpus) String() string { return proto.CompactTextString(m) }
func (*CorpusRootsReply_Corpus) ProtoMessage()    {}
func (*CorpusRootsReply_Corpus) Descriptor() ([]byte, []int) {
	return fileDescriptorFiletree, []int{1, 0}
}

func (m *CorpusRootsReply_Corpus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CorpusRootsReply_Corpus) GetRoot() []string {
	if m != nil {
		return m.Root
	}
	return nil
}

type DirectoryRequest struct {
	Corpus string `protobuf:"bytes,1,opt,name=corpus,proto3" json:"corpus,omitempty"`
	Root   string `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	Path   string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (m *DirectoryRequest) Reset()                    { *m = DirectoryRequest{} }
func (m *DirectoryRequest) String() string            { return proto.CompactTextString(m) }
func (*DirectoryRequest) ProtoMessage()               {}
func (*DirectoryRequest) Descriptor() ([]byte, []int) { return fileDescriptorFiletree, []int{2} }

func (m *DirectoryRequest) GetCorpus() string {
	if m != nil {
		return m.Corpus
	}
	return ""
}

func (m *DirectoryRequest) GetRoot() string {
	if m != nil {
		return m.Root
	}
	return ""
}

func (m *DirectoryRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type DirectoryReply struct {
	// Set of tickets for each contained sub-directory's corpus, root, and path.
	Subdirectory []string `protobuf:"bytes,1,rep,name=subdirectory" json:"subdirectory,omitempty"`
	// Set of file tickets contained within this directory.
	File []string `protobuf:"bytes,2,rep,name=file" json:"file,omitempty"`
}

func (m *DirectoryReply) Reset()                    { *m = DirectoryReply{} }
func (m *DirectoryReply) String() string            { return proto.CompactTextString(m) }
func (*DirectoryReply) ProtoMessage()               {}
func (*DirectoryReply) Descriptor() ([]byte, []int) { return fileDescriptorFiletree, []int{3} }

func (m *DirectoryReply) GetSubdirectory() []string {
	if m != nil {
		return m.Subdirectory
	}
	return nil
}

func (m *DirectoryReply) GetFile() []string {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*CorpusRootsRequest)(nil), "kythe.proto.CorpusRootsRequest")
	proto.RegisterType((*CorpusRootsReply)(nil), "kythe.proto.CorpusRootsReply")
	proto.RegisterType((*CorpusRootsReply_Corpus)(nil), "kythe.proto.CorpusRootsReply.Corpus")
	proto.RegisterType((*DirectoryRequest)(nil), "kythe.proto.DirectoryRequest")
	proto.RegisterType((*DirectoryReply)(nil), "kythe.proto.DirectoryReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FileTreeService service

type FileTreeServiceClient interface {
	// CorpusRoots returns all known corpus/root pairs for stored files.
	CorpusRoots(ctx context.Context, in *CorpusRootsRequest, opts ...grpc.CallOption) (*CorpusRootsReply, error)
	// Directory returns the file/sub-directory contents of the given directory.
	Directory(ctx context.Context, in *DirectoryRequest, opts ...grpc.CallOption) (*DirectoryReply, error)
}

type fileTreeServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileTreeServiceClient(cc *grpc.ClientConn) FileTreeServiceClient {
	return &fileTreeServiceClient{cc}
}

func (c *fileTreeServiceClient) CorpusRoots(ctx context.Context, in *CorpusRootsRequest, opts ...grpc.CallOption) (*CorpusRootsReply, error) {
	out := new(CorpusRootsReply)
	err := grpc.Invoke(ctx, "/kythe.proto.FileTreeService/CorpusRoots", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileTreeServiceClient) Directory(ctx context.Context, in *DirectoryRequest, opts ...grpc.CallOption) (*DirectoryReply, error) {
	out := new(DirectoryReply)
	err := grpc.Invoke(ctx, "/kythe.proto.FileTreeService/Directory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileTreeService service

type FileTreeServiceServer interface {
	// CorpusRoots returns all known corpus/root pairs for stored files.
	CorpusRoots(context.Context, *CorpusRootsRequest) (*CorpusRootsReply, error)
	// Directory returns the file/sub-directory contents of the given directory.
	Directory(context.Context, *DirectoryRequest) (*DirectoryReply, error)
}

func RegisterFileTreeServiceServer(s *grpc.Server, srv FileTreeServiceServer) {
	s.RegisterService(&_FileTreeService_serviceDesc, srv)
}

func _FileTreeService_CorpusRoots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CorpusRootsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTreeServiceServer).CorpusRoots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kythe.proto.FileTreeService/CorpusRoots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTreeServiceServer).CorpusRoots(ctx, req.(*CorpusRootsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTreeService_Directory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTreeServiceServer).Directory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kythe.proto.FileTreeService/Directory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTreeServiceServer).Directory(ctx, req.(*DirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileTreeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kythe.proto.FileTreeService",
	HandlerType: (*FileTreeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CorpusRoots",
			Handler:    _FileTreeService_CorpusRoots_Handler,
		},
		{
			MethodName: "Directory",
			Handler:    _FileTreeService_Directory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kythe/proto/filetree.proto",
}

func (m *CorpusRootsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CorpusRootsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *CorpusRootsReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CorpusRootsReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Corpus) > 0 {
		for _, msg := range m.Corpus {
			dAtA[i] = 0xa
			i++
			i = encodeVarintFiletree(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CorpusRootsReply_Corpus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CorpusRootsReply_Corpus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFiletree(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Root) > 0 {
		for _, s := range m.Root {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *DirectoryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DirectoryRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Corpus) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFiletree(dAtA, i, uint64(len(m.Corpus)))
		i += copy(dAtA[i:], m.Corpus)
	}
	if len(m.Root) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFiletree(dAtA, i, uint64(len(m.Root)))
		i += copy(dAtA[i:], m.Root)
	}
	if len(m.Path) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFiletree(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	return i, nil
}

func (m *DirectoryReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DirectoryReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Subdirectory) > 0 {
		for _, s := range m.Subdirectory {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.File) > 0 {
		for _, s := range m.File {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64Filetree(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Filetree(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFiletree(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CorpusRootsRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *CorpusRootsReply) Size() (n int) {
	var l int
	_ = l
	if len(m.Corpus) > 0 {
		for _, e := range m.Corpus {
			l = e.Size()
			n += 1 + l + sovFiletree(uint64(l))
		}
	}
	return n
}

func (m *CorpusRootsReply_Corpus) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFiletree(uint64(l))
	}
	if len(m.Root) > 0 {
		for _, s := range m.Root {
			l = len(s)
			n += 1 + l + sovFiletree(uint64(l))
		}
	}
	return n
}

func (m *DirectoryRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Corpus)
	if l > 0 {
		n += 1 + l + sovFiletree(uint64(l))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovFiletree(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovFiletree(uint64(l))
	}
	return n
}

func (m *DirectoryReply) Size() (n int) {
	var l int
	_ = l
	if len(m.Subdirectory) > 0 {
		for _, s := range m.Subdirectory {
			l = len(s)
			n += 1 + l + sovFiletree(uint64(l))
		}
	}
	if len(m.File) > 0 {
		for _, s := range m.File {
			l = len(s)
			n += 1 + l + sovFiletree(uint64(l))
		}
	}
	return n
}

func sovFiletree(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFiletree(x uint64) (n int) {
	return sovFiletree(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CorpusRootsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFiletree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CorpusRootsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CorpusRootsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFiletree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFiletree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CorpusRootsReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFiletree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CorpusRootsReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CorpusRootsReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Corpus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFiletree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFiletree
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Corpus = append(m.Corpus, &CorpusRootsReply_Corpus{})
			if err := m.Corpus[len(m.Corpus)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFiletree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFiletree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CorpusRootsReply_Corpus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFiletree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Corpus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Corpus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFiletree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFiletree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFiletree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFiletree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFiletree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFiletree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DirectoryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFiletree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DirectoryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DirectoryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Corpus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFiletree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFiletree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Corpus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFiletree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFiletree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFiletree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFiletree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFiletree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFiletree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DirectoryReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFiletree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DirectoryReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DirectoryReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subdirectory", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFiletree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFiletree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subdirectory = append(m.Subdirectory, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFiletree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFiletree
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = append(m.File, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFiletree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFiletree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFiletree(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFiletree
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
					return 0, ErrIntOverflowFiletree
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
					return 0, ErrIntOverflowFiletree
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFiletree
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFiletree
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
				next, err := skipFiletree(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthFiletree = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFiletree   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("kythe/proto/filetree.proto", fileDescriptorFiletree) }

var fileDescriptorFiletree = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xbb, 0xff, 0xfe, 0x29, 0x74, 0x2a, 0x5a, 0x16, 0x91, 0x10, 0x31, 0x2d, 0x8b, 0x87,
	0x9e, 0x52, 0xad, 0x57, 0x4f, 0x2a, 0xa2, 0x27, 0x21, 0xfa, 0x02, 0x6d, 0x3a, 0xb6, 0xc1, 0xd4,
	0x89, 0x9b, 0x4d, 0x21, 0x57, 0x9f, 0xc2, 0x97, 0xf0, 0x3d, 0x3c, 0xfa, 0x08, 0x12, 0x5f, 0x44,
	0x76, 0x93, 0xc8, 0xa6, 0x50, 0x4f, 0x3b, 0xf3, 0xdb, 0x99, 0x8f, 0x6f, 0x3e, 0x70, 0x9f, 0x72,
	0xb5, 0xc4, 0x71, 0x22, 0x49, 0xd1, 0xf8, 0x31, 0x8a, 0x51, 0x49, 0x44, 0xdf, 0xb4, 0xbc, 0x67,
	0xfe, 0xca, 0x46, 0xec, 0x03, 0xbf, 0x24, 0x99, 0x64, 0x69, 0x40, 0xa4, 0xd2, 0x00, 0x5f, 0x32,
	0x4c, 0x95, 0x78, 0x65, 0xd0, 0x6f, 0xe0, 0x24, 0xce, 0xf9, 0x39, 0x74, 0x42, 0xc3, 0x1c, 0x36,
	0x6c, 0x8f, 0x7a, 0x93, 0x63, 0xdf, 0x12, 0xf2, 0x37, 0xc7, 0x6b, 0x50, 0xed, 0xb8, 0x27, 0xd0,
	0x29, 0x09, 0xe7, 0xf0, 0xff, 0x79, 0xba, 0x42, 0x87, 0x0d, 0xd9, 0xa8, 0x1b, 0x98, 0x5a, 0x33,
	0x49, 0xa4, 0x9c, 0x7f, 0xc3, 0xb6, 0x66, 0xba, 0x16, 0x01, 0xf4, 0xaf, 0x22, 0x89, 0xa1, 0x22,
	0x99, 0x57, 0xc6, 0xf8, 0x81, 0xe5, 0x41, 0x6f, 0x57, 0x9d, 0xb5, 0xcf, 0xea, 0x7d, 0xcd, 0x92,
	0xa9, 0x5a, 0x3a, 0xed, 0x92, 0xe9, 0x5a, 0xdc, 0xc0, 0xae, 0xa5, 0xa9, 0xaf, 0x12, 0xb0, 0x93,
	0x66, 0xb3, 0x79, 0x0d, 0xcd, 0x6d, 0xdd, 0xa0, 0xc1, 0xb4, 0x92, 0xce, 0xb0, 0x76, 0xa7, 0xeb,
	0xc9, 0x3b, 0x83, 0xbd, 0xeb, 0x28, 0xc6, 0x07, 0x89, 0x78, 0x8f, 0x72, 0x1d, 0x85, 0xc8, 0xef,
	0xa0, 0x67, 0xc5, 0xc0, 0x07, 0xdb, 0x03, 0x32, 0xd7, 0xb8, 0x47, 0x7f, 0x26, 0x28, 0x5a, 0xfc,
	0x16, 0xba, 0xbf, 0x76, 0x79, 0x73, 0x7a, 0x33, 0x1a, 0xf7, 0x70, 0xdb, 0xb7, 0x91, 0xba, 0x38,
	0xfd, 0x28, 0x3c, 0xf6, 0x59, 0x78, 0xec, 0xab, 0xf0, 0xd8, 0xdb, 0xb7, 0xd7, 0x82, 0x41, 0x48,
	0x2b, 0x7f, 0x41, 0xb4, 0x88, 0xd1, 0x9f, 0xe3, 0x5a, 0x11, 0xc5, 0xa9, 0xad, 0x31, 0xeb, 0x98,
	0xe7, 0xec, 0x27, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xec, 0x35, 0x30, 0x4d, 0x02, 0x00, 0x00,
}
