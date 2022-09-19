package random

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-08

import (
	"crypto/rand"
	"encoding/hex"
)

// RandString follow the same format as buffalo
func RandString(i int) string {
	if i == 0 {
		i = 64
	}
	b := make([]byte, i)
	rand.Read(b)
	return hex.EncodeToString(b)
}
