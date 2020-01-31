package util

import (
	"crypto/md5"
	"encoding/hex"
)

func StrMd5(str string) string {
	maker := md5.New()
	maker.Write([]byte(str))
	return hex.EncodeToString(maker.Sum(nil))
}
