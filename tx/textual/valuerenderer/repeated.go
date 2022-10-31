package valuerenderer

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type repeatedValueRenderer struct {
	tr      *Textual
	msgDesc protoreflect.MessageDescriptor
	field   string
	kind    string
	vr      ValueRenderer
}

func NewRepeatedValueRenderer(t *Textual, msgDesc protoreflect.MessageDescriptor, field, kind string, v ValueRenderer) ValueRenderer {
	return &repeatedValueRenderer{
		tr:      t,
		msgDesc: msgDesc,
		field:   field,
		kind:    kind,
		vr:      v,
	}
}

func (mr *repeatedValueRenderer) name() string {
	name := mr.kind
	if mr.msgDesc != nil {
		name = string(mr.msgDesc.Name())
	}
	return name
}

func (mr *repeatedValueRenderer) header(len int) string {
	return fmt.Sprintf("%d %s", len, formatFieldName(mr.name()))
}

func (mr *repeatedValueRenderer) Format(ctx context.Context, v protoreflect.Value) ([]Screen, error) {
	l := v.List()

	if l == nil {
		return nil, fmt.Errorf("non-List value")
	}

	screens := make([]Screen, 1)
	screens[0].Text = mr.header(l.Len())

	for i := 0; i < l.Len(); i++ {
		subscreens, err := mr.vr.Format(ctx, l.Get(i))
		if err != nil {
			return nil, err
		}

		if len(subscreens) == 0 {
			return nil, fmt.Errorf("empty rendering")
		}

		headerScreen := Screen{
			Text:   fmt.Sprintf("%s (%d/%d): %s", formatFieldName(mr.field), i+1, l.Len(), subscreens[0].Text),
			Indent: subscreens[0].Indent + 1,
			Expert: subscreens[0].Expert,
		}
		screens = append(screens, headerScreen)

		for i := 1; i < len(subscreens); i++ {
			extraScreen := Screen{
				Text:   subscreens[i].Text,
				Indent: subscreens[i].Indent + 1,
				Expert: subscreens[i].Expert,
			}
			screens = append(screens, extraScreen)
		}
	}

	return screens, nil
}

func (mr *repeatedValueRenderer) Parse(ctx context.Context, screens []Screen) (protoreflect.Value, error) {
	return nilValue, fmt.Errorf("TODO")
}
