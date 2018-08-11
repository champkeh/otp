package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

type Param struct {
	Digits int
}

func GenerateHOTPValue(key []byte, counter uint64, param Param) (string, error) {
	if len(key) < 15 {
		return "", errors.New(fmt.Sprintf("secret must not less 120 bit (your key length:%d)", len(key)*8))
	}
	if param.Digits < 6 {
		return "", errors.New(fmt.Sprintf("digits must not less than 6 (your digits is %d)", param.Digits))
	}

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, counter)

	mac := hmac.New(sha1.New, key)
	mac.Write(buf)
	sum := mac.Sum(nil)

	offset := sum[19] & 0xf

	code := (int32(sum[offset]&0x7f) << 24) |
		(int32(sum[offset+1]&0xff) << 16) |
		(int32(sum[offset+2]&0xff) << 8) |
		int32(sum[offset+3]&0xff)

	format := fmt.Sprintf("%%0%dd", param.Digits)
	return fmt.Sprintf(format, code%int32(math.Pow10(param.Digits))), nil
}
