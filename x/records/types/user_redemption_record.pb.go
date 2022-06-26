// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: records/user_redemption_record.proto

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

type UserRedemptionRecord struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender      string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver    string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount      uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom       string `protobuf:"bytes,5,opt,name=denom,proto3" json:"denom,omitempty"`
	HostZoneId  string `protobuf:"bytes,6,opt,name=hostZoneId,proto3" json:"hostZoneId,omitempty"`
	EpochNumber int32  `protobuf:"varint,7,opt,name=epochNumber,proto3" json:"epochNumber,omitempty"`
	IsClaimable bool   `protobuf:"varint,8,opt,name=isClaimable,proto3" json:"isClaimable,omitempty"`
}

func (m *UserRedemptionRecord) Reset()         { *m = UserRedemptionRecord{} }
func (m *UserRedemptionRecord) String() string { return proto.CompactTextString(m) }
func (*UserRedemptionRecord) ProtoMessage()    {}
func (*UserRedemptionRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_297c97fa417e7065, []int{0}
}
func (m *UserRedemptionRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserRedemptionRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserRedemptionRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserRedemptionRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRedemptionRecord.Merge(m, src)
}
func (m *UserRedemptionRecord) XXX_Size() int {
	return m.Size()
}
func (m *UserRedemptionRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRedemptionRecord.DiscardUnknown(m)
}

var xxx_messageInfo_UserRedemptionRecord proto.InternalMessageInfo

func (m *UserRedemptionRecord) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRedemptionRecord) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *UserRedemptionRecord) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *UserRedemptionRecord) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *UserRedemptionRecord) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *UserRedemptionRecord) GetHostZoneId() string {
	if m != nil {
		return m.HostZoneId
	}
	return ""
}

func (m *UserRedemptionRecord) GetEpochNumber() int32 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func (m *UserRedemptionRecord) GetIsClaimable() bool {
	if m != nil {
		return m.IsClaimable
	}
	return false
}

func init() {
	proto.RegisterType((*UserRedemptionRecord)(nil), "StrideLabs.stride.records.UserRedemptionRecord")
}

func init() {
	proto.RegisterFile("records/user_redemption_record.proto", fileDescriptor_297c97fa417e7065)
}

var fileDescriptor_297c97fa417e7065 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4a, 0xc4, 0x30,
	0x1c, 0xc6, 0x2f, 0xe7, 0x5d, 0xad, 0x11, 0x1c, 0xc2, 0x21, 0xd1, 0x21, 0x14, 0x71, 0xe8, 0x72,
	0xed, 0xe0, 0x1b, 0xe8, 0x20, 0x82, 0x38, 0x44, 0x5c, 0x6e, 0x39, 0xda, 0xe6, 0x8f, 0x0d, 0x5c,
	0x9a, 0x92, 0xa4, 0xa2, 0x6f, 0xe1, 0x63, 0x39, 0xde, 0xe8, 0x28, 0xed, 0xea, 0x43, 0x48, 0xd3,
	0x7a, 0x74, 0xcb, 0xf7, 0x7d, 0xf9, 0xfd, 0x87, 0x1f, 0xbe, 0x36, 0x50, 0x68, 0x23, 0x6c, 0xda,
	0x58, 0x30, 0x5b, 0x03, 0x02, 0x54, 0xed, 0xa4, 0xae, 0xb6, 0x43, 0x9f, 0xd4, 0x46, 0x3b, 0x4d,
	0x2e, 0x9e, 0x9d, 0x91, 0x02, 0x1e, 0xb3, 0xdc, 0x26, 0xd6, 0x3f, 0x93, 0x91, 0xbb, 0xfa, 0x45,
	0x78, 0xf5, 0x62, 0xc1, 0xf0, 0x03, 0xca, 0xfd, 0x42, 0xce, 0xf0, 0x5c, 0x0a, 0x8a, 0x22, 0x14,
	0x2f, 0xf8, 0x5c, 0x0a, 0x72, 0x8e, 0x03, 0x0b, 0x95, 0x00, 0x43, 0xe7, 0x11, 0x8a, 0x4f, 0xf8,
	0x98, 0xc8, 0x25, 0x0e, 0x0d, 0x14, 0x20, 0xdf, 0xc0, 0xd0, 0x23, 0xbf, 0x1c, 0x72, 0xcf, 0x64,
	0x4a, 0x37, 0x95, 0xa3, 0x0b, 0x7f, 0x67, 0x4c, 0x64, 0x85, 0x97, 0x02, 0x2a, 0xad, 0xe8, 0xd2,
	0x03, 0x43, 0x20, 0x0c, 0xe3, 0x52, 0x5b, 0xb7, 0xd1, 0x15, 0x3c, 0x08, 0x1a, 0xf8, 0x69, 0xd2,
	0x90, 0x08, 0x9f, 0x42, 0xad, 0x8b, 0xf2, 0xa9, 0x51, 0x39, 0x18, 0x7a, 0x1c, 0xa1, 0x78, 0xc9,
	0xa7, 0x55, 0xff, 0x43, 0xda, 0xbb, 0x5d, 0x26, 0x55, 0x96, 0xef, 0x80, 0x86, 0x11, 0x8a, 0x43,
	0x3e, 0xad, 0x6e, 0xef, 0xbf, 0x5a, 0x86, 0xf6, 0x2d, 0x43, 0x3f, 0x2d, 0x43, 0x9f, 0x1d, 0x9b,
	0xed, 0x3b, 0x36, 0xfb, 0xee, 0xd8, 0x6c, 0xb3, 0x7e, 0x95, 0xae, 0x6c, 0xf2, 0xa4, 0xd0, 0x2a,
	0x1d, 0x74, 0xad, 0x7b, 0x5f, 0xe9, 0xe0, 0x2b, 0x7d, 0x4f, 0xff, 0x4d, 0xbb, 0x8f, 0x1a, 0x6c,
	0x1e, 0x78, 0xb3, 0x37, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0xdc, 0x6d, 0x8d, 0x81, 0x01,
	0x00, 0x00,
}

