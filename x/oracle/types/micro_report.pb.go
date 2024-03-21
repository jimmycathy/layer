// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/oracle/micro_report.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MicroReport represents data for a single report
type MicroReport struct {
	// reporter is the address of the reporter
	Reporter string `protobuf:"bytes,1,opt,name=reporter,proto3" json:"reporter,omitempty"`
	// the power of the reporter based on total tokens normalized
	Power int64 `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
	// string identifier of the data spec
	QueryType string `protobuf:"bytes,3,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	// hash of the query data
	QueryId string `protobuf:"bytes,4,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	// aggregate method to use for aggregating all the reports for the query id
	AggregateMethod string `protobuf:"bytes,5,opt,name=aggregate_method,json=aggregateMethod,proto3" json:"aggregate_method,omitempty"`
	// hex string of the response value
	Value string `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	// timestamp of when the report was created
	Timestamp time.Time `protobuf:"bytes,7,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	// indicates if the report's query id is in the cyclelist
	Cyclelist bool `protobuf:"varint,8,opt,name=cyclelist,proto3" json:"cyclelist,omitempty"`
	// block number of when the report was created
	BlockNumber int64 `protobuf:"varint,9,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (m *MicroReport) Reset()         { *m = MicroReport{} }
func (m *MicroReport) String() string { return proto.CompactTextString(m) }
func (*MicroReport) ProtoMessage()    {}
func (*MicroReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_c39350954f878191, []int{0}
}
func (m *MicroReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MicroReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MicroReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MicroReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MicroReport.Merge(m, src)
}
func (m *MicroReport) XXX_Size() int {
	return m.Size()
}
func (m *MicroReport) XXX_DiscardUnknown() {
	xxx_messageInfo_MicroReport.DiscardUnknown(m)
}

var xxx_messageInfo_MicroReport proto.InternalMessageInfo

func (m *MicroReport) GetReporter() string {
	if m != nil {
		return m.Reporter
	}
	return ""
}

func (m *MicroReport) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *MicroReport) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *MicroReport) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *MicroReport) GetAggregateMethod() string {
	if m != nil {
		return m.AggregateMethod
	}
	return ""
}

func (m *MicroReport) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *MicroReport) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *MicroReport) GetCyclelist() bool {
	if m != nil {
		return m.Cyclelist
	}
	return false
}

func (m *MicroReport) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*MicroReport)(nil), "layer.oracle.MicroReport")
}

func init() { proto.RegisterFile("layer/oracle/micro_report.proto", fileDescriptor_c39350954f878191) }

var fileDescriptor_c39350954f878191 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x51, 0xbd, 0x6e, 0xdb, 0x30,
	0x10, 0x16, 0xed, 0xda, 0x96, 0x68, 0x03, 0x2d, 0x08, 0x0f, 0xac, 0xd0, 0x4a, 0x6a, 0x27, 0x79,
	0xa8, 0x04, 0xb4, 0x6f, 0xe0, 0x4e, 0x19, 0x9c, 0x41, 0xf0, 0x94, 0x45, 0x90, 0xe4, 0x0b, 0x2d,
	0x84, 0x0a, 0x15, 0x9a, 0x4a, 0xa2, 0xb7, 0xf0, 0x1b, 0x65, 0xf5, 0xe8, 0x31, 0x53, 0x12, 0xd8,
	0x2f, 0x12, 0x88, 0x84, 0xed, 0x8d, 0xdf, 0xcf, 0xdd, 0xf1, 0xbb, 0xc3, 0x3e, 0xcf, 0x5a, 0x90,
	0xb1, 0x90, 0x59, 0xc1, 0x21, 0xae, 0xca, 0x42, 0x8a, 0x54, 0x42, 0x2d, 0xa4, 0x8a, 0x6a, 0x29,
	0x94, 0x20, 0x13, 0x6d, 0x88, 0x8c, 0xc1, 0x9d, 0x32, 0xc1, 0x84, 0x16, 0xe2, 0xee, 0x65, 0x3c,
	0xae, 0xcf, 0x84, 0x60, 0x1c, 0x62, 0x8d, 0xf2, 0xe6, 0x36, 0x56, 0x65, 0x05, 0x1b, 0x95, 0x55,
	0xb5, 0x31, 0xfc, 0x7e, 0xe9, 0xe1, 0xf1, 0xa2, 0xeb, 0x9d, 0xe8, 0xd6, 0xc4, 0xc5, 0xb6, 0x19,
	0x02, 0x92, 0xa2, 0x00, 0x85, 0x4e, 0x72, 0xc6, 0x64, 0x8a, 0x07, 0xb5, 0x78, 0x02, 0x49, 0x7b,
	0x01, 0x0a, 0xfb, 0x89, 0x01, 0xe4, 0x27, 0xc6, 0x0f, 0x0d, 0xc8, 0x36, 0x55, 0x6d, 0x0d, 0xb4,
	0xaf, 0x6b, 0x1c, 0xcd, 0x2c, 0xdb, 0x1a, 0xc8, 0x77, 0x6c, 0x1b, 0xb9, 0x5c, 0xd1, 0x2f, 0x5a,
	0x1c, 0x69, 0x7c, 0xb5, 0x22, 0x33, 0xfc, 0x2d, 0x63, 0x4c, 0x02, 0xcb, 0x14, 0xa4, 0x15, 0xa8,
	0xb5, 0x58, 0xd1, 0x81, 0xb6, 0x7c, 0x3d, 0xf3, 0x0b, 0x4d, 0x77, 0xa3, 0x1f, 0x33, 0xde, 0x00,
	0x1d, 0x6a, 0xdd, 0x00, 0x32, 0xc7, 0xce, 0x39, 0x0f, 0x1d, 0x05, 0x28, 0x1c, 0xff, 0x75, 0x23,
	0x93, 0x38, 0x3a, 0x25, 0x8e, 0x96, 0x27, 0xc7, 0xdc, 0xde, 0xbd, 0xf9, 0xd6, 0xf6, 0xdd, 0x47,
	0xc9, 0xa5, 0x8c, 0xfc, 0xc0, 0x4e, 0xd1, 0x16, 0x1c, 0x78, 0xb9, 0x51, 0xd4, 0x0e, 0x50, 0x68,
	0x27, 0x17, 0x82, 0xfc, 0xc2, 0x93, 0x9c, 0x8b, 0xe2, 0x2e, 0xbd, 0x6f, 0xaa, 0x1c, 0x24, 0x75,
	0x74, 0xf2, 0xb1, 0xe6, 0xae, 0x35, 0x35, 0xff, 0xbf, 0x3b, 0x78, 0x68, 0x7f, 0xf0, 0xd0, 0xc7,
	0xc1, 0x43, 0xdb, 0xa3, 0x67, 0xed, 0x8f, 0x9e, 0xf5, 0x7a, 0xf4, 0xac, 0x9b, 0x19, 0x2b, 0xd5,
	0xba, 0xc9, 0xa3, 0x42, 0x54, 0xb1, 0x02, 0xce, 0x85, 0xfc, 0x53, 0x8a, 0xd8, 0x9c, 0xf5, 0xf9,
	0x74, 0xd8, 0x6e, 0x6b, 0x9b, 0x7c, 0xa8, 0xbf, 0xfb, 0xef, 0x33, 0x00, 0x00, 0xff, 0xff, 0x0f,
	0x39, 0x4c, 0x75, 0xf5, 0x01, 0x00, 0x00,
}

