// Code generated by protoc-gen-go.
// source: v2ray.com/core/config.proto
// DO NOT EDIT!

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/config.proto

It has these top-level messages:
	AllocationStrategyConcurrency
	AllocationStrategyRefresh
	AllocationStrategy
	InboundConnectionConfig
	OutboundConnectionConfig
	Config
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_loader "v2ray.com/core/common/loader"
import v2ray_core_common_net "v2ray.com/core/common/net"
import v2ray_core_common_net1 "v2ray.com/core/common/net"
import v2ray_core_common_log "v2ray.com/core/common/log"
import v2ray_core_transport_internet "v2ray.com/core/transport/internet"
import v2ray_core_transport "v2ray.com/core/transport"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigFormat int32

const (
	ConfigFormat_Protobuf ConfigFormat = 0
	ConfigFormat_JSON     ConfigFormat = 1
)

var ConfigFormat_name = map[int32]string{
	0: "Protobuf",
	1: "JSON",
}
var ConfigFormat_value = map[string]int32{
	"Protobuf": 0,
	"JSON":     1,
}

func (x ConfigFormat) String() string {
	return proto.EnumName(ConfigFormat_name, int32(x))
}
func (ConfigFormat) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllocationStrategy_Type int32

const (
	// Always allocate all connection handlers.
	AllocationStrategy_Always AllocationStrategy_Type = 0
	// Randomly allocate specific range of handlers.
	AllocationStrategy_Random AllocationStrategy_Type = 1
	// External. Not supported yet.
	AllocationStrategy_External AllocationStrategy_Type = 2
)

var AllocationStrategy_Type_name = map[int32]string{
	0: "Always",
	1: "Random",
	2: "External",
}
var AllocationStrategy_Type_value = map[string]int32{
	"Always":   0,
	"Random":   1,
	"External": 2,
}

func (x AllocationStrategy_Type) String() string {
	return proto.EnumName(AllocationStrategy_Type_name, int32(x))
}
func (AllocationStrategy_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type AllocationStrategyConcurrency struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *AllocationStrategyConcurrency) Reset()                    { *m = AllocationStrategyConcurrency{} }
func (m *AllocationStrategyConcurrency) String() string            { return proto.CompactTextString(m) }
func (*AllocationStrategyConcurrency) ProtoMessage()               {}
func (*AllocationStrategyConcurrency) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllocationStrategyRefresh struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *AllocationStrategyRefresh) Reset()                    { *m = AllocationStrategyRefresh{} }
func (m *AllocationStrategyRefresh) String() string            { return proto.CompactTextString(m) }
func (*AllocationStrategyRefresh) ProtoMessage()               {}
func (*AllocationStrategyRefresh) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AllocationStrategy struct {
	Type AllocationStrategy_Type `protobuf:"varint,1,opt,name=type,enum=v2ray.core.AllocationStrategy_Type" json:"type,omitempty"`
	// Number of handlers (ports) running in parallel.
	Concurrency *AllocationStrategyConcurrency `protobuf:"bytes,2,opt,name=concurrency" json:"concurrency,omitempty"`
	// Number of minutes before a handler is regenerated.
	Refresh *AllocationStrategyRefresh `protobuf:"bytes,3,opt,name=refresh" json:"refresh,omitempty"`
}

func (m *AllocationStrategy) Reset()                    { *m = AllocationStrategy{} }
func (m *AllocationStrategy) String() string            { return proto.CompactTextString(m) }
func (*AllocationStrategy) ProtoMessage()               {}
func (*AllocationStrategy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AllocationStrategy) GetConcurrency() *AllocationStrategyConcurrency {
	if m != nil {
		return m.Concurrency
	}
	return nil
}

func (m *AllocationStrategy) GetRefresh() *AllocationStrategyRefresh {
	if m != nil {
		return m.Refresh
	}
	return nil
}

// Config for an inbound connection handler.
type InboundConnectionConfig struct {
	Settings               *v2ray_core_common_loader.TypedSettings     `protobuf:"bytes,1,opt,name=settings" json:"settings,omitempty"`
	PortRange              *v2ray_core_common_net.PortRange            `protobuf:"bytes,2,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
	ListenOn               *v2ray_core_common_net1.IPOrDomain          `protobuf:"bytes,3,opt,name=listen_on,json=listenOn" json:"listen_on,omitempty"`
	Tag                    string                                      `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	AllocationStrategy     *AllocationStrategy                         `protobuf:"bytes,5,opt,name=allocation_strategy,json=allocationStrategy" json:"allocation_strategy,omitempty"`
	StreamSettings         *v2ray_core_transport_internet.StreamConfig `protobuf:"bytes,6,opt,name=stream_settings,json=streamSettings" json:"stream_settings,omitempty"`
	AllowPassiveConnection bool                                        `protobuf:"varint,7,opt,name=allow_passive_connection,json=allowPassiveConnection" json:"allow_passive_connection,omitempty"`
}

