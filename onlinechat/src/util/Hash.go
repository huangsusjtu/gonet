package util

import (
	"crypto/md5"
)

// MD5
func HashMd5(str string) string {
	var hashMD5 = md5.New()
	hashMD5.Write([]byte(str))
	return string(hashMD5.Sum(nil))
}

func HashMd5Raw(raw []byte) []byte {
	var hashMD5 = md5.New()
	hashMD5.Write(raw)
	return hashMD5.Sum(nil)
}
