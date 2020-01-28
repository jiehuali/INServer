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
	Command_KEEP_ALIVE                 Command = 0
	Command_RESP                       Command = 1
	Command_GATE                       Command = 2
	Command_SERVER_STATE               Command = 3
	Command_SERVER_LIST_CHANGED        Command = 4
	Command_SERVER_RESET               Command = 5
	Command_CONNECT_GATE_REQ           Command = 6
	Command_SESSION_CERT_NTF           Command = 7
	Command_CCHAT_CHAT                 Command = 8
	Command_DISPATCH                   Command = 9
	Command_LD_CREATE_PLAYER_REQ       Command = 10
	Command_GD_LOAD_PLAYER_REQ         Command = 11
	Command_GD_RELEASE_PLAYER_NTF      Command = 12
	Command_GD_CREATE_ROLE_REQ         Command = 13
	Command_GD_LOAD_ROLE_REQ           Command = 14
	Command_RELOAD_ETC_REQ             Command = 15
	Command_UPDATE_ETC_NTF             Command = 16
	Command_MOVE_NTF                   Command = 17
	Command_NEAR_ENTITIES_NTF          Command = 18
	Command_ENTITY_DATA_REQ            Command = 19
	Command_ROLE_ENTER                 Command = 20
	Command_UPDATE_PLAYER_ADDRESS_NTF  Command = 21
	Command_REMOVE_PLAYER_ADDRESS_NTF  Command = 22
	Command_UPDATE_MAP_ADDRESS_NTF     Command = 23
	Command_REMOVE_MAP_ADDRESS_NTF     Command = 24
	Command_GET_MAP_ADDRESS_REQ        Command = 25
	Command_LOAD_STATIC_MAP_REQ        Command = 26
	Command_LOAD_DYNAMIC_MAP_REQ       Command = 27
	Command_UPDATE_STATIC_MAP_UUID_NTF Command = 28
	Command_GET_STATIC_MAP_UUID_REQ    Command = 29
	Command_SAVE_STATIC_MAP_REQ        Command = 30
	Command_SAVE_DYNAMIC_MAP_REQ       Command = 31
	Command_SAVE_ROLE                  Command = 32
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
	21: "UPDATE_PLAYER_ADDRESS_NTF",
	22: "REMOVE_PLAYER_ADDRESS_NTF",
	23: "UPDATE_MAP_ADDRESS_NTF",
	24: "REMOVE_MAP_ADDRESS_NTF",
	25: "GET_MAP_ADDRESS_REQ",
	26: "LOAD_STATIC_MAP_REQ",
	27: "LOAD_DYNAMIC_MAP_REQ",
	28: "UPDATE_STATIC_MAP_UUID_NTF",
	29: "GET_STATIC_MAP_UUID_REQ",
	30: "SAVE_STATIC_MAP_REQ",
	31: "SAVE_DYNAMIC_MAP_REQ",
	32: "SAVE_ROLE",
}

