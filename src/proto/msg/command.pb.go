// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: command.proto

package msg

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

type Command int32

const (
	Command_KEEP_ALIVE            Command = 0
	Command_RESP                  Command = 1
	Command_GATE                  Command = 2
	Command_SERVER_STATE          Command = 3
	Command_SERVER_LIST_CHANGED   Command = 4
	Command_SERVER_RESET          Command = 5
	Command_CONNECT_GATE_REQ      Command = 6
	Command_SESSION_CERT_NTF      Command = 7
	Command_CCHAT_CHAT            Command = 8
	Command_DISPATCH              Command = 9
	Command_LD_CREATE_PLAYER_REQ  Command = 10
	Command_GD_LOAD_PLAYER_REQ    Command = 11
	Command_GD_RELEASE_PLAYER_NTF Command = 12
	Command_GD_CREATE_ROLE_REQ    Command = 13
	Command_GD_LOAD_ROLE_REQ      Command = 14
	Command_RELOAD_ETC_REQ        Command = 15
	Command_UPDATE_ETC_NTF        Command = 16
	Command_MOVE_NTF              Command = 17
	Command_NEAR_ENTITIES_NTF     Command = 18
	Command_ENTITY_DATA_REQ       Command = 19
	Command_ROLE_ENTER            Command = 20
)

var Command_name = map[int32]string{
	0:  "KEEP_ALIVE",
	1:  "RESP",
	2:  "GATE",
	3:  "SERVER_STATE",
	4:  "SERVER_LIST_CHANGED",
	5:  "SERVER_RESET",
	6:  "CONNECT_GATE_REQ",
	7:  "SESSION_CERT_NTF",
	8:  "CCHAT_CHAT",
	9:  "DISPATCH",
	10: "LD_CREATE_PLAYER_REQ",
	11: "GD_LOAD_PLAYER_REQ",
	12: "GD_RELEASE_PLAYER_NTF",
	13: "GD_CREATE_ROLE_REQ",
	14: "GD_LOAD_ROLE_REQ",
	15: "RELOAD_ETC_REQ",
	16: "UPDATE_ETC_NTF",
	17: "MOVE_NTF",
	18: "NEAR_ENTITIES_NTF",
	19: "ENTITY_DATA_REQ",
	20: "ROLE_ENTER",
}

var Command_value = map[string]int32{
	"KEEP_ALIVE":            0,
	"RESP":                  1,
	"GATE":                  2,
	"SERVER_STATE":          3,
	"SERVER_LIST_CHANGED":   4,
	"SERVER_RESET":          5,
	"CONNECT_GATE_REQ":      6,
	"SESSION_CERT_NTF":      7,
	"CCHAT_CHAT":            8,
	"DISPATCH":              9,
	"LD_CREATE_PLAYER_REQ":  10,
	"GD_LOAD_PLAYER_REQ":    11,
	"GD_RELEASE_PLAYER_NTF": 12,
	"GD_CREATE_ROLE_REQ":    13,
	"GD_LOAD_ROLE_REQ":      14,
	"RELOAD_ETC_REQ":        15,
	"UPDATE_ETC_NTF":        16,
	"MOVE_NTF":              17,
	"NEAR_ENTITIES_NTF":     18,
	"ENTITY_DATA_REQ":       19,
	"ROLE_ENTER":            20,
}

func (x Command) String() string {
	return proto.EnumName(Command_name, int32(x))
}

func (Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_213c0bb044472049, []int{0}
}

func init() {
	proto.RegisterEnum("Command", Command_name, Command_value)
}

func init() { proto.RegisterFile("command.proto", fileDescriptor_213c0bb044472049) }

