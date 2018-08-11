package main

import (
	"fmt"

	"github.com/champkeh/otp/hotp"
)

func main() {
	k := []byte("this is a very long secret")
	var c uint64 = 1213822223333

	value, err := hotp.GenerateHOTPValue(k, c, hotp.Param{Digits: 7})
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