func (m *UserRedemptionRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserRedemptionRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserRedemptionRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsClaimable {
		i--
		if m.IsClaimable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.EpochNumber != 0 {
		i = encodeVarintUserRedemptionRecord(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x38
	}
	if len(m.HostZoneId) > 0 {
		i -= len(m.HostZoneId)
		copy(dAtA[i:], m.HostZoneId)
		i = encodeVarintUserRedemptionRecord(dAtA, i, uint64(len(m.HostZoneId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintUserRedemptionRecord(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Amount != 0 {
		i = encodeVarintUserRedemptionRecord(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintUserRedemptionRecord(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintUserRedemptionRecord(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUserRedemptionRecord(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUserRedemptionRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserRedemptionRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserRedemptionRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUserRedemptionRecord(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovUserRedemptionRecord(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovUserRedemptionRecord(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovUserRedemptionRecord(uint64(m.Amount))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovUserRedemptionRecord(uint64(l))
	}
	l = len(m.HostZoneId)
	if l > 0 {
		n += 1 + l + sovUserRedemptionRecord(uint64(l))
	}
	if m.EpochNumber != 0 {
		n += 1 + sovUserRedemptionRecord(uint64(m.EpochNumber))
	}
	if m.IsClaimable {
		n += 2
	}
	return n
}

func sovUserRedemptionRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserRedemptionRecord(x uint64) (n int) {
	return sovUserRedemptionRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserRedemptionRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserRedemptionRecord
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
			return fmt.Errorf("proto: UserRedemptionRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserRedemptionRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRedemptionRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRedemptionRecord
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
				return ErrInvalidLengthUserRedemptionRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserRedemptionRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRedemptionRecord
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
				return ErrInvalidLengthUserRedemptionRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserRedemptionRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRedemptionRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRedemptionRecord
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
				return ErrInvalidLengthUserRedemptionRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserRedemptionRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZoneId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRedemptionRecord
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
				return ErrInvalidLengthUserRedemptionRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserRedemptionRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostZoneId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRedemptionRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsClaimable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRedemptionRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsClaimable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipUserRedemptionRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserRedemptionRecord
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
func skipUserRedemptionRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserRedemptionRecord
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
					return 0, ErrIntOverflowUserRedemptionRecord
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
					return 0, ErrIntOverflowUserRedemptionRecord
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
				return 0, ErrInvalidLengthUserRedemptionRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserRedemptionRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserRedemptionRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserRedemptionRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserRedemptionRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserRedemptionRecord = fmt.Errorf("proto: unexpected end of group")
)
