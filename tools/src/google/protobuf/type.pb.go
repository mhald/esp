// Code generated by protoc-gen-go.
// source: google/protobuf/type.proto
// DO NOT EDIT!

package descriptor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The syntax in which a protocol buffer element is defined.
type Syntax int32

const (
	// Syntax `proto2`.
	Syntax_SYNTAX_PROTO2 Syntax = 0
	// Syntax `proto3`.
	Syntax_SYNTAX_PROTO3 Syntax = 1
)

var Syntax_name = map[int32]string{
	0: "SYNTAX_PROTO2",
	1: "SYNTAX_PROTO3",
}
var Syntax_value = map[string]int32{
	"SYNTAX_PROTO2": 0,
	"SYNTAX_PROTO3": 1,
}

func (x Syntax) String() string {
	return proto.EnumName(Syntax_name, int32(x))
}
func (Syntax) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// Basic field types.
type Field_Kind int32

const (
	// Field type unknown.
	Field_TYPE_UNKNOWN Field_Kind = 0
	// Field type double.
	Field_TYPE_DOUBLE Field_Kind = 1
	// Field type float.
	Field_TYPE_FLOAT Field_Kind = 2
	// Field type int64.
	Field_TYPE_INT64 Field_Kind = 3
	// Field type uint64.
	Field_TYPE_UINT64 Field_Kind = 4
	// Field type int32.
	Field_TYPE_INT32 Field_Kind = 5
	// Field type fixed64.
	Field_TYPE_FIXED64 Field_Kind = 6
	// Field type fixed32.
	Field_TYPE_FIXED32 Field_Kind = 7
	// Field type bool.
	Field_TYPE_BOOL Field_Kind = 8
	// Field type string.
	Field_TYPE_STRING Field_Kind = 9
	// Field type group. Proto2 syntax only, and deprecated.
	Field_TYPE_GROUP Field_Kind = 10
	// Field type message.
	Field_TYPE_MESSAGE Field_Kind = 11
	// Field type bytes.
	Field_TYPE_BYTES Field_Kind = 12
	// Field type uint32.
	Field_TYPE_UINT32 Field_Kind = 13
	// Field type enum.
	Field_TYPE_ENUM Field_Kind = 14
	// Field type sfixed32.
	Field_TYPE_SFIXED32 Field_Kind = 15
	// Field type sfixed64.
	Field_TYPE_SFIXED64 Field_Kind = 16
	// Field type sint32.
	Field_TYPE_SINT32 Field_Kind = 17
	// Field type sint64.
	Field_TYPE_SINT64 Field_Kind = 18
)

var Field_Kind_name = map[int32]string{
	0:  "TYPE_UNKNOWN",
	1:  "TYPE_DOUBLE",
	2:  "TYPE_FLOAT",
	3:  "TYPE_INT64",
	4:  "TYPE_UINT64",
	5:  "TYPE_INT32",
	6:  "TYPE_FIXED64",
	7:  "TYPE_FIXED32",
	8:  "TYPE_BOOL",
	9:  "TYPE_STRING",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "TYPE_UINT32",
	14: "TYPE_ENUM",
	15: "TYPE_SFIXED32",
	16: "TYPE_SFIXED64",
	17: "TYPE_SINT32",
	18: "TYPE_SINT64",
}
var Field_Kind_value = map[string]int32{
	"TYPE_UNKNOWN":  0,
	"TYPE_DOUBLE":   1,
	"TYPE_FLOAT":    2,
	"TYPE_INT64":    3,
	"TYPE_UINT64":   4,
	"TYPE_INT32":    5,
	"TYPE_FIXED64":  6,
	"TYPE_FIXED32":  7,
	"TYPE_BOOL":     8,
	"TYPE_STRING":   9,
	"TYPE_GROUP":    10,
	"TYPE_MESSAGE":  11,
	"TYPE_BYTES":    12,
	"TYPE_UINT32":   13,
	"TYPE_ENUM":     14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32":   17,
	"TYPE_SINT64":   18,
}

func (x Field_Kind) String() string {
	return proto.EnumName(Field_Kind_name, int32(x))
}
func (Field_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1, 0} }

