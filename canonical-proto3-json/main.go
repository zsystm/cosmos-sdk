package canonical_proto3_json

/*******************************************************************************

QUESTIONS:
- Do we want to support:
    - Protobuf: Any?
		- ANSWER: Yes, all of them
- Is it okay to copy bits of code from canonicaljson-go library?
	- ANSWER: Yes, but need to credit them in relevant files and check with Cory that all is good
- What are our performance needs? Is this naive implementation enough or is it
  worth optimising?
	- ANSWER: Yes, optimise
- Testing strategy?
	- ANSWER: Do some units tests with expected values. Also check some of the simple invariants. Property-testing for some subsets.

- Default values?
	- NullValue -> null or nothing?
	- Empty -> {} or nothing?
	- Map -> Default field value or nothing?

TODO:
- Check condition 3 of canonical JSON spec is met (UCS sorting of fields)
    - How do we do this? Google it!

NOTES:
- We'll eventually do this our own way, so in the long term will want to synchronise with Tyler's work

*******************************************************************************/

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	canonicaljson "github.com/gibson042/canonicaljson-go"
	"github.com/gogo/protobuf/types"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

const (
	hex                          = "0123456789ABCDEF"
	RFC3339TrailingZeroes string = "2006-01-02T15:04:05.000000000Z"
)

func main(t *testing.T) {
	registerImports()
	compareCanonicalJSON(t, &testTimeTest)
	compareCanonicalJSON(t, &testMsgCreateDog)
	compareCanonicalJSON(t, &testCat)
	compareCanonicalJSON(t, &testCompleteTest1)
	compareCanonicalJSON(t, &testCompleteTest2)
}

func compareCanonicalJSON(t *testing.T, m gogoMessage) {
	ourResult := MarshalMessage(gogoReflect(m))
	canonicalResult, err := canonicaljson.Marshal(m)
	if err != nil {
		// panic(err.Error())
		t.Logf("canonicaljson-go produced no results: %s", err.Error())
	}
	if bytes.Equal([]byte(ourResult), canonicalResult) {
		t.Log("---------- They match!! ----------")
		t.Log(ourResult)
	} else {
		t.Logf("---------- They're different ----------\nOurs:\n%s\nCanonical:\n%s\n", ourResult, canonicalResult)
	}
}

func isDefault(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
	switch kind := fd.Kind(); {
	case kind.GoString() == "EnumKind":
		return false
	case fd.IsList():
		return v.List().Len() == 0
	case fd.IsMap():
		return v.Map().Len() == 0
	case kind.GoString() == "MessageKind":
		switch m := v.Message(); m.Descriptor().FullName() {
		case "google.protobuf.Any":
			// TODO: Remove this to not filter out Any
			return true
		case "google.protobuf.BoolValue":
			return isWrapperDefault(m)
		case "google.protobuf.BytesValue":
			return isWrapperDefault(m)
		case "google.protobuf.DoubleValue":
			return isWrapperDefault(m)
		case "google.protobuf.Duration":
			fields := m.Descriptor().Fields()
			secs := m.Get(fields.ByName("seconds")).Int()
			nanos := m.Get(fields.ByName("nanos")).Int()
			return secs == 0 && nanos == 0
		case "google.protobuf.FloatValue":
			return isWrapperDefault(m)
		case "google.protobuf.Int32Value":
			return isWrapperDefault(m)
		case "google.protobuf.Int64Value":
			return isWrapperDefault(m)
		case "google.protobuf.StringValue":
			return isWrapperDefault(m)
		case "google.protobuf.Struct":
			fields := m.Descriptor().Fields()
			structFields := m.Get(fields.ByName("fields")).Map()
			return structFields.Len() == 0
		case "google.protobuf.Timestamp":
			fields := m.Descriptor().Fields()
			secs := m.Get(fields.ByName("seconds")).Int()
			nanos := m.Get(fields.ByName("nanos")).Int()
			return secs == 0 && nanos == 0
		case "google.protobuf.UInt32Value":
			return isWrapperDefault(m)
		case "google.protobuf.UInt64Value":
			return isWrapperDefault(m)
		}
	}

	return false
}

func isWrapperDefault(m protoreflect.Message) bool {
	valueFieldDescriptor := m.Descriptor().Fields().ByName("value")
	value := m.Get(valueFieldDescriptor)
	switch valueFieldDescriptor.Kind().GoString() {
	case "BoolKind":
		return value.Bool() == false
	case "BytesKind":
		return len(value.Bytes()) == 0
	case "DoubleKind", "FloatKind":
		return value.Float() == 0.0
	case "Int32Kind", "Int64Kind":
		return value.Int() == 0
	case "Uint32Kind", "Uint64Kind":
		return value.Uint() == 0
	case "StringKind":
		return value.String() == ""
	default:
		return false
	}
}

