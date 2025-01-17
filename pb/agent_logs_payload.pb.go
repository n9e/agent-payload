// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/logs/agent_logs_payload.proto

package pb

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

type Log struct {
	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// from host
	Hostname string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// from config
	Service string `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	Source  string `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	// from config, container tags, ...
	Tags                 []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad4736ffde8ea121, []int{0}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Log.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return m.Size()
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Log) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Log) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Log) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Log) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Log) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Log) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Log)(nil), "pb.Log")
}

func init() {
	proto.RegisterFile("proto/logs/agent_logs_payload.proto", fileDescriptor_ad4736ffde8ea121)
}

var fileDescriptor_ad4736ffde8ea121 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x45, 0xc9, 0x74, 0xec, 0xd8, 0x20, 0x08, 0x59, 0x48, 0x10, 0xa9, 0x45, 0x37, 0xdd, 0xd8,
	0x2e, 0x5c, 0xb9, 0x74, 0xd6, 0x22, 0x43, 0x97, 0x6e, 0x86, 0xa4, 0x0d, 0x99, 0xc2, 0xa4, 0x2f,
	0xf4, 0xbd, 0x0a, 0x7e, 0x99, 0xbf, 0xe0, 0xd2, 0x4f, 0x90, 0x7e, 0x89, 0x34, 0xad, 0xce, 0xee,
	0x9e, 0x3c, 0xc2, 0xe1, 0x5e, 0x7e, 0xef, 0x7b, 0x20, 0x28, 0x8f, 0x60, 0xb1, 0x54, 0xd6, 0x74,
	0xb4, 0x9f, 0xe2, 0xde, 0xab, 0x8f, 0x23, 0xa8, 0xa6, 0x08, 0x57, 0xb1, 0xf2, 0xfa, 0xee, 0x93,
	0xf1, 0xe8, 0x05, 0xac, 0x90, 0x7c, 0xe3, 0x0c, 0xa2, 0xb2, 0x46, 0xb2, 0x8c, 0xe5, 0x49, 0xf5,
	0x87, 0xe2, 0x8a, 0xc7, 0x48, 0x8a, 0x06, 0x94, 0xab, 0x70, 0x58, 0x48, 0xdc, 0xf0, 0x84, 0x5a,
	0x67, 0x90, 0x94, 0xf3, 0x32, 0xca, 0x58, 0x1e, 0x55, 0xa7, 0x07, 0x71, 0xcd, 0xcf, 0x0f, 0x80,
	0xd4, 0x29, 0x67, 0xe4, 0x3a, 0xfc, 0xfb, 0xe7, 0xc9, 0x85, 0xa6, 0x7f, 0x6f, 0x6b, 0x23, 0xcf,
	0x66, 0xd7, 0x82, 0xc1, 0x05, 0x43, 0x5f, 0x1b, 0x19, 0x2f, 0xae, 0x40, 0x42, 0xf0, 0x35, 0x29,
	0x8b, 0x72, 0x93, 0x45, 0x79, 0x52, 0x85, 0xbc, 0x7d, 0xfd, 0x1a, 0x53, 0xf6, 0x3d, 0xa6, 0xec,
	0x67, 0x4c, 0x19, 0xbf, 0xac, 0xc1, 0x15, 0x4d, 0x53, 0x84, 0xb2, 0x85, 0xd7, 0xdb, 0x8b, 0xe7,
	0x29, 0xed, 0xe6, 0xc2, 0x3b, 0xf6, 0x76, 0x6b, 0x5b, 0x3a, 0x0c, 0xba, 0xa8, 0xc1, 0x95, 0xdd,
	0x93, 0x99, 0x57, 0x79, 0x58, 0x06, 0x29, 0xbd, 0xd6, 0x71, 0x18, 0xe5, 0xf1, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0x96, 0x8f, 0x5d, 0x3b, 0x01, 0x00, 0x00,
}

func (m *Log) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Log) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Service) > 0 {
		i -= len(m.Service)
		copy(dAtA[i:], m.Service)
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Service)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Hostname) > 0 {
		i -= len(m.Hostname)
		copy(dAtA[i:], m.Hostname)
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Hostname)))
		i--
		dAtA[i] = 0x22
	}
	if m.Timestamp != 0 {
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAgentLogsPayload(dAtA []byte, offset int, v uint64) int {
	offset -= sovAgentLogsPayload(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Log) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovAgentLogsPayload(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovAgentLogsPayload(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovAgentLogsPayload(uint64(m.Timestamp))
	}
	l = len(m.Hostname)
	if l > 0 {
		n += 1 + l + sovAgentLogsPayload(uint64(l))
	}
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovAgentLogsPayload(uint64(l))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovAgentLogsPayload(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovAgentLogsPayload(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAgentLogsPayload(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAgentLogsPayload(x uint64) (n int) {
	return sovAgentLogsPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Log) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgentLogsPayload
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
			return fmt.Errorf("proto: Log: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Log: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgentLogsPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgentLogsPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgentLogsPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hostname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgentLogsPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgentLogsPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgentLogsPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgentLogsPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAgentLogsPayload
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
func skipAgentLogsPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAgentLogsPayload
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
					return 0, ErrIntOverflowAgentLogsPayload
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
					return 0, ErrIntOverflowAgentLogsPayload
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
				return 0, ErrInvalidLengthAgentLogsPayload
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAgentLogsPayload
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAgentLogsPayload
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAgentLogsPayload        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgentLogsPayload          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAgentLogsPayload = fmt.Errorf("proto: unexpected end of group")
)
