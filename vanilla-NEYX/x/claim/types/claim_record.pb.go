// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/claim/claim_record.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type ClaimRecord struct {
	Address           string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Claimable         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=claimable,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"claimable"`
	CompletedMissions []uint64                               `protobuf:"varint,3,rep,packed,name=completedMissions,proto3" json:"completedMissions,omitempty"`
	ClaimedMissions   []uint64                               `protobuf:"varint,4,rep,packed,name=claimedMissions,proto3" json:"claimedMissions,omitempty"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e8dcdb3caceb3d1, []int{0}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetCompletedMissions() []uint64 {
	if m != nil {
		return m.CompletedMissions
	}
	return nil
}

func (m *ClaimRecord) GetClaimedMissions() []uint64 {
	if m != nil {
		return m.ClaimedMissions
	}
	return nil
}

func init() {
	proto.RegisterType((*ClaimRecord)(nil), "modules.claim.ClaimRecord")
}

func init() { proto.RegisterFile("modules/claim/claim_record.proto", fileDescriptor_4e8dcdb3caceb3d1) }

var fileDescriptor_4e8dcdb3caceb3d1 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0xce, 0x49, 0xcc, 0xcc, 0x85, 0x90, 0xf1, 0x45, 0xa9, 0xc9, 0xf9,
	0x45, 0x29, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x50, 0x15, 0x7a, 0x60, 0x39, 0x29,
	0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x8c, 0x3e, 0x88, 0x05, 0x51, 0x24, 0x25, 0x99, 0x9c, 0x5f,
	0x9c, 0x9b, 0x5f, 0x1c, 0x0f, 0x91, 0x80, 0x70, 0x20, 0x52, 0x4a, 0x5f, 0x19, 0xb9, 0xb8, 0x9d,
	0x41, 0x5a, 0x83, 0xc0, 0xa6, 0x0a, 0x19, 0x71, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17,
	0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0x49, 0x5c, 0xda, 0xa2, 0x2b, 0x02, 0xd5, 0xe2, 0x08,
	0x91, 0x09, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x0f, 0x82, 0x29, 0x14, 0x8a, 0xe2, 0xe2, 0x04, 0xdb,
	0x9e, 0x98, 0x94, 0x93, 0x2a, 0xc1, 0x04, 0xd6, 0x65, 0x73, 0xe2, 0x9e, 0x3c, 0xc3, 0xad, 0x7b,
	0xf2, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0x50, 0x7b, 0xa1, 0x94,
	0x6e, 0x71, 0x4a, 0xb6, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0xb1, 0x9e, 0x67, 0x5e, 0xc9, 0xa5, 0x2d,
	0xba, 0x5c, 0x50, 0x3b, 0x3c, 0xf3, 0x4a, 0x82, 0x10, 0xc6, 0x09, 0xe9, 0x70, 0x09, 0x26, 0xe7,
	0xe7, 0x16, 0xe4, 0xa4, 0x96, 0xa4, 0xa6, 0xf8, 0x66, 0x16, 0x17, 0x67, 0xe6, 0xe7, 0x15, 0x4b,
	0x30, 0x2b, 0x30, 0x6b, 0xb0, 0x04, 0x61, 0x4a, 0x08, 0x69, 0x70, 0xf1, 0x83, 0xb5, 0x22, 0xa9,
	0x65, 0x01, 0xab, 0x45, 0x17, 0x76, 0x72, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0x75, 0x24, 0x27, 0x67, 0xa6, 0xe7, 0x65, 0x96, 0xa4, 0xea, 0xc3, 0x62, 0xa1, 0x02,
	0x1a, 0x0f, 0x60, 0x77, 0x27, 0xb1, 0x81, 0x43, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x09,
	0x0f, 0x6f, 0x43, 0xa5, 0x01, 0x00, 0x00,
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedMissions) > 0 {
		dAtA2 := make([]byte, len(m.ClaimedMissions)*10)
		var j1 int
		for _, num := range m.ClaimedMissions {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintClaimRecord(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CompletedMissions) > 0 {
		dAtA4 := make([]byte, len(m.CompletedMissions)*10)
		var j3 int
		for _, num := range m.CompletedMissions {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintClaimRecord(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Claimable.Size()
		i -= size
		if _, err := m.Claimable.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaimRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaimRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaimRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaimRecord(uint64(l))
	}
	l = m.Claimable.Size()
	n += 1 + l + sovClaimRecord(uint64(l))
	if len(m.CompletedMissions) > 0 {
		l = 0
		for _, e := range m.CompletedMissions {
			l += sovClaimRecord(uint64(e))
		}
		n += 1 + sovClaimRecord(uint64(l)) + l
	}
	if len(m.ClaimedMissions) > 0 {
		l = 0
		for _, e := range m.ClaimedMissions {
			l += sovClaimRecord(uint64(e))
		}
		n += 1 + sovClaimRecord(uint64(l)) + l
	}
	return n
}

func sovClaimRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaimRecord(x uint64) (n int) {
	return sovClaimRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimRecord
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claimable", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Claimable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CompletedMissions = append(m.CompletedMissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaimRecord
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaimRecord
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.CompletedMissions) == 0 {
					m.CompletedMissions = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaimRecord
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CompletedMissions = append(m.CompletedMissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletedMissions", wireType)
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ClaimedMissions = append(m.ClaimedMissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaimRecord
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaimRecord
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ClaimedMissions) == 0 {
					m.ClaimedMissions = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaimRecord
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ClaimedMissions = append(m.ClaimedMissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedMissions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaimRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimRecord
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
func skipClaimRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
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
				return 0, ErrInvalidLengthClaimRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaimRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaimRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaimRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaimRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaimRecord = fmt.Errorf("proto: unexpected end of group")
)