func collectFields(m map[string]string) func(protoreflect.FieldDescriptor, protoreflect.Value) bool {
	return func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if !isDefault(fd, v) {
			m[MarshalString(fd.TextName())] = MarshalField(fd, v)
		}
		return true
	}
}

func collectMapFields(valueFieldDescriptor protoreflect.FieldDescriptor, m map[string]string) func(protoreflect.MapKey, protoreflect.Value) bool {
	return func(k protoreflect.MapKey, v protoreflect.Value) bool {
		m[MarshalString(k.String())] = MarshalField(valueFieldDescriptor, v)
		return true
	}
}

func JSONValueString(kind protoreflect.Kind, v protoreflect.Value) string {
	var sb strings.Builder
	switch kind.GoString() {
	case "BoolKind":
		sb.WriteString(v.String())
	case "BytesKind":
		sb.WriteByte('"')
		bytes := v.Bytes()
		encoded := make([]byte, base64.StdEncoding.EncodedLen(len(bytes)))
		base64.StdEncoding.Encode(encoded, bytes)
		sb.Write(encoded)
		sb.WriteByte('"')
	case "DoubleKind":
		sb.WriteString(normalizeNumber(strconv.FormatFloat(v.Float(), 'E', -1, 64)))
	case "EnumKind":
		panic("Should already have processed Enum")
	case "FloatKind":
		sb.WriteString(normalizeNumber(strconv.FormatFloat(v.Float(), 'E', -1, 32)))
	case "GroupKind":
	case "Fixed32Kind", "Int32Kind", "Sfixed32Kind", "Sint32Kind", "Uint32Kind":
		sb.WriteString(v.String())
	case "Fixed64Kind", "Int64Kind", "Sfixed64Kind", "Sint64Kind", "Uint64Kind":
		sb.WriteByte('"')
		sb.WriteString(v.String())
		sb.WriteByte('"')
	case "MessageKind":
		sb.WriteString(MarshalMessage(v.Message()))
	case "StringKind":
		sb.WriteString(MarshalString(v.String()))
	}
	return sb.String()
}

func MarshalField(fd protoreflect.FieldDescriptor, v protoreflect.Value) string {
	switch kind := fd.Kind(); {
	case kind.GoString() == "EnumKind":
		return MarshalEnum(fd.Enum(), v.Enum())
	case fd.IsList():
		return MarshalList(kind, v.List())
	case fd.IsMap():
		return MarshalMap(fd.MapValue(), v.Map())
	default:
		return JSONValueString(kind, v)
	}
}

func MarshalEnum(ed protoreflect.EnumDescriptor, enumNumber protoreflect.EnumNumber) string {
	if ed.FullName() == "google.protobuf.NullValue" {
		return "null"
	} else {
		var sb strings.Builder
		sb.WriteByte('"')
		sb.WriteString(string(ed.Values().Get(int(enumNumber)).Name()))
		sb.WriteByte('"')
		return sb.String()
	}
}

func MarshalMessage(m protoreflect.Message) string {
	switch m.Descriptor().FullName() {
	case "google.protobuf.Any":
		return "TODO"
	case "google.protobuf.BoolValue":
		return MarshalWrapper(m)
	case "google.protobuf.BytesValue":
		return MarshalWrapper(m)
	case "google.protobuf.DoubleValue":
		return MarshalWrapper(m)
	case "google.protobuf.Duration":
		return MarshalDuration(m)
	case "google.protobuf.Empty":
		return "{}"
	case "google.protobuf.FieldMask":
		return MarshalFieldMask(m)
	case "google.protobuf.FloatValue":
		return MarshalWrapper(m)
	case "google.protobuf.Int32Value":
		return MarshalWrapper(m)
	case "google.protobuf.Int64Value":
		return MarshalWrapper(m)
	case "google.protobuf.ListValue":
		return MarshalListValue(m)
	case "google.protobuf.StringValue":
		return MarshalWrapper(m)
	case "google.protobuf.Struct":
		return MarshalStruct(m)
	case "google.protobuf.Timestamp":
		return MarshalTimestamp(m)
	case "google.protobuf.UInt32Value":
		return MarshalWrapper(m)
	case "google.protobuf.UInt64Value":
		return MarshalWrapper(m)
	case "google.protobuf.Value":
		return MarshalValue(m)
	default:
		// Collect fields into a map
		messageMap := make(map[string]string)
		m.Range(collectFields(messageMap))

		return JSONObjectString(messageMap)
	}
}

