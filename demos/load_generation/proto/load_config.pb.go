// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: demos/load_generation/proto/load_config.proto

package load_generation

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type LocustConfigFile struct {
	Phases []*LocustPhaseConfig `protobuf:"bytes,1,rep,name=phases,proto3" json:"phases,omitempty"`
}

func (m *LocustConfigFile) Reset()      { *m = LocustConfigFile{} }
func (*LocustConfigFile) ProtoMessage() {}
func (*LocustConfigFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_d354fe8a1f1f8d15, []int{0}
}
func (m *LocustConfigFile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocustConfigFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocustConfigFile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocustConfigFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocustConfigFile.Merge(m, src)
}
func (m *LocustConfigFile) XXX_Size() int {
	return m.Size()
}
func (m *LocustConfigFile) XXX_DiscardUnknown() {
	xxx_messageInfo_LocustConfigFile.DiscardUnknown(m)
}

var xxx_messageInfo_LocustConfigFile proto.InternalMessageInfo

func (m *LocustConfigFile) GetPhases() []*LocustPhaseConfig {
	if m != nil {
		return m.Phases
	}
	return nil
}

type LocustPhaseConfig struct {
	UserTypes     []*UserType `protobuf:"bytes,1,rep,name=user_types,json=userTypes,proto3" json:"user_types,omitempty"`
	Duration      int32       `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	NumUsers      int32       `protobuf:"varint,3,opt,name=num_users,json=numUsers,proto3" json:"num_users,omitempty"`
	HatchRatePerS int32       `protobuf:"varint,4,opt,name=hatch_rate_per_s,json=hatchRatePerS,proto3" json:"hatch_rate_per_s,omitempty"`
}

func (m *LocustPhaseConfig) Reset()      { *m = LocustPhaseConfig{} }
func (*LocustPhaseConfig) ProtoMessage() {}
func (*LocustPhaseConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d354fe8a1f1f8d15, []int{1}
}
func (m *LocustPhaseConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocustPhaseConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocustPhaseConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocustPhaseConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocustPhaseConfig.Merge(m, src)
}
func (m *LocustPhaseConfig) XXX_Size() int {
	return m.Size()
}
func (m *LocustPhaseConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LocustPhaseConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LocustPhaseConfig proto.InternalMessageInfo

func (m *LocustPhaseConfig) GetUserTypes() []*UserType {
	if m != nil {
		return m.UserTypes
	}
	return nil
}

func (m *LocustPhaseConfig) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *LocustPhaseConfig) GetNumUsers() int32 {
	if m != nil {
		return m.NumUsers
	}
	return 0
}

func (m *LocustPhaseConfig) GetHatchRatePerS() int32 {
	if m != nil {
		return m.HatchRatePerS
	}
	return 0
}

type UserType struct {
	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ratio     int32  `protobuf:"varint,2,opt,name=ratio,proto3" json:"ratio,omitempty"`
	MinWaitMs int32  `protobuf:"varint,3,opt,name=min_wait_ms,json=minWaitMs,proto3" json:"min_wait_ms,omitempty"`
	MaxWaitMs int32  `protobuf:"varint,4,opt,name=max_wait_ms,json=maxWaitMs,proto3" json:"max_wait_ms,omitempty"`
}

func (m *UserType) Reset()      { *m = UserType{} }
func (*UserType) ProtoMessage() {}
func (*UserType) Descriptor() ([]byte, []int) {
	return fileDescriptor_d354fe8a1f1f8d15, []int{2}
}
func (m *UserType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserType.Merge(m, src)
}
func (m *UserType) XXX_Size() int {
	return m.Size()
}
func (m *UserType) XXX_DiscardUnknown() {
	xxx_messageInfo_UserType.DiscardUnknown(m)
}

var xxx_messageInfo_UserType proto.InternalMessageInfo

func (m *UserType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserType) GetRatio() int32 {
	if m != nil {
		return m.Ratio
	}
	return 0
}

func (m *UserType) GetMinWaitMs() int32 {
	if m != nil {
		return m.MinWaitMs
	}
	return 0
}

func (m *UserType) GetMaxWaitMs() int32 {
	if m != nil {
		return m.MaxWaitMs
	}
	return 0
}

func init() {
	proto.RegisterType((*LocustConfigFile)(nil), "pl.demos.load_generation.LocustConfigFile")
	proto.RegisterType((*LocustPhaseConfig)(nil), "pl.demos.load_generation.LocustPhaseConfig")
	proto.RegisterType((*UserType)(nil), "pl.demos.load_generation.UserType")
}

func init() {
	proto.RegisterFile("demos/load_generation/proto/load_config.proto", fileDescriptor_d354fe8a1f1f8d15)
}

var fileDescriptor_d354fe8a1f1f8d15 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8a, 0xd4, 0x30,
	0x1c, 0xc6, 0x1b, 0x77, 0x76, 0xd9, 0x66, 0x11, 0xd6, 0xe0, 0xa1, 0xac, 0x10, 0x86, 0x5e, 0x1c,
	0x90, 0x6d, 0x41, 0x8f, 0x9e, 0x74, 0xc0, 0x93, 0xc2, 0x52, 0x95, 0x05, 0x2f, 0x21, 0xed, 0x64,
	0xdb, 0x40, 0x93, 0x94, 0x26, 0xc1, 0x99, 0x9b, 0x8f, 0xe0, 0x63, 0xf8, 0x08, 0x3e, 0x82, 0xc7,
	0x39, 0xce, 0xd1, 0xe9, 0x5c, 0x3c, 0xce, 0x23, 0x48, 0x93, 0x4e, 0x0f, 0x23, 0xe3, 0x2d, 0xdf,
	0xf7, 0xfd, 0xbf, 0xdf, 0x77, 0x09, 0xbc, 0x5d, 0x30, 0xa1, 0x74, 0x5a, 0x2b, 0xba, 0x20, 0x25,
	0x93, 0xac, 0xa5, 0x86, 0x2b, 0x99, 0x36, 0xad, 0x32, 0xca, 0xbb, 0x85, 0x92, 0x0f, 0xbc, 0x4c,
	0x9c, 0x83, 0xa2, 0xa6, 0x4e, 0x5c, 0x23, 0x39, 0x6a, 0xdc, 0xdc, 0x96, 0xdc, 0x54, 0x36, 0x4f,
	0x0a, 0x25, 0xd2, 0x52, 0x95, 0xca, 0x23, 0x72, 0xfb, 0xe0, 0x94, 0xe7, 0xf5, 0x2f, 0x0f, 0x8a,
	0xef, 0xe1, 0xf5, 0x7b, 0x55, 0x58, 0x6d, 0xe6, 0x0e, 0xff, 0x8e, 0xd7, 0x0c, 0xcd, 0xe1, 0x45,
	0x53, 0x51, 0xcd, 0x74, 0x04, 0xa6, 0x67, 0xb3, 0xab, 0x97, 0x2f, 0x92, 0x53, 0x6b, 0x89, 0xef,
	0xde, 0xf5, 0xd7, 0x1e, 0x90, 0x0d, 0xd5, 0xf8, 0x27, 0x80, 0x4f, 0xfe, 0x49, 0xd1, 0x1b, 0x08,
	0xad, 0x66, 0x2d, 0x31, 0xab, 0x66, 0xc4, 0xc7, 0xa7, 0xf1, 0x9f, 0x35, 0x6b, 0x3f, 0xad, 0x1a,
	0x96, 0x85, 0x76, 0x78, 0x69, 0x74, 0x03, 0x2f, 0x17, 0xd6, 0xe7, 0xd1, 0xa3, 0x29, 0x98, 0x9d,
	0x67, 0xa3, 0x46, 0xcf, 0x60, 0x28, 0xad, 0x20, 0xfd, 0xb1, 0x8e, 0xce, 0x7c, 0x28, 0xad, 0xe8,
	0x31, 0x1a, 0x3d, 0x87, 0xd7, 0x15, 0x35, 0x45, 0x45, 0x5a, 0x6a, 0x18, 0x69, 0x58, 0x4b, 0x74,
	0x34, 0x71, 0x37, 0x8f, 0x9d, 0x9f, 0x51, 0xc3, 0xee, 0x58, 0xfb, 0x31, 0x36, 0xf0, 0xf2, 0x30,
	0x8c, 0x10, 0x9c, 0x48, 0x2a, 0x58, 0x04, 0xa6, 0x60, 0x16, 0x66, 0xee, 0x8d, 0x9e, 0xc2, 0x73,
	0xb7, 0x37, 0xcc, 0x7b, 0x81, 0x30, 0xbc, 0x12, 0x5c, 0x92, 0xaf, 0x94, 0x1b, 0x22, 0x0e, 0xeb,
	0xa1, 0xe0, 0xf2, 0x9e, 0x72, 0xf3, 0x41, 0xbb, 0x9c, 0x2e, 0xc7, 0x7c, 0x32, 0xe4, 0x74, 0xe9,
	0xf3, 0xb7, 0xab, 0xf5, 0x16, 0x07, 0x9b, 0x2d, 0x0e, 0xf6, 0x5b, 0x0c, 0xbe, 0x75, 0x18, 0xfc,
	0xe8, 0x30, 0xf8, 0xd5, 0x61, 0xb0, 0xee, 0x30, 0xf8, 0xdd, 0x61, 0xf0, 0xa7, 0xc3, 0xc1, 0xbe,
	0xc3, 0xe0, 0xfb, 0x0e, 0x07, 0xeb, 0x1d, 0x0e, 0x36, 0x3b, 0x1c, 0x7c, 0x99, 0x37, 0x7c, 0xc9,
	0x59, 0x4d, 0x73, 0x9d, 0x50, 0x9e, 0x8e, 0x22, 0xfd, 0xdf, 0x77, 0x7a, 0x7d, 0x64, 0xe7, 0x17,
	0xce, 0x7f, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x06, 0xa9, 0x1f, 0x95, 0x85, 0x02, 0x00, 0x00,
}

func (this *LocustConfigFile) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LocustConfigFile)
	if !ok {
		that2, ok := that.(LocustConfigFile)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Phases) != len(that1.Phases) {
		return false
	}
	for i := range this.Phases {
		if !this.Phases[i].Equal(that1.Phases[i]) {
			return false
		}
	}
	return true
}
func (this *LocustPhaseConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LocustPhaseConfig)
	if !ok {
		that2, ok := that.(LocustPhaseConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.UserTypes) != len(that1.UserTypes) {
		return false
	}
	for i := range this.UserTypes {
		if !this.UserTypes[i].Equal(that1.UserTypes[i]) {
			return false
		}
	}
	if this.Duration != that1.Duration {
		return false
	}
	if this.NumUsers != that1.NumUsers {
		return false
	}
	if this.HatchRatePerS != that1.HatchRatePerS {
		return false
	}
	return true
}
func (this *UserType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserType)
	if !ok {
		that2, ok := that.(UserType)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Ratio != that1.Ratio {
		return false
	}
	if this.MinWaitMs != that1.MinWaitMs {
		return false
	}
	if this.MaxWaitMs != that1.MaxWaitMs {
		return false
	}
	return true
}
func (this *LocustConfigFile) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&load_generation.LocustConfigFile{")
	if this.Phases != nil {
		s = append(s, "Phases: "+fmt.Sprintf("%#v", this.Phases)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *LocustPhaseConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&load_generation.LocustPhaseConfig{")
	if this.UserTypes != nil {
		s = append(s, "UserTypes: "+fmt.Sprintf("%#v", this.UserTypes)+",\n")
	}
	s = append(s, "Duration: "+fmt.Sprintf("%#v", this.Duration)+",\n")
	s = append(s, "NumUsers: "+fmt.Sprintf("%#v", this.NumUsers)+",\n")
	s = append(s, "HatchRatePerS: "+fmt.Sprintf("%#v", this.HatchRatePerS)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UserType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&load_generation.UserType{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Ratio: "+fmt.Sprintf("%#v", this.Ratio)+",\n")
	s = append(s, "MinWaitMs: "+fmt.Sprintf("%#v", this.MinWaitMs)+",\n")
	s = append(s, "MaxWaitMs: "+fmt.Sprintf("%#v", this.MaxWaitMs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringLoadConfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *LocustConfigFile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocustConfigFile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocustConfigFile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Phases) > 0 {
		for iNdEx := len(m.Phases) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Phases[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLoadConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LocustPhaseConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocustPhaseConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocustPhaseConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HatchRatePerS != 0 {
		i = encodeVarintLoadConfig(dAtA, i, uint64(m.HatchRatePerS))
		i--
		dAtA[i] = 0x20
	}
	if m.NumUsers != 0 {
		i = encodeVarintLoadConfig(dAtA, i, uint64(m.NumUsers))
		i--
		dAtA[i] = 0x18
	}
	if m.Duration != 0 {
		i = encodeVarintLoadConfig(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UserTypes) > 0 {
		for iNdEx := len(m.UserTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLoadConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UserType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxWaitMs != 0 {
		i = encodeVarintLoadConfig(dAtA, i, uint64(m.MaxWaitMs))
		i--
		dAtA[i] = 0x20
	}
	if m.MinWaitMs != 0 {
		i = encodeVarintLoadConfig(dAtA, i, uint64(m.MinWaitMs))
		i--
		dAtA[i] = 0x18
	}
	if m.Ratio != 0 {
		i = encodeVarintLoadConfig(dAtA, i, uint64(m.Ratio))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintLoadConfig(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLoadConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovLoadConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LocustConfigFile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Phases) > 0 {
		for _, e := range m.Phases {
			l = e.Size()
			n += 1 + l + sovLoadConfig(uint64(l))
		}
	}
	return n
}

func (m *LocustPhaseConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.UserTypes) > 0 {
		for _, e := range m.UserTypes {
			l = e.Size()
			n += 1 + l + sovLoadConfig(uint64(l))
		}
	}
	if m.Duration != 0 {
		n += 1 + sovLoadConfig(uint64(m.Duration))
	}
	if m.NumUsers != 0 {
		n += 1 + sovLoadConfig(uint64(m.NumUsers))
	}
	if m.HatchRatePerS != 0 {
		n += 1 + sovLoadConfig(uint64(m.HatchRatePerS))
	}
	return n
}

func (m *UserType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLoadConfig(uint64(l))
	}
	if m.Ratio != 0 {
		n += 1 + sovLoadConfig(uint64(m.Ratio))
	}
	if m.MinWaitMs != 0 {
		n += 1 + sovLoadConfig(uint64(m.MinWaitMs))
	}
	if m.MaxWaitMs != 0 {
		n += 1 + sovLoadConfig(uint64(m.MaxWaitMs))
	}
	return n
}

func sovLoadConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLoadConfig(x uint64) (n int) {
	return sovLoadConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *LocustConfigFile) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForPhases := "[]*LocustPhaseConfig{"
	for _, f := range this.Phases {
		repeatedStringForPhases += strings.Replace(f.String(), "LocustPhaseConfig", "LocustPhaseConfig", 1) + ","
	}
	repeatedStringForPhases += "}"
	s := strings.Join([]string{`&LocustConfigFile{`,
		`Phases:` + repeatedStringForPhases + `,`,
		`}`,
	}, "")
	return s
}
func (this *LocustPhaseConfig) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForUserTypes := "[]*UserType{"
	for _, f := range this.UserTypes {
		repeatedStringForUserTypes += strings.Replace(f.String(), "UserType", "UserType", 1) + ","
	}
	repeatedStringForUserTypes += "}"
	s := strings.Join([]string{`&LocustPhaseConfig{`,
		`UserTypes:` + repeatedStringForUserTypes + `,`,
		`Duration:` + fmt.Sprintf("%v", this.Duration) + `,`,
		`NumUsers:` + fmt.Sprintf("%v", this.NumUsers) + `,`,
		`HatchRatePerS:` + fmt.Sprintf("%v", this.HatchRatePerS) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UserType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UserType{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Ratio:` + fmt.Sprintf("%v", this.Ratio) + `,`,
		`MinWaitMs:` + fmt.Sprintf("%v", this.MinWaitMs) + `,`,
		`MaxWaitMs:` + fmt.Sprintf("%v", this.MaxWaitMs) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringLoadConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *LocustConfigFile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoadConfig
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
			return fmt.Errorf("proto: LocustConfigFile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocustConfigFile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phases", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoadConfig
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
				return ErrInvalidLengthLoadConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoadConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phases = append(m.Phases, &LocustPhaseConfig{})
			if err := m.Phases[len(m.Phases)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoadConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoadConfig
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
func (m *LocustPhaseConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoadConfig
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
			return fmt.Errorf("proto: LocustPhaseConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocustPhaseConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserTypes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoadConfig
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
				return ErrInvalidLengthLoadConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoadConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserTypes = append(m.UserTypes, &UserType{})
			if err := m.UserTypes[len(m.UserTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoadConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumUsers", wireType)
			}
			m.NumUsers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoadConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumUsers |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HatchRatePerS", wireType)
			}
			m.HatchRatePerS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoadConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HatchRatePerS |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLoadConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoadConfig
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
func (m *UserType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoadConfig
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
			return fmt.Errorf("proto: UserType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoadConfig
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
				return ErrInvalidLengthLoadConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoadConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ratio", wireType)
			}
			m.Ratio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoadConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ratio |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinWaitMs", wireType)
			}
			m.MinWaitMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoadConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinWaitMs |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxWaitMs", wireType)
			}
			m.MaxWaitMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoadConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxWaitMs |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLoadConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoadConfig
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
func skipLoadConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoadConfig
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
					return 0, ErrIntOverflowLoadConfig
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
					return 0, ErrIntOverflowLoadConfig
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
				return 0, ErrInvalidLengthLoadConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLoadConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLoadConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLoadConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoadConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLoadConfig = fmt.Errorf("proto: unexpected end of group")
)