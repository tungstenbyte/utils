package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	var passord = "tech5"

	pwd := sha1.New()
	pwd.Write([]byte(passord))
	pwd.Write([]byte("hash_salt"))

	senha := fmt.Sprintf("%x", pwd.Sum(nil))

	fmt.Println(senha)

}