func MarshalDuration(m protoreflect.Message) string {
	fields := m.Descriptor().Fields()
	secs := m.Get(fields.ByName("seconds")).Int()
	nanos := m.Get(fields.ByName("nanos")).Int()
	return fmt.Sprintf("\"%v.%09ds\"", secs, nanos)
}

func MarshalFieldMask(m protoreflect.Message) string {
	pathsFieldDescriptor := m.Descriptor().Fields().ByName("paths")
	pathsList := m.Get(pathsFieldDescriptor).List()

	var sb strings.Builder

	for i := 0; i < pathsList.Len(); i++ {
		sb.WriteString(pathsList.Get(i).String())
		if i != pathsList.Len()-1 {
			sb.WriteByte(',')
		}
	}

	return MarshalString(sb.String())
}

func MarshalListValue(m protoreflect.Message) string {
	valuesFieldDescriptor := m.Descriptor().Fields().ByName("values")
	valuesList := m.Get(valuesFieldDescriptor).List()
	return MarshalList(valuesFieldDescriptor.Kind(), valuesList)
}

func MarshalStruct(m protoreflect.Message) string {
	fields := m.Descriptor().Fields()
	structFieldsDescriptor := fields.ByName("fields")
	structFields := m.Get(structFieldsDescriptor).Map()
	return MarshalMap(structFieldsDescriptor.MapValue(), structFields)
}

func MarshalTimestamp(m protoreflect.Message) string {
	fields := m.Descriptor().Fields()
	secs := m.Get(fields.ByName("seconds")).Int()
	nanos := m.Get(fields.ByName("nanos")).Int()
	t := time.Unix(secs, nanos)
	return fmt.Sprintf("\"%s\"", t.UTC().Format(RFC3339TrailingZeroes))
}

func MarshalValue(m protoreflect.Message) string {
	kindOneOf := m.Descriptor().Oneofs().ByName("kind")
	kindFieldDescriptor := m.WhichOneof(kindOneOf)
	return MarshalField(kindFieldDescriptor, m.Get(kindFieldDescriptor))
}

func MarshalWrapper(m protoreflect.Message) string {
	valueFieldDescriptor := m.Descriptor().Fields().ByName("value")
	return MarshalField(valueFieldDescriptor, m.Get(valueFieldDescriptor))
}

func MarshalMap(valueFieldDescriptor protoreflect.FieldDescriptor, m protoreflect.Map) string {
	// Collect fields into a map
	fieldMap := make(map[string]string)
	m.Range(collectMapFields(valueFieldDescriptor, fieldMap))

	return JSONObjectString(fieldMap)
}

func MarshalList(listValueKind protoreflect.Kind, l protoreflect.List) string {
	var sb strings.Builder

	sb.WriteByte('[')
	for i := 0; i < l.Len(); i++ {
		sb.WriteString(JSONValueString(listValueKind, l.Get(i)))
		if i != l.Len()-1 {
			sb.WriteByte(',')
		}
	}
	sb.WriteByte(']')

	return sb.String()
}

func JSONObjectString(fieldMap map[string]string) string {
	numFields := len(fieldMap)

	// Sort keys (TODO: Make sure this is the UCS sorting we need)
	keys := make([]string, 0, numFields)
	for k := range fieldMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Output the map in sorted order
	var sb strings.Builder
	sb.WriteByte('{')
	for i, k := range keys {
		sb.WriteString(k)
		sb.WriteByte(':')
		sb.WriteString(fieldMap[k])
		if i != numFields-1 {
			sb.WriteByte(',')
		}
	}
	sb.WriteByte('}')

	return sb.String()
}

// Force capital-E exponent, remove + signs and leading zeroes
// Source: github.com/gibson042/canonicaljson-go
var expNormalizer = regexp.MustCompile("(?:E(?:[+]0*|(-|)0+)|e(?:[+]|(-|))0*)([0-9])")
var expNormalizerReplacement = "E$1$2$3"

// Find the significant digits in a numeric string
// Source: github.com/gibson042/canonicaljson-go
var significantDigits = regexp.MustCompile(`^(?:0(?:\.0*|)|)([0-9.]+?)(?:0*\.?0*|)(?:E|$)`)

