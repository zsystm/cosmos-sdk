package canonical_proto3_json

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	gogoproto "github.com/gogo/protobuf/proto"
	proto2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

type gogoMessage interface {
	Descriptor() ([]byte, []int)
	Marshal() ([]byte, error)
}

// Register gogoproto imports with the Google Protobuf global file registry
func registerImports() {
	registerImport("gogoproto/gogo.proto")
	registerImport("google/protobuf/empty.proto")
	registerImport("google/protobuf/struct.proto")
	registerImport("testdata.proto")
}

// Reflect a gogoproto message to a Google Protobuf V2 message by manually
// serialising and deserialising using the FileDescriptor
func gogoReflect(m gogoMessage) protoreflect.Message {
	// Get the MessageDescriptor of t from the FileDescriptor
	gzippedpb, indices := m.Descriptor()

	fdProto, err := fileDescriptorProtoFromBytes(gzippedpb)
	if err != nil {
		panic(err.Error())
	}

	fd, err := protodesc.NewFile(fdProto, protoregistry.GlobalFiles)
	if err != nil {
		panic(err.Error())
	}

	md := messageDescriptorFromIndex(fd, indices)

	// Construct a new dynamic message from the MessageDescriptor
	dynMessage := dynamicpb.NewMessage(md)

	// Marshal the current message
	b, err := m.Marshal()
	if err != nil {
		panic(err.Error())
	}

	// Unmarshal it using the new dynamic message
	err = proto2.Unmarshal(b, dynMessage)
	if err != nil {
		panic(err.Error())
	}

	return dynMessage
}

func registerImport(filename string) {
	isGoGo := filename == "gogoproto/gogo.proto"
	if isGoGo {
		filename = "gogo.proto"
	}

	importBytes := gogoproto.FileDescriptor(filename)
	if len(importBytes) == 0 {
		panic("No import bytes!")
	}

	fdProto, err := fileDescriptorProtoFromBytes(importBytes)
	if err != nil {
		panic(err.Error())
	}

	if isGoGo {
		gogoName := "gogoproto/gogo.proto"
		fdProto.Name = &gogoName
	}

	fd, err := protodesc.NewFile(fdProto, protoregistry.GlobalFiles)
	if err != nil {
		panic(err.Error())
	}

	protoregistry.GlobalFiles.RegisterFile(fd)
}

func messageDescriptorFromIndex(fd protoreflect.FileDescriptor, indices []int) protoreflect.MessageDescriptor {
	md := fd.Messages().Get(indices[0])
	for _, i := range indices[1:] {
		md = md.Messages().Get(indices[i])
	}
	return md
}

func fileDescriptorProtoFromBytes(bs []byte) (*descriptorpb.FileDescriptorProto, error) {
	gzr, err := gzip.NewReader(bytes.NewReader(bs))
	if err != nil {
		panic(err.Error())
	}
	protoBlob, err := ioutil.ReadAll(gzr)
	if err != nil {
		panic(err.Error())
	}

	fdescproto := new(descriptorpb.FileDescriptorProto)
	if err := proto2.Unmarshal(protoBlob, fdescproto); err != nil {
		panic(err.Error())
	}

	return fdescproto, nil
}
