// go_cryptopals project main.go
package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const str1 = "1c0111001f010100061a024b53535009181c"
	const str2 = "686974207468652062756c6c277320657965"
	const expectedStr = "746865206b696420646f6e277420706c6179"

	decoded1, err := hex.DecodeString(str1)
	if err != nil {
		log.Fatal(err)
	}
	decoded2, err := hex.DecodeString(str2)
	if err != nil {
		log.Fatal(err)
	}
	expected, err := hex.DecodeString(expectedStr)
	if err != nil {
		log.Fatal(err)
	}

	output := make([]byte, len(decoded1))
	for i := 0; i < len(decoded1); i++ {
		output[i] = decoded1[i] ^ decoded2[i]
	}

	if bytes.Equal(output, expected) {
		fmt.Println("SUCCESS!")
	}
}
