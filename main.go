package main

import (
	"fmt"

	"github.com/liuyh73/DES-Golang-Binary/des"
)

func main() {
	key := "1101001010110101100110110111011110011001111010001101111001110110"
	clear_text := "1111111011011100101110101001100001110110010101000011001000010000"
	cipher_text := des.Encrypt(clear_text, key)
	fmt.Println("密钥：     ", key)
	fmt.Println("初始明文： ", clear_text)

	fmt.Println("明文加密后:", cipher_text)
	fmt.Println("密文解密后:", des.Decrypt(cipher_text, key))
}
