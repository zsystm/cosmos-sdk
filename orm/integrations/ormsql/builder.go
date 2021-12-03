package ormsql

import "google.golang.org/protobuf/encoding/protojson"

type builder struct {
	jsonMarshalOptions protojson.MarshalOptions
}

func newBuilder(jsonMarshalOptions protojson.MarshalOptions) *builder {
	return &builder{jsonMarshalOptions: jsonMarshalOptions}
}
