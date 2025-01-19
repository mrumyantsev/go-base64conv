package base64conv_test

import (
	"math"
	"testing"

	"github.com/mrumyantsev/base64conv-go"
)

func TestAll(t *testing.T) {
	testEncoderAndDecoder(t, base64conv.ItobStd, base64conv.BtoiStd, "Std")
	testEncoderAndDecoder(t, base64conv.ItobUrl, base64conv.BtoiUrl, "Url")
	testEncoderAndDecoder(t, base64conv.ItobRawStd, base64conv.BtoiRawStd, "RawStd")
	testEncoderAndDecoder(t, base64conv.ItobRawUrl, base64conv.BtoiRawUrl, "RawUrl")
}

func testEncoderAndDecoder(
	t *testing.T,
	enc func(i int64) string,
	dec func(b string) (int64, error),
	names string,
) {
	const areaRange = 1_000
	const begMin = 0
	const begMax = areaRange
	const midMin = (math.MaxInt64 >> 1) - (areaRange >> 1)
	const midMax = (math.MaxInt64 >> 1) + (areaRange >> 1)
	const endMin = math.MaxInt64 - areaRange
	const endMax = math.MaxInt64 - 1

	testInAreas(t, enc, dec, names, begMin, begMax)
	testInAreas(t, enc, dec, names, midMin, midMax)
	testInAreas(t, enc, dec, names, endMin, endMax)
}

func testInAreas(
	t *testing.T,
	enc func(i int64) string,
	dec func(b string) (int64, error),
	names string,
	min int64,
	max int64,
) {
	var current int64
	var restored int64
	var err error

	for current = min; current <= max; current++ {
		restored, err = dec(enc(current))
		if err != nil {
			t.Fatalf("unexpected error occured (%s): current: %d, error: %v", names, current, err)
		}

		if current != restored {
			t.Fatalf("current not equal restored (%s): current: %d, restored: %d", names, current, restored)
		}
	}
}
