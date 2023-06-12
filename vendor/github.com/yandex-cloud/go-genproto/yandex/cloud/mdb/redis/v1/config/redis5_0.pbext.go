// Code generated by protoc-gen-goext. DO NOT EDIT.

package redis

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *RedisConfig5_0) SetMaxmemoryPolicy(v RedisConfig5_0_MaxmemoryPolicy) {
	m.MaxmemoryPolicy = v
}

func (m *RedisConfig5_0) SetTimeout(v *wrapperspb.Int64Value) {
	m.Timeout = v
}

func (m *RedisConfig5_0) SetPassword(v string) {
	m.Password = v
}

func (m *RedisConfig5_0) SetDatabases(v *wrapperspb.Int64Value) {
	m.Databases = v
}

func (m *RedisConfig5_0) SetSlowlogLogSlowerThan(v *wrapperspb.Int64Value) {
	m.SlowlogLogSlowerThan = v
}

func (m *RedisConfig5_0) SetSlowlogMaxLen(v *wrapperspb.Int64Value) {
	m.SlowlogMaxLen = v
}

func (m *RedisConfig5_0) SetNotifyKeyspaceEvents(v string) {
	m.NotifyKeyspaceEvents = v
}

func (m *RedisConfig5_0) SetClientOutputBufferLimitPubsub(v *RedisConfig5_0_ClientOutputBufferLimit) {
	m.ClientOutputBufferLimitPubsub = v
}

func (m *RedisConfig5_0) SetClientOutputBufferLimitNormal(v *RedisConfig5_0_ClientOutputBufferLimit) {
	m.ClientOutputBufferLimitNormal = v
}

func (m *RedisConfig5_0_ClientOutputBufferLimit) SetHardLimit(v *wrapperspb.Int64Value) {
	m.HardLimit = v
}

func (m *RedisConfig5_0_ClientOutputBufferLimit) SetSoftLimit(v *wrapperspb.Int64Value) {
	m.SoftLimit = v
}

func (m *RedisConfig5_0_ClientOutputBufferLimit) SetSoftSeconds(v *wrapperspb.Int64Value) {
	m.SoftSeconds = v
}

func (m *RedisConfigSet5_0) SetEffectiveConfig(v *RedisConfig5_0) {
	m.EffectiveConfig = v
}

func (m *RedisConfigSet5_0) SetUserConfig(v *RedisConfig5_0) {
	m.UserConfig = v
}

func (m *RedisConfigSet5_0) SetDefaultConfig(v *RedisConfig5_0) {
	m.DefaultConfig = v
}
