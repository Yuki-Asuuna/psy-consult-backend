package helper

import (
	"crypto/md5"
	"fmt"
)

func S2MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