// normalizeNumber normalizes a valid JSON numeric string and writes the result on its encoder.
// Source: github.com/gibson042/canonicaljson-go (modified to fit our serialiser)
func normalizeNumber(s string) string {
	s = expNormalizer.ReplaceAllString(s, expNormalizerReplacement)

	// detect negative sign
	negative := s[0] == '-'
	if negative {
		s = s[1:]
	}

	// detect exponent
	exp := 0
	expPos := strings.LastIndexByte(s, 'E')
	if expPos == -1 {
		expPos = len(s)
	} else {
		var err error
		exp, err = strconv.Atoi(s[expPos+1:])
		if err != nil {
			panic(err.Error())
		}
	}

	// characterize the significant digits, finishing early on zero
	significantRange := significantDigits.FindStringSubmatchIndex(s)
	if significantRange == nil {
		panic(fmt.Errorf("canonical-proto3-json: no significand in %q", s))
	}
	significantRange = significantRange[2:4]
	if s[significantRange[0]] == '0' {
		return "0"
	}
	pointPos := strings.IndexByte(s, '.')
	effectivePointPos := pointPos
	if effectivePointPos == -1 {
		effectivePointPos = expPos
	}
	var significand string
	if pointPos > significantRange[0] && pointPos < significantRange[1] {
		significand = s[significantRange[0]:pointPos] + s[pointPos+1:significantRange[1]]
	} else {
		significand = s[significantRange[0]:significantRange[1]]
	}

	// detect integers (i.e., significand fits within exponent)
	isInt := (exp <= 0 && (effectivePointPos+exp) >= significantRange[1]) ||
		(exp > 0 && (effectivePointPos+exp+1) >= significantRange[1])

	// effective exponent increases/decreases by excess/deficient significand magnitude
	if effectivePointPos > 1 {
		exp += effectivePointPos - 1
	} else if significantRange[0] > 0 {
		exp -= significantRange[0] - 1
	}

	// write result
	var sb strings.Builder
	if negative {
		sb.WriteByte('-')
	}
	if isInt {
		// integer: render without exponent
		sb.WriteString(significand)
		if len(significand) < exp+1 {
			sb.WriteString(strings.Repeat("0", exp+1-len(significand)))
		}
	} else {
		// non-integer: render minimal significand with decimal point and non-empty exponent
		if len(significand) == 1 {
			significand += "0"
		}
		sb.WriteString(significand[:1] + "." + significand[1:] + "E" + strconv.Itoa(exp))
	}

	return sb.String()
}

// Marshal a string making sure it uses minimal UTF-8 encoding
// Source: github.com/gibson042/canonicaljson-go (modified to fit our serialiser)
func MarshalString(s string) string {
	return MarshalStringBytes([]byte(s))
}

// Marshal a string making sure it uses minimal UTF-8 encoding
// Source: github.com/gibson042/canonicaljson-go (modified to fit our serialiser)
func MarshalStringBytes(s []byte) string {
	var sb strings.Builder
	sb.WriteByte('"')
	start := 0
	rejectLowSurrogateAt := -1
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if 0x20 <= b && b != '\\' && b != '"' {
				i++
				continue
			}
			if start < i {
				sb.Write(s[start:i])
			}
			sb.WriteByte('\\')
			switch b {
			case '\\', '"':
				sb.WriteByte(b)
			case '\n':
				sb.WriteByte('n')
			case '\r':
				sb.WriteByte('r')
			case '\t':
				sb.WriteByte('t')
			case '\x08': // \b
				sb.WriteByte('b')
			case '\x0c': // \f
				sb.WriteByte('f')
			default:
				// This encodes other bytes < 0x20 as \u00xx.
				sb.WriteString("u00")
				sb.WriteByte(hex[b>>4])
				sb.WriteByte(hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRune(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				sb.Write(s[start:i])
			}

			// Check for a "WTF-8" lone surrogate (and escape it).
			// High surrogate (U+D800 through U+DBFF): 1110 1101   10 10bbbb   10 bbbbbb
			// Low surrogate  (U+DC00 through U+DFFF): 1110 1101   10 11bbbb   10 bbbbbb
			if s[i] == 0xED && i+2 < len(s) {
				c2 := s[i+1]
				c3 := s[i+2]
				if c2 >= 0xA0 && c2 <= 0xBF && (c3&0xC0) == 0x80 {
					// Don't let this special-case logic sneak through a valid surrogate pair.
					if isHigh := (c2 & 0x10) == 0; isHigh || i != rejectLowSurrogateAt {
						if isHigh {
							rejectLowSurrogateAt = i + 3
						}
						sb.WriteString(`\u`)
						sb.WriteByte(hex[0xD])
						sb.WriteByte(hex[(c2>>2)&0x0F])
						sb.WriteByte(hex[((c2<<2)&0x0C)|((c3>>4)&0x03)])
						sb.WriteByte(hex[c3&0x0F])
						i += 3
						start = i
						continue
					}
				}
			}

			panic(fmt.Errorf("Unsupported value: %q", string(s)))
		}
		i += size
	}
	if start < len(s) {
		sb.Write(s[start:])
	}
	sb.WriteByte('"')
	return sb.String()
}

