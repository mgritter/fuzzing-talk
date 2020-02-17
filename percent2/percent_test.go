package percent2

import (
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"net/url"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("encode can be unescaped", prop.ForAll(
		// Property should be true
		func(example string) bool {
			t.Logf("Input:   %q", example)
			encoded, err := PercentEncode(example)
			if err != nil {
				return false
			}
			t.Logf("Encoded: %q", encoded)
			decoded, err := url.QueryUnescape(encoded)
			if err != nil {
				return false
			}
			t.Logf("Decoded: %q", decoded)
			return example == decoded
		},
		gen.AnyString(),
	))

	properties.TestingRun(t)
}

/*
func TestEncodeDecodeBytes(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("bytes encoded can be unescaped", prop.ForAll(
		// Property should be true
		func(eb []byte) bool {
			example := string(eb)

			t.Logf("Input:   %q", example)
			encoded, err := PercentEncode(example)
			if err != nil {
				return false
			}
			t.Logf("Encoded: %q", encoded)
			decoded, err := url.QueryUnescape(encoded)
			if err != nil {
				return false
			}
			t.Logf("Decoded: %q", decoded)
			return example == decoded
		},
		gen.SliceOf(gen.UInt8()),
	))

	properties.TestingRun(t)
}
*/

/*
func TestEncodeVsEncode(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("my encode matches QueryEscape", prop.ForAll(
		func(example string) bool {
			//t.Logf("Input:    %q", example)
			encoded, err := PercentEncode(example)
			if err != nil {
				return false
			}
			//t.Logf("Encoded:  %q", encoded)
			encoded2 := url.QueryEscape(example)
			//t.Logf("Encoded2: %q", encoded2)
			return encoded == encoded2
		},
		gen.AnyString(),
	))

	properties.TestingRun(t)
}
*/

/*
func TestEncodeDecodeFilter(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("string encoded can be unescaped", prop.ForAll(
		// Property should be true
		func(example string) bool {
			t.Logf("Input:   %q", example)
			encoded, err := PercentEncode(example)
			if err != nil {
				return false
			}
			t.Logf("Encoded: %q", encoded)
			decoded, err := url.QueryUnescape(encoded)
			if err != nil {
				return false
			}
			t.Logf("Decoded: %q", decoded)
			return example == decoded
		},
		genAsciiString(),
	))

	properties.TestingRun(t)
}
*/

func genAsciiString() gopter.Gen {
	runeGen := gen.RuneRange('\x00', '\xff')
	return gen.SliceOf(runeGen).Map(runesToString).WithShrinker(gen.StringShrinker)
}

func runesToString(v []rune) string {
	return string(v)
}
