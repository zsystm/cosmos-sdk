package valuerenderer

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type repeatedValueRenderer struct {
	tr      *Textual
	msgDesc protoreflect.MessageDescriptor
	vr      ValueRenderer
}

func NewRepeatedValueRenderer(t *Textual, msgDesc protoreflect.MessageDescriptor, v ValueRenderer) ValueRenderer {
	return &repeatedValueRenderer{tr: t, msgDesc: msgDesc, vr: v}
}

func (mr *repeatedValueRenderer) header(len int, fname string) string {
	return fmt.Sprintf("%s: %d %s", mr.msgDesc.Name(), len, fname)
}

func (mr *repeatedValueRenderer) Format(ctx context.Context, v protoreflect.Value) ([]Screen, error) {
	l := v.List()

	if l == nil {
		return nil, fmt.Errorf("non-List value")
	}

	if l.Len() == 0 {
		return []Screen{}, nil
	}

	screens := make([]Screen, 1)

	//Is this the best way to obtain the field name?
	screens[0].Text = mr.header(l.Len(), formatFieldName(string(l.Get(0).Message().Descriptor().Name())))

	for i := 0; i < l.Len(); i++ {
		subscreens, err := mr.vr.Format(ctx, l.Get(i))
		if err != nil {
			return nil, err
		}

		if len(subscreens) == 0 {
			return nil, fmt.Errorf("empty rendering")
		}

		headerScreen := Screen{
			Text:   fmt.Sprintf("%s (%d/%d): %s", formatFieldName(string(mr.msgDesc.Fields().Get(0).Name())), i+1, l.Len(), subscreens[0].Text),
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