// Whether a field is optional, required, or repeated.
type Field_Cardinality int32

const (
	// For fields with unknown cardinality.
	Field_CARDINALITY_UNKNOWN Field_Cardinality = 0
	// For optional fields.
	Field_CARDINALITY_OPTIONAL Field_Cardinality = 1
	// For required fields. Proto2 syntax only.
	Field_CARDINALITY_REQUIRED Field_Cardinality = 2
	// For repeated fields.
	Field_CARDINALITY_REPEATED Field_Cardinality = 3
)

var Field_Cardinality_name = map[int32]string{
	0: "CARDINALITY_UNKNOWN",
	1: "CARDINALITY_OPTIONAL",
	2: "CARDINALITY_REQUIRED",
	3: "CARDINALITY_REPEATED",
}
var Field_Cardinality_value = map[string]int32{
	"CARDINALITY_UNKNOWN":  0,
	"CARDINALITY_OPTIONAL": 1,
	"CARDINALITY_REQUIRED": 2,
	"CARDINALITY_REPEATED": 3,
}

func (x Field_Cardinality) String() string {
	return proto.EnumName(Field_Cardinality_name, int32(x))
}
func (Field_Cardinality) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1, 1} }

// A protocol buffer message type.
type Type struct {
	// The fully qualified message name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The list of fields.
	Fields []*Field `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	// The list of types appearing in `oneof` definitions in this type.
	Oneofs []string `protobuf:"bytes,3,rep,name=oneofs" json:"oneofs,omitempty"`
	// The protocol buffer options.
	Options []*Option `protobuf:"bytes,4,rep,name=options" json:"options,omitempty"`
	// The source context.
	SourceContext *SourceContext `protobuf:"bytes,5,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	// The source syntax.
	Syntax Syntax `protobuf:"varint,6,opt,name=syntax,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Type) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Type) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

// A single field of a message type.
type Field struct {
	// The field type.
	Kind Field_Kind `protobuf:"varint,1,opt,name=kind,enum=google.protobuf.Field_Kind" json:"kind,omitempty"`
	// The field cardinality.
	Cardinality Field_Cardinality `protobuf:"varint,2,opt,name=cardinality,enum=google.protobuf.Field_Cardinality" json:"cardinality,omitempty"`
	// The field number.
	Number int32 `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
	// The field name.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// The field type URL, without the scheme, for message or enumeration
	// types. Example: `"type.googleapis.com/google.protobuf.Timestamp"`.
	TypeUrl string `protobuf:"bytes,6,opt,name=type_url,json=typeUrl" json:"type_url,omitempty"`
	// The index of the field type in `Type.oneofs`, for message or enumeration
	// types. The first type has index 1; zero means the type is not in the list.
	OneofIndex int32 `protobuf:"varint,7,opt,name=oneof_index,json=oneofIndex" json:"oneof_index,omitempty"`
	// Whether to use alternative packed wire representation.
	Packed bool `protobuf:"varint,8,opt,name=packed" json:"packed,omitempty"`
	// The protocol buffer options.
	Options []*Option `protobuf:"bytes,9,rep,name=options" json:"options,omitempty"`
	// The field JSON name.
	JsonName string `protobuf:"bytes,10,opt,name=json_name,json=jsonName" json:"json_name,omitempty"`
	// The string value of the default value of this field. Proto2 syntax only.
	DefaultValue string `protobuf:"bytes,11,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Field) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

// Enum type definition.
type Enum struct {
	// Enum type name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Enum value definitions.
	Enumvalue []*EnumValue `protobuf:"bytes,2,rep,name=enumvalue" json:"enumvalue,omitempty"`
	// Protocol buffer options.
	Options []*Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	// The source context.
	SourceContext *SourceContext `protobuf:"bytes,4,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	// The source syntax.
	Syntax Syntax `protobuf:"varint,5,opt,name=syntax,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
}

func (m *Enum) Reset()                    { *m = Enum{} }
func (m *Enum) String() string            { return proto.CompactTextString(m) }
func (*Enum) ProtoMessage()               {}
func (*Enum) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Enum) GetEnumvalue() []*EnumValue {
	if m != nil {
		return m.Enumvalue
	}
	return nil
}

func (m *Enum) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Enum) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

// Enum value definition.
type EnumValue struct {
	// Enum value name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Enum value number.
	Number int32 `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	// Protocol buffer options.
	Options []*Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
}

