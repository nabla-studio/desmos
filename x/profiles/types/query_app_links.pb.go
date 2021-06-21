// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/query_app_links.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryUserApplicationLinksRequest represent the request used when querying the
// application links of a specific user
type QueryUserApplicationLinksRequest struct {
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Pagination defines an optional pagination for the request
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryUserApplicationLinksRequest) Reset()         { *m = QueryUserApplicationLinksRequest{} }
func (m *QueryUserApplicationLinksRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUserApplicationLinksRequest) ProtoMessage()    {}
func (*QueryUserApplicationLinksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d2a2b45fd238cb, []int{0}
}
func (m *QueryUserApplicationLinksRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserApplicationLinksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserApplicationLinksRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserApplicationLinksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserApplicationLinksRequest.Merge(m, src)
}
func (m *QueryUserApplicationLinksRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserApplicationLinksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserApplicationLinksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserApplicationLinksRequest proto.InternalMessageInfo

func (m *QueryUserApplicationLinksRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *QueryUserApplicationLinksRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryUserApplicationLinksResponse represents the response to the query used
// to get the application links for a specific user
type QueryUserApplicationLinksResponse struct {
	Links []ApplicationLink `protobuf:"bytes,1,rep,name=links,proto3" json:"links"`
	// Pagination defines the pagination response
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryUserApplicationLinksResponse) Reset()         { *m = QueryUserApplicationLinksResponse{} }
func (m *QueryUserApplicationLinksResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUserApplicationLinksResponse) ProtoMessage()    {}
func (*QueryUserApplicationLinksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d2a2b45fd238cb, []int{1}
}
func (m *QueryUserApplicationLinksResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserApplicationLinksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserApplicationLinksResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserApplicationLinksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserApplicationLinksResponse.Merge(m, src)
}
func (m *QueryUserApplicationLinksResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserApplicationLinksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserApplicationLinksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserApplicationLinksResponse proto.InternalMessageInfo

func (m *QueryUserApplicationLinksResponse) GetLinks() []ApplicationLink {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *QueryUserApplicationLinksResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryUserApplicationLinksRequest)(nil), "desmos.profiles.v1beta1.QueryUserApplicationLinksRequest")
	proto.RegisterType((*QueryUserApplicationLinksResponse)(nil), "desmos.profiles.v1beta1.QueryUserApplicationLinksResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/query_app_links.proto", fileDescriptor_18d2a2b45fd238cb)
}

var fileDescriptor_18d2a2b45fd238cb = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0xbb, 0xff, 0x3f, 0x9a, 0x58, 0x6e, 0x8d, 0x89, 0xd0, 0x98, 0x5a, 0x39, 0x68, 0x63,
	0xc2, 0x6e, 0xc0, 0x27, 0x90, 0x18, 0xf5, 0xe0, 0x41, 0x9b, 0x78, 0xf1, 0x42, 0xb6, 0xb0, 0xd4,
	0xc6, 0xd2, 0x59, 0xba, 0x5b, 0x23, 0x17, 0x9f, 0xc1, 0x67, 0xf1, 0x29, 0x38, 0x72, 0xf4, 0x64,
	0x0c, 0xbc, 0x88, 0xe9, 0xee, 0x22, 0x84, 0x88, 0xde, 0x66, 0xf2, 0x7d, 0xfb, 0xcd, 0x6f, 0x76,
	0xec, 0x66, 0x9f, 0x89, 0x21, 0x08, 0xc2, 0x73, 0x18, 0x24, 0x29, 0x13, 0xe4, 0xa9, 0x15, 0x31,
	0x49, 0x5b, 0x64, 0x54, 0xb0, 0x7c, 0xdc, 0xa5, 0x9c, 0x77, 0xd3, 0x24, 0x7b, 0x14, 0x98, 0xe7,
	0x20, 0xc1, 0xd9, 0xd3, 0x76, 0xbc, 0xb0, 0x63, 0x63, 0x77, 0x77, 0x63, 0x88, 0x41, 0x79, 0x48,
	0x59, 0x69, 0xbb, 0xbb, 0x1f, 0x03, 0xc4, 0x29, 0x23, 0x94, 0x27, 0x84, 0x66, 0x19, 0x48, 0x2a,
	0x13, 0xc8, 0x4c, 0x98, 0x5b, 0x37, 0xaa, 0xea, 0xa2, 0x62, 0x40, 0x68, 0x36, 0x36, 0x12, 0xde,
	0x84, 0x35, 0x84, 0x3e, 0x4b, 0xc5, 0x3a, 0x97, 0x5b, 0xef, 0x41, 0xe9, 0xef, 0x6a, 0x02, 0xdd,
	0x18, 0xe9, 0x44, 0x77, 0x24, 0xa2, 0x82, 0xe9, 0xad, 0xbe, 0xc3, 0x38, 0x8d, 0x93, 0x4c, 0x21,
	0x69, 0x6f, 0xe3, 0xc5, 0xf6, 0x6f, 0x4b, 0xc7, 0x9d, 0x60, 0xf9, 0x19, 0xe7, 0x69, 0xd2, 0x53,
	0xea, 0x75, 0x39, 0x29, 0x64, 0xa3, 0x82, 0x09, 0xe9, 0x38, 0x76, 0xa5, 0x10, 0x2c, 0xaf, 0x21,
	0x1f, 0x05, 0x3b, 0xa1, 0xaa, 0x9d, 0x0b, 0xdb, 0x5e, 0x66, 0xd5, 0xfe, 0xf9, 0x28, 0xa8, 0xb6,
	0x8f, 0xb0, 0xc1, 0x28, 0x07, 0x63, 0x35, 0x78, 0xf1, 0x5b, 0xf8, 0x86, 0xc6, 0xcc, 0xe4, 0x85,
	0x2b, 0x2f, 0x1b, 0x6f, 0xc8, 0x3e, 0xfc, 0x05, 0x40, 0x70, 0xc8, 0x04, 0x73, 0xce, 0xed, 0x2d,
	0xb5, 0x7b, 0x0d, 0xf9, 0xff, 0x83, 0x6a, 0x3b, 0xc0, 0x1b, 0x8e, 0x82, 0xd7, 0x12, 0x3a, 0x95,
	0xc9, 0xc7, 0x81, 0x15, 0xea, 0xc7, 0xce, 0xe5, 0x0f, 0xcc, 0xc7, 0x7f, 0x32, 0x6b, 0x84, 0x55,
	0xe8, 0xce, 0xd5, 0x64, 0xe6, 0xa1, 0xe9, 0xcc, 0x43, 0x9f, 0x33, 0x0f, 0xbd, 0xce, 0x3d, 0x6b,
	0x3a, 0xf7, 0xac, 0xf7, 0xb9, 0x67, 0xdd, 0xe3, 0x38, 0x91, 0x0f, 0x45, 0x84, 0x7b, 0x30, 0x24,
	0x9a, 0xb1, 0x99, 0xd2, 0x48, 0x98, 0x9a, 0x3c, 0x2f, 0xcf, 0x2b, 0xc7, 0x9c, 0x89, 0x68, 0x5b,
	0x5d, 0xe1, 0xf4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x04, 0x99, 0x2a, 0xde, 0x95, 0x02, 0x00, 0x00,
}

