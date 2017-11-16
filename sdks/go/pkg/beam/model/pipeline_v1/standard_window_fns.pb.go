// Code generated by protoc-gen-go. DO NOT EDIT.
// source: standard_window_fns.proto

package pipeline_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// beam:windowfn:fixed_windows:v0.1
type FixedWindowsPayload struct {
	Size   *google_protobuf1.Duration  `protobuf:"bytes,1,opt,name=size" json:"size,omitempty"`
	Offset *google_protobuf2.Timestamp `protobuf:"bytes,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *FixedWindowsPayload) Reset()                    { *m = FixedWindowsPayload{} }
func (m *FixedWindowsPayload) String() string            { return proto.CompactTextString(m) }
func (*FixedWindowsPayload) ProtoMessage()               {}
func (*FixedWindowsPayload) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *FixedWindowsPayload) GetSize() *google_protobuf1.Duration {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *FixedWindowsPayload) GetOffset() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Offset
	}
	return nil
}

// beam:windowfn:sliding_windows:v0.1
type SlidingWindowsPayload struct {
	Size   *google_protobuf1.Duration  `protobuf:"bytes,1,opt,name=size" json:"size,omitempty"`
	Offset *google_protobuf2.Timestamp `protobuf:"bytes,2,opt,name=offset" json:"offset,omitempty"`
	Period *google_protobuf1.Duration  `protobuf:"bytes,3,opt,name=period" json:"period,omitempty"`
}

func (m *SlidingWindowsPayload) Reset()                    { *m = SlidingWindowsPayload{} }
func (m *SlidingWindowsPayload) String() string            { return proto.CompactTextString(m) }
func (*SlidingWindowsPayload) ProtoMessage()               {}
func (*SlidingWindowsPayload) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *SlidingWindowsPayload) GetSize() *google_protobuf1.Duration {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *SlidingWindowsPayload) GetOffset() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *SlidingWindowsPayload) GetPeriod() *google_protobuf1.Duration {
	if m != nil {
		return m.Period
	}
	return nil
}

// beam:windowfn:session_windows:v0.1
type SessionsPayload struct {
	GapSize *google_protobuf1.Duration `protobuf:"bytes,1,opt,name=gap_size,json=gapSize" json:"gap_size,omitempty"`
}

func (m *SessionsPayload) Reset()                    { *m = SessionsPayload{} }
func (m *SessionsPayload) String() string            { return proto.CompactTextString(m) }
func (*SessionsPayload) ProtoMessage()               {}
func (*SessionsPayload) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *SessionsPayload) GetGapSize() *google_protobuf1.Duration {
	if m != nil {
		return m.GapSize
	}
	return nil
}

func init() {
	proto.RegisterType((*FixedWindowsPayload)(nil), "org.apache.beam.model.pipeline.v1.FixedWindowsPayload")
	proto.RegisterType((*SlidingWindowsPayload)(nil), "org.apache.beam.model.pipeline.v1.SlidingWindowsPayload")
	proto.RegisterType((*SessionsPayload)(nil), "org.apache.beam.model.pipeline.v1.SessionsPayload")
}

func init() { proto.RegisterFile("standard_window_fns.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0xa9, 0x4a, 0x95, 0xed, 0x41, 0x8c, 0x08, 0x69, 0x0e, 0x7e, 0xe4, 0xe4, 0xc5, 0x2d,
	0xa9, 0xfe, 0x82, 0x2a, 0xf5, 0x2a, 0x8d, 0x20, 0x78, 0x09, 0x13, 0x77, 0xb2, 0x0e, 0x24, 0x3b,
	0x4b, 0x76, 0xfb, 0x61, 0xff, 0x93, 0xff, 0x51, 0xc8, 0x87, 0x07, 0x3d, 0xd4, 0x93, 0xe7, 0x79,
	0xde, 0xf7, 0x19, 0x76, 0x47, 0x8c, 0x9d, 0x07, 0xa3, 0xa0, 0x56, 0xd9, 0x9a, 0x8c, 0xe2, 0x75,
	0x56, 0x18, 0x27, 0x6d, 0xcd, 0x9e, 0x83, 0x2b, 0xae, 0xb5, 0x04, 0x0b, 0x6f, 0xef, 0x28, 0x73,
	0x84, 0x4a, 0x56, 0xac, 0xb0, 0x94, 0x96, 0x2c, 0x96, 0x64, 0x50, 0xae, 0x92, 0xe8, 0x5c, 0x33,
	0xeb, 0x12, 0x27, 0x4d, 0x20, 0x5f, 0x16, 0x13, 0xb5, 0xac, 0xc1, 0x13, 0x9b, 0xb6, 0x22, 0xba,
	0xf8, 0x39, 0xf7, 0x54, 0xa1, 0xf3, 0x50, 0xd9, 0x16, 0x88, 0x37, 0xe2, 0x74, 0x4e, 0x1b, 0x54,
	0x2f, 0x8d, 0xdc, 0x3d, 0xc1, 0x47, 0xc9, 0xa0, 0x82, 0x1b, 0x71, 0xe0, 0x68, 0x8b, 0xe1, 0xe0,
	0x72, 0x70, 0x3d, 0x9a, 0x8e, 0x65, 0x5b, 0x23, 0xfb, 0x1a, 0xf9, 0xd0, 0x69, 0x16, 0x0d, 0x16,
	0x4c, 0xc5, 0x90, 0x8b, 0xc2, 0xa1, 0x0f, 0xf7, 0x9a, 0x40, 0xf4, 0x2b, 0xf0, 0xdc, 0x7b, 0x17,
	0x1d, 0x19, 0x7f, 0x0e, 0xc4, 0x59, 0x5a, 0x92, 0x22, 0xa3, 0xff, 0x5d, 0x1e, 0x24, 0x62, 0x68,
	0xb1, 0x26, 0x56, 0xe1, 0xfe, 0x2e, 0x49, 0x07, 0xc6, 0x8f, 0xe2, 0x38, 0x45, 0xe7, 0x88, 0xcd,
	0xf7, 0xa2, 0x77, 0xe2, 0x48, 0x83, 0xcd, 0xfe, 0xb6, 0xec, 0xa1, 0x06, 0x9b, 0xd2, 0x16, 0x67,
	0xf7, 0x62, 0xf7, 0xc7, 0xce, 0x4e, 0xd2, 0xee, 0x2c, 0xda, 0xb7, 0x99, 0x1b, 0xf7, 0x3a, 0xea,
	0xe7, 0xd9, 0x2a, 0xc9, 0x87, 0x8d, 0xe0, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x82, 0x58, 0x88,
	0x91, 0x3f, 0x02, 0x00, 0x00,
}