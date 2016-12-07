package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
)

func main() {
	input := "ugkcyxxp"
	password := ""

	for i := 0; i < int(^uint(0)>>1); i++ {
		hash := md5.New()
		io.WriteString(hash, input+strconv.Itoa(i))
		if hex.EncodeToString(hash.Sum(nil))[:5] == "00000" {
			password += string(hex.EncodeToString(hash.Sum(nil))[5])
			if len(password) == 8 {
				break
			}
		}
	}

	println(password)
}
