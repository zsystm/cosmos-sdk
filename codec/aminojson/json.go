package aminojson

import (
	"bytes"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protorange"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Marshal(message proto.Message) ([]byte, error) {
	// Print a message in a humanly readable format.
	var indent []byte
	buf := &bytes.Buffer{}
	err := protorange.Options{
		Stable: true,
	}.Range(message.ProtoReflect(),
		func(p protopath.Values) error {
			// Print the key.
			var fd protoreflect.FieldDescriptor
			last := p.Index(-1)
			beforeLast := p.Index(-2)
			switch last.Step.Kind() {
			case protopath.FieldAccessStep:
				fd = last.Step.FieldDescriptor()
				fmt.Fprintf(buf, "%s%q: ", indent, fd.Name())
			case protopath.ListIndexStep:
				fd = beforeLast.Step.FieldDescriptor() // lists always appear in the context of a repeated field
				fmt.Fprintf(buf, "%s%d: ", indent, last.Step.ListIndex())
			case protopath.MapIndexStep:
				fd = beforeLast.Step.FieldDescriptor() // maps always appear in the context of a repeated field
				fmt.Fprintf(buf, "%s%v: ", indent, last.Step.MapIndex().Interface())
			case protopath.AnyExpandStep:
				fmt.Fprintf(buf, "%s[%v]: ", indent, last.Value.Message().Descriptor().FullName())
			case protopath.UnknownAccessStep:
				fmt.Fprintf(buf, "%s?: ", indent)
			}

			// Starting printing the value.
			switch v := last.Value.Interface().(type) {
			case protoreflect.Message:
				fmt.Fprintf(buf, "{\n")
				indent = append(indent, '\t')
			case protoreflect.List:
				fmt.Fprintf(buf, "[\n")
				indent = append(indent, '\t')
			case protoreflect.Map:
				fmt.Fprintf(buf, "{\n")
				indent = append(indent, '\t')
			case protoreflect.EnumNumber:
				var ev protoreflect.EnumValueDescriptor
				if fd != nil {
					ev = fd.Enum().Values().ByNumber(v)
				}
				if ev != nil {
					fmt.Fprintf(buf, "%v\n", ev.Name())
				} else {
					fmt.Fprintf(buf, "%v\n", v)
				}
			case string, []byte:
				fmt.Fprintf(buf, "%q,\n", v)
			default:
				fmt.Fprintf(buf, "%v,\n", v)
			}
			return nil
		},
		func(p protopath.Values) error {
			// Finish printing the value.
			last := p.Index(-1)
			switch last.Value.Interface().(type) {
			case protoreflect.Message:
				indent = indent[:len(indent)-1]
				fmt.Fprintf(buf, "%s},\n", indent)
			case protoreflect.List:
				indent = indent[:len(indent)-1]
				fmt.Fprintf(buf, "%s],\n", indent)
			case protoreflect.Map:
				indent = indent[:len(indent)-1]
				fmt.Fprintf(buf, "%s},\n", indent)
			}
			return nil
		},
	)
	return buf.Bytes(), err
}