func (m *MicroReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MicroReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MicroReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockNumber != 0 {
		i = encodeVarintMicroReport(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x48
	}
	if m.Cyclelist {
		i--
		if m.Cyclelist {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintMicroReport(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintMicroReport(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.AggregateMethod) > 0 {
		i -= len(m.AggregateMethod)
		copy(dAtA[i:], m.AggregateMethod)
		i = encodeVarintMicroReport(dAtA, i, uint64(len(m.AggregateMethod)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.QueryId) > 0 {
		i -= len(m.QueryId)
		copy(dAtA[i:], m.QueryId)
		i = encodeVarintMicroReport(dAtA, i, uint64(len(m.QueryId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintMicroReport(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Power != 0 {
		i = encodeVarintMicroReport(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Reporter) > 0 {
		i -= len(m.Reporter)
		copy(dAtA[i:], m.Reporter)
		i = encodeVarintMicroReport(dAtA, i, uint64(len(m.Reporter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMicroReport(dAtA []byte, offset int, v uint64) int {
	offset -= sovMicroReport(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MicroReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Reporter)
	if l > 0 {
		n += 1 + l + sovMicroReport(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovMicroReport(uint64(m.Power))
	}
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovMicroReport(uint64(l))
	}
	l = len(m.QueryId)
	if l > 0 {
		n += 1 + l + sovMicroReport(uint64(l))
	}
	l = len(m.AggregateMethod)
	if l > 0 {
		n += 1 + l + sovMicroReport(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovMicroReport(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovMicroReport(uint64(l))
	if m.Cyclelist {
		n += 2
	}
	if m.BlockNumber != 0 {
		n += 1 + sovMicroReport(uint64(m.BlockNumber))
	}
	return n
}

func sovMicroReport(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMicroReport(x uint64) (n int) {
	return sovMicroReport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MicroReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMicroReport
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
			return fmt.Errorf("proto: MicroReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MicroReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
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
				return ErrInvalidLengthMicroReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMicroReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
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
				return ErrInvalidLengthMicroReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMicroReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
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
				return ErrInvalidLengthMicroReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMicroReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
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
				return ErrInvalidLengthMicroReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMicroReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
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
				return ErrInvalidLengthMicroReport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMicroReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
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
				return ErrInvalidLengthMicroReport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMicroReport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cyclelist", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
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
			m.Cyclelist = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMicroReport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMicroReport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMicroReport
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
func skipMicroReport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMicroReport
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
					return 0, ErrIntOverflowMicroReport
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
					return 0, ErrIntOverflowMicroReport
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
				return 0, ErrInvalidLengthMicroReport
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMicroReport
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMicroReport
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMicroReport        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMicroReport          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMicroReport = fmt.Errorf("proto: unexpected end of group")
)
