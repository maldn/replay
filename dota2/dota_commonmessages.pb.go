// Code generated by protoc-gen-go.
// source: dota_commonmessages.proto
// DO NOT EDIT!

package dota2

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CDOTAMsg_LocationPing struct {
	X                *int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                *int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Target           *int32 `protobuf:"varint,3,opt,name=target" json:"target,omitempty"`
	DirectPing       *bool  `protobuf:"varint,4,opt,name=direct_ping" json:"direct_ping,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAMsg_LocationPing) Reset()         { *m = CDOTAMsg_LocationPing{} }
func (m *CDOTAMsg_LocationPing) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_LocationPing) ProtoMessage()    {}

func (m *CDOTAMsg_LocationPing) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_LocationPing) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_LocationPing) GetTarget() int32 {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return 0
}

func (m *CDOTAMsg_LocationPing) GetDirectPing() bool {
	if m != nil && m.DirectPing != nil {
		return *m.DirectPing
	}
	return false
}

type CDOTAMsg_MapLine struct {
	X                *int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                *int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Initial          *bool  `protobuf:"varint,3,opt,name=initial" json:"initial,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAMsg_MapLine) Reset()         { *m = CDOTAMsg_MapLine{} }
func (m *CDOTAMsg_MapLine) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_MapLine) ProtoMessage()    {}

func (m *CDOTAMsg_MapLine) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_MapLine) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_MapLine) GetInitial() bool {
	if m != nil && m.Initial != nil {
		return *m.Initial
	}
	return false
}

func init() {
}
