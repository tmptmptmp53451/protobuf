// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3_proto/proto3.proto

package proto3_proto

import (
	fmt "fmt"
	proto "github.com/tmptmptmp53451/protobuf/proto"
	test_proto "github.com/tmptmptmp53451/protobuf/proto/test_proto"
	any "github.com/tmptmptmp53451/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message_Humour int32

const (
	Message_UNKNOWN     Message_Humour = 0
	Message_PUNS        Message_Humour = 1
	Message_SLAPSTICK   Message_Humour = 2
	Message_BILL_BAILEY Message_Humour = 3
)

var Message_Humour_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUNS",
	2: "SLAPSTICK",
	3: "BILL_BAILEY",
}

var Message_Humour_value = map[string]int32{
	"UNKNOWN":     0,
	"PUNS":        1,
	"SLAPSTICK":   2,
	"BILL_BAILEY": 3,
}

func (x Message_Humour) String() string {
	return proto.EnumName(Message_Humour_name, int32(x))
}

func (Message_Humour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{0, 0}
}

type Message struct {
	Name                 string                             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hilarity             Message_Humour                     `protobuf:"varint,2,opt,name=hilarity,proto3,enum=proto3_proto.Message_Humour" json:"hilarity,omitempty"`
	HeightInCm           uint32                             `protobuf:"varint,3,opt,name=height_in_cm,json=heightInCm,proto3" json:"height_in_cm,omitempty"`
	Data                 []byte                             `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ResultCount          int64                              `protobuf:"varint,7,opt,name=result_count,json=resultCount,proto3" json:"result_count,omitempty"`
	TrueScotsman         bool                               `protobuf:"varint,8,opt,name=true_scotsman,json=trueScotsman,proto3" json:"true_scotsman,omitempty"`
	Score                float32                            `protobuf:"fixed32,9,opt,name=score,proto3" json:"score,omitempty"`
	Key                  []uint64                           `protobuf:"varint,5,rep,packed,name=key,proto3" json:"key,omitempty"`
	ShortKey             []int32                            `protobuf:"varint,19,rep,packed,name=short_key,json=shortKey,proto3" json:"short_key,omitempty"`
	Nested               *Nested                            `protobuf:"bytes,6,opt,name=nested,proto3" json:"nested,omitempty"`
	RFunny               []Message_Humour                   `protobuf:"varint,16,rep,packed,name=r_funny,json=rFunny,proto3,enum=proto3_proto.Message_Humour" json:"r_funny,omitempty"`
	Terrain              map[string]*Nested                 `protobuf:"bytes,10,rep,name=terrain,proto3" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Proto2Field          *test_proto.SubDefaults            `protobuf:"bytes,11,opt,name=proto2_field,json=proto2Field,proto3" json:"proto2_field,omitempty"`
	Proto2Value          map[string]*test_proto.SubDefaults `protobuf:"bytes,13,rep,name=proto2_value,json=proto2Value,proto3" json:"proto2_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Anything             *any.Any                           `protobuf:"bytes,14,opt,name=anything,proto3" json:"anything,omitempty"`
	ManyThings           []*any.Any                         `protobuf:"bytes,15,rep,name=many_things,json=manyThings,proto3" json:"many_things,omitempty"`
	Submessage           *Message                           `protobuf:"bytes,17,opt,name=submessage,proto3" json:"submessage,omitempty"`
	Children             []*Message                         `protobuf:"bytes,18,rep,name=children,proto3" json:"children,omitempty"`
	StringMap            map[string]string                  `protobuf:"bytes,20,rep,name=string_map,json=stringMap,proto3" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetHilarity() Message_Humour {
	if m != nil {
		return m.Hilarity
	}
	return Message_UNKNOWN
}

func (m *Message) GetHeightInCm() uint32 {
	if m != nil {
		return m.HeightInCm
	}
	return 0
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetResultCount() int64 {
	if m != nil {
		return m.ResultCount
	}
	return 0
}

func (m *Message) GetTrueScotsman() bool {
	if m != nil {
		return m.TrueScotsman
	}
	return false
}

func (m *Message) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Message) GetKey() []uint64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetShortKey() []int32 {
	if m != nil {
		return m.ShortKey
	}
	return nil
}

func (m *Message) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Message) GetRFunny() []Message_Humour {
	if m != nil {
		return m.RFunny
	}
	return nil
}

