package canonical_proto3_json

import (
	"strconv"
	"testing"
	"unicode/utf8"

	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

func TestMain(t *testing.T) {
	main(t)
}

// Point 1 of the canonical JSON spec:
// "JSON text in canonical form MUST be encoded in UTF-8"
func testValidUTF8(t *rapid.T) {
	completeTest := genCompleteTest.Draw(t, "completeTest").(*CompleteTest)
	res := MarshalMessage(gogoReflect(completeTest))
	require.True(t, utf8.ValidString(res))
}

// Point 2 of the canonical JSON spec:
// "JSON text in canonical form MUST NOT include insignificant (i.e.,
// inter-token) whitespace (defined in section 2 of RFC 7159)"
func testInsignificantWhitespace(t *rapid.T) {
	// TODO
}

// Point 3 of the canonical JSON spec:
// "JSON text in canonical form MUST order the members of all objects
// lexicographically by the UCS (Unicode Character Set) code points of their
// names"
// TODO: Fix this! We're getting a nice bug catch here :)
// TODO: Check point 3.1
func testLexicographicalOrdering(t *rapid.T) {
	completeTest := genCompleteTest.Draw(t, "completeTest").(*CompleteTest)
	res := MarshalMessage(gogoReflect(completeTest))
	t.Log(res)

	err := jsonparser.ObjectEach([]byte(res), orderingHandler(t))
	require.NoError(t, err)
}

// orderingHandler creates a callback to check that JSON object keys are in
// ascending lexicographical order by Unicode codepoint
func orderingHandler(t *rapid.T) func([]byte, []byte, jsonparser.ValueType, int) error {
	var last []byte
	return func(key []byte, value []byte, dataType jsonparser.ValueType, _ int) error {
		// Special case for when last is empty and our first key might be empty
		// TODO: Check that we allow empty keys
		// TODO: Check that we disallow duplicate keys
		if len(last) == 0 {
			last = key
			return nil
		}

		lastRunes := []rune(string(last))
		keyRunes := []rune(string(key))

		// Require that each key is greater than the last
		require.Equal(t, -1, compareRunes(lastRunes, keyRunes))

		// Recurse into sub-objects
		if dataType == jsonparser.Object {
			err := jsonparser.ObjectEach(value, orderingHandler(t))
			if err != nil {
				return err
			}
		}

		// Set the last value to the new key
		last = key

		return nil
	}
}

// compareRunes does a lexicographic comparison on two slices of runes
func compareRunes(a []rune, b []rune) int {
	// Always compare with a shorter than b to avoid array index overflow
	if len(a) > len(b) {
		return -compareRunes(b, a)
	}

	// Run through and compare lexicographically
	for i := range a {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}

	// If they're the same length, they're equal, otherwise the longer one
	// is greater
	if len(a) == len(b) {
		return 0
	} else {
		return -1
	}
}

// Point 4 of the canonical JSON spec
func testIntegerNumbers(t *rapid.T) {

}

// Point 5 of the canonical JSON spec
func testNonIntegerNumbers(t *rapid.T) {

}

// Point 6 of the canonical JSON spec
func testMinimalStringEncoding(t *rapid.T) {

}

func testNormalizeNumber(t *rapid.T) {
	n := rapid.Int64().Draw(t, "n").(int64)
	normalizedN := normalizeNumber(strconv.FormatInt(n, 10))
	require.Equal(t, normalizedN, normalizeNumber(normalizedN))
}

// // Right now we can't say this is true, because of some escaping stuff
// func testNormalizedString(t *rapid.T) {
// 	normalizeString := func(s string) string {
// 		return strings.TrimPrefix(strings.TrimSuffix(MarshalString(s), "\""), "\"")
// 	}

// 	s := rapid.String().Draw(t, "s").(string)
// 	normalizedS := normalizeString(s)

// 	// normalizeString(s) == normalizeString(normalizeString(s))
// 	require.Equal(t, normalizedS, normalizeString(normalizedS))
// }

func TestProperties(t *testing.T) {
	// 6 points of the canonical JSON spec
	t.Run("TestValidUTF8", rapid.MakeCheck(testValidUTF8))
	t.Run("TestInsignificantWhitespace", rapid.MakeCheck(testInsignificantWhitespace))
	t.Run("TestLexicographicalOrdering", rapid.MakeCheck(testLexicographicalOrdering))
	t.Run("TestIntegerNumbers", rapid.MakeCheck(testIntegerNumbers))
	t.Run("TestNonIntegerNumbers", rapid.MakeCheck(testNonIntegerNumbers))
	t.Run("TestMinimalStringEncoding", rapid.MakeCheck(testMinimalStringEncoding))

	t.Run("TestNormalizeNumber", rapid.MakeCheck(testNormalizeNumber))
	// t.Run("TestNormalizedString", rapid.MakeCheck(testNormalizedString))
}
