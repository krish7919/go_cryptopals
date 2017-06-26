// go_cryptopals project main.go
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	decoded, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)

	const expectedOutput = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	encodedLen := base64.StdEncoding.EncodedLen(len(decoded))
	output := make([]byte, encodedLen)
	base64.StdEncoding.Encode(output, decoded)

	fmt.Printf("%s\n", output)
	fmt.Printf("%s\n", expectedOutput)

	if bytes.Equal(output, []byte(expectedOutput)) {
		fmt.Println("SUCCESS!")
	}
}
