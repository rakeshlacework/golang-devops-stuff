// Code generated by protoc-gen-gogo.
// source: stream.proto
// DO NOT EDIT!

package warden

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type StreamRequest struct {
	Handle           *string `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	JobId            *uint32 `protobuf:"varint,2,req,name=job_id" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}

func (m *StreamRequest) GetHandle() string {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return ""
}

func (m *StreamRequest) GetJobId() uint32 {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return 0
}

type StreamResponse struct {
	Name             *string       `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Data             *string       `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	ExitStatus       *uint32       `protobuf:"varint,3,opt,name=exit_status" json:"exit_status,omitempty"`
	Info             *InfoResponse `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}

func (m *StreamResponse) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *StreamResponse) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func (m *StreamResponse) GetExitStatus() uint32 {
	if m != nil && m.ExitStatus != nil {
		return *m.ExitStatus
	}
	return 0
}

func (m *StreamResponse) GetInfo() *InfoResponse {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
}