var (
	testTimeTest = TimeTest{
		TestTimestamp: &testTime,
	}

	testMsgCreateDog = testdata.MsgCreateDog{
		Dog: &testdata.Dog{
			Size_: "big",
			Name:  "Woofer",
		},
	}

	testCat = testdata.Cat{
		Moniker: "Catface",
		Lives:   136,
	}

	testCompleteTest1 = CompleteTest{
		NumericalTest:       &CompleteTest_NumericalTest{},
		TestEnum:            0,
		TestMap:             map[int32]int64{},
		TestTimeMap:         map[int32]*types.Duration{},
		TestRepeatedFixed32: []uint32{},
		TestBool:            false,
		TestString:          "",
		TestBytes:           []byte{},
		TestAny:             &cosmostypes.Any{},
		TestTimestamp:       nil,
		TestStdTimestamp:    &time.Time{},
		TestDuration:        &types.Duration{},
		TestStruct:          &types.Struct{},
		TestBoolValue:       nil,
		TestBytesValue:      &types.BytesValue{},
		TestDoubleValue:     &types.DoubleValue{},
		TestFloatValue:      &types.FloatValue{},
		TestInt32Value:      &types.Int32Value{},
		TestInt64Value:      &types.Int64Value{},
		TestStringValue:     &types.StringValue{},
		TestUint32Value:     &types.UInt32Value{},
		TestUint64Value:     &types.UInt64Value{},
		TestFieldMask:       nil,
		TestEmpty:           nil,
	}

	testTime         = time.Unix(10000, 100)
	testTimestamp, _ = types.TimestampProto(testTime)
	testDuration     = types.DurationProto(time.Since(time.Now()))

	testCompleteTest2 = CompleteTest{
		NumericalTest: &CompleteTest_NumericalTest{
			TestInt32:   150,
			TestFixed32: 200,
			TestUint32:  300,
			TestInt64:   -24,
			TestFixed64: -0,
			TestUint64:  100,
			TestFloat:   -1.0000,
			TestDouble:  4.40001e2,
		},
		TestEnum:            1,
		TestMap:             map[int32]int64{32: 48, -0: 2e3},
		TestTimeMap:         map[int32]*types.Duration{400: testDuration},
		TestRepeatedFixed32: []uint32{123, 456},
		TestBool:            true,
		TestString:          "hello",
		TestBytes:           []byte{'1', '2', '3'},
		TestAny:             &cosmostypes.Any{},
		TestTimestamp:       testTimestamp,
		TestStdTimestamp:    &testTime,
		TestDuration:        testDuration,
		TestStruct: &types.Struct{
			Fields: map[string]*types.Value{
				"test_struct_bool": {
					Kind: &types.Value_BoolValue{
						BoolValue: true,
					},
				},
				"test_struct_list": {
					Kind: &types.Value_ListValue{
						ListValue: &types.ListValue{
							Values: []*types.Value{
								{
									Kind: &types.Value_StringValue{
										StringValue: "Hello!",
									},
								},
							},
						},
					},
				},
				"test_struct_null": {
					Kind: &types.Value_NullValue{
						NullValue: 0,
					},
				},
			},
		},
		TestBoolValue: &types.BoolValue{
			Value: true,
		},
		TestBytesValue: &types.BytesValue{
			Value: []byte{0x1},
		},
		TestDoubleValue: &types.DoubleValue{
			Value: 0.0003,
		},
		TestFloatValue: &types.FloatValue{
			Value: 0.00004,
		},
		TestInt32Value: &types.Int32Value{
			Value: 13456,
		},
		TestInt64Value: &types.Int64Value{
			Value: -1000000000002,
		},
		TestStringValue: &types.StringValue{
			Value: "abcdefg",
		},
		TestUint32Value: &types.UInt32Value{
			Value: 123456,
		},
		TestUint64Value: &types.UInt64Value{
			Value: 12346576869779,
		},
		TestFieldMask: &types.FieldMask{
			Paths: []string{"f.baz", "h", "g.foo.bar"},
		},
		TestEmpty: &types.Empty{},
	}
)
