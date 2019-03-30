package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var f = flag.String("f", "sha256", "cryptographic hash functions")

func main() {
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	switch *f {
	case "sha256":
		hash := sha256.Sum256(input.Bytes())
		fmt.Printf("%x\n", hash)
	case "sha384":
		hash := sha512.Sum384(input.Bytes())
		fmt.Printf("%x\n", hash)
	case "sha512":
		hash := sha512.Sum512(input.Bytes())
		fmt.Printf("%x\n", hash)
	}
}