func (m *EnumValue) Reset()                    { *m = EnumValue{} }
func (m *EnumValue) String() string            { return proto.CompactTextString(m) }
func (*EnumValue) ProtoMessage()               {}
func (*EnumValue) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *EnumValue) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

// A protocol buffer option, which can be attached to a message, field,
// enumeration, etc.
type Option struct {
	// The option's name. For example, `"java_package"`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The option's value. For example, `"com.google.protobuf"`.
	Value *google_protobuf.Any `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Option) Reset()                    { *m = Option{} }
func (m *Option) String() string            { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()               {}
func (*Option) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *Option) GetValue() *google_protobuf.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "google.protobuf.Type")
	proto.RegisterType((*Field)(nil), "google.protobuf.Field")
	proto.RegisterType((*Enum)(nil), "google.protobuf.Enum")
	proto.RegisterType((*EnumValue)(nil), "google.protobuf.EnumValue")
	proto.RegisterType((*Option)(nil), "google.protobuf.Option")
	proto.RegisterEnum("google.protobuf.Syntax", Syntax_name, Syntax_value)
	proto.RegisterEnum("google.protobuf.Field_Kind", Field_Kind_name, Field_Kind_value)
	proto.RegisterEnum("google.protobuf.Field_Cardinality", Field_Cardinality_name, Field_Cardinality_value)
}

func init() { proto.RegisterFile("google/protobuf/type.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x8e, 0xda, 0x56,
	0x14, 0x8e, 0x7f, 0xf0, 0xe0, 0xe3, 0x81, 0xb9, 0xb9, 0x89, 0x12, 0x67, 0x22, 0xa5, 0x88, 0x76,
	0x81, 0xb2, 0x60, 0x54, 0x18, 0x8d, 0xba, 0x35, 0x83, 0x87, 0x5a, 0x43, 0x6c, 0xf7, 0x62, 0x9a,
	0xb0, 0x42, 0x1e, 0x30, 0x11, 0x89, 0xb9, 0x46, 0xd8, 0x6e, 0x87, 0x87, 0xe8, 0x3b, 0x54, 0x5d,
	0x76, 0xdd, 0x87, 0xe8, 0x5b, 0xb5, 0xba, 0xd7, 0x60, 0xcc, 0x4f, 0xa5, 0xb4, 0xdd, 0x71, 0xbe,
	0xf3, 0x9d, 0xef, 0xfc, 0xdc, 0xe3, 0x03, 0x5c, 0x7e, 0x8c, 0xa2, 0x8f, 0x61, 0x70, 0xb5, 0x5c,
	0x45, 0x49, 0xf4, 0x90, 0xce, 0xae, 0x92, 0xf5, 0x32, 0x68, 0x72, 0x0b, 0x5f, 0x64, 0xbe, 0xe6,
	0xd6, 0x77, 0xf9, 0xea, 0x90, 0xec, 0xd3, 0x75, 0xe6, 0xbd, 0xfc, 0xe6, 0xd0, 0x15, 0x47, 0xe9,
	0x6a, 0x12, 0x8c, 0x27, 0x11, 0x4d, 0x82, 0xc7, 0x24, 0x63, 0xd5, 0x7f, 0x11, 0x41, 0xf6, 0xd6,
	0xcb, 0x00, 0x63, 0x90, 0xa9, 0xbf, 0x08, 0x74, 0xa1, 0x26, 0x34, 0x54, 0xc2, 0x7f, 0xe3, 0x26,
	0x28, 0xb3, 0x79, 0x10, 0x4e, 0x63, 0x5d, 0xac, 0x49, 0x0d, 0xad, 0xf5, 0xa2, 0x79, 0x90, 0xbf,
	0x79, 0xc7, 0xdc, 0x64, 0xc3, 0xc2, 0x2f, 0x40, 0x89, 0x68, 0x10, 0xcd, 0x62, 0x5d, 0xaa, 0x49,
	0x0d, 0x95, 0x6c, 0x2c, 0xfc, 0x2d, 0x9c, 0x45, 0xcb, 0x64, 0x1e, 0xd1, 0x58, 0x97, 0xb9, 0xd0,
	0xcb, 0x23, 0x21, 0x87, 0xfb, 0xc9, 0x96, 0x87, 0x4d, 0xa8, 0xee, 0xd7, 0xab, 0x97, 0x6a, 0x42,
	0x43, 0x6b, 0xbd, 0x39, 0x8a, 0x1c, 0x70, 0xda, 0x6d, 0xc6, 0x22, 0x95, 0xb8, 0x68, 0xe2, 0x2b,
	0x50, 0xe2, 0x35, 0x4d, 0xfc, 0x47, 0x5d, 0xa9, 0x09, 0x8d, 0xea, 0x89, 0xc4, 0x03, 0xee, 0x26,
	0x1b, 0x5a, 0xfd, 0x0f, 0x05, 0x4a, 0xbc, 0x29, 0x7c, 0x05, 0xf2, 0xe7, 0x39, 0x9d, 0xf2, 0x81,
	0x54, 0x5b, 0xaf, 0x4f, 0xb7, 0xde, 0xbc, 0x9f, 0xd3, 0x29, 0xe1, 0x44, 0xdc, 0x05, 0x6d, 0xe2,
	0xaf, 0xa6, 0x73, 0xea, 0x87, 0xf3, 0x64, 0xad, 0x8b, 0x3c, 0xae, 0xfe, 0x0f, 0x71, 0xb7, 0x3b,
	0x26, 0x29, 0x86, 0xb1, 0x19, 0xd2, 0x74, 0xf1, 0x10, 0xac, 0x74, 0xa9, 0x26, 0x34, 0x4a, 0x64,
	0x63, 0xe5, 0xef, 0x23, 0x17, 0xde, 0xe7, 0x15, 0x94, 0xd9, 0x72, 0x8c, 0xd3, 0x55, 0xc8, 0xfb,
	0x53, 0xc9, 0x19, 0xb3, 0x87, 0xab, 0x10, 0x7f, 0x05, 0x1a, 0x1f, 0xfe, 0x78, 0x4e, 0xa7, 0xc1,
	0xa3, 0x7e, 0xc6, 0xb5, 0x80, 0x43, 0x16, 0x43, 0x58, 0x9e, 0xa5, 0x3f, 0xf9, 0x1c, 0x4c, 0xf5,
	0x72, 0x4d, 0x68, 0x94, 0xc9, 0xc6, 0x2a, 0xbe, 0x95, 0xfa, 0x85, 0x6f, 0xf5, 0x1a, 0xd4, 0x4f,
	0x71, 0x44, 0xc7, 0xbc, 0x3e, 0xe0, 0x75, 0x94, 0x19, 0x60, 0xb3, 0x1a, 0xbf, 0x86, 0xca, 0x34,
	0x98, 0xf9, 0x69, 0x98, 0x8c, 0x7f, 0xf2, 0xc3, 0x34, 0xd0, 0x35, 0x4e, 0x38, 0xdf, 0x80, 0x3f,
	0x32, 0xac, 0xfe, 0xa7, 0x08, 0x32, 0x9b, 0x24, 0x46, 0x70, 0xee, 0x8d, 0x5c, 0x73, 0x3c, 0xb4,
	0xef, 0x6d, 0xe7, 0xbd, 0x8d, 0x9e, 0xe0, 0x0b, 0xd0, 0x38, 0xd2, 0x75, 0x86, 0x9d, 0xbe, 0x89,
	0x04, 0x5c, 0x05, 0xe0, 0xc0, 0x5d, 0xdf, 0x31, 0x3c, 0x24, 0xe6, 0xb6, 0x65, 0x7b, 0x37, 0xd7,
	0x48, 0xca, 0x03, 0x86, 0x19, 0x20, 0x17, 0x09, 0xed, 0x16, 0x2a, 0xe5, 0x39, 0xee, 0xac, 0x0f,
	0x66, 0xf7, 0xe6, 0x1a, 0x29, 0xfb, 0x48, 0xbb, 0x85, 0xce, 0x70, 0x05, 0x54, 0x8e, 0x74, 0x1c,
	0xa7, 0x8f, 0xca, 0xb9, 0xe6, 0xc0, 0x23, 0x96, 0xdd, 0x43, 0x6a, 0xae, 0xd9, 0x23, 0xce, 0xd0,
	0x45, 0x90, 0x2b, 0xbc, 0x33, 0x07, 0x03, 0xa3, 0x67, 0x22, 0x2d, 0x67, 0x74, 0x46, 0x9e, 0x39,
	0x40, 0xe7, 0x7b, 0x65, 0xb5, 0x5b, 0xa8, 0x92, 0xa7, 0x30, 0xed, 0xe1, 0x3b, 0x54, 0xc5, 0x4f,
	0xa1, 0x92, 0xa5, 0xd8, 0x16, 0x71, 0x71, 0x00, 0xdd, 0x5c, 0x23, 0xb4, 0x2b, 0x24, 0x53, 0x79,
	0xba, 0x07, 0xdc, 0x5c, 0x23, 0x5c, 0x4f, 0x40, 0x2b, 0xec, 0x16, 0x7e, 0x09, 0xcf, 0x6e, 0x0d,
	0xd2, 0xb5, 0x6c, 0xa3, 0x6f, 0x79, 0xa3, 0xc2, 0x5c, 0x75, 0x78, 0x5e, 0x74, 0x38, 0xae, 0x67,
	0x39, 0xb6, 0xd1, 0x47, 0xc2, 0xa1, 0x87, 0x98, 0x3f, 0x0c, 0x2d, 0x62, 0x76, 0x91, 0x78, 0xec,
	0x71, 0x4d, 0xc3, 0x33, 0xbb, 0x48, 0xaa, 0xff, 0x25, 0x80, 0x6c, 0xd2, 0x74, 0x71, 0xf2, 0x8c,
	0x7c, 0x07, 0x6a, 0x40, 0xd3, 0x45, 0xf6, 0xfc, 0xd9, 0x25, 0xb9, 0x3c, 0x5a, 0x2a, 0x16, 0xcd,
	0x97, 0x81, 0xec, 0xc8, 0xc5, 0x65, 0x94, 0xfe, 0xf3, 0xe1, 0x90, 0xff, 0xdf, 0xe1, 0x28, 0x7d,
	0xd9, 0xe1, 0xf8, 0x04, 0x6a, 0xde, 0xc2, 0xc9, 0x29, 0xec, 0x3e, 0x6c, 0x71, 0xef, 0xc3, 0xfe,
	0xf7, 0x3d, 0xd6, 0xbf, 0x07, 0x25, 0x83, 0x4e, 0x26, 0x7a, 0x0b, 0xa5, 0xed, 0xa8, 0x59, 0xe3,
	0xcf, 0x8f, 0xe4, 0x0c, 0xba, 0x26, 0x19, 0xe5, 0x6d, 0x13, 0x94, 0xac, 0x0f, 0xb6, 0x6c, 0x83,
	0x91, 0xed, 0x19, 0x1f, 0xc6, 0x2e, 0x71, 0x3c, 0xa7, 0x85, 0x9e, 0x1c, 0x42, 0x6d, 0x24, 0x74,
	0xfa, 0xf0, 0x6c, 0x12, 0x2d, 0x0e, 0x15, 0x3b, 0x2a, 0xfb, 0x0b, 0x71, 0x99, 0xe5, 0x0a, 0xbf,
	0x0a, 0xc2, 0x6f, 0xa2, 0xd4, 0x73, 0x3b, 0xbf, 0x8b, 0x6f, 0x7a, 0x19, 0xcf, 0xdd, 0x66, 0x7e,
	0x1f, 0x84, 0xe1, 0x3d, 0x8d, 0x7e, 0xa6, 0x8c, 0x1f, 0x3f, 0x28, 0x5c, 0xa0, 0xfd, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x95, 0xbb, 0xeb, 0x52, 0xf3, 0x06, 0x00, 0x00,
}
