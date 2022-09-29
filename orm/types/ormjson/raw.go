package ormjson

import (
	"bytes"
	"encoding/json"
	"io"

	"google.golang.org/protobuf/runtime/protoiface"

	"cosmossdk.io/core/appmodule"
)

type rawMessageSource struct {
	m map[string]json.RawMessage
}

// NewRawMessageSource returns a new ReadSource for the provided
// json.RawMessage where it is assumed that the raw message is a JSON
// map where each table's JSON referenced by the map key corresponding
// to the tables full protobuf name.
func NewRawMessageSource(message json.RawMessage) (appmodule.GenesisSource, error) {
	var m map[string]json.RawMessage
	err := json.Unmarshal(message, &m)
	if err != nil {
		return nil, err
	}
	return &rawMessageSource{m}, err
}

func (r rawMessageSource) OpenReader(field string) (io.ReadCloser, error) {
	j, ok := r.m[field]
	if !ok {
		return nil, nil
	}
	return readCloserWrapper{bytes.NewReader(j)}, nil
}

func (r rawMessageSource) ReadMessage(v1 protoiface.MessageV1) error {
	//TODO implement me
	panic("implement me")
}

func (r rawMessageSource) ReadRawJSON() (json.RawMessage, error) {
	//TODO implement me
	panic("implement me")
}

type readCloserWrapper struct {
	io.Reader
}

func (r readCloserWrapper) Close() error { return nil }

var _ appmodule.GenesisSource = rawMessageSource{}

// RawMessageTarget is a WriteTarget wrapping a raw JSON map.
type RawMessageTarget struct {
	m map[string]json.RawMessage
}

// NewRawMessageTarget returns a new WriteTarget where each table's JSON
// is written to a map key corresponding to the table's full protobuf name.
func NewRawMessageTarget() *RawMessageTarget {
	return &RawMessageTarget{}
}

func (r *RawMessageTarget) OpenWriter(field string) (io.WriteCloser, error) {
	if r.m == nil {
		r.m = map[string]json.RawMessage{}
	}

	return &rawWriter{Buffer: &bytes.Buffer{}, sink: r, field: field}, nil
}

func (r *RawMessageTarget) WriteMessage(v1 protoiface.MessageV1) error {
	//TODO implement me
	panic("implement me")
}

func (r *RawMessageTarget) WriteRawJSON(message json.RawMessage) error {
	//TODO implement me
	panic("implement me")
}

// JSON returns the JSON map that was written as a json.RawMessage.
func (r *RawMessageTarget) JSON() (json.RawMessage, error) {
	return json.MarshalIndent(r.m, "", "  ")
}

type rawWriter struct {
	*bytes.Buffer
	field string
	sink  *RawMessageTarget
}

func (r rawWriter) Close() error {
	r.sink.m[r.field] = r.Buffer.Bytes()
	return nil
}

var _ appmodule.GenesisTarget = &RawMessageTarget{}
