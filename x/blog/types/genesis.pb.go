// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/blog/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the blog module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params           Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId           string        `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	PostList         []Post        `protobuf:"bytes,3,rep,name=postList,proto3" json:"postList"`
	PostCount        uint64        `protobuf:"varint,4,opt,name=postCount,proto3" json:"postCount,omitempty"`
	SentPostList     []SentPost    `protobuf:"bytes,5,rep,name=sentPostList,proto3" json:"sentPostList"`
	SentPostCount    uint64        `protobuf:"varint,6,opt,name=sentPostCount,proto3" json:"sentPostCount,omitempty"`
	TimeoutPostList  []TimeoutPost `protobuf:"bytes,7,rep,name=timeoutPostList,proto3" json:"timeoutPostList"`
	TimeoutPostCount uint64        `protobuf:"varint,8,opt,name=timeoutPostCount,proto3" json:"timeoutPostCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb8e9842ca79b33, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetPostList() []Post {
	if m != nil {
		return m.PostList
	}
	return nil
}

func (m *GenesisState) GetPostCount() uint64 {
	if m != nil {
		return m.PostCount
	}
	return 0
}

func (m *GenesisState) GetSentPostList() []SentPost {
	if m != nil {
		return m.SentPostList
	}
	return nil
}

func (m *GenesisState) GetSentPostCount() uint64 {
	if m != nil {
		return m.SentPostCount
	}
	return 0
}

func (m *GenesisState) GetTimeoutPostList() []TimeoutPost {
	if m != nil {
		return m.TimeoutPostList
	}
	return nil
}

func (m *GenesisState) GetTimeoutPostCount() uint64 {
	if m != nil {
		return m.TimeoutPostCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "example.blog.GenesisState")
}

func init() { proto.RegisterFile("example/blog/genesis.proto", fileDescriptor_abb8e9842ca79b33) }

var fileDescriptor_abb8e9842ca79b33 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x4f, 0xca, 0xc9, 0x4f, 0xd7, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0xca, 0xe9, 0x81, 0xe4, 0xa4, 0x04, 0x13,
	0x73, 0x33, 0xf3, 0xf2, 0xf5, 0xc1, 0x24, 0x44, 0x81, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x98,
	0xa9, 0x0f, 0x62, 0x41, 0x45, 0x25, 0x51, 0x8c, 0x2c, 0x48, 0x2c, 0x4a, 0xcc, 0x85, 0x9a, 0x28,
	0x25, 0x8e, 0x2a, 0x95, 0x5f, 0x5c, 0x02, 0x95, 0x90, 0x41, 0x91, 0x28, 0x4e, 0xcd, 0x2b, 0x89,
	0x47, 0x92, 0x95, 0x47, 0x91, 0x2d, 0xc9, 0xcc, 0x4d, 0xcd, 0x2f, 0x45, 0x56, 0xa0, 0x34, 0x99,
	0x99, 0x8b, 0xc7, 0x1d, 0xe2, 0xf6, 0xe0, 0x92, 0xc4, 0x92, 0x54, 0x21, 0x73, 0x2e, 0x36, 0x88,
	0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x22, 0x7a, 0xc8, 0x7e, 0xd1, 0x0b, 0x00, 0xcb,
	0x39, 0x71, 0x9e, 0xb8, 0x27, 0xcf, 0xb0, 0xe2, 0xf9, 0x06, 0x2d, 0xc6, 0x20, 0xa8, 0x72, 0x21,
	0x71, 0x2e, 0xf6, 0x82, 0xfc, 0xa2, 0x92, 0xf8, 0xcc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x36, 0x10, 0xd7, 0x33, 0x45, 0xc8, 0x84, 0x8b, 0x03, 0x64, 0xa1, 0x4f, 0x66, 0x71, 0x89,
	0x04, 0xb3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x10, 0x9a, 0x99, 0xf9, 0xc5, 0x25, 0x4e, 0x2c, 0x20,
	0x13, 0x83, 0xe0, 0x2a, 0x85, 0x64, 0xb8, 0x38, 0x41, 0x6c, 0xe7, 0xfc, 0xd2, 0xbc, 0x12, 0x09,
	0x16, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x84, 0x80, 0x90, 0x03, 0x17, 0x0f, 0xc8, 0xab, 0x01, 0x30,
	0x73, 0x59, 0xc1, 0xe6, 0x8a, 0xa1, 0x9a, 0x1b, 0x0c, 0x55, 0x01, 0x35, 0x1b, 0x45, 0x87, 0x90,
	0x0a, 0x17, 0x2f, 0x8c, 0x0f, 0xb1, 0x83, 0x0d, 0x6c, 0x07, 0xaa, 0xa0, 0x90, 0x27, 0x17, 0x3f,
	0x34, 0xd0, 0xe0, 0x56, 0xb1, 0x83, 0xad, 0x92, 0x44, 0xb5, 0x2a, 0x04, 0xa1, 0x08, 0x6a, 0x1b,
	0xba, 0x3e, 0x21, 0x2d, 0x2e, 0x01, 0x24, 0x21, 0x88, 0x9d, 0x1c, 0x60, 0x3b, 0x31, 0xc4, 0x9d,
	0xf4, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x04, 0x16, 0xa1, 0x15,
	0xd0, 0x28, 0xad, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x47, 0xa6, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x82, 0xd0, 0x37, 0xd3, 0x94, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutPostCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TimeoutPostCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.TimeoutPostList) > 0 {
		for iNdEx := len(m.TimeoutPostList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimeoutPostList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.SentPostCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SentPostCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SentPostList) > 0 {
		for iNdEx := len(m.SentPostList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SentPostList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.PostCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PostCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PostList) > 0 {
		for iNdEx := len(m.PostList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PostList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.PostList) > 0 {
		for _, e := range m.PostList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PostCount != 0 {
		n += 1 + sovGenesis(uint64(m.PostCount))
	}
	if len(m.SentPostList) > 0 {
		for _, e := range m.SentPostList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.SentPostCount != 0 {
		n += 1 + sovGenesis(uint64(m.SentPostCount))
	}
	if len(m.TimeoutPostList) > 0 {
		for _, e := range m.TimeoutPostList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TimeoutPostCount != 0 {
		n += 1 + sovGenesis(uint64(m.TimeoutPostCount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostList = append(m.PostList, Post{})
			if err := m.PostList[len(m.PostList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostCount", wireType)
			}
			m.PostCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentPostList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SentPostList = append(m.SentPostList, SentPost{})
			if err := m.SentPostList[len(m.SentPostList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentPostCount", wireType)
			}
			m.SentPostCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SentPostCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutPostList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeoutPostList = append(m.TimeoutPostList, TimeoutPost{})
			if err := m.TimeoutPostList[len(m.TimeoutPostList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutPostCount", wireType)
			}
			m.TimeoutPostCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutPostCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
