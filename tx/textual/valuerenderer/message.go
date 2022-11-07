package valuerenderer

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type messageValueRenderer struct {
	tr      *Textual
	msgDesc protoreflect.MessageDescriptor
	fds     []protoreflect.FieldDescriptor
}

func NewMessageValueRenderer(t *Textual, msgDesc protoreflect.MessageDescriptor) ValueRenderer {
	fields := msgDesc.Fields()
	fds := make([]protoreflect.FieldDescriptor, 0, fields.Len())
	for i := 0; i < fields.Len(); i++ {
		fds = append(fds, fields.Get(i))
	}
	sort.Slice(fds, func(i, j int) bool { return fds[i].Number() < fds[j].Number() })

	return &messageValueRenderer{tr: t, msgDesc: msgDesc, fds: fds}
}

func (mr *messageValueRenderer) header() string {
	return fmt.Sprintf("%s object", mr.msgDesc.Name())
}

func (mr *messageValueRenderer) Format(ctx context.Context, v protoreflect.Value) ([]Screen, error) {
	fullName := v.Message().Descriptor().FullName()
	wantFullName := mr.msgDesc.FullName()
	if fullName != wantFullName {
		return nil, fmt.Errorf(`bad message type: want "%s", got "%s"`, wantFullName, fullName)
	}

	screens := make([]Screen, 1)
	screens[0].Text = mr.header()

	for _, fd := range mr.fds {
		vr, err := mr.tr.GetValueRenderer(fd)
		if err != nil {
			return nil, err
		}
		// Skip default values.
		if !v.Message().Has(fd) {
			continue
		}

		subscreens := make([]Screen, 0)
		if fd.IsList() && !vr.handlesRepeated() {
			// If the field is a list, we need to format each element of the list
			subscreens, err = mr.formatRepeated(ctx, v.Message().Get(fd), fd)
		} else {
			// If the field is not list, we need to format the field
			subscreens, err = vr.Format(ctx, v.Message().Get(fd))
		}

		if err != nil {
			return nil, err
		}
		if len(subscreens) == 0 {
			return nil, fmt.Errorf("empty rendering for field %s", fd.Name())
		}

		headerScreen := Screen{
			Text:   fmt.Sprintf("%s: %s", toSentenceCase(string(fd.Name())), subscreens[0].Text),
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

func (mr *messageValueRenderer) formatRepeated(ctx context.Context, v protoreflect.Value, fd protoreflect.FieldDescriptor) ([]Screen, error) {
	vr, err := mr.tr.GetValueRenderer(fd)
	if err != nil {
		return nil, err
	}

	l := v.List()
	if l == nil {
		return nil, fmt.Errorf("non-List value")
	}

	screens := make([]Screen, 1)
	// <message_name>: <int> <field_kind>
	screens[0].Text = fmt.Sprintf("%d %s", l.Len(), toPluralKind(l.Len(), getKind(fd)))

	for i := 0; i < l.Len(); i++ {
		subscreens, err := vr.Format(ctx, l.Get(i))
		if err != nil {
			return nil, err
		}

		if len(subscreens) == 0 {
			return nil, fmt.Errorf("empty rendering")
		}

		headerScreen := Screen{
			// <field_name> (<int>/<int>): <value rendered 1st line>
			Text:   fmt.Sprintf("%s (%d/%d): %s", toSentenceCase(string(fd.Name())), i+1, l.Len(), subscreens[0].Text),
			Indent: subscreens[0].Indent + 1,
			Expert: subscreens[0].Expert,
		}
		screens = append(screens, headerScreen)

		// <optional value rendered in the next lines>
		for i := 1; i < len(subscreens); i++ {
			extraScreen := Screen{
				Text:   subscreens[i].Text,
				Indent: subscreens[i].Indent + 1,
				Expert: subscreens[i].Expert,
			}
			screens = append(screens, extraScreen)
		}
	}

	// End of <field_name>.
	terminalScreen := Screen{
		Text: fmt.Sprintf("End of %s", toSentenceCase(string(fd.Name()))),
	}
	screens = append(screens, terminalScreen)
	return screens, nil
}

// getKind returns the field kind: if the field is a protobuf
// message, then we return the message's name. Or else, we
// return the protobuf kind.
func getKind(fd protoreflect.FieldDescriptor) string {
	if fd.Kind() == protoreflect.MessageKind {
		return string(fd.Message().Name())
	}

	return fd.Kind().String()
}

// toPluralKind makes an honest attempt at making the kind plural, if
// the length is not one.  Note: It makes no attempts to handle the various oddities
// of pluralization.  For instance.. Oddity will become Odditys (instead of Oddities)
func toPluralKind(length int, kind string) string {
	formatted := toSentenceCase(kind)
	if length == 1 {
		return formatted
	}
	pluralized := []rune(formatted)
	pluralizedLen := utf8.RuneCountInString(formatted)
	lastRune := pluralized[pluralizedLen-1]
	ess := rune('s')
	if lastRune != ess {
		pluralized = append(pluralized, ess)
	}
	return string(pluralized)
}

// toSentenceCase formats a field name in sentence case, as specified in:
// https://github.com/cosmos/cosmos-sdk/blob/b6f867d0b674d62e56b27aa4d00f5b6042ebac9e/docs/architecture/adr-050-sign-mode-textual-annex1.md?plain=1#L110
func toSentenceCase(name string) string {
	if len(name) == 0 {
		return name
	}
	return strings.ToTitle(name[0:1]) + strings.ReplaceAll(name[1:], "_", " ")
}

var nilValue = protoreflect.Value{}

func (mr *messageValueRenderer) Parse(ctx context.Context, screens []Screen) (protoreflect.Value, error) {
	if len(screens) == 0 {
		return nilValue, fmt.Errorf("expect at least one screen")
	}

	wantHeader := fmt.Sprintf("%s object", mr.msgDesc.Name())
	if screens[0].Text != wantHeader {
		return nilValue, fmt.Errorf(`bad header: want "%s", got "%s"`, wantHeader, screens[0].Text)
	}
	if screens[0].Indent != 0 {
		return nilValue, fmt.Errorf("bad message indentation: want 0, got %d", screens[0].Indent)
	}

	msgType, err := protoregistry.GlobalTypes.FindMessageByName(mr.msgDesc.FullName())
	if err != nil {
		return nilValue, err
	}
	msg := msgType.New()
	idx := 1

	for _, fd := range mr.fds {
		if idx >= len(screens) {
			// remaining fields are default
			break
		}

		vr, err := mr.tr.GetValueRenderer(fd)
		if err != nil {
			return nilValue, err
		}

		if screens[idx].Indent != 1 {
			return nilValue, fmt.Errorf("bad message indentation: want 1, got %d", screens[idx].Indent)
		}

		prefix := toSentenceCase(string(fd.Name())) + ": "
		if !strings.HasPrefix(screens[idx].Text, prefix) {
			// we must have skipped this fd because of a default value
			continue
		}

		// Make a new screen without the prefix
		subscreens := make([]Screen, 1)
		subscreens[0] = screens[idx]
		subscreens[0].Text = strings.TrimPrefix(screens[idx].Text, prefix)
		subscreens[0].Indent--
		idx++

		// Gather nested screens
		for idx < len(screens) && screens[idx].Indent > 1 {
			scr := screens[idx]
			scr.Indent--
			subscreens = append(subscreens, scr)
			idx++
		}

		val, err := vr.Parse(ctx, subscreens)
		if err != nil {
			return nilValue, err
		}
		msg.Set(fd, val)
	}

	if idx > len(screens) {
		return nilValue, fmt.Errorf("leftover screens")
	}

	return protoreflect.ValueOfMessage(msg), nil
}

func (vr messageValueRenderer) handlesRepeated() bool { return false }
