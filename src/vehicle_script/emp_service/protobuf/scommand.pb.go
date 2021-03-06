// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scommand.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Command_TaskType int32

const (
	Command_DEFAULT_TASKTYPE Command_TaskType = 0
	Command_GW_SET           Command_TaskType = 1
	Command_STRATEGY_ADD     Command_TaskType = 3
	Command_STRATEGY_SET     Command_TaskType = 4
)

var Command_TaskType_name = map[int32]string{
	0: "DEFAULT_TASKTYPE",
	1: "GW_SET",
	3: "STRATEGY_ADD",
	4: "STRATEGY_SET",
}

var Command_TaskType_value = map[string]int32{
	"DEFAULT_TASKTYPE": 0,
	"GW_SET":           1,
	"STRATEGY_ADD":     3,
	"STRATEGY_SET":     4,
}

func (x Command_TaskType) String() string {
	return proto.EnumName(Command_TaskType_name, int32(x))
}

func (Command_TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1bdf113f3d6e5a7a, []int{0, 0}
}

type GwSetParam_Type int32

const (
	GwSetParam_DEFAULT GwSetParam_Type = 0
	GwSetParam_PROTECT GwSetParam_Type = 1
	GwSetParam_RESTART GwSetParam_Type = 2
)

var GwSetParam_Type_name = map[int32]string{
	0: "DEFAULT",
	1: "PROTECT",
	2: "RESTART",
}

var GwSetParam_Type_value = map[string]int32{
	"DEFAULT": 0,
	"PROTECT": 1,
	"RESTART": 2,
}

func (x GwSetParam_Type) String() string {
	return proto.EnumName(GwSetParam_Type_name, int32(x))
}

func (GwSetParam_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1bdf113f3d6e5a7a, []int{1, 0}
}

type StrategyAddParam_Type int32

const (
	StrategyAddParam_TYPEDEFAULT StrategyAddParam_Type = 0
	StrategyAddParam_WHITEMODE   StrategyAddParam_Type = 1
	StrategyAddParam_BLACKMODE   StrategyAddParam_Type = 2
)

var StrategyAddParam_Type_name = map[int32]string{
	0: "TYPEDEFAULT",
	1: "WHITEMODE",
	2: "BLACKMODE",
}

var StrategyAddParam_Type_value = map[string]int32{
	"TYPEDEFAULT": 0,
	"WHITEMODE":   1,
	"BLACKMODE":   2,
}

func (x StrategyAddParam_Type) String() string {
	return proto.EnumName(StrategyAddParam_Type_name, int32(x))
}

func (StrategyAddParam_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1bdf113f3d6e5a7a, []int{2, 0}
}

type StrategyAddParam_HandleMode int32

const (
	StrategyAddParam_MODEDEFAULT    StrategyAddParam_HandleMode = 0
	StrategyAddParam_PREVENTWARNING StrategyAddParam_HandleMode = 1
	StrategyAddParam_WARNING        StrategyAddParam_HandleMode = 2
)

var StrategyAddParam_HandleMode_name = map[int32]string{
	0: "MODEDEFAULT",
	1: "PREVENTWARNING",
	2: "WARNING",
}

var StrategyAddParam_HandleMode_value = map[string]int32{
	"MODEDEFAULT":    0,
	"PREVENTWARNING": 1,
	"WARNING":        2,
}

func (x StrategyAddParam_HandleMode) String() string {
	return proto.EnumName(StrategyAddParam_HandleMode_name, int32(x))
}

func (StrategyAddParam_HandleMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1bdf113f3d6e5a7a, []int{2, 1}
}

type StrategySetParam_HandleMode int32

const (
	StrategySetParam_MODEDEFAULT    StrategySetParam_HandleMode = 0
	StrategySetParam_PREVENTWARNING StrategySetParam_HandleMode = 1
	StrategySetParam_WARNING        StrategySetParam_HandleMode = 2
)

var StrategySetParam_HandleMode_name = map[int32]string{
	0: "MODEDEFAULT",
	1: "PREVENTWARNING",
	2: "WARNING",
}

var StrategySetParam_HandleMode_value = map[string]int32{
	"MODEDEFAULT":    0,
	"PREVENTWARNING": 1,
	"WARNING":        2,
}

func (x StrategySetParam_HandleMode) String() string {
	return proto.EnumName(StrategySetParam_HandleMode_name, int32(x))
}

func (StrategySetParam_HandleMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1bdf113f3d6e5a7a, []int{3, 0}
}

