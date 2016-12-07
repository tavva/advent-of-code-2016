package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
	"strings"
)

func main() {
	input := "ugkcyxxp"
	password := []byte("--------")

	for i := 0; i < int(^uint(0)>>1); i++ {
		hash := md5.New()
		io.WriteString(hash, input+strconv.Itoa(i))
		if hex.EncodeToString(hash.Sum(nil))[:5] == "00000" {
			index := string(hex.EncodeToString(hash.Sum(nil))[5])
			char := string(hex.EncodeToString(hash.Sum(nil))[6])

			index_int, err := strconv.ParseInt(index, 0, 64)
			if err != nil {
				continue
			}

			if index_int >= 0 && index_int <= 7 {
				println(index_int)
				println(char)
				if password[index_int] == '-' {
					password[index_int] = hex.EncodeToString(hash.Sum(nil))[6]
				}
			}

			println(string(password))

			if !strings.Contains(string(password), "-") {
				break
			}
		}
	}
}
