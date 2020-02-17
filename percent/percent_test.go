package percent

import (
	"net/url"
	"testing"
	"testing/quick"
)

// Random strings
func TestEncodeDecode(t *testing.T) {
	f := func(example string) bool {
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
	}
	if err := quick.Check(f, &quick.Config{
		MaxCount: 100,
	}); err != nil {
		t.Error(err)
	}
}

/*
// Random byte sequences, converted to strings.
func TestEncodeDecodeBytes(t *testing.T) {
	f := func(b []byte) bool {
		example := string(b)
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
	}
	if err := quick.Check(f, &quick.Config{
		MaxCount: 100,
	}); err != nil {
		t.Error(err)
	}
}
*/

/*
// Random strings, compared with the real QueryEscape implementation.
func TestEncodeVsEncode(t *testing.T) {
	f := func(example string) bool {
		t.Logf("Input:    %q", example)
		encoded, err := PercentEncode(example)
		if err != nil {
			return false
		}
		t.Logf("Encoded:  %q", encoded)
		encoded2 := url.QueryEscape(example)
		t.Logf("Encoded2: %q", encoded2)
		return encoded == encoded2
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
*/
