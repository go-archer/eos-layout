package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func New(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