func (m *Message) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

func (m *Message) GetProto2Field() *test_proto.SubDefaults {
	if m != nil {
		return m.Proto2Field
	}
	return nil
}

func (m *Message) GetProto2Value() map[string]*test_proto.SubDefaults {
	if m != nil {
		return m.Proto2Value
	}
	return nil
}

func (m *Message) GetAnything() *any.Any {
	if m != nil {
		return m.Anything
	}
	return nil
}

func (m *Message) GetManyThings() []*any.Any {
	if m != nil {
		return m.ManyThings
	}
	return nil
}

func (m *Message) GetSubmessage() *Message {
	if m != nil {
		return m.Submessage
	}
	return nil
}

func (m *Message) GetChildren() []*Message {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Message) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

type Nested struct {
	Bunny                string   `protobuf:"bytes,1,opt,name=bunny,proto3" json:"bunny,omitempty"`
	Cute                 bool     `protobuf:"varint,2,opt,name=cute,proto3" json:"cute,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()    {}
func (*Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{1}
}

func (m *Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested.Unmarshal(m, b)
}
func (m *Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested.Marshal(b, m, deterministic)
}
func (m *Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested.Merge(m, src)
}
func (m *Nested) XXX_Size() int {
	return xxx_messageInfo_Nested.Size(m)
}
func (m *Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Nested proto.InternalMessageInfo

func (m *Nested) GetBunny() string {
	if m != nil {
		return m.Bunny
	}
	return ""
}

func (m *Nested) GetCute() bool {
	if m != nil {
		return m.Cute
	}
	return false
}

type MessageWithMap struct {
	ByteMapping          map[bool][]byte `protobuf:"bytes,1,rep,name=byte_mapping,json=byteMapping,proto3" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MessageWithMap) Reset()         { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()    {}
func (*MessageWithMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{2}
}

func (m *MessageWithMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithMap.Unmarshal(m, b)
}
func (m *MessageWithMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithMap.Marshal(b, m, deterministic)
}
func (m *MessageWithMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithMap.Merge(m, src)
}
func (m *MessageWithMap) XXX_Size() int {
	return xxx_messageInfo_MessageWithMap.Size(m)
}
func (m *MessageWithMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithMap.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithMap proto.InternalMessageInfo

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

type IntMap struct {
	Rtt                  map[int32]int32 `protobuf:"bytes,1,rep,name=rtt,proto3" json:"rtt,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IntMap) Reset()         { *m = IntMap{} }
func (m *IntMap) String() string { return proto.CompactTextString(m) }
func (*IntMap) ProtoMessage()    {}
func (*IntMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{3}
}

func (m *IntMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMap.Unmarshal(m, b)
}
func (m *IntMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMap.Marshal(b, m, deterministic)
}
func (m *IntMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMap.Merge(m, src)
}
func (m *IntMap) XXX_Size() int {
	return xxx_messageInfo_IntMap.Size(m)
}
func (m *IntMap) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMap.DiscardUnknown(m)
}

var xxx_messageInfo_IntMap proto.InternalMessageInfo

func (m *IntMap) GetRtt() map[int32]int32 {
	if m != nil {
		return m.Rtt
	}
	return nil
}

type IntMaps struct {
	Maps                 []*IntMap `protobuf:"bytes,1,rep,name=maps,proto3" json:"maps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IntMaps) Reset()         { *m = IntMaps{} }
func (m *IntMaps) String() string { return proto.CompactTextString(m) }
func (*IntMaps) ProtoMessage()    {}
func (*IntMaps) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{4}
}

func (m *IntMaps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMaps.Unmarshal(m, b)
}
func (m *IntMaps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMaps.Marshal(b, m, deterministic)
}
func (m *IntMaps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMaps.Merge(m, src)
}
func (m *IntMaps) XXX_Size() int {
	return xxx_messageInfo_IntMaps.Size(m)
}
func (m *IntMaps) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMaps.DiscardUnknown(m)
}

var xxx_messageInfo_IntMaps proto.InternalMessageInfo

func (m *IntMaps) GetMaps() []*IntMap {
	if m != nil {
		return m.Maps
	}
	return nil
}

type TestUTF8 struct {
	Scalar string   `protobuf:"bytes,1,opt,name=scalar,proto3" json:"scalar,omitempty"`
	Vector []string `protobuf:"bytes,2,rep,name=vector,proto3" json:"vector,omitempty"`
	// Types that are valid to be assigned to Oneof:
	//	*TestUTF8_Field
	Oneof                isTestUTF8_Oneof `protobuf_oneof:"oneof"`
	MapKey               map[string]int64 `protobuf:"bytes,4,rep,name=map_key,json=mapKey,proto3" json:"map_key,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapValue             map[int64]string `protobuf:"bytes,5,rep,name=map_value,json=mapValue,proto3" json:"map_value,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TestUTF8) Reset()         { *m = TestUTF8{} }
func (m *TestUTF8) String() string { return proto.CompactTextString(m) }
func (*TestUTF8) ProtoMessage()    {}
func (*TestUTF8) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{5}
}

func (m *TestUTF8) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestUTF8.Unmarshal(m, b)
}
func (m *TestUTF8) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestUTF8.Marshal(b, m, deterministic)
}
func (m *TestUTF8) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestUTF8.Merge(m, src)
}
func (m *TestUTF8) XXX_Size() int {
	return xxx_messageInfo_TestUTF8.Size(m)
}
func (m *TestUTF8) XXX_DiscardUnknown() {
	xxx_messageInfo_TestUTF8.DiscardUnknown(m)
}