func (m *InboundConnectionConfig) Reset()                    { *m = InboundConnectionConfig{} }
func (m *InboundConnectionConfig) String() string            { return proto.CompactTextString(m) }
func (*InboundConnectionConfig) ProtoMessage()               {}
func (*InboundConnectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InboundConnectionConfig) GetSettings() *v2ray_core_common_loader.TypedSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *InboundConnectionConfig) GetPortRange() *v2ray_core_common_net.PortRange {
	if m != nil {
		return m.PortRange
	}
	return nil
}

func (m *InboundConnectionConfig) GetListenOn() *v2ray_core_common_net1.IPOrDomain {
	if m != nil {
		return m.ListenOn
	}
	return nil
}

func (m *InboundConnectionConfig) GetAllocationStrategy() *AllocationStrategy {
	if m != nil {
		return m.AllocationStrategy
	}
	return nil
}

func (m *InboundConnectionConfig) GetStreamSettings() *v2ray_core_transport_internet.StreamConfig {
	if m != nil {
		return m.StreamSettings
	}
	return nil
}

type OutboundConnectionConfig struct {
	Settings       *v2ray_core_common_loader.TypedSettings     `protobuf:"bytes,1,opt,name=settings" json:"settings,omitempty"`
	SendThrough    *v2ray_core_common_net1.IPOrDomain          `protobuf:"bytes,2,opt,name=send_through,json=sendThrough" json:"send_through,omitempty"`
	StreamSettings *v2ray_core_transport_internet.StreamConfig `protobuf:"bytes,3,opt,name=stream_settings,json=streamSettings" json:"stream_settings,omitempty"`
	Tag            string                                      `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
}

func (m *OutboundConnectionConfig) Reset()                    { *m = OutboundConnectionConfig{} }
func (m *OutboundConnectionConfig) String() string            { return proto.CompactTextString(m) }
func (*OutboundConnectionConfig) ProtoMessage()               {}
func (*OutboundConnectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OutboundConnectionConfig) GetSettings() *v2ray_core_common_loader.TypedSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *OutboundConnectionConfig) GetSendThrough() *v2ray_core_common_net1.IPOrDomain {
	if m != nil {
		return m.SendThrough
	}
	return nil
}

func (m *OutboundConnectionConfig) GetStreamSettings() *v2ray_core_transport_internet.StreamConfig {
	if m != nil {
		return m.StreamSettings
	}
	return nil
}

type Config struct {
	Inbound   []*InboundConnectionConfig                `protobuf:"bytes,1,rep,name=inbound" json:"inbound,omitempty"`
	Outbound  []*OutboundConnectionConfig               `protobuf:"bytes,2,rep,name=outbound" json:"outbound,omitempty"`
	Log       *v2ray_core_common_log.Config             `protobuf:"bytes,3,opt,name=log" json:"log,omitempty"`
	App       []*v2ray_core_common_loader.TypedSettings `protobuf:"bytes,4,rep,name=app" json:"app,omitempty"`
	Transport *v2ray_core_transport.Config              `protobuf:"bytes,5,opt,name=transport" json:"transport,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Config) GetInbound() []*InboundConnectionConfig {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *Config) GetOutbound() []*OutboundConnectionConfig {
	if m != nil {
		return m.Outbound
	}
	return nil
}

func (m *Config) GetLog() *v2ray_core_common_log.Config {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Config) GetApp() []*v2ray_core_common_loader.TypedSettings {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Config) GetTransport() *v2ray_core_transport.Config {
	if m != nil {
		return m.Transport
	}
	return nil
}

func init() {
	proto.RegisterType((*AllocationStrategyConcurrency)(nil), "v2ray.core.AllocationStrategyConcurrency")
	proto.RegisterType((*AllocationStrategyRefresh)(nil), "v2ray.core.AllocationStrategyRefresh")
	proto.RegisterType((*AllocationStrategy)(nil), "v2ray.core.AllocationStrategy")
	proto.RegisterType((*InboundConnectionConfig)(nil), "v2ray.core.InboundConnectionConfig")
	proto.RegisterType((*OutboundConnectionConfig)(nil), "v2ray.core.OutboundConnectionConfig")
	proto.RegisterType((*Config)(nil), "v2ray.core.Config")
	proto.RegisterEnum("v2ray.core.ConfigFormat", ConfigFormat_name, ConfigFormat_value)
	proto.RegisterEnum("v2ray.core.AllocationStrategy_Type", AllocationStrategy_Type_name, AllocationStrategy_Type_value)
}

