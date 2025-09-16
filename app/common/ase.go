package common

import (
	"go-ssh-forward/app/common/Const"
	"go-ssh-forward/app/common/ase"
	"strings"
)

var aseKey = "1234567890123456"

func init() {
	key, err := Const.GetMachineid()
	if err != nil {
		return
	}
	key = strings.ReplaceAll(key, "-", "")
	num := len(key)
	switch true {
	case num >= 32:
		aseKey = key[0:32]
		return
	case num < 32 && num >= 24:
		aseKey = key[0:24]
		return
	case num < 24 && num >= 16:
		aseKey = key[0:16]
		return
	default:
		return
	}
}

// AES解码
func Decode(str string) string {
	res, err := ase.AesDecryptGCM([]byte(aseKey), []byte(str))
	if err != nil {
		return ""
	}
	return string(res)
}

// AES编码
func Encode(str string) string {
	res, err := ase.AesEncryptGCM([]byte(aseKey), []byte(str))
	if err != nil {
		return ""
	}
	return string(res)
}