var fileDescriptor_213c0bb044472049 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcb, 0x52, 0x22, 0x31,
	0x14, 0x86, 0xe9, 0x81, 0x01, 0xe6, 0x0c, 0x97, 0x70, 0xb8, 0xcc, 0xcc, 0xa6, 0xd7, 0x53, 0xb3,
	0x18, 0x16, 0x3e, 0x41, 0x4c, 0x8e, 0x4d, 0x97, 0x6d, 0xba, 0x49, 0x22, 0x55, 0xb8, 0x49, 0x29,
	0x52, 0xae, 0x10, 0xab, 0xb1, 0x7c, 0x16, 0xb7, 0xbe, 0x8d, 0x4b, 0x1f, 0xc1, 0xc2, 0x17, 0xb1,
	0x12, 0x90, 0x62, 0x77, 0xf2, 0x55, 0xce, 0xf7, 0xff, 0xa9, 0x40, 0x7b, 0xb1, 0x5e, 0xad, 0xae,
	0xef, 0x6f, 0xff, 0x3f, 0x94, 0xeb, 0xc7, 0xf5, 0xbf, 0x97, 0x2a, 0x34, 0xc4, 0x8e, 0x60, 0x07,
	0xe0, 0x9c, 0xa8, 0x70, 0x3c, 0x4b, 0x67, 0xc4, 0x2a, 0xd8, 0x84, 0x9a, 0x26, 0x53, 0xb0, 0xc8,
	0x4f, 0x09, 0xb7, 0xc4, 0xbe, 0x21, 0x83, 0x96, 0x21, 0x3d, 0x23, 0xed, 0x8c, 0xf5, 0xa4, 0x8a,
	0xbf, 0xa0, 0xbf, 0x27, 0x59, 0x6a, 0xac, 0x13, 0x13, 0xae, 0x12, 0x92, 0xac, 0x76, 0x74, 0x55,
	0x93, 0x21, 0xcb, 0xbe, 0xe3, 0x00, 0x98, 0xc8, 0x95, 0x22, 0x61, 0x9d, 0xd7, 0x39, 0x4d, 0x53,
	0x56, 0xf7, 0xd4, 0x90, 0x31, 0x69, 0xae, 0x9c, 0x20, 0x6d, 0x9d, 0xb2, 0x67, 0xac, 0xe1, 0xcb,
	0x08, 0x31, 0xe1, 0x41, 0x68, 0x59, 0x13, 0x5b, 0xd0, 0x94, 0xa9, 0x29, 0xb8, 0x15, 0x13, 0xf6,
	0x03, 0x7f, 0xc3, 0x20, 0x93, 0x4e, 0x68, 0xf2, 0x9a, 0x22, 0xe3, 0xf3, 0x90, 0x32, 0x65, 0x80,
	0x23, 0xc0, 0x44, 0xba, 0x2c, 0xe7, 0xf2, 0x98, 0xff, 0xc4, 0x3f, 0x30, 0x4c, 0xa4, 0xd3, 0x94,
	0x11, 0x37, 0x87, 0x15, 0x1f, 0xd5, 0xda, 0xaf, 0xec, 0x65, 0x3a, 0xcf, 0x76, 0xc5, 0xda, 0xbe,
	0xd8, 0x97, 0xea, 0x40, 0x3b, 0x88, 0xd0, 0xd1, 0x14, 0x20, 0x59, 0x11, 0x58, 0xd7, 0xb3, 0xcb,
	0x42, 0xfa, 0x75, 0xcf, 0xbc, 0x95, 0xf9, 0xc2, 0x17, 0xf9, 0x8c, 0xc2, 0xa9, 0x87, 0x43, 0xe8,
	0x29, 0xe2, 0xda, 0x91, 0xb2, 0xa9, 0x4d, 0xc9, 0x04, 0x8c, 0xd8, 0x87, 0x6e, 0x20, 0x73, 0x27,
	0xb9, 0xe5, 0xc1, 0xd6, 0xf7, 0x4f, 0x0f, 0x79, 0xa4, 0x2c, 0x69, 0x36, 0x38, 0xfd, 0xfb, 0xba,
	0x8d, 0xa3, 0xb7, 0x6d, 0x1c, 0xbd, 0x6f, 0xe3, 0xe8, 0xf9, 0x23, 0xae, 0x5c, 0x8d, 0x52, 0x65,
	0x96, 0xe5, 0xd3, 0xb2, 0x1c, 0x6f, 0xca, 0xc5, 0x38, 0xfc, 0xe4, 0x78, 0xb5, 0xb9, 0xbb, 0xa9,
	0x87, 0xf1, 0xe4, 0x33, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x5b, 0x7c, 0x77, 0xe5, 0x01, 0x00, 0x00,
}