func (m *QueryUserApplicationLinksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserApplicationLinksRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserApplicationLinksRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryUserApplicationLinksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserApplicationLinksResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserApplicationLinksResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Links) > 0 {
		for iNdEx := len(m.Links) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Links[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryAppLinks(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryAppLinks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryUserApplicationLinksRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	return n
}

func (m *QueryUserApplicationLinksResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovQueryAppLinks(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	return n
}

func sovQueryAppLinks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryAppLinks(x uint64) (n int) {
	return sovQueryAppLinks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryUserApplicationLinksRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryAppLinks
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
			return fmt.Errorf("proto: QueryUserApplicationLinksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserApplicationLinksRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryAppLinks
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
func (m *QueryUserApplicationLinksResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryAppLinks
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
			return fmt.Errorf("proto: QueryUserApplicationLinksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserApplicationLinksResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, ApplicationLink{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryAppLinks
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
func skipQueryAppLinks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryAppLinks
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
					return 0, ErrIntOverflowQueryAppLinks
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryAppLinks
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
				return 0, ErrInvalidLengthQueryAppLinks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryAppLinks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryAppLinks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryAppLinks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryAppLinks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryAppLinks = fmt.Errorf("proto: unexpected end of group")
)