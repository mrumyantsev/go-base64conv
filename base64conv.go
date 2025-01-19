package base64conv

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"math/big"
)

var (
	ErrOutputNumberOverflowsTypeInt64 = errors.New("output number overflows type int64")
)

// ItobStd encodes int64 number to base64 string, using standard base64
// encoding, as defined in RFC 4648.
func ItobStd(i int64) string {
	return encode(base64.StdEncoding, i)
}

// ItobUrl encodes int64 number to base64 string, using alternate
// base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
func ItobUrl(i int64) string {
	return encode(base64.URLEncoding, i)
}

// ItobRawStd encodes int64 number to base64 string, using standard
// raw, unpadded base64 encoding, as defined in RFC 4648 section 3.2.
// This is the same as ItobStd but omits padding characters.
func ItobRawStd(i int64) string {
	return encode(base64.RawStdEncoding, i)
}

// ItobRawUrl encodes int64 number to base64 string, using unpadded
// alternate base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
// This is the same as ItobUrl but omits padding characters.
func ItobRawUrl(i int64) string {
	return encode(base64.RawURLEncoding, i)
}

// BtoiStd decodes int64 number from base64 string, using standard
// base64 encoding, as defined in RFC 4648.
func BtoiStd(b string) (int64, error) {
	return decode(base64.StdEncoding, b)
}

// BtoiUrl decodes int64 number from base64 string, using alternate
// base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
func BtoiUrl(b string) (int64, error) {
	return decode(base64.URLEncoding, b)
}

// BtoiRawStd decodes int64 number from base64 string, using standard
// raw, unpadded base64 encoding, as defined in RFC 4648 section 3.2.
// This is the same as BtoiStd but omits padding characters.
func BtoiRawStd(b string) (int64, error) {
	return decode(base64.RawStdEncoding, b)
}

// BtoiRawUrl decodes int64 number from base64 string, using unpadded
// alternate base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
// This is the same as BtoiUrl but omits padding characters.
func BtoiRawUrl(b string) (int64, error) {
	return decode(base64.RawURLEncoding, b)
}

// encode provides int64-to-base64 encoding with a given encoder from
// encoding/base64 package.
func encode(enc *base64.Encoding, i int64) string {
	bigint := big.NewInt(i)
	return enc.EncodeToString(bigint.Bytes())
}

// decode provides base64-to-int64 decoding with a given encoder from
// encoding/base64 package.
func decode(enc *base64.Encoding, b string) (int64, error) {
	i, err := enc.DecodeString(b)
	if err != nil {
		return 0, err
	}
	const int64bytes = 8
	length := len(i)
	if length > int64bytes {
		return 0, ErrOutputNumberOverflowsTypeInt64
	}
	zeros := make([]byte, int64bytes-length, int64bytes)
	return int64(binary.BigEndian.Uint64(append(zeros, i...))), nil
}