var Command_value = map[string]int32{
	"KEEP_ALIVE":                 0,
	"RESP":                       1,
	"GATE":                       2,
	"SERVER_STATE":               3,
	"SERVER_LIST_CHANGED":        4,
	"SERVER_RESET":               5,
	"CONNECT_GATE_REQ":           6,
	"SESSION_CERT_NTF":           7,
	"CCHAT_CHAT":                 8,
	"DISPATCH":                   9,
	"LD_CREATE_PLAYER_REQ":       10,
	"GD_LOAD_PLAYER_REQ":         11,
	"GD_RELEASE_PLAYER_NTF":      12,
	"GD_CREATE_ROLE_REQ":         13,
	"GD_LOAD_ROLE_REQ":           14,
	"RELOAD_ETC_REQ":             15,
	"UPDATE_ETC_NTF":             16,
	"MOVE_NTF":                   17,
	"NEAR_ENTITIES_NTF":          18,
	"ENTITY_DATA_REQ":            19,
	"ROLE_ENTER":                 20,
	"UPDATE_PLAYER_ADDRESS_NTF":  21,
	"REMOVE_PLAYER_ADDRESS_NTF":  22,
	"UPDATE_MAP_ADDRESS_NTF":     23,
	"REMOVE_MAP_ADDRESS_NTF":     24,
	"GET_MAP_ADDRESS_REQ":        25,
	"LOAD_STATIC_MAP_REQ":        26,
	"LOAD_DYNAMIC_MAP_REQ":       27,
	"UPDATE_STATIC_MAP_UUID_NTF": 28,
	"GET_STATIC_MAP_UUID_REQ":    29,
	"SAVE_STATIC_MAP_REQ":        30,
	"SAVE_DYNAMIC_MAP_REQ":       31,
	"SAVE_ROLE":                  32,
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
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x1b, 0xe8, 0x4f, 0x7a, 0x49, 0x52, 0xf7, 0xe6, 0xaf, 0x49, 0xe9, 0xc0, 0x12, 0xb1,
	0x20, 0x0b, 0x9e, 0xc0, 0xd8, 0x97, 0x64, 0xc4, 0x64, 0x66, 0x6a, 0x3b, 0x91, 0xc2, 0xc6, 0x82,
	0x52, 0xb1, 0x0a, 0x41, 0x29, 0xe2, 0x59, 0x78, 0x24, 0x96, 0x3c, 0x02, 0x0a, 0x2f, 0x82, 0xee,
	0xcd, 0x50, 0xa5, 0x11, 0x3b, 0xe7, 0x7c, 0xbe, 0xe7, 0x1c, 0xc7, 0x1e, 0x68, 0xde, 0xac, 0x96,
	0xcb, 0x0f, 0x5f, 0x3e, 0xbd, 0xfa, 0xba, 0x5e, 0x7d, 0x5b, 0xbd, 0xdc, 0x1c, 0xc1, 0x89, 0xd9,
	0x2a, 0xd8, 0x02, 0x78, 0x47, 0x54, 0x46, 0x9d, 0xa5, 0x73, 0x52, 0x07, 0x58, 0x87, 0x43, 0x47,
	0xbe, 0x54, 0x35, 0x5e, 0x8d, 0x75, 0x20, 0xf5, 0x08, 0x15, 0x34, 0x3c, 0xb9, 0x39, 0xb9, 0xe8,
	0x03, 0x2b, 0x8f, 0xb1, 0x0f, 0xed, 0x4a, 0xc9, 0x52, 0x1f, 0xa2, 0x99, 0xe8, 0x7c, 0x4c, 0x56,
	0x1d, 0xee, 0x6c, 0x75, 0xe4, 0x29, 0xa8, 0x23, 0xec, 0x80, 0x32, 0x45, 0x9e, 0x93, 0x09, 0x91,
	0xed, 0xa2, 0xa3, 0x6b, 0x75, 0xcc, 0xaa, 0x27, 0xef, 0xd3, 0x22, 0x8f, 0x86, 0x5c, 0x88, 0x79,
	0x78, 0xab, 0x4e, 0xb8, 0x8c, 0x31, 0x13, 0x2d, 0x86, 0x41, 0xd5, 0xb1, 0x01, 0x75, 0x9b, 0xfa,
	0x52, 0x07, 0x33, 0x51, 0xa7, 0x78, 0x01, 0x9d, 0xcc, 0x46, 0xe3, 0x88, 0x6d, 0xca, 0x4c, 0x2f,
	0x24, 0xe5, 0x5a, 0x01, 0xf6, 0x00, 0xc7, 0x36, 0x66, 0x85, 0xb6, 0xbb, 0xfa, 0x13, 0x1c, 0x40,
	0x77, 0x6c, 0xa3, 0xa3, 0x8c, 0xb4, 0xbf, 0x1f, 0xe1, 0xa8, 0x46, 0x35, 0x52, 0x99, 0xb9, 0x22,
	0xdb, 0x16, 0x6b, 0x72, 0xb1, 0x7f, 0x56, 0xf7, 0x6a, 0x0b, 0x11, 0x5a, 0x8e, 0x44, 0xa4, 0x60,
	0x44, 0x3b, 0x63, 0x6d, 0x56, 0x5a, 0x1e, 0x67, 0x8d, 0x5d, 0x15, 0x17, 0x9e, 0x16, 0x73, 0x92,
	0x5f, 0xe7, 0xd8, 0x85, 0xf3, 0x9c, 0xb4, 0x8b, 0x94, 0x87, 0x34, 0xa4, 0xe4, 0x45, 0x46, 0x6c,
	0xc3, 0x99, 0x28, 0x8b, 0x68, 0x75, 0xd0, 0xe2, 0xd6, 0xe6, 0xa3, 0x4b, 0x1e, 0xe5, 0x81, 0x9c,
	0xea, 0xe0, 0x15, 0x0c, 0x2a, 0xf7, 0xaa, 0xb6, 0xb6, 0xd6, 0x91, 0xdf, 0x7a, 0x74, 0x19, 0x3b,
	0x92, 0xa8, 0xff, 0xe0, 0x1e, 0x0e, 0xa1, 0x57, 0x4d, 0x4f, 0x75, 0xf9, 0x80, 0xf5, 0x99, 0x55,
	0xa3, 0xfb, 0xec, 0x82, 0xef, 0x75, 0x4c, 0xe1, 0x01, 0xe0, 0x7a, 0x03, 0x06, 0x72, 0x7c, 0x7e,
	0x00, 0xa9, 0x91, 0x0d, 0x0c, 0x86, 0x72, 0x29, 0x0c, 0xec, 0x22, 0xd7, 0xd3, 0x1d, 0x72, 0x89,
	0x09, 0x0c, 0xab, 0x0e, 0x3b, 0x43, 0xb3, 0x59, 0x6a, 0x25, 0xeb, 0x29, 0x5e, 0x42, 0x9f, 0xb3,
	0xf6, 0x21, 0x0f, 0x5f, 0xc9, 0x03, 0xd3, 0x73, 0xda, 0xcf, 0x4b, 0x38, 0x4f, 0xc0, 0x7e, 0xde,
	0x33, 0x6c, 0xc2, 0xa9, 0x10, 0xfe, 0x1b, 0xd5, 0xf3, 0x37, 0x2f, 0x7e, 0x6e, 0x92, 0xda, 0xaf,
	0x4d, 0x52, 0xfb, 0xbd, 0x49, 0x6a, 0x3f, 0xfe, 0x24, 0x07, 0xef, 0x7b, 0x69, 0xee, 0x6f, 0xd7,
	0xdf, 0x6f, 0xd7, 0xa3, 0xbb, 0xf5, 0xcd, 0x48, 0x3e, 0x85, 0xd1, 0xf2, 0xee, 0xf3, 0xc7, 0x63,
	0x59, 0xbe, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x76, 0xd7, 0x57, 0x48, 0x26, 0x03, 0x00, 0x00,
}