//基础最外层命令
type Command struct {
	ItemType             Command_TaskType `protobuf:"varint,1,opt,name=item_type,json=itemType,proto3,enum=protobuf.Command_TaskType" json:"item_type,omitempty"`
	Param                []byte           `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
	CmdID                string           `protobuf:"bytes,3,opt,name=cmdID,proto3" json:"cmdID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bdf113f3d6e5a7a, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetItemType() Command_TaskType {
	if m != nil {
		return m.ItemType
	}
	return Command_DEFAULT_TASKTYPE
}

func (m *Command) GetParam() []byte {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *Command) GetCmdID() string {
	if m != nil {
		return m.CmdID
	}
	return ""
}

//小v设置
type GwSetParam struct {
	Type                 GwSetParam_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protobuf.GwSetParam_Type" json:"type,omitempty"`
	Switch               bool            `protobuf:"varint,2,opt,name=switch,proto3" json:"switch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GwSetParam) Reset()         { *m = GwSetParam{} }
func (m *GwSetParam) String() string { return proto.CompactTextString(m) }
func (*GwSetParam) ProtoMessage()    {}
func (*GwSetParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bdf113f3d6e5a7a, []int{1}
}

func (m *GwSetParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GwSetParam.Unmarshal(m, b)
}
func (m *GwSetParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GwSetParam.Marshal(b, m, deterministic)
}
func (m *GwSetParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GwSetParam.Merge(m, src)
}
func (m *GwSetParam) XXX_Size() int {
	return xxx_messageInfo_GwSetParam.Size(m)
}
func (m *GwSetParam) XXX_DiscardUnknown() {
	xxx_messageInfo_GwSetParam.DiscardUnknown(m)
}

var xxx_messageInfo_GwSetParam proto.InternalMessageInfo

func (m *GwSetParam) GetType() GwSetParam_Type {
	if m != nil {
		return m.Type
	}
	return GwSetParam_DEFAULT
}

func (m *GwSetParam) GetSwitch() bool {
	if m != nil {
		return m.Switch
	}
	return false
}

//策略添加
type StrategyAddParam struct {
	StrategyId           string                      `protobuf:"bytes,1,opt,name=strategy_id,json=strategyId,proto3" json:"strategy_id,omitempty"`
	HandleMode           StrategyAddParam_HandleMode `protobuf:"varint,2,opt,name=handle_mode,json=handleMode,proto3,enum=protobuf.StrategyAddParam_HandleMode" json:"handle_mode,omitempty"`
	DefenseType          StrategyAddParam_Type       `protobuf:"varint,3,opt,name=defense_type,json=defenseType,proto3,enum=protobuf.StrategyAddParam_Type" json:"defense_type,omitempty"`
	DIPList              []string                    `protobuf:"bytes,5,rep,name=dIP_list,json=dIPList,proto3" json:"dIP_list,omitempty"`
	URLList              []string                    `protobuf:"bytes,6,rep,name=URL_list,json=URLList,proto3" json:"URL_list,omitempty"`
	Enable               bool                        `protobuf:"varint,7,opt,name=enable,proto3" json:"enable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *StrategyAddParam) Reset()         { *m = StrategyAddParam{} }
func (m *StrategyAddParam) String() string { return proto.CompactTextString(m) }
func (*StrategyAddParam) ProtoMessage()    {}
func (*StrategyAddParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bdf113f3d6e5a7a, []int{2}
}

func (m *StrategyAddParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StrategyAddParam.Unmarshal(m, b)
}
func (m *StrategyAddParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StrategyAddParam.Marshal(b, m, deterministic)
}
func (m *StrategyAddParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategyAddParam.Merge(m, src)
}
func (m *StrategyAddParam) XXX_Size() int {
	return xxx_messageInfo_StrategyAddParam.Size(m)
}
func (m *StrategyAddParam) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategyAddParam.DiscardUnknown(m)
}

var xxx_messageInfo_StrategyAddParam proto.InternalMessageInfo

func (m *StrategyAddParam) GetStrategyId() string {
	if m != nil {
		return m.StrategyId
	}
	return ""
}

func (m *StrategyAddParam) GetHandleMode() StrategyAddParam_HandleMode {
	if m != nil {
		return m.HandleMode
	}
	return StrategyAddParam_MODEDEFAULT
}

func (m *StrategyAddParam) GetDefenseType() StrategyAddParam_Type {
	if m != nil {
		return m.DefenseType
	}
	return StrategyAddParam_TYPEDEFAULT
}

func (m *StrategyAddParam) GetDIPList() []string {
	if m != nil {
		return m.DIPList
	}
	return nil
}

func (m *StrategyAddParam) GetURLList() []string {
	if m != nil {
		return m.URLList
	}
	return nil
}

func (m *StrategyAddParam) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