func init() { proto.RegisterFile("v2ray.com/core/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0xdd, 0x6e, 0xd3, 0x3e,
	0x18, 0xc6, 0x97, 0xb6, 0xeb, 0xda, 0xb7, 0xfb, 0xef, 0x5f, 0x19, 0x04, 0x61, 0x30, 0x54, 0xba,
	0xaf, 0xf2, 0xa1, 0x54, 0x14, 0x21, 0x3e, 0x24, 0x18, 0x5b, 0x07, 0xd2, 0x40, 0xa2, 0x95, 0xbb,
	0x23, 0x4e, 0x2a, 0x2f, 0xf5, 0xb2, 0x48, 0x89, 0x1d, 0x39, 0xee, 0x46, 0x2f, 0x81, 0x03, 0x2e,
	0x86, 0x5b, 0xe1, 0x8a, 0x90, 0x1d, 0x37, 0xcd, 0x68, 0x33, 0x26, 0x21, 0xce, 0x9c, 0xf8, 0xf9,
	0xbd, 0xb6, 0x9f, 0xe7, 0xb5, 0xe1, 0xee, 0x79, 0x47, 0x90, 0x89, 0xe3, 0xf2, 0xb0, 0xed, 0x72,
	0x41, 0xdb, 0x2e, 0x67, 0xa7, 0xbe, 0xe7, 0x44, 0x82, 0x4b, 0x8e, 0x60, 0x3a, 0x29, 0xe8, 0xfa,
	0xee, 0x9c, 0x30, 0x0c, 0x39, 0x6b, 0x07, 0x9c, 0x8c, 0xa8, 0x68, 0xcb, 0x49, 0x44, 0x13, 0x68,
	0x7d, 0x6b, 0xb1, 0x90, 0x51, 0xd9, 0x8e, 0xb8, 0x90, 0x46, 0xb5, 0x9b, 0xaf, 0x22, 0xa3, 0x91,
	0xa0, 0x71, 0x6c, 0x84, 0x3b, 0x79, 0xeb, 0x7a, 0x97, 0xf6, 0xba, 0xee, 0xfc, 0xa6, 0x93, 0x82,
	0xb0, 0x58, 0x2d, 0xd8, 0xf6, 0x99, 0xa4, 0x42, 0x15, 0xbe, 0xa4, 0xdf, 0xce, 0xd5, 0x67, 0x65,
	0xcd, 0xe7, 0xb0, 0xb1, 0x1f, 0x04, 0xdc, 0x25, 0xd2, 0xe7, 0x6c, 0x20, 0x05, 0x91, 0xd4, 0x9b,
	0x74, 0x39, 0x73, 0xc7, 0x42, 0x50, 0xe6, 0x4e, 0xd0, 0x4d, 0x58, 0x3e, 0x27, 0xc1, 0x98, 0xda,
	0x56, 0xc3, 0x6a, 0xfd, 0x87, 0x93, 0x8f, 0xe6, 0x53, 0xb8, 0x33, 0x8f, 0x61, 0x7a, 0x2a, 0x68,
	0x7c, 0x96, 0x83, 0x7c, 0x2b, 0x00, 0x9a, 0x67, 0xd0, 0x0b, 0x28, 0x29, 0x73, 0xb5, 0x76, 0xad,
	0xb3, 0xe9, 0xcc, 0x22, 0x71, 0xe6, 0xd5, 0xce, 0xf1, 0x24, 0xa2, 0x58, 0x03, 0xe8, 0x13, 0xd4,
	0xdc, 0xd9, 0x3e, 0xed, 0x42, 0xc3, 0x6a, 0xd5, 0x3a, 0x0f, 0xaf, 0xe6, 0x33, 0x07, 0xc3, 0x59,
	0x1a, 0xed, 0xc1, 0x8a, 0x48, 0x76, 0x6f, 0x17, 0x75, 0xa1, 0xed, 0xab, 0x0b, 0x99, 0xa3, 0xe2,
	0x29, 0xd5, 0x7c, 0x02, 0x25, 0xb5, 0x37, 0x04, 0x50, 0xde, 0x0f, 0x2e, 0xc8, 0x24, 0xae, 0x2f,
	0xa9, 0x31, 0x26, 0x6c, 0xc4, 0xc3, 0xba, 0x85, 0x56, 0xa1, 0xf2, 0xfe, 0xab, 0xca, 0x89, 0x04,
	0xf5, 0x42, 0xf3, 0x67, 0x11, 0x6e, 0x1f, 0xb1, 0x13, 0x3e, 0x66, 0xa3, 0x2e, 0x67, 0x8c, 0xba,
	0xaa, 0x76, 0x57, 0xe7, 0x82, 0xba, 0x50, 0x89, 0xa9, 0x94, 0x3e, 0xf3, 0x62, 0x6d, 0x4a, 0xad,
	0xb3, 0x9b, 0xdd, 0x4b, 0xd2, 0x1f, 0x4e, 0xd2, 0x97, 0xda, 0x8f, 0xd1, 0xc0, 0xc8, 0x71, 0x0a,
	0xa2, 0x3d, 0x00, 0x95, 0xf5, 0x50, 0x10, 0xe6, 0x51, 0xe3, 0x4d, 0x63, 0x41, 0x19, 0x46, 0xa5,
	0xd3, 0xe7, 0x42, 0x62, 0xa5, 0xc3, 0xd5, 0x68, 0x3a, 0x44, 0x6f, 0xa1, 0x1a, 0xf8, 0xb1, 0xa4,
	0x6c, 0xc8, 0x99, 0xb1, 0xe4, 0x41, 0x0e, 0x7f, 0xd4, 0xef, 0x89, 0x43, 0x1e, 0x12, 0x9f, 0xe1,
	0x4a, 0xc2, 0xf4, 0x18, 0xaa, 0x43, 0x51, 0x12, 0xcf, 0x2e, 0x35, 0xac, 0x56, 0x15, 0xab, 0x21,
	0xea, 0xc1, 0x0d, 0x92, 0xfa, 0x38, 0x8c, 0x8d, 0x91, 0xf6, 0xb2, 0xae, 0x7d, 0xff, 0x0f, 0x76,
	0x23, 0x32, 0xdf, 0x39, 0xc7, 0xf0, 0x7f, 0x2c, 0x05, 0x25, 0xe1, 0x30, 0xf5, 0xab, 0xac, 0x8b,
	0x3d, 0xce, 0x16, 0x4b, 0xfb, 0xde, 0x99, 0xde, 0x13, 0x67, 0xa0, 0xa9, 0xc4, 0x6e, 0xbc, 0x96,
	0xd4, 0x98, 0x7a, 0x88, 0x5e, 0x82, 0xad, 0xd6, 0xba, 0x18, 0x46, 0x24, 0x8e, 0xfd, 0x73, 0x3a,
	0x74, 0xd3, 0x80, 0xec, 0x95, 0x86, 0xd5, 0xaa, 0xe0, 0x5b, 0x7a, 0xbe, 0x9f, 0x4c, 0xcf, 0xe2,
	0x6b, 0x7e, 0x2f, 0x80, 0xdd, 0x1b, 0xcb, 0x7f, 0x98, 0xea, 0x21, 0xac, 0xc6, 0x94, 0x8d, 0x86,
	0xf2, 0x4c, 0xf0, 0xb1, 0x77, 0x66, 0x72, 0xbd, 0x46, 0x2e, 0x35, 0x85, 0x1d, 0x27, 0xd4, 0x22,
	0xdf, 0x8a, 0x7f, 0xef, 0xdb, 0x5c, 0xe0, 0xcd, 0x1f, 0x05, 0x28, 0x9b, 0xd3, 0xbf, 0x81, 0x15,
	0x3f, 0x69, 0x77, 0xdb, 0x6a, 0x14, 0x5b, 0xb5, 0xcb, 0xf7, 0x3c, 0xe7, 0x26, 0xe0, 0x29, 0x83,
	0xde, 0x41, 0x85, 0x1b, 0x63, 0xed, 0x82, 0xe6, 0xb7, 0xb2, 0x7c, 0x9e, 0xe9, 0x38, 0xa5, 0x50,
	0x1b, 0x8a, 0x01, 0xf7, 0xcc, 0x39, 0x37, 0x16, 0x3a, 0xef, 0x39, 0x86, 0x52, 0x4a, 0xf4, 0x0a,
	0x8a, 0x24, 0x8a, 0xec, 0x92, 0x5e, 0xed, 0xda, 0x51, 0x29, 0x06, 0xbd, 0x86, 0x6a, 0x6a, 0x9e,
	0x69, 0xef, 0x7b, 0x8b, 0x9d, 0x35, 0x0b, 0xce, 0xe4, 0x8f, 0x76, 0x60, 0x35, 0xf9, 0xf9, 0x81,
	0x8b, 0x90, 0x48, 0xf5, 0x6c, 0xf4, 0xd5, 0x3b, 0x7d, 0x32, 0x3e, 0xad, 0x2f, 0xa1, 0x0a, 0x94,
	0x3e, 0x0e, 0x7a, 0x9f, 0xeb, 0xd6, 0xc1, 0x26, 0xac, 0xb9, 0x3c, 0xcc, 0x54, 0x3d, 0xa8, 0x25,
	0x9c, 0x56, 0x7f, 0x29, 0xa9, 0x5f, 0x27, 0x65, 0xfd, 0xc4, 0x3f, 0xfb, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x1c, 0xfa, 0xca, 0xe6, 0x04, 0x07, 0x00, 0x00,
}