var xxx_messageInfo_TestUTF8 proto.InternalMessageInfo

func (m *TestUTF8) GetScalar() string {
	if m != nil {
		return m.Scalar
	}
	return ""
}

func (m *TestUTF8) GetVector() []string {
	if m != nil {
		return m.Vector
	}
	return nil
}

type isTestUTF8_Oneof interface {
	isTestUTF8_Oneof()
}

type TestUTF8_Field struct {
	Field string `protobuf:"bytes,3,opt,name=field,proto3,oneof"`
}

func (*TestUTF8_Field) isTestUTF8_Oneof() {}

func (m *TestUTF8) GetOneof() isTestUTF8_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (m *TestUTF8) GetField() string {
	if x, ok := m.GetOneof().(*TestUTF8_Field); ok {
		return x.Field
	}
	return ""
}

func (m *TestUTF8) GetMapKey() map[string]int64 {
	if m != nil {
		return m.MapKey
	}
	return nil
}

func (m *TestUTF8) GetMapValue() map[int64]string {
	if m != nil {
		return m.MapValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestUTF8) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestUTF8_Field)(nil),
	}
}

func init() {
	proto.RegisterEnum("proto3_proto.Message_Humour", Message_Humour_name, Message_Humour_value)
	proto.RegisterType((*Message)(nil), "proto3_proto.Message")
	proto.RegisterMapType((map[string]*test_proto.SubDefaults)(nil), "proto3_proto.Message.Proto2ValueEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto3_proto.Message.StringMapEntry")
	proto.RegisterMapType((map[string]*Nested)(nil), "proto3_proto.Message.TerrainEntry")
	proto.RegisterType((*Nested)(nil), "proto3_proto.Nested")
	proto.RegisterType((*MessageWithMap)(nil), "proto3_proto.MessageWithMap")
	proto.RegisterMapType((map[bool][]byte)(nil), "proto3_proto.MessageWithMap.ByteMappingEntry")
	proto.RegisterType((*IntMap)(nil), "proto3_proto.IntMap")
	proto.RegisterMapType((map[int32]int32)(nil), "proto3_proto.IntMap.RttEntry")
	proto.RegisterType((*IntMaps)(nil), "proto3_proto.IntMaps")
	proto.RegisterType((*TestUTF8)(nil), "proto3_proto.TestUTF8")
	proto.RegisterMapType((map[string]int64)(nil), "proto3_proto.TestUTF8.MapKeyEntry")
	proto.RegisterMapType((map[int64]string)(nil), "proto3_proto.TestUTF8.MapValueEntry")
}