// 策略修改
type StrategySetParam struct {
	StrategyId           string                      `protobuf:"bytes,1,opt,name=strategy_id,json=strategyId,proto3" json:"strategy_id,omitempty"`
	HandleMode           StrategySetParam_HandleMode `protobuf:"varint,2,opt,name=handle_mode,json=handleMode,proto3,enum=protobuf.StrategySetParam_HandleMode" json:"handle_mode,omitempty"`
	Enable               bool                        `protobuf:"varint,7,opt,name=enable,proto3" json:"enable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *StrategySetParam) Reset()         { *m = StrategySetParam{} }
func (m *StrategySetParam) String() string { return proto.CompactTextString(m) }
func (*StrategySetParam) ProtoMessage()    {}
func (*StrategySetParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bdf113f3d6e5a7a, []int{3}
}

func (m *StrategySetParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StrategySetParam.Unmarshal(m, b)
}
func (m *StrategySetParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StrategySetParam.Marshal(b, m, deterministic)
}
func (m *StrategySetParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategySetParam.Merge(m, src)
}
func (m *StrategySetParam) XXX_Size() int {
	return xxx_messageInfo_StrategySetParam.Size(m)
}
func (m *StrategySetParam) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategySetParam.DiscardUnknown(m)
}

var xxx_messageInfo_StrategySetParam proto.InternalMessageInfo

func (m *StrategySetParam) GetStrategyId() string {
	if m != nil {
		return m.StrategyId
	}
	return ""
}

func (m *StrategySetParam) GetHandleMode() StrategySetParam_HandleMode {
	if m != nil {
		return m.HandleMode
	}
	return StrategySetParam_MODEDEFAULT
}

func (m *StrategySetParam) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func init() {
	proto.RegisterEnum("protobuf.Command_TaskType", Command_TaskType_name, Command_TaskType_value)
	proto.RegisterEnum("protobuf.GwSetParam_Type", GwSetParam_Type_name, GwSetParam_Type_value)
	proto.RegisterEnum("protobuf.StrategyAddParam_Type", StrategyAddParam_Type_name, StrategyAddParam_Type_value)
	proto.RegisterEnum("protobuf.StrategyAddParam_HandleMode", StrategyAddParam_HandleMode_name, StrategyAddParam_HandleMode_value)
	proto.RegisterEnum("protobuf.StrategySetParam_HandleMode", StrategySetParam_HandleMode_name, StrategySetParam_HandleMode_value)
	proto.RegisterType((*Command)(nil), "protobuf.Command")
	proto.RegisterType((*GwSetParam)(nil), "protobuf.GwSetParam")
	proto.RegisterType((*StrategyAddParam)(nil), "protobuf.StrategyAddParam")
	proto.RegisterType((*StrategySetParam)(nil), "protobuf.StrategySetParam")
}

func init() { proto.RegisterFile("scommand.proto", fileDescriptor_1bdf113f3d6e5a7a) }

var fileDescriptor_1bdf113f3d6e5a7a = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x8d, 0x81, 0x00, 0x1e, 0x53, 0x6a, 0xad, 0xa2, 0xc8, 0xe9, 0x25, 0xc8, 0x52, 0x25, 0x2e,
	0xf1, 0x21, 0x55, 0xd5, 0x5b, 0x25, 0x07, 0x6f, 0x88, 0x15, 0x87, 0x58, 0xeb, 0xa5, 0x28, 0x27,
	0xcb, 0xb0, 0x9b, 0x62, 0x15, 0x63, 0x84, 0x5d, 0x45, 0x5c, 0xab, 0x7e, 0x58, 0xff, 0xa1, 0x3f,
	0x54, 0xed, 0xae, 0xc1, 0x6d, 0xa5, 0xf4, 0xd2, 0x9c, 0xbc, 0x6f, 0xe6, 0xed, 0xf8, 0xcd, 0x7b,
	0x0b, 0xfd, 0x62, 0x91, 0x67, 0x59, 0xb2, 0x66, 0xce, 0x66, 0x9b, 0x97, 0x39, 0xea, 0xca, 0xcf,
	0xfc, 0xeb, 0xa3, 0xfd, 0x43, 0x83, 0xce, 0x48, 0xf5, 0xd0, 0x07, 0xd0, 0xd3, 0x92, 0x67, 0x71,
	0xb9, 0xdb, 0x70, 0x4b, 0x1b, 0x68, 0xc3, 0xfe, 0xe5, 0x1b, 0x67, 0xcf, 0x74, 0x2a, 0x96, 0x43,
	0x93, 0xe2, 0x0b, 0xdd, 0x6d, 0x38, 0xe9, 0x0a, 0xb2, 0x38, 0xa1, 0x13, 0x38, 0xde, 0x24, 0xdb,
	0x24, 0xb3, 0x1a, 0x03, 0x6d, 0xd8, 0x23, 0x0a, 0x88, 0xea, 0x22, 0x63, 0xbe, 0x67, 0x35, 0x07,
	0xda, 0x50, 0x27, 0x0a, 0xd8, 0x21, 0x74, 0xf7, 0x13, 0xd0, 0x09, 0x98, 0x1e, 0xbe, 0x76, 0xa7,
	0x01, 0x8d, 0xa9, 0x1b, 0xdd, 0xd2, 0x87, 0x10, 0x9b, 0x47, 0x08, 0xa0, 0x3d, 0x9e, 0xc5, 0x11,
	0xa6, 0xa6, 0x86, 0x4c, 0xe8, 0x45, 0x94, 0xb8, 0x14, 0x8f, 0x1f, 0x62, 0xd7, 0xf3, 0xcc, 0xe6,
	0x1f, 0x15, 0xc1, 0x69, 0xd9, 0xdf, 0x34, 0x80, 0xf1, 0x53, 0xc4, 0xcb, 0x50, 0xfe, 0xf6, 0x02,
	0x5a, 0xbf, 0x2d, 0x70, 0x56, 0x2f, 0x50, 0x73, 0x1c, 0xa9, 0x5f, 0xd2, 0xd0, 0x29, 0xb4, 0x8b,
	0xa7, 0xb4, 0x5c, 0x2c, 0xa5, 0xf8, 0x2e, 0xa9, 0x90, 0x7d, 0x01, 0x2d, 0xa9, 0xd1, 0x80, 0x4e,
	0xa5, 0xd1, 0x3c, 0x12, 0x20, 0x24, 0xf7, 0x14, 0x8f, 0x84, 0x36, 0x03, 0x3a, 0x04, 0x47, 0xd4,
	0x25, 0xd4, 0x6c, 0xd8, 0xdf, 0x9b, 0x60, 0x46, 0xe5, 0x36, 0x29, 0xf9, 0xe7, 0x9d, 0xcb, 0x98,
	0x92, 0x72, 0x0e, 0x46, 0x51, 0xd5, 0xe2, 0x94, 0x49, 0x45, 0x3a, 0x81, 0x7d, 0xc9, 0x67, 0xe8,
	0x1a, 0x8c, 0x65, 0xb2, 0x66, 0x2b, 0x1e, 0x67, 0x39, 0xe3, 0x52, 0x41, 0xff, 0xf2, 0x6d, 0x2d,
	0xf9, 0xef, 0x89, 0xce, 0x8d, 0x64, 0xdf, 0xe5, 0x8c, 0x13, 0x58, 0x1e, 0xce, 0xe8, 0x0a, 0x7a,
	0x8c, 0x3f, 0xf2, 0x75, 0xc1, 0x55, 0x78, 0x4d, 0x39, 0xe8, 0xfc, 0x1f, 0x83, 0xa4, 0x03, 0x46,
	0x75, 0x49, 0x2e, 0x7a, 0x06, 0x5d, 0xe6, 0x87, 0xf1, 0x2a, 0x2d, 0x4a, 0xeb, 0x78, 0xd0, 0x1c,
	0xea, 0xa4, 0xc3, 0xfc, 0x30, 0x48, 0x8b, 0x52, 0xb4, 0xa6, 0x24, 0x50, 0xad, 0xb6, 0x6a, 0x4d,
	0x49, 0x20, 0x5b, 0xa7, 0xd0, 0xe6, 0xeb, 0x64, 0xbe, 0xe2, 0x56, 0x47, 0xd9, 0xa7, 0x90, 0xfd,
	0xbe, 0xb2, 0xef, 0x35, 0x18, 0x22, 0xd6, 0xda, 0xc2, 0x57, 0xa0, 0xcf, 0x6e, 0x7c, 0x8a, 0xef,
	0xee, 0x3d, 0x6c, 0x6a, 0x02, 0x5e, 0x05, 0xee, 0xe8, 0x56, 0xc2, 0x86, 0xfd, 0x11, 0xa0, 0x5e,
	0x51, 0x5c, 0x16, 0xf5, 0xfa, 0x32, 0x82, 0x7e, 0x48, 0xf0, 0x27, 0x3c, 0xa1, 0x33, 0x97, 0x4c,
	0xfc, 0xc9, 0x58, 0xc5, 0xb0, 0x07, 0x0d, 0xfb, 0xa7, 0x56, 0xc7, 0x70, 0x78, 0x11, 0x2f, 0x16,
	0xc3, 0xe1, 0xfd, 0x3c, 0x13, 0xc3, 0x73, 0x66, 0xfc, 0xe7, 0x56, 0xf3, 0xb6, 0x54, 0xf2, 0xee,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0xd4, 0x05, 0xc6, 0xc7, 0x03, 0x00, 0x00,
}
