// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lottery/lottery/stored_bid.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type StoredBid struct {
	Index     string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	LotteryId uint64 `protobuf:"varint,2,opt,name=lotteryId,proto3" json:"lotteryId,omitempty"`
	BetAmount uint64 `protobuf:"varint,3,opt,name=betAmount,proto3" json:"betAmount,omitempty"`
	Bidder    string `protobuf:"bytes,4,opt,name=bidder,proto3" json:"bidder,omitempty"`
}

func (m *StoredBid) Reset()         { *m = StoredBid{} }
func (m *StoredBid) String() string { return proto.CompactTextString(m) }
func (*StoredBid) ProtoMessage()    {}
func (*StoredBid) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9f84a1670cf0c76, []int{0}
}
func (m *StoredBid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredBid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredBid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredBid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredBid.Merge(m, src)
}
func (m *StoredBid) XXX_Size() int {
	return m.Size()
}
func (m *StoredBid) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredBid.DiscardUnknown(m)
}

var xxx_messageInfo_StoredBid proto.InternalMessageInfo

func (m *StoredBid) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredBid) GetLotteryId() uint64 {
	if m != nil {
		return m.LotteryId
	}
	return 0
}

func (m *StoredBid) GetBetAmount() uint64 {
	if m != nil {
		return m.BetAmount
	}
	return 0
}

func (m *StoredBid) GetBidder() string {
	if m != nil {
		return m.Bidder
	}
	return ""
}

func init() {
	proto.RegisterType((*StoredBid)(nil), "dedicateddev.lottery.lottery.StoredBid")
}

func init() { proto.RegisterFile("lottery/lottery/stored_bid.proto", fileDescriptor_c9f84a1670cf0c76) }

var fileDescriptor_c9f84a1670cf0c76 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0xc9, 0x2f, 0x29,
	0x49, 0x2d, 0xaa, 0xd4, 0x87, 0xd1, 0xc5, 0x25, 0xf9, 0x45, 0xa9, 0x29, 0xf1, 0x49, 0x99, 0x29,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x32, 0x29, 0xa9, 0x29, 0x99, 0xc9, 0x89, 0x25, 0xa9,
	0x29, 0x29, 0xa9, 0x65, 0x7a, 0x50, 0x65, 0x30, 0x5a, 0xa9, 0x94, 0x8b, 0x33, 0x18, 0xac, 0xc3,
	0x29, 0x33, 0x45, 0x48, 0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x08, 0xc2, 0x11, 0x92, 0xe1, 0xe2, 0x84, 0xaa, 0xf6, 0x4c, 0x91, 0x60, 0x52, 0x60,
	0xd4, 0x60, 0x09, 0x42, 0x08, 0x80, 0x64, 0x93, 0x52, 0x4b, 0x1c, 0x73, 0xf3, 0x4b, 0xf3, 0x4a,
	0x24, 0x98, 0x21, 0xb2, 0x70, 0x01, 0x21, 0x31, 0x2e, 0xb6, 0xa4, 0xcc, 0x94, 0x94, 0xd4, 0x22,
	0x09, 0x16, 0xb0, 0x91, 0x50, 0x9e, 0x93, 0xe7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0xe9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xbb, 0xc0,
	0x5c, 0xee, 0x92, 0x5a, 0x06, 0xf7, 0x60, 0x05, 0x9c, 0x55, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0xf6, 0xa6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x09, 0x88, 0x30, 0x38, 0x0a, 0x01, 0x00,
	0x00,
}

func (m *StoredBid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredBid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredBid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintStoredBid(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x22
	}
	if m.BetAmount != 0 {
		i = encodeVarintStoredBid(dAtA, i, uint64(m.BetAmount))
		i--
		dAtA[i] = 0x18
	}
	if m.LotteryId != 0 {
		i = encodeVarintStoredBid(dAtA, i, uint64(m.LotteryId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredBid(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredBid(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredBid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredBid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredBid(uint64(l))
	}
	if m.LotteryId != 0 {
		n += 1 + sovStoredBid(uint64(m.LotteryId))
	}
	if m.BetAmount != 0 {
		n += 1 + sovStoredBid(uint64(m.BetAmount))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovStoredBid(uint64(l))
	}
	return n
}

func sovStoredBid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredBid(x uint64) (n int) {
	return sovStoredBid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredBid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredBid
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
			return fmt.Errorf("proto: StoredBid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredBid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredBid
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
				return ErrInvalidLengthStoredBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LotteryId", wireType)
			}
			m.LotteryId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LotteryId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetAmount", wireType)
			}
			m.BetAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BetAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredBid
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
				return ErrInvalidLengthStoredBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStoredBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredBid
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
func skipStoredBid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredBid
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
					return 0, ErrIntOverflowStoredBid
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
					return 0, ErrIntOverflowStoredBid
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
				return 0, ErrInvalidLengthStoredBid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredBid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredBid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredBid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredBid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredBid = fmt.Errorf("proto: unexpected end of group")
)