func init() {
	proto.RegisterFile("proto3_proto/proto3.proto", fileDescriptor_1c50d9b824d4ac38)
}

var fileDescriptor_1c50d9b824d4ac38 = []byte{
	// 896 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x6f, 0x6f, 0xdb, 0xb6,
	0x13, 0xae, 0x2c, 0xff, 0x91, 0xcf, 0x76, 0xea, 0x1f, 0x7f, 0x6e, 0xc7, 0x7a, 0x1b, 0xa0, 0x79,
	0xc3, 0x20, 0x0c, 0xab, 0xb2, 0xb9, 0xc8, 0x90, 0xb5, 0xc5, 0x86, 0x24, 0x6b, 0x50, 0x23, 0xb1,
	0x67, 0xd0, 0xce, 0x82, 0xbd, 0x12, 0x68, 0x87, 0xb6, 0x85, 0x59, 0x94, 0x27, 0x52, 0x05, 0xf4,
	0x05, 0xf6, 0x41, 0xf6, 0x95, 0xf6, 0x85, 0x06, 0x92, 0x72, 0x2a, 0x17, 0xea, 0xf2, 0x4a, 0xbc,
	0x47, 0xcf, 0xdd, 0x73, 0xbc, 0x3b, 0x1e, 0x3c, 0xdb, 0x25, 0xb1, 0x8c, 0x5f, 0x04, 0xfa, 0x73,
	0x6c, 0x0c, 0x5f, 0x7f, 0x50, 0xbb, 0xf8, 0xab, 0xff, 0x6c, 0x1d, 0xc7, 0xeb, 0x2d, 0x33, 0x94,
	0x45, 0xba, 0x3a, 0xa6, 0x3c, 0x33, 0xc4, 0xfe, 0x13, 0xc9, 0x84, 0xcc, 0x23, 0xa8, 0xa3, 0x81,
	0x07, 0x7f, 0x35, 0xa1, 0x31, 0x66, 0x42, 0xd0, 0x35, 0x43, 0x08, 0xaa, 0x9c, 0x46, 0x0c, 0x5b,
	0xae, 0xe5, 0x35, 0x89, 0x3e, 0xa3, 0x53, 0x70, 0x36, 0xe1, 0x96, 0x26, 0xa1, 0xcc, 0x70, 0xc5,
	0xb5, 0xbc, 0xa3, 0xe1, 0x67, 0x7e, 0x51, 0xd2, 0xcf, 0x9d, 0xfd, 0xb7, 0x69, 0x14, 0xa7, 0x09,
	0xb9, 0x67, 0x23, 0x17, 0xda, 0x1b, 0x16, 0xae, 0x37, 0x32, 0x08, 0x79, 0xb0, 0x8c, 0xb0, 0xed,
	0x5a, 0x5e, 0x87, 0x80, 0xc1, 0x46, 0xfc, 0x22, 0x52, 0x7a, 0x77, 0x54, 0x52, 0x5c, 0x75, 0x2d,
	0xaf, 0x4d, 0xf4, 0x19, 0x7d, 0x01, 0xed, 0x84, 0x89, 0x74, 0x2b, 0x83, 0x65, 0x9c, 0x72, 0x89,
	0x1b, 0xae, 0xe5, 0xd9, 0xa4, 0x65, 0xb0, 0x0b, 0x05, 0xa1, 0x2f, 0xa1, 0x23, 0x93, 0x94, 0x05,
	0x62, 0x19, 0x4b, 0x11, 0x51, 0x8e, 0x1d, 0xd7, 0xf2, 0x1c, 0xd2, 0x56, 0xe0, 0x2c, 0xc7, 0x50,
	0x0f, 0x6a, 0x62, 0x19, 0x27, 0x0c, 0x37, 0x5d, 0xcb, 0xab, 0x10, 0x63, 0xa0, 0x2e, 0xd8, 0x7f,
	0xb0, 0x0c, 0xd7, 0x5c, 0xdb, 0xab, 0x12, 0x75, 0x44, 0x9f, 0x42, 0x53, 0x6c, 0xe2, 0x44, 0x06,
	0x0a, 0xff, 0xbf, 0x6b, 0x7b, 0x35, 0xe2, 0x68, 0xe0, 0x8a, 0x65, 0xe8, 0x5b, 0xa8, 0x73, 0x26,
	0x24, 0xbb, 0xc3, 0x75, 0xd7, 0xf2, 0x5a, 0xc3, 0xde, 0xe1, 0xd5, 0x27, 0xfa, 0x1f, 0xc9, 0x39,
	0xe8, 0x04, 0x1a, 0x49, 0xb0, 0x4a, 0x39, 0xcf, 0x70, 0xd7, 0xb5, 0x1f, 0xac, 0x54, 0x3d, 0xb9,
	0x54, 0x5c, 0xf4, 0x1a, 0x1a, 0x92, 0x25, 0x09, 0x0d, 0x39, 0x06, 0xd7, 0xf6, 0x5a, 0xc3, 0x41,
	0xb9, 0xdb, 0xdc, 0x90, 0xde, 0x70, 0x99, 0x64, 0x64, 0xef, 0x82, 0x5e, 0x82, 0x99, 0x80, 0x61,
	0xb0, 0x0a, 0xd9, 0xf6, 0x0e, 0xb7, 0x74, 0xa2, 0x9f, 0xf8, 0xef, 0xbb, 0xed, 0xcf, 0xd2, 0xc5,
	0x2f, 0x6c, 0x45, 0xd3, 0xad, 0x14, 0xa4, 0x65, 0xc8, 0x97, 0x8a, 0x8b, 0x46, 0xf7, 0xbe, 0xef,
	0xe8, 0x36, 0x65, 0xb8, 0xa3, 0xe5, 0xbf, 0x2e, 0x97, 0x9f, 0x6a, 0xe6, 0x6f, 0x8a, 0x68, 0x52,
	0xc8, 0x43, 0x69, 0x04, 0x7d, 0x07, 0x0e, 0xe5, 0x99, 0xdc, 0x84, 0x7c, 0x8d, 0x8f, 0xf2, 0x5a,
	0x99, 0x59, 0xf4, 0xf7, 0xb3, 0xe8, 0x9f, 0xf1, 0x8c, 0xdc, 0xb3, 0xd0, 0x09, 0xb4, 0x22, 0xca,
	0xb3, 0x40, 0x5b, 0x02, 0x3f, 0xd6, 0xda, 0xe5, 0x4e, 0xa0, 0x88, 0x73, 0xcd, 0x43, 0x27, 0x00,
	0x22, 0x5d, 0x44, 0x26, 0x29, 0xfc, 0x3f, 0x2d, 0xf5, 0xa4, 0x34, 0x63, 0x52, 0x20, 0xa2, 0xef,
	0xc1, 0x59, 0x6e, 0xc2, 0xed, 0x5d, 0xc2, 0x38, 0x46, 0x5a, 0xea, 0x23, 0x4e, 0xf7, 0x34, 0x74,
	0x01, 0x20, 0x64, 0x12, 0xf2, 0x75, 0x10, 0xd1, 0x1d, 0xee, 0x69, 0xa7, 0xaf, 0xca, 0x6b, 0x33,
	0xd3, 0xbc, 0x31, 0xdd, 0x99, 0xca, 0x34, 0xc5, 0xde, 0xee, 0x4f, 0xa1, 0x5d, 0xec, 0xdb, 0x7e,
	0x00, 0xcd, 0x0b, 0xd3, 0x03, 0xf8, 0x0d, 0xd4, 0x4c, 0xf5, 0x2b, 0xff, 0x31, 0x62, 0x86, 0xf2,
	0xb2, 0x72, 0x6a, 0xf5, 0x6f, 0xa1, 0xfb, 0x61, 0x2b, 0x4a, 0xa2, 0x3e, 0x3f, 0x8c, 0xfa, 0xd1,
	0x79, 0x28, 0x04, 0x7e, 0x0d, 0x47, 0x87, 0xf7, 0x28, 0x09, 0xdb, 0x2b, 0x86, 0x6d, 0x16, 0xbc,
	0x07, 0x3f, 0x43, 0xdd, 0xcc, 0x35, 0x6a, 0x41, 0xe3, 0x66, 0x72, 0x35, 0xf9, 0xf5, 0x76, 0xd2,
	0x7d, 0x84, 0x1c, 0xa8, 0x4e, 0x6f, 0x26, 0xb3, 0xae, 0x85, 0x3a, 0xd0, 0x9c, 0x5d, 0x9f, 0x4d,
	0x67, 0xf3, 0xd1, 0xc5, 0x55, 0xb7, 0x82, 0x1e, 0x43, 0xeb, 0x7c, 0x74, 0x7d, 0x1d, 0x9c, 0x9f,
	0x8d, 0xae, 0xdf, 0xfc, 0xde, 0xb5, 0x07, 0x43, 0xa8, 0x9b, 0xcb, 0x2a, 0x91, 0x85, 0x7e, 0x45,
	0x46, 0xd8, 0x18, 0x6a, 0x59, 0x2c, 0x53, 0x69, 0x94, 0x1d, 0xa2, 0xcf, 0x83, 0xbf, 0x2d, 0x38,
	0xca, 0x7b, 0x70, 0x1b, 0xca, 0xcd, 0x98, 0xee, 0xd0, 0x14, 0xda, 0x8b, 0x4c, 0x32, 0xd5, 0xb3,
	0x9d, 0x1a, 0x46, 0x4b, 0xf7, 0xed, 0x79, 0x69, 0xdf, 0x72, 0x1f, 0xff, 0x3c, 0x93, 0x6c, 0x6c,
	0xf8, 0xf9, 0x68, 0x2f, 0xde, 0x23, 0xfd, 0x9f, 0xa0, 0xfb, 0x21, 0xa1, 0x58, 0x19, 0xa7, 0xa4,
	0x32, 0xed, 0x62, 0x65, 0xfe, 0x84, 0xfa, 0x88, 0x4b, 0x95, 0xdb, 0x31, 0xd8, 0x89, 0x94, 0x79,
	0x4a, 0x9f, 0x1f, 0xa6, 0x64, 0x28, 0x3e, 0x91, 0xd2, 0xa4, 0xa0, 0x98, 0xfd, 0x1f, 0xc0, 0xd9,
	0x03, 0x45, 0xc9, 0x5a, 0x89, 0x64, 0xad, 0x28, 0xf9, 0x02, 0x1a, 0x26, 0x9e, 0x40, 0x1e, 0x54,
	0x23, 0xba, 0x13, 0xb9, 0x68, 0xaf, 0x4c, 0x94, 0x68, 0xc6, 0xe0, 0x9f, 0x0a, 0x38, 0x73, 0x26,
	0xe4, 0xcd, 0xfc, 0xf2, 0x14, 0x3d, 0x85, 0xba, 0x58, 0xd2, 0x2d, 0x4d, 0xf2, 0x26, 0xe4, 0x96,
	0xc2, 0xdf, 0xb1, 0xa5, 0x8c, 0x13, 0x5c, 0x71, 0x6d, 0x85, 0x1b, 0x0b, 0x3d, 0x85, 0x9a, 0xd9,
	0x3f, 0x6a, 0xcb, 0x37, 0xdf, 0x3e, 0x22, 0xc6, 0x44, 0xaf, 0xa0, 0x11, 0xd1, 0x9d, 0x5e, 0xae,
	0xd5, 0xb2, 0xe5, 0xb6, 0x17, 0xf4, 0xc7, 0x74, 0x77, 0xc5, 0x32, 0x73, 0xf7, 0x7a, 0xa4, 0x0d,
	0x74, 0x06, 0x4d, 0xe5, 0x6c, 0x2e, 0x59, 0x2b, 0x7b, 0x80, 0x45, 0xf7, 0xc2, 0x6a, 0x72, 0xa2,
	0xdc, 0xec, 0xff, 0x08, 0xad, 0x42, 0xe4, 0x87, 0x26, 0xda, 0x2e, 0xbe, 0x87, 0x57, 0xd0, 0x39,
	0x88, 0x5a, 0x74, 0xb6, 0x1f, 0x78, 0x0e, 0xe7, 0x0d, 0xa8, 0xc5, 0x9c, 0xc5, 0xab, 0x45, 0xdd,
	0xe4, 0xfb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x74, 0x17, 0x7f, 0xc3, 0x07, 0x00, 0x00,
}
